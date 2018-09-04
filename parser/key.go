package parser

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/kryptn/confg/containers"
)

func keyFromGroup(g *group, keyName string) containers.Key {
	key := containers.Key{
		Key:      keyName,
		Priority: g.Priority,
		Dest:     g.Name,
		Backend:  g.Backend,
	}

	return key
}

func (p *Parser) parseKeys() (fatal bool, err error) {
	groups := p.parseGroups()

	for _, group := range groups {
		for keyName, keyPrimitive := range group.Keys {
			key, err := decodeKey(group, keyName, keyPrimitive, p.md)
			if err != nil {
				continue
			}
			p.confg.Keys = append(p.confg.Keys, key)
		}
	}
	return false, nil
}

func complexKey(g *group, keyName string, prim toml.Primitive, md toml.MetaData) (*containers.Key, error) {
	key := keyFromGroup(g, keyName)
	if err := md.PrimitiveDecode(prim, &key); err != nil {
		return nil, err
	}
	fmt.Printf("complx key -- name: %s, %+v\n", keyName, key)
	return &key, nil
}

func simpleKey(g *group, keyName string, prim toml.Primitive, md toml.MetaData) (*containers.Key, error) {
	key := keyFromGroup(g, keyName)
	var lookup string
	if err := md.PrimitiveDecode(prim, &lookup); err != nil {
		return nil, err
	}
	key.Lookup = lookup
	fmt.Printf("simple key -- name: %s, %+v\n", keyName, key)
	return &key, nil
}

func decodeKey(g *group, keyName string, prim toml.Primitive, md toml.MetaData) (key *containers.Key, err error) {
	tomlKey := fmt.Sprintf("%s.keys.%s", g.Name, keyName)
	switch md.Type(tomlKey) {
	case "Hash":
		key, err = complexKey(g, keyName, prim, md)
	default:
		key, err = simpleKey(g, keyName, prim, md)
	}
	return key, err
}
