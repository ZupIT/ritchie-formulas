# Github merge branch

## Command

```bash
rit github merge branch
```

## Requirements

- git installed

## How to generate personal access token

To generate a personal access token click
 [here](https://github.com/settings/tokens)

## How it works

```bash
git pull origin {{current_branch}}
git fetch
git branch -D {{dest_branch}}
git checkout {{dest_branch}}
git pull origin {{current_branch}}
```

If push is true

```bash
 git push
 git checkout {{current_branch}}
```
