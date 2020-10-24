# SH
SOURCE_FILE=src/src.csproj
BIN_FOLDER=bin
BIN_FOLDER_LINUX=linux
BIN_FOLDER_DARWIN=darwin
BIN_UNIX=src.dll
SH=$(BIN_FOLDER)/run.sh

build: linux-build macOS-build sh-unix docker

linux-build:
	mkdir -p $(BIN_FOLDER)
	cp -r src/* $(BIN_FOLDER)
	dotnet build $(SOURCE_FILE) -o $(BIN_FOLDER)/$(BIN_FOLDER_LINUX) --configuration Release

macOS-build:
	mkdir -p $(BIN_FOLDER)/$(BIN_FOLDER_DARWIN)
	dotnet build $(SOURCE_FILE) -o $(BIN_FOLDER)/$(BIN_FOLDER_DARWIN) --configuration Release --runtime osx-x64

sh-unix:
	echo '#!/bin/sh' > $(SH)
	echo 'if [ $$(uname) = "Darwin" ]; then' >> $(SH)
	echo '  dotnet "$$(dirname "$$0")"/$(BIN_FOLDER_DARWIN)/$(BIN_UNIX)' >> $(SH)
	echo 'else' >> $(SH)
	echo '	dotnet "$$(dirname "$$0")"/$(BIN_FOLDER_LINUX)/$(BIN_UNIX)' >> $(SH)
	echo 'fi' >> $(SH)
	chmod +x $(SH)

docker:
	cp Dockerfile set_umask.sh $(BIN_FOLDER)
