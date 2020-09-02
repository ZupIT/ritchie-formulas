:: Ruby parameters
echo off
SETLOCAL
SET BIN_FOLDER=bin
SET BAT_FILE=%BIN_FOLDER%\run.bat
SET SH_FILE=%BIN_FOLDER%\run.sh
:build
    mkdir %BIN_FOLDER%
    xcopy /E /I src %BIN_FOLDER%
    CALL :BAT_WINDOWS
    CALL :SH_LINUX
    CALL :CP_DOCKER
    GOTO DONE

:BAT_WINDOWS
    echo @ECHO OFF > %BAT_FILE%
    echo SET mypath=%%~dp0 >> %BAT_FILE%
    echo call bundle install --gemfile %%mypath:~0,-1%%Gemfile >> %BAT_FILE%
    echo ruby %%mypath:~0,-1%%/index.rb >> %BAT_FILE%
    GOTO DONE

:SH_LINUX
    echo #!/bin/sh > %SH_FILE%
	echo bundle install --quiet --gemfile "$(dirname "$0")"/Gemfile >> %SH_FILE%
	echo ruby "$(dirname "$0")"/index.rb >> %SH_FILE%
    GOTO DONE

:CP_DOCKER
    copy Dockerfile %BIN_FOLDER%
    copy set_umask.sh %BIN_FOLDER%
    GOTO DONE

:DONE