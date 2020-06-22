package tpl

const (
	Maintf = `variable "region" {
  default = ""
}
provider "aws" {
  region  = var.region
  version = "~> 2.8"
}

terraform {
  required_version = "0.12.13"
  required_providers {
    aws        = "~> 2.8"
  }

  backend "s3" {
  }
}
	`
)
