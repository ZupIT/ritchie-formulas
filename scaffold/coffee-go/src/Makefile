# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOTEST=$(GOCMD) test
BINARY_NAME=coffee
CMD_PATH=./main.go
DIST=../dist
DIST_MAC=$(DIST)/darwin/bin/$(BINARY_NAME)-darwin
DIST_LINUX=$(DIST)/linux/bin/$(BINARY_NAME)-linux
DIST_WIN=$(DIST)/windows/bin/$(BINARY_NAME)-windows.exe

build:
	mkdir -p $(DIST)
	export MODULE=$(GO111MODULE=on go list -m)
	#LINUX
	GOOS=linux GOARCH=amd64 $(GOBUILD) -tags release -ldflags '-X $(MODULE)/cmd.Version=$(VERSION) -X $(MODULE)/cmd.BuildDate=$(DATE)' -o ./$(DIST_LINUX) -v $(CMD_PATH)
	#MAC
	GOOS=darwin GOARCH=amd64 $(GOBUILD) -tags release -ldflags '-X $(MODULE)/cmd.Version=$(VERSION) -X $(MODULE)/cmd.BuildDate=$(DATE)' -o ./$(DIST_MAC) -v $(CMD_PATH)
	#WINDOWS 64
	GOOS=windows GOARCH=amd64 $(GOBUILD) -tags release -ldflags '-X $(MODULE)/cmd.Version=$(VERSION) -X $(MODULE)/cmd.BuildDate=$(DATE)' -o ./$(DIST_WIN) -v $(CMD_PATH)

test:
	$(GOTEST) -short `go list ./... | grep -v vendor/`
