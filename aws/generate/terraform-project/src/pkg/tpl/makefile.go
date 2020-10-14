package tpl

const (
	Makefile = `.ONESHELL:
.SHELL := /usr/bin/bash
.PHONY: apply plan prep
ENV = $(ENVIRONMENT)
VARS="variables/$(ENV).tfvars"

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

set-env:
	@if [ -z $(ENV) ]; then \
		echo "ENV was not set"; \
		ERROR=1; \
	 fi
	@if [ ! -z $${ERROR} ] && [ $${ERROR} -eq 1 ]; then \
		echo "Example usage: ENVIRONMENT=demo make plan "; \
		exit 1; \
	 fi
	@if [ ! -f "$(VARS)" ]; then \
		echo "Could not find variables file: $(VARS)"; \
		exit 1; \
	 fi

prep: set-env ## Prepare the environment if needed
	@echo "Configuring the terraform backend"
	@terraform init -var-file=./variables/common.tfvars -var-file=$(VARS) -reconfigure -backend-config=$(ENV).tfbackend

plan: prep ## Show what terraform wants to do
	@echo "Planning the terraform steps"
	@terraform plan -var-file=./variables/common.tfvars -var-file=$(VARS)

apply: plan ## Let Terraform do his thing
	@echo "Applying the terraform plan"
	@terraform apply -var-file=./variables/common.tfvars -var-file=$(VARS)

apply-ci: plan ## Let Terraform do his thing on CI
	@echo "Applying the terraform plan"
	@terraform apply -var-file=./variables/common.tfvars -var-file=$(VARS) -auto-approve

destroy: set-env ## Let Terraform destroy everything
	@echo "Destroying the terraform environment"
	@terraform destroy -var-file=./variables/common.tfvars -var-file=$(VARS)
`
)
