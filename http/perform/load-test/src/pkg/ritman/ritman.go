package ritman

import (
	"crypto/tls"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/user"
	"path"
	"runtime"
	"strconv"
	"sync"
	"time"

	"github.com/fatih/color"
)

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

// Input contaning command data
type Input struct {
	Duration   string
	MaxThreads string
}

type testParameters struct {
	duration   time.Duration
	maxThreads int
	homePath   string
	score      LoadBalanceTestScore
}

// Result of load balance test
type Result struct {
	Started    int64
	Success    bool  `json:"success"`
	StatusCode int   `json:"statusCode"`
	Latency    int64 `json:"ms"`
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

func getInputParams(in Input) (int, int) {
	dur, err := strconv.Atoi(in.Duration)

	if err != nil {
		color.Red(fmt.Sprintf("Duration parameter is a integer field, error: %s", err.Error()))
		// EINVAL
		os.Exit(22)
	}

	maxThreads, err := strconv.Atoi(in.MaxThreads)

	if err != nil {
		color.Red(fmt.Sprintf("Max threads parameter is a integer field, error: %s", err.Error()))
		// EINVAL
		os.Exit(22)
	}

	return dur, maxThreads
}

func getHomeDir() string {
	usr, err := user.Current()

	if err != nil {
		color.Red("Could not get your current user path")
		// EACCES
		os.Exit(13)
	}

	return usr.HomeDir
}

func getTestParams(in Input) testParameters {
	dur, maxThreads := getInputParams(in)

	return testParameters{
		duration:   time.Duration(dur) * time.Second,
		maxThreads: maxThreads,
		homePath:   getHomeDir(),
		score:      LoadBalanceTestScore{},
	}
}

// Run ritchie CLI integration
func (in Input) Run() {
	ritmanTest := getTestParams(in)
	ritman := NewRitman(ritmanTest.maxThreads)
	request := NewRequestTarget(path.Join(ritmanTest.homePath, ".rit", testTargetFileName))

	ritmanTest.score.init()
	ritmanTest.score.StartedAt = ritman.started.Format(time.RFC3339)

	fmt.Printf("Starting test with duration of %d seconds, Target: %s - %s\n", int64(ritmanTest.duration.Seconds()), request.Method, request.Target)

	for res := range ritman.LoadTesting(request, ritmanTest.duration) {
		ritmanTest.score.Add(res)
	}

	ritmanTest.score.measureLatencyAndRps()
	ritmanTest.score.writeToFile(ritmanTest.homePath)
}
