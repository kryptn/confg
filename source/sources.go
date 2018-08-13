package source

import (
	"errors"
	"github.com/kryptn/confg/containers"
	envSource "github.com/kryptn/confg/source/env"
	etcdSource "github.com/kryptn/confg/source/etcd"
	"log"
)

type SourceClient interface {
	Gather([]*containers.Key)
	Lookup(lookup string) (interface{}, bool)
}

func GetSource(backend *containers.Backend) (SourceClient, error) {
	var client SourceClient
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

var sources = map[string]SourceClient{}

func getOrInitSource(backend *containers.Backend) (SourceClient, error) {
	source, ok := sources[backend.Source]
	if ok {
		return source, nil
	}

	source, err := GetSource(backend)
	if err != nil {
		return nil, err
	}
	sources[backend.Source] = source
	return source, nil
}

func Gather(backend *containers.Backend, keys []*containers.Key) error {
	source, err := getOrInitSource(backend)
	if err != nil {
		return err
	}

	log.Printf("source %+v", source)
	log.Printf("keys %+v", keys)
	for _, key := range keys {
		log.Printf("-- key %+v", key)
	}
	source.Gather(keys)
	return nil
}
