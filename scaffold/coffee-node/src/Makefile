# Make Run Node
BINARY_NAME_UNIX=coffee-node.sh
BINARY_NAME_WINDOWS=coffee-node.bat
DIST=../dist
DIST_DIR=$(DIST)/commons/bin

build:
	mkdir -p $(DIST_DIR)
	cp run_template $(BINARY_NAME_UNIX) && chmod +x $(BINARY_NAME_UNIX)
	cp run_template $(BINARY_NAME_WINDOWS) && chmod +x $(BINARY_NAME_WINDOWS)

	cp -r $(BINARY_NAME_WINDOWS) $(BINARY_NAME_UNIX) index.js package.json coffee $(DIST_DIR)

	#Clean files
	rm $(BINARY_NAME_UNIX) $(BINARY_NAME_WINDOWS)
