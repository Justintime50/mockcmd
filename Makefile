## help - Display help about make targets for this Makefile
help:
	@cat Makefile | grep '^## ' --color=never | cut -c4- | sed -e "`printf 's/ - /\t- /;'`" | column -s "`printf '\t'`" -t

## build - Build the project
build:
	go build ./mockcmd

## clean - Clean the project
clean:
	rm -rf dist
	rm $(go env GOPATH)/bin/mockcmd

## coverage - Get test coverage and open it in a browser
coverage: 
	go clean -testcache && go test ./... -coverprofile=covprofile && go tool cover -html=covprofile

## install - Install globally from source
install:
	go build -o $(go env GOPATH)/bin/mockcmd

## lint - Lint the project
lint:
	golangci-lint run

## test - Test the project
test:
	go clean -testcache && go test ./...

.PHONY: help build clean coverage install lint test
