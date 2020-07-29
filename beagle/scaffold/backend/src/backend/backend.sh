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

createSpringPackage() {
  if [[ ! $1 == "com/base/bff/beagle/project/demo" ]]; then
    mkdir -p src/main/kotlin/$1
    cp -r src/main/kotlin/com/base/bff/beagle/project/demo/* src/main/kotlin/$1
    rm -rf src/main/kotlin/com/base/bff/beagle/project/demo

    mkdir -p src/test/kotlin/$1
    cp -r src/test/kotlin/com/base/bff/beagle/project/demo/* src/test/kotlin/$1
    rm -rf src/test/kotlin/com/base/bff/beagle/project/demo

    find src/ -empty -type d -delete
  fi
}

createMicronautPackage() {
  if [[ ! $1 == "com/bff/micronaut" ]]; then
    mkdir -p src/main/kotlin/$1
    cp -r src/main/kotlin/com/bff/micronaut/* src/main/kotlin/$1
    rm -rf src/main/kotlin/com/bff/micronaut

    mkdir -p src/test/kotlin/$1
    cp -r src/test/kotlin/com/bff/micronaut/* src/test/kotlin/$1
    rm -rf src/test/kotlin/com/bff/micronaut

    find src/ -empty -type d -delete
  fi
}

replacePackageName() {
  package_name=$1
  shift

  local arr=("$@")
  for file in "${arr[@]}"
  do
    sed "s/\${package_name}/$package_name/" -i $file
  done
}

run() {
  slug_name=$(createSlug "$PROJECT_NAME")

  checkProjectName $slug_name

  mkdir $CURRENT_PWD/$slug_name

  if [[ $FRAMEWORK == "Spring" ]]; then
    cp -r backend/beagle-spring/. $CURRENT_PWD/$slug_name

    cd $CURRENT_PWD/$slug_name

    sed "s/\${artifact_name}/$slug_name/" -i pom.xml
    sed "s/\${jdk}/$JDK/" -i pom.xml
    sed "s/\${kotlin_version}/$KOTLIN_VERSION/" -i pom.xml
    sed "s/\${beagle_version}/$BEAGLE_VERSION/" -i pom.xml

    formatted_package_name=$(formatPackageName $PACKAGE_NAME)
    createSpringPackage $formatted_package_name

    files=(
      pom.xml
      src/main/kotlin/$formatted_package_name/config/CorsConfig.kt
      src/main/kotlin/$formatted_package_name/controller/MyController.kt
      src/main/kotlin/$formatted_package_name/service/MyService.kt
      src/main/kotlin/$formatted_package_name/BffBeagleSpringApplication.kt
      src/test/kotlin/$formatted_package_name/BffBeagleSpringApplicationTests.kt
    )

    replacePackageName $PACKAGE_NAME "${files[@]}"

    if [[ $CORS == "true" ]]; then
      read -p "Allowed origin url: (ex: https://localhost:3000) " url
      sed "s,\${cors_url},$url," -i src/main/kotlin/$formatted_package_name/config/CorsConfig.kt
    else
      rm -rf src/main/kotlin/$formatted_package_name/config
    fi
  else
    cp -r backend/beagle-micronaut/. $CURRENT_PWD/$slug_name

    cd $CURRENT_PWD/$slug_name

    sed "s/\${app_name}/$slug_name/" -i src/main/resources/application.yml
    sed "s/\${artifact_name}/$slug_name/" -i pom.xml
    sed "s/\${jdk}/$JDK/" -i pom.xml
    sed "s/\${kotlin_version}/$KOTLIN_VERSION/" -i pom.xml
    sed "s/\${beagle_version}/$BEAGLE_VERSION/" -i pom.xml

    formatted_package_name=$(formatPackageName $PACKAGE_NAME)
    createMicronautPackage $formatted_package_name

    files=(
      micronaut-cli.yml
      pom.xml
      src/main/kotlin/$formatted_package_name/Application.kt
      src/main/kotlin/$formatted_package_name/MyController.kt
      src/main/kotlin/$formatted_package_name/MyService.kt
      src/test/kotlin/$formatted_package_name/BffMicronautTest.kt
    )

    replacePackageName $PACKAGE_NAME "${files[@]}"

    if [[ $CORS == "true" ]]; then
      read -p "Allowed origin url: (ex: https://localhost:3000) " url
      sed "s,\${cors_url},$url," -i src/main/resources/application.yml
    else
      sed "s/cors://" -i src/main/resources/application.yml
      sed "s/configurations://" -i src/main/resources/application.yml
      sed "s/allowedOrigins: \${cors_url}//" -i src/main/resources/application.yml
      sed "s/allowedMethods: GET, PUT, POST//" -i src/main/resources/application.yml
      sed "s/allowedHeaders: Cache-Control//" -i src/main/resources/application.yml
    fi
  fi

  if [[ $DOCKER_EXECUTION ]]; then
    chown 1000:1000 -R $CURRENT_PWD/$slug_name
  fi

  echo "Project successfully created!!"
  echo "📁  ./$slug_name"
}
