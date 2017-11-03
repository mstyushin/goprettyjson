# Basic go commands (copypasted from somewhere)
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
# Binary names
BINARY_NAME=pretty

default: build install

build:
	$(GOBUILD) -o out/$(BINARY_NAME) -v

clean:
	$(GOCLEAN)
	rm -rf out

install:
	cp out/$(BINARY_NAME) /usr/local/bin/
	rm -rf out
	$(GOCLEAN)

uninstall:
	rm -f /usr/local/bin/$(BINARY_NAME)
