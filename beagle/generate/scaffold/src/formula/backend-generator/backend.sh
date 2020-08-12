#!/bin/bash

formatPackageName() {
  echo "$1" | tr "." "/"
}

createSpringPackage() {
  if [[ ! $1 == "com/base/bff/beagle/project/demo" ]]; then
    mkdir -p src/main/kotlin/"$1"
    cp -r src/main/kotlin/com/base/bff/beagle/project/demo/* src/main/kotlin/"$1"
    rm -rf src/main/kotlin/com/base/bff/beagle/project/demo

    mkdir -p src/test/kotlin/"$1"
    cp -r src/test/kotlin/com/base/bff/beagle/project/demo/* src/test/kotlin/"$1"
    rm -rf src/test/kotlin/com/base/bff/beagle/project/demo

    find src/ -empty -type d -delete
  fi
}

createMicronautPackage() {
  if [[ ! $1 == "com/bff/micronaut" ]]; then
    mkdir -p src/main/kotlin/"$1"
    cp -r src/main/kotlin/com/bff/micronaut/* src/main/kotlin/"$1"
    rm -rf src/main/kotlin/com/bff/micronaut

    mkdir -p src/test/kotlin/"$1"
    cp -r src/test/kotlin/com/bff/micronaut/* src/test/kotlin/"$1"
    rm -rf src/test/kotlin/com/bff/micronaut

    find src/ -empty -type d -delete
  fi
}

replacePackageName() {
  local package_name=$1
  shift

  local arr=("$@")
  for file in "${arr[@]}"
  do
    sed -i -e "s/\${package_name}/$package_name/" "$file"
  done
}

removeCors() {
  local framework=$1
  local package_name=$2
  local slug=$3
  local formatted_package_name

  formatted_package_name=$(formatPackageName "$package_name")

  cd "$CURRENT_PWD"/"$slug" || exit

  if [[ $1 == "spring" ]]; then
    rm -rf src/main/kotlin/"$formatted_package_name"/config
  else
    sed -i -e "1d" src/main/resources/application.properties
  fi
}

runBackend() {
  slug=$1
  package_name=$2
  jdk_version=$3
  kotlin_version=$4
  beagle_version=$5
  framework=$6

  mkdir "$CURRENT_PWD"/"$slug"

  if [[ $framework == "spring" ]]; then
    cp -r formula/backend-generator/_vendor/beagle-spring/. "$CURRENT_PWD"/"$slug"

    cd "$CURRENT_PWD"/"$slug" || exit

    sed -i -e "s/\${artifact_name}/$slug/" pom.xml
    sed -i -e "s/\${jdk}/$jdk_version/" pom.xml
    sed -i -e "s/\${kotlin_version}/$kotlin_version/" pom.xml
    sed -i -e "s/\${beagle_version}/$beagle_version/" pom.xml

    formatted_package_name=$(formatPackageName "$package_name")
    createSpringPackage "$formatted_package_name"

    files=(
      pom.xml
      src/main/kotlin/"$formatted_package_name"/config/CorsConfig.kt
      src/main/kotlin/"$formatted_package_name"/controller/MyController.kt
      src/main/kotlin/"$formatted_package_name"/service/MyService.kt
      src/main/kotlin/"$formatted_package_name"/BffBeagleSpringApplication.kt
      src/test/kotlin/"$formatted_package_name"/BffBeagleSpringApplicationTests.kt
    )

    replacePackageName "$package_name" "${files[@]}"
  else
    cp -r formula/backend-generator/_vendor/beagle-micronaut/. "$CURRENT_PWD"/"$slug"

    cd "$CURRENT_PWD"/"$slug" || exit

    sed -i -e "s/\${artifact_name}/$slug/" pom.xml
    sed -i -e "s/\${jdk}/$jdk_version/" pom.xml
    sed -i -e "s/\${kotlin_version}/$kotlin_version/" pom.xml
    sed -i -e "s/\${beagle_version}/$beagle_version/" pom.xml

    formatted_package_name=$(formatPackageName "$package_name")
    createMicronautPackage "$formatted_package_name"

    files=(
      micronaut-cli.yml
      pom.xml
      src/main/kotlin/"$formatted_package_name"/Application.kt
      src/main/kotlin/"$formatted_package_name"/MyController.kt
      src/main/kotlin/"$formatted_package_name"/MyService.kt
      src/test/kotlin/"$formatted_package_name"/BffMicronautTest.kt
    )

    replacePackageName "$package_name" "${files[@]}"
  fi

  if [[ $DOCKER_EXECUTION ]]; then
    chown 1000:1000 -R "$CURRENT_PWD"/"$slug"
  fi
}
