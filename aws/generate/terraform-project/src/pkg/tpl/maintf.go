package tpl

const (
	Maintf = `variable "region" {
  default = ""
}
provider "aws" {
  region  = var.region
  version = "2.57.0"
}

terraform {
  required_version = "0.12.13"
  required_providers {
    aws        = "2.57.0"
  }

  backend "s3" {
  }
}
	`
)
