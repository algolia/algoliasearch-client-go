package search

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSynonymUnmarshalType(t *testing.T) {
	for _, c := range []struct {
		payload  string
	}{
		{`{"type": "synonym"}`},
		{`{"type": "oneWaySynonym"}`},
		{`{"type": "altCorrection1"}`},
		{`{"type": "altCorrection2"}`},
		{`{"type": "placeholder"}`},
		{`{"type": "onewaysynonym"}`},
		{`{"type": "altcorrection1"}`},
		{`{"type": "altcorrection2"}`},
	} {
		var s rawSynonym
		err := json.Unmarshal([]byte(c.payload), &s)
		require.NoError(t, err, "payload=%q", c.payload)
	}
}
