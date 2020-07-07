package main

import (
	listTopic "github.com/ZupIT/ritchie-formulas/kafka/src/pkg/list/topic"
	"os"
)

func main() {
	if command := commandResolver(); command != nil {
		command.Run()
	}
}

func commandResolver() CommandHandler {
	u := os.Getenv("URLS")
	return listTopic.NewInputs(u)
}

type CommandHandler interface {
	Run()
}
