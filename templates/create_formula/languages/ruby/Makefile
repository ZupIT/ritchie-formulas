# Make Run Ruby
BINARY_NAME_UNIX=run.sh
BINARY_NAME_WINDOWS=run.bat
BIN_FOLDER=bin

build: ruby-build docker

ruby-build:
	mkdir -p $(BIN_FOLDER)
	echo '#!/bin/sh' > $(BIN_FOLDER)/$(BINARY_NAME_UNIX)
	echo 'php -f "$$(dirname "$$0")"//index.php' >>  $(BIN_FOLDER)/$(BINARY_NAME_UNIX)
	echo 'php -f index.php' > $(BIN_FOLDER)/$(BINARY_NAME_WINDOWS)
	cp -r src/* $(BIN_FOLDER)
	chmod +x $(BIN_FOLDER)/$(BINARY_NAME_UNIX)

docker:
	cp Dockerfile set_umask.sh $(BIN_FOLDER)

#build:
#	mkdir -p $(DIST_DIR)
#	cp run_template $(BINARY_NAME_UNIX) && chmod +x $(BINARY_NAME_UNIX)
#	echo './index.rb' >> $(DIST_DIR)/$(BINARY_NAME_WINDOWS)
#
#	cp -r . $(DIST_DIR)
#
#	#Clean files
#	rm $(BINARY_NAME_UNIX)