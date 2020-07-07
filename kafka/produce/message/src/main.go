package main

import (
	"github.com/ZupIT/ritchie-formulas/kafka/src/pkg/produce"
	"os"
)

func main() {
	if command := commandResolver(); command != nil {
		command.Run()
	}
}

func commandResolver() CommandHandler {
	u := os.Getenv("URLS")

	t := os.Getenv("TOPIC")

	return produce.NewInputs(u, t)
}

type CommandHandler interface {
	Run()
}
