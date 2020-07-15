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
:build
    cd src
    mkdir %DIST_WIN_DIR%
    SET GO111MODULE=on
    for /f %%i in ('go list -m') do set MODULE=%%i
    CALL :windows
    GOTO DONE
:windows
    SET CGO_ENABLED=
	SET GOOS=windows
    SET GOARCH=amd64
    %GOBUILD% -tags release -o %DIST_WIN_DIR%\%BIN_WIN% -v %CMD_PATH%
    echo @ECHO OFF > %BIN_FOLDER%\run.bat
    echo cd windows  >> %BIN_FOLDER%\run.bat
    echo start /B /WAIT main.exe  >> %BIN_FOLDER%\run.bat
    GOTO DONE
:DONE