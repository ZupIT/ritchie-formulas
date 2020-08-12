# SH
BINARY_NAME=run.sh
BIN_FOLDER=bin

build: bash-build

bash-build:
	mkdir -p $(BIN_FOLDER)
	cp -r src/* $(BIN_FOLDER)
	mv $(BIN_FOLDER)/main.sh $(BIN_FOLDER)/$(BINARY_NAME)
	chmod +x $(BIN_FOLDER)/$(BINARY_NAME)
