#!/bin/bash

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

run() {
  slug_name=$(createSlug "$PROJECT_NAME")

  checkProjectName $slug_name

  mkdir $CURRENT_PWD/$slug_name

  if [[ $FRAMEWORK == "ReactJS" ]]; then
    cp -r web/beagle-react/* $CURRENT_PWD/$slug_name

    cd $CURRENT_PWD/$slug_name

    sed "s,\${bff_url},$BFF_URL," -i src/beagle/beagle-service.ts
  else
    cp -r web/beagle-angular/* $CURRENT_PWD/$slug_name

    cd $CURRENT_PWD/$slug_name

    sed "s,\${bff_url},$BFF_URL," -i src/app/beagle.module.ts
  fi

  sed "s,\${project_name},$slug_name," -i package.json
  sed "s,\${beagle_version},$BEAGLE_VERSION," -i package.json

  if [[ $DOCKER_EXECUTION ]]; then
    chown 1000:1000 -R $CURRENT_PWD/$slug_name
  fi

  echo "Project successfully created!!"
  echo "📁  ./$slug_name"
}
