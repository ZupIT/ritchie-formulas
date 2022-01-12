<!-- markdownlint-disable MD041 MD033 MD013-->
[![CircleCI](https://circleci.com/gh/ZupIT/ritchie-formulas/tree/ritchie-2.0.0.svg?style=shield)](https://circleci.com/gh/ZupIT/ritchie-formulas)
![pre-commit](https://img.shields.io/badge/pre--commit-enabled-brightgreen?logo=pre-commit&logoColor=white)

<img class="special-img-class" src="/docs/img/ritchie-banner.png" />

## **Table of contents**

### 1. [**Ritchie's commons formula repository**](#ritchie's-commons-formula-repository)
> #### 1.1. [**Create a new formula**](#create-a-new-formula)
> #### 1.2. [**Add support to other languages on create formula command**](#add-support-to-other-languages-on-create-formula-command)
### 2. [**Documentation**](#documentation)
### 3. [**Contributing**](#contributing)
### 4. [**License**](#license)
### 5. [**Community**](#community)

# **Ritchie's commons formula repository**

This repository contains the community formulas [**ritchie-cli**](https://github.com/ZupIT/ritchie-cli) executes.

## **Create a new formula**
You can create your own formula by following these steps:

**Step 1.** Fork the repository;

**Step 2.** Create a branch: `git checkout -b <branch_name>`;

**Step 3.** Create a new formula, using the forked repository as a Ritchie
workspace: **`rit create formula`**
If you need help check out [**how to create formulas on Ritchie**](https://docs.ritchiecli.io/getting-started/creating-formulas) on our documentation.

**Step 4.** Build and use the new formula: **`rit build formula`** or use **--watch** to keep observing changes on formula code live: **`rit build formula --watch`**

**Step 5.** Run `pre-commit.sh` to lint your code;

**Step 6.** Run **`go test -v ./.github/workflows/validation/...`** to test your code and formula structure. _(GoLang Required)_;

**Step 7.** Commit your implementation;

**Step 8.** Push your branch;

**Step 9.** Open a pull request on the repository for analysis.

## **Add support to other languages on create formula command**

The rit create formula command uses the **`/templates/create_formula`** folder
to list the languages options. If you like to edit a language template
or to add more language to create formula command, please access the
[**Languages Template Tutorial**](https://github.com/ZupIT/ritchie-formulas/tree/master/templates/create_formula) tutorial.

## [**Documentation**](https://docs.ritchiecli.io)
For more information, access [**Ritchie's documentation**](https://docs.ritchiecli.io).

[![Documentation](/docs/img/documentation-ritchie.png)](https://docs.ritchiecli.io)

## **Contributing**

Feel free to use, recommend improvements, or contribute to new implementations.

Check out our [**contributing guide**](https://github.com/ZupIT/ritchie-cli/blob/master/CONTRIBUTING.md) to learn about our development process, how to suggest bug fixes and improvements.

### **Developer Certificate of Origin - DCO**

 This is a security layer for the project and for the developers. It is mandatory.
 
 Follow one of these two methods to add DCO to your commits:
 
**1. Command line**
 Follow the steps: 
 **Step 1:** Configure your local git environment adding the same name and e-mail configured at your GitHub account. It helps to sign commits manually during reviews and suggestions.

 ```
git config --global user.name “Name”
git config --global user.email “email@domain.com.br”
```
**Step 2:** Add the Signed-off-by line with the `'-s -S'` flag in the git commit command:

```
$ git commit -s -m "This is my commit message"
```

**2. GitHub website**
You can also manually sign your commits during GitHub reviews and suggestions, follow the steps below: 

**Step 1:** When the commit changes box opens, manually type or paste your signature in the comment box, see the example:

```
Signed-off-by: Name < e-mail address >
```

For this method, your name and e-mail must be the same registered on your GitHub account.


## **License**
[**Apache License 2.0**](https://github.com/ZupIT/ritchie-formulas/blob/main/LICENSE).

## **Community**

Do you have any question about Ritchie? Let's chat in our [**forum**](https://forum.zup.com.br/).
