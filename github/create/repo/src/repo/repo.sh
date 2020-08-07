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

  checkProjectName $PROJECT_NAME

  mkdir $PROJECT_NAME
  cd $PROJECT_NAME || exit

  git init
  echo $PROJECT_DESCRIPTION >> README.md

  if [[ $DOCKER_EXECUTION ]]; then
    read -p "Enter your email: " email
    git config --local user.name $USERNAME
    git config --local user.email $email
  fi
  
  git add .
  git commit -m "Initial Commit"
  
  curl -H 'Authorization: token '$TOKEN https://api.github.com/user/repos -d '{"name":"'$PROJECT_NAME'", "private":'$PRIVATE'}' &&
  git remote add origin https://$USERNAME:$TOKEN@github.com/$USERNAME/$PROJECT_NAME.git &&
  git push origin master

  if [[ $DOCKER_EXECUTION ]]; then
    cd ..
    chown 1000:1000 -R $PROJECT_NAME
  fi

  echo "Project successfully created!!"
  echo "📁  ./$PROJECT_NAME"
}
