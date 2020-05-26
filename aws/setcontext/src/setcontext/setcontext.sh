#!/bin/bash
run() {
	cat ~/.aws/credentials | head -n -3 > /tmp/temp_cred
	echo '[default]' >> /tmp/temp_cred
	cat ~/.aws/credentials | grep $CONTEXT_NAME -A 2 | tail -n 2 >> /tmp/temp_cred
	cat /tmp/temp_cred > ~/.aws/credentials
}

run
