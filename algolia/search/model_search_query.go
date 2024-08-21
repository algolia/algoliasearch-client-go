// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
package search

import (
	"encoding/json"
	"fmt"

	"github.com/algolia/algoliasearch-client-go/v4/algolia/utils"
)

// SearchQuery - struct for SearchQuery.
type SearchQuery struct {
	SearchForFacets *SearchForFacets
	SearchForHits   *SearchForHits
}

// SearchForFacetsAsSearchQuery is a convenience function that returns SearchForFacets wrapped in SearchQuery.
func SearchForFacetsAsSearchQuery(v *SearchForFacets) *SearchQuery {
	return &SearchQuery{
		SearchForFacets: v,
	}
}

// SearchForHitsAsSearchQuery is a convenience function that returns SearchForHits wrapped in SearchQuery.
func SearchForHitsAsSearchQuery(v *SearchForHits) *SearchQuery {
	return &SearchQuery{
		SearchForHits: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct.
func (dst *SearchQuery) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup if possible, if not we will try every possibility
	var jsonDict map[string]any
	_ = newStrictDecoder(data).Decode(&jsonDict)
	if utils.HasKey(jsonDict, "facet") && utils.HasKey(jsonDict, "type") {
		// try to unmarshal data into SearchForFacets
		err = newStrictDecoder(data).Decode(&dst.SearchForFacets)
		if err == nil && validateStruct(dst.SearchForFacets) == nil {
			return nil // found the correct type
		} else {
			dst.SearchForFacets = nil
		}
	}
	// try to unmarshal data into SearchForHits
	err = newStrictDecoder(data).Decode(&dst.SearchForHits)
	if err == nil && validateStruct(dst.SearchForHits) == nil {
		return nil // found the correct type
	} else {
		dst.SearchForHits = nil
	}

	return fmt.Errorf("Data failed to match schemas in oneOf(SearchQuery)")
}

// Marshal data from the first non-nil pointers in the struct to JSON.
func (src SearchQuery) MarshalJSON() ([]byte, error) {
	if src.SearchForFacets != nil {
		serialized, err := json.Marshal(&src.SearchForFacets)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal one of SearchForFacets of SearchQuery: %w", err)
		}

		return serialized, nil
	}

	if src.SearchForHits != nil {
		serialized, err := json.Marshal(&src.SearchForHits)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal one of SearchForHits of SearchQuery: %w", err)
		}

		return serialized, nil
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance.
func (obj SearchQuery) GetActualInstance() any {
	if obj.SearchForFacets != nil {
		return *obj.SearchForFacets
	}

	if obj.SearchForHits != nil {
		return *obj.SearchForHits
	}

	// all schemas are nil
	return nil
}
