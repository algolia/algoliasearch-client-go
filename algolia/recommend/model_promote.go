// File generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation.
package recommend

import (
	"encoding/json"
	"fmt"
)

// Promote - struct for Promote.
type Promote struct {
	PromoteObjectID  *PromoteObjectID
	PromoteObjectIDs *PromoteObjectIDs
}

// PromoteObjectIDsAsPromote is a convenience function that returns PromoteObjectIDs wrapped in Promote.
func PromoteObjectIDsAsPromote(v *PromoteObjectIDs) *Promote {
	return &Promote{
		PromoteObjectIDs: v,
	}
}

// PromoteObjectIDAsPromote is a convenience function that returns PromoteObjectID wrapped in Promote.
func PromoteObjectIDAsPromote(v *PromoteObjectID) *Promote {
	return &Promote{
		PromoteObjectID: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct.
func (dst *Promote) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal data into PromoteObjectID
	err = newStrictDecoder(data).Decode(&dst.PromoteObjectID)
	if err == nil && validateStruct(dst.PromoteObjectID) == nil {
		jsonPromoteObjectID, _ := json.Marshal(dst.PromoteObjectID)
		if string(jsonPromoteObjectID) == "{}" { // empty struct
			dst.PromoteObjectID = nil
		} else {
			return nil
		}
	} else {
		dst.PromoteObjectID = nil
	}

	// try to unmarshal data into PromoteObjectIDs
	err = newStrictDecoder(data).Decode(&dst.PromoteObjectIDs)
	if err == nil && validateStruct(dst.PromoteObjectIDs) == nil {
		jsonPromoteObjectIDs, _ := json.Marshal(dst.PromoteObjectIDs)
		if string(jsonPromoteObjectIDs) == "{}" { // empty struct
			dst.PromoteObjectIDs = nil
		} else {
			return nil
		}
	} else {
		dst.PromoteObjectIDs = nil
	}

	return fmt.Errorf("Data failed to match schemas in oneOf(Promote)")
}

// Marshal data from the first non-nil pointers in the struct to JSON.
func (src Promote) MarshalJSON() ([]byte, error) {
	if src.PromoteObjectID != nil {
		serialized, err := json.Marshal(&src.PromoteObjectID)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal one of PromoteObjectID of Promote: %w", err)
		}

		return serialized, nil
	}

	if src.PromoteObjectIDs != nil {
		serialized, err := json.Marshal(&src.PromoteObjectIDs)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal one of PromoteObjectIDs of Promote: %w", err)
		}

		return serialized, nil
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance.
func (obj *Promote) GetActualInstance() any {
	if obj == nil {
		return nil
	}
	if obj.PromoteObjectID != nil {
		return obj.PromoteObjectID
	}

	if obj.PromoteObjectIDs != nil {
		return obj.PromoteObjectIDs
	}

	// all schemas are nil
	return nil
}

type NullablePromote struct {
	value *Promote
	isSet bool
}

func (v NullablePromote) Get() *Promote {
	return v.value
}

func (v *NullablePromote) Set(val *Promote) {
	v.value = val
	v.isSet = true
}

func (v NullablePromote) IsSet() bool {
	return v.isSet
}

func (v *NullablePromote) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePromote(val *Promote) *NullablePromote {
	return &NullablePromote{value: val, isSet: true}
}

func (v NullablePromote) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value) //nolint:wrapcheck
}

func (v *NullablePromote) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value) //nolint:wrapcheck
}
