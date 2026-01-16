package utils_test

import (
	"reflect"
	"testing"

	"github.com/algolia/algoliasearch-client-go/v4/algolia/utils"
)

func TestDeepMerge(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		dst      map[string]any
		src      map[string]any
		expected map[string]any
	}{
		{
			name:     "simple merge - no overlap",
			dst:      map[string]any{"a": 1},
			src:      map[string]any{"b": 2},
			expected: map[string]any{"a": 1, "b": 2},
		},
		{
			name:     "override value",
			dst:      map[string]any{"a": 1},
			src:      map[string]any{"a": 2},
			expected: map[string]any{"a": 2},
		},
		{
			name:     "nested merge",
			dst:      map[string]any{"nested": map[string]any{"a": 1, "b": 2}},
			src:      map[string]any{"nested": map[string]any{"b": 3, "c": 4}},
			expected: map[string]any{"nested": map[string]any{"a": 1, "b": 3, "c": 4}},
		},
		{
			name:     "array replacement - not merged",
			dst:      map[string]any{"arr": []int{1, 2}},
			src:      map[string]any{"arr": []int{3, 4, 5}},
			expected: map[string]any{"arr": []int{3, 4, 5}},
		},
		{
			name:     "type override - int to string",
			dst:      map[string]any{"val": 123},
			src:      map[string]any{"val": "string"},
			expected: map[string]any{"val": "string"},
		},
		{
			name:     "nil src value overrides",
			dst:      map[string]any{"a": 1},
			src:      map[string]any{"a": nil},
			expected: map[string]any{"a": nil},
		},
		{
			name:     "empty src - no change",
			dst:      map[string]any{"a": 1},
			src:      map[string]any{},
			expected: map[string]any{"a": 1},
		},
		{
			name:     "empty dst - copy src",
			dst:      map[string]any{},
			src:      map[string]any{"a": 1, "b": 2},
			expected: map[string]any{"a": 1, "b": 2},
		},
		{
			name: "deeply nested - 3 levels",
			dst: map[string]any{
				"l1": map[string]any{
					"l2": map[string]any{
						"l3":       1,
						"existing": "keep",
					},
				},
			},
			src: map[string]any{
				"l1": map[string]any{
					"l2": map[string]any{
						"l3":  2,
						"new": 3,
					},
				},
			},
			expected: map[string]any{
				"l1": map[string]any{
					"l2": map[string]any{
						"l3":       2,
						"existing": "keep",
						"new":      3,
					},
				},
			},
		},
		{
			name:     "src map overwrites dst non-map",
			dst:      map[string]any{"key": "string"},
			src:      map[string]any{"key": map[string]any{"nested": true}},
			expected: map[string]any{"key": map[string]any{"nested": true}},
		},
		{
			name:     "src non-map overwrites dst map",
			dst:      map[string]any{"key": map[string]any{"nested": true}},
			src:      map[string]any{"key": "string"},
			expected: map[string]any{"key": "string"},
		},
		{
			name: "mixed types in nested",
			dst: map[string]any{
				"config": map[string]any{
					"enabled": true,
					"count":   10,
					"tags":    []string{"a", "b"},
				},
			},
			src: map[string]any{
				"config": map[string]any{
					"count":   20,
					"newFlag": true,
				},
			},
			expected: map[string]any{
				"config": map[string]any{
					"enabled": true,
					"count":   20,
					"tags":    []string{"a", "b"},
					"newFlag": true,
				},
			},
		},
		{
			name:     "boolean values",
			dst:      map[string]any{"flag": true},
			src:      map[string]any{"flag": false},
			expected: map[string]any{"flag": false},
		},
		{
			name:     "float values",
			dst:      map[string]any{"score": 1.5},
			src:      map[string]any{"score": 2.5},
			expected: map[string]any{"score": 2.5},
		},
	}

	for _, tt := range tests {
		tt := tt // capture range variable
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			// Make copies to avoid mutation across parallel tests
			dst := make(map[string]any)
			for k, v := range tt.dst {
				dst[k] = v
			}

			src := make(map[string]any)
			for k, v := range tt.src {
				src[k] = v
			}

			utils.DeepMerge(dst, src)

			if !reflect.DeepEqual(dst, tt.expected) {
				t.Errorf("DeepMerge() = %v, want %v", dst, tt.expected)
			}
		})
	}
}

