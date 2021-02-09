# Description

This formula receives 5 inputs (region, vpc name, vpc cidr, vpc azs, customer name)
and adds vpc module configuration into the terraform project.

Be in the terraform project directory and run formula command.

After execution, the main.tf and qa.tfvars files will add new settings for creating a VPC.

## Command

```bash
rit aws add terraform-vpc
```

## Requirements

- Golang

## Demonstration

![Demo](hhttps://media.giphy.com/media/BdONqMwV5rbcvhKOSr/giphy.gif)