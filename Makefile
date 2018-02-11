BINNAME = accepta
GO = go

test:
	export CWD=${PWD} && $(GO) test -cover $(PKGS)

test-verbose:
	export CWD=${PWD} && $(GO) test -cover $(PKGS)

test-short:
	export CWD=${PWD} && $(GO) test -cover $(PKGS) -short

build:
	$(GO) build $(PKGS)

swagger:
	go get -u github.com/go-swagger/go-swagger/cmd/swagger
	swagger generate client -f swaggers/api136.yaml -A core

run:
	$(GO) run *.go

.PHONY: test test-verbose build run
