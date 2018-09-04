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

func (c *Confg) Reduce() (*Confg, error) {
	c.ReduceKeys()
	return c, nil
}

type okays []bool

func (o okays) All() bool {
	for _, ok := range o {
		if !ok {
			return false
		}
	}
	return true
}

func (c *Confg) Validate() (bool, []error) {
	var oks []bool
	var errors []error
	for _, backend := range c.Backends {
		ok, errs := backend.Validate()
		oks = append(oks, ok)
		errors = append(errors, errs...)
	}
	for _, key := range c.Keys {
		ok, errs := key.Validate()
		oks = append(oks, ok)
		errors = append(errors, errs...)
	}

	return okays(oks).All(), errors

}
