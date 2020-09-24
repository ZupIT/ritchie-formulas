<!-- markdownlint-disable MD041 MD033 MD013-->
[![CircleCI](https://circleci.com/gh/ZupIT/ritchie-formulas/tree/ritchie-2.0.0.svg?style=shield)](https://circleci.com/gh/ZupIT/ritchie-formulas)
[![pre-commit](https://img.shields.io/badge/pre--commit-enabled-brightgreen?logo=pre-commit&logoColor=white)](https://github.com/pre-commit/pre-commit)

<img class="special-img-class" src="/docs/img/ritchie-banner.png" />

# Ritchie's commons formula repository

This repository contains the community formulas which can be executed by the [ritchie-cli](https://github.com/ZupIT/ritchie-cli).

## Create a new formula

1. Fork the repository.
2. Create a branch:`git checkout -b <branch_name>`
3. Create a new formula, using the forked repository as a Ritchie
workspace: `rit create formula` if you need help please visit
 [how to create formulas on Ritchie](https://docs.ritchiecli.io/getting-started/creating-formulas)
4. Build and use the new formula: `rit build formula`
 or use --watch to keep observing changes on formula code live: `rit build formula --watch`
5. Run `pre-commit.sh` to lint your code
6. Run `go test -v ./.circleci/validation/...` to test your code and formula
structure. _(GoLang Required)_
7. Commit your implementation.
8. Push your branch.
9. Open a pull request on the repository for analysis.

## Add support to other languages on create formula command

The rit create formula command use the folder `/templates/create_formula`
to list the languages options. If you like to edit some language template
or to add more language to create formula command please access
the following tutorial:
[Languages Template Tutorial](https://github.com/ZupIT/ritchie-formulas/tree/master/templates/create_formula)

## Full Documentation

[Gitbook](https://docs.ritchiecli.io)

## Contributing

[Contribute to the Ritchie community](https://github.com/ZupIT/ritchie-cli/blob/master/CONTRIBUTING.md)

## Zup Products

[Zup open source](https://opensource.zup.com.br)
