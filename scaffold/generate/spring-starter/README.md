# Spring Boot Scaffold Generator


## Premisses

- [Golang installed](https://golang.org/doc/install)
- [Ritchie installed](https://docs.ritchiecli.io/v/v2.0-pt/getting-started/installation)

## Adding ritchie-formulas to Ritchie CLI

- Run command 
```bash
echo '{"provider":"Github", "name":"demo", "version":"2.0.0", "url":"https://github.com/ZupIT/ritchie-formulas", "token": null, "priority":1}' | rit add repo --stdin
```

- or add via [add repo](https://docs.ritchiecli.io/v/v2.0-pt/tutorials/formulas/como-executar-formulas/formula-hello-world)
```bash
rit add repo
```

## Command

- Prompt
```bash
rit scaffold generate spring-starter
```

## Description

This formula generate a scaffold project using spring boot and dependencies that you to choose 


## Demo

<img src="demo.gif">
