package template

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/fatih/color"
)

type Input struct {
}

type fileWriter struct {
	fileContent []byte
}

type requestTarget struct {
	Target  string            `json:"target"`
	Method  string            `json:"method"`
	Headers map[string]string `json:"headers"`
	Body    map[string]string `json:"body"`
}

const defaultRitmanConfigFile = "ritman-target.json"

const defaultURLTarget = "https://postman-echo.com/post"
const defaultMethod = "POST"

var defaultHeaders = map[string]string{
	"Accept":          "application/json; charset=utf-8",
	"Accept-Encoding": "gzip, deflate, br",
	"Connection":      "Keep-alive",
	"Content-Type":    "application/json",
}

var defaultBody = map[string]string{
	"name":     "ritchie",
	"lastName": "cli",
	"command":  "ritman",
	"url":      "https://ritchiecli.io/",
}

func (fw fileWriter) getContent() fileWriter {
	data := requestTarget{Target: defaultURLTarget, Method: defaultMethod, Headers: defaultHeaders, Body: defaultBody}

	if file, err := json.MarshalIndent(data, "", " "); err == nil {
		fw.fileContent = file
	} else {
		color.Red(fmt.Sprintf("could not create config file: %s\n", err))
	}

	return fw
}

func (fw fileWriter) writeToFile() {
	if err := ioutil.WriteFile(defaultRitmanConfigFile, fw.fileContent, 0644); err == nil {
		mydir, _ := os.Getwd()
		path := filepath.FromSlash(mydir + "/" + defaultRitmanConfigFile)
		color.Green(fmt.Sprintf("config file created: %s\n", path))
	} else {
		color.Red(fmt.Sprintf("could not write config file: %s\n", err))
	}
}

func (in Input) Run() {
	var writer fileWriter

	writer.getContent().writeToFile()
}
