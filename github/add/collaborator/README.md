# Add a new collaborator to your repository.

## Description

This formula allows adding a new collaborator, by typing only two parameters. (Collaborator username and repository name).

## Requirements

* NodeJS
* NPM

### You need to set your credentials.
1. You need generate an access token with access to the **repo**  scope. [Generate Here!](https://github.com/settings/tokens)
2. Set your credentials at rit. [Use case.](https://docs.ritchiecli.io/use-cases/using-first-commands/credentials)

## Examples

```sh
$ rit github add collaborator
```
![Example](doc/github.gif)

#### STDIN Example

```sh
$ echo '{"collaborator_user":"value", "repository_name":"value"}' | rit github add collaborator --stdin
```