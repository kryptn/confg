package outputter

import (
	"bytes"
	"github.com/BurntSushi/toml"
)

func prepareToml(mapping Mapping) ([]byte, error) {
	buf := new(bytes.Buffer)
	err := toml.NewEncoder(buf).Encode(mapping)
	return buf.Bytes(), err
}
