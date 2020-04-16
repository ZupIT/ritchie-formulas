# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
BINARY_NAME=spring-starter
CMD_PATH=./main.go
DIST=../bin
DIST_MAC=$(DIST)/$(BINARY_NAME)-darwin
DIST_LINUX=$(DIST)/$(BINARY_NAME)-linux
DIST_WIN=$(DIST)/$(BINARY_NAME)-windows.exe

build:
	mkdir -p $(DIST)
	export MODULE=$(GO111MODULE=on go list -m)
	#LINUX
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -tags release -ldflags '-X $(MODULE)/cmd.Version=$(VERSION) -X $(MODULE)/cmd.BuildDate=$(DATE)' -o ./$(DIST_LINUX) -v $(CMD_PATH)
	#MAC
	GOOS=darwin GOARCH=amd64 $(GOBUILD) -tags release -ldflags '-X $(MODULE)/cmd.Version=$(VERSION) -X $(MODULE)/cmd.BuildDate=$(DATE)' -o ./$(DIST_MAC) -v $(CMD_PATH)
	#WINDOWS 64
	GOOS=windows GOARCH=amd64 $(GOBUILD) -tags release -ldflags '-X $(MODULE)/cmd.Version=$(VERSION) -X $(MODULE)/cmd.BuildDate=$(DATE)' -o ./$(DIST_WIN) -v $(CMD_PATH)
