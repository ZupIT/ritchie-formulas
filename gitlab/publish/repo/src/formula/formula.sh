#!/bin/bash
# shellcheck disable=SC2181
# shellcheck disable=SC2086
# shellcheck disable=SC2164
# shellcheck disable=SC1001
  removeSpaces() {
    echo "${1}" | xargs | tr " " -
  }

cleanName() {
  tmp="$1"
  if [[ "$1" = *" "* ]]; then
    echo >&2 "Removing spaces from Project Name...: "
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

  runFormula() {

  slug_name=$(cleanName "$PROJECTNAME")
  checkProjectName $slug_name

  cd $WORKSPACE_PATH

  if [ $? != 0 ]; then
   cd -
   exit 1
  fi

  echo "---------------------------------------------------------------------------"

  if git rev-parse --git-dir > /dev/null 2>&1; then
    echo "🚧 This repository already exists. Preparing new commit..."
    git add . > /dev/null
    git commit -m "$VERSION Commit" > /dev/null
  else
    echo "🚥 Repository creation. Preparing first commit..."
    git init
    if [[ $DOCKER_EXECUTION ]]; then
      git config --local user.name $USERNAME
      git config --local user.email $EMAIL
    fi
    git add . > /dev/null
    git commit -m "Initial Commit" > /dev/null
    curl -H 'PRIVATE-TOKEN: '$TOKEN -X POST 'https://gitlab.com/api/v4/projects?name='$slug_name'&visibility='$PRIVACY > /dev/null

    if [ $? != 0 ]; then
      echo -e "✘️ Fail creating new repository";
      exit 1;
    fi
    git remote add origin https://oauth2:$TOKEN@gitlab.com/$USERNAME/$slug_name.git
  fi

  git push origin master > /dev/null
  projectID=$(curl --header "Private-Token: $TOKEN" -X GET https://gitlab.com/api/v4/projects\?search\=$slug_name\ | cut -c "8-15")
  if [[ $DOCKER_EXECUTION ]]; then
    chown 1000:1000 -R $CURRENT_PWD/$slug_name
  fi

  echo "---------------------------------------------------------------------------"
  echo "✅ Project added on Gitlab"
  echo "✅ The id of project is: $projectID"
  echo "📎 Run: $ git clone https://gitlab.com/maurineimirandazup/$slug_name.git"
  sleep 1s

  echo "---------------------------------------------------------------------------"
  echo "🛠 Generating release $VERSION"
  API_JSON=$(printf '{"name":"%s", "tag_name": "%s", "description": "Release of version %s", "ref":"master"}' $VERSION $VERSION $VERSION)
  curl --header 'Content-Type: application/json' --header "Private-Token: $TOKEN" --data "$API_JSON" --request POST 'https://gitlab.com/api/v4/projects/'$projectID'/releases' > /dev/null
  if [ $? != 0 ]; then
      echo -e "✘️ Fail generating release $VERSION";
      exit 1;
  fi
  echo "🚀 Release $VERSION successfully generated"
  sleep 1s

  echo "---------------------------------------------------------------------------"
  echo "📁 Removing local build"
  cd ..
  rm -rf ~/.rit/repos/local
  if [ $? != 0 ]; then
      echo -e "✘️ Fail removing local build";
      exit 1;
  fi
  echo "🗑  Local build removed successfully"
  sleep 1s

  echo "---------------------------------------------------------------------------"
  echo "🐙 Adding Gitlab repository https://gitlab.com/$USERNAME/$slug_name to Ritchie"
  sleep 10s
  echo '{"provider":"Gitlab", "name":"'$slug_name'", "version":"'$VERSION'", "url":"'https://gitlab.com/$USERNAME/$slug_name'", "token":"'$TOKEN'", "priority":2}' | rit add repo --stdin
  if [ $? != 0 ]; then
      echo -e "✘️ Fail adding Github repository to Ritchie ($ rit add repo)";
      exit 1;
  fi
  echo "🔁 Updating Ritchie repository"
  echo '{"name":"'$slug_name'", "version":"'$VERSION'"}' | rit update repo --stdin
  if [ $? != 0 ]; then
      echo -e "✘️ Fail updating repository on Github ($ rit update repo)";
      exit 1;
  fi
  echo "👏👏👏 New workspace published and imported successfully"
}
