// File generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation.
package search

import (
	"encoding/json"
	"fmt"
)

// PromoteObjectID Record to promote.
type PromoteObjectID struct {
	// Unique identifier of the record to promote.
	ObjectID string `json:"objectID"`
	// The position to promote the records to. If you pass objectIDs, the records are placed at this position as a group. For example, if you pronmote four objectIDs to position 0, the records take the first four positions.
	Position int32 `json:"position"`
}

// NewPromoteObjectID instantiates a new PromoteObjectID object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewPromoteObjectID(objectID string, position int32) *PromoteObjectID {
	this := &PromoteObjectID{}
	this.ObjectID = objectID
	this.Position = position
	return this
}

// NewEmptyPromoteObjectID return a pointer to an empty PromoteObjectID object.
func NewEmptyPromoteObjectID() *PromoteObjectID {
	return &PromoteObjectID{}
}

// GetObjectID returns the ObjectID field value.
func (o *PromoteObjectID) GetObjectID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectID
}

// GetObjectIDOk returns a tuple with the ObjectID field value
// and a boolean to check if the value has been set.
func (o *PromoteObjectID) GetObjectIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectID, true
}

// SetObjectID sets field value.
func (o *PromoteObjectID) SetObjectID(v string) *PromoteObjectID {
	o.ObjectID = v
	return o
}

// GetPosition returns the Position field value.
func (o *PromoteObjectID) GetPosition() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Position
}

// GetPositionOk returns a tuple with the Position field value
// and a boolean to check if the value has been set.
func (o *PromoteObjectID) GetPositionOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Position, true
}

// SetPosition sets field value.
func (o *PromoteObjectID) SetPosition(v int32) *PromoteObjectID {
	o.Position = v
	return o
}

func (o PromoteObjectID) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if true {
		toSerialize["objectID"] = o.ObjectID
	}
	if true {
		toSerialize["position"] = o.Position
	}
	serialized, err := json.Marshal(toSerialize)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal PromoteObjectID: %w", err)
	}

	return serialized, nil
}

func (o PromoteObjectID) String() string {
	out := ""
	out += fmt.Sprintf("  objectID=%v\n", o.ObjectID)
	out += fmt.Sprintf("  position=%v\n", o.Position)
	return fmt.Sprintf("PromoteObjectID {\n%s}", out)
}

type NullablePromoteObjectID struct {
	value *PromoteObjectID
	isSet bool
}

func (v NullablePromoteObjectID) Get() *PromoteObjectID {
	return v.value
}

func (v *NullablePromoteObjectID) Set(val *PromoteObjectID) {
	v.value = val
	v.isSet = true
}

func (v NullablePromoteObjectID) IsSet() bool {
	return v.isSet
}

func (v *NullablePromoteObjectID) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePromoteObjectID(val *PromoteObjectID) *NullablePromoteObjectID {
	return &NullablePromoteObjectID{value: val, isSet: true}
}

func (v NullablePromoteObjectID) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value) //nolint:wrapcheck
}

func (v *NullablePromoteObjectID) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value) //nolint:wrapcheck
}
