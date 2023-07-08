GIT_VERSION := $(shell git describe --abbrev=0 --tags)

default: build

test:
	go test ./...

build:
	go build -ldflags "-s -w -X github.com/thaim/tflint-ruleset-formatter/project.version=${GIT_VERSION}"

install: build
	mkdir -p ~/.tflint.d/plugins
	mv ./tflint-ruleset-formatter ~/.tflint.d/plugins
