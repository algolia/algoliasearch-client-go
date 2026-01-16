package utils

import (
	"encoding/json"
	"fmt"
)

// DeepMerge merges src into dst recursively.
//
// Merge rules:
//   - If both src and dst values for a key are maps, they are recursively merged
//   - Otherwise, the src value replaces the dst value (src wins)
//   - Arrays are replaced entirely, not merged
//
// This function modifies dst in place.
func DeepMerge(dst, src map[string]any) {
	for key, srcVal := range src {
		if srcMap, srcIsMap := srcVal.(map[string]any); srcIsMap {
			if dstVal, exists := dst[key]; exists {
				if dstMap, dstIsMap := dstVal.(map[string]any); dstIsMap {
					// Both are maps: recurse
					DeepMerge(dstMap, srcMap)

					continue
				}
			}
			// dst[key] doesn't exist or isn't a map: deep copy src map
			newMap := make(map[string]any)
			DeepMerge(newMap, srcMap)
			dst[key] = newMap

			continue
		}
		// src value is not a map: override dst
		dst[key] = srcVal
	}
}

// StructToMap converts any value to map[string]any via JSON round-trip.
// The returned map is a deep copy - modifications won't affect the original.
func StructToMap(v any) (map[string]any, error) {
	if v == nil {
		return make(map[string]any), nil
	}

	data, err := json.Marshal(v)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal to JSON: %w", err)
	}

	var result map[string]any

	err = json.Unmarshal(data, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal to map: %w", err)
	}

	if result == nil {
		result = make(map[string]any)
	}

	return result, nil
}

// MergeBodyParams merges custom body parameters into a typed request body.
//
// If customParams is empty, returns typedBody unchanged (fast path).
// Otherwise, converts typedBody to a map and deep-merges customParams into it.
//
// This is used by API clients to support WithBodyParam and WithBodyParams
// request options.
func MergeBodyParams(typedBody any, customParams map[string]any) (any, error) {
	// Fast path: no custom params
	if len(customParams) == 0 {
		return typedBody, nil
	}

	// Convert typed body to map
	bodyMap, err := StructToMap(typedBody)
	if err != nil {
		return nil, fmt.Errorf("failed to convert body to map: %w", err)
	}

	// Deep merge custom params (customParams wins on conflict)
	DeepMerge(bodyMap, customParams)

	return bodyMap, nil
}
