#!/bin/bash

run() {
  checkCommand git
  echo -e "\\e[0;32m✔ \\e[1;30mUpdating repo local: "$CURRENT_PWD"\\e[0m";
  echo -e "\\e[0;32m✔ \\e[1;30mUpdating branch: "$BRANCH"\\e[0m";
  cd "$CURRENT_PWD" || exit
  status=$(git status | grep "nothing to commit\|nada a submeter")
  if [ "$status" != "nothing to commit, working tree clean" ] && [ "$status" != "nada a submeter, diretório de trabalho vazio" ] ;then
    echo -e "✘️ \\e[91mError:\\e[0m Commit or discard changes in working directory";
    exit 1;
  fi

  if $SETUPSTREAM; then
    read -p "Type original repo: (https://github.com/ORIGINAL_OWNER/ORIGINAL_REPOSITORY.git) " originalrepo
    git remote add upstream "$originalrepo"
    if $? != 0; then
      echo -e "✘️ \\e[91mError: \\e[0mFail $ git remote add upstream $originalrepo";
      exit 1;
    else
      echo -e "\\e[0;32m✔ \\e[1;30mRemote upstream added\\e[0m";
    fi
  fi

  git checkout "$BRANCH"
  if [ $? != 0 ]; then
    echo -e "✘️ \\e[91mError: \\e[0mFail $ git checkout $BRANCH";
    exit 1;
  else
    echo -e "\\e[0;32m✔ \\e[1;30mCheckout $BRANCH done\\e[0m";
  fi

  git fetch upstream
  if [ $? != 0 ]; then
    echo -e "✘️ \\e[91mError:\\e[0mFail $ git fetch upstream";
    exit 1;
  else
    echo -e "\\e[0;32m✔ \\e[1;30mFetch upstream done\\e[0m";
  fi

  git rebase upstream/"$BRANCH" "$BRANCH"
  if [ $? != 0 ]; then
    echo -e "✘️ \\e[91mError: \\e[0mFail $ git rebase upstream/$BRANCH $BRANCH";
    exit 1;
  else
    echo -e "\\e[0;32m✔ \\e[1;30mRebase upstream/$BRANCH $BRANCH done\\e[0m";
  fi

  if $PUSH; then
    git push
    if [ $? != 0 ]; then
      echo -e "✘️ \\e[91mError: \\e[0mFail $ git push";
      exit 1;
    else
      echo -e "\\e[0;32m✔ \\e[1;30mPush done\\e[0m";
    fi
  fi
}


checkCommand () {
    if ! command -v "$1" >/dev/null; then
      echo -e "✘️ \\e[91mError: \\e[33;1m$1 \\e[0mis required";
      exit 1;
    fi
}