package containers

type Key struct {
	Key      string
	Lookup   string
	Priority int

	Default interface{}
	Value   interface{}

	Resolved bool

	Dest    string
	Backend string
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
