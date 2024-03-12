// File generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation.
package recommend

import (
	"encoding/json"
	"fmt"
)

// AutomaticFacetFilters - Filter to be applied to the search.  You can use this to respond to search queries that match a facet value. For example, if users search for \"comedy\", which matches a facet value of the \"genre\" facet, you can filter the results to show the top-ranked comedy movies.
type AutomaticFacetFilters struct {
	ArrayOfAutomaticFacetFilter *[]AutomaticFacetFilter
	ArrayOfString               *[]string
}

// []AutomaticFacetFilterAsAutomaticFacetFilters is a convenience function that returns []AutomaticFacetFilter wrapped in AutomaticFacetFilters.
func ArrayOfAutomaticFacetFilterAsAutomaticFacetFilters(v []AutomaticFacetFilter) *AutomaticFacetFilters {
	return &AutomaticFacetFilters{
		ArrayOfAutomaticFacetFilter: &v,
	}
}

// []stringAsAutomaticFacetFilters is a convenience function that returns []string wrapped in AutomaticFacetFilters.
func ArrayOfStringAsAutomaticFacetFilters(v []string) *AutomaticFacetFilters {
	return &AutomaticFacetFilters{
		ArrayOfString: &v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct.
func (dst *AutomaticFacetFilters) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal data into ArrayOfAutomaticFacetFilter
	err = newStrictDecoder(data).Decode(&dst.ArrayOfAutomaticFacetFilter)
	if err == nil && validateStruct(dst.ArrayOfAutomaticFacetFilter) == nil {
		jsonArrayOfAutomaticFacetFilter, _ := json.Marshal(dst.ArrayOfAutomaticFacetFilter)
		if string(jsonArrayOfAutomaticFacetFilter) == "{}" { // empty struct
			dst.ArrayOfAutomaticFacetFilter = nil
		} else {
			return nil
		}
	} else {
		dst.ArrayOfAutomaticFacetFilter = nil
	}

	// try to unmarshal data into ArrayOfString
	err = newStrictDecoder(data).Decode(&dst.ArrayOfString)
	if err == nil && validateStruct(dst.ArrayOfString) == nil {
		jsonArrayOfString, _ := json.Marshal(dst.ArrayOfString)
		if string(jsonArrayOfString) == "{}" { // empty struct
			dst.ArrayOfString = nil
		} else {
			return nil
		}
	} else {
		dst.ArrayOfString = nil
	}

	return fmt.Errorf("Data failed to match schemas in oneOf(AutomaticFacetFilters)")
}

// Marshal data from the first non-nil pointers in the struct to JSON.
func (src AutomaticFacetFilters) MarshalJSON() ([]byte, error) {
	if src.ArrayOfAutomaticFacetFilter != nil {
		serialized, err := json.Marshal(&src.ArrayOfAutomaticFacetFilter)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal one of ArrayOfAutomaticFacetFilter of AutomaticFacetFilters: %w", err)
		}

		return serialized, nil
	}

	if src.ArrayOfString != nil {
		serialized, err := json.Marshal(&src.ArrayOfString)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal one of ArrayOfString of AutomaticFacetFilters: %w", err)
		}

		return serialized, nil
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance.
func (obj AutomaticFacetFilters) GetActualInstance() any {
	if obj.ArrayOfAutomaticFacetFilter != nil {
		return *obj.ArrayOfAutomaticFacetFilter
	}

	if obj.ArrayOfString != nil {
		return *obj.ArrayOfString
	}

	// all schemas are nil
	return nil
}

type NullableAutomaticFacetFilters struct {
	value *AutomaticFacetFilters
	isSet bool
}

func (v NullableAutomaticFacetFilters) Get() *AutomaticFacetFilters {
	return v.value
}

func (v *NullableAutomaticFacetFilters) Set(val *AutomaticFacetFilters) {
	v.value = val
	v.isSet = true
}

func (v NullableAutomaticFacetFilters) IsSet() bool {
	return v.isSet
}

func (v *NullableAutomaticFacetFilters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAutomaticFacetFilters(val *AutomaticFacetFilters) *NullableAutomaticFacetFilters {
	return &NullableAutomaticFacetFilters{value: val, isSet: true}
}

func (v NullableAutomaticFacetFilters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value) //nolint:wrapcheck
}

func (v *NullableAutomaticFacetFilters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value) //nolint:wrapcheck
}
