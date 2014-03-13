SRC=$(wildcard src/*.go)
PROJECT=algoliasearch

algoliasearch: ${SRC}
	go build ${SRC}

test: ${SRC}
	go test ${SRC}

clean:
	${RM} ${PROJECT}
