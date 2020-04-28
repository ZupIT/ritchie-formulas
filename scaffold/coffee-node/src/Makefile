# SH
BINARY_NAME=coffee-node
DIST=../dist
DIST_MAC_DIR=$(DIST)/darwin/bin
DIST_LINUX_DIR=$(DIST)/linux/bin
DIST_WIN_DIR=$(DIST)/windows/bin

build:
	mkdir -p $(DIST_MAC_DIR) $(DIST_LINUX_DIR) $(DIST_WIN_DIR)
	pkg -t node12-linux index.js -o coffee-linux
	pkg -t node12-macos index.js -o coffee-darwin
	pkg -t node12-windows index.js -o coffee-windows.exe
	mv coffee-* $(DIST)