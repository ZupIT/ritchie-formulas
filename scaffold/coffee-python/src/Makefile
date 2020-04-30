# SH
BINARY_NAME=coffee-python.py
BINARY_NAME_WINDOWS=coffee.bat
DIST=../dist
DIST_DIR=$(DIST)/commons/bin
build:
	mkdir -p $(DIST_DIR)
	cp main.py $(DIST_DIR)/$(BINARY_NAME) && cp -r coffee $(DIST_DIR)
	chmod +x $(DIST_DIR)/$(BINARY_NAME)
	echo 'python coffee-python.py' >> $(DIST_DIR)/$(BINARY_NAME_WINDOWS)