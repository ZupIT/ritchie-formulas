#!/bin/bash

run() {
  echo "Updating repo local: $CURRENT_PWD"
  cd "$CURRENT_PWD" || exit
  status=$(git status | grep "nothing to commit")
  if [ "$status" != "nothing to commit, working tree clean" ] ;then
    echo -e "✘️ \e[91mError: \e[0mis Commit or discard changes in working directory";exit 1;
    exit 1;
  fi
  if $SETUPSTREAM; then
    read -r -p "Type original repo: (https://github.com/ORIGINAL_OWNER/ORIGINAL_REPOSITORY.git) " originalrepo
    git remote add upstream "$originalrepo"
  fi
  git checkout master
  git fetch upstream
  git rebase upstream/master master
  if $PUSH; then
    echo "Push"
    git push
  fi
}
