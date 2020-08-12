#!/bin/bash

runFormula() {
  XCODE_USER_TEMPLATES_DIR=~/Library/Developer/Xcode/Templates/File\ Templates

  rm -fR "$XCODE_USER_TEMPLATES_DIR"/Clean\ Swift
}
