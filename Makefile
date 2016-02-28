tools:
	go get -u github.com/golang/lint/golint

lint:
	golint

style:
	gofmt -w .

test-integration:
	packer build \
		-var-file=$(HOME)/.aws/credentials.json \
		test-integration/success.json

.PHONY: tools lint style test-integration
