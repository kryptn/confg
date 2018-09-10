package outputter

import (
	"bytes"
	"github.com/BurntSushi/toml"
)

func outputToml(filename string, mapping Mapping) error {
	buf := new(bytes.Buffer)
	if err := toml.NewEncoder(buf).Encode(mapping); err != nil {
		return err
	}
	err := writeToFile(filename, buf.Bytes())
	return err
}
