# Make Run PHP
BINARY_NAME_UNIX=run.sh
BINARY_NAME_WINDOWS=run.bat
DIST=bin

build:
	mkdir -p $(DIST)
	echo '#!/bin/sh' > $(DIST)/$(BINARY_NAME_UNIX)
	echo 'php -f index.php' >>  $(DIST)/$(BINARY_NAME_UNIX)
	echo 'php -f index.php' > $(DIST)/$(BINARY_NAME_WINDOWS)
	cp -r src/* $(DIST)
	chmod +x $(DIST)/$(BINARY_NAME_UNIX)
