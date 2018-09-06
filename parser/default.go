package parser

import (
	"errors"
	"fmt"
	"github.com/kryptn/confg/containers"
)

func (p *Parser) parseDefaults() (fatal bool, err error) {
	defaultsPrimitive, ok := p.root["default"]
	if !ok {
		return false, errors.New("parser.defaults: no root defaults defined")
	}

	err = p.md.PrimitiveDecode(defaultsPrimitive, &p.defaults)
	if err != nil {
		return false, err
	}

	return false, nil
}

var empty = struct{}{}

func (p *Parser) populateUndefinedDefaults() (fatal bool, err error) {

	definedKeys := make(map[string]struct{})

	for _, key := range p.confg.Keys {
		defined := fmt.Sprintf("%s.%s", key.Dest, key.Key)
		definedKeys[defined] = empty
	}

	for groupName, group := range p.defaults {
		for keyName, defaultValue := range group {
			defined := fmt.Sprintf("%s.%s", groupName, keyName)
			if _, ok := definedKeys[defined]; !ok {
				key := containers.Key{
					Key:     keyName,
					Default: defaultValue,
					Dest:    groupName,
				}

				p.confg.Keys = append(p.confg.Keys, &key)
			}
		}
	}

	return false, nil
}

func (p *Parser) populateValuesWithDefaults() (fatal bool, err error) {
	for _, key := range p.confg.Keys {
		if key.Default != nil {
			key.Value = key.Default
		}
	}
	return false, nil
}
