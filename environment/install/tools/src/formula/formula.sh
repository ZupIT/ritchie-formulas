#!/bin/sh

# shellcheck source=/dev/null
. "$(dirname "$0")"/formula/linux.sh --source-only

SYSTEM=$(uname -s)

runFormula() {
  # todo: Detect different versions of OS
  case "$SYSTEM" in
  Linux*)
<<<<<<< HEAD
    runConfigLinux "$SYSTEM" "$CONFIGURATION"
=======
    runConfigLinux "$SYSTEM" "$CONFIGURATION" "$GIT_NAME" "$GIT_EMAIL"
>>>>>>> Fixes for Shell syntax checking
    ;;
  Darwin*)
    config "MacOS"
    # todo: add MacOS
    ;;
  CYGWIN*)
    config "CYGWIN"
    # todo: add CYGWIN
    ;;
  *)
    printf "Unknown operating system.\\n"
    ;;
  esac
}
