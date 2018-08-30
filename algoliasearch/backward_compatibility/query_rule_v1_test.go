package backward_compatibility

import (
	"encoding/json"
	"testing"

	"github.com/algolia/algoliasearch-client-go/algoliasearch"

	"github.com/stretchr/testify/require"
)

func TestQueryRulesV1(t *testing.T) {
	t.SkipNow()
	t.Parallel()

	queryRuleMap := map[string]interface{}{
		"objectID": "query_edits",
		"condition": map[string]interface{}{
			"anchoring": "is",
			"pattern":   "mobile phone",
		},
		"consequence": map[string]interface{}{
			"params": map[string]interface{}{
				"query": map[string]interface{}{
					"remove": []string{"mobile", "phone"},
				},
			},
		},
	}

	data, err := json.Marshal(&queryRuleMap)
	require.NoError(t, err)

	var rule algoliasearch.Rule
	err = json.Unmarshal(data, &rule)
	require.NoError(t, err)

	require.Equal(t, "one", rule.ObjectID)
	require.Equal(t,
		algoliasearch.RuleCondition{Anchoring: algoliasearch.Contains, Pattern: "mobile phone"},
		rule.Condition,
	)
	require.Equal(t,
		algoliasearch.RuleConsequence{
			Params: map[string]interface{}{
				"query": map[string]interface{}{
					"edits": []algoliasearch.Edit{
						algoliasearch.DeleteEdit("mobile"),
						algoliasearch.DeleteEdit("phone"),
					},
				},
			},
		},
		rule.Consequence,
	)
}
