#!/bin/sh

runFormula() {
  echo "Deleting item"
  aws sqs delete-message --queue-url "$RIT_QUEUE_URL" --receipt-handle "$RIT_RECEIPT_HANDLE"
  sleep 5000
}