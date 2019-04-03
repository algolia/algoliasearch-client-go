all: deps generate lint test

deps:
	dep ensure -vendor-only

generate:
	go generate ./...

lint:
	golangci-lint run ./algolia/...

test: unit-tests integration-tests

unit-tests:
	GOCACHE=off gotest -v ./algoliasearch/...

integration-tests:
	GOCACHE=off gotest -v ./it/...

clean:
	rm -f `grep -R -l --include \*.go -F 'DO NOT EDIT' * | grep -v -F 'algolia/internal/gen/'`
