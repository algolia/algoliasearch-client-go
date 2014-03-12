SRC=$(wildcard src/*.go)
PROJECT=algoliasearch

algoliasearch: ${SRC}
	go build ${SRC}

clean:
	${RM} ${PROJECT}
