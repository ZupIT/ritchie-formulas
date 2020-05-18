package ritman

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"sync"
	"time"

	"github.com/fatih/color"
)

// Input contaning command data
type Input struct {
	Duration   string
	MaxThreads string
}

const testTargetFileName = "ritman-target.json"

var (
	// Using Brian Goetz thread pool formula, assuming "wait time" (LAN service) to
	// have a latency of 20ms and "Service time" latency of 5ms
	// Number of threads = Number of Available Cores * (1 + Wait time / Service time)
	defaultWorkers = runtime.NumCPU() * (1 + (20 / 5))
	localAddr      = net.IPAddr{IP: net.IPv4zero}
	dialer         = &net.Dialer{
		LocalAddr: &net.TCPAddr{IP: localAddr.IP, Zone: localAddr.Zone},
		KeepAlive: 30 * time.Second,
	}
)

//RequestTarget JSON Request target
type RequestTarget struct {
	Target  string                 `json:"target"`
	Method  string                 `json:"method"`
	Headers map[string]string      `json:"headers"`
	Body    map[string]interface{} `json:"body"`
}

// Result of load balance test
type Result struct {
	Started    int64
	Success    bool  `json:"success"`
	StatusCode int   `json:"statusCode"`
	Latency    int64 `json:"ms"`
}

// LoadBalanceTestScore computed test data
type LoadBalanceTestScore struct {
	Hits          int64             `json:"hits"`
	RatePerSecond float64           `json:"avgRps"`
	AvgMs         int64             `json:"avgMs"`
	MaxMs         int64             `json:"maxMs"`
	MinMs         int64             `json:"minMs"`
	StartedAt     string            `json:"startedAt"`
	EndAt         string            `json:"endedAt"`
	Histogram     map[int]Histogram `json:"histogram"`
}

// Histogram to organize data in a chart structure
type Histogram struct {
	Hits       int64       `json:"hits"`
	AvgMs      int64       `json:"avgMs"`
	MaxMs      int64       `json:"maxMs"`
	MinMs      int64       `json:"minMs"`
	StatusCode map[int]int `json:"statusCode"`
}

// Ritman client for http load testing
type Ritman struct {
	client     http.Client
	started    time.Time
	maxThreads int
}

// LoadTesting creates the worker pool to test an givin service performance
func (r *Ritman) LoadTesting(request RequestTarget, duration time.Duration) <-chan *Result {
	var workers int
	var wg sync.WaitGroup

	results := make(chan *Result)

	if r.maxThreads > 0 {
		workers = r.maxThreads
	} else {
		workers = defaultWorkers
	}

	fmt.Printf("Spawning %d async workers\n", workers)

	for i := 0; i < workers; i++ {
		wg.Add(1)
		go r.handle(request, &wg, duration, results)
	}

	go func() {
		defer close(results)
		defer wg.Wait()

		for time.Since(r.started).Milliseconds() <= duration.Milliseconds() {
			time.Sleep(1 * time.Second)
			fmt.Printf(".")
		}
		fmt.Printf("\n")
		fmt.Println("Done! Collecting results...")
	}()

	return results
}

func (r *Ritman) handle(request RequestTarget, wg *sync.WaitGroup, duration time.Duration, ch chan<- *Result) {
	defer wg.Done()

	for time.Since(r.started).Milliseconds() <= duration.Milliseconds() {
		ch <- r.doRequest(request)
	}
}

func (r *Ritman) doRequest(request RequestTarget) *Result {
	started := time.Now()
	var res = Result{}

	resp, err := r.client.Do(request.createRequest())

	res.Latency = time.Since(started).Milliseconds()
	res.Started = started.Unix()
	res.Success = err == nil

	if err == nil {
		res.StatusCode = resp.StatusCode
	} else {
		// Probably a socket timeout/socket err
		res.StatusCode = -1
	}

	return &res
}

// NewRitman returns an default Ritman which will be used to test a given target
// Http.Transport.MaxConnsPerHost = 0 - Not limit, the OS will manage the socket buffer size
func NewRitman(threads int) *Ritman {
	var httpClient = http.Client{
		Timeout: 60 * time.Second,
		Transport: &http.Transport{
			Proxy:               http.ProxyFromEnvironment,
			Dial:                dialer.Dial,
			TLSClientConfig:     &tls.Config{InsecureSkipVerify: true},
			MaxIdleConnsPerHost: 10000,
			MaxConnsPerHost:     0,
		},
	}

	return &Ritman{client: httpClient, started: time.Now(), maxThreads: threads}
}

func newHistogram(res Result) Histogram {
	statusCode := make(map[int]int)

	statusCode[res.StatusCode]++

	return Histogram{Hits: 1, AvgMs: res.Latency, MaxMs: res.Latency, MinMs: res.Latency, StatusCode: statusCode}
}

func (score *LoadBalanceTestScore) init() {
	if score.Histogram == nil {
		score.Histogram = make(map[int]Histogram)
	}
	// Starts with big number just to avoid always 0 as minimum latency
	score.MinMs = 99999999
}

