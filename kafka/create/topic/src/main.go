package main

import (
	"fmt"
	"os"
	"strconv"

	createTopic "github.com/ZupIT/ritchie-formulas/kafka/src/pkg/create/topic"
)

func main() {
	if command := commandResolver(); command != nil {
		command.Run()
	}
}

func commandResolver() CommandHandler {
	u := os.Getenv("URLS")

	n := os.Getenv("NAME")
	r := os.Getenv("REPLICATION")
	p := os.Getenv("PARTITIONS")

	re, err := strconv.Atoi(r)
	if err != nil {
		fmt.Println("Replication must be a number")
		return nil
	}

	pa, err := strconv.Atoi(p)
	if err != nil {
		fmt.Println("Partitions must be a number")
		return nil
	}

	return createTopic.NewInputs(u, n, int16(re), int32(pa))
}

type CommandHandler interface {
	Run()
}
