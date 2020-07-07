# SH
BINARY_NAME=run.sh
DIST=../bin
FOLDER_SRC=helm-configmap

build:
	mkdir -p $(DIST)
	cp main.sh $(DIST)/$(BINARY_NAME) && cp -r $(FOLDER_SRC) $(DIST)
	chmod +x $(DIST)/$(BINARY_NAME)