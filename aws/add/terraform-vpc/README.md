<!-- markdownlint-disable-file MD013 -->
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

![Demo](https://github.com/maurineimirandazup/ritchie-formulas/blob/feature-awsvpc-readme/docs/img/terraform-vpc.gif)
