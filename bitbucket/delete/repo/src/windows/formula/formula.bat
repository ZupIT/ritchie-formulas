@echo off
SETLOCAL

call:%~1
goto exit

:runFormula
  curl -X DELETE -u %USERNAME%:%TOKEN% https://api.bitbucket.org/2.0/repositories/%USERNAME%/%PROJECT_NAME%

  echo Repository successfully deleted from Bitbucket
  goto exit

:exit
  ENDLOCAL
  exit /b
