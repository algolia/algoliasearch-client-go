package opt

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/algolia/algoliasearch-client-go/algolia/opt"
	"github.com/stretchr/testify/require"
)

func TestValidUntil(t *testing.T) {
	for _, c := range []struct {
		opts     []interface{}
		expected *opt.ValidUntilOption
	}{
		{
			opts:     []interface{}{nil},
			expected: opt.ValidUntil(time.Time{}),
		},
		{
			opts:     []interface{}{opt.ValidUntil(time.Time{})},
			expected: opt.ValidUntil(time.Time{}),
		},
		{
			opts:     []interface{}{opt.ValidUntil(time.Date(2019, 4, 1, 10, 49, 0, 0, time.UTC))},
			expected: opt.ValidUntil(time.Date(2019, 4, 1, 10, 49, 0, 0, time.UTC)),
		},
	} {
		var (
			in  = ExtractValidUntil(c.opts...)
			out opt.ValidUntilOption
		)
		data, err := json.Marshal(&in)
		require.NoError(t, err)
		err = json.Unmarshal(data, &out)
		require.NoError(t, err)
		require.Equal(t, *c.expected, out)
	}
}
