package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/ZupIT/ritchie-formulas/kafka/src/pkg/consume"
	createTopic "github.com/ZupIT/ritchie-formulas/kafka/src/pkg/create/topic"
	listTopic "github.com/ZupIT/ritchie-formulas/kafka/src/pkg/list/topic"
	"github.com/ZupIT/ritchie-formulas/kafka/src/pkg/produce"
)

const (
	produceCmd     = "produce"
	consumeCmd     = "consume"
	listTopicCmd   = "listTopic"
	createTopicCmd = "createTopic"
)

func main() {
	if command := commandResolver(); command != nil {
		command.Run()
	}
}

func commandResolver() CommandHandler {
	cmd := os.Getenv("COMMAND")
	u := os.Getenv("URLS")

	switch cmd {
	case consumeCmd:
		b, _ := strconv.ParseBool(os.Getenv("BEGINNING"))
		t := os.Getenv("TOPIC")

		return consume.NewInputs(u, t, b)
	case produceCmd:
		t := os.Getenv("TOPIC")

		return produce.NewInputs(u, t)
	case createTopicCmd:
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
	case listTopicCmd:
		return listTopic.NewInputs(u)
	default:
		fmt.Println("Command not found")
		return nil
	}
}

type CommandHandler interface {
	Run()
}
