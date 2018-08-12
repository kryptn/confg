package main

import (
	"github.com/kryptn/confg/containers"
	"github.com/kryptn/confg/parser"
	"github.com/kryptn/confg/source"
	"log"
)

type SourceClient interface {
	Register(*containers.Key)
	//Resolve()
	Collect() []*containers.Key
}

func confgFromParsed(p *parser.Parsed) (*containers.Confg, error) {

	confg := containers.Confg{
		Backends: p.Backends,
		Keys:     p.Keys}

	backendMap := map[string]*containers.Backend{}
	backendSourceMap := map[string]SourceClient{}

	for _, backend := range confg.Backends {
		client, err := source.GetSource(backend)
		if err != nil {
			return nil, err
		}
		backendSourceMap[backend.Name] = client
		backendMap[backend.Name] = backend
	}

	for _, key := range confg.Keys {
		client, ok := backendSourceMap[key.Backend]
		if !ok {
			log.Fatalf("Undefined backend %s", key.Backend)
		}
		client.Register(key)
	}

	var resolvedKeys []*containers.Key
	for _, client := range backendSourceMap {
		resolvedKeys = append(resolvedKeys, client.Collect()...)
	}

	renderedKeys := map[string]map[string]interface{}{}
	for _, key := range confg.Keys {
		log.Printf("%s %s %s", key.Key, key.Value, key.Backend)
		_, ok := renderedKeys[key.Dest]
		if !ok {
			renderedKeys[key.Dest] = map[string]interface{}{}
		}
		renderedKeys[key.Dest][key.Key] = key.Value
	}
	confg.Rendered = renderedKeys

	return &confg, nil
}
