ci: clean deps lint style test build

clean:
	rm -rf bin/ pkg/ src/ packer-post-processor-json-updater*

deps:
	go mod download

build:
	gox

lint:
	go mod verify
	golint

style:
	gofmt -w *.go

test:
	rm -f testdata/test/some_ignored_file.json
	go test -v ./...

test-integration:
	BUILD_ID=`date '+%Y%m%d%H%M%S'` \
	  packer build \
		testdata/test-integration/success_template.json

.PHONY: ci deps build lint style test test-integration
