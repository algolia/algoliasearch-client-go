package opt

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAttributesToRetrieve(t *testing.T) {
	for _, c := range []struct {
		opts     []interface{}
		expected []string
	}{
		{
			opts:     []interface{}{nil},
			expected: nil,
		},
		{
			opts:     []interface{}{AttributesToRetrieve(nil)},
			expected: nil,
		},
		{
			opts: []interface{}{
				AttributesToRetrieve([]string{"attr1", "attr2"}),
			},
			expected: []string{"attr1", "attr2"},
		},
		{
			opts: []interface{}{
				AttributesToRetrieve([]string{"attr1", "attr2"}),
				AttributesToRetrieve([]string{"attr3", "attr4"}),
			},
			expected: []string{"attr1", "attr2", "attr3", "attr4"},
		},
		{
			opts: []interface{}{
				AttributesToRetrieve([]string{"attr1", "attr2"}),
				AttributesToRetrieve([]string{"attr1", "attr2"}),
			},
			expected: []string{"attr1", "attr2"},
		},
	} {
		res := ExtractAttributesToRetrieve(c.opts...)
		require.ElementsMatch(t, c.expected, res)
	}
}
