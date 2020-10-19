package main

import (
	"os"
	"ritman/pkg/ritman"
)

func main() {
	ritman.Input{
		Duration:   os.Getenv("TEST_DURATION"),
		MaxThreads: os.Getenv("MAX_THREADS"),
	}.Run()
}
