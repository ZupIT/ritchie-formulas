#!/bin/bash
run() {
	< ~/.aws/credentials tr ' ' _ | nl
	head -n -3 > /tmp/temp_cred
	echo '[default]' >> /tmp/temp_cred
	< ~/.aws/credentials tr ' ' _ | nl
	grep "$CONTEXT_NAME" -A 2 | tail -n 2 >> /tmp/temp_cred
	cat /tmp/temp_cred > ~/.aws/credentials
}

run
