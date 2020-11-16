:: Python parameters
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
    pip3 install -r %BIN_FOLDER%/requirements.txt
    GOTO DONE

:BAT_WINDOWS
    echo @ECHO OFF > %BAT_FILE%
    echo SET mypath=%%~dp0 >> %BAT_FILE%
    echo start /B /WAIT python %%mypath:~0,-1%%/main.py >> %BAT_FILE%

:SH_LINUX
	echo #!/bin/bash > %SH_FILE%
    echo if [ -f /.dockerenv ] ; then >> %SH_FILE%
    echo pip3 install -r "$(dirname "$0")"/requirements.txt ^>^> /dev/null >> %SH_FILE%
    echo fi >> %SH_FILE%
    echo python3 "$(dirname "$0")"/main.py >> %SH_FILE%
    GOTO DONE

:CP_DOCKER
    copy Dockerfile %BIN_FOLDER%
    copy set_umask.sh %BIN_FOLDER%
    GOTO DONE

:DONE
