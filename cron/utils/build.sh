#!/bin/sh

BINARY_NAME=run.sh
BIN_FOLDER=bin

checkCommand () {
  if ! command -v "$1" >/dev/null; then
    echo "$1 required" >&2;
		exit 1;
	fi
}

#python-build:
	checkCommand python3
	checkCommand pip3
	mkdir -p $BIN_FOLDER
	cp -r src/* $BIN_FOLDER
	pip3 install -r $BIN_FOLDER/requirements.txt --user --disable-pip-version-check

#sh_unix:
	{
	echo "#!/bin/bash"
	echo "if [ -f /.dockerenv ] ; then"
	echo "pip3 install -r \$(dirname \"\$0\")/requirements.txt --user --disable-pip-version-check >> /dev/null"
	echo "fi"
	echo "python3 \$(dirname \"\$0\")/main.py"
	} >> $BIN_FOLDER/$BINARY_NAME
	chmod +x $BIN_FOLDER/$BINARY_NAME

#docker:
	cp Dockerfile set_umask.sh $BIN_FOLDER
