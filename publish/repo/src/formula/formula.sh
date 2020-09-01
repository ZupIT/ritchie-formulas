#!/bin/sh
runFormula() {
  if [ "Github" == $PROVIDER ] 
  then
    echo "🐙 Github provider selected"
    echo '{"privacy":"'$PRIVACY'", "project_name":"'$PROJECT_NAME'", "workspace_path":"'$WORKSPACE_PATH'", "version":"'$VERSION'"}' | rit github publish repo --stdin
  elif [ "Gitlab" == $PROVIDER ]
  then
    echo "🦊 Gitlab provider selected"
    echo '{"privacy":"'$PRIVACY'", "project_name":"'$PROJECT_NAME'", "workspace_path":"'$WORKSPACE_PATH'", "version":"'$VERSION'"}' | rit gitlab publish repo --stdin
  else
    echo "🤖 Unexpected Provider informed. Check it please and try again."
  fi
}
