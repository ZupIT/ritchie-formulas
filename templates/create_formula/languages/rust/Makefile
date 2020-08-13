# Make Run Rust
BINARY_NAME_UNIX=run.sh
BINARY_NAME_UNIX_DOCKER=runDocker.sh
BINARY_NAME_WINDOWS=run.bat
BIN_FOLDER=bin

build: rust-build sh-unix bat-windows docker

rust-build:
	mkdir -p $(BIN_FOLDER)
	cp -r src $(BIN_FOLDER)
	cp Cargo.toml $(BIN_FOLDER)

sh-unix:
	echo '#!/bin/sh' > $(BIN_FOLDER)/$(BINARY_NAME_UNIX)
	echo 'cargo run --quiet' >>  $(BIN_FOLDER)/$(BINARY_NAME_UNIX)
	echo 'rm Cargo.lock' >>  $(BIN_FOLDER)/$(BINARY_NAME_UNIX)
	echo 'rm -rf target' >>  $(BIN_FOLDER)/$(BINARY_NAME_UNIX)
	echo 'cargo run --quiet' > $(BIN_FOLDER)/$(BINARY_NAME_WINDOWS)
	chmod +x $(BIN_FOLDER)/$(BINARY_NAME_UNIX)

bat-windows:
	echo 'cargo run --quiet' > $(BIN_FOLDER)/$(BINARY_NAME_WINDOWS)

docker:
	echo '#!/bin/sh' > $(BIN_FOLDER)/$(BINARY_NAME_UNIX_DOCKER)
	echo 'cargo run --manifest-path /rit/Cargo.toml --quiet' >>  $(BIN_FOLDER)/$(BINARY_NAME_UNIX_DOCKER)
	chmod +x $(BIN_FOLDER)/$(BINARY_NAME_UNIX_DOCKER)
	cp Dockerfile set_umask.sh $(BIN_FOLDER)
