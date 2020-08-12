#!/bin/bash

runFormula() {
  XCODE_USER_TEMPLATES_DIR=~/Library/Developer/Xcode/Templates/File\ Templates

  mkdir -p "$XCODE_USER_TEMPLATES_DIR"
  rm -fR "$XCODE_USER_TEMPLATES_DIR"/Clean\ Swift
  cp -R formula/clean-swift "$XCODE_USER_TEMPLATES_DIR"
}
