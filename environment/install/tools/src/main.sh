#!/bin/bash

# shellcheck source=/dev/null
. "$(dirname "$0")"/formula/formula.sh --source-only

runFormula "$CONFIGURATION"
