#!/bin/bash

. $PWD/backend/backend.sh --source-only

run $PROJECT_NAME $PACKAGE_NAME $JDK $KOTLIN_VERSION $BEAGLE_VERSION $BFF_URL $CORS