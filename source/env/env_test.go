package env_test

import (
	"github.com/kryptn/confg/containers"
	"github.com/kryptn/confg/source/env"
	"io/ioutil"
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
	Lookup(lookup string) (interface{}, bool)
}

func TestEnvSource_Lookup(t *testing.T) {

}

type kv struct {
	key, value string
}

type envGetTest struct {
	useFile      bool
	fileContents string
	envContents  []kv
	backend      *containers.Backend
	expectError  bool
	testKey      string
	expValue     string
}

var envGetTests = map[string]envGetTest{
	"get with file": {
		useFile:      true,
		fileContents: "a=c\n",
		envContents:  []kv{{"a", "b"}},
		backend:      &containers.Backend{Source: "env"},
		testKey:      "a",
		expValue:     "c",
	},
	"get with file -- invalid lines": {
		useFile:      true,
		fileContents: "a=c\nb\n#comment=true",
		envContents:  []kv{{"a", "b"}},
		backend:      &containers.Backend{Source: "env"},
		testKey:      "a",
		expValue:     "c",
	},
	"get with file -- invalid filename": {
		useFile:      true,
		fileContents: "a=c\nb\n#comment=true",
		envContents:  []kv{{"a", "b"}},
		backend:      &containers.Backend{Source: "env", EnvFile: "/tmp/112358"},
		expectError:  true,
	},
	"get without file": {
		useFile:     false,
		envContents: []kv{{"a", "b"}},
		backend:     &containers.Backend{Source: "env"},
		testKey:     "a",
		expValue:    "b",
	},
	"get unknown key": {
		useFile:     false,
		envContents: []kv{{"a", "b"}},
		backend:     &containers.Backend{Source: "env"},
		testKey:     "ccc",
		expectError: true,
	},
	"no source defined fail": {
		backend:     &containers.Backend{},
		expectError: true,
	},
}

func TestGet(t *testing.T) {
	for desc, test := range envGetTests {
		if test.useFile {
			tf, err := ioutil.TempFile("/tmp", "confg")
			if err != nil {
				t.Logf("Issue when making temp file")
				t.FailNow()
			}
			defer tf.Close()
			tf.WriteString(test.fileContents)
			if test.backend.EnvFile == "" {
				test.backend.EnvFile = tf.Name()
			}
		}
		for _, kv := range test.envContents {
			os.Setenv(kv.key, kv.value)
		}

		es, err := env.Get(test.backend)
		if err != nil && !test.expectError {
			t.Logf("%s -- failed on get", desc)
			t.Fail()
		}
		if test.expectError {
			continue
		}

		result, err := es.Lookup(test.testKey)
		if err != nil || result.(string) != test.expValue {
			t.Logf("%s -- failed on lookup", desc)
			t.Fail()
		}

	}

}
