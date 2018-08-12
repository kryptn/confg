package parser

import "github.com/BurntSushi/toml"

type group struct {
	Name     string
	Backend  string
	Priority int
	Keys     map[string]toml.Primitive
}

func (p *Parsed) parseGroup(name string, primitive toml.Primitive) (*group, error) {
	g := group{Name: name}
	if err := p.md.PrimitiveDecode(primitive, &g); err != nil {
		return nil, err
	}
	return &g, nil
}

func (p *Parsed) parseGroups() (map[string]*group, error) {
	groups := map[string]*group{}
	for name, primitive := range p.root {
		if _, protected := protectedKeys[name]; protected {
			continue
		}
		parsedGroup, err := p.parseGroup(name, primitive)
		if err != nil {
			continue
		}
		groups[name] = parsedGroup
	}
	return groups, nil
}
