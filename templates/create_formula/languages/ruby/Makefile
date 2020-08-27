# Make Run Ruby
BINARY_NAME_UNIX=run.sh
BINARY_NAME_WINDOWS=run.bat
BIN_FOLDER=bin

build: ruby-build docker

ruby-build:
	mkdir -p $(BIN_FOLDER)
	cp -r src/* $(BIN_FOLDER)
	echo '#!/bin/sh' > $(BIN_FOLDER)/$(BINARY_NAME_UNIX)
	echo 'bundle install --gemfile "$$(dirname "$$0")"/Gemfile'  >> $(BIN_FOLDER)/$(BINARY_NAME_UNIX)
	echo 'ruby "$$(dirname "$$0")"/index.rb' >> $(BIN_FOLDER)/$(BINARY_NAME_UNIX)

	echo 'bundle install --gemfile "$$(dirname "$$0")"/Gemfile' >> $(BIN_FOLDER)/$(BINARY_NAME_WINDOWS)
	echo 'ruby "$$(dirname "$$0")"/index.rb' >> $(BIN_FOLDER)/$(BINARY_NAME_WINDOWS)
	chmod +x $(BIN_FOLDER)/$(BINARY_NAME_UNIX)

docker:
	cp Dockerfile set_umask.sh $(BIN_FOLDER)
