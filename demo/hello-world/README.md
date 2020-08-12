<!-- markdownlint-disable-file MD013 -->

# Ritchie Formula Hello World

## Premisses

- [Golang installed](https://golang.org/doc/install)

## Command

- Prompt

```bash
rit demo hello-world
```

- Docker

```bash
rit demo hello-world --docker
```

- Stdin

```bash
echo '{"input_text":"Dennis", "input_bool":"false", "input_list":"false", "input_password":"Ritchie"}' | rit demo hello-world --stdin
```

- Stdin + Docker

```bash
echo '{"input_text":"Dennis", "input_bool":"false", "input_list":"false", "input_password":"Ritchie"}' | rit demo hello-world --stdin --docker
```

## Description

This Hello World command has been implemented for new user to discover Ritchie.

It allows the user to inform 4 different kinds of inputs:

- a text

- a boolean

- a list

- and a password (secret)

## Demo

![Alt Text](https://media.giphy.com/media/VdQGuZoyozL9J1Lhhl/giphy.gif)
