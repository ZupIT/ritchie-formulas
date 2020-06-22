package main

import (
	"env/pkg/env"
	"os"
)

func main() {
	token := os.Getenv("TOKEN")
	repoOwner := os.Getenv("REPO_OWNER")
	repoName := os.Getenv("REPO_NAME")
	envName := os.Getenv("ENV_NAME")
	envValue := os.Getenv("ENV_VALUE")

	env.Input{
    	Token:    token,
    	RepoOwner:    repoOwner,
    	RepoName: repoName,
    	ENVName: envName,
    	ENVValue: envValue,
    }.Run()
}