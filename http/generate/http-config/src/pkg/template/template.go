package template

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/user"
	"path"

	"github.com/fatih/color"
)

const (
	defaultMethod    = "POST"
	defaultURLTarget = "https://postman-echo.com/post"
	configFileName   = "ritman-target.json"
)

var (
	defaultHeaders = map[string]string{
		"Accept":          "application/json; charset=utf-8",
		"Accept-Encoding": "gzip, deflate, br",
		"Connection":      "Keep-alive",
		"Content-Type":    "application/json",
	}
	defaultBody = map[string]string{
		"name":     "ritchie",
		"lastName": "cli",
		"command":  "ritman",
		"url":      "https://ritchiecli.io/",
	}
)

type fileWriter struct {
	fileContent []byte
}

type requestTarget struct {
	Target  string            `json:"target"`
	Method  string            `json:"method"`
	Headers map[string]string `json:"headers"`
	Body    map[string]string `json:"body"`
}

func (fw fileWriter) getContent() fileWriter {
	data := requestTarget{Target: defaultURLTarget, Method: defaultMethod, Headers: defaultHeaders, Body: defaultBody}

	if file, err := json.MarshalIndent(data, "", " "); err == nil {
		fw.fileContent = file
	} else {
		color.Red(fmt.Sprintf("could not create config file: %s\n", err))
		// EACCES
		os.Exit(13)
	}

	return fw
}

func (fw fileWriter) writeToFile() {
	usr, err := user.Current()

	if err != nil {
		color.Red("Could not get your current user path")
		// EACCES
		os.Exit(13)
	}

	path := path.Join(usr.HomeDir, ".rit", configFileName)

	if err := ioutil.WriteFile(path, fw.fileContent, 0600); err == nil {
		color.Green(fmt.Sprintf("config file created: %s\n", path))
	} else {
		color.Red(fmt.Sprintf("could not write config file: %s\n", err))
	}
}

// Run - Ritchie formula runner
func Run() {
	var writer fileWriter

	writer.getContent().writeToFile()
}
