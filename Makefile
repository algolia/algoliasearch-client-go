all: deps generate lint test

deps:
	go get -u ./...

generate:
	go generate ./...

lint:
	golangci-lint run ./algolia/...

test: unit-tests integration-tests

unit-tests:
	gotest -v ./algolia/...

integration-tests:
	gotest -v ./cts/...

clean:
	rm -f `grep -R -l --include \*.go -F 'DO NOT EDIT' * | grep -v -F 'algolia/internal/gen/'`
