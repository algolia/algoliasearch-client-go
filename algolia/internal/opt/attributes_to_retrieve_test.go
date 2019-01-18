package opt

import (
	"encoding/json"
	"testing"

	"github.com/algolia/algoliasearch-client-go/algolia/opt"
	"github.com/stretchr/testify/require"
)

func TestAttributesToRetrieve(t *testing.T) {
	for _, c := range []struct {
		opts     []interface{}
		expected opt.AttributesToRetrieveOption
	}{
		{
			opts:     []interface{}{nil},
			expected: opt.AttributesToRetrieveOption{},
		},
		{
			opts:     []interface{}{opt.AttributesToRetrieve()},
			expected: opt.AttributesToRetrieveOption{},
		},
		{
			opts:     []interface{}{opt.AttributesToRetrieve("attr1")},
			expected: opt.AttributesToRetrieve("attr1"),
		},
		{
			opts:     []interface{}{opt.AttributesToRetrieve("attr1", "attr2")},
			expected: opt.AttributesToRetrieve("attr1", "attr2"),
		},
		{
			opts: []interface{}{
				opt.AttributesToRetrieve("attr1", "attr2"),
				opt.AttributesToRetrieve("attr3", "attr4"),
			},
			expected: opt.AttributesToRetrieve("attr1", "attr2", "attr3", "attr4"),
		},
		{
			opts: []interface{}{
				opt.AttributesToRetrieve("attr1", "attr2"),
				opt.AttributesToRetrieve("attr1", "attr2"),
			},
			expected: opt.AttributesToRetrieve("attr1", "attr2"),
		},
		{
			opts: []interface{}{
				opt.AttributesToRetrieve("attr1", "attr2"),
				opt.AttributesToRetrieve("attr2", "attr3"),
			},
			expected: opt.AttributesToRetrieve("attr1", "attr2", "attr3"),
		},
	} {
		var (
			in  = ExtractAttributesToRetrieve(c.opts...)
			out opt.AttributesToRetrieveOption
		)
		data, err := json.Marshal(&in)
		require.NoError(t, err)
		err = json.Unmarshal(data, &out)
		require.NoError(t, err)
		require.ElementsMatch(t, c.expected.Get(), out.Get())
	}
}
