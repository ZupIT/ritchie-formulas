#!/bin/bash

checkCommand() {
  if ! command -v $1 >/dev/null; then
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

  cd $CURRENT_PWD

  PROJECT_NAME=$(checkSpace "$PROJECT_NAME")

  checkProjectName $PROJECT_NAME
  
  mkdir $PROJECT_NAME
  cd $PROJECT_NAME
  
  git init
  echo $PROJECT_DESCRIPTION >> README.md
  
  if [[ $EXECUTION_DOCKER ]]; then
    read -p "Enter your email: " email
    git config --local user.name $USERNAME
    git config --local user.email $email
  fi

  git add .
  git commit -m "Initial Commit"
  
  curl -H 'Authorization: token '$TOKEN https://api.github.com/user/repos -d '{"name":"'$PROJECT_NAME'", "private":'$PRIVATE'}'
  
  git push https://$USERNAME:$TOKEN@github.com/$USERNAME/$PROJECT_NAME.git HEAD

  git remote add origin https://github.com/$USERNAME/$PROJECT_NAME
}
