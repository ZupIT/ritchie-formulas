package main

import (
	"aws/pkg/aws"
	"os"
)

func main() {
	name := os.Getenv("PROJECT_NAME")
	repo := os.Getenv("PROJECT_REPO")

	aws.Input{
		ProjectName: name,
		ProjectPath: repo,
	}.Run()
}
