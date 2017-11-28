package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/hashicorp/packer/packer"
  "github.com/Jeffail/gabs"
)

// Ensure that a file exists. If there are missing directories/sub-directories,
// they will be created. If the file does not exist, an empty JSON file will be created.
func EnsureJSONFileExists(file string) error {

  dir := filepath.Dir(file)
	mkdir_err := os.MkdirAll(dir, os.ModeDir)
	if mkdir_err != nil {
		return mkdir_err
	}

	if _, stat_err := os.Stat(file); os.IsNotExist(stat_err) {
		content := []byte("{}")
		write_err := ioutil.WriteFile(file, content, 0644)
		if write_err != nil {
			return write_err
		}
	}

	return nil
}

// UpdateJSONFile sets the value of the JSON paths within the specified file
// with a new value string, e.g. the AMI ID.
// The JSON file will also be pretty-printed with 4-spaces indentation.
func UpdateJSONFile(file string, paths []string, newValue string, ui packer.Ui) error {

	content, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}

	json, err := gabs.ParseJSON(content)
	if err != nil {
		return err
	}

	for _, path := range paths {
		oldValue, _ := json.Path(path).Data().(string)
		ui.Say(fmt.Sprintf("Updating %s from '%s' to '%s' in %s...", path, oldValue, newValue, file))
		json.SetP(newValue, path)
	}

	err = ioutil.WriteFile(file, []byte(json.StringIndent("", "    ")), 0644)
	return err
}
