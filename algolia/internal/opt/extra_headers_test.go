package opt

import (
	"testing"

	"github.com/algolia/algoliasearch-client-go/algolia/opt"
	"github.com/stretchr/testify/require"
)

func TestExtractExtraHeaders(t *testing.T) {
	for _, c := range []struct {
		opts     []interface{}
		expected map[string]string
	}{
		{
			opts:     []interface{}{nil},
			expected: map[string]string{},
		},
		{
			opts:     []interface{}{opt.ExtraHeaders(map[string]string{})},
			expected: map[string]string{},
		},
		{
			opts:     []interface{}{opt.ExtraHeaders(map[string]string{"key": "value"})},
			expected: map[string]string{"key": "value"},
		},
		{
			opts: []interface{}{
				opt.ExtraHeaders(map[string]string{"key1": "value1", "key2": "value2"}),
				opt.ExtraHeaders(map[string]string{"key3": "value3", "key4": "value4"}),
			},
			expected: map[string]string{
				"key1": "value1",
				"key2": "value2",
				"key3": "value3",
				"key4": "value4",
			},
		},
		{
			opts: []interface{}{
				opt.ExtraHeaders(map[string]string{"key1": "value1", "key2": "value2"}),
				opt.ExtraHeaders(map[string]string{"key3": "value3", "key4": "value4"}),
			},
			expected: map[string]string{
				"key1": "value1",
				"key2": "value2",
				"key3": "value3",
				"key4": "value4",
			},
		},
		{
			opts: []interface{}{
				opt.ExtraHeaders(map[string]string{"key1": "value1", "key2": "value2"}),
				opt.ExtraHeaders(map[string]string{"key1": "value11", "key3": "value3"}),
			},
			expected: map[string]string{
				"key1": "value1,value11",
				"key2": "value2",
				"key3": "value3",
			},
		},
	} {
		out := ExtractExtraHeaders(c.opts...)
		require.Equal(t, c.expected, out.Get())
	}
}
