package ritman

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path"
	"strconv"
	"time"

	"github.com/fatih/color"
)

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
	score.Hits++
	score.updateLatency(*res)
	key := time.Unix(res.Started, 0).Second()

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

func (score LoadBalanceTestScore) writeToFile(home string) {
	path := path.Join(home, getResultFileName())
	jsonData, err := json.MarshalIndent(score, "", "\t")

	if err == nil {
		err = ioutil.WriteFile(path, jsonData, 0644)
		if err != nil {
			panic(fmt.Sprintf("Could not write results file: %s\nPermission required +644", err))
		} else {
			color.Green(fmt.Sprintf("Test results wrote in file: %s\n", path))
		}
	} else {
		panic(fmt.Sprintf("Could not serialize test results: %s", err))
	}
}

func getResultFileName() string {
	return "ritman-test-result-" + strconv.FormatInt(time.Now().Unix(), 10) + ".json"
}
