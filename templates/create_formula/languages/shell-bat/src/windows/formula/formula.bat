@echo off
SETLOCAL

call:%~1
goto exit

:runFormula
  echo Hello World!
  echo You receive %SAMPLE_TEXT% in text.
  echo You receive %SAMPLE_LIST% in list.
  echo You receive %SAMPLE_BOOL% in boolean.
  goto exit

:exit
  ENDLOCAL
  exit /b
