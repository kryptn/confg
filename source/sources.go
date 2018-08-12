package source

import (
	"errors"
	"github.com/kryptn/confg/containers"
	envSource "github.com/kryptn/confg/source/env"
)

type SourceClient interface {
	Register(*containers.Key)
	//Resolve()
	Collect() []*containers.Key
}

func GetSource(backend *containers.Backend) (SourceClient, error) {
	var client SourceClient
	var err error

	switch backend.Source {
	case "env":
		client, err = envSource.Get(backend)
	default:
		client, err = nil, errors.New("invalid source")

	}

	return client, err
}
