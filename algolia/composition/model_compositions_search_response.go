// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
package composition

import (
	"encoding/json"
	"fmt"
)

// CompositionsSearchResponse struct for CompositionsSearchResponse.
type CompositionsSearchResponse struct {
	Run                  []CompositionRunSearchResponse `json:"run"`
	AdditionalProperties map[string]any
}

type _CompositionsSearchResponse CompositionsSearchResponse

// NewCompositionsSearchResponse instantiates a new CompositionsSearchResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCompositionsSearchResponse(run []CompositionRunSearchResponse) *CompositionsSearchResponse {
	this := &CompositionsSearchResponse{}
	this.Run = run
	return this
}

// NewEmptyCompositionsSearchResponse return a pointer to an empty CompositionsSearchResponse object.
func NewEmptyCompositionsSearchResponse() *CompositionsSearchResponse {
	return &CompositionsSearchResponse{}
}

// GetRun returns the Run field value.
func (o *CompositionsSearchResponse) GetRun() []CompositionRunSearchResponse {
	if o == nil {
		var ret []CompositionRunSearchResponse
		return ret
	}

	return o.Run
}

// GetRunOk returns a tuple with the Run field value
// and a boolean to check if the value has been set.
func (o *CompositionsSearchResponse) GetRunOk() ([]CompositionRunSearchResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.Run, true
}

// SetRun sets field value.
func (o *CompositionsSearchResponse) SetRun(v []CompositionRunSearchResponse) *CompositionsSearchResponse {
	o.Run = v
	return o
}

func (o *CompositionsSearchResponse) SetAdditionalProperty(key string, value any) *CompositionsSearchResponse {
	if o.AdditionalProperties == nil {
		o.AdditionalProperties = make(map[string]any)
	}

	o.AdditionalProperties[key] = value

	return o
}

func (o CompositionsSearchResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if true {
		toSerialize["run"] = o.Run
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	serialized, err := json.Marshal(toSerialize)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal CompositionsSearchResponse: %w", err)
	}

	return serialized, nil
}

func (o *CompositionsSearchResponse) UnmarshalJSON(bytes []byte) error {
	varCompositionsSearchResponse := _CompositionsSearchResponse{}

	err := json.Unmarshal(bytes, &varCompositionsSearchResponse)
	if err != nil {
		return fmt.Errorf("failed to unmarshal CompositionsSearchResponse: %w", err)
	}

	*o = CompositionsSearchResponse(varCompositionsSearchResponse)

	additionalProperties := make(map[string]any)

	err = json.Unmarshal(bytes, &additionalProperties)
	if err != nil {
		return fmt.Errorf("failed to unmarshal additionalProperties in CompositionsSearchResponse: %w", err)
	}

	delete(additionalProperties, "run")
	o.AdditionalProperties = additionalProperties

	return nil
}

func (o CompositionsSearchResponse) String() string {
	out := ""
	out += fmt.Sprintf("  run=%v\n", o.Run)
	for key, value := range o.AdditionalProperties {
		out += fmt.Sprintf("  %s=%v\n", key, value)
	}
	return fmt.Sprintf("CompositionsSearchResponse {\n%s}", out)
}
