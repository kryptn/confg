package parser

import (
	"github.com/BurntSushi/toml"
)

type group struct {
	Name     string
	Backend  string
	Priority int
	Defaults map[string]interface{} `toml:"default"`
	Keys     map[string]toml.Primitive
}

func (p *Parser) parseGroup(name string, primitive toml.Primitive) (*group, error) {
	g := group{Name: name}
	if err := p.md.PrimitiveDecode(primitive, &g); err != nil {
		return nil, err
	}
	return &g, nil
}

func (p *Parser) parseGroups() map[string]*group {
	groups := map[string]*group{}

	for key, primitive := range p.root {
		if _, protected := protectedKeys[key]; protected {
			continue
		}
		group, err := p.parseGroup(key, primitive)
		if err != nil {
			continue
		}

		for keyName, defaultValue := range group.Defaults {
			if _, ok := p.defaults[group.Name]; !ok {
				p.defaults[group.Name] = map[string]interface{}{}
			}
			p.defaults[group.Name][keyName] = defaultValue
		}

		groups[key] = group
	}

	return groups
}
