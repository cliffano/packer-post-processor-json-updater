test-integration:
	packer build \
		-var-file=$(HOME)/.aws/credentials.json \
		test-integration/success.json

.PHONY: test-integration
