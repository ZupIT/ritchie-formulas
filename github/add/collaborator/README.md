<!-- markdownlint-disable-file MD013 -->
# Add a collaborator on a repository

## Command

```bash
rit github add collaborator
```

![Example](https://github.com/ZupIT/ritchie-formulas/raw/master/github/add/collaborator/src/docs/github.gif)

## Description

This formula allows adding a new collaborator by typing only two parameters.
(collaborator username and repository name)

### STDIN Example

```bash
echo '{"collaborator_user":"value", "repository_name":"value"}' | rit github add collaborator --stdin
```
