package main

import (
	"aws/pkg/aws"
	"os"
)

func main() {
	name := os.Getenv("PROJECT_NAME")
	loc := os.Getenv("PROJECT_LOCATION")

	aws.Input{
		ProjectName:     name,
		ProjectLocation: loc,
	}.Run()
}
