package parser

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/kryptn/confg/containers"
	"log"
)

func initKey(g *group, keyName string) containers.Key {
	key := containers.Key{
		Key:      keyName,
		Priority: g.Priority,
		Dest:     g.Name,
		Backend:  g.Backend,
	}

	return key
}

func (p *Parser) checkForDefault(key *containers.Key) {
	if _, ok := p.defaults[key.Dest]; ok {
		if value, ok := p.defaults[key.Dest][key.Key]; ok {
			key.Default = value
		}
	}
}

func (p *Parser) parseKeys() (fatal bool, err error) {
	groups := p.parseGroups()

	for _, group := range groups {
		for keyName, keyPrimitive := range group.Keys {
			key := initKey(group, keyName)

			err := decodeKey(&key, keyPrimitive, p.md)
			if err != nil {
				log.Printf("error on decode: %v", err)
				continue
			}

			if key.Default == nil {
				p.checkForDefault(&key)
			}

			p.confg.Keys = append(p.confg.Keys, &key)
		}
	}
	return false, nil
}

func decodeKey(key *containers.Key, primitive toml.Primitive, md toml.MetaData) (err error) {
	tomlKey := fmt.Sprintf("%s.keys.%s", key.Dest, key.Key)
	switch md.Type(tomlKey) {
	case "Hash":
		err = md.PrimitiveDecode(primitive, key)
	default:
		var lookup string
		err = md.PrimitiveDecode(primitive, &lookup)
		key.Lookup = lookup
	}
	return err
}
