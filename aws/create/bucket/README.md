<!-- markdownlint-disable-file MD013 -->
<!-- markdownlint-disable-file MD033 -->
<!-- markdownlint-disable-file MD041 -->

## Requirements

- Golang
- Set AWS credentials (\$ rit set credential) with accesskeyid, secretaccesskey

```bash
rit set credential
```

## Command

```bash
rit aws create bucket
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
