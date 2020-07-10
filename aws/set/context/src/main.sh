#!/bin/bash

# shellcheck source=/dev/null
. "$PWD"/setcontext/setcontext.sh --source-only

run "$CONTEXT_NAME"
