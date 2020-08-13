#!/bin/sh

runFormula() {
  XCODE_USER_TEMPLATES_DIR=~/Library/Developer/Xcode/Templates/File\ Templates

  mkdir -p "$XCODE_USER_TEMPLATES_DIR" &&
  rm -fR "$XCODE_USER_TEMPLATES_DIR"/Clean\ Swift &&
  cp -R formula/_vendor/Clean\ Swift "$XCODE_USER_TEMPLATES_DIR"

  retval=$?
  if [ $retval != 0 ]; then
    echo "✘️ Fail to install templates";
    exit 1;
  else
    echo "✔ Templates succesfully installed";
  fi
}
