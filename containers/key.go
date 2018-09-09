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
	// sorting so less is more, so higher preference
	left, right := ks[i], ks[j]

	preferPriority := func() bool {
		return left.Priority > right.Priority
	}

	if left.Resolved && right.Resolved {
		// if both resolved, take highest priority
		return preferPriority()
	}

	if left.Resolved || right.Resolved {
		// prefer the resolved value
		return left.Resolved
	}

	if left.Value != nil && right.Value != nil {
		// if both unresolved (guaranteed) and both
		// don't have a default defined
		return preferPriority()
	}

	if left.Value != nil || right.Value != nil {
		// prefer the one with the default defined
		return left.Value != nil
	}

	return preferPriority()
}

func (ks KeySet) FirstValid() *Key {
	if len(ks) > 0 {
		sort.Sort(ks)
		return ks[0]
	}
	return nil
}
