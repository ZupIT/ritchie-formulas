#!/bin/bash

# shellcheck source=/dev/null
. "$PWD"/repo/repo.sh --source-only

run "$PROJECT_NAME" "$PROJECT_DESCRIPTION" "$PRIVATE $USERNAME" "$TOKEN"