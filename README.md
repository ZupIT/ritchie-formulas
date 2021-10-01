<!-- markdownlint-disable MD041 MD033 MD013-->
[![CircleCI](https://circleci.com/gh/ZupIT/ritchie-formulas/tree/ritchie-2.0.0.svg?style=shield)](https://circleci.com/gh/ZupIT/ritchie-formulas)
[![pre-commit](https://img.shields.io/badge/pre--commit-enabled-brightgreen?logo=pre-commit&logoColor=white)](https://github.com/pre-commit/pre-commit)

<img class="special-img-class" src="/docs/img/ritchie-banner.png" />

# **Ritchie's commons formula repository**

This repository contains the community formulas [**ritchie-cli**](https://github.com/ZupIT/ritchie-cli) executes.

## **Create a new formula**
You can create your own formula, follow the steps below to create one: 

**Step 1.** Fork the repository;

**Step 2.** Create a branch: `git checkout -b <branch_name>`;

**Step 3.** Create a new formula, using the forked repository as a Ritchie
workspace: **`rit create formula`**
If you need help check out [**how to create formulas on Ritchie**](https://docs.ritchiecli.io/getting-started/creating-formulas) on documentation.

**Step 4.** Build and use the new formula: **`rit build formula`** or use **--watch** to keep observing changes on formula code live: **`rit build formula --watch`**

**Step 5.** Run `pre-commit.sh` to lint your code;

**Step 6.** Run **`go test -v ./.github/workflows/validation/...`** to test your code and formula structure. _(GoLang Required)_;

**Step 7.** Commit your implementation;

**Step 8.** Push your branch;

**Step 9.** Open a pull request on the repository for analysis.

## **Add support to other languages on create formula command**

The rit create formula command uses the **`/templates/create_formula`** folder
to list the languages options. If you like to edit a language template
or to add more language to create formula command please access the
[**Languages Template Tutorial**](https://github.com/ZupIT/ritchie-formulas/tree/master/templates/create_formula) tutorial. 

## [**Documentation**](https://docs.ritchiecli.io)
For more information, access [**Ritchie's documentation**](https://docs.ritchiecli.io).

[![Documentation](/docs/img/documentation-ritchie.png)](https://docs.ritchiecli.io)

## **Contributing**

Feel free to use, recommend improvements, or contribute to new implementations.

Check out our [**contributing guide**](https://github.com/ZupIT/ritchie-cli/blob/master/CONTRIBUTING.md) to learn about our development process, how to suggest bugfixes and improvements. 

## **Community**

Feel free to reach out to us at:

### [**Zup Open Source Forum**](https://forum.zup.com.br/c/en/9)

[![Zup forum](/docs/img/zup-forum-topics.png)](https://forum.zup.com.br/c/en/9)