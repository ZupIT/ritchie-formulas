#!/bin/bash

# shellcheck source=/dev/null
. "$(dirname "$0")"/repo/repo.sh --source-only

run "$PROJECT_NAME" "$PROJECT_DESCRIPTION" "$PRIVATE $USERNAME" "$TOKEN"
