#!/bin/bash
run() {
	# replace de default context
	cat ~/.aws/credentials > /tmp/temp_cred
	echo -e "\n __________________________________"
	echo -e "\n ~/.aws/credentials BEFORE CHANGE: \n"
	cat ~/.aws/credentials
	sed -i "s/default/backupDefault/" /tmp/temp_cred
	sed -i "s/$CONTEXT_NAME/default/" /tmp/temp_cred
	sed -i "s/backupDefault/$CONTEXT_NAME/" /tmp/temp_cred
	cat /tmp/temp_cred > ~/.aws/credentials
	echo -e "\n __________________________________"
	echo -e "\n ~/.aws/credentials AFTER CHANGE: \n"
	cat ~/.aws/credentials
}
