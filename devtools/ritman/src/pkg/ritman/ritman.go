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
)

// Input contaning command data
type Input struct {
	Duration string
	URL      string
}

var (
	// Using Brian Goetz thread pool formula.
	// Assuming "wait time" to have a latency of 20ms and "Service time" latency of 5ms
	// Number of threads = Number of Available Cores * (1 + Wait time / Service time)
	defaultWorkers = runtime.NumCPU() * (1 + (20 / 5))
	localAddr      = net.IPAddr{IP: net.IPv4zero}
	dialer         = &net.Dialer{
		LocalAddr: &net.TCPAddr{IP: localAddr.IP, Zone: localAddr.Zone},
		KeepAlive: 30 * time.Second,
	}
)

// Result of load balance test
type Result struct {
	Started    int64
	Success    bool   `json:"success"`
	Body       string `json:"body"`
	StatusCode int    `json:"statusCode"`
	Latency    int64  `json:"ms"`
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
	Hits        int64       `json:"hits"`
	AvgMs       int64       `json:"avgMs"`
	MaxMs       int64       `json:"maxMs"`
	MinMs       int64       `json:"minMs"`
	StatusCodes map[int]int `json:"statusCodes"`
}

// Ritman client for http load testing
type Ritman struct {
	client  http.Client
	started time.Time
}

// LoadTesting creates the worker pool to test an given service performance
func (r *Ritman) LoadTesting(req *http.Request, duration time.Duration) <-chan *Result {
	var wg sync.WaitGroup

	results := make(chan *Result)

	for i := 0; i < defaultWorkers; i++ {
		wg.Add(1)
		go r.handle(req, &wg, duration, results)
	}

	go func() {
		defer close(results)
		defer wg.Wait()

		for time.Since(r.started).Milliseconds() <= duration.Milliseconds() {
			time.Sleep(1 * time.Second)
			fmt.Printf(".")
		}
		fmt.Println("Done! Collecting results...")
	}()

	return results
}

func (r *Ritman) handle(req *http.Request, wg *sync.WaitGroup, duration time.Duration, ch chan<- *Result) {
	defer wg.Done()

	for time.Since(r.started).Milliseconds() <= duration.Milliseconds() {
		ch <- r.doRequest(req)
	}
}

func (r *Ritman) doRequest(req *http.Request) *Result {
	started := time.Now()
	var res = Result{}

	resp, err := r.client.Do(req)

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
func NewRitman() *Ritman {
	var httpClient = http.Client{
		Timeout: 30 * time.Second,
		Transport: &http.Transport{
			Proxy:               http.ProxyFromEnvironment,
			Dial:                dialer.Dial,
			TLSClientConfig:     &tls.Config{InsecureSkipVerify: true},
			MaxIdleConns:        1024,
			MaxIdleConnsPerHost: 10000,
			MaxConnsPerHost:     0,
		},
	}

	return &Ritman{client: httpClient, started: time.Now()}
}

func newHistogram(res Result) Histogram {
	statusCodes := make(map[int]int)

	statusCodes[res.StatusCode] = 1

	return Histogram{Hits: 1, AvgMs: res.Latency, MaxMs: res.Latency, MinMs: res.Latency, StatusCodes: statusCodes}
}

func (score *LoadBalanceTestScore) init() {
	if score.Histogram == nil {
		score.Histogram = make(map[int]Histogram)
	}
	// Add a generic big number just to avoid always 0 as minimum latency
	score.MinMs = 99999999
}

// Add helper function to compute data
func (score *LoadBalanceTestScore) Add(res *Result) {
	key := time.Unix(res.Started, 0).Second()

	score.Hits++

	score.updateLatency(*res)
	if histogram, ok := score.Histogram[key]; ok {
		histogram.Hits++
		histogram.StatusCodes[res.StatusCode]++

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
		fmt.Println("Test results wrote in file: ", path)
	} else {
		panic(fmt.Sprintf("Could not get the results path: %s", err.Error()))
	}
}

func getResultFileName() string {
	return "ritman-test-result-" + strconv.FormatInt(time.Now().Unix(), 10) + ".json"
}

func createRequest(method string, url string) *http.Request {
	var body []byte

	req, err := http.NewRequest(method, url, bytes.NewReader(body))

	if err != nil {
		panic(fmt.Sprintf("Could not create HTTP Request: %s", err.Error()))
	}

	req.Header.Add("User-Agent", "RitmanLoadTesterRuntime/1.0.0b")
	req.Header.Add("Accept", "application/json; charset=utf-8")
	req.Header.Add("Accept-Encoding", "gzip, deflate, br")
	req.Header.Add("Connection", "Keep-alive")

	return req
}

// Ritman command runner
func (in Input) Run() {
	dur, err := strconv.Atoi(in.Duration)
	if err != nil {
		panic(fmt.Sprintf("Duration parameter is a integer field, error: %s", err.Error()))
	}

	if len(in.URL) < 15 {
		panic(fmt.Sprintf("Invalid URL %s", in.URL))
	}

	var method = "GET"
	duration := time.Duration(dur) * time.Second
	ritman := NewRitman()
	req := createRequest(method, in.URL)

	var score LoadBalanceTestScore
	score.init()
	score.StartedAt = ritman.started.Format(time.RFC3339)

	fmt.Printf("Starting test with duration of %d seconds, Target: %s - %s\n", int64(duration.Seconds()), method, in.URL)

	for res := range ritman.LoadTesting(req, duration) {
		score.Add(res)
	}

	score.measureLatencyAndRps()

	score.writeToFile()
}
