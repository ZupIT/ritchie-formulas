# Make Run PHP
BINARY_NAME_UNIX=run.sh
BINARY_NAME_WINDOWS=run.bat
BIN_FOLDER=bin

build: php-build docker

php-build:
	mkdir -p $(BIN_FOLDER)
	echo '#!/bin/sh' > $(BIN_FOLDER)/$(BINARY_NAME_UNIX)
	echo 'php -f "$$(dirname "$$0")"//index.php' >>  $(BIN_FOLDER)/$(BINARY_NAME_UNIX)
	echo 'php -f index.php' > $(BIN_FOLDER)/$(BINARY_NAME_WINDOWS)
	cp -r src/* $(BIN_FOLDER)
	chmod +x $(BIN_FOLDER)/$(BINARY_NAME_UNIX)

docker:
	cp Dockerfile set_umask.sh $(BIN_FOLDER)
