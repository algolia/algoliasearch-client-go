build: generate
	go build ./...

deps:
	dep ensure -vendor-only

test: ut it

ut:
	GOCACHE=off gotest -v ./algoliasearch/...

it:
	GOCACHE=off gotest -v ./it/...

generate:
	go generate ./...

clean-generate:
	rm -f `grep -R -l --include \*.go -F 'DO NOT EDIT' * | grep -v -F 'algolia/internal/gen/'`
