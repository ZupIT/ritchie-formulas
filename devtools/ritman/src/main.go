package main

import (
	"os"
	"ritman/pkg/ritman"
)

func main() {
	duration := os.Getenv("test_duration")
	url := os.Getenv("test_url")

	ritman.Input{
		Duration: duration,
		Url:      url,
	}.Run()
}
