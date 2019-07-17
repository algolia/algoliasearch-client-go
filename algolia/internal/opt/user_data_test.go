// Code generated by go generate. DO NOT EDIT.

package opt

import (
	"encoding/json"
	"testing"

	"github.com/algolia/algoliasearch-client-go/algolia/opt"
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
	} {
		var (
			in  = ExtractUserData(c.opts...)
			out opt.UserDataOption
		)
		data, err := json.Marshal(&in)
		require.NoError(t, err)
		err = json.Unmarshal(data, &out)
		require.NoError(t, err)
		require.Equal(t, *c.expected, out)
	}
}
