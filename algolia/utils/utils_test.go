package utils_test

import (
	"testing"
	"time"

	"github.com/algolia/algoliasearch-client-go/v4/algolia/utils"
)

type oneOfWrapper struct {
	actual any
}

func (o oneOfWrapper) GetActualInstance() any {
	return o.actual
}

func TestParameterToString(t *testing.T) {
	t.Parallel()

	fixedTime := time.Date(2024, 6, 15, 10, 30, 0, 0, time.UTC)
	strPtr := utils.ToPtr("pointed")

	var nilPtr *string

	tests := []struct {
		name     string
		input    any
		expected string
	}{
		{
			name:     "nil",
			input:    nil,
			expected: "",
		},
		{
			name:     "nil pointer",
			input:    nilPtr,
			expected: "",
		},
		{
			name:     "pointer to string",
			input:    strPtr,
			expected: "pointed",
		},
		{
			name:     "string slice",
			input:    []string{"a", "b", "c"},
			expected: "a,b,c",
		},
		{
			name:     "map",
			input:    map[string]any{"key": "value", "num": 42},
			expected: `{"key":"value","num":42}`,
		},
		{
			name:     "time.Time",
			input:    fixedTime,
			expected: "2024-06-15T10:30:00Z",
		},
		{
			name:     "oneOf wrapper struct",
			input:    oneOfWrapper{actual: "unwrapped"},
			expected: "unwrapped",
		},
		{
			name:     "int",
			input:    42,
			expected: "42",
		},
		{
			name:     "bool",
			input:    true,
			expected: "true",
		},
		{
			name:     "float",
			input:    3.14,
			expected: "3.14",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result := utils.ParameterToString(tt.input)
			if result != tt.expected {
				t.Errorf("ParameterToString() = %q, want %q", result, tt.expected)
			}
		})
	}
}

func TestQueryParameterToString(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    any
		expected string
	}{
		{
			name:     "string with spaces",
			input:    "hello world",
			expected: "hello%20world",
		},
		{
			name:     "special characters",
			input:    "key=value&other=123",
			expected: "key%3Dvalue%26other%3D123",
		},
		{
			name:     "map with special chars",
			input:    map[string]string{"query": "shoes & bags"},
			expected: "%7B%22query%22%3A%22shoes%20%26%20bags%22%7D",
		},
		{
			name:     "simple int",
			input:    123,
			expected: "123",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result := utils.QueryParameterToString(tt.input)
			if result != tt.expected {
				t.Errorf("QueryParameterToString() = %q, want %q", result, tt.expected)
			}
		})
	}
}
