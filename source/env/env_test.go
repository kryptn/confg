package env_test

import (
	"github.com/kryptn/confg/containers"
	"github.com/kryptn/confg/source/env"
	"os"
	"testing"
)

type EnvKeyTest struct {
	target   string
	given    string
	set      bool
	key      *containers.Key
	expected interface{}
}

var envKeyTests = []EnvKeyTest{
	{
		// tests happy path
		"env_name_a", "text here", true,
		&containers.Key{
			Key:    "a",
			Lookup: "env_name_a",
		},
		"text here",
	}, {
		// tests default given
		target: "env_name_b",
		key: &containers.Key{
			Key:     "a",
			Lookup:  "env_name_b",
			Default: "default"},
		expected: "default",
	},
}

type testSourceClient interface {
	Register(key *containers.Key)
	Collect() []*containers.Key
}

func TestEnvSource_Collect(t *testing.T) {
	for _, test := range envKeyTests {
		if test.set {
			os.Setenv(test.target, test.given)
		}

		es, _ := env.Get(&containers.Backend{Source: "env"})

		es.Register(test.key)
		result := es.Collect()
		if result[0].Value != test.expected {
			t.Logf("Expected %v got %v", test.expected, result[0].Value)
			t.Fail()
		}

	}
}

type envGetTests struct {

}

func TestGet(t *testing.T) {



}
