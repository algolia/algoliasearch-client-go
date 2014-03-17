SRC=$(wildcard *.go)
PROJECT=algoliasearch

.PHONY: test algoliasearch

algoliasearch: ${SRC}
	go build test/algoliasearch.go

test: ${SRC}
	go test ${SRC}

clean:
	${RM} ${PROJECT}
