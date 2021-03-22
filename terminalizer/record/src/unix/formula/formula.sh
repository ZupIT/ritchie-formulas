#!/bin/sh

runFormula() {
  if [ "$RIT_PERSONALIZE" == "yes" ]; then
    if [ "$RIT_HOW" == "config file" ]; then
      terminalizer record "$RIT_GIF_NAME" -c "$RIT_CONFIG_PATH"
    else
      terminalizer config

      sed -i -e "s/cols: auto/cols: $RIT_WIDTH/" config.yml
      sed -i -e "s/rows: auto/rows: $RIT_HEIGHT/" config.yml
      sed -i -e "s/title: Terminalizer/title: $RIT_TITLE/" config.yml

      if [ $RIT_CWD ]; then
        sed -i -e "s,cwd: null,cwd: $RIT_CWD," config.yml
      else
        sed -i -e "s,cwd: null,cwd: $CURRENT_PWD," config.yml
      fi

      sed -i -e "s/cursorStyle: block/cursorStyle: $RIT_CURSOR_STYLE/" config.yml

      terminalizer record "$RIT_GIF_NAME" -c config.yml
    fi
  else
    terminalizer record "$RIT_GIF_NAME"
  fi

  sed -i -e "s/$USER/$RIT_USERNAME/g" "$RIT_GIF_NAME".yml
  sed -i -e "s/$HOSTNAME/$RIT_HOSTNAME/g" "$RIT_GIF_NAME".yml

  if [ "$RIT_RENDER" == "yes" ]; then
    terminalizer render "$RIT_GIF_NAME".yml -o "$RIT_GIF_NAME".gif
  fi
}
