package opt

import (
	"testing"
	"time"

	"github.com/algolia/algoliasearch-client-go/v3/algolia/opt"
	"github.com/stretchr/testify/require"
)

func TestWaitConfiguration(t *testing.T) {
	for _, c := range []struct {
		name     string
		opts     []interface{}
		expected *opt.WaitConfigurationOption
	}{
		{
			name:     "no conf",
			opts:     []interface{}{},
			expected: opt.DefaultWaitConfiguration(),
		},
		{
			name:     "custom conf",
			opts:     []interface{}{&opt.WaitConfigurationOption{DelayGrowth: func(*time.Duration) time.Duration { return 42 * time.Second }}},
			expected: &opt.WaitConfigurationOption{DelayGrowth: func(*time.Duration) time.Duration { return 42 * time.Second }},
		},
	} {
		t.Run(c.name, func(t *testing.T) {
			in := ExtractWaitConfiguration(c.opts...)
			// Lambda can't be compared, the first value is used as an identifier
			require.Equal(t, c.expected.DelayGrowth(nil), in.DelayGrowth(nil))
		})
	}
}
