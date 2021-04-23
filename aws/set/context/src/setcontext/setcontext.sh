#!/bin/bash
run() {
	# check if context exist
	exist=$(grep "$CONTEXT_NAME" < ~/.aws/credentials)
	if [ "[$CONTEXT_NAME]" != "$exist" ]; then
	echo "'$CONTEXT_NAME' context does not exist."
	echo "Run 'cat ~/.aws/credentials' to see the avaible contexts." >&2
	exit
	fi

	# replace de default context
	cat ~/.aws/credentials > /tmp/temp_cred
	echo -e "______________~/.aws/credentials BEFORE change:_________________\
	"
	cut -b -30 ~/.aws/credentials
	sed -i "s/default/backupDefault/" /tmp/temp_cred
	sed -i "s/$CONTEXT_NAME/default/" /tmp/temp_cred
	sed -i "s/backupDefault/$CONTEXT_NAME/" /tmp/temp_cred
	cat /tmp/temp_cred > ~/.aws/credentials
	echo -e "______________~/.aws/credentials AFTER change:__________________\
	"
	cut -b -30 ~/.aws/credentials
}
