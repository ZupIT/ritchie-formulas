@ECHO OFF
SETLOCAL

SET BINARY_NAME=run.bat
SET BIN_FOLDER=bin
SET ENTRY_POINT=main.ps1

:build
  mkdir %BIN_FOLDER%
  xcopy src %BIN_FOLDER% /e/h/i/c
  echo @ECHO OFF > %BIN_FOLDER%\%BINARY_NAME%
  echo Powershell.exe -executionpolicy remotesigned -File %~dp0src\%ENTRY_POINT% >> %BIN_FOLDER%\%BINARY_NAME%
  ENDLOCAL
  exit /b 0
