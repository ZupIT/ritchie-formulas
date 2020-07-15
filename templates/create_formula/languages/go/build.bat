:: Go parameters
echo off
SETLOCAL
SET BINARY_NAME=main
SET GOCMD=go
SET GOBUILD=%GOCMD% build
SET CMD_PATH=main.go
SET BIN_FOLDER=..\bin
SET DIST_WIN_DIR=%BIN_FOLDER%\windows
SET BIN_WIN=%BINARY_NAME%.exe
SET BAT_FILE=%BIN_FOLDER%\run.bat

:build
    cd src
    mkdir %DIST_WIN_DIR%
    SET GO111MODULE=on
    for /f %%i in ('go list -m') do set MODULE=%%i
    CALL :windows
    GOTO CP_DOCKER
    GOTO DONE
    cd ..

:windows
    SET CGO_ENABLED=
	SET GOOS=windows
    SET GOARCH=amd64
    %GOBUILD% -tags release -o %DIST_WIN_DIR%\%BIN_WIN% -v %CMD_PATH%
    echo @ECHO OFF > %BAT_FILE%
    echo SET mypath=%%~dp0 >> %BAT_FILE%
    echo start /B /WAIT %%mypath:~0,-1%%/windows/main.exe >> %BAT_FILE%
    GOTO DONE

:CP_DOCKER
    copy ..\Dockerfile %BIN_FOLDER%
    copy ..\set_umask.sh %BIN_FOLDER%
    GOTO DONE
:DONE