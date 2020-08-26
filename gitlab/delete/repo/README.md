<!-- markdownlint-disable-file MD013 -->
<!-- markdownlint-disable-file MD033 -->
<!-- markdownlint-disable-file MD034 -->

# Ritchie Formula

## Premisses

- Set GITLAB credentials ($ rit set credentials) with USERNAME, TOKEN

## Command

- Prompt

```bash
rit gitlab delete repo
```

- Docker

```bash
rit gitlab delete repo --docker
```

- Stdin

```bash
echo '{"project_name":"dennis-ritchie"}' | rit gitlab delete repo --stdin
```

- Stdin + Docker

```bash
echo '{"project_name":"dennis-ritchie"}' | rit gitlab delete repo --stdin --docker
```

## Description

This Gitlab delete command allows the user to delete a Gitlab PUBLIC or PRIVATE repository on the user domain.

The user has to inform only one input:

- the repository name on Gitlab

Ex: The NAME would be "ritchie-formulas" for URL https://github.com/ZupIT/ritchie-formulas

## How it works

![gif](https://media.giphy.com/media/RJJMed5qzQeynuztBS/giphy.gif)
