GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
BINARY_NAME=passphrase-web
DEPCMD=dep
RICECMD=rice

.PHONY: all dep rice build clean

all: test dep rice build

test:
	$(GOTEST) -v ./...

dep:
	$(DEPCMD) ensure

rice:
	$(RICECMD) embed-go

build:
	$(GOBUILD) -o $(BINARY_NAME) -v

clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME) rice-box.go
