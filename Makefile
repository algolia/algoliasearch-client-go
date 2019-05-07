all: deps generate lint test

deps:
	dep ensure -vendor-only
	go get ./...

generate:
	go generate ./...

lint:
	golangci-lint run ./...

test: unit-tests integration-tests

unit-tests:
	gotest -v ./algolia/...

integration-tests:
	gotest -v ./cts/...

clean:
	rm -f `grep -R -l --include \*.go -F 'DO NOT EDIT' * | grep -v -F 'algolia/internal/gen/'`
	go clean -cache -testcache ./...
