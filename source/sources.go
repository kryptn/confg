package source

import (
	"errors"
	"github.com/kryptn/confg/containers"

	envSource "github.com/kryptn/confg/source/env"
	etcdSource "github.com/kryptn/confg/source/etcd"
)

var sources map[string]Client

func init() {
	sources = map[string]Client{}
}

type Client interface {
	Lookup(lookup string) (interface{}, bool)
}

func getSource(backend *containers.Backend) (Client, error) {
	var client Client
	var err error

	switch backend.Source {
	case "env":
		client, err = envSource.Get(backend)
	case "etcd":
		client, err = etcdSource.Get(backend)
	default:
		client, err = nil, errors.New("invalid source")
	}

	return client, err
}

func GetSource(backend *containers.Backend) (Client, error) {
	source, ok := sources[backend.Source]
	if ok {
		return source, nil
	}

	source, err := getSource(backend)
	if err != nil {
		return nil, err
	}
	sources[backend.Source] = source
	return source, nil
}
