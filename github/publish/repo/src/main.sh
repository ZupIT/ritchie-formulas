#!/bin/bash

# shellcheck source=/dev/null
. "$(dirname "$0")"/formula/formula.sh --source-only

runFormula "$USERNAME" "$TOKEN" "$EMAIL" "$PRIVACY" "$PROJECT_NAME" "$WORKSPACE_PATH" "$VERSION"
