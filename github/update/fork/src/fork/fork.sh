#!/bin/bash

run() {
  checkCommand git
  echo "Updating repo local: $CURRENT_PWD"
  cd $CURRENT_PWD
  status=$(git status | grep "nothing to commit")
  if [ "$status" != "nothing to commit, working tree clean" ] ;then
    echo -e "✘️ \e[91mError: \e[0mis Commit or discard changes in working directory";exit 1;
    exit 1;
  fi
  if $SETUPSTREAM; then
    read -p "Type original repo: (https://github.com/ORIGINAL_OWNER/ORIGINAL_REPOSITORY.git) " originalrepo
    git remote add upstream $originalrepo
  fi
  git checkout master
  echo -e '\e[0;32m✔ \e[1;30mChecktou master\e[0m'
  git fetch upstream
  echo -e '\e[0;32m✔ \e[1;30mFetch upstream\e[0m'
  git rebase upstream/master master
  echo -e '\e[0;32m✔ \e[1;30mRebase upstream/master master\e[0m'
  if $PUSH; then
    git push
    echo -e '\e[0;32m✔ \e[1;30mPush\e[0m'
  fi
}


checkCommand () {
    if ! command -v $1 >/dev/null; then
      echo -e "✘️ \e[91mError: \e[33;1m$1 \e[0mis required";exit 1;
    fi
}
