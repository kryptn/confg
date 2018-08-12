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

//func OmDecodeKey(g group, keyName string, prim toml.Primitive, md toml.MetaData) (kc *containers.Key, err error) {
//	key := keyFromGroup(g, keyName)
//
//	tomlKey := fmt.Sprintf("%s.keys.%s", g.Name, keyName)
//	switch md.Type(tomlKey) {
//	case "Hash":
//		log.Print("hash type")
//		err = md.PrimitiveDecode(prim, &key)
//	default:
//		log.Print("another type")
//		err = md.PrimitiveDecode(prim, &key.Lookup)
//	}
//	if err != nil {
//		return nil, err
//	}
//	return &key, nil
//}
