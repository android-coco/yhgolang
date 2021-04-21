# Self documented Makefile
# http://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get


.PHONY: all
all: clean build

help:
	@grep -E '^[%.a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) \
		| awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-15s\033[0m %s\n", $$1, $$2}'

clean:
	@go clean

deps:
	$(GOGET) github.com/gorilla/websocket
	$(GOGET) github.com/stripe/stripe-go

doc:
	rm -rf docs/
	swag init . 

build:
	@GOOS=linux GOARCH=amd64 $(GOBUILD)  -o ./ main.go

upload:
	scp bin/linux/node-server yc-test:

uploaddev:
	scp bin/linux/node-server yc-dev:

uploaduat:
	scp bin/linux/node-server yc-demo-2: & \
	wait 
	@echo "upload uat over"

uploadprod:
	scp bin/linux/node-server yc-node: & \
	wait 
	@echo "upload prod over"

dev:
	./bin/linux/node-server --env=dev


dock: 
	docker build . --build-arg SSH_PRIVATE_KEY="$$(cat ~/.ssh/id_rsa)"  --no-cache -t docker-code.aliyun.com:youchuan-blockchain/node-server
