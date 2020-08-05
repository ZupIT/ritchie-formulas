package main

import (
	"os"
	"terraform/pkg/terraform"
)

func main() {
	loadInputs().Run()
}

func loadInputs() terraform.Inputs {
	return terraform.Inputs{
		Repository:         os.Getenv("REPOSITORY"),
		TerraformPath:      os.Getenv("TERRAFORM_PATH"),
		Environment:        os.Getenv("ENVIRONMENT"),
		GitUser:            os.Getenv("GIT_USER"),
		GitToken:           os.Getenv("GIT_TOKEN"),
		AwsAccessKeyId:     os.Getenv("AWS_ACCESS_KEY_ID"),
		AwsSecretAccessKey: os.Getenv("AWS_SECRET_ACCESS_KEY"),
		Pwd:                os.Getenv("CURRENT_PWD"),
	}
}
