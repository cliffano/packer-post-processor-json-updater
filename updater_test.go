package main

import (
	"bytes"
	"github.com/mitchellh/packer/packer"
	"testing"
)

// This is totally not isolated from the filesystem
func TestUpdateJsonFileWithString(t *testing.T) {
	file := "/tmp/testpathstring.json"
	paths := []string{"testpathstring"}
	ui := &packer.BasicUi{
		Reader:      new(bytes.Buffer),
		Writer:      new(bytes.Buffer),
		ErrorWriter: new(bytes.Buffer),
	}
	value := "simplevalue_string"
	_ = UpdateJsonFile(file, paths, value, ui, true)
}

func TestUpdateJsonFileWithMap(t *testing.T) {
	file := "/tmp/testpathmap.json"
	paths := []string{"testpathmap"}
	ui := &packer.BasicUi{
		Reader:      new(bytes.Buffer),
		Writer:      new(bytes.Buffer),
		ErrorWriter: new(bytes.Buffer),
	}
	value := map[string]string{"us-east-1": "ami-123456"}
	_ = UpdateJsonFile(file, paths, value, ui, true)
}
