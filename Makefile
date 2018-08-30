install:
	go install ./algoliasearch

deps:
	dep ensure -vendor-only

test: ut it

ut:
	GOCACHE=off gotest -v ./algoliasearch/...

it:
	GOCACHE=off gotest -v ./it/...


.PHONY: install deps test ut it
