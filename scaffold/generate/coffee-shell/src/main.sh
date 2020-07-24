#!/bin/bash

# shellcheck source=/dev/null
. "$(dirname "$0")"/coffee/coffee.sh --source-only
#In sh for receive inputs of CLI use: $NAME, $COFFEE_TYPE and $DELIVERY for this exemple

run "$NAME" "$COFFEE_TYPE" "$DELIVERY"
