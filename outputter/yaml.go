package outputter

import (
	"gopkg.in/yaml.v2"
)

func outputYaml(filename string, mapping Mapping) error {
	b, err := yaml.Marshal(mapping)
	if err != nil {
		return err
	}
	err = writeToFile(filename, b)
	return err
}