func TestStructToMap(t *testing.T) {
	t.Parallel()

	type TestStruct struct {
		Name  string `json:"name"`
		Value int    `json:"value"`
	}

	type NestedStruct struct {
		Inner TestStruct `json:"inner"`
		Tags  []string   `json:"tags"`
	}

	tests := []struct {
		name     string
		input    any
		expected map[string]any
		wantErr  bool
	}{
		{
			name:     "nil input",
			input:    nil,
			expected: map[string]any{},
			wantErr:  false,
		},
		{
			name:     "already a map",
			input:    map[string]any{"a": 1, "b": "two"},
			expected: map[string]any{"a": 1, "b": "two"},
			wantErr:  false,
		},
		{
			name:     "simple struct",
			input:    TestStruct{Name: "test", Value: 42},
			expected: map[string]any{"name": "test", "value": float64(42)}, // JSON numbers are float64
			wantErr:  false,
		},
		{
			name: "nested struct",
			input: NestedStruct{
				Inner: TestStruct{Name: "inner", Value: 10},
				Tags:  []string{"a", "b"},
			},
			expected: map[string]any{
				"inner": map[string]any{"name": "inner", "value": float64(10)},
				"tags":  []any{"a", "b"},
			},
			wantErr: false,
		},
		{
			name:     "pointer to struct",
			input:    &TestStruct{Name: "ptr", Value: 99},
			expected: map[string]any{"name": "ptr", "value": float64(99)},
			wantErr:  false,
		},
		{
			name:     "empty struct",
			input:    TestStruct{},
			expected: map[string]any{"name": "", "value": float64(0)},
			wantErr:  false,
		},
		{
			name:     "empty map passthrough",
			input:    map[string]any{},
			expected: map[string]any{},
			wantErr:  false,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result, err := utils.StructToMap(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("StructToMap() error = %v, wantErr %v", err, tt.wantErr)

				return
			}

			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("StructToMap() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestStructToMap_DoesNotModifyOriginal(t *testing.T) {
	t.Parallel()

	original := map[string]any{
		"a": 1,
		"b": 2,
		"nested": map[string]any{
			"x": 10,
			"y": 20,
		},
	}

	result, err := utils.StructToMap(original)
	if err != nil {
		t.Fatalf("StructToMap() unexpected error: %v", err)
	}

	result["a"] = 999
	result["c"] = 3

	nestedResult, ok := result["nested"].(map[string]any)
	if !ok {
		t.Fatalf("nested is not a map")
	}

	nestedResult["x"] = 999
	nestedResult["z"] = 30

	if original["a"] != 1 {
		t.Errorf("Original top-level was modified: a = %v, want 1", original["a"])
	}

	if _, exists := original["c"]; exists {
		t.Errorf("Original has unexpected key 'c'")
	}

	nestedOriginal, ok := original["nested"].(map[string]any)
	if !ok {
		t.Fatalf("original nested is not a map")
	}

	if nestedOriginal["x"] != 10 {
		t.Errorf("Original nested map was modified: x = %v, want 10", nestedOriginal["x"])
	}

	if _, exists := nestedOriginal["z"]; exists {
		t.Errorf("Original nested map has unexpected key 'z'")
	}
}

func TestMergeBodyParams(t *testing.T) {
	t.Parallel()

	type SearchParams struct {
		Query       string `json:"query"`
		HitsPerPage int    `json:"hitsPerPage"`
	}

	tests := []struct {
		name         string
		typedBody    any
		customParams map[string]any
		expected     any
		wantErr      bool
	}{
		{
			name:         "empty custom params - return typed body",
			typedBody:    SearchParams{Query: "shoes", HitsPerPage: 20},
			customParams: map[string]any{},
			expected:     SearchParams{Query: "shoes", HitsPerPage: 20},
			wantErr:      false,
		},
		{
			name:         "nil custom params - return typed body",
			typedBody:    SearchParams{Query: "shoes", HitsPerPage: 20},
			customParams: nil,
			expected:     SearchParams{Query: "shoes", HitsPerPage: 20},
			wantErr:      false,
		},
		{
			name:         "add new field",
			typedBody:    SearchParams{Query: "shoes", HitsPerPage: 20},
			customParams: map[string]any{"customField": "value"},
			expected: map[string]any{
				"query":       "shoes",
				"hitsPerPage": float64(20),
				"customField": "value",
			},
			wantErr: false,
		},
		{
			name:         "override existing field",
			typedBody:    SearchParams{Query: "shoes", HitsPerPage: 20},
			customParams: map[string]any{"hitsPerPage": 100},
			expected: map[string]any{
				"query":       "shoes",
				"hitsPerPage": 100,
			},
			wantErr: false,
		},
		{
			name:      "nil typed body with custom params",
			typedBody: nil,
			customParams: map[string]any{
				"customOnly": "value",
			},
			expected: map[string]any{
				"customOnly": "value",
			},
			wantErr: false,
		},
		{
			name:      "typed body is already a map",
			typedBody: map[string]any{"existing": "value"},
			customParams: map[string]any{
				"new": "field",
			},
			expected: map[string]any{
				"existing": "value",
				"new":      "field",
			},
			wantErr: false,
		},
		{
			name: "nested custom params",
			typedBody: map[string]any{
				"settings": map[string]any{
					"enabled": true,
					"count":   10,
				},
			},
			customParams: map[string]any{
				"settings": map[string]any{
					"count":     20,
					"newOption": "value",
				},
			},
			expected: map[string]any{
				"settings": map[string]any{
					"enabled":   true,
					"count":     20,
					"newOption": "value",
				},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result, err := utils.MergeBodyParams(tt.typedBody, tt.customParams)
			if (err != nil) != tt.wantErr {
				t.Errorf("MergeBodyParams() error = %v, wantErr %v", err, tt.wantErr)

				return
			}

			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("MergeBodyParams() = %v, want %v", result, tt.expected)
			}
		})
	}
}
