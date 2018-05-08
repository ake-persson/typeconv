all:	build

format:
	gofmt -s -w .

test:
	golint -set_exit_status ./...
	go vet ./...
	errcheck ./...
	gocyclo -top 25 .
	go test ./... -v -covermode=atomic

build: format test
	go build

.PHONY: format test build
