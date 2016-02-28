package main

import (
	"fmt"
	"io/ioutil"

	"github.com/jeffail/gabs"
	"github.com/mitchellh/packer/packer"
)

func UpdateJsonFile(file string, paths []string, value interface{}, ui packer.Ui, force bool) error {

	// TODO: add a flag to create a file if it does not exist
	content, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}

	// Parse file and create new JSON object if invalid detected and overwrite
	json, err := gabs.ParseJSON(content)
	if err != nil {
		ui.Say("Invalid JSON Template")
		if force {
			json = gabs.New()
		} else {
			return err
		}
	}

	if newValueMap, success := value.(map[string]string); success {
		for _, path := range paths {
			oldValue, _ := json.Path(path).Data().(string)
			ui.Say(fmt.Sprintf("Updating %s from %s to %s in %s...", path, oldValue, newValueMap, file))
			for key, value := range newValueMap {
				ui.Say(fmt.Sprintf("[DEBUG] Setting %s to %s at %s", key, value, path))
				// TODO: this is pretty brittle here since if the type of value of node
				// is not changeable. This might lead to bad render of JSON where previously
				// erroreous values are prevailed over consecutive runs.
				json.Set("", path, key)
				json.Set(value, path, key)
				ui.Say(fmt.Sprintf("[DEBUG] JSON Dump %s", json.String()))

			}
		}
	} else {
		newValue, _ := value.(string)
		for _, path := range paths {
			oldValue, _ := json.Path(path).Data().(string)
			ui.Say(fmt.Sprintf("Updating %s from %s to %s in %s...", path, oldValue, newValue, file))
			ui.Say(fmt.Sprintf("%s", path))
			json.SetP(newValue, path)
		}
	}
	// pretty printed JSON dump
	err = ioutil.WriteFile(file, []byte(json.StringIndent("", "    ")), 0644)
	return err
}
