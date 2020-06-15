# SH
BINARY_NAME=helm-configmap.sh
DIST=../dist
DIST_DIR=$(DIST)/commons/bin
build:
	mkdir -p $(DIST_DIR)
	cp main.sh $(DIST_DIR)/$(BINARY_NAME) && cp -r helm-configmap $(DIST_DIR) && cp Dockerfile $(DIST_DIR)
	chmod +x $(DIST_DIR)/$(BINARY_NAME)