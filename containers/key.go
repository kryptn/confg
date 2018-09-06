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
		Reason string
	}
}

type ClientGetter func(string) (interface{}, error)

func (k *Key) Resolve(getter ClientGetter) {
	// attempt to resolve value -- if an error is returned consider
	// the resolution a failure but set the value if not nil
	value, err := getter(k.Lookup)

	if value != nil {
		k.Value = value
	}

	if err != nil {
		k.Meta.Reason = err.Error()
	}

	// consider resolved only if a returned value and no error
	k.Resolved = value != nil && err == nil
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
	// priority order:
	// first check resolved: if both, priority, otherwise the resolved one is true
	// then check if they have a value at all, if both then priority, if one use that one

	left, right := ks[i], ks[j]

	if left.Resolved && right.Resolved {
		return left.Priority > right.Priority
	}

	if left.Resolved || right.Resolved {
		return left.Resolved
	}

	if left.Value != nil && right.Value != nil {
		return left.Priority > right.Priority
	}

	if left.Value != nil || right.Value != nil {
		return left.Value != nil
	}

	return left.Priority > right.Priority
}

func (ks KeySet) FirstValid() *Key {
	if len(ks) > 0 {
		sort.Sort(ks)
		return ks[0]
	}
	return nil
}
