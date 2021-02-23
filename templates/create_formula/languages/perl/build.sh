#!/bin/sh

BINARY_NAME_UNIX=run.sh
BINARY_NAME_WINDOWS=run.bat
BIN_FOLDER=bin


# Perl-Build:
	mkdir -p $BIN_FOLDER
	cp -r src/* "$BIN_FOLDER"

	#Unix
	{
	echo "#!/bin/sh"
	echo "perl -I \$(dirname \"\$0\") \$(dirname \"\$0\")/main.pl"
	} >>  $BIN_FOLDER/$BINARY_NAME_UNIX
	chmod +x "$BIN_FOLDER/$BINARY_NAME_UNIX"

	#Windows
	{
  echo "@ECHO OFF"
  echo "SET mypath=%%~dp0"
  echo "start /B /D %%mypath%% /WAIT perl -I %%mypath%% main.pl" 
  } >> $BIN_FOLDER/$BINARY_NAME_WINDOWS

#Docker Files:
	cp Dockerfile set_umask.sh $BIN_FOLDER
