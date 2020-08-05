# SH
BINARY_NAME=run.sh
BINARY_NAME_WINDOWS=run.bat
BIN_FOLDER=bin

build: python-build docker

python-build:
	mkdir -p $(BIN_FOLDER)
	echo '#!/bin/sh' > $(BIN_FOLDER)/$(BINARY_NAME)
	echo 'python3 "$$(dirname "$$0")"/main.py' >>  $(BIN_FOLDER)/$(BINARY_NAME)
	cp -r src/* $(BIN_FOLDER)
	chmod +x $(BIN_FOLDER)/$(BINARY_NAME)
	echo 'python main.py' >> $(BIN_FOLDER)/$(BINARY_NAME_WINDOWS)

docker:
	cp Dockerfile set_umask.sh $(BIN_FOLDER)
