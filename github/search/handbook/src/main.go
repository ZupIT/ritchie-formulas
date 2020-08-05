package main

import (
	"handbook/pkg/handbook"
	"os"
)

func main() {
	loadInputs().Run()
}

func loadInputs() handbook.Inputs {
	return handbook.Inputs{
		GitUser:  os.Getenv("GIT_USER"),
		GitToken: os.Getenv("GIT_TOKEN"),
	}
}
