package search

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestShouldSendBatch(t *testing.T) {
	type Object struct{}

	for _, c := range []struct {
		name         string
		maxBatchSize int
		batch        []interface{}
		object       interface{}
		expected     bool
	}{
		{
			"zero values",
			0,
			[]interface{}{},
			Object{},
			false,
		},
		{
			"batch not big enough with remaining objects",
			1,
			[]interface{}{},
			Object{},
			false,
		},
		{
			"batch big enough with remaining objects",
			1,
			[]interface{}{Object{}},
			Object{},
			true,
		},
		{
			"batch not big enough without remaining objects",
			2,
			[]interface{}{Object{}},
			nil,
			true,
		},
		{
			"batch big enough without remaining objects",
			2,
			[]interface{}{Object{}, Object{}},
			nil,
			true,
		},
	} {
		require.Equal(
			t,
			c.expected,
			shouldSendBatch(c.maxBatchSize, c.batch, c.object),
			"%q test case failed", c.name,
		)
	}
}
