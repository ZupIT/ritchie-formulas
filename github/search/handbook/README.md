<!-- markdownlint-disable-file MD013 -->
<!-- markdownlint-disable-file MD033 -->
# Description

This formula search a handbook on a github repository

## Command

```bash
rit github search handbook
```

### Adding ritchie-formulas to Ritchie CLI

- Run command

```bash
echo '{"provider":"Github", "name":"demo", "version":"2.0.0", "url":"https://github.com/ZupIT/ritchie-formulas", "token": null, "priority":1}' | rit add repo --stdin
```

- or add via [add repo](https://docs.ritchiecli.io/v/v2.0-pt/tutorials/formulas/como-executar-formulas/formula-hello-world)

```bash
rit add repo
```

## Requirements

- [Golang installed](https://golang.org/doc/install)
- Github Account

## Demonstration

<img src="https://github.com/ZupIT/ritchie-formulas/raw/master/github/search/handbook/demo.gif">
