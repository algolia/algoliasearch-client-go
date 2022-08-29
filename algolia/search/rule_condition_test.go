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
				Filters:      "",
			},
			`{}`,
		},
		{
			RuleCondition{
				Anchoring:    Is,
				Pattern:      "",
				Context:      "",
				Alternatives: nil,
				Filters:      "",
			},
			`{"anchoring": "is", "pattern": ""}`,
		},
		{
			RuleCondition{
				Anchoring:    "",
				Pattern:      "Pattern",
				Context:      "",
				Alternatives: nil,
				Filters:      "",
			},
			`{}`,
		},
		{
			RuleCondition{
				Anchoring:    Is,
				Pattern:      "Pattern",
				Context:      "Context",
				Alternatives: AlternativesEnabled(),
				Filters:      "",
			},
			`{"anchoring": "is", "pattern": "Pattern", "context": "Context", "alternatives": true}`,
		},
		{
			RuleCondition{
				Anchoring:    Is,
				Pattern:      "Pattern",
				Context:      "Context",
				Alternatives: AlternativesDisabled(),
				Filters:      "",
			},
			`{"anchoring": "is", "pattern": "Pattern", "context": "Context", "alternatives": false}`,
		},
		{
			RuleCondition{
				Anchoring:    "",
				Pattern:      "",
				Context:      "",
				Alternatives: nil,
				Filters:      "type:house AND (amenity:fireplace OR amenity:place)",
			},
			`{"filters": "type:house AND (amenity:fireplace OR amenity:place)"}`,
		},
		{
			RuleCondition{
				Anchoring:    Is,
				Pattern:      "",
				Context:      "Context",
				Alternatives: nil,
				Filters:      "type:house AND (amenity:fireplace OR amenity:place)",
			},
			`{"anchoring": "is", "pattern": "", "context": "Context", "filters": "type:house AND (amenity:fireplace OR amenity:place)"}`,
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
		if c.condition.Anchoring == "" {
			require.Equal(t, condition.Pattern, "")
		} else {
			require.Equal(t, c.condition.Pattern, condition.Pattern)
		}
		require.Equal(t, c.condition.Context, condition.Context)
		if c.condition.Alternatives == nil {
			require.Nil(t, condition.Alternatives)
		} else {
			require.NotNil(t, condition.Alternatives)
			require.Equal(t, *c.condition.Alternatives, *condition.Alternatives)
		}

	}
}
