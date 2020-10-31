<!-- markdownlint-disable-file MD013 -->
<!-- markdownlint-disable-file MD033 -->

# Create bucket aws

## Premisses

- Set AWS credentials ($ rit set credential) with accesskeyid, secretaccesskey

```bash
rit set credential
```

## command

- Prompt

```bash
rit aws create bucket
```

- Docker

```bash
rit aws create bucket --docker
```

- Stdin

```bash
echo '{"region":"us-east-1", "bucket":"ritchie-formulas-demo-stdin"}' | rit aws create bucket --stdin
```

- Stdin + Docker

```bash
echo '{"region":"us-east-1", "bucket":"ritchie-formulas-demo-stdin-docker"}' | rit aws create bucket --stdin --docker
```

## Description

This AWS Create Bucket command allows the user to create a bucket in AWS S3
If the repository already exists, the user don't receive error.

The user has to inform 2 different kinds of inputs:

- the Region Bucket name

- the Bucket name

## Demonstration

- Command execution

![Alt Text](https://media.giphy.com/media/5OI7ywkzgCkbS14LB5/source.gif)

- Created bucket after executing command
