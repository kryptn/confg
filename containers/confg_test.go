package containers_test

import (
	"github.com/kryptn/confg/containers"
	"testing"
)

func TestConfg_Overlay(t *testing.T) {
	c1 := &containers.Confg{
		Backends: map[string]*containers.Backend{
			"a": {Name: "a", Source: "env"},
		},
		Keys: []*containers.Key{
			{Key: "a", Dest: "a"},
		},
	}
	c2 := &containers.Confg{
		Backends: map[string]*containers.Backend{
			"a": {Name: "a", Source: "env"},
		},
		Keys: []*containers.Key{
			{Key: "a", Dest: "a"},
		},
	}

	combined := (&containers.Confg{}).Overlay(c1, c2)

	if len(combined.Backends) != 1 {
		t.Fail()
	}
	if len(combined.Keys) != 2 {
		t.Fail()
	}

}
