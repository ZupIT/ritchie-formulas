@echo off
SETLOCAL

call:%~1
goto exit

:runFormula
  cd %WORKSPACE_PATH% || goto exit

  echo ---------------------------------------------------------------------------

  if exist .git (
    echo This repository already exists. Preparing new commit...
    git add .
    git commit -m "%VERSION% Commit"
  ) else (
    echo Repository creation. Preparing first commit...

    git init
    git add . > nul
    git commit -m "Initial Commit" > nul
    curl -H "Authorization: token %TOKEN%" https://api.github.com/user/repos -d "{\"name\":\"%PROJECT_NAME%\", \"private\": %PRIVACY%}" > nul
    if %errorlevel% == 1 (
      echo Fail creating new repository
      goto exit
    )
    git remote add origin https://%USERNAME%:%TOKEN%@github.com/%USERNAME%/%PROJECT_NAME%.git > nul
  )

  git push origin master > nul

  echo ---------------------------------------------------------------------------
  echo Project added on Github
  echo Run: $ git clone https://github.com/%USERNAME%/%PROJECT_NAME%.git

  timeout 2 > nul

  echo ---------------------------------------------------------------------------
  echo Generating release %VERSION%

  SET API_JSON="{\"tag_name\": \"%VERSION%\",\"target_commitish\": \"master\",\"name\": \"%VERSION%\",\"body\": \"Release of version %VERSION%\",\"draft\": false,\"prerelease\": false}"
  curl --data %API_JSON% https://api.github.com/repos/%USERNAME%/%PROJECT_NAME%/releases?access_token=%TOKEN% > nul
  if %errorlevel% == 1 (
    echo Fail generating release %VERSION%
    goto exit
  )

  echo Release %VERSION% successfully generated
  timeout 2 > nul

  echo ---------------------------------------------------------------------------
  echo Removing local build

  rmdir /q/s %USERPROFILE%\.rit\repos\local
  if %errorlevel% == 1 (
    echo Fail removing local build
    goto exit
  )

  echo Local build removed successfully
  timeout 2 > nul

  echo ---------------------------------------------------------------------------
  echo Adding Github repository https://github.com/%USERNAME%/%PROJECT_NAME% to Ritchie
  timeout 6 > nul

  echo {"provider":"Github", "name":"%PROJECT_NAME%", "version":"%VERSION%", "url":"https://github.com/%USERNAME%/%PROJECT_NAME%", "token":"%TOKEN%", "priority":2} | rit add repo --stdin
  if %errorlevel% == 1 (
    echo Fail adding Github repository to Ritchie $ rit add repo
    goto exit
  )

  echo Updating Ritchie repository
  echo {"name":"%PROJECT_NAME%", "version":"%VERSION%"} | rit update repo --stdin
  if %errorlevel% == 1 (
    echo Fail updating repository on Github $ rit update repo
    goto exit
  )

  echo New workspace published and imported successfully

  goto exit

:exit
  ENDLOCAL
  exit /b
