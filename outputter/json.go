package outputter

import (
	"encoding/json"
)

func outputJson(filename string, mapping Mapping) error {

	b, err := json.MarshalIndent(mapping, "", "  ")
	if err != nil {
		return err
	}
	err = writeToFile(filename, b)
	return err
}
