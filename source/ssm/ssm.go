package ssm

import (
	"errors"
	"github.com/kryptn/confg/containers"
)

type SsmSource struct {
}

func (ss *SsmSource) Lookup(lookup string) (interface{}, bool) {
	return nil, false
}

func (ss *SsmSource) Gather(keys []*containers.Key) {
	for _, key := range keys {
		v, ok := ss.Lookup(key.Lookup)
		key.Inject(v, ok)
	}
}

func Get(backend *containers.Backend) (*SsmSource, error) {
	if backend.Source != "ssm" {
		return nil, errors.New("source.ssm invalid backend")
	}
	ss := SsmSource{}

	return &ss, nil

}
