package opt

import (
	"encoding/json"
	"testing"

	"github.com/algolia/algoliasearch-client-go/v3/algolia/opt"
	"github.com/stretchr/testify/require"
)

func TestExtractUserData(t *testing.T) {
	for _, c := range []struct {
		opts     []interface{}
		expected *opt.UserDataOption
	}{
		{
			opts:     []interface{}{nil},
			expected: opt.UserData(map[string]interface{}{}),
		},
		{
			opts:     []interface{}{opt.UserData(map[string]interface{}{})},
			expected: opt.UserData(map[string]interface{}{}),
		},
		{
			opts:     []interface{}{opt.UserData(map[string]interface{}{"k1": 1, "k2": []string{"2", "3"}})},
			expected: opt.UserData(map[string]interface{}{"k1": 1.0, "k2": []interface{}{"2", "3"}}),
		},
		{
			opts:     []interface{}{opt.UserData(42)},
			expected: opt.UserData(42),
		},
		{
			opts:     []interface{}{opt.UserData("random string")},
			expected: opt.UserData("random string"),
		},
	} {
		var (
			in  = ExtractUserData(c.opts...)
			out opt.UserDataOption
		)
		data, err := json.Marshal(&in)
		require.NoError(t, err)

		err = json.Unmarshal(data, &out)
		require.NoError(t, err)

		require.True(t, c.expected.Equal(&out), "expected %#v but got %#v", *c.expected, out)
		require.True(t, out.Equal(c.expected), "expected %#v but got %#v", *c.expected, out)
	}
}
