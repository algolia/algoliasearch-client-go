package search

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/algolia/algoliasearch-client-go/v3/algolia/opt"
)

func TestSearchSynonymsParams_Type(t *testing.T) {
	params := newSearchSynonymsParams(
		"q",
		opt.Type(
			string(OneWaySynonymType),
			string(AltCorrection1Type),
		),
	)
	data, err := json.Marshal(params)
	require.NoError(t, err)

	var m map[string]string
	err = json.Unmarshal(data, &m)
	require.NoError(t, err)

	require.Equal(t, "q", m["query"])
	require.Equal(t, "oneWaySynonym,altCorrection1", m["type"])
}
