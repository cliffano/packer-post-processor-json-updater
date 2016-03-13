package main

import (
	"bytes"
	"testing"

	"github.com/mitchellh/packer/packer"
	"github.com/stretchr/testify/assert"
)

func TestUpdateJSONFileShouldGiveErrorWhenFileDoesNotExist(t *testing.T) {
	file := "testdata/some_inexisting_file.json"
	paths := []string{"some.path.foo.bar"}
	ui := &packer.BasicUi{
		Reader:      new(bytes.Buffer),
		Writer:      new(bytes.Buffer),
		ErrorWriter: new(bytes.Buffer),
	}
	value := "ami-12345"
	err := UpdateJSONFile(file, paths, value, ui)
	assert.Equal(t, err.Error(), "open testdata/some_inexisting_file.json: no such file or directory")
}
