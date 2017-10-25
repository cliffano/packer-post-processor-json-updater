package main

import (
	"fmt"
	"io/ioutil"

	"github.com/Jeffail/gabs"
	"github.com/mitchellh/packer/packer"
)

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
