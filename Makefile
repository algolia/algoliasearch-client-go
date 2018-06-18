PROJECT=algoliasearch

install:
	go install ./$(PROJECT)

deps:
	dep ensure -vendor-only

test:
	gotest -v ./$(PROJECT)

.PHONY: install deps test
