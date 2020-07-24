:: Java parameters
echo off
SETLOCAL
SET BIN_FOLDER=bin
SET BAT_FILE=%BIN_FOLDER%\run.bat
:build
    rmdir /Q /S %BIN_FOLDER%
    mkdir %BIN_FOLDER%
    xcopy /E /I src %BIN_FOLDER%
    cd %BIN_FOLDER%
    call npm install
    cd ..
    GOTO BAT_WINDOWS
    GOTO CP_DOCKER
    GOTO DONE

:BAT_WINDOWS
    echo @ECHO OFF > %BAT_FILE%
    echo SET mypath=%%~dp0 >> %BAT_FILE%
    echo start /B /WAIT node %%mypath:~0,-1%%/main.js >> %BAT_FILE%

:CP_DOCKER
    copy Dockerfile %BIN_FOLDER%
    copy set_umask.sh %BIN_FOLDER%
    GOTO DONE

:DONE
