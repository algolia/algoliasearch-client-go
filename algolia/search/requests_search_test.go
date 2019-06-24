package search

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/algolia/algoliasearch-client-go/algolia/opt"
)

func TestNewSearchParams_ExtraOptionsOverride(t *testing.T) {
	params := newSearchParams("", opt.Analytics(false), opt.ExtraOptions(map[string]interface{}{"analytics": true}))
	data, err := json.Marshal(params)
	require.NoError(t, err)

	var m map[string]interface{}
	err = json.Unmarshal(data, &m)
	require.NoError(t, err)

	require.Len(t, m, 2)

	itf, ok := m["analytics"]
	require.True(t, ok)

	v, ok := itf.(bool)
	require.True(t, ok)
	require.True(t, v)
}
