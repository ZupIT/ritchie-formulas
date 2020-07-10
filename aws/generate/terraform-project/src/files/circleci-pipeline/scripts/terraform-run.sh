#!/bin/bash

echo "Downloading terraform"

curl -fsSL "https://releases.hashicorp.com/terraform/${TERRAFORM_VERSION}/terraform_${TERRAFORM_VERSION}_linux_amd64.zip" -o terraform.zip

unzip terraform.zip

mv terraform /usr/local/bin

echo "Applying changes to ${TERRAFORM_ENV} ${ENVIRONMENT}"

cd src && AWS_ACCESS_KEY_ID=${AWS_ACCESS_KEY_ID} AWS_SECRET_ACCESS_KEY=${AWS_SECRET_ACCESS_KEY} make apply-ci
