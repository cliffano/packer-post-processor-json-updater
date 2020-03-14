ci: clean tools deps lint style test build

clean:
	rm -rf bin/ pkg/ src/ packer-post-processor-json-updater*

tools:
	go get -u golang.org/x/lint/golint
	go get -u github.com/mitchellh/gox

deps:
	go get -u github.com/cliffano/packer-post-processor-json-updater
	go get -u github.com/Jeffail/gabs
	go get -u github.com/hashicorp/hcl/v2/hcldec
	go get -u github.com/hashicorp/packer/common
	go get -u github.com/hashicorp/packer/packer
	go get -u github.com/hashicorp/packer/packer/plugin
	go get -u github.com/stretchr/testify/assert

build:
	bin/gox -osarch="!plan9/386"

lint:
	bin/golint

style:
	gofmt -w *.go

test:
	rm -f testdata/test/some_ignored_file.json
	go test

test-integration:
	BUILD_ID=`date '+%Y%m%d%H%M%S'` \
	  packer build \
		testdata/test-integration/success_template.json

.PHONY: ci tools deps build lint style test test-integration
