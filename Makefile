PROJECT=algoliasearch
COVERAGE_FILE=coverage.out

install:
	go install ./$(PROJECT)

deps:
	glide install

test: test-unit

test-unit:
	go test -v ./$(PROJECT)

coverage:
	go list -f '{{if gt (len .TestGoFiles) 0}}"go test -covermode count -coverprofile {{.Name}}.coverprofile -coverpkg ./... {{.ImportPath}}"{{end}}' ./... | xargs -I {} bash -c {}
	gocovmerge `ls *.coverprofile` > $(COVERAGE_FILE)
	go tool cover -html=$(COVERAGE_FILE)

lol:
	curl -X POST -H 'Content-Type: application/json' -d '{"TRAVIS_JOB_ID":$(shell echo ${TRAVIS_JOB_ID})}' https://api-key-dealer.herokuapp.com/1/travis/keys/new

.PHONY: install test clean
