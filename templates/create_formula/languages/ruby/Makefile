# Make Run Ruby
BINARY_NAME_UNIX=run.sh
BINARY_NAME_WINDOWS=run.bat
BIN_FOLDER=bin

build: ruby-build docker

ruby-build:
	mkdir -p $(BIN_FOLDER)
	echo '#!/bin/sh' 		 > $(BIN_FOLDER)/$(BINARY_NAME_UNIX)
	echo 'bundle install --path vendor/bundle' 	>> $(BIN_FOLDER)/$(BINARY_NAME_UNIX)
	echo './index.rb'		>> $(BIN_FOLDER)/$(BINARY_NAME_UNIX)

	echo 'bundle install --path vendor/bundle' 	>  $(BIN_FOLDER)/$(BINARY_NAME_WINDOWS)
	echo './index.rb'		>> $(BIN_FOLDER)/$(BINARY_NAME_WINDOWS)

	cp -r src/* $(BIN_FOLDER)
	chmod +x $(BIN_FOLDER)/$(BINARY_NAME_UNIX)

docker:
	cp Dockerfile set_umask.sh $(BIN_FOLDER)