#!/bin/bash

# shellcheck source=/dev/null
. "$PWD"/fork/fork.sh --source-only

run "$BRANCH" "$PUSH" "$SETUPSTREAM"
