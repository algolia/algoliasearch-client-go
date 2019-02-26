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
	rm `grep -R -l --include \*.go -F 'DO NOT EDIT' *`
	go generate ./...
