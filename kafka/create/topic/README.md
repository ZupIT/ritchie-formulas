<!-- markdownlint-disable-file MD013 -->
<!-- markdownlint-disable-file MD033 -->
# Description

This formula create a topic on the inform kafka URL.
You can test it using docker with the `rit docker generate compose` formula.

```bash
kafka server: Topic with this name already exists. - Topic 'XXXX' already exists.
```

## Command

```bash
rit kafka create topic
```

## Requirements

- Have an available Kafka URL.
- [Golang installed](https://golang.org/doc/install)

## Demonstration

- Command execution

![Alt Text](https://media.giphy.com/media/VuSYXuKCgxJnIkPTYJ/giphy.gif)

- Created bucket after executing command
