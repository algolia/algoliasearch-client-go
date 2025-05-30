// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
package ingestion

import (
	"encoding/json"
	"fmt"
)

// SourceDocker struct for SourceDocker.
type SourceDocker struct {
	// Name of the connector.
	Image string `json:"image"`
	// Configuration of the spec.
	Configuration map[string]any `json:"configuration"`
}

// NewSourceDocker instantiates a new SourceDocker object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSourceDocker(image string, configuration map[string]any) *SourceDocker {
	this := &SourceDocker{}
	this.Image = image
	this.Configuration = configuration
	return this
}

// NewEmptySourceDocker return a pointer to an empty SourceDocker object.
func NewEmptySourceDocker() *SourceDocker {
	return &SourceDocker{}
}

// GetImage returns the Image field value.
func (o *SourceDocker) GetImage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Image
}

// GetImageOk returns a tuple with the Image field value
// and a boolean to check if the value has been set.
func (o *SourceDocker) GetImageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Image, true
}

// SetImage sets field value.
func (o *SourceDocker) SetImage(v string) *SourceDocker {
	o.Image = v
	return o
}

// GetConfiguration returns the Configuration field value.
func (o *SourceDocker) GetConfiguration() map[string]any {
	if o == nil {
		var ret map[string]any
		return ret
	}

	return o.Configuration
}

// GetConfigurationOk returns a tuple with the Configuration field value
// and a boolean to check if the value has been set.
func (o *SourceDocker) GetConfigurationOk() (map[string]any, bool) {
	if o == nil {
		return nil, false
	}
	return o.Configuration, true
}

// SetConfiguration sets field value.
func (o *SourceDocker) SetConfiguration(v map[string]any) *SourceDocker {
	o.Configuration = v
	return o
}

func (o SourceDocker) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	toSerialize["image"] = o.Image
	toSerialize["configuration"] = o.Configuration
	serialized, err := json.Marshal(toSerialize)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal SourceDocker: %w", err)
	}

	return serialized, nil
}

func (o SourceDocker) String() string {
	out := ""
	out += fmt.Sprintf("  image=%v\n", o.Image)
	out += fmt.Sprintf("  configuration=%v\n", o.Configuration)
	return fmt.Sprintf("SourceDocker {\n%s}", out)
}
