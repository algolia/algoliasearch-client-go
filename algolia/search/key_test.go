package search

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/algolia/algoliasearch-client-go/v3/algolia/opt"
)

func TestAreKeysEqual(t *testing.T) {
	for _, c := range []struct {
		k1       Key
		k2       Key
		areEqual bool
	}{
		{Key{}, Key{}, true},

		{Key{ACL: []string{"search"}}, Key{ACL: []string{"search"}}, true},
		{Key{ACL: []string{"search"}}, Key{ACL: []string{"browse"}}, false},
		{Key{ACL: []string{"search"}}, Key{ACL: []string{"search", "browse"}}, false},

		{Key{CreatedAt: time.Now()}, Key{CreatedAt: time.Now().Add(1 * time.Hour)}, true},

		{Key{Description: "a description"}, Key{Description: "a description"}, true},
		{Key{Description: "a description"}, Key{Description: "A DESCRIPTION"}, false},

		{Key{Indexes: []string{"index1"}}, Key{Indexes: []string{"index1"}}, true},
		{Key{Indexes: []string{"index1"}}, Key{Indexes: []string{"index2"}}, false},
		{Key{Indexes: []string{"index1"}}, Key{Indexes: []string{"index1", "index2"}}, false},

		{Key{MaxQueriesPerIPPerHour: 1}, Key{MaxQueriesPerIPPerHour: 1}, true},
		{Key{MaxQueriesPerIPPerHour: 1}, Key{MaxQueriesPerIPPerHour: 2}, false},

		{Key{MaxHitsPerQuery: 1}, Key{MaxHitsPerQuery: 1}, true},
		{Key{MaxHitsPerQuery: 1}, Key{MaxHitsPerQuery: 2}, false},

		{*(&Key{}).SetQueryParameters(opt.RestrictSources("192.168.1.0/24")), *(&Key{}).SetQueryParameters(opt.RestrictSources("192.168.1.0/24")), true},
		{*(&Key{}).SetQueryParameters(opt.RestrictSources("192.168.1.0/24")), *(&Key{}).SetQueryParameters(opt.RestrictSources("192.168.1.0/32")), false},

		{*(&Key{}).SetQueryParameters(opt.TypoToleranceStrict()), *(&Key{}).SetQueryParameters(opt.TypoToleranceStrict()), true},
		{*(&Key{}).SetQueryParameters(opt.TypoToleranceStrict()), *(&Key{}).SetQueryParameters(opt.TypoToleranceMin()), false},

		{Key{Referers: []string{"referer1"}}, Key{Referers: []string{"referer1"}}, true},
		{Key{Referers: []string{"referer1"}}, Key{Referers: []string{"referer2"}}, false},
		{Key{Referers: []string{"referer1"}}, Key{Referers: []string{"referer1", "referer2"}}, false},

		{Key{Validity: time.Duration(1)}, Key{Validity: time.Duration(2)}, true},

		{Key{Value: "XYZ"}, Key{Value: "XYZ"}, true},
		{Key{Value: "XYZ"}, Key{Value: "ABC"}, false},
	} {
		c1, c2 := c.k1.CreatedAt, c.k2.CreatedAt
		v1, v2 := c.k1.Validity, c.k2.Validity

		require.Equal(t, c.areEqual, c.k1.Equal(c.k2), "keys:\n%#v\n%#v", c.k1, c.k2)
		require.Equal(t, c.areEqual, c.k2.Equal(c.k1), "keys:\n%#v\n%#v", c.k2, c.k1)

		require.Equal(t, c1, c.k1.CreatedAt)
		require.Equal(t, c2, c.k2.CreatedAt)
		require.Equal(t, v1, c.k1.Validity)
		require.Equal(t, v2, c.k2.Validity)
	}
}

func TestKeyDeserializer(t *testing.T) {
	for _, c := range []struct {
		payload  string
		expected Key
	}{
		{
			`{}`,
			Key{CreatedAt: time.Unix(0, 0)},
		},
		{
			`{"queryParameters":"facetFilters=[[%22privatePlatformCodes:-ha%22], [%22biddingStatus:-Closed%22, %22saleStatus:None%22]]"}`,
			Key{
				CreatedAt: time.Unix(0, 0),
				QueryParameters: KeyQueryParams{
					QueryParams: QueryParams{
						FacetFilters: opt.FacetFilterAnd(
							opt.FacetFilter("privatePlatformCodes:-ha"),
							opt.FacetFilterOr(
								opt.FacetFilter("biddingStatus:-Closed"),
								opt.FacetFilter("saleStatus:None"),
							),
						),
					},
				},
			},
		},
	} {
		var got Key
		err := json.Unmarshal([]byte(c.payload), &got)
		require.NoError(t, err, "cannot unmarshal key payload %s", c.payload)
		require.Equal(t, c.expected, got, "keys are not equal for payload %s", c.payload)
	}
}
