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

formatPackageName() {
  echo "$1" | tr "." "/"
}

createPackage() {
  if [[ ! $1 == "br/com/basebeagle" ]]; then
    mkdir -p app/src/androidTest/java/$1
    cp -r app/src/androidTest/java/br/com/basebeagle/* app/src/androidTest/java/$1
    rm -rf app/src/androidTest/java/br/com/basebeagle

    mkdir -p app/src/main/java/$1
    cp -r app/src/main/java/br/com/basebeagle/* app/src/main/java/$1
    rm -rf app/src/main/java/br/com/basebeagle

    mkdir -p app/src/test/java/$1
    cp -r app/src/test/java/br/com/basebeagle/* app/src/test/java/$1
    rm -rf app/src/test/java/br/com/basebeagle

    find app/src/ -empty -type d -delete
  fi
}

replacePackageName() {
  package_name=$1
  shift

  local arr=("$@")
  for file in "${arr[@]}"
  do
    echo $file
    sed "s/\${package_name}/$package_name/" -i $file
  done
}

run() {
  slug_name=$(createSlug "$PROJECT_NAME")

  checkProjectName $slug_name

  mkdir $CURRENT_PWD/$slug_name

  cp -r android/project/* $CURRENT_PWD/$slug_name

  cd $CURRENT_PWD/$slug_name
  
  sed "s,\${project_name},$PROJECT_NAME," -i settings.gradle

  sed "s,\${project_name},$PROJECT_NAME," -i app/src/main/res/values/strings.xml

  sed "s,\${beagle_version},$BEAGLE_VERSION," -i app/build.gradle

  formatted_package_name=$(formatPackageName $PACKAGE_NAME)
  createPackage $formatted_package_name

  sed "s,\${beagle_url},$BEAGLE_URL," -i app/src/main/java/$formatted_package_name/config/AppBeagleConfig.kt

  sed "s,\${min_sdk},$MIN_SDK," -i app/build.gradle

  sed "s,\${target_sdk},$TARGET_SDK," -i app/build.gradle

  files=(
    app/src/main/AndroidManifest.xml
    app/build.gradle
    app/src/androidTest/java/$formatted_package_name/ExampleInstrumentedTest.kt
    app/src/main/java/$formatted_package_name/MainActivity.kt
    app/src/main/java/$formatted_package_name/activities/AppBeagleActivity.kt
    app/src/main/java/$formatted_package_name/config/AppAplication.kt
    app/src/main/java/$formatted_package_name/config/AppBeagleConfig.kt
    app/src/test/java/$formatted_package_name/ExampleUnitTest.kt
  )

  replacePackageName $PACKAGE_NAME "${files[@]}"

  if [[ $DOCKER_EXECUTION ]]; then
    chown 1000:1000 -R $CURRENT_PWD/$slug_name
  fi

  echo "Project successfully created!!"
  echo "📁  ./$slug_name"
}
