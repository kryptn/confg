package parser

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/kryptn/confg/containers"
	"log"
)

type Parsed struct {
	Backends []*containers.Backend
	Keys     []*containers.Key

	md   toml.MetaData
	root map[string]toml.Primitive
}

func (p *Parsed) Print() {

}

func verbosePrintMetadataKeys(md toml.MetaData) {
	log.Print("printing keys from metadata")
	for _, key := range md.Keys() {
		log.Printf("\tkey: %v, %v", key, md.Type(key.String()))
	}
}

func verbosePrintUndecodedKeys(md toml.MetaData) {
	log.Print("printing undecoded keys")
	for _, key := range md.Undecoded() {
		log.Printf("\tkey: %v, %v", key, md.Type(key.String()))
	}
}

var protectedKeys = map[string]struct{}{"backend": {}}

func (p *Parsed) Parse() error {
	backendsPrimitive, ok := p.root["backend"]
	if !ok {
		log.Fatalf("No backends found in config")
	}
	backends, err := parseBackends(backendsPrimitive, p.md)
	if err != nil {
		log.Fatalf("broke decoding backends %v", err)
	}
	p.Backends = backends

	groups, err := p.parseGroups()
	if err != nil {
		log.Fatalf("broke %v", err)
	}

	keys := []*containers.Key{}
	for _, group := range groups {
		for keyName, primitive := range group.Keys {
			key, err := decodeKey(group, keyName, primitive, p.md)
			if err != nil {
				continue
			}
			fmt.Printf("decoded key -- name: %s, %+v\n", keyName, key)
			keys = append(keys, key)
		}
	}
	p.Keys = keys
	//log.Printf("undecoded: %+s", p.md.Undecoded())

	return nil
}

func ParsedFromFile(filename string) (parsed *Parsed, err error) {
	parsed = &Parsed{}
	parsed.md, err = toml.DecodeFile(filename, &parsed.root)
	if err != nil {
		return nil, err
	}
	return parsed, nil

}