// Add helper function to compute data
func (score *LoadBalanceTestScore) Add(res *Result) {
	key := time.Unix(res.Started, 0).Second()

	score.Hits++

	score.updateLatency(*res)
	if histogram, ok := score.Histogram[key]; ok {
		histogram.Hits++

		histogram.StatusCode[res.StatusCode]++

		histogram.updateLatency(*res)

		score.Histogram[key] = histogram
	} else {
		score.Histogram[key] = newHistogram(*res)
	}
}

func (score *LoadBalanceTestScore) updateLatency(res Result) {
	score.AvgMs += res.Latency

	if score.MaxMs < res.Latency {
		score.MaxMs = res.Latency
	}
	if score.MinMs > res.Latency {
		score.MinMs = res.Latency
	}
}

func (h *Histogram) updateLatency(res Result) {
	h.AvgMs += res.Latency

	if h.MaxMs < res.Latency {
		h.MaxMs = res.Latency
	}
	if h.MinMs > res.Latency {
		h.MinMs = res.Latency
	}
}

func (score *LoadBalanceTestScore) measureLatencyAndRps() {

	score.EndAt = time.Now().Format(time.RFC3339)

	if len(score.Histogram) > 0 {
		score.RatePerSecond = float64(score.Hits / int64(len(score.Histogram)))
		score.AvgMs = score.AvgMs / score.Hits

		for key, value := range score.Histogram {
			avgMs := value.AvgMs / value.Hits
			histogram := score.Histogram[key]
			histogram.AvgMs = avgMs
			score.Histogram[key] = histogram
		}
	}
}

func (score LoadBalanceTestScore) toJSON() (string, error) {
	jsonData, err := json.Marshal(score)

	var prettyJSON bytes.Buffer

	json.Indent(&prettyJSON, jsonData, "", "\t")

	if err == nil {
		return string(prettyJSON.Bytes()), nil
	}

	return "", err
}

func (score LoadBalanceTestScore) writeToFile() {
	fileName := getResultFileName()
	jsonData, err := json.MarshalIndent(score, "", "\t")

	if err == nil {
		err = ioutil.WriteFile(fileName, jsonData, 0644)
		if err != nil {
			panic(fmt.Sprintf("Could not write results file: %s\nPermission required +644", err))
		} else {
			printResultPath(fileName)
		}
	} else {
		panic(fmt.Sprintf("Could not serialize test results: %s", err))
	}
}

func printResultPath(fileName string) {
	mydir, err := os.Getwd()

	if err == nil {
		path := filepath.FromSlash(mydir + "/" + fileName)
		color.Green(fmt.Sprintf("Test results wrote in file: %s\n", path))
	} else {
		panic(fmt.Sprintf("Could not get the results path: %s", err.Error()))
	}
}

func getResultFileName() string {
	return "ritman-test-result-" + strconv.FormatInt(time.Now().Unix(), 10) + ".json"
}

func (request RequestTarget) createRequest() *http.Request {
	req, err := http.NewRequest(request.Method, request.Target, request.getBody())

	if err != nil {
		panic(fmt.Sprintf("Could not create HTTP Request: %s", err.Error()))
	}

	for key, value := range request.Headers {
		req.Header.Add(key, value)
	}

	req.Header.Add("User-Agent", "RitmanLoadTesterRuntime/1.0.0-rc")

	return req
}

func (request RequestTarget) getBody() *bytes.Reader {
	var (
		body []byte
		err  error
	)

	switch request.Method {
	case
		"PUT",
		"POST",
		"PATCH":
		body, err = json.Marshal(request.Body)
		if err != nil {
			panic(fmt.Sprintf("Error adding body: %s", err.Error()))
		}
	}

	return bytes.NewReader(body)
}

// NewRequestTarget - Creates a new RequestTarget struct from test-target.json file
func NewRequestTarget() RequestTarget {
	var request RequestTarget
	file, err := ioutil.ReadFile(testTargetFileName)

	if err != nil {
		color.Red(fmt.Sprintf("Could not find ritman config file, please run command: rit devtools template-generator and make sure the HTTP target, method, headers and body matches your test criteria."))
		// ENOENT - No such file or directory
		os.Exit(2)
	}

	if err := json.Unmarshal([]byte(file), &request); err != nil {
		panic(err)
	}

	return request
}

// Run ritchie CLI integration
func (in Input) Run() {
	dur, err := strconv.Atoi(in.Duration)

	if err != nil {
		panic(fmt.Sprintf("Duration parameter is a integer field, error: %s", err.Error()))
	}

	maxThreads, err := strconv.Atoi(in.MaxThreads)

	if err != nil {
		panic(fmt.Sprintf("Max threads parameter is a integer field, error: %s", err.Error()))
	}

	ritman := NewRitman(maxThreads)
	request := NewRequestTarget()
	duration := time.Duration(dur) * time.Second

	var score LoadBalanceTestScore
	score.init()
	score.StartedAt = ritman.started.Format(time.RFC3339)

	fmt.Printf("Starting test with duration of %d seconds, Target: %s - %s\n", int64(duration.Seconds()), request.Method, request.Target)

	for res := range ritman.LoadTesting(request, duration) {
		score.Add(res)
	}

	score.measureLatencyAndRps()

	score.writeToFile()
}
