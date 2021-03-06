package adapter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilters(t *testing.T) {
	tests := []struct {
		start    state
		filters  *metricFilters
		expected state
	}{
		{
			state{action: actIgnore, name: "test"},
			nil,
			state{action: actIgnore, name: "test"},
		},
		{
			state{action: actIgnore, name: "test"},
			makeFilters(),
			state{action: actIgnore, name: "test"},
		},
		{
			state{action: actIgnore, name: "test"},
			makeFilters(
				WhitelistIf(func(_ string) bool { return true }),
			),
			state{action: actAccept, name: "test"},
		},
		{
			state{action: actIgnore, name: "test"},
			makeFilters(
				WhitelistIf(func(_ string) bool { return false }),
			),
			state{action: actIgnore, name: "test"},
		},
		{
			state{action: actIgnore, name: "test"},
			makeFilters(Whitelist("other")),
			state{action: actIgnore, name: "test"},
		},
		{
			state{action: actIgnore, name: "test"},
			makeFilters(Whitelist("test")),
			state{action: actAccept, name: "test"},
		},
		{
			state{action: actIgnore, name: "test"},
			makeFilters(Rename("test", "new")),
			state{action: actAccept, name: "new"},
		},
		{
			state{action: actIgnore, name: "t-e-s-t"},
			makeFilters(NameReplace("-", ".")),
			state{action: actIgnore, name: "t.e.s.t"},
		},
		{
			state{action: actIgnore, name: "test"},
			makeFilters(ToUpperName),
			state{action: actIgnore, name: "TEST"},
		},
		{
			state{action: actIgnore, name: "TEST"},
			makeFilters(ToLowerName),
			state{action: actIgnore, name: "test"},
		},
	}

	for i, test := range tests {
		t.Logf("run test (%v): %v => %v", i, test.start, test.expected)

		actual := test.filters.apply(test.start)
		assert.Equal(t, test.expected, actual)
	}
}
