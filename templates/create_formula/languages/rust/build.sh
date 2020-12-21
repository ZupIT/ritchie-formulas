#!/bin/sh

BIN_NAME=formula
BIN_FOLDER=bin
UNIX_SH=run.sh
WINDOWS_BAT=run.bat

# rust-build:
	mkdir -p $BIN_FOLDER
	cp -r src/* $BIN_FOLDER
	cargo build --manifest-path $BIN_FOLDER/Cargo.toml --release

	# Unix
	{
	echo "#!/bin/sh"
	echo "\$(dirname \"\$0\")/target/release/$BIN_NAME"
	} >>  $BIN_FOLDER/$UNIX_SH
	chmod +x $BIN_FOLDER/$UNIX_SH
	# Windows
	{
    echo "@ECHO OFF"
    echo "SET mypath=%%~dp0"
    echo "start /B /WAIT %%mypath:~0,-1%%/target/release/$BIN_NAME"
	} >> $BIN_FOLDER/$WINDOWS_BAT
# docker:
	cp Dockerfile set_umask.sh $BIN_FOLDER
