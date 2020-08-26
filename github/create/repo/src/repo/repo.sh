#!/bin/bash

checkCommand() {
  if ! command -v "$1" >/dev/null; then
    echo "Error: $1 command required"
    exit 1
  fi
}

removeSpaces() {
  echo "${1}" | xargs | tr " " -
}

checkSpace() {
  tmp="$1"

  if [[ "$1" = *" "* ]]; then
    echo >&2 "Removing spaces from Project name..."
    tmp=$(removeSpaces "$1")
    echo >&2 "Project name without spaces: $tmp"
  fi

  echo "$tmp"
}

checkProjectName() {
  if [[ ! "$1" =~ ^[a-zA-Z0-9-]+$ ]]; then
    echo "Project name cannot contain special characters"
    exit 1
  fi
}

run() {
  checkCommand git

  PROJECT_NAME=$(checkSpace "$PROJECT_NAME")

  checkProjectName "$PROJECT_NAME"

  if [[ "$WORKSPACE_PATH" != " " ]]; then
    cd "$WORKSPACE_PATH" || exit 1
  else
    mkdir "$PROJECT_NAME"
    cd "$PROJECT_NAME" || exit 1
    echo "$PROJECT_DESCRIPTION" >> README.md
  fi

  git init

  git add .
  git commit -m "Initial Commit"

  curl -H 'Authorization: token '"$TOKEN" https://api.github.com/user/repos -d '{"name":"'"$PROJECT_NAME"'", "private":'"$PRIVATE"'}' &&
  git remote add origin https://"$USERNAME":"$TOKEN"@github.com/"$USERNAME"/"$PROJECT_NAME".git &&
  git push origin master

  echo "✅ Repository successfully initialized with git and added on Github!!"
}
