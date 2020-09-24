@echo off
SETLOCAL

call:%~1
goto exit

:runFormula
  if "%WORKSPACE_PATH%" == " " (
    mkdir %CURRENT_PWD%\%PROJECT_NAME%
    cd %CURRENT_PWD%\%PROJECT_NAME%
    echo %PROJECT_DESCRIPTION% >> README.md
  ) else (
    cd %WORKSPACE_PATH%
  )

  git init
  git add .
  git commit -m "Initial Commit"
  curl -X POST -u %USERNAME%:%TOKEN% -H "Content-Type: application/json" https://api.bitbucket.org/2.0/repositories/%USERNAME%/%PROJECT_NAME% -d "{\"scm\": \"git\", \"is_private\": %PRIVATE%}" > nul
  git remote add origin https://%USERNAME%:%TOKEN%@bitbucket.org/%USERNAME%/%PROJECT_NAME%.git
  git push origin master

  echo Repository successfully initialized with git and added on Bitbucket!!
  goto exit

:exit
  ENDLOCAL
  exit /b
