#!/bin/bash
<<<<<<< HEAD
=======
# shellcheck disable=SC2181
# shellcheck disable=SC2086
# shellcheck disable=SC2164
>>>>>>> upstream/master

removeSpaces() {
  echo "${1}" | xargs | tr " " -
}

createSlug() {
  tmp="$1"

  if [[ "$1" = *" "* ]]; then
    echo >&2 "Removing spaces from Project name..."
    tmp=$(removeSpaces "$1")
    echo >&2 "Project name without spaces: $tmp"
  fi

  echo "$tmp" | tr '[:upper:]' '[:lower:]'
}

checkProjectName() {
  if [[ ! "$1" =~ ^[a-zA-Z0-9-]+$ ]]; then
    echo "Project name cannot contain special characters"
    exit 1
  fi
}

formatPackageName() {
  echo "$1" | tr "." "/"
}

runFormula() {

  slug_name=$(createSlug "$PROJECT_NAME")

  checkProjectName $slug_name

  cd $WORKSPACE_PATH
<<<<<<< HEAD
=======
  if [ $? != 0 ]; then
      cd -
      exit 1;
  fi
>>>>>>> upstream/master

  echo "---------------------------------------------------------------------------"

  if git rev-parse --git-dir > /dev/null 2>&1; then
    echo "🚧 This repository already exists. Preparing new commit..."
    git add .
    git commit -m "$VERSION Commit"
  else
    echo "🚥 Repository creation ($PRIVACY). Preparing first commit..."
    git init
      if [[ $DOCKER_EXECUTION ]]; then
        git config --local user.name $USERNAME
        git config --local user.email $EMAIL
      fi
    git add .
    git commit -m "Initial Commit"
    curl -H 'Authorization: token '$TOKEN https://api.github.com/user/repos -d '{"name":"'$slug_name'", "private": '$PRIVACY'}'
    if [ $? != 0 ]; then
      echo -e "✘️ \\e[91mError:\\e[0mFail creating $PRIVACY repository";
      exit 1;
    fi
    git remote add origin https://$USERNAME:$TOKEN@github.com/$USERNAME/$slug_name.git
  fi

  git push origin master

  if [[ $DOCKER_EXECUTION ]]; then
    chown 1000:1000 -R $CURRENT_PWD/$slug_name
  fi

  echo "---------------------------------------------------------------------------"
  echo "✅ Project successfully added on Github!!"
  echo "📎 Run: $ git clone https://github.com/$USERNAME/$slug_name.git"
  sleep 1s

  echo "---------------------------------------------------------------------------"
  echo "🛠 Generating release $VERSION"
  API_JSON=$(printf '{"tag_name": "%s","target_commitish": "master","name": "%s","body": "Release of version %s","draft": false,"prerelease": false}' $VERSION $VERSION $VERSION)
  curl --data "$API_JSON" https://api.github.com/repos/$USERNAME/$slug_name/releases?access_token=$TOKEN
  if [ $? != 0 ]; then
      echo -e "✘️ \\e[91mError:\\e[0mFail generating release $VERSION";
      exit 1;
  fi
  echo "🚀  Release $VERSION successfully generated !"
  sleep 1s

  echo "---------------------------------------------------------------------------"
  echo "📁 Removing local build"
  cd ..
  rm -rf ~/.rit/repos/local
  if [ $? != 0 ]; then
      echo -e "✘️ \\e[91mError:\\e[0mFail removing local build";
      exit 1;
  fi
  echo "🗑 Local build removed successfully !"
  sleep 1s

  echo "---------------------------------------------------------------------------"
  echo "🐙 Adding https://github.com/$USERNAME/$slug_name Github repository to Ritchie"
  sleep 10s
  echo '{"provider":"Github", "name":"'$slug_name'", "version":"'$VERSION'", "url":"'https://github.com/$USERNAME/$slug_name'", "token":"'$TOKEN'", "priority":2}' | rit add repo --stdin
  if [ $? != 0 ]; then
      echo -e "✘️ \\e[91mError:\\e[0mFail adding Github repository to Ritchie ($ rit add repo)";
      exit 1;
  fi
  echo "🔁 Updating Ritchie repository"
  echo '{"name":"'$slug_name'", "version":"'$VERSION'"}' | rit update repo --stdin
  if [ $? != 0 ]; then
      echo -e "✘️ \\e[91mError:\\e[0mFail updating repository on Github ($ rit update repo)";
      exit 1;
  fi
  echo "👏👏👏 New workspace published successfully !"
}
