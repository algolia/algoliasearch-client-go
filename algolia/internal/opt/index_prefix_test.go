// Code generated by go generate. DO NOT EDIT.

package opt

import (
	"encoding/json"
	"testing"

	"github.com/algolia/algoliasearch-client-go/v3/algolia/opt"
	"github.com/stretchr/testify/require"
)

func TestIndexPrefix(t *testing.T) {
	for _, c := range []struct {
		opts     []interface{}
		expected *opt.IndexPrefixOption
	}{
		{
			opts:     []interface{}{nil},
			expected: opt.IndexPrefix(""),
		},
		{
			opts:     []interface{}{opt.IndexPrefix("")},
			expected: opt.IndexPrefix(""),
		},
		{
			opts:     []interface{}{opt.IndexPrefix("content of the string value")},
			expected: opt.IndexPrefix("content of the string value"),
		},
	} {
		var (
			in  = ExtractIndexPrefix(c.opts...)
			out opt.IndexPrefixOption
		)
		data, err := json.Marshal(&in)
		require.NoError(t, err)
		err = json.Unmarshal(data, &out)
		require.NoError(t, err)
		require.Equal(t, *c.expected, out)
	}
}