package cmd

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

// readFileJSON is a helper function to read a JSON file specified by path and
// decode the data into v.
func readFileJSON(path string, v interface{}) error {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	return json.Unmarshal(data, v)
}

// writeFileJSON is a helper function to JSON encode v and write the result to
// a file specified by path.
func writeFileJSON(path string, v interface{}, perm os.FileMode) error {
	data, err := json.Marshal(v)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(path, data, perm)
	if err != nil {
		return err
	}

	return nil
}
