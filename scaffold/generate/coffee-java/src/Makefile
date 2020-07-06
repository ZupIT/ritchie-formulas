# SH
BINARY_NAME_UNIX=run.sh
BINARY_NAME_WINDOWS=run.bat
DIST=../bin

build:
	mkdir -p $(DIST)
	javac -source 1.8 -target 1.8 *.java
	echo "Main-Class: Main" > manifest.txt
	jar cvfm Main.jar manifest.txt *.class coffee/*.class
	echo '#!/bin/sh' > $(DIST)/$(BINARY_NAME_UNIX)
	echo 'java -jar Main.jar' >>  $(DIST)/$(BINARY_NAME_UNIX)
	chmod +x $(DIST)/$(BINARY_NAME_UNIX)
	echo 'java -jar Main.jar' > $(DIST)/$(BINARY_NAME_WINDOWS)
	cp Main.jar $(DIST)
	#Clean files
	rm Main.jar manifest.txt *.class coffee/*.class