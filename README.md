<!-- markdownlint-disable MD041 MD033 -->
[![CircleCI](https://circleci.com/gh/ZupIT/ritchie-formulas/tree/ritchie-2.0.0.svg?style=shield)](https://circleci.com/gh/ZupIT/ritchie-formulas)
[![pre-commit](https://img.shields.io/badge/pre--commit-enabled-brightgreen?logo=pre-commit&logoColor=white)](https://github.com/pre-commit/pre-commit)

<img class="special-img-class" src="/docs/img/ritchie-banner.png" />

# Ritchie commons formula repository

This repository contains the community formulas which can be executed by the [ritchie-cli](https://github.com/ZupIT/ritchie-cli).

## Contribute to the Repository with your formulas

1. Fork the repository.
2. Create a branch: 
```bash
    git checkout -b <branch_name>
``` 
3. Create a new formula, using the forked repository as a Ritchie workspace:
```bash
    rit create formula
```
4. Build and use the new formula:
```bash
    rit build formula
```
or use watch to watch changes on formula code
```bash
    rit build formula --watch
```
5. Commit your implementation:
```bash
    git add *
    git commit -m '<commit_message>'
```
6. Push your branch: 
```bash
    git push origin
```
7. Open a pull request on the repository for analysis.

## Full Documentation

- [Gitbook](https://docs.ritchiecli.io)


## Contributing

[Contribute to the Ritchie community](https://github.com/ZupIT/ritchie-cli/blob/master/CONTRIBUTING.md)


## Zup Products

- [Zup open source](https://opensource.zup.com.br)