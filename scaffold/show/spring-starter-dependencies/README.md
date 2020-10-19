# Spring Starter Dependencies Show Scaffold

## Premisses

- [Ritchie installed](https://docs.ritchiecli.io/v/v2.0-pt/getting-started/installation)

## Command

- Prompt

```bash
rit scaffold show spring-starter-dependencies
```

*It is necessary to have [Golang](https://golang.org/doc/install) installed for this command to work*

- Docker

```bash
rit scaffold show spring-starter-dependencies --docker
```

*It is necessary to have [Docker installed](https://docs.docker.com/get-docker) for this command to work*

## Description

This formula will list the dependencies of the spring boot starter grouped by type and with the following information:
- Id
- Name
- Description

You can use these ids to generate a scaffold project using spring boot.

## Demonstration

<img src="demo.gif">
