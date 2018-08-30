package backward_compatibility

import (
	"encoding/json"
	"testing"

	"github.com/algolia/algoliasearch-client-go/algoliasearch"

	"github.com/stretchr/testify/require"
)

func TestOldSettings(t *testing.T) {
	t.SkipNow()
	t.Parallel()

	settingsMap := map[string]interface{}{
		"attributesToIndex":        []string{"attr1", "attr2"},
		"numericAttributesToIndex": []string{"attr1", "attr2"},
		"slaves":                   []string{"index1", "index2"},
	}

	data, err := json.Marshal(settingsMap)
	require.NoError(t, err)

	var settings algoliasearch.Settings
	err = json.Unmarshal(data, &settings)
	require.NoError(t, err)

	require.ElementsMatch(t, settingsMap["attributesToIndex"], settings.SearchableAttributes)
	require.ElementsMatch(t, settingsMap["numericAttributesToIndex"], settings.NumericAttributesForFiltering)
	require.ElementsMatch(t, settingsMap["slaves"], settings.Replicas)
}
