:: Ruby parameters
echo off
SETLOCAL
SET BIN_FOLDER=bin
SET BAT_FILE=%BIN_FOLDER%\run.bat
SET SH_FILE=%BIN_FOLDER%\run.sh
:build
    mkdir %BIN_FOLDER%
    xcopy /E /I src %BIN_FOLDER%
    CALL bundle config set path vendor/bundle
    CALL bundle install --gemfile %%BIN_FOLDER%%/Gemfile
    CALL :BAT_WINDOWS
    CALL :SH_LINUX
    CALL :CP_DOCKER
    GOTO DONE

:BAT_WINDOWS
    echo @ECHO OFF > %BAT_FILE%
    echo ruby /index.rb >> %BAT_FILE%
    GOTO DONE

:SH_LINUX
    echo #!/bin/sh > %SH_FILE%
    echo cd "$(dirname "$0")" >> %SH_FILE%
    echo bundle config set path vendor/bundle >> %SH_FILE%
	echo ruby ./index.rb >> %SH_FILE%
    GOTO DONE

:CP_DOCKER
    copy Dockerfile %BIN_FOLDER%
    copy set_umask.sh %BIN_FOLDER%
    GOTO DONE

:DONE
