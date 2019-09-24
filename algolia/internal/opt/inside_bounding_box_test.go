package opt

import (
	"encoding/json"
	"testing"

	"github.com/algolia/algoliasearch-client-go/v3/algolia/opt"
	"github.com/stretchr/testify/require"
)

func TestInsideBoundingBox(t *testing.T) {
	for _, c := range []struct {
		opts     []interface{}
		expected *opt.InsideBoundingBoxOption
	}{
		{
			opts:     []interface{}{nil},
			expected: &opt.InsideBoundingBoxOption{},
		},
		{
			opts: []interface{}{opt.InsideBoundingBox(
				[][4]float64{
					{1.0, 2.0, 3.0, 4.0},
					{5.0, 6.0, 7.0, 8.0},
				},
			)},
			expected: opt.InsideBoundingBox(
				[][4]float64{
					{1.0, 2.0, 3.0, 4.0},
					{5.0, 6.0, 7.0, 8.0},
				},
			),
		},
		{
			opts: []interface{}{opt.InsideBoundingBoxFromCoordinates(
				"1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0",
			)},
			expected: opt.InsideBoundingBoxFromCoordinates(
				"1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0",
			),
		},
	} {
		var (
			in  = ExtractInsideBoundingBox(c.opts...)
			out opt.InsideBoundingBoxOption
		)
		data, err := json.Marshal(&in)
		require.NoError(t, err)
		err = json.Unmarshal(data, &out)
		require.NoError(t, err)
		require.Equal(t, *c.expected, out)
	}
}
