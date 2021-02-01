#!/bin/bash

run() {
  checkCommand git
  echo -e "\\e[0;32m✔ \\e[1;30mUpdating repo local: $CURRENT_PWD\\e[0m";
  echo -e "\\e[0;32m✔ \\e[1;30mUpdating branch: $BRANCH\\e[0m";
  cd "$CURRENT_PWD" || exit
  status=$(git status | grep "nothing to commit\\|nada a submeter")
  if [ "$status" != "nothing to commit, working tree clean" ] && [ "$status" != "nada a submeter, diretório de trabalho vazio" ] ;then
    echo -e "✘️ \\e[91mError:\\e[0m Commit or discard changes in working directory";
    exit 1;
  fi

  if $SETUPSTREAM; then
    echo -n "Type original repo: (https://repository-site.com/ORIGINAL_OWNER/ORIGINAL_REPOSITORY.git) "
    read -r originalrepo
    if git remote add upstream "$originalrepo"; then
      echo -e "\\e[0;32m✔ \\e[1;30mRemote upstream added\\e[0m";
    else
      echo -e "✘️ \\e[91mError: \\e[0mFail $ git remote add upstream $originalrepo";
      exit 1;
    fi
  fi

  if git checkout "$BRANCH"; then
    echo -e "\\e[0;32m✔ \\e[1;30mCheckout $BRANCH done\\e[0m";
  else
    echo -e "✘️ \\e[91mError: \\e[0mFail $ git checkout $BRANCH";
    exit 1;
  fi

  if git fetch upstream; then
    echo -e "\\e[0;32m✔ \\e[1;30mFetch upstream done\\e[0m";
  else
    echo -e "✘️ \\e[91mError:\\e[0mFail $ git fetch upstream";
    exit 1;
  fi

  if git rebase upstream/"$BRANCH" "$BRANCH"; then
    echo -e "\\e[0;32m✔ \\e[1;30mRebase upstream/$BRANCH $BRANCH done\\e[0m";
  else
    echo -e "✘️ \\e[91mError: \\e[0mFail $ git rebase upstream/$BRANCH $BRANCH";
    exit 1;
  fi

  if $PUSH; then
    if git push; then
      echo -e "\\e[0;32m✔ \\e[1;30mPush done\\e[0m";
    else
      echo -e "✘️ \\e[91mError: \\e[0mFail $ git push";
      exit 1;
    fi
  fi
}

checkCommand () {
    if ! command -v "$1" >/dev/null; then
      echo -e "✘️ \\e[91mError: \\e[33;1m$1 \\e[0mis required";
      exit 1;
    fi
}
