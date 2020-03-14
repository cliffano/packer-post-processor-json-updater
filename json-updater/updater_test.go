package json-updater

import (
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/Jeffail/gabs"
	"github.com/hashicorp/packer/packer"
	"github.com/stretchr/testify/assert"
)

func TestEnsureJSONFileExistsShouldReturnErrorWhenDirectoryCannotBeCreated(t *testing.T) {
	file := "/inexisting_path/inexisting_file.json"
	err := EnsureJSONFileExists(file)
	assert.Equal(t, "mkdir /inexisting_path: permission denied", err.Error())
}

func TestEnsureJSONFileExistsShouldCreateJSONFileWhenDirectoryCanBeCreatedAndFileDoesNotExist(t *testing.T) {
	file := "testdata/test/ignored_file.json"
	_ = EnsureJSONFileExists(file)
	content, _ := ioutil.ReadFile(file)
	assert.Equal(t, "{}", string(content))
}

func TestEnsureJSONFileExistsShouldLeaveExistingJSONFileAsIsWhenFileAlreadyExists(t *testing.T) {
	file := "testdata/test/unmodifiable_file.json"
	_ = EnsureJSONFileExists(file)
	content, _ := ioutil.ReadFile(file)
	assert.Equal(t, "{\n    \"variables\": {\n        \"variable\": \"should-not-be-modified\"\n    }\n}\n", string(content))
}

func TestUpdateJSONFileShouldReturnErrorWhenFileDoesNotExist(t *testing.T) {
	file := "testdata/inexisting_file.json"
	paths := []string{"some.path.foo.bar"}
	ui := &packer.BasicUi{
		Reader:      new(bytes.Buffer),
		Writer:      new(bytes.Buffer),
		ErrorWriter: new(bytes.Buffer),
	}
	value := "ami-12345"
	err := UpdateJSONFile(file, paths, value, ui)
	assert.Equal(t, "open testdata/inexisting_file.json: no such file or directory", err.Error())
}

func TestUpdateJSONFileShouldSetValuesOnMultiPaths(t *testing.T) {
	file := "testdata/test/modifiable_file.json"
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
	assert.Equal(t, "ami-12345", value1)
	value2 := json.Path("variables.variable2").Data().(string)
	assert.Equal(t, "ami-12345", value2)
}
