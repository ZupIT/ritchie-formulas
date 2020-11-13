<!-- markdownlint-disable-file MD013 -->
<!-- markdownlint-disable-file MD033 -->

## Requirements

- Golang
- The bucket is clean
- Set AWS credentials (\$ rit set credential) with accesskeyid, secretaccesskey

```bash
rit set credential
```

## Command

```bash
rit aws delete bucket
```

## Description

This AWS Delete Bucket command allows the user to delete a bucket in AWS S3

If the command using stdin method, the question for list and confirm name of the bucket is suppressed.

The user has to inform 2 different kinds of inputs:

- the Region Bucket name

- the Bucket name

If the bucket has content, you receive the error:

```bash
error: BucketNotEmpty: The bucket you tried to delete is not empty
```

For this error you can execute before `rit aws clean bucket`

## Demonstration

- Command execution

![Alt Text](https://media.giphy.com/media/UwN20TEphzatoNaSWg/source.gif)

- deleted bucket after executing command
