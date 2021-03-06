package containers_test

import (
	"github.com/pkg/errors"
	"sort"
	"testing"

	"github.com/kryptn/confg/containers"
)

var priorityTest = []*containers.Key{
	{Key: "pass", Priority: 10},
	{Key: "fail", Priority: 0},
}

var resolvedTest = []*containers.Key{
	{Key: "pass", Priority: 10, Resolved: true},
	{Key: "fail", Priority: 20},
}

var samePriorityTest = []*containers.Key{
	{Key: "pass", Priority: 10, Resolved: true},
	{Key: "fail", Priority: 10, Resolved: true},
}

var samePriorityFirstTest = []*containers.Key{
	{Key: "fail", Priority: 10, Resolved: false},
	{Key: "pass", Priority: 10, Resolved: true},
	{Key: "fail", Priority: 10, Resolved: true},
}

var manyConditions = []*containers.Key{
	{Key: "fail", Priority: 15, Resolved: false},
	{Key: "fail", Priority: 10, Resolved: true},
	{Key: "fail", Priority: 19, Resolved: true},
	{Key: "pass", Priority: 20, Resolved: true},
	{Key: "fail", Priority: 20, Resolved: false},
	{Key: "fail", Priority: 18, Resolved: true},
	{Key: "fail", Priority: 17, Resolved: false},
}

var withDefaults = []*containers.Key{
	{Key: "fail", Priority: 10, Resolved: false, Value: true},
	{Key: "fail", Priority: 10, Resolved: false, Value: nil},
	{Key: "pass", Priority: 20, Resolved: true, Value: true},
	{Key: "fail", Priority: 20, Resolved: false, Value: nil},
	{Key: "fail", Priority: 30, Resolved: false, Value: nil},
	{Key: "fail", Priority: 30, Resolved: false, Value: true},
}

type sortTest struct {
	desc string
	keys []*containers.Key
}

var sortTests = []sortTest{
	{"highest priority first", priorityTest},
	{"resolved over unresolved", resolvedTest},
	{"first declared if same priority", samePriorityTest},
	{"first declared first resolved if same priority", samePriorityFirstTest},
	{"nothing specific", manyConditions},
	{"lower priority resolved over higher unresolved with value", withDefaults},
}

func TestKeySet_Less(t *testing.T) {
	for _, test := range sortTests {
		sort.Sort(containers.KeySet(test.keys))
		highest := test.keys[0]
		if highest.Key != "pass" {
			t.Logf("%s: Expected \"pass\", got %s", test.desc, highest.Key)
			t.Fail()
		}
	}
}

func TestKeySet_FirstValid(t *testing.T) {
	for _, test := range sortTests {
		keys := containers.KeySet(test.keys)
		first := keys.FirstValid()
		if first.Key != "pass" {
			t.Logf("%s: Expected \"pass\", got %s", test.desc, first.Key)
			t.Fail()
		}
	}

	ks := containers.KeySet{}
	first := ks.FirstValid()
	if first != nil {
		t.Logf("empty keyset should return nil")
		t.Fail()
	}
}

func getterReturns(v interface{}, err error) func(string) (v interface{}, err error) {
	return func(s string) (interface{}, error) {
		return v, err
	}
}

type Result struct {
	value    interface{}
	resolved bool
}

type resolverTest struct {
	key    *containers.Key
	getter func(string) (v interface{}, err error)
	result Result
}

var resolverTests = map[string]resolverTest{
	"resolves value": {
		key:    &containers.Key{},
		getter: getterReturns("a", nil),
		result: Result{"a", true},
	},
	"value and error": {
		key:    &containers.Key{},
		getter: getterReturns("a", errors.New("err")),
		result: Result{"a", false},
	},
	"error": {
		key:    &containers.Key{},
		getter: getterReturns(nil, errors.New("err")),
		result: Result{nil, false},
	},
	"default but resolves": {
		key:    &containers.Key{Value: false},
		getter: getterReturns(true, nil),
		result: Result{true, true},
	},
	"defaults": {
		key:    &containers.Key{Value: false},
		getter: getterReturns(nil, errors.New("err")),
		result: Result{false, false},
	},
}

func tFuncForResolve(test resolverTest) func(t *testing.T) {
	return func(t *testing.T) {
		key := test.key
		key.Resolve(test.getter)

		if key.Value != test.result.value {
			t.Logf("%s: expected %v got %v", t.Name(), test.result.value, key.Value)
			t.Fail()
		}

		if key.Resolved != test.result.resolved {
			t.Logf("%s: expected %v got %v", t.Name(), test.result.resolved, key.Resolved)
			t.Fail()
		}
	}
}

func TestKey_Resolve(t *testing.T) {
	for desc, test := range resolverTests {
		t.Run(desc, tFuncForResolve(test))
	}
}
