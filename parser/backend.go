package parser

import (
	"errors"
)

func (p *Parser) parseBackends() (fatal bool, err error) {
	backendPrimitive, ok := p.root["backend"]
	if !ok {
		return false, errors.New("parser.backend: no backend defined")
	}
	err = p.md.PrimitiveDecode(backendPrimitive, &p.confg.Backends)
	if err != nil {
		return false, err
	}

	for _, backend := range p.confg.Backends {
		backend.ConfigPath = p.path
	}

	return false, nil
}
