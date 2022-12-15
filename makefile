SHELL := /bin/bash
VERSION := 1.0

all: jh

build:
	go build -ldflags "-X main.build=local"

jh:
	docker build \
		-f zarf/dockerfile \
		-t jh-api-amd64:$(VERSION) \
		--build-arg APPID=$(APP_ID) \
		--build-arg BUILD_REF=$(VERSION) \
		--build-arg BUILD_DATE=`date -u +"%Y-%m-%dT%H:%M:%SZ"` \
		.	

container-stop:
	docker stop "$(docker ps -a -q)" && docker rm "$(docker ps -a -q)"

tidy:
	go mod tidy
	go mod vendor

deps-upgrade:
	go get -u -v ./...
	go mod tidy
	go mod vendor	