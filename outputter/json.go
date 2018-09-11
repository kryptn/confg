package outputter

import "encoding/json"

func prepareJson(mapping Mapping) ([]byte, error) {
	data, err := json.MarshalIndent(mapping, "", "  ")
	return data, err
}
