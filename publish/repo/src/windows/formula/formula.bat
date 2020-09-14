@echo off
SETLOCAL

call:%~1
goto exit

:runFormula
  echo %PROVIDER% 
  if "%PROVIDER%" == "Github" (
    echo Github provider selected
    echo {"privacy":"%PRIVACY%", "project_name":"%PROJECT_NAME%", "workspace_path":"%WORKSPACE_PATH%", "version":"%VERSION%"} | rit github publish repo --stdin
    goto exit
  ) else (
    if "%PROVIDER%" == "Gitlab" (
      echo Gitlab provider selected
      echo {"privacy":"%PRIVACY%", "project_name":"%PROJECT_NAME%", "workspace_path":"%WORKSPACE_PATH%", "version":"%VERSION%"} | rit gitlab publish repo --stdin
      goto exit
    ) else (
      echo Unexpected Provider informed. Check it please and try again.
      goto exit
    )
  )

:exit
  ENDLOCAL
  exit /b
