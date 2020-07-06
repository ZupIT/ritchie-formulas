# SH
BINARY_NAME=run.sh
BINARY_NAME_WINDOWS=run.bat
DIST=../bin

build:
	mkdir -p $(DIST)
	echo '#!/bin/sh' > $(DIST)/$(BINARY_NAME)
	echo 'python main.py' >>  $(DIST)/$(BINARY_NAME)
	cp main.py $(DIST)/ && cp -r coffee $(DIST)
	chmod +x $(DIST)/$(BINARY_NAME)
	echo 'python main.py' >> $(DIST)/$(BINARY_NAME_WINDOWS)