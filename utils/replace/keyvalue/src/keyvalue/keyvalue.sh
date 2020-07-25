#!/bin/sh
run() {

  case $FORMAT in
    yml)
      SED_ARG="s,$KEY.*$,$KEY: $VALUE,g";;
    hcl)
      SED_ARG="s,$KEY.*$,$KEY=$VALUE,g";;
  esac

  if [ $(uname -s) == "Darwin" ]
  then
    find $CURRENT_PWD -type f | xargs sed -i "" "$SED_ARG"
  else
    find $CURRENT_PWD -type f | xargs sed -i "$SED_ARG"
  fi

  echo "\033[0;32mDone!"
}
