package formula

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

type OutputJSON struct {
	Output map[string]interface{} `json:"output"`
}

var outputFile = RitchieDir() + "/output.json"

func RunFormula(command string) error {
	cmdSplit := strings.Split(command, " ")
	cmd := exec.Command(cmdSplit[0], cmdSplit[1:]...)
	if err := cmd.Start(); err != nil {
		return err
	}

	cmd.Wait()

	return nil
}

func WriteOutput(output map[string]interface{}) error {
	data := OutputJSON{output}
	outputJSON, err := json.Marshal(data)
	if err != nil {
		return err
	}

	if err := ioutil.WriteFile(outputFile, outputJSON, os.ModePerm); err != nil {
		return err
	}

	return nil
}

func ReadOutput() (OutputJSON, error) {
	file, _ := ioutil.ReadFile(outputFile)

	var data OutputJSON
	if err := json.Unmarshal(file, &data); err != nil {
		return data, err
	}

	return data, nil
}

func RemoveFile() error {
	if err := os.Remove(outputFile); err != nil {
		return err
	}

	return nil
}

func RitchieDir() string {
	myDir, _ := os.UserHomeDir()

	return filepath.Join(myDir, ".rit")
}
