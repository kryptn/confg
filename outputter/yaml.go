package outputter

import "gopkg.in/yaml.v2"

func prepareYaml(mapping Mapping) ([]byte, error) {
	data, err := yaml.Marshal(mapping)
	return data, err
}
