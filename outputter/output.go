package outputter

import (
	"github.com/pkg/errors"
	"io/ioutil"
	"strings"
)

type Mapping map[string]map[string]interface{}

func Output(filename string, mapping Mapping) error {
	var err error
	if strings.HasSuffix(filename, ".toml") {
		err = outputToml(filename, mapping)
		return err
	}
	if strings.HasSuffix(filename, ".yaml") {
		err = outputYaml(filename, mapping)
		return err
	}
	if strings.HasSuffix(filename, ".json") {
		err = outputJson(filename, mapping)
		return err
	}
	return errors.New("output type not supported")
}

func writeToFile(filename string, contents []byte) error {
	err := ioutil.WriteFile(filename, contents, 0644)
	if err != nil {
		return err
	}
	return nil
}
