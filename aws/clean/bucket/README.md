<!-- markdownlint-disable-file MD013 -->
<!-- markdownlint-disable-file MD033 -->

# clean bucket aws

## Premisses

- Set AWS credentials ($ rit set credential) with accesskeyid, secretaccesskey

```bash
rit set credential
```

## command

- Prompt

```bash
rit aws clean bucket
```

- Docker

```bash
rit aws clean bucket --docker
```

- Stdin

```bash
echo '{"region":"us-east-1", "bucket":"ritchie-formulas-demo-stdin"}' | rit aws clean bucket --stdin
```

- Stdin + Docker

```bash
echo '{"region":"us-east-1", "bucket":"ritchie-formulas-demo-stdin-docker"}' | rit aws clean bucket --stdin --docker
```

## Description

This AWS Clean Bucket command allows the user to clean a bucket in AWS S3

If the command using stdin method, the question for list and confirm name of the bucket is suppressed.

The user has to inform 2 different kinds of inputs:

- the Region Bucket name

- the Bucket name

## Demonstration

- Command execution

![Alt Text](https://media.giphy.com/media/R5OMk8P6yCXp1n2Uqo/source.gif)

- cleand bucket after executing command
