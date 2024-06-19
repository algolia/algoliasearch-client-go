// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
package search

import (
	"encoding/json"
	"fmt"
)

// TagFilters - Filter the search by values of the special `_tags` attribute.  **Prefer using the `filters` parameter, which supports all filter types and combinations with boolean operators.**  Different from regular facets, `_tags` can only be used for filtering (including or excluding records). You won't get a facet count. The same combination and escaping rules apply as for `facetFilters`.
type TagFilters struct {
	ArrayOfTagFilters *[]TagFilters
	String            *string
}

// []TagFiltersAsTagFilters is a convenience function that returns []TagFilters wrapped in TagFilters.
func ArrayOfTagFiltersAsTagFilters(v []TagFilters) *TagFilters {
	return &TagFilters{
		ArrayOfTagFilters: &v,
	}
}

// stringAsTagFilters is a convenience function that returns string wrapped in TagFilters.
func StringAsTagFilters(v string) *TagFilters {
	return &TagFilters{
		String: &v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct.
func (dst *TagFilters) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal data into ArrayOfTagFilters
	err = newStrictDecoder(data).Decode(&dst.ArrayOfTagFilters)
	if err == nil && validateStruct(dst.ArrayOfTagFilters) == nil {
		jsonArrayOfTagFilters, _ := json.Marshal(dst.ArrayOfTagFilters)
		if string(jsonArrayOfTagFilters) == "{}" { // empty struct
			dst.ArrayOfTagFilters = nil
		} else {
			return nil
		}
	} else {
		dst.ArrayOfTagFilters = nil
	}

	// try to unmarshal data into String
	err = newStrictDecoder(data).Decode(&dst.String)
	if err == nil && validateStruct(dst.String) == nil {
		jsonString, _ := json.Marshal(dst.String)
		if string(jsonString) == "{}" { // empty struct
			dst.String = nil
		} else {
			return nil
		}
	} else {
		dst.String = nil
	}

	return fmt.Errorf("Data failed to match schemas in oneOf(TagFilters)")
}

// Marshal data from the first non-nil pointers in the struct to JSON.
func (src TagFilters) MarshalJSON() ([]byte, error) {
	if src.ArrayOfTagFilters != nil {
		serialized, err := json.Marshal(&src.ArrayOfTagFilters)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal one of ArrayOfTagFilters of TagFilters: %w", err)
		}

		return serialized, nil
	}

	if src.String != nil {
		serialized, err := json.Marshal(&src.String)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal one of String of TagFilters: %w", err)
		}

		return serialized, nil
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance.
func (obj TagFilters) GetActualInstance() any {
	if obj.ArrayOfTagFilters != nil {
		return *obj.ArrayOfTagFilters
	}

	if obj.String != nil {
		return *obj.String
	}

	// all schemas are nil
	return nil
}
