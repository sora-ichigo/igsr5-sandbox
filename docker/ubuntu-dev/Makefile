start: build
	docker container run -it --mount type=bind,source=$(shell pwd)/work,target=/work ubuntu-dev:latest /bin/bash

build:
	docker build . -t ubuntu-dev
