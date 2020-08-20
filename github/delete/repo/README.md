<!-- markdownlint-disable-file MD013 -->
<!-- markdownlint-disable-file MD033 -->
<!-- markdownlint-disable-file MD034 -->

# Ritchie Formula

## Premisses

- Set GITHUB credentials ($ rit set credentials) with USERNAME, TOKEN

## Command

- Prompt

```bash
rit github delete repo
```

- Docker

```bash
rit github delete repo --docker
```

- Stdin

```bash
echo '{"project_name":"dennis-ritchie"}' | rit github delete repo --stdin
```

- Stdin + Docker

```bash
echo '{"project_name":"dennis-ritchie"}' | rit github delete repo --stdin --docker
```

## Description

This Github delete command allows the user to delete a Github PUBLIC or PRIVATE repository on the user domain.

The user has to inform only one input: 

- the repository name on Github 

Ex: The NAME would be "ritchie-formulas" for URL https://github.com/ZupIT/ritchie-formulas

## Demonstration

- Command execution

![Alt Text](https://media.giphy.com/media/RK5XCK1ZczsOBHpgoM/giphy.gif)
