// File generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation.
package monitoring

import (
	"encoding/json"
	"fmt"
)

// InventoryResponse struct for InventoryResponse.
type InventoryResponse struct {
	Inventory []Server `json:"inventory,omitempty"`
}

type InventoryResponseOption func(f *InventoryResponse)

func WithInventoryResponseInventory(val []Server) InventoryResponseOption {
	return func(f *InventoryResponse) {
		f.Inventory = val
	}
}

// NewInventoryResponse instantiates a new InventoryResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewInventoryResponse(opts ...InventoryResponseOption) *InventoryResponse {
	this := &InventoryResponse{}
	for _, opt := range opts {
		opt(this)
	}
	return this
}

// NewEmptyInventoryResponse return a pointer to an empty InventoryResponse object.
func NewEmptyInventoryResponse() *InventoryResponse {
	return &InventoryResponse{}
}

// GetInventory returns the Inventory field value if set, zero value otherwise.
func (o *InventoryResponse) GetInventory() []Server {
	if o == nil || o.Inventory == nil {
		var ret []Server
		return ret
	}
	return o.Inventory
}

// GetInventoryOk returns a tuple with the Inventory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryResponse) GetInventoryOk() ([]Server, bool) {
	if o == nil || o.Inventory == nil {
		return nil, false
	}
	return o.Inventory, true
}

// HasInventory returns a boolean if a field has been set.
func (o *InventoryResponse) HasInventory() bool {
	if o != nil && o.Inventory != nil {
		return true
	}

	return false
}

// SetInventory gets a reference to the given []Server and assigns it to the Inventory field.
func (o *InventoryResponse) SetInventory(v []Server) *InventoryResponse {
	o.Inventory = v
	return o
}

func (o InventoryResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Inventory != nil {
		toSerialize["inventory"] = o.Inventory
	}
	serialized, err := json.Marshal(toSerialize)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal InventoryResponse: %w", err)
	}

	return serialized, nil
}

func (o InventoryResponse) String() string {
	out := ""
	out += fmt.Sprintf("  inventory=%v\n", o.Inventory)
	return fmt.Sprintf("InventoryResponse {\n%s}", out)
}

type NullableInventoryResponse struct {
	value *InventoryResponse
	isSet bool
}

func (v NullableInventoryResponse) Get() *InventoryResponse {
	return v.value
}

func (v *NullableInventoryResponse) Set(val *InventoryResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableInventoryResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableInventoryResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInventoryResponse(val *InventoryResponse) *NullableInventoryResponse {
	return &NullableInventoryResponse{value: val, isSet: true}
}

func (v NullableInventoryResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value) //nolint:wrapcheck
}

func (v *NullableInventoryResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value) //nolint:wrapcheck
}
