# Make Run Python
BINARY_NAME=template.py
BINARY_NAME_WINDOWS=template.bat
DIST=../dist
DIST_DIR=$(DIST)/commons/bin
build:
	mkdir -p $(DIST_DIR)
	cp main.py $(DIST_DIR) && cp -r template Dockerfile set_umask.sh $(DIST_DIR)
	chmod +x $(DIST_DIR)/main.py
	echo 'python main.py' >> $(DIST_DIR)/$(BINARY_NAME_WINDOWS)