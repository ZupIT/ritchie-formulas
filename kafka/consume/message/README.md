# Description

This formula consume a kafka topic in realtime and print messages.
You can test it using docker with the `rit docker generate compose` formula.

## Command

```bash
rit kafka consume message
```

## Requirements

- Have an available Kafka URL.
- [Golang installed](https://golang.org/doc/install)

## Demonstration

- Command execution

![Alt Text](https://github.com/ZupIT/ritchie-formulas/raw/master/kafka/consume/message/docs/kafka-consume-message.gif)
