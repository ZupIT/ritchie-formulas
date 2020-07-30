#!/bin/bash

. $PWD/fork/fork.sh --source-only

run $BRANCH $PUSH $SETUPSTREAM
