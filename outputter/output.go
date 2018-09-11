package outputter

import (
	"github.com/pkg/errors"
	"io/ioutil"
	"strings"
)

type Mapping map[string]map[string]interface{}

var suffixToMapper = map[string]func(Mapping) ([]byte, error){
	".toml": prepareToml,
	".json": prepareJson,
	".yaml": prepareYaml,
	".yml":  prepareYaml,
}

func Output(filename string, mapping Mapping) error {
	var err error

	pathParts := strings.Split(filename, ".")
	suffix := pathParts[len(pathParts)-1]

	outputter, ok := suffixToMapper[suffix]
	if !ok {
		return errors.New("output type not supported")
	}

	data, err := outputter(mapping)
	if err != nil {
		return err
	}

	err = writeToFile(filename, data)
	if err != nil {
		return err
	}

	return nil
}

func writeToFile(filename string, contents []byte) error {
	err := ioutil.WriteFile(filename, contents, 0644)
	if err != nil {
		return err
	}
	return nil
}
