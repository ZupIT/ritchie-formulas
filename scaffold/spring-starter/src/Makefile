# Go parameters
BINARY_NAME=spring-starter
GOCMD=go
GOBUILD=$(GOCMD) build
GOTEST=$(GOCMD) test
CMD_PATH=./main.go
DIST=../dist
DIST_MAC_DIR=$(DIST)/darwin/bin
BIN_MAC=$(BINARY_NAME)-darwin
DIST_LINUX_DIR=$(DIST)/linux/bin
BIN_LINUX=$(BINARY_NAME)-linux
DIST_WIN_DIR=$(DIST)/windows/bin
BIN_WIN=$(BINARY_NAME)-windows.exe

build:
	mkdir -p $(DIST_MAC_DIR) $(DIST_LINUX_DIR) $(DIST_WIN_DIR)
	export MODULE=$(GO111MODULE=on go list -m)
	#LINUX
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -tags release -ldflags '-X $(MODULE)/cmd.Version=$(VERSION) -X $(MODULE)/cmd.BuildDate=$(DATE)' -o '$(DIST_LINUX_DIR)/$(BIN_LINUX)' -v $(CMD_PATH)
	#MAC
	GOOS=darwin GOARCH=amd64 $(GOBUILD) -tags release -ldflags '-X $(MODULE)/cmd.Version=$(VERSION) -X $(MODULE)/cmd.BuildDate=$(DATE)' -o '$(DIST_MAC_DIR)/$(BIN_MAC)' -v $(CMD_PATH)
	#WINDOWS 64
	GOOS=windows GOARCH=amd64 $(GOBUILD) -tags release -ldflags '-X $(MODULE)/cmd.Version=$(VERSION) -X $(MODULE)/cmd.BuildDate=$(DATE)' -o '$(DIST_WIN_DIR)/$(BIN_WIN)' -v $(CMD_PATH)

test:
	$(GOTEST) -short `go list ./... | grep -v vendor/`