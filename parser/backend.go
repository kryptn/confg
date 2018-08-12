package parser

import (
	"github.com/BurntSushi/toml"
	"github.com/kryptn/confg/containers"
)

func parseBackends(prim toml.Primitive, md toml.MetaData) ([]*containers.Backend, error) {
	backends := map[string]*containers.Backend{}

	err := md.PrimitiveDecode(prim, &backends)
	if err != nil {
		return nil, err
	}

	be := []*containers.Backend{}
	for name, backend := range backends {
		backend.Name = name
		be = append(be, backend)
	}

	return be, nil

}
