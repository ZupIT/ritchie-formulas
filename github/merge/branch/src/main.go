package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func main() {
	destinationBranch := os.Getenv("BRANCH")
	push := os.Getenv("PUSH")
	pwdEnv := os.Getenv("CURRENT_PWD")
	err := os.Chdir(pwdEnv)
	if err != nil {
		log.Fatalf("Failed to execute formula. Erro: %v", err)
	}
	currentBranch := execCommand("git rev-parse --abbrev-ref HEAD")

	execCommand("git pull origin " + currentBranch)
	execCommand("git fetch")
	execCommand("git branch -D " + destinationBranch)
	execCommand("git checkout " + destinationBranch)
	execCommand("git pull origin " + currentBranch)

	if push == "true" {
		execCommand("git push")
		execCommand("git checkout " + currentBranch)
	}
}

func execCommand(value string) string {
	command := strings.Split(value, " ")[0]
	params := strings.Split(value, " ")[1:]
	log.Printf("Executing command: %v params: %v\n", command, params)
	cmd := exec.Command(command, params...)
	stdout, _ := cmd.StdoutPipe()
	var outError bytes.Buffer
	cmd.Stderr = &outError
	errS := cmd.Start()
	if errS != nil {
		log.Fatalf("Error to start command")
	}
	scanner := bufio.NewScanner(stdout)
	scanner.Split(bufio.ScanLines)
	commandResultMessage := ""
	for scanner.Scan() {
		m := scanner.Text()
		fmt.Println(m)
		commandResultMessage += m
	}
	err := cmd.Wait()
	if err != nil {
		log.Fatalf("Failed to execute command %v\nParams: %v\nError: %v", command, params, outError.String())
	}
	return commandResultMessage
}
