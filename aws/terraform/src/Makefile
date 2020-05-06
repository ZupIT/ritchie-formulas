# Go parameters
BINARY_NAME=terraform-cli
GOCMD=go
GOBUILD=$(GOCMD) build
GOTEST=$(GOCMD) test
CMD_PATH=./main.go
DIST=../dist
DIST_MAC_DIR=$(DIST)/darwin/bin
BIN_MAC=$(BINARY_NAME)-darwin
DIST_LINUX_DIR=$(DIST)/linux/bin
BIN_LINUX=$(BINARY_NAME)-linux

build:
	mkdir -p $(DIST_MAC_DIR) $(DIST_LINUX_DIR)
	export MODULE=$(GO111MODULE=on go list -m)
	#LINUX
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -tags release -ldflags '-X $(MODULE)/cmd.Version=$(VERSION) -X $(MODULE)/cmd.BuildDate=$(DATE)' -o '$(DIST_LINUX_DIR)/$(BIN_LINUX)' -v $(CMD_PATH)
	#MAC
	GOOS=darwin GOARCH=amd64 $(GOBUILD) -tags release -ldflags '-X $(MODULE)/cmd.Version=$(VERSION) -X $(MODULE)/cmd.BuildDate=$(DATE)' -o '$(DIST_MAC_DIR)/$(BIN_MAC)' -v $(CMD_PATH)

test:
	$(GOTEST) -short `go list ./... | grep -v vendor/`
