package opt

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestContext(t *testing.T) {
	{
		expected := context.Background()
		ctx := ExtractContext(expected)
		require.Equal(t, expected, ctx)
	}

	{
		expected := context.TODO()
		ctx := ExtractContext(expected)
		require.Equal(t, expected, ctx)
	}

	{
		expected, _ := context.WithTimeout(context.Background(), 42*time.Second)
		ctx := ExtractContext(expected)
		require.Equal(t, expected, ctx)
	}
}
