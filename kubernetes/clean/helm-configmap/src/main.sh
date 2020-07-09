#!/bin/bash

# shellcheck source=/dev/null
. "$PWD"/helm-configmap/helm-configmap.sh --source-only

run "$SAMPLE_TEXT" "$SAMPLE_LIST" "$SAMPLE_BOOL"