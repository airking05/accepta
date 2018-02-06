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

run:
	$(GO) run *.go

.PHONY: test test-verbose build run
