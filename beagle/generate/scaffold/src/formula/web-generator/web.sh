#!/bin/bash

runWeb() {
  slug=$1
  beagle_version=$2
  framework=$3

  mkdir "$CURRENT_PWD"/"$slug"

  if [[ $framework == "react" ]]; then
    cp -r formula/web-generator/beagle-react/* "$CURRENT_PWD"/"$slug"

    cd "$CURRENT_PWD"/"$slug" || exit
  else
    cp -r formula/web-generator/beagle-angular/* "$CURRENT_PWD"/"$slug"

    cd "$CURRENT_PWD"/"$slug" || exit
  fi

  sed -i -e "s,\${project_name},$slug," package.json
  sed -i -e "s,\${beagle_version},$beagle_version," package.json

  if [[ $DOCKER_EXECUTION ]]; then
    chown 1000:1000 -R "$CURRENT_PWD"/"$slug"
  fi
}
