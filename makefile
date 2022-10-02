GO=go

BIN_PATH=bin
SRC_PATH=src

all: clean build install

build:
	$(GO) build -o $(BIN_PATH)/mdocker $(SRC_PATH)/*

install:
	cp bin/mdocker /usr/bin/mdocker
	cp bin/mdocker /usr/local/bin/mdocker

unistall:
	rm -rf /usr/bin/mdocker  /usr/local/bin/mdocker
	rm -rf bin/mdocker
	
clean: unistall

