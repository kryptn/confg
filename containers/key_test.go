package containers_test

import (
	"github.com/kryptn/confg/containers"
	"sort"
	"testing"
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