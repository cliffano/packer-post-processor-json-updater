package main

import (
	"fmt"
	"io/ioutil"

	"github.com/jeffail/gabs"
	"github.com/mitchellh/packer/packer"
)

func UpdateJsonFile(file string, paths []string, value string, ui packer.Ui) error {

	content, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}

	json, err := gabs.ParseJSON(content)
	if err != nil {
		return err
	}

	for _, path := range paths {
		var value string
		value, _ = json.Path(path).Data().(string)
		ui.Say(fmt.Sprintf("Updating %s to %s in %s...", path, value, file))
		json.SetP(value, path)
	}

	err = ioutil.WriteFile(file, []byte(json.StringIndent("", "    ")), 0644)
	return err
}
