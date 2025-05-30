// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
package search

import (
	"encoding/json"
	"fmt"

	"github.com/algolia/algoliasearch-client-go/v4/algolia/utils"
)

// BrowseParams - struct for BrowseParams.
type BrowseParams struct {
	BrowseParamsObject *BrowseParamsObject
	SearchParamsString *SearchParamsString
}

// SearchParamsStringAsBrowseParams is a convenience function that returns SearchParamsString wrapped in BrowseParams.
func SearchParamsStringAsBrowseParams(v *SearchParamsString) *BrowseParams {
	return &BrowseParams{
		SearchParamsString: v,
	}
}

// BrowseParamsObjectAsBrowseParams is a convenience function that returns BrowseParamsObject wrapped in BrowseParams.
func BrowseParamsObjectAsBrowseParams(v *BrowseParamsObject) *BrowseParams {
	return &BrowseParams{
		BrowseParamsObject: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct.
func (dst *BrowseParams) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup if possible, if not we will try every possibility
	var jsonDict map[string]any
	_ = json.Unmarshal(data, &jsonDict)
	if utils.HasKey(jsonDict, "params") {
		// try to unmarshal data into SearchParamsString
		err = json.Unmarshal(data, &dst.SearchParamsString)
		if err == nil {
			return nil // found the correct type
		} else {
			dst.SearchParamsString = nil
		}
	}
	// try to unmarshal data into BrowseParamsObject
	err = json.Unmarshal(data, &dst.BrowseParamsObject)
	if err == nil {
		return nil // found the correct type
	} else {
		dst.BrowseParamsObject = nil
	}

	return fmt.Errorf("Data failed to match schemas in oneOf(BrowseParams)")
}

// Marshal data from the first non-nil pointers in the struct to JSON.
func (src BrowseParams) MarshalJSON() ([]byte, error) {
	if src.BrowseParamsObject != nil {
		serialized, err := json.Marshal(&src.BrowseParamsObject)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal one of BrowseParamsObject of BrowseParams: %w", err)
		}

		return serialized, nil
	}

	if src.SearchParamsString != nil {
		serialized, err := json.Marshal(&src.SearchParamsString)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal one of SearchParamsString of BrowseParams: %w", err)
		}

		return serialized, nil
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance.
func (obj BrowseParams) GetActualInstance() any {
	if obj.BrowseParamsObject != nil {
		return *obj.BrowseParamsObject
	}

	if obj.SearchParamsString != nil {
		return *obj.SearchParamsString
	}

	// all schemas are nil
	return nil
}
