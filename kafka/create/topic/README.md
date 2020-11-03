<!-- markdownlint-disable-file MD013 -->
<!-- markdownlint-disable-file MD033 -->

# create kafka topic

## Premisses

- The kafka cluster is accessible

## Premisse

Have an available Kafka URL.

## Command

- Prompt

```bash
rit kafka create topic
```

- Docker

```bash
rit kafka create topic --docker
```

- Stdin

```bash
echo '{"urls":"localhost:9092", "name":"topic-stdin", "replication":"1", "partitions":"1"}' | rit kafka create topic --stdin
```

- Stdin + Docker

```bash
echo '{"urls":"localhost:9092", "name":"topic-stdin", "replication":"1", "partitions":"1"}' | rit kafka create topic --stdin --docker
```

## Description

This create kafka topic command allows the user to delete a bucket in AWS S3

If the command using stdin method, the question for list and confirm name of the bucket is suppressed.

If the kafka has exists, you receive the error:

```bash
kafka server: Topic with this name already exists. - Topic 'XXXX' already exists.
```

## Demonstration

- Command execution

![Alt Text](https://media.giphy.com/media/VuSYXuKCgxJnIkPTYJ/giphy.gif)

- Created bucket after executing command
