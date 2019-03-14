package opt

import (
	"testing"

	"github.com/algolia/algoliasearch-client-go/algolia/opt"
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
			opt:      opt.SearchableAttributes("attr1", "attr2"),
			expected: []interface{}{opt.SearchableAttributes("attr1", "attr2")},
		},
		{
			inputs:   []interface{}{},
			opt:      opt.SearchableAttributes("attr1", "attr2"),
			expected: []interface{}{opt.SearchableAttributes("attr1", "attr2")},
		},
		{
			inputs:   []interface{}{opt.SearchableAttributes("attr1", "attr2")},
			opt:      nil,
			expected: []interface{}{opt.SearchableAttributes("attr1", "attr2")},
		},
		{
			inputs:   []interface{}{opt.SearchableAttributes("attr1", "attr2")},
			opt:      opt.Distinct(1),
			expected: []interface{}{opt.SearchableAttributes("attr1", "attr2"), opt.Distinct(1)},
		},
		{
			inputs:   []interface{}{opt.SearchableAttributes("attr1", "attr2")},
			opt:      opt.AttributesToRetrieve("attr1", "attr2"),
			expected: []interface{}{opt.SearchableAttributes("attr1", "attr2"), opt.AttributesToRetrieve("attr1", "attr2")},
		},
		{
			inputs:   []interface{}{opt.SearchableAttributes("attr1", "attr2")},
			opt:      opt.SearchableAttributes("attr3"),
			expected: []interface{}{opt.SearchableAttributes("attr3")},
		},
		{
			inputs:   []interface{}{opt.SearchableAttributes("attr1", "attr2"), opt.AttributesToRetrieve("attr1"), opt.Distinct(1)},
			opt:      opt.AttributesToRetrieve("attr3"),
			expected: []interface{}{opt.SearchableAttributes("attr1", "attr2"), opt.AttributesToRetrieve("attr3"), opt.Distinct(1)},
		},
	} {
		opts := InsertOrReplaceOption(c.inputs, c.opt)
		require.ElementsMatch(t, opts, c.expected, "inputs: %#v\nopt: %#v\nexpected: %#v", c.inputs, c.opt, c.expected)
	}
}
