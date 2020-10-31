# Make Run PHP

BIN_FOLDER=bin
BINARY_NAME_UNIX=run.sh
BINARY_NAME_WINDOWS=run.bat

build: php-build sh-unix bat-windows docker

php-build:
	mkdir -p $(BIN_FOLDER)
	cp -r src/* $(BIN_FOLDER)
	composer install -q -d $(BIN_FOLDER)

sh-unix:
	echo '#!/bin/sh' > $(BIN_FOLDER)/$(BINARY_NAME_UNIX)
	echo 'php -f "$$(dirname "$$0")"//index.php' >>  $(BIN_FOLDER)/$(BINARY_NAME_UNIX)
	chmod +x $(BIN_FOLDER)/$(BINARY_NAME_UNIX)

bat-windows:
	echo 'echo off' > $(BIN_FOLDER)/$(BINARY_NAME_WINDOWS)
	echo 'php -f index.php' >> $(BIN_FOLDER)/$(BINARY_NAME_WINDOWS)

docker:
	cp Dockerfile set_umask.sh $(BIN_FOLDER)
