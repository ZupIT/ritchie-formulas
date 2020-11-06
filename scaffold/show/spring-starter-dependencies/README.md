<!-- markdownlint-disable-file MD013 -->
<!-- markdownlint-disable-file MD033 -->

# Spring Starter Dependencies Show Scaffold

## Premisses

- [Ritchie installed](https://docs.ritchiecli.io/v/v2.0-pt/getting-started/installation)

## Command

- Prompt

```bash
rit scaffold show spring-starter-dependencies
```

_It is necessary to have [Golang](https://golang.org/doc/install) installed for this command to work_

- Docker

```bash
rit scaffold show spring-starter-dependencies --docker
```

_It is necessary to have [Docker installed](https://docs.docker.com/get-docker) for this command to work_

## Description

This formula will list the dependencies of the spring boot starter grouped by type and with the following information:

- Id
- Name
- Description

You can use these ids to generate a scaffold project using spring boot.

## Demonstration

<img src="demo.gif">
