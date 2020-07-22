#!/bin/bash

. $PWD/android/android.sh --source-only

run $PROJECT_NAME $PACKAGE_NAME $BEAGLE_VERSION $BEAGLE_URL $TARGET_SDK $MIN_SDK 