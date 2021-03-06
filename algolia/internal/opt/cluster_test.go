// Code generated by go generate. DO NOT EDIT.

package opt

import (
	"encoding/json"
	"testing"

	"github.com/algolia/algoliasearch-client-go/v3/algolia/opt"
	"github.com/stretchr/testify/require"
)

func TestCluster(t *testing.T) {
	for _, c := range []struct {
		opts     []interface{}
		expected *opt.ClusterOption
	}{
		{
			opts:     []interface{}{nil},
			expected: opt.Cluster(""),
		},
		{
			opts:     []interface{}{opt.Cluster("")},
			expected: opt.Cluster(""),
		},
		{
			opts:     []interface{}{opt.Cluster("content of the string value")},
			expected: opt.Cluster("content of the string value"),
		},
	} {
		var (
			in  = ExtractCluster(c.opts...)
			out opt.ClusterOption
		)
		data, err := json.Marshal(&in)
		require.NoError(t, err)
		err = json.Unmarshal(data, &out)
		require.NoError(t, err)
		require.Equal(t, *c.expected, out)
	}
}
