// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
package search

import (
	"encoding/json"
	"fmt"
)

// InsideBoundingBox - struct for InsideBoundingBox.
type InsideBoundingBox struct {
	ArrayOfArrayOfFloat64 *[][]float64
	String                *string
}

// stringAsInsideBoundingBox is a convenience function that returns string wrapped in InsideBoundingBox.
func StringAsInsideBoundingBox(v string) *InsideBoundingBox {
	return &InsideBoundingBox{
		String: &v,
	}
}

// [][]float64AsInsideBoundingBox is a convenience function that returns [][]float64 wrapped in InsideBoundingBox.
func ArrayOfArrayOfFloat64AsInsideBoundingBox(v [][]float64) *InsideBoundingBox {
	return &InsideBoundingBox{
		ArrayOfArrayOfFloat64: &v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct.
func (dst *InsideBoundingBox) UnmarshalJSON(data []byte) error {
	var err error
	// this object is nullable so check if the payload is null or empty string
	if string(data) == "" || string(data) == "{}" {
		return nil
	}

	// try to unmarshal data into String
	err = json.Unmarshal(data, &dst.String)
	if err == nil {
		return nil // found the correct type
	} else {
		dst.String = nil
	}
	// try to unmarshal data into ArrayOfArrayOfFloat64
	err = json.Unmarshal(data, &dst.ArrayOfArrayOfFloat64)
	if err == nil {
		return nil // found the correct type
	} else {
		dst.ArrayOfArrayOfFloat64 = nil
	}

	return fmt.Errorf("Data failed to match schemas in oneOf(InsideBoundingBox)")
}

// Marshal data from the first non-nil pointers in the struct to JSON.
func (src InsideBoundingBox) MarshalJSON() ([]byte, error) {
	if src.ArrayOfArrayOfFloat64 != nil {
		serialized, err := json.Marshal(&src.ArrayOfArrayOfFloat64)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal one of ArrayOfArrayOfFloat64 of InsideBoundingBox: %w", err)
		}

		return serialized, nil
	}

	if src.String != nil {
		serialized, err := json.Marshal(&src.String)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal one of String of InsideBoundingBox: %w", err)
		}

		return serialized, nil
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance.
func (obj InsideBoundingBox) GetActualInstance() any {
	if obj.ArrayOfArrayOfFloat64 != nil {
		return *obj.ArrayOfArrayOfFloat64
	}

	if obj.String != nil {
		return *obj.String
	}

	// all schemas are nil
	return nil
}
