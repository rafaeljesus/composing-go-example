.PHONY: all deps test build

all: deps test build

deps:
	@go get -u github.com/golang/dep/cmd/dep
	@dep ensure

test:
	@go vet ./{cmd,user}/...
	@go test -v -race -cover ./{cmd,user}/...

build:
	@go build ./{cmd,user}/...
