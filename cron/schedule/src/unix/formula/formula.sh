#!/bin/sh

runFormula() {
  echo "Hello World!"
  echoColor "cyan" "Command: $RIT_INPUT_COMMAND."
  echoColor "cyan" "Frequency: $RIT_INPUT_FREQUENCY."
  echoColor "cyan" "Day of Week: $RIT_INPUT_DAY_OF_WEEK."
  echoColor "cyan" "Day of Month: $RIT_INPUT_DAY_OF_MONTH."
  echoColor "cyan" "Hour: $RIT_INPUT_HOUR."


  #comando crontab -l | { cat; echo "00 09 * * 1-5 echo hello"; } | crontab -
  
  HOUR=${RIT_INPUT_HOUR:0:2} #get the first 2 characters
  MINUTES=${RIT_INPUT_HOUR: -2} #get the last 2 characters
  
  if [ -z "$RIT_INPUT_DAY_OF_MONTH" ]
    then
      RIT_INPUT_DAY_OF_MONTH=*
  fi

  case $RIT_INPUT_DAY_OF_WEEK in
    Sunday)
      DAY_OF_WEEK=0
      ;;
    Monday)
      DAY_OF_WEEK=1
      ;;
    Tuesday)
      DAY_OF_WEEK=2
      ;;
    Wednesday)
      DAY_OF_WEEK=3
      ;;  
    Thursday)
      DAY_OF_WEEK=4
      ;;
    Friday)
      DAY_OF_WEEK=5
      ;;
    Saturday)
      DAY_OF_WEEK=6
      ;;
    *)
      DAY_OF_WEEK=*
      ;;
  esac
  

  #echoColor "yellow" "Command: $MINUTES $HOUR $RIT_INPUT_DAY_OF_MONTH * $DAY_OF_WEEK $RIT_INPUT_COMMAND"
  
  crontab -l | { cat; echo "$MINUTES $HOUR $RIT_INPUT_DAY_OF_MONTH * $DAY_OF_WEEK $RIT_INPUT_COMMAND"; } | crontab -
}

echoColor() {
  case $1 in
    red)
      echo "$(printf '\033[31m')$2$(printf '\033[0m')"
      ;;
    green)
      echo "$(printf '\033[32m')$2$(printf '\033[0m')"
      ;;
    yellow)
      echo "$(printf '\033[33m')$2$(printf '\033[0m')"
      ;;
    blue)
      echo "$(printf '\033[34m')$2$(printf '\033[0m')"
      ;;
    cyan)
      echo "$(printf '\033[36m')$2$(printf '\033[0m')"
      ;;
    esac
}
