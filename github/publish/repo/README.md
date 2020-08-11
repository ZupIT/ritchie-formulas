<!-- markdownlint-disable-file MD013 -->
<!-- markdownlint-disable-file MD033 -->

# Ritchie Formula

## Premisses

- Set GITHUB credentials ($ rit set credentials) with USERNAME, TOKEN & EMAIL

## Command

- Prompt

```rit github publish repo```

- Docker

```rit github publish repo --docker```

- Stdin

```echo '{"privacy":"true", "project_name":"ritchie-formulas-demo", "workspace_path":"/home/users/dennis/ritchie-formulas-local", "version":"v1.0.0"}' | rit github publish repo --stdin```

- Stdin + Docker

```echo '{"privacy":"true", "project_name":"ritchie-formulas-demo", "workspace_path":"/home/users/dennis/ritchie-formulas-local", "version":"v1.0.0"}' | rit github publish repo --stdin --docker```

## Description

This Github publish command allows the user to create a Github Ritchie formulas repository based on a local repository.
It also generates a release and add it to Ritchie repositories ($ rit add repo).
If the repository already exists, it will commit the new code and generate the new release version informed, before updating Ritchie repositories ($ rit update repo)

The user has to inform 4 different kinds of inputs:

- the repository's privacy

- the Github repository's name

- the local repository path you wish to publish

- the release version to generate

## Demo

- Command execution 

![Alt Text](https://media.giphy.com/media/KAqByBf4loMxXbv3NY/giphy.gif)

- Published Repo after executing command

<img class="special-img-class" src="/github/publish/repo/docs/img/repo-published-on-github.png" />

- Ritchie Listed Repo after executing command

<img class="special-img-class" src="/github/publish/repo/docs/img/rit-list-repo.png" />
