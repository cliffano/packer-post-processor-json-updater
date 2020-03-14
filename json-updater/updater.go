package json-updater

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/Jeffail/gabs"
	"github.com/hashicorp/packer/packer"
)

// EnsureJSONFileExists ensures that a file exists. If there are missing directories/sub-directories,
// they will be created. If the file does not exist, an empty JSON file will be created.
func EnsureJSONFileExists(file string) error {

	dir := filepath.Dir(file)
	mkdirErr := os.MkdirAll(dir, os.ModeDir)
	if mkdirErr != nil {
		return mkdirErr
	}

	if _, statErr := os.Stat(file); os.IsNotExist(statErr) {
		content := []byte("{}")
		writeErr := ioutil.WriteFile(file, content, 0644)
		if writeErr != nil {
			return writeErr
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
