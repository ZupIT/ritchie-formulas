# Make Run Rust
BINARY_NAME_UNIX=run.sh
BINARY_NAME_WINDOWS=run.bat
BIN_FOLDER=bin

build: rust-build docker

rust-build:
	mkdir -p $(BIN_FOLDER)
	cargo install --path .
	cp -r src $(BIN_FOLDER)
	cp -r target $(BIN_FOLDER)
	echo '#!/bin/sh' > $(BIN_FOLDER)/$(BINARY_NAME_UNIX)
	echo './target/release/rust' >>  $(BIN_FOLDER)/$(BINARY_NAME_UNIX)
	echo 'target/release/rust.exe' > $(BIN_FOLDER)/$(BINARY_NAME_WINDOWS)
	chmod +x $(BIN_FOLDER)/$(BINARY_NAME_UNIX)

docker:
	cp Dockerfile set_umask.sh $(BIN_FOLDER)
