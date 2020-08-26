#!/bin/bash
# shellcheck disable=SC2181
# shellcheck disable=SC2086
# shellcheck disable=SC2164

removeSpaces() {
  echo "${1}" | xargs | tr " " -
}

createSlug() {
  tmp="$1"

  if [[ "$1" = *" "* ]]; then
    echo >&2 "Removing spaces from Project name..."
    tmp=$(removeSpaces "$1")
    echo >&2 "Project name without spaces: $tmp"
  fi

  echo "$tmp" | tr '[:upper:]' '[:lower:]'
}

checkProjectName() {
  if [[ ! "$1" =~ ^[a-zA-Z0-9-]+$ ]]; then
    echo "Project name cannot contain special characters"
    exit 1
  fi
}

runFormula() {

  slug_name=$(createSlug "$PROJECT_NAME")
  checkProjectName $slug_name

  if [[ $DOCKER_EXECUTION ]]; then
    git config --local user.name $USERNAME
  fi

  echo "---------------------------------------------------------------------------"

  echo "📡 Checking if Github url https://github.com/$USERNAME/$slug_name.git exists."

  git ls-remote "https://github.com/$USERNAME/$slug_name.git" > /dev/null 2>&1

  if [ "$?" -ne 0 ]; then
    sleep 1s
    echo "🚨 Unable to read from https://github.com/$USERNAME/$slug_name.git"
    exit 1;
  else
    echo "🐙 https://github.com/$USERNAME/$slug_name.git exists."
    echo "🚧 Start deleting https://github.com/$USERNAME/$slug_name.git repository."
    sleep 1s
    curl -X DELETE -H 'Authorization: token '$TOKEN https://api.github.com/repos/$USERNAME/$slug_name > /dev/null
  fi

  git ls-remote "https://github.com/$USERNAME/$slug_name.git" > /dev/null 2>&1

  if [ "$?" -ne 0 ]; then
    sleep 1s
    echo "✅ Repository successfully deleted from Github"
  else
    sleep 1s
    echo "🚨 Unable to delete https://github.com/$USERNAME/$slug_name.git repository"
    echo "🔧 Check your github token authorizations."
  fi
}
