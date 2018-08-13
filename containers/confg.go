package containers

type Confg struct {
	version string

	Backends []*Backend
	backends map[string]*Backend
	Keys     []*Key

	Reduced map[string]map[string]interface{}
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

func (c *Confg) GetBackend(name string) *Backend {
	if c.backends == nil {
		c.backends = map[string]*Backend{}
	}

	backend, ok := c.backends[name]
	if ok {
		return backend
	}

	for _, backend := range c.Backends {
		if backend.Name == name {
			c.backends[name] = backend
			return backend
		}
	}

	return nil
}

type keySetMapping map[string]map[string][]*Key
type reducedKeyMapping map[string]map[string]interface{}

func keyMappingAssigner(mapping keySetMapping) func(*Key) {
	return func(key *Key) {
		_, ok := mapping[key.Dest]
		if !ok {
			mapping[key.Dest] = map[string][]*Key{}
		}
		mapping[key.Dest][key.Key] = append(mapping[key.Dest][key.Key], key)
	}
}

func reducedKeyMappingAssigner(mapping reducedKeyMapping) func(string, string, interface{}) {
	return func(groupName, keyName string, v interface{}) {
		_, ok := mapping[groupName]
		if !ok {
			mapping[groupName] = map[string]interface{}{}
		}
		mapping[groupName][keyName] = v
	}
}

func (c *Confg) ReduceKeys() error {
	keySetMapping := make(keySetMapping)
	keySetAssigner := keyMappingAssigner(keySetMapping)

	for _, key := range c.Keys {
		key.Prepare()
		keySetAssigner(key)
	}

	reducedMapping := make(reducedKeyMapping)
	reducedAssigner := reducedKeyMappingAssigner(reducedMapping)
	for groupName, keySetMap := range keySetMapping {
		for keyName, keys := range keySetMap {

			if len(keys) > 0 {
				key := KeySet(keys).FirstValid()
				reducedAssigner(groupName, keyName, key.Value)
			}
		}
	}
	c.Reduced = reducedMapping

	return nil
}
