tools:
	go get -u github.com/tools/godep
	go get -u github.com/golang/lint/golint
	go get -u github.com/mitchellh/gox

deps:
	go get github.com/jeffail/gabs
	go get github.com/mitchellh/packer/common
	#go get github.com/mitchellh/packer/packer
	#go get github.com/mitchellh/packer/packer/plugin
	#godep save -r

build:
	gox -osarch="!plan9/386"

lint:
	golint

style:
	gofmt -w .

test:
	go test

test-integration:
	packer build \
		-var-file=$(HOME)/.aws/credentials.json \
		test-integration/success.json

.PHONY: tools deps build lint style test test-integration
