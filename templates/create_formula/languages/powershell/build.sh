#!/bin/sh

BIN_FOLDER=bin
SH=$BIN_FOLDER/run.sh
BAT=$BIN_FOLDER/run.bat

# bash-build:
	mkdir -p $BIN_FOLDER
	cp -r src/* $BIN_FOLDER
	chmod +x "$BIN_FOLDER/$BINARY_NAME"

# sh-unix:
	{
	echo "#!/bin/sh"
	echo "pwsh \$(dirname \"\$0\")/main.ps1"
	} >> $SH
	chmod +x $SH

# bat-windows:
	{
	echo "@ECHO OFF"
	echo "SET mypath=%~dp0"
	echo "start /B /WAIT %mypath:~0,-1%/main.ps1"
	} >> $BAT

# docker:
	cp Dockerfile set_umask.sh $BIN_FOLDER
