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

	fatal, err = parser.parseBackends()
	//log.Printf("parseBackends %t -- %+v", fatal, err)
	if err != nil && fatal {
		return nil, err
	}

	fatal, err = parser.parseDefaults()
	if err != nil && fatal {
		return nil, err
	}

	fatal, err = parser.parseKeys()
	//log.Printf("parseKeys %t -- %+v", fatal, err)
	if err != nil && fatal {
		return nil, err
	}

	fatal, err = parser.populateUndefinedDefaults()
	if err != nil && fatal {
		return nil, err
	}

	fatal, err = parser.populateValuesWithDefaults()
	if err != nil && fatal {
		return nil, err
	}

	return parser.confg, nil
}
