package opt

import (
	"encoding/json"
	"testing"

	"github.com/algolia/algoliasearch-client-go/v3/algolia/opt"
	"github.com/stretchr/testify/require"
)

func TestInsidePolygon(t *testing.T) {
	for _, c := range []struct {
		opts     []interface{}
		expected *opt.InsidePolygonOption
	}{
		{
			opts:     []interface{}{nil},
			expected: &opt.InsidePolygonOption{},
		},
		{
			opts: []interface{}{opt.InsidePolygon(
				[][]float64{
					{1.0, 2.0, 3.0, 4.0},
					{5.0, 6.0, 7.0, 8.0},
				},
			)},
			expected: opt.InsidePolygon(
				[][]float64{
					{1.0, 2.0, 3.0, 4.0},
					{5.0, 6.0, 7.0, 8.0},
				},
			),
		},
		{
			opts: []interface{}{opt.InsidePolygonFromCoordinates(
				"1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0",
			)},
			expected: opt.InsidePolygonFromCoordinates(
				"1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0",
			),
		},
	} {
		var (
			in  = ExtractInsidePolygon(c.opts...)
			out opt.InsidePolygonOption
		)
		data, err := json.Marshal(&in)
		require.NoError(t, err)
		err = json.Unmarshal(data, &out)
		require.NoError(t, err)
		require.Equal(t, *c.expected, out)
	}
}
