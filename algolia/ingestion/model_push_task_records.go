// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
package ingestion

import (
	"encoding/json"
	"fmt"
)

// PushTaskRecords struct for PushTaskRecords.
type PushTaskRecords struct {
	// Unique record identifier.
	ObjectID             string `json:"objectID"`
	AdditionalProperties map[string]any
}

type _PushTaskRecords PushTaskRecords

// NewPushTaskRecords instantiates a new PushTaskRecords object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewPushTaskRecords(objectID string) *PushTaskRecords {
	this := &PushTaskRecords{}
	this.ObjectID = objectID
	return this
}

// NewEmptyPushTaskRecords return a pointer to an empty PushTaskRecords object.
func NewEmptyPushTaskRecords() *PushTaskRecords {
	return &PushTaskRecords{}
}

// GetObjectID returns the ObjectID field value.
func (o *PushTaskRecords) GetObjectID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectID
}

// GetObjectIDOk returns a tuple with the ObjectID field value
// and a boolean to check if the value has been set.
func (o *PushTaskRecords) GetObjectIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectID, true
}

// SetObjectID sets field value.
func (o *PushTaskRecords) SetObjectID(v string) *PushTaskRecords {
	o.ObjectID = v
	return o
}

func (o *PushTaskRecords) SetAdditionalProperty(key string, value any) *PushTaskRecords {
	if o.AdditionalProperties == nil {
		o.AdditionalProperties = make(map[string]any)
	}

	o.AdditionalProperties[key] = value

	return o
}

func (o PushTaskRecords) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	toSerialize["objectID"] = o.ObjectID

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	serialized, err := json.Marshal(toSerialize)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal PushTaskRecords: %w", err)
	}

	return serialized, nil
}

func (o *PushTaskRecords) UnmarshalJSON(bytes []byte) error {
	varPushTaskRecords := _PushTaskRecords{}

	err := json.Unmarshal(bytes, &varPushTaskRecords)
	if err != nil {
		return fmt.Errorf("failed to unmarshal PushTaskRecords: %w", err)
	}

	*o = PushTaskRecords(varPushTaskRecords)

	additionalProperties := make(map[string]any)

	err = json.Unmarshal(bytes, &additionalProperties)
	if err != nil {
		return fmt.Errorf("failed to unmarshal additionalProperties in PushTaskRecords: %w", err)
	}

	delete(additionalProperties, "objectID")
	o.AdditionalProperties = additionalProperties

	return nil
}

func (o PushTaskRecords) String() string {
	out := ""
	out += fmt.Sprintf("  objectID=%v\n", o.ObjectID)
	for key, value := range o.AdditionalProperties {
		out += fmt.Sprintf("  %s=%v\n", key, value)
	}
	return fmt.Sprintf("PushTaskRecords {\n%s}", out)
}
