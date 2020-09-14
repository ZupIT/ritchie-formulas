@echo off
SETLOCAL

call:%~1
goto exit

:run
  if %WORKSPACE_PATH% == " " (
    mkdir %WORKSPACE_PATH%
    cd %WORKSPACE_PATH%
    echo %PROJECT_DESCRIPTION% >> README.md
  ) else (
    cd %WORKSPACE_PATH%
  )

  git init
  git add .
  git commit -m "Initial Commit"

  curl -H "Authorization: token %TOKEN%" https://api.github.com/user/repos -d "{\"name\":\"%PROJECT_NAME%\", \"private\": %PRIVATE%}"
  git remote add origin https://%USERNAME%:%TOKEN%@github.com/%USERNAME%/%PROJECT_NAME%.git
  git push origin master

  echo Repository successfully initialized with git and added on Github!!

:exit
  ENDLOCAL
  exit /b
