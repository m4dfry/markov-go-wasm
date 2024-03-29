# More a script than a makefile

# Go parameters
GOCMD=go
GORUN=$(GOCMD) run
GOBUILD=$(GOCMD) build

wasm:
	GOARCH=wasm GOOS=js $(GOBUILD) -o docs/main.wasm main.go
serv:
	$(GORUN) goserv.go