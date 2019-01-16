package opt

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFacetFilters(t *testing.T) {
	for _, c := range []struct {
		opts         []interface{}
		expectedJSON string
	}{
		{
			opts:         []interface{}{nil},
			expectedJSON: "null",
		},
		{
			opts:         []interface{}{FacetFilters("filter1", "value1")},
			expectedJSON: `"filter1:value1"`,
		},
		{
			opts: []interface{}{
				FacetFiltersAnd(
					FacetFilters("filter1", "value1"),
					FacetFilters("filter2", "value2"),
				),
			},
			expectedJSON: `[["filter1:value1"],["filter2:value2"]]`,
		},
		{
			opts: []interface{}{
				FacetFiltersOr(
					FacetFilters("filter1", "value1"),
					FacetFilters("filter2", "value2"),
				),
			},
			expectedJSON: `[["filter1:value1","filter2:value2"]]`,
		},
		{
			opts: []interface{}{
				FacetFiltersAnd(
					FacetFiltersOr(
						FacetFilters("filter1", "value1"),
						FacetFilters("filter2", "value2"),
					),
				),
			},
			expectedJSON: `[["filter1:value1","filter2:value2"]]`,
		},
		{
			opts: []interface{}{
				FacetFiltersAnd(
					FacetFiltersOr(
						FacetFilters("filter1", "value1"),
						FacetFilters("filter2", "value2"),
					),
					FacetFiltersOr(
						FacetFilters("filter3", "value3"),
						FacetFilters("filter4", "value4"),
					),
				),
			},
			expectedJSON: `[["filter1:value1","filter2:value2"],["filter3:value3","filter4:value4"]]`,
		},
		{
			opts: []interface{}{
				FacetFiltersAnd(
					FacetFilters("filter1", "value1"),
					FacetFiltersOr(
						FacetFilters("filter2", "value2"),
						FacetFilters("filter3", "value3"),
					),
				),
			},
			expectedJSON: `[["filter1:value1"],["filter2:value2","filter3:value3"]]`,
		},
	} {
		res := ExtractFacetFilters(c.opts...)
		data, err := json.Marshal(&res)
		require.NoError(t, err)
		require.Equal(t, c.expectedJSON, string(data))
	}
}
