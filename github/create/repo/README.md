# Create a local and remote repository

## Command

### Prompt

`$ rit github create repo`

### Stdin Unix

`$ echo "{\"project_name\":\"your_project_name\", \"project_description\":\"your_description\", \"private\":\"false or true\"}" | rit github create repo --stdin`

### Stdin Windows PowerShell

`$ echo '{"project_name":"your_project_name", "project_description":"your_description", "private":"false or true"}' | rit github create repo --stdin`

## Requirements

- git installed
- set github username and token credentials using `$ rit set credential`

## How to generate personal access token

To generate a personal access token click [here](https://github.com/settings/tokens)

## How it works

![gif](https://media.giphy.com/media/U5bfisA8omNg52kEG1/giphy.gif)
