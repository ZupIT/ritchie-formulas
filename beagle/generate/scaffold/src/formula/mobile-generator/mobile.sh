#!/bin/bash

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
  local package_name=$1
  shift

  local arr=("$@")
  for file in "${arr[@]}"
  do
    sed -i -e "s/\${package_name}/$package_name/" $file
  done
}

formatHeaders() {
  local project_name=$1
  local organization_name=$2
  local complete_date=$(date +'%d/%m/%Y')
  local year=$(date +'%Y')
  shift
  shift

  local arr=("$@")
  for file in "${arr[@]}"
  do
    sed -i -e "s/\${project_name}/$project_name/" $file
    sed -i -e "s/\${organization_name}/$organization_name/" $file
    sed -i -e "s,\${date},$complete_date," $file
    sed -i -e "s,\${year},$year," $file
  done
}

runMobile() {
  os=$1
  project_name=$2
  slug=$3

  mkdir $CURRENT_PWD/$slug

  if [[ $os == "ios" ]]; then
    organization_name=$4
    organization_id=$5
    beagle_version=$6
    bff_url=$7
    sourcery=$8
    dependecy_management=$9

    cp -r formula/mobile-generator/ios/. $CURRENT_PWD/$slug

    cd $CURRENT_PWD/$slug

    mv cocoapods $slug
    mv cocoapods.xcodeproj $slug.xcodeproj

    sed -i -e "s,\${project_name},$slug,g" $slug.xcodeproj/project.pbxproj
    
    sed -i -e "s,\${organization_name},$organization_name," $slug.xcodeproj/project.pbxproj

    sed -i -e "s,\${organization_id},$organization_id," $slug.xcodeproj/project.pbxproj

    sed -i -e "s,\${bff_url},$bff_url," $slug/Commom/Constants.swift

    sed -i -e "s,\${project_name},$slug," Podfile

    sed -i -e "s,\${project_name},$slug," .sourcery.yml

    if [[ $beagle_version != "latest" ]]; then
      sed -i -e "s/pod 'Beagle'/pod 'Beagle', '$beagle_version'/" Podfile
    fi

    if [[ $sourcery == "n" ]]; then
      rm .sourcery.yml
      sed -i -e "159,179d" $slug.xcodeproj/project.pbxproj # removing lines 159-179
      sed -i -e "99d" $slug.xcodeproj/project.pbxproj # removing line 99
    fi

    files=(
      $slug/ViewController.swift
      $slug/SceneDelegate.swift
      $slug/AppDelegate.swift
      $slug/Commom/Constants.swift
      $slug/BeagleConfig/BeagleConfig.swift
    )

    formatHeaders "$project_name" "$organization_name" "${files[@]}"
  else
    package_name=$4
    min_sdk=$5
    target_sdk=$6
    kotlin_version=$7
    beagle_version=$8
    bff_url=$9

    cp -r formula/mobile-generator/android/* $CURRENT_PWD/$slug

    cd $CURRENT_PWD/$slug

    sed -i -e "s,\${kotlin_v},$kotlin_version," build.gradle

    sed -i -e "s,\${project_name},$project_name," settings.gradle

    sed -i -e "s,\${project_name},$project_name," app/src/main/res/values/strings.xml

    sed -i -e "s,\${beagle_version},$beagle_version," app/build.gradle

    formatted_package_name=$(formatPackageName $package_name)
    createPackage $formatted_package_name

    sed -i -e "s,\${bff_url},$bff_url," app/src/main/java/$formatted_package_name/config/AppBeagleConfig.kt

    sed -i -e "s,\${min_sdk},$min_sdk," app/build.gradle

    sed -i -e "s,\${target_sdk},$target_sdk," app/build.gradle

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

    replacePackageName $package_name "${files[@]}"
  fi

  if [[ $DOCKER_EXECUTION ]]; then
    chown 1000:1000 -R $CURRENT_PWD/$slug
  fi
}