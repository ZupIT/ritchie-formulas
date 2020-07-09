package env

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/fatih/color"
)

const (
	URL = "https://circleci.com/api/v1.1/project/gh/%s/%s/envvar?circle-token=%s"
)

type Input struct {
	Token     string
	RepoOwner string
	RepoName  string
	ENVName   string
	ENVValue  string
}

func (in Input) Run() {
	env := struct {
		Name  string `json:"name"`
		Value string `json:"value"`
	}{
		in.ENVName,
		in.ENVValue,
	}

	b, err := json.Marshal(&env)
	if err != nil {
		color.Red(fmt.Sprintf("error marshalling env: %q, error: %q", env, err))
		os.Exit(1)
	}

	url := fmt.Sprintf(URL, in.RepoOwner, in.RepoName, in.Token)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(b))
	if err != nil {
		color.Red(fmt.Sprintf("error sending env: %q", err))
		os.Exit(1)
	}

	if resp.StatusCode != http.StatusCreated {
		defer resp.Body.Close()
		b, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			color.Red(fmt.Sprintf("error reading response body: %q", err))
			os.Exit(1)
		}
		msg := fmt.Sprintf("%s - %s", resp.Status, string(b))
		color.Red(fmt.Sprintf("error sending env: %q", msg))
		os.Exit(1)
	}

	color.Green(fmt.Sprintf("env (%s) added successfully!\n", in.ENVName))
	color.Green(fmt.Sprintf("project: %s/%s\n", in.RepoOwner, in.RepoName))
}
