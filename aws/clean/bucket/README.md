<!-- markdownlint-disable-file MD013 -->
<!-- markdownlint-disable-file MD033 -->

# Ritchie Formula

## Requirements

- Golang
- Set AWS credentials (\$ rit set credential) with accesskeyid, secretaccesskey

```bash
rit set credential
```

## Command

```bash
rit aws clean bucket
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
