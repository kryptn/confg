package containers

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
