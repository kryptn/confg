package gatherer

import (
	"github.com/kryptn/confg/containers"
	"github.com/kryptn/confg/source"
)

type SourceClient interface {
	Lookup(lookup string) (value interface{}, ok bool)
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
			continue
		}

		client, err := source.GetSource(backend)
		if err != nil {
			key.Resolved = false
			continue
		}

		v, ok := client.Lookup(key.Lookup)
		key.Inject(v, ok)
	}

	return g.confg, nil
}
