#!/bin/sh

runFormula() {
  XCODE_USER_TEMPLATES_DIR=~/Library/Developer/Xcode/Templates/File\ Templates

  rm -fR "$XCODE_USER_TEMPLATES_DIR"/Clean\ Swift

  retval=$?
  if [ $retval != 0 ]; then
    echo "✘️ Fail to uninstall templates";
    exit 1;
  else
    echo "✔ Templates succesfully uninstalled";
  fi
}
