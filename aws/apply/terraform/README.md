# Terraform aws apply

## Premisses

- [Terraform installed](https://www.terraform.io/downloads.html)
- Set Github credentials
- Set AWS credentials

You can set credentials by running *rit set credential* and providing USERNAME, TOKEN and EMAIL for Github and ACCESS KEY ID and SECRET ACCESS KEY for AWS.

## Command

- Prompt 

```bash
rit aws apply terraform
```

*It is necessary to have [Golang installed](https://golang.org/doc/install) for this command to work*

- Docker 


```bash
rit aws apply terraform --docker
```

*It is necessary to have [Docker installed](https://docs.docker.com/get-docker) for this command to work*


- Stdin 

```bash
echo '{"repository":"https://github.com/eduardorcury/ritchie-demo", "terraform_path":"/terraform", "environment":"dev"}' | rit aws apply terraform --stdin
```

*It is necessary to have [Golang installed](https://golang.org/doc/install) for this command to work*

- Stdin + Docker 

```bash
echo '{"repository":"https://github.com/eduardorcury/ritchie-demo", "terraform_path":"/terraform", "environment":"dev"}' | rit aws apply terraform --stdin --docker
```

*It is necessary to have [Docker installed](https://docs.docker.com/get-docker) for this command to work*

## Description

This command allows the user to execute terraform init, terraform plan and terraform apply command on a given repository. The command also loads the variables located on the files './variables/common.tfvars' and 'variables/**env**.tfvars', where **env** is the environment name provided.

Note that this command will clone the provided repository in the present working directory.

The user has to provide 3 inputs:

- the repository URL

- the path to the terraform files

- the environment name

The equivalent terraform apply command is:

```bash
terraform apply -var-file=./variables/common.tfvars -var-file=variables/{ENV}.tfvars -auto-approve
```

## Demonstration

- Command execution

![Demo gif](https://github.com/eduardorcury/ritchie-demo/blob/main/media/rit-demo.gif)

- The created resource

In this demo, an AWS IAM user was created using terraform and ritchie.

![Img](https://github.com/eduardorcury/ritchie-demo/blob/main/media/resource-img.png)
