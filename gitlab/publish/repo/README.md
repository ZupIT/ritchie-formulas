<!-- markdownlint-disable-file MD013 -->
<!-- markdownlint-disable-file MD033 -->

# Ritchie Formula

## Premisses

- Set GITLAB credentials (\$ rit set credentials) with USERNAME, TOKEN & EMAIL

## Command

- Prompt

```bash
rit gitlab publish repo
```

- Docker

```bash
rit gitlab publish repo --docker (WIP)
```

- Stdin

```bash
echo '{"privacy":"true", "project_name":"ritchie-formulas-demo", "workspace_path":"/home/users/dennis/ritchie-formulas-local", "version":"v1.0.0"}' | rit gitlab publish repo --stdin
```

- Stdin + Docker

```bash
echo '{"privacy":"true", "project_name":"ritchie-formulas-demo", "workspace_path":"/home/users/dennis/ritchie-formulas-local", "version":"v1.0.0"}' | rit gitlab publish repo --stdin --docker
```

## Description

This Gitlab publish command allows the user to create a Gitlab Ritchie formulas repository based on a local repository.
It also generates a release and add it to Ritchie repositories ($ rit add repo).
If the repository already exists, it will commit the new code and generate the new release version informed, before updating Ritchie repositories ($ rit update repo).

The user has to inform 4 different kinds of inputs:

- the repository's privacy

- the Gitlab repository's name

- the local repository path you wish to publish

- the release version to generate

## Demo

![gif](https://github.com/ZupIT/ritchie-formulas/raw/master/gitlab/publish/repo/doc/gif.gif)
