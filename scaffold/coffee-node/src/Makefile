# Make Run Node
BINARY_NAME=coffee-node
DIST=../dist
DIST_DIR=$(DIST)/commons/bin
DIST_LINUX_DIR=$(DIST)/linux/bin
DIST_WIN_DIR=$(DIST)/windows/bin

build:
	mkdir -p $(DIST_DIR)
	cp run_template $(BINARY_NAME) && chmod +x $(BINARY_NAME)

	cp -r $(BINARY_NAME) index.js package.json coffee $(DIST_DIR)

	#Clean files
	rm $(BINARY_NAME)