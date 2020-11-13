<!-- markdownlint-disable-file MD013 -->
<!-- markdownlint-disable-file MD033 -->
# Description

This Github publish command allows the user to create a Github Ritchie formulas repository based on a local repository.
It also generates a release and add it to Ritchie repositories ($ rit add repo).
If the repository already exists, it will commit the new code and generate the new release version informed, before updating Ritchie repositories ($ rit update repo).

The user has to inform 4 different kinds of inputs:

- the repository's privacy

- the Github repository's name

- the local repository path you wish to publish

- the release version to generate

## Command

```bash
rit github publish repo
```

## Requirements

- Git
- Set Github Credentials

## Demonstration

- Command execution

![Alt Text](https://media.giphy.com/media/KAqByBf4loMxXbv3NY/giphy.gif)

- Published Repo after executing command

<img class="special-img-class" src="https://github.com/ZupIT/ritchie-formulas/raw/master/github/publish/repo/docs/img/repo-published-on-github.png" />

- Ritchie Listed Repo after executing command

<img class="special-img-class" src="https://github.com/ZupIT/ritchie-formulas/raw/master/github/publish/repo/docs/img/rit-list-repo.png" />
