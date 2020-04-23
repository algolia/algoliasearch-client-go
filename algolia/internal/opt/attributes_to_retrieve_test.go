// Code generated by go generate. DO NOT EDIT.

package opt

import (
	"encoding/json"
	"testing"

	"github.com/algolia/algoliasearch-client-go/v3/algolia/opt"
	"github.com/stretchr/testify/require"
)

func TestAttributesToRetrieve(t *testing.T) {
	for _, c := range []struct {
		opts     []interface{}
		expected *opt.AttributesToRetrieveOption
	}{
		{
			opts:     []interface{}{nil},
			expected: opt.AttributesToRetrieve([]string{"*"}...),
		},
		{
			opts:     []interface{}{opt.AttributesToRetrieve("value1")},
			expected: opt.AttributesToRetrieve("value1"),
		},
		{
			opts:     []interface{}{opt.AttributesToRetrieve("value1", "value2", "value3")},
			expected: opt.AttributesToRetrieve("value1", "value2", "value3"),
		},
	} {
		var (
			in  = ExtractAttributesToRetrieve(c.opts...)
			out opt.AttributesToRetrieveOption
		)
		data, err := json.Marshal(&in)
		require.NoError(t, err)
		err = json.Unmarshal(data, &out)
		require.NoError(t, err)
		require.ElementsMatch(t, c.expected.Get(), out.Get())
	}
}

func TestAttributesToRetrieve_CommaSeparatedString(t *testing.T) {
	for _, c := range []struct {
		payload  string
		expected *opt.AttributesToRetrieveOption
	}{
		{
			payload:  `""`,
			expected: opt.AttributesToRetrieve([]string{}...),
		},
		{
			payload:  `"value1"`,
			expected: opt.AttributesToRetrieve("value1"),
		},
		{
			payload:  `"value1,value2,value3"`,
			expected: opt.AttributesToRetrieve("value1", "value2", "value3"),
		},
	} {
		var got opt.AttributesToRetrieveOption
		err := json.Unmarshal([]byte(c.payload), &got)
		require.NoError(t, err)
		require.ElementsMatch(t, c.expected.Get(), got.Get())
	}
}
