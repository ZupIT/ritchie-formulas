# SH
DIST=../bin

build:
	mkdir -p $(DIST)
	pkg -t node12-linux index.js -o coffee-linux
	pkg -t node12-macos index.js -o coffee-darwin
	pkg -t node12-windows index.js -o coffee-windows.exe
	mv coffee-* $(DIST)