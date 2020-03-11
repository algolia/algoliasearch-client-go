package opt

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestType(t *testing.T) {
	for _, expected := range []*TypeOption{
		new(TypeOption),
		Type([]string{}...),
		Type("oneWaySynonym", "altCorrection1"),
	} {
		data, err := json.Marshal(expected)
		require.NoError(t, err)
		expectedPayload := `"` + strings.Join(expected.Get(), ",") + `"`
		require.Equal(t, expectedPayload, string(data))

		var got TypeOption
		err = json.Unmarshal(data, &got)
		require.NoError(t, err)

		require.ElementsMatch(t, expected.Get(), got.Get())
	}
}
