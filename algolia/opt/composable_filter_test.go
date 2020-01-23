package opt

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestComposableFilterOption_UnmarshalJSON(t *testing.T) {
	for _, c := range []struct {
		payload  string
		expected composableFilterOption
	}{
		{
			`[]`,
			composableFilterOption{},
		},
		{
			`[[]]`,
			composableFilterOption{},
		},
		{
			`"color:green"`,
			composableFilterOption{[][]string{
				{`color:green`},
			}},
		},
		{
			`" color:green "`,
			composableFilterOption{[][]string{
				{`color:green`},
			}},
		},
		{
			`"color:green,color:yellow"`,
			composableFilterOption{[][]string{
				{`color:green`, `color:yellow`},
			}},
		},
		{
			`" color:green , color:yellow "`,
			composableFilterOption{[][]string{
				{`color:green`, `color:yellow`},
			}},
		},
		{
			`["color:green","color:yellow"]`,
			composableFilterOption{[][]string{
				{`color:green`, `color:yellow`},
			}},
		},
		{
			`[" color:green "," color:yellow "]`,
			composableFilterOption{[][]string{
				{`color:green`, `color:yellow`},
			}},
		},
		{
			`[["color:green"],["color:yellow"]]`,
			composableFilterOption{[][]string{
				{`color:green`},
				{`color:yellow`},
			}},
		},
		{
			`[[" color:green "],[" color:yellow "]]`,
			composableFilterOption{[][]string{
				{`color:green`},
				{`color:yellow`},
			}},
		},
	} {
		var got composableFilterOption
		err := json.Unmarshal([]byte(c.payload), &got)
		require.NoError(t, err, "cannot unmarshal payload %q", c.payload)

		fGot := got.Get()
		fExpected := c.expected.Get()

		require.Equal(
			t,
			len(fGot),
			len(fExpected),
			"expected %v as deserialized filters instead of %v for payload %q",
			fExpected,
			fGot,
			c.payload,
		)
	}
}
