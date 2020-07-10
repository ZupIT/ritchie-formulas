package main

import (
	"aws/bucket/pkg/bucket"
	"os"
)

func main() {
	loadInputs().Run()
}

func loadInputs() bucket.Inputs {
	return bucket.Inputs{
		Key:    os.Getenv("ACCESS_KEY"),
		Secret: os.Getenv("SECRET_ACCESS_KEY"),
		Region: os.Getenv("REGION"),
	}
}
