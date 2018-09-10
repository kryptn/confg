package containers

type Confg struct {
	version string

	Backends map[string]*Backend
	Keys     []*Key

	Reduced map[string]map[string]interface{}
}

func (c *Confg) overlay(other *Confg) *Confg {
	if c.Backends == nil {
		c.Backends = map[string]*Backend{}
	}

	for key, value := range other.Backends {
		c.Backends[key] = value
	}
	c.Keys = append(c.Keys, other.Keys...)

	return c
}

func (c *Confg) Overlay(others ...*Confg) *Confg {
	for _, other := range others {
		c.overlay(other)
	}
	return c
}
