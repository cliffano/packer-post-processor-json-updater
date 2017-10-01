package main

import (
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/hashicorp/packer/packer"
	"github.com/jeffail/gabs"
	"github.com/stretchr/testify/assert"
)

func TestUpdateJSONFileShouldReturnErrorWhenFileDoesNotExist(t *testing.T) {
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

func TestUpdateJSONFileShouldSetValuesOnMultiPaths(t *testing.T) {
	file := "testdata/test/update_json_file.json"
	paths := []string{"variables.variable1", "variables.variable2"}
	ui := &packer.BasicUi{
		Reader:      new(bytes.Buffer),
		Writer:      new(bytes.Buffer),
		ErrorWriter: new(bytes.Buffer),
	}
	value := "ami-12345"
	_ = UpdateJSONFile(file, paths, value, ui)
	content, _ := ioutil.ReadFile(file)
	json, _ := gabs.ParseJSON(content)
	value1 := json.Path("variables.variable1").Data().(string)
	assert.Equal(t, value1, "ami-12345")
	value2 := json.Path("variables.variable2").Data().(string)
	assert.Equal(t, value2, "ami-12345")
}
