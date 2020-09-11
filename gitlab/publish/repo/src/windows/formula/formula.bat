@echo off
SETLOCAL

call:%~1
goto exit

:runFormula
  cd %WORKSPACE_PATH% || goto exit

  if exist .git (
    echo This repository already exists. Preparing new commit...
    git add .
    git commit -m "%VERSION% Commit"
  ) else (
    echo Repository creation. Preparing first commit...

    git init
    git add . > nul
    git commit -m "Initial Commit" > nul
    curl -H "PRIVATE-TOKEN: %TOKEN%" -X POST "https://gitlab.com/api/v4/projects?name=%PROJECT_NAME%&visibility=%PRIVACY%" > nul
    if %errorlevel% == 1 (
      echo Fail creating new repository
      goto exit
    )
    git remote add origin https://oauth2:%TOKEN%@gitlab.com/%USERNAME%/%PROJECT_NAME%.git > nul
  )

  git push origin master > nul

  echo ---------------------------------------------------------------------------
  echo Project added on Gitlab
  echo Run: $ git clone https://gitlab.com/%USERNAME%/%PROJECT_NAME%.git
  timeout 2 > nul

  echo ---------------------------------------------------------------------------
  echo Generating release %VERSION%

  SET API_JSON="{\"name\":\"%VERSION%\", \"tag_name\":\"%VERSION%\", \"description\":\"Release of version %VERSION%\", \"ref\":\"master\"}"
  SET url_encoded=%USERNAME%%%2f%PROJECT_NAME%
  curl --header "Content-Type: application/json" --header "Private-Token: %TOKEN%" --data %API_JSON% --request POST "https://gitlab.com/api/v4/projects/%url_encoded%/releases" > nul
  if %errorlevel% == 1 (
    echo Fail generating release %VERSION%
    goto exit
  )

  echo Release %VERSION% successfully generated
  timeout 2 > nul

  echo ---------------------------------------------------------------------------
  echo Adding Gitlab repository https://gitlab.com/%USERNAME%/%PROJECT_NAME% to Ritchie
  timeout 6 > nul

  echo {"provider":"Gitlab", "name":"%PROJECT_NAME%", "version":"%VERSION%", "url":"https://gitlab.com/%USERNAME%/%PROJECT_NAME%", "token":"%TOKEN%", "priority":2} | rit add repo --stdin
  if %errorlevel% == 1 (
    echo Fail adding Gitlab repository to Ritchie $ rit add repo
    goto exit
  )

  echo Updating Ritchie repository
  echo {"name":"%PROJECT_NAME%", "version":"%VERSION%"} | rit update repo --stdin
  if %errorlevel% == 1 (
    echo Fail updating repository on Gitlab $ rit update repo
    goto exit
  )

  echo New workspace published and imported successfully

  goto exit

:exit
  ENDLOCAL
  exit /b
