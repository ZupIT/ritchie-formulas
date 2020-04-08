# SH
DIST=../bin

build:
	mkdir -p $(DIST)
	javac *.java
	echo "Main-Class: Main" > manifest.txt
	jar cvfm Main.jar manifest.txt *.class coffee/*.class
	cat run.sh Main.jar > coffee.run
	chmod +x coffee.run
	mv coffee.run $(DIST)
	rm Main.jar manifest.txt *.class coffee/*.class