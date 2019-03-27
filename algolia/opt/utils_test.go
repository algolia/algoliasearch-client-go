package opt

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestInsertOrReplaceOption(t *testing.T) {
	for _, c := range []struct {
		inputs   []interface{}
		opt      interface{}
		expected []interface{}
	}{
		{
			inputs:   nil,
			opt:      nil,
			expected: nil,
		},
		{
			inputs:   []interface{}{},
			opt:      nil,
			expected: []interface{}{},
		},
		{
			inputs:   nil,
			opt:      SearchableAttributes("attr1", "attr2"),
			expected: []interface{}{SearchableAttributes("attr1", "attr2")},
		},
		{
			inputs:   []interface{}{},
			opt:      SearchableAttributes("attr1", "attr2"),
			expected: []interface{}{SearchableAttributes("attr1", "attr2")},
		},
		{
			inputs:   []interface{}{SearchableAttributes("attr1", "attr2")},
			opt:      nil,
			expected: []interface{}{SearchableAttributes("attr1", "attr2")},
		},
		{
			inputs:   []interface{}{SearchableAttributes("attr1", "attr2")},
			opt:      Distinct(true),
			expected: []interface{}{SearchableAttributes("attr1", "attr2"), Distinct(true)},
		},
		{
			inputs:   []interface{}{SearchableAttributes("attr1", "attr2")},
			opt:      AttributesToRetrieve("attr1", "attr2"),
			expected: []interface{}{SearchableAttributes("attr1", "attr2"), AttributesToRetrieve("attr1", "attr2")},
		},
		{
			inputs:   []interface{}{SearchableAttributes("attr1", "attr2")},
			opt:      SearchableAttributes("attr3"),
			expected: []interface{}{SearchableAttributes("attr3")},
		},
		{
			inputs:   []interface{}{SearchableAttributes("attr1", "attr2"), AttributesToRetrieve("attr1"), Distinct(true)},
			opt:      AttributesToRetrieve("attr3"),
			expected: []interface{}{SearchableAttributes("attr1", "attr2"), AttributesToRetrieve("attr3"), Distinct(true)},
		},
	} {
		opts := InsertOrReplaceOption(c.inputs, c.opt)
		require.ElementsMatch(t, opts, c.expected, "inputs: %#v\nopt: %#v\nexpected: %#v", c.inputs, c.opt, c.expected)
	}
}
