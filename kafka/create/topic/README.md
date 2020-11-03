<!-- markdownlint-disable-file MD013 -->
<!-- markdownlint-disable-file MD033 -->

# create kafka topic

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

This formula create a topic on the inform kafka URL.
You can test it using docker with the `rit docker generate compose` formula.

```bash
kafka server: Topic with this name already exists. - Topic 'XXXX' already exists.
```

## Demonstration

- Command execution

![Alt Text](https://media.giphy.com/media/VuSYXuKCgxJnIkPTYJ/giphy.gif)

- Created bucket after executing command
