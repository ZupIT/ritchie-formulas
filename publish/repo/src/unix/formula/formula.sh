#!/bin/bash
# shellcheck disable=SC2181
# shellcheck disable=SC2086
# shellcheck disable=SC2164

runFormula() {
  if [ "Github" == $PROVIDER ]
  then
    echo "🐙 Github provider selected"
#    echo '{"privacy":"'$PRIVACY'", "project_name":"'$PROJECT_NAME'", "workspace_path":"'$WORKSPACE_PATH'", "version":"'$VERSION'"}' | rit github publish repo --stdin
    rit github publish repo --privacy=$PRIVACY --project_name=$PROJECT_NAME --workspace_path=$WORKSPACE_PATH --version=$VERSION
  elif [ "Gitlab" == $PROVIDER ]
  then
    echo "🦊 Gitlab provider selected"
#    echo '{"privacy":"'$PRIVACY'", "project_name":"'$PROJECT_NAME'", "workspace_path":"'$WORKSPACE_PATH'", "version":"'$VERSION'"}' | rit gitlab publish repo --stdin
    rit gitlab publish repo --privacy=$PRIVACY --project_name=$PROJECT_NAME --workspace_path=$WORKSPACE_PATH --version=$VERSION
  else
    echo "🤖 Unexpected Provider informed. Check it please and try again."
  fi
}
