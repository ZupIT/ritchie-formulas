package tpl

const Circleciconfig = `
version: 2.1

orbs:
  windows: circleci/windows@2.4.0

references:
  images:
    ubuntu: &UBUNTU_IMAGE cimg/base:2020.01

  environment: &ENVIRONMENT
    TERRAFORM_ENV : "{{.ProjectName}}"
    TERRAFORM_VERSION: "0.12.13"


  filters: &FILTERS_DELIVERY
    branches:
      only:
        - qa
        - master

executors:
  ubuntu-executor:
    docker:
      - image: *UBUNTU_IMAGE
        user: root
    working_directory: /workspace

jobs:
  terraform:
    environment:
      <<: *ENVIRONMENT
    executor: ubuntu-executor
    steps:
      - checkout
      - run:
          name: Applying terraform changes to the environment
          command: |
            . ./.circleci/scripts/credentials.sh
            . ./.circleci/scripts/terraform-env.sh
            . ./.circleci/scripts/terraform-run.sh

workflows:

  delivery:
    jobs:
      - terraform:
          filters:
            <<: *FILTERS_DELIVERY
`
