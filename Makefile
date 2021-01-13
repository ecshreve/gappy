mod-init:
	go mod init github.com/ecshreve/gappy

mod-tidy:
	go mod tidy
	
build:
	go build -o bin/gappy github.com/ecshreve/gappy/cmd/gappy

install:
	go install -i github.com/ecshreve/gappy/cmd/gappy

run-only:
	bin/gappy

run: build run-only

test:
	go test github.com/ecshreve/gappy/...

testv:
	go test -v github.com/ecshreve/gappy/...

testc:
	go test -race -coverprofile=coverage.txt -covermode=atomic github.com/ecshreve/gappy/...