package tpl

const (
	Maintf = `
variable "vpc_name" {
  type = string
}
variable "vpc_cidr" {
  type = string
}

variable "vpc_azs" {
  type = list(string)
}

variable "customer_name" {
  default = ""
}
module "vpc" {
  source = "terraform-aws-modules/vpc/aws"

  version = "~> v2.0"

  name = var.vpc_name
  cidr = var.vpc_cidr

  azs             = var.vpc_azs
  private_subnets = [
  for num in [1,2,3]:
    cidrsubnet(var.vpc_cidr, 5, num)
  ]
  public_subnets  = [
  for num in [4,5,6]:
  cidrsubnet(var.vpc_cidr, 5, num)
  ]

  enable_nat_gateway = true
  enable_vpn_gateway = false

  tags = {
    Terraform = "true"
    Environment = var.customer_name
  }
}

`
)
