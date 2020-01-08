package search

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRuleCondition_MarshalJSON(t *testing.T) {
	for _, c := range []struct {
		condition RuleCondition
		expected  string
	}{
		{
			RuleCondition{},
			`{}`,
		},
		{
			RuleCondition{
				Anchoring:    "",
				Pattern:      "",
				Context:      "",
				Alternatives: nil,
			},
			`{}`,
		},
		{
			RuleCondition{
				Anchoring:    Is,
				Pattern:      "",
				Context:      "",
				Alternatives: nil,
			},
			`{"anchoring": "is", "pattern": ""}`,
		},
		{
			RuleCondition{
				Anchoring:    "",
				Pattern:      "Pattern",
				Context:      "",
				Alternatives: nil,
			},
			`{"anchoring": "", "pattern": "Pattern"}`,
		},
		{
			RuleCondition{
				Anchoring:    Is,
				Pattern:      "Pattern",
				Context:      "Context",
				Alternatives: AlternativesEnabled(),
			},
			`{"anchoring": "is", "pattern": "Pattern", "context": "Context", "alternatives": true}`,
		},
		{
			RuleCondition{
				Anchoring:    Is,
				Pattern:      "Pattern",
				Context:      "Context",
				Alternatives: AlternativesDisabled(),
			},
			`{"anchoring": "is", "pattern": "Pattern", "context": "Context", "alternatives": false}`,
		},
	} {
		// Encode the RuleCondition to JSON
		data, err := json.Marshal(c.condition)
		require.NoError(t, err, "should marshal %#v without error", c.condition)

		// Compare the two JSON content
		var i1, i2 interface{}
		err = json.Unmarshal([]byte(c.expected), &i1)
		require.NoError(t, err)
		err = json.Unmarshal(data, &i2)
		require.NoError(t, err)
		require.Equal(t, i1, i2)

		// Decode the encoded JSON to a RuleCondition
		var condition RuleCondition
		err = json.Unmarshal(data, &condition)
		require.NoError(t, err, "should unmarshal %q without error", string(data))

		// Compare the two RuleConditions
		require.Equal(t, c.condition.Anchoring, condition.Anchoring)
		require.Equal(t, c.condition.Pattern, condition.Pattern)
		require.Equal(t, c.condition.Context, condition.Context)
		if c.condition.Alternatives == nil {
			require.Nil(t, condition.Alternatives)
		} else {
			require.NotNil(t, condition.Alternatives)
			require.Equal(t, *c.condition.Alternatives, *condition.Alternatives)
		}

	}
}
