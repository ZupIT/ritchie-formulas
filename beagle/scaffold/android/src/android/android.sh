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
  cat app/src/main/AndroidManifest.xml | sed -e "s,\${package_name},$1," -i app/src/main/AndroidManifest.xml

  cat app/build.gradle | sed -e "s,\${package_name},$1," -i app/build.gradle

  cat app/src/androidTest/java/$2/ExampleInstrumentedTest.kt | sed -e "s,\${package_name},$1," -i app/src/androidTest/java/$2/ExampleInstrumentedTest.kt
  
  cat app/src/main/java/$2/MainActivity.kt | sed -e "s,\${package_name},$1," -i app/src/main/java/$2/MainActivity.kt
  cat app/src/main/java/$2/activities/AppBeagleActivity.kt | sed -e "s,\${package_name},$1," -i app/src/main/java/$2/activities/AppBeagleActivity.kt
  cat app/src/main/java/$2/config/AppAplication.kt | sed -e "s,\${package_name},$1," -i app/src/main/java/$2/config/AppAplication.kt
  cat app/src/main/java/$2/config/AppBeagleConfig.kt | sed -e "s,\${package_name},$1," -i app/src/main/java/$2/config/AppBeagleConfig.kt

  cat app/src/test/java/$2/ExampleUnitTest.kt | sed -e "s,\${package_name},$1," -i app/src/test/java/$2/ExampleUnitTest.kt
}

run() {
  slug_name=$(createSlug "$PROJECT_NAME")

  checkProjectName $slug_name

  mkdir $CURRENT_PWD/$slug_name

  cp -r android/project/* $CURRENT_PWD/$slug_name

  cd $CURRENT_PWD/$slug_name
  
  cat app/src/main/res/values/strings.xml | sed -e "s,\${project_name},$PROJECT_NAME," -i app/src/main/res/values/strings.xml

  cat app/build.gradle | sed -e "s,\${beagle_version},$BEAGLE_VERSION," -i app/build.gradle

  formatted_package_name=$(formatPackageName $PACKAGE_NAME)
  createPackage $formatted_package_name

  cat app/src/main/java/$formatted_package_name/config/AppBeagleConfig.kt | sed -e "s,\${beagle_url},$BEAGLE_URL," -i app/src/main/java/$formatted_package_name/config/AppBeagleConfig.kt

  cat app/build.gradle | sed -e "s,\${min_sdk},$MIN_SDK," -i app/build.gradle

  cat app/build.gradle | sed -e "s,\${target_sdk},$TARGET_SDK," -i app/build.gradle

  replacePackageName $PACKAGE_NAME $formatted_package_name

  if [[ $DOCKER_EXECUTION ]]; then
    chown 1000:1000 -R $CURRENT_PWD/$slug_name
  fi

  echo "Project successfully created!!"
  echo "📁  ./$slug_name"
}
