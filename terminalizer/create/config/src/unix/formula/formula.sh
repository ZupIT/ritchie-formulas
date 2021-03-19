#!/bin/sh

runFormula() {
  terminalizer config

  sed -i -e "s/cols: auto/cols: $RIT_WIDTH/" config.yml
  sed -i -e "s/rows: auto/rows: $RIT_HEIGHT/" config.yml
  sed -i -e "s/title: Terminalizer/title: $RIT_TITLE/" config.yml

  if [ $RIT_CWD ]; then
    sed -i -e "s,cwd: null,cwd: $RIT_CWD," config.yml
  fi

  sed -i -e "s/cursorStyle: block/cursorStyle: $RIT_CURSOR_STYLE/" config.yml
}
