.PHONY: all deps test build

all: deps test build

deps:
	@go get -u github.com/golang/dep/cmd/dep
	@dep ensure

test:
	@go vet ./...
	@go test -v -race -cover ./...

build:
	@go build ./{cmd,user}/...
