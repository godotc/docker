GO=go

BIN_PATH=bin
SRC_PATH=src

all: clean build install

docker:
	go build -o $(BIN_PATH)/docker $(SRC_PATH)/*

install:
	cp bin/docker /usr/bin/docker
	cp bin/docker /usr/local
