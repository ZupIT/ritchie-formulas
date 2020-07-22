# Make Run Python
BINARY_NAME=ml_template.py
BINARY_NAME_WINDOWS=ml_template.bat
DIST=../dist
DIST_DIR=$(DIST)/commons/bin
build:
	mkdir -p $(DIST_DIR)
	cp main.py $(DIST_DIR) && cp -r ml_template Dockerfile set_umask.sh $(DIST_DIR)
	chmod +x $(DIST_DIR)/main.py
	echo 'python main.py' >> $(DIST_DIR)/$(BINARY_NAME_WINDOWS)