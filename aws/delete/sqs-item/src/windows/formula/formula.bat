@echo off
SETLOCAL

call:%~1
goto exit

:runFormula
  echo Deleting item
  aws sqs delete-message --queue-url %RIT_QUEUE_URL% --receipt-handle %RIT_RECEIPT_HANDLE%

  goto exit

:exit
  ENDLOCAL
  exit /b 0
