#!/bin/bash
# shellcheck disable=SC2181

checkCommand() {
  if ! command -v "$1" >/dev/null; then
    echo "Error: $1 command required"
    exit 1
  fi
}

runFormula() {
  checkCommand git

  curl -X DELETE -u "$USERNAME":"$TOKEN" https://api.bitbucket.org/2.0/repositories/"$USERNAME"/"$PROJECT_NAME"

  if [ "$?" -ne 0 ]; then
    echo "✅ Repository successfully deleted from Bitbucket"
  else
    echo
    echo "Could not find repository with name $PROJECT_NAME"
  fi
}
