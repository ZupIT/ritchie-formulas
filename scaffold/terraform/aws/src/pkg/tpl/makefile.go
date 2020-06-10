package tpl

const (
	Makefile = `.ONESHELL:
.SHELL := /usr/bin/bash
.PHONY: apply plan prep
ENV = $(ENVIRONMENT)
VARS="variables/$(ENV).tfvars"
BOLD=$(shell tput bold)
RED=$(shell tput setaf 1)
GREEN=$(shell tput setaf 2)
YELLOW=$(shell tput setaf 3)
RESET=$(shell tput sgr0)

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

set-env:
	@if [ -z $(ENV) ]; then \
		echo "$(BOLD)$(RED)ENV was not set$(RESET)"; \
		ERROR=1; \
	 fi
	@if [ ! -z $${ERROR} ] && [ $${ERROR} -eq 1 ]; then \
		echo "$(BOLD)Example usage: ENVIRONMENT=demo make plan $(RESET)"; \
		exit 1; \
	 fi
	@if [ ! -f "$(VARS)" ]; then \
		echo "$(BOLD)$(RED)Could not find variables file: $(VARS)$(RESET)"; \
		exit 1; \
	 fi

prep: set-env ## Prepare the environment if needed
	@echo "$(BOLD)Configuring the terraform backend$(RESET)"
	@terraform init -var-file=./variables/common.tfvars -var-file=$(VARS) -reconfigure -backend-config=$(ENV).tfbackend

plan: prep ## Show what terraform wants to do
	@echo "$(BOLD)Planning the terraform steps$(RESET)"
	@terraform plan -var-file=./variables/common.tfvars -var-file=$(VARS)

apply: plan ## Let Terraform do his thing
	@echo "$(BOLD)Apllying the terraform plan$(RESET)"
	@terraform apply -var-file=./variables/common.tfvars -var-file=$(VARS)

apply-ci: plan ## Let Terraform do his thing on CI
	@echo "$(BOLD)Apllying the terraform plan$(RESET)"
	@terraform apply -var-file=./variables/common.tfvars -var-file=$(VARS) -auto-approve

destroy: set-env ## Let Terraform destroy everything
	@echo "$(BOLD)Destroying the terraform environment$(RESET)"
	@terraform destroy -var-file=./variables/common.tfvars -var-file=$(VARS)
	
	`
)
