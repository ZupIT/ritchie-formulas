package main

import (
	"os"
	"strconv"

	"github.com/ZupIT/ritchie-formulas/kafka/src/pkg/consume"
)

func main() {
	if command := commandResolver(); command != nil {
		command.Run()
	}
}

func commandResolver() CommandHandler {
	u := os.Getenv("URLS")
	b, _ := strconv.ParseBool(os.Getenv("BEGINNING"))
	t := os.Getenv("TOPIC")

	return consume.NewInputs(u, t, b)
}

type CommandHandler interface {
	Run()
}
