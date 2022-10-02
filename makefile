GO=go

BIN_PATH=bin
SRC_PATH=src

project=docker

all: clean build install

build:
	$(GO) build -o $(BIN_PATH)/docker $(SRC_PATH)/*

install:
	cp bin/docker /usr/bin/docker
	cp bin/docker /usr/local/bin/docker
	mkdir -p /var/lib/docker/images
	mkdir -p /var/lib/docker/volumes
	mkdir -p /var/lib/docker/containers

unistall:
	rm -rf /usr/bin/docker  /usr/local/bin/docker
	rm -rf bin/docker
	
clean: unistall

run: 
	docker run /bin/bash
