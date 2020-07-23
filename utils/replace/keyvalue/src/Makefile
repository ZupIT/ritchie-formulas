# SH
BINARY_NAME=keyvalue.sh
DIST=../dist
DIST_DIR=$(DIST)/commons/bin
build:
	mkdir -p $(DIST_DIR)
	cp main.sh $(DIST_DIR)/$(BINARY_NAME) && cp -r keyvalue Dockerfile set_umask.sh $(DIST_DIR)
	chmod +x $(DIST_DIR)/$(BINARY_NAME)