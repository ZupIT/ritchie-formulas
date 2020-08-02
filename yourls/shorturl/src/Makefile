# Java Parameters
BINARY_NAME_UNIX=shorturl.sh
BINARY_NAME_WINDOWS=shorturl.bat
DIST=../dist
DIST_DIR=$(DIST)/commons/bin
LIB_DIR=../lib/*

build:
	mkdir -p $(DIST_DIR)
	javac -source 1.8 -target 1.8 -cp "$(LIB_DIR):." *.java
	echo "Main-Class: Main" > manifest.txt
	echo "Class-Path: json-simple-1.1.1.jar" >> manifest.txt
	jar cvfm Main.jar manifest.txt *.class shorturl/*.class
	cp run_template $(BINARY_NAME_UNIX) && chmod +x $(BINARY_NAME_UNIX)
	sed '1d' run_template > $(BINARY_NAME_WINDOWS) && chmod +x $(BINARY_NAME_WINDOWS)
	cp Main.jar $(BINARY_NAME_WINDOWS) $(BINARY_NAME_UNIX) Dockerfile set_umask.sh $(LIB_DIR) $(DIST_DIR)
	#Clean files
	rm Main.jar manifest.txt *.class shorturl/*.class $(BINARY_NAME_WINDOWS) $(BINARY_NAME_UNIX)