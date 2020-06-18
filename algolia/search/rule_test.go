package search

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRule_MarshalJSON(t *testing.T) {
	for _, c := range []struct {
		rule     Rule
		expected string
	}{
		{
			Rule{},
			`{"consequence":{}}`,
		},
		{
			Rule{
				Condition: RuleCondition{},
			},
			`{"consequence":{}}`,
		},
		{
			Rule{
				Conditions: []RuleCondition{},
			},
			`{"consequence":{}}`,
		},
		{
			Rule{
				Conditions: []RuleCondition{
					{
						Anchoring:    Is,
						Pattern:      "Pattern",
						Context:      "Context",
						Alternatives: AlternativesEnabled(),
					},
				},
			},
			`{"conditions":[{"anchoring":"is","pattern":"Pattern","context":"Context","alternatives":true}],"consequence":{}}`,
		},
		{
			Rule{
				Condition: RuleCondition{
					Anchoring:    Is,
					Pattern:      "Pattern",
					Context:      "Context",
					Alternatives: AlternativesEnabled(),
				},
			},
			`{"condition":{"anchoring":"is","pattern":"Pattern","context":"Context","alternatives":true},"consequence":{}}`,
		},
	} {
		// Encode the Rule to JSON
		data, err := json.Marshal(&c.rule)
		require.NoError(t, err, "should not get an error while encoding %#v", c.rule)

		// Compare the two JSON content
		var i1, i2 interface{}
		err = json.Unmarshal([]byte(c.expected), &i1)
		require.NoError(t, err)
		err = json.Unmarshal(data, &i2)
		require.NoError(t, err)
		require.Equal(t, i1, i2)

		// Decode the encoded JSON to a Rule
		var rule Rule
		err = json.Unmarshal(data, &rule)
		require.NoError(t, err, "should unmarshal %q without error", string(data))

		// Compare the two Rules
		require.True(t, true, c.rule.Equal(rule))
	}
}
