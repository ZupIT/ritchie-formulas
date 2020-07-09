:: Python parameters
echo off
SETLOCAL
SET BINARY_NAME_UNIX=ml_template.sh
SET BINARY_NAME_WINDOWS=ml_template.bat
SET DIST=..\dist
SET DIST_DIR=%DIST%\commons\bin
:build
    mkdir %DIST_DIR%
    echo python main.py >> %DIST_DIR%\%BINARY_NAME_WINDOWS%
    xcopy ml_template %DIST_DIR%\ml_template /E /C /I
    for %%i in (main.py Dockerfile set_umask.sh) do copy %%i %DIST_DIR%
    GOTO DONE
:DONE