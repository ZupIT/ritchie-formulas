# SH
BINARY_NAME=coffee-java
DIST=../dist
DIST_MAC_DIR=$(DIST)/darwin/bin
DIST_LINUX_DIR=$(DIST)/linux/bin
DIST_WIN_DIR=$(DIST)/windows/bin

build:
	mkdir -p $(DIST_MAC_DIR) $(DIST_LINUX_DIR) $(DIST_WIN_DIR)
	javac *.java
	echo "Main-Class: Main" > manifest.txt
	jar cvfm Main.jar manifest.txt *.class coffee/*.class
	cp run_template $(BINARY_NAME) && chmod +x $(BINARY_NAME)

	#Mac
	cp Main.jar $(BINARY_NAME) $(DIST_MAC_DIR)

	#Linux
	cp Main.jar $(BINARY_NAME) $(DIST_LINUX_DIR)

	#Windows
	cp Main.jar $(BINARY_NAME) $(DIST_WIN_DIR)

	#Clean files
	rm Main.jar manifest.txt *.class coffee/*.class $(BINARY_NAME)