// File generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation.
package suggestions

import (
	"encoding/json"
	"fmt"
)

// GetLogFile200Response struct for GetLogFile200Response.
type GetLogFile200Response struct {
	// Timestamp in [ISO-8601](https://wikipedia.org/wiki/ISO_8601) format.
	Timestamp *string   `json:"timestamp,omitempty"`
	Level     *LogLevel `json:"level,omitempty"`
	// Details about this log entry.
	Message *string `json:"message,omitempty"`
	// Level indicating the position of a suggestion in a hierarchy of records.   For example, a `contextLevel` of 1 indicates that this suggestion belongs to a previous suggestion with `contextLevel` 0.
	ContextLevel *int32 `json:"contextLevel,omitempty"`
}

type GetLogFile200ResponseOption func(f *GetLogFile200Response)

func WithGetLogFile200ResponseTimestamp(val string) GetLogFile200ResponseOption {
	return func(f *GetLogFile200Response) {
		f.Timestamp = &val
	}
}

func WithGetLogFile200ResponseLevel(val LogLevel) GetLogFile200ResponseOption {
	return func(f *GetLogFile200Response) {
		f.Level = &val
	}
}

func WithGetLogFile200ResponseMessage(val string) GetLogFile200ResponseOption {
	return func(f *GetLogFile200Response) {
		f.Message = &val
	}
}

func WithGetLogFile200ResponseContextLevel(val int32) GetLogFile200ResponseOption {
	return func(f *GetLogFile200Response) {
		f.ContextLevel = &val
	}
}

// NewGetLogFile200Response instantiates a new GetLogFile200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewGetLogFile200Response(opts ...GetLogFile200ResponseOption) *GetLogFile200Response {
	this := &GetLogFile200Response{}
	for _, opt := range opts {
		opt(this)
	}
	return this
}

// NewGetLogFile200ResponseWithDefaults instantiates a new GetLogFile200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewGetLogFile200ResponseWithDefaults() *GetLogFile200Response {
	this := &GetLogFile200Response{}
	return this
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *GetLogFile200Response) GetTimestamp() string {
	if o == nil || o.Timestamp == nil {
		var ret string
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLogFile200Response) GetTimestampOk() (*string, bool) {
	if o == nil || o.Timestamp == nil {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *GetLogFile200Response) HasTimestamp() bool {
	if o != nil && o.Timestamp != nil {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given string and assigns it to the Timestamp field.
func (o *GetLogFile200Response) SetTimestamp(v string) {
	o.Timestamp = &v
}

// GetLevel returns the Level field value if set, zero value otherwise.
func (o *GetLogFile200Response) GetLevel() LogLevel {
	if o == nil || o.Level == nil {
		var ret LogLevel
		return ret
	}
	return *o.Level
}

// GetLevelOk returns a tuple with the Level field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLogFile200Response) GetLevelOk() (*LogLevel, bool) {
	if o == nil || o.Level == nil {
		return nil, false
	}
	return o.Level, true
}

// HasLevel returns a boolean if a field has been set.
func (o *GetLogFile200Response) HasLevel() bool {
	if o != nil && o.Level != nil {
		return true
	}

	return false
}

// SetLevel gets a reference to the given LogLevel and assigns it to the Level field.
func (o *GetLogFile200Response) SetLevel(v LogLevel) {
	o.Level = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *GetLogFile200Response) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLogFile200Response) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *GetLogFile200Response) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *GetLogFile200Response) SetMessage(v string) {
	o.Message = &v
}

// GetContextLevel returns the ContextLevel field value if set, zero value otherwise.
func (o *GetLogFile200Response) GetContextLevel() int32 {
	if o == nil || o.ContextLevel == nil {
		var ret int32
		return ret
	}
	return *o.ContextLevel
}

// GetContextLevelOk returns a tuple with the ContextLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLogFile200Response) GetContextLevelOk() (*int32, bool) {
	if o == nil || o.ContextLevel == nil {
		return nil, false
	}
	return o.ContextLevel, true
}

// HasContextLevel returns a boolean if a field has been set.
func (o *GetLogFile200Response) HasContextLevel() bool {
	if o != nil && o.ContextLevel != nil {
		return true
	}

	return false
}

// SetContextLevel gets a reference to the given int32 and assigns it to the ContextLevel field.
func (o *GetLogFile200Response) SetContextLevel(v int32) {
	o.ContextLevel = &v
}

func (o GetLogFile200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Timestamp != nil {
		toSerialize["timestamp"] = o.Timestamp
	}
	if o.Level != nil {
		toSerialize["level"] = o.Level
	}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	if o.ContextLevel != nil {
		toSerialize["contextLevel"] = o.ContextLevel
	}
	return json.Marshal(toSerialize)
}

func (o GetLogFile200Response) String() string {
	out := ""
	out += fmt.Sprintf("  timestamp=%v\n", o.Timestamp)
	out += fmt.Sprintf("  level=%v\n", o.Level)
	out += fmt.Sprintf("  message=%v\n", o.Message)
	out += fmt.Sprintf("  contextLevel=%v\n", o.ContextLevel)
	return fmt.Sprintf("GetLogFile200Response {\n%s}", out)
}

type NullableGetLogFile200Response struct {
	value *GetLogFile200Response
	isSet bool
}

func (v NullableGetLogFile200Response) Get() *GetLogFile200Response {
	return v.value
}

func (v *NullableGetLogFile200Response) Set(val *GetLogFile200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetLogFile200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetLogFile200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetLogFile200Response(val *GetLogFile200Response) *NullableGetLogFile200Response {
	return &NullableGetLogFile200Response{value: val, isSet: true}
}

func (v NullableGetLogFile200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetLogFile200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
