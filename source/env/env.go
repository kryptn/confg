package env

import (
	"errors"
	"fmt"
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

type keyPair struct {
	key, value string
}

func parseEnvLine(line string) (keyPair, bool) {
	kp := keyPair{}
	if line == "" {
		return kp, false
	}
	if !strings.Contains(line, "=") {
		return kp, false
	}
	if strings.HasPrefix(line, "#") {
		return kp, false
	}

	split := strings.SplitN(line, "=", 2)
	kp.key = strings.TrimSpace(split[0])
	kp.value = strings.TrimSpace(split[1])
	return kp, true
}

func mappingFromEnvLines(lines []string) map[string]string {
	mapping := map[string]string{}
	for _, line := range lines {
		kp, ok := parseEnvLine(line)
		if ok {
			mapping[kp.key] = kp.value
		}
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

func (es *EnvSource) Lookup(lookup string) (interface{}, error) {
	value, ok := es.env[lookup]
	if !ok {
		errorText := fmt.Sprintf("confg.source.env lookup error: %s", lookup)
		return nil, errors.New(errorText)
	}
	return value, nil
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
