PLUGIN_DIR=./bin/protoc-gen-example
SRC=$(shell find ./src -name '*.proto'
DIST=./dist

.PHONY: gen
gen:
	protoc --plugin=${PLUGIN_DIR} --example_out=. ./example.proto
.PHONY: build
build:
	go build -o ${PLUGIN_DIR} ./main.go

