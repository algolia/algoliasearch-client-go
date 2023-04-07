package opt

import (
	"encoding/json"
	"testing"

	"github.com/algolia/algoliasearch-client-go/v3/algolia/opt"
	"github.com/stretchr/testify/require"
)

func TestReRankingApplyFilter(t *testing.T) {
	for _, c := range []struct {
		opts     []interface{}
		expected *opt.ReRankingApplyFilterOption
	}{
		{
			opts:     []interface{}{nil},
			expected: &opt.ReRankingApplyFilterOption{},
		},
		{
			opts: []interface{}{opt.ReRankingApplyFilter("filter1:value1")},
			expected: opt.ReRankingApplyFilterAnd(
				[]string{"filter1:value1"},
			),
		},
		{
			opts: []interface{}{opt.ReRankingApplyFilterOr(
				"filter1:value1",
				"filter2:value2",
			)},
			expected: opt.ReRankingApplyFilterAnd(
				[]string{"filter1:value1", "filter2:value2"},
			),
		},
		{
			opts: []interface{}{opt.ReRankingApplyFilterOr(
				"filter1:value1",
				opt.ReRankingApplyFilter("filter2:value2"),
			)},
			expected: opt.ReRankingApplyFilterAnd(
				[]string{"filter1:value1", "filter2:value2"},
			),
		},
		{
			opts: []interface{}{opt.ReRankingApplyFilterAnd(
				[]string{"filter1:value1"},
				[]string{"filter2:value2"},
			)},
			expected: opt.ReRankingApplyFilterAnd(
				[]string{"filter1:value1"},
				[]string{"filter2:value2"},
			),
		},
		{
			opts: []interface{}{opt.ReRankingApplyFilterAnd(
				"filter1:value1",
				[]string{"filter2:value2"},
			)},
			expected: opt.ReRankingApplyFilterAnd(
				[]string{"filter1:value1"},
				[]string{"filter2:value2"},
			),
		},
		{
			opts: []interface{}{opt.ReRankingApplyFilterAnd(
				"filter1:value1",
				[]string{"filter2:value2", "filter3:value3"},
			)},
			expected: opt.ReRankingApplyFilterAnd(
				[]string{"filter1:value1"},
				[]string{"filter2:value2", "filter3:value3"},
			),
		},
		{
			opts: []interface{}{opt.ReRankingApplyFilterAnd(
				"filter1:value1",
				opt.ReRankingApplyFilterOr("filter2:value2", "filter3:value3"),
			)},
			expected: opt.ReRankingApplyFilterAnd(
				[]string{"filter1:value1"},
				[]string{"filter2:value2", "filter3:value3"},
			),
		},
	} {
		var (
			in  = ExtractReRankingApplyFilter(c.opts...)
			out opt.ReRankingApplyFilterOption
		)
		data, err := json.Marshal(&in)
		require.NoError(t, err)
		err = json.Unmarshal(data, &out)
		require.NoError(t, err)
		require.Equal(t, *c.expected, out)
	}
}
