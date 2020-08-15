#!/bin/sh
run() {
  #removing previous configuration if exists
  rm -rf $HOME/.yourls
  #(re)creating local directory to store configuration information
  mkdir $HOME/.yourls
  echo "yourls_api_endpoint = $YOURLS_API_ENDPOINT" > $HOME/.yourls/yourls.properties
  echo "yourls_api_secret = $YOURLS_API_SECRET" >> $HOME/.yourls/yourls.properties
}
