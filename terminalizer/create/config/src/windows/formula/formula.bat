@echo off
SETLOCAL

call:%~1
goto exit

:runFormula
  echo This formula does not have support for Windows

  goto exit

:exit
  ENDLOCAL
  exit /b 0
