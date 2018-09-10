package parser

import (
	"github.com/BurntSushi/toml"
	"github.com/kryptn/confg/containers"
)

type Parser struct {
	confg *containers.Confg

	path string

	defaults map[string]map[string]interface{}

	md   toml.MetaData
	root map[string]toml.Primitive
}

var protectedKeys = map[string]struct{}{
	"version": {},
	"backend": {},
	"default": {},
}

func ConfgFromFile(filename string) (*containers.Confg, error) {
	var err error
	var fatal bool
	parser := Parser{
		confg: &containers.Confg{},
		root:  map[string]toml.Primitive{},
		path:  filename,
	}

	parser.md, err = toml.DecodeFile(filename, &parser.root)
	//log.Printf("decode %t -- %+v", fatal, err)
	if err != nil {
		return nil, err
	}

	parseSteps := []func() (bool, error){
		parser.parseBackends,
		parser.parseDefaults,
		parser.parseKeys,
		parser.populateUndefinedDefaults,
		parser.populateValuesWithDefaults,
	}

	for _, step := range parseSteps {
		fatal, err = step()
		if err != nil && fatal {
			return nil, err
		}
	}

	return parser.confg, nil
}
