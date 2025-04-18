// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
package ingestion

import (
	"encoding/json"
	"fmt"
)

// DockerStreamsInput The selected streams of an airbyte connector.
type DockerStreamsInput struct {
	Streams []DockerStreams `json:"streams"`
}

// NewDockerStreamsInput instantiates a new DockerStreamsInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDockerStreamsInput(streams []DockerStreams) *DockerStreamsInput {
	this := &DockerStreamsInput{}
	this.Streams = streams
	return this
}

// NewEmptyDockerStreamsInput return a pointer to an empty DockerStreamsInput object.
func NewEmptyDockerStreamsInput() *DockerStreamsInput {
	return &DockerStreamsInput{}
}

// GetStreams returns the Streams field value.
func (o *DockerStreamsInput) GetStreams() []DockerStreams {
	if o == nil {
		var ret []DockerStreams
		return ret
	}

	return o.Streams
}

// GetStreamsOk returns a tuple with the Streams field value
// and a boolean to check if the value has been set.
func (o *DockerStreamsInput) GetStreamsOk() ([]DockerStreams, bool) {
	if o == nil {
		return nil, false
	}
	return o.Streams, true
}

// SetStreams sets field value.
func (o *DockerStreamsInput) SetStreams(v []DockerStreams) *DockerStreamsInput {
	o.Streams = v
	return o
}

func (o DockerStreamsInput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	toSerialize["streams"] = o.Streams
	serialized, err := json.Marshal(toSerialize)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal DockerStreamsInput: %w", err)
	}

	return serialized, nil
}

func (o DockerStreamsInput) String() string {
	out := ""
	out += fmt.Sprintf("  streams=%v\n", o.Streams)
	return fmt.Sprintf("DockerStreamsInput {\n%s}", out)
}
