package main

import (
	"github.com/kryptn/confg/containers"
	"github.com/kryptn/confg/source"
	"log"
)

func GatherAllKeys(confg containers.Confg) {
	backendKeyMap := map[*containers.Backend][]*containers.Key{}
	for _, key := range confg.Keys {
		backend := confg.GetBackend(key.Backend)
		backendKeyMap[backend] = append(backendKeyMap[backend], key)
	}

	for backend, keys := range backendKeyMap {
		err := source.Gather(backend, keys)
		if err != nil {
			log.Printf("Error when gathering from backend %s: %v", backend.Name, err)
		}
	}
}
