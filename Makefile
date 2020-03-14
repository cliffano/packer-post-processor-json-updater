ci: clean deps lint style test build

clean:
	rm -f packer-post-processor-json-updater*

tools:
	go get -u github.com/tools/godep
	go get -u github.com/golang/lint/golint
	go get -u github.com/mitchellh/gox

deps:
	go get -u github.com/Jeffail/gabs
	go get -u github.com/mitchellh/packer/common
	go get -u github.com/mitchellh/packer/packer
	go get -u github.com/mitchellh/packer/packer/plugin
	go get -u github.com/stretchr/testify/assert
	#godep save -r

build:
	gox -osarch="!plan9/386"

lint:
	golint

style:
	gofmt -w .

test:
	rm -f testdata/test/some_ignored_file.json
	go test

test-integration:
	BUILD_ID=`date '+%Y%m%d%H%M%S'` \
	  packer build \
		testdata/test-integration/success_template.json

.PHONY: ci tools deps build lint style test test-integration
