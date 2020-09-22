@echo off
SETLOCAL
SET BINARY_NAME=run.bat
SET BIN_FOLDER=bin
SET ENTRY_POINT=main.bat

:build
  mkdir %BIN_FOLDER%
  xcopy src %BIN_FOLDER% /e/h/i/c
  cd %BIN_FOLDER%
  rename %ENTRY_POINT% %BINARY_NAME%
  exit /b 0

ENDLOCAL
