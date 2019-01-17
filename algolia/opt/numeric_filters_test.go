package opt

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNumericFilters(t *testing.T) {
	// Unicode encoding is used for `<`, `>` symbols in the following expected
	// JSON strings. This is due to the default behaviour of the JSON encoder
	// from the Go standard library. More informations about this in the
	// official documentation: https://golang.org/pkg/encoding/json/#Marshal.
	//
	// As far as the Algolia API is concerned, it doesn't change anything. The
	// Algolia API accepts and understands both strings in the same way, e.g.
	// characters `<` and `>` are equivalent to `\u003c` and `\u003e`
	// respectively.
	for _, c := range []struct {
		opts         []interface{}
		expectedJSON string
	}{
		{
			opts:         []interface{}{nil},
			expectedJSON: "null",
		},
		{
			opts:         []interface{}{NumericFilters("inStock > 0")},
			expectedJSON: `"inStock \u003e 0"`,
		},
		{
			opts: []interface{}{
				NumericFiltersAnd(
					NumericFilters("inStock > 0"),
					NumericFilters("price < 1000"),
				),
			},
			expectedJSON: `[["inStock \u003e 0"],["price \u003c 1000"]]`,
		},
		{
			opts: []interface{}{
				NumericFiltersOr(
					NumericFilters("inStock > 0"),
					NumericFilters("price < 1000"),
				),
			},
			expectedJSON: `[["inStock \u003e 0","price \u003c 1000"]]`,
		},
		{
			opts: []interface{}{
				NumericFiltersAnd(
					NumericFiltersOr(
						NumericFilters("inStock > 0"),
						NumericFilters("price < 1000"),
					),
				),
			},
			expectedJSON: `[["inStock \u003e 0","price \u003c 1000"]]`,
		},
		{
			opts: []interface{}{
				NumericFiltersAnd(
					NumericFiltersOr(
						NumericFilters("inStock < 1000"),
						NumericFilters("deliveryDate < 1441755506"),
					),
					NumericFiltersOr(
						NumericFilters("inStock > 0"),
						NumericFilters("price < 1000"),
					),
				),
			},
			expectedJSON: `[["inStock \u003c 1000","deliveryDate \u003c 1441755506"],["inStock \u003e 0","price \u003c 1000"]]`,
		},
		{
			opts: []interface{}{
				NumericFiltersAnd(
					NumericFilters("deliveryDate < 1441755506"),
					NumericFiltersOr(
						NumericFilters("inStock > 0"),
						NumericFilters("price < 1000"),
					),
				),
			},
			expectedJSON: `[["deliveryDate \u003c 1441755506"],["inStock \u003e 0","price \u003c 1000"]]`,
		},
	} {
		res := ExtractNumericFilters(c.opts...)
		data, err := json.Marshal(&res)
		require.NoError(t, err)
		require.Equal(t, c.expectedJSON, string(data))
	}
}
