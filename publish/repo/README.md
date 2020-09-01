<!-- markdownlint-disable-file MD013 -->
<!-- markdownlint-disable-file MD033 -->
<!-- markdownlint-disable-file MD024 -->

# Ritchie Publish Repo Formula

## Command

- Set GITHUB or GITLAB credentials ($ rit set credentials) with USERNAME, TOKEN & EMAIL.

- This formula can currently only be executed locally on MacOs or Linux.

## Command

- Prompt

```bash
rit publish repo
```

- Stdin

```bash
echo '{"provider":"Github","privacy":"true", "project_name":"ritchie-formulas-demo", "workspace_path":"/home/users/dennis/ritchie-formulas-local", "version":"v1.0.0"}' | rit publish repo --stdin
```

## Description

This publish command allows the user to create a Githab or Gitlab Ritchie formulas repository based on a local repository.

It encapsulates the `rit github publish repo` and the `rit gitlab publish repo` formulas.

Therefore, it also generates a release and add it to Ritchie repositories ($ rit add repo).
If the repository already exists, it will commit the new code and generate the new release version informed, before updating Ritchie repositories ($ rit update repo).

The user has to inform 5 different kinds of inputs:

- the provider (Github or Gitlab)

- the repository's privacy

- the Gitlab repository's name

- the local repository path you wish to publish

- the release version to generate

## Demonstration

- Command execution

<img class="special-img-class" src="/publish/repo/docs/img/Github.png" />

<img class="special-img-class" src="/publish/repo/docs/img/Gitlab.png" />
