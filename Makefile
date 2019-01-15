GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
BINARY_NAME=passphrase-web
DEPCMD=dep
RICECMD=rice
NPMCMD=npm

.PHONY: all ui dep rice build clean

all: ui test dep rice build

ui:
	cd web && $(NPMCMD) install
	cd web && $(NPMCMD) run build

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
	rm -rf web/dist
