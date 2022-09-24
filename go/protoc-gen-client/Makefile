PLUGIN_DIR=./bin/protoc-gen-client
SRC=$(shell find ./src -name '*.proto')
DIST=./dist

.PHONY: gen
gen: build
	protoc --plugin=${PLUGIN_DIR} --client_out=${DIST} ${SRC}

.PHONY: build
build:
	go build -o ${PLUGIN_DIR} ./main.go

