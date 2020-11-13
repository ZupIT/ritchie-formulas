<!-- markdownlint-disable-file MD013 -->
<!-- markdownlint-disable-file MD033 -->
<!-- markdownlint-disable-file MD041 -->

## Requirements

- Set GITLAB credentials (\$ rit set credentials) with USERNAME, TOKEN & EMAIL

## Command

```bash
rit gitlab publish repo
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

## Demonstration

![gif](https://github.com/ZupIT/ritchie-formulas/raw/master/gitlab/publish/repo/doc/gif.gif)
