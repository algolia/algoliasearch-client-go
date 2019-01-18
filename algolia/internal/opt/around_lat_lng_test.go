package opt

import (
	"encoding/json"
	"testing"

	"github.com/algolia/algoliasearch-client-go/algolia/opt"
	"github.com/stretchr/testify/require"
)

func TestAroundLatLng(t *testing.T) {
	for _, c := range []struct {
		opts     []interface{}
		expected opt.AroundLatLngOption
	}{
		{
			opts:     []interface{}{nil},
			expected: opt.AroundLatLng(""),
		},
		{
			opts:     []interface{}{opt.AroundLatLng("40.71, -74.01")},
			expected: opt.AroundLatLng("40.71, -74.01"),
		},
	} {
		var (
			in  = ExtractAroundLatLng(c.opts...)
			out opt.AroundLatLngOption
		)
		data, err := json.Marshal(&in)
		require.NoError(t, err)
		err = json.Unmarshal(data, &out)
		require.NoError(t, err)
		require.Equal(t, c.expected, out)
	}
}
