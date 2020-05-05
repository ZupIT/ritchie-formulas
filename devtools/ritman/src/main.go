package main

import (
	"log"
	"os"
	"ritman/pkg/ritman"
	"ritman/pkg/template"
)

const (
	ritmanCmd   = "httpTest"
	templateCmd = "template"
)

func main() {
	loadInputs().Run()
}

func loadInputs() commandHandler {
	command := os.Getenv("COMMAND")
	switch command {
	case ritmanCmd:
		return ritman.Input{
			Duration:   os.Getenv("test_duration"),
			MaxThreads: os.Getenv("max_threads"),
		}
	case templateCmd:
		return template.Input{}
	default:
		log.Fatalln("Command not found")
	}
	return nil
}

type commandHandler interface {
	Run()
}
