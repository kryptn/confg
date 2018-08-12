package env

import (
	"errors"
	"fmt"
	"github.com/kryptn/confg/containers"
	"os"
)

type EnvSource struct {
	keys []*containers.Key


}

func readEnvFromKey(key *containers.Key) {
	if key.Default != nil {
		key.Value = key.Default
	}
	value, resolved := os.LookupEnv(key.Lookup)
	key.Resolved = resolved
	if resolved {
		key.Value = value
	}

	fmt.Printf("before -- name: %s, %+v\n", key.Key, key)
	result := os.Getenv(key.Lookup)
	if result != "" {
		key.Value = result
	}
	fmt.Printf("after  -- name: %s, %+v\n", key.Key, key)
}

func (es *EnvSource) Register(key *containers.Key) {
	es.keys = append(es.keys, key)
}

func (es *EnvSource) Resolve() {

	for _, key := range es.keys {
		fmt.Printf("resolving source: %s, key: %s\n", key.Backend, key.Key)
		readEnvFromKey(key)
	}
}

func (es *EnvSource) Collect() []*containers.Key {
	es.Resolve()
	return es.keys
}

func Get(backend *containers.Backend) (*EnvSource, error) {
	if backend.Source != "env" {
		return nil, errors.New("source.env invalid backend")
	}


	return &EnvSource{[]*containers.Key{}}, nil
}
