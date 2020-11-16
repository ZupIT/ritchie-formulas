# Describe Instances AWS EC2

## Premisses

- [Ritchie installed](https://docs.ritchiecli.io/v/v2.0-pt/getting-started/installation)
- Set AWS credentials (\$ rit set credentials).

## Command

- Prompt

```bash
rit aws describe-instances ec2
```

_It is necessary to have [Python](https://www.python.org/downloads/) installed_

- Docker

```bash
rit aws describe-instances ec2 --docker
```

## Description

This formula will describes all instances in the EC2 service.

Select your region name and filter instance state:

- AWS region i.e., us-east-1, us-west-1, sa-east-1
- Instance state i.e., all, running and stopped
