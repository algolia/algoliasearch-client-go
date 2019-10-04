all: generate deps lint test

generate:
	go generate ./...

deps:
	go get ./...

lint:
	golangci-lint run ./...

test: unit-tests integration-tests

unit-tests:
	go test -v ./algolia/...

integration-tests:
	go test -v ./cts/...

clean:
	rm -f `grep -R -l --include \*.go -F 'DO NOT EDIT' * | grep -v -F 'algolia/internal/gen/'`
	go clean -cache -testcache ./...
