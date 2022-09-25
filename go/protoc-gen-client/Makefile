PLUGIN_DIR=./bin/protoc-gen-client
SRC=$(shell find ./proto -name '*.proto')
DIST=./dist

.PHONY: debug
debug: gen
	go test ./dist

.PHONY: gen
gen: build
	protoc --plugin=${PLUGIN_DIR} --client_out=${DIST} --go_out=. --go-grpc_out=require_unimplemented_servers=false:. ${SRC}

.PHONY: build
build:
	go build -o ${PLUGIN_DIR} ./main.go

