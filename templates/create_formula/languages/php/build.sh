#!/bin/sh

BINARY_NAME_UNIX=run.sh
BINARY_NAME_WINDOWS=run.bat
BIN_FOLDER=bin

# php-build:
	mkdir -p $BIN_FOLDER
	cp -r src/* $BIN_FOLDER
	composer install -q -d $BIN_FOLDER

	# Unix
	{
	echo "#!/bin/sh"
	echo "php -f \$(dirname \"\$0\")/index.php"
	} >>  $BIN_FOLDER/$BINARY_NAME_UNIX
	chmod +x $BIN_FOLDER/$BINARY_NAME_UNIX

	# Windows
	{
	echo "@echo off"
	echo "php -f index.php"
	} >> $BIN_FOLDER/$BINARY_NAME_WINDOWS

# docker:
	cp Dockerfile set_umask.sh $BIN_FOLDER
