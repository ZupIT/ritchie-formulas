package main

import (
	"env/pkg/env"
	"os"
    "github.com/fatih/color"
)

func main() {
	token := os.Getenv("TOKEN")
	repoOwner := os.Getenv("REPO_OWNER")
	repoName := os.Getenv("REPO_NAME")
	envName := os.Getenv("ENV_NAME")
	envValue := os.Getenv("ENV_VALUE")

	color.Green(fmt.Sprintf("env (%s) added successfully!\n", envName))
	color.Green(fmt.Sprintf("project: %s/%s\n", repoOwner, repoName))
}
