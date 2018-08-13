package env

import (
	"errors"
	"github.com/kryptn/confg/containers"
	"io/ioutil"
	"os"
	"strings"
)

type EnvSource struct {
	env map[string]string
}

func (es EnvSource) insertMapping(mapping map[string]string) {
	for key, value := range mapping {
		es.env[key] = value
	}
}

func mappingFromEnvLines(lines []string) map[string]string {
	mapping := map[string]string{}
	for _, envLine := range lines {
		isEmpty := envLine == ""
		doesNotAssign := !strings.Contains(envLine, "=")
		comment := strings.HasPrefix(envLine, "#")

		if isEmpty || doesNotAssign || comment {
			continue
		}

		keyValue := strings.SplitN(envLine, "=", 2)
		//log.Printf("-- KEYVALUE -- %+v", keyValue)
		key := strings.TrimSpace(keyValue[0])
		value := strings.TrimSpace(keyValue[1])
		mapping[key] = value
	}
	return mapping
}

func (es EnvSource) projectFromFile(filename string) error {
	rawContent, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	content := string(rawContent)
	lines := strings.Split(content, "\n")
	mapping := mappingFromEnvLines(lines)
	es.insertMapping(mapping)
	return nil
}

func (es EnvSource) projectEnv() error {
	mapping := mappingFromEnvLines(os.Environ())
	es.insertMapping(mapping)
	return nil
}

func (es *EnvSource) Lookup(lookup string) (interface{}, bool) {
	value, ok := es.env[lookup]
	return value, ok
}

func (es *EnvSource) Gather(keys []*containers.Key) {
	for _, key := range keys {
		v, ok := es.Lookup(key.Lookup)
		key.Inject(v, ok)
	}
}

func Get(backend *containers.Backend) (*EnvSource, error) {
	if backend.Source != "env" {
		return nil, errors.New("source.env invalid backend")
	}
	es := &EnvSource{map[string]string{}}

	es.projectEnv()

	if backend.EnvFile != "" {
		es.projectFromFile(backend.EnvFile)
	}

	return es, nil
}
