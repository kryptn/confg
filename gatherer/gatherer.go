package gatherer

import (
	"fmt"
	"github.com/kryptn/confg/containers"
	"github.com/kryptn/confg/source"
)

type SourceClient interface {
	Lookup(lookup string) (value interface{}, err error)
}

type Gatherer struct {
	confg *containers.Confg

	sources map[string]SourceClient
}

func NewGatherer(confg *containers.Confg) *Gatherer {
	return &Gatherer{
		confg:   confg,
		sources: map[string]SourceClient{},
	}
}

func (g *Gatherer) Resolve() (*containers.Confg, error) {
	for _, key := range g.confg.Keys {
		backend, ok := g.confg.Backends[key.Backend]
		if !ok {
			key.Resolved = false
			key.Meta.Reason = fmt.Sprintf("backend %s not declared", key.Backend)
			continue
		}

		client, err := source.GetSource(backend)
		if err != nil {
			key.Resolved = false
			key.Meta.Reason = fmt.Sprintf("source %s not declared", backend.Source)
			continue
		}

		key.Resolve(client.Lookup)
	}

	return g.confg, nil
}
