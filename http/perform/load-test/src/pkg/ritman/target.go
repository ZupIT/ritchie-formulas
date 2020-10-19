package ritman

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/fatih/color"
)

//RequestTarget JSON Request target
type RequestTarget struct {
	Target  string                 `json:"target"`
	Method  string                 `json:"method"`
	Headers map[string]string      `json:"headers"`
	Body    map[string]interface{} `json:"body"`
}

func (request RequestTarget) createRequest() *http.Request {
	req, err := http.NewRequest(request.Method, request.Target, request.getBody())

	if err != nil {
		panic(fmt.Sprintf("Could not create HTTP Request: %s", err.Error()))
	}

	for key, value := range request.Headers {
		req.Header.Add(key, value)
	}

	req.Header.Add("User-Agent", "RitmanLoadTest/0.0.1")

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
func NewRequestTarget(filePath string) RequestTarget {
	var request RequestTarget
	file, err := ioutil.ReadFile(filePath)

	if err != nil {
		color.Red(fmt.Sprintf("Could not find ritman config file, please run command: (rit http generate http-config) and make sure the HTTP target, method, headers and body matches your test criteria."))
		// ENOENT - No such file or directory
		os.Exit(2)
	}

	if err := json.Unmarshal([]byte(file), &request); err != nil {
		panic(err)
	}

	return request
}
