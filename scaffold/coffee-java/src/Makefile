# SH
BINARY_NAME_UNIX=coffee-java.sh
BINARY_NAME_WINDOWS=coffee-java.bat
DIST=../dist
DIST_DIR=$(DIST)/commons/bin

build:
	mkdir -p $(DIST_DIR)
	javac -source 1.8 -target 1.8 *.java
	echo "Main-Class: Main" > manifest.txt
	jar cvfm Main.jar manifest.txt *.class coffee/*.class
	cp run_template $(BINARY_NAME_UNIX) && chmod +x $(BINARY_NAME_UNIX)
	cp run_template $(BINARY_NAME_WINDOWS) && chmod +x $(BINARY_NAME_WINDOWS)
	cp Main.jar $(BINARY_NAME_WINDOWS) $(BINARY_NAME_UNIX) $(DIST_DIR)
	#Clean files
	rm Main.jar manifest.txt *.class coffee/*.class $(BINARY_NAME_WINDOWS) $(BINARY_NAME_UNIX)