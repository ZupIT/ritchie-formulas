#!/bin/bash

if expr "$CIRCLE_BRANCH" : 'qa' >/dev/null; then
  export ENVIRONMENT="qa"

elif expr "$CIRCLE_BRANCH" : 'master' >/dev/null; then
  export ENVIRONMENT="prod"
else
  echo ""
fi
