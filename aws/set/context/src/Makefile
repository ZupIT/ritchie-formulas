# SH
BINARY_NAME=run.sh
BIN=../bin
build:
	mkdir -p $(BIN)
	cp main.sh $(BIN)/$(BINARY_NAME) && cp -r setcontext $(BIN) && cp Dockerfile $(BIN)
	chmod +x $(BIN)/$(BINARY_NAME)