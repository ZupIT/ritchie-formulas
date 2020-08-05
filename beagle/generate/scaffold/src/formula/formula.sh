#!/bin/bash

. "$(dirname "$0")"/formula/backend-generator/backend.sh --source-only
. "$(dirname "$0")"/formula/web-generator/web.sh --source-only
. "$(dirname "$0")"/formula/mobile-generator/mobile.sh --source-only

removeSpaces() {
  echo "${1}" | xargs | tr " " -
}

createSlug() {
  local tmp="$1"

  if [[ "$1" = *" "* ]]; then
    tmp=$(removeSpaces "$1")
  fi

  echo "$tmp" | tr '[:upper:]' '[:lower:]'
}

binaryRead() {
  local var
  while true; do
    read -p "$1" var
    if [[ "$var" == "y" || "$var" == "Y" || "$var" == "n" || "$var" == "N" ]]; then
      echo "$var" | tr '[:upper:]' '[:lower:]'
      break
    else
      echo >&2 "Please, enter with 'y' or 'n'"
    fi
  done
}

simpleRead() {
  local var
  while true; do
    read -p "$1" var
    if [[ -z "$var" ]]; then
      echo >&2 "Please, input must not be empty"
    else
      echo "$var"
      break
    fi
  done
}

readWithDefaultValue() {
  local var
  while true; do
    read -p "$1" var
    if [[ -z "$var" ]]; then
      echo >&2 "$2"
      echo "$2"
      break
    else
      echo "$var"
      break
    fi
  done
}

readProjectName() {
  local project_name
  while true; do
    read -p "$1" project_name
    local slug=$(createSlug "$project_name")
    if [[ ! "$slug" =~ ^[a-zA-Z0-9-]+$ ]]; then
      echo >&2 "Project name cannot contain special characters"
    else
      reply=("$project_name" "$slug")
      break
    fi
  done
}

readTwoOptions() {
  local var
  while true; do
    read -p "Choose [$1/$2]: " var
    if [[ $var == $1 || $var == $2 ]]; then
      echo "$var"
      break
    else
      echo >&2 "Please, enter with '$1' or '$2'"
    fi
  done
}

createBackendProject() {
  readProjectName "Backend project name: "
  local project_name=${reply[0]}
  backend_slug=${reply[1]}
  backend_package_name=$(simpleRead "Package name (ex: com.example): ")
  local jdk_version=$(readWithDefaultValue "JDK version(8+) (default: 13): " "13")
  local kotlin_version=$(readWithDefaultValue "Kotlin version(1.3+) (default: 1.3.72): " "1.3.72")
  local beagle_version=$(readWithDefaultValue "Beagle version (default: 1.0.2): " "1.0.2")
  backend_framework=$(readTwoOptions "spring" "micronaut")

  $(runBackend "$project_name" $backend_slug $backend_package_name $jdk_version $kotlin_version $beagle_version $backend_framework)
}

createWebProject() {
  readProjectName "Web project name: "
  local project_name=${reply[0]}
  web_slug=${reply[1]}
  local beagle_version=$(readWithDefaultValue "Beagle version (default: 1.1.0): " "1.1.0")
  local framework=$(readTwoOptions "react" "angular")

  $(runWeb "$project_name" $web_slug $beagle_version $framework)
}

createMobileProject() {
  readProjectName "Mobile project name: "
  local project_name=${reply[0]}
  mobile_slug=${reply[1]}
  local os=$(readTwoOptions "ios" "android")

  if [[ $os == "ios" ]]; then
    local organization_name=$(simpleRead "Organization name: ")
    local organization_id=$(simpleRead "Organizaion ID: ")
    local beagle_version=$(readWithDefaultValue "Beagle version (ex: 1.0.0, default: latest): " "latest")
    local bff_url=$(readWithDefaultValue "BFF url (default: http://localhost:8080): " "http://localhost:8080")
    local sourcery=$(binaryRead "Do you want to use Sourcery? [y/n]: ")

    $(runMobile $os "$project_name" $mobile_slug "$organization_name" $organization_id $beagle_version $bff_url $sourcery)
  else
    local package_name=$(simpleRead "Package name (ex: com.example): ")
    local min_sdk=$(readWithDefaultValue "Min version SDK android: (default: 21): " "21")
    local target_sdk=$(readWithDefaultValue "Target version SDK: (default: 29): " "21")
    local kotlin_version=$(readWithDefaultValue "Kotlin version(1.3+) (default: 1.3.72): " "1.3.72")
    local beagle_version=$(readWithDefaultValue "Beagle version (default: 1.0.0): " "1.0.0")
    local bff_url=$(readWithDefaultValue "BFF url (default: http://localhost:8080): " "http://localhost:8080")

    $(runMobile $os "$project_name" $mobile_slug $package_name $min_sdk $target_sdk $kotlin_version $beagle_version $bff_url)
  fi
}

removeCorsFromBackend() {
  $(removeCors $backend_framework $backend_package_name $backend_slug)
}

printResult() {
  if [[ $backend_slug ]]; then
    echo "Backend project successfully created!!"
    echo "📁  ./$backend_slug"
    echo
  fi

  if [[ $web_slug ]]; then
    echo "Web project successfully created!!"
    echo "📁  ./$web_slug"
    echo
  fi

  if [[ $mobile_slug ]]; then
    echo "mobile project successfully created!!"
    echo "📁  ./$mobile_slug"
    echo
  fi
}

runFormula() {
  createBackendProject

  echo

  local answer=$(binaryRead "Do you want to create a web project? [y/n]: ")
  if [[ $answer == "y" ]]; then
    createWebProject

    echo

    answer=$(binaryRead "Do you want to create a mobile project? [y/n]: ")
    if [[ $answer == "y" ]]; then
      createMobileProject
    fi
  else
    answer=$(binaryRead "Do you want to create a mobile project? [y/n]: ")
    if [[ $answer == "y" ]]; then
      createMobileProject
    fi

    removeCorsFromBackend
  fi

  echo

  printResult
}
