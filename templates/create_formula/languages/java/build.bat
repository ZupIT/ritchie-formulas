:: Java parameters
echo off
SETLOCAL
SET BIN_FOLDER=bin
SET BIN_NAME=Main.jar
SET BAT_FILE=%BIN_FOLDER%\run.bat
:build
    call mvn clean install
    mkdir %BIN_FOLDER%
    copy target\%BIN_NAME% %BIN_FOLDER%\%BIN_NAME%
    rmdir /Q /S target
    GOTO BAT_WINDOWS
    GOTO CP_DOCKER
    GOTO DONE

:BAT_WINDOWS
    	echo @ECHO OFF > %BAT_FILE%
    	echo java -jar %BIN_NAME% >> %BAT_FILE%

:CP_DOCKER
    copy Dockerfile %BIN_FOLDER%
    copy set_umask.sh %BIN_FOLDER%
    GOTO DONE

:DONE