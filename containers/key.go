package containers

import (
	"errors"
	"sort"
)

type Key struct {
	Key      string
	Lookup   string
	Priority int

	Default interface{}
	Value   interface{}

	Resolved bool

	Dest    string
	Backend string

	Meta struct {
		reason string
	}
}

func (k *Key) Prepare() {
	if k.Default != nil {
		k.Value = k.Default
	}
}

func (k *Key) Inject(v interface{}, ok bool) {
	k.Resolved = ok
	if ok {
		k.Value = v
	}
}

func (k *Key) Validate() (bool, []error) {
	ok := true
	errs := []error{}
	if k.Key == "" {
		ok = false
		errs = append(errs, errors.New("Key must be defined"))
	}
	if k.Lookup == "" {
		ok = false
		errs = append(errs, errors.New("Looup must be defined"))
	}
	return ok, errs
}

type KeySet []*Key

func (ks KeySet) Len() int {
	return len(ks)
}

func (ks KeySet) Swap(i, j int) {
	ks[i], ks[j] = ks[j], ks[i]
}

func (ks KeySet) Less(i, j int) bool {
	if ks[i].Resolved != ks[j].Resolved {
		if ks[j].Resolved {
			return false
		}
		return true
	}
	return ks[i].Priority > ks[j].Priority
}

func (ks KeySet) FirstValid() *Key {
	if len(ks) > 0 {
		sort.Sort(ks)
		return ks[0]
	}
	return nil
}
