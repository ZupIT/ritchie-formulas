package main

import (
	"os"
	"project/pkg/project"
)

func main() {
	name := os.Getenv("PROJECT_NAME")
	repo := os.Getenv("PROJECT_REPO")

	project.Input{
		ProjectName: name,
		ProjectPath: repo,
	}.Run()
}
