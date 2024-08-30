// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
package suggestions

import (
	"encoding/json"
	"fmt"
)

// ConfigStatus struct for ConfigStatus.
type ConfigStatus struct {
	// Name of the Query Suggestions index (case-sensitive).
	IndexName *string `json:"indexName,omitempty"`
	// Whether the creation or update of the Query Suggestions index is in progress.
	IsRunning *bool `json:"isRunning,omitempty"`
	// Date and time when the Query Suggestions index was last built, in RFC 3339 format.
	LastBuiltAt *string `json:"lastBuiltAt,omitempty"`
	// Date and time when the Query Suggestions index was last updated successfully.
	LastSuccessfulBuiltAt *string `json:"lastSuccessfulBuiltAt,omitempty"`
	// Duration of the last successful build in seconds.
	LastSuccessfulBuildDuration *string `json:"lastSuccessfulBuildDuration,omitempty"`
}

type ConfigStatusOption func(f *ConfigStatus)

func WithConfigStatusIndexName(val string) ConfigStatusOption {
	return func(f *ConfigStatus) {
		f.IndexName = &val
	}
}

func WithConfigStatusIsRunning(val bool) ConfigStatusOption {
	return func(f *ConfigStatus) {
		f.IsRunning = &val
	}
}

func WithConfigStatusLastBuiltAt(val string) ConfigStatusOption {
	return func(f *ConfigStatus) {
		f.LastBuiltAt = &val
	}
}

func WithConfigStatusLastSuccessfulBuiltAt(val string) ConfigStatusOption {
	return func(f *ConfigStatus) {
		f.LastSuccessfulBuiltAt = &val
	}
}

func WithConfigStatusLastSuccessfulBuildDuration(val string) ConfigStatusOption {
	return func(f *ConfigStatus) {
		f.LastSuccessfulBuildDuration = &val
	}
}

// NewConfigStatus instantiates a new ConfigStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewConfigStatus(opts ...ConfigStatusOption) *ConfigStatus {
	this := &ConfigStatus{}
	for _, opt := range opts {
		opt(this)
	}
	return this
}

// NewEmptyConfigStatus return a pointer to an empty ConfigStatus object.
func NewEmptyConfigStatus() *ConfigStatus {
	return &ConfigStatus{}
}

// GetIndexName returns the IndexName field value if set, zero value otherwise.
func (o *ConfigStatus) GetIndexName() string {
	if o == nil || o.IndexName == nil {
		var ret string
		return ret
	}
	return *o.IndexName
}

// GetIndexNameOk returns a tuple with the IndexName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigStatus) GetIndexNameOk() (*string, bool) {
	if o == nil || o.IndexName == nil {
		return nil, false
	}
	return o.IndexName, true
}

// HasIndexName returns a boolean if a field has been set.
func (o *ConfigStatus) HasIndexName() bool {
	if o != nil && o.IndexName != nil {
		return true
	}

	return false
}

// SetIndexName gets a reference to the given string and assigns it to the IndexName field.
func (o *ConfigStatus) SetIndexName(v string) *ConfigStatus {
	o.IndexName = &v
	return o
}

// GetIsRunning returns the IsRunning field value if set, zero value otherwise.
func (o *ConfigStatus) GetIsRunning() bool {
	if o == nil || o.IsRunning == nil {
		var ret bool
		return ret
	}
	return *o.IsRunning
}

// GetIsRunningOk returns a tuple with the IsRunning field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigStatus) GetIsRunningOk() (*bool, bool) {
	if o == nil || o.IsRunning == nil {
		return nil, false
	}
	return o.IsRunning, true
}

// HasIsRunning returns a boolean if a field has been set.
func (o *ConfigStatus) HasIsRunning() bool {
	if o != nil && o.IsRunning != nil {
		return true
	}

	return false
}

// SetIsRunning gets a reference to the given bool and assigns it to the IsRunning field.
func (o *ConfigStatus) SetIsRunning(v bool) *ConfigStatus {
	o.IsRunning = &v
	return o
}

// GetLastBuiltAt returns the LastBuiltAt field value if set, zero value otherwise.
func (o *ConfigStatus) GetLastBuiltAt() string {
	if o == nil || o.LastBuiltAt == nil {
		var ret string
		return ret
	}
	return *o.LastBuiltAt
}

// GetLastBuiltAtOk returns a tuple with the LastBuiltAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigStatus) GetLastBuiltAtOk() (*string, bool) {
	if o == nil || o.LastBuiltAt == nil {
		return nil, false
	}
	return o.LastBuiltAt, true
}

// HasLastBuiltAt returns a boolean if a field has been set.
func (o *ConfigStatus) HasLastBuiltAt() bool {
	if o != nil && o.LastBuiltAt != nil {
		return true
	}

	return false
}

// SetLastBuiltAt gets a reference to the given string and assigns it to the LastBuiltAt field.
func (o *ConfigStatus) SetLastBuiltAt(v string) *ConfigStatus {
	o.LastBuiltAt = &v
	return o
}

// GetLastSuccessfulBuiltAt returns the LastSuccessfulBuiltAt field value if set, zero value otherwise.
func (o *ConfigStatus) GetLastSuccessfulBuiltAt() string {
	if o == nil || o.LastSuccessfulBuiltAt == nil {
		var ret string
		return ret
	}
	return *o.LastSuccessfulBuiltAt
}

// GetLastSuccessfulBuiltAtOk returns a tuple with the LastSuccessfulBuiltAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigStatus) GetLastSuccessfulBuiltAtOk() (*string, bool) {
	if o == nil || o.LastSuccessfulBuiltAt == nil {
		return nil, false
	}
	return o.LastSuccessfulBuiltAt, true
}

// HasLastSuccessfulBuiltAt returns a boolean if a field has been set.
func (o *ConfigStatus) HasLastSuccessfulBuiltAt() bool {
	if o != nil && o.LastSuccessfulBuiltAt != nil {
		return true
	}

	return false
}

// SetLastSuccessfulBuiltAt gets a reference to the given string and assigns it to the LastSuccessfulBuiltAt field.
func (o *ConfigStatus) SetLastSuccessfulBuiltAt(v string) *ConfigStatus {
	o.LastSuccessfulBuiltAt = &v
	return o
}

// GetLastSuccessfulBuildDuration returns the LastSuccessfulBuildDuration field value if set, zero value otherwise.
func (o *ConfigStatus) GetLastSuccessfulBuildDuration() string {
	if o == nil || o.LastSuccessfulBuildDuration == nil {
		var ret string
		return ret
	}
	return *o.LastSuccessfulBuildDuration
}

// GetLastSuccessfulBuildDurationOk returns a tuple with the LastSuccessfulBuildDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigStatus) GetLastSuccessfulBuildDurationOk() (*string, bool) {
	if o == nil || o.LastSuccessfulBuildDuration == nil {
		return nil, false
	}
	return o.LastSuccessfulBuildDuration, true
}

// HasLastSuccessfulBuildDuration returns a boolean if a field has been set.
func (o *ConfigStatus) HasLastSuccessfulBuildDuration() bool {
	if o != nil && o.LastSuccessfulBuildDuration != nil {
		return true
	}

	return false
}

// SetLastSuccessfulBuildDuration gets a reference to the given string and assigns it to the LastSuccessfulBuildDuration field.
func (o *ConfigStatus) SetLastSuccessfulBuildDuration(v string) *ConfigStatus {
	o.LastSuccessfulBuildDuration = &v
	return o
}

func (o ConfigStatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.IndexName != nil {
		toSerialize["indexName"] = o.IndexName
	}
	if o.IsRunning != nil {
		toSerialize["isRunning"] = o.IsRunning
	}
	if o.LastBuiltAt != nil {
		toSerialize["lastBuiltAt"] = o.LastBuiltAt
	}
	if o.LastSuccessfulBuiltAt != nil {
		toSerialize["lastSuccessfulBuiltAt"] = o.LastSuccessfulBuiltAt
	}
	if o.LastSuccessfulBuildDuration != nil {
		toSerialize["lastSuccessfulBuildDuration"] = o.LastSuccessfulBuildDuration
	}
	serialized, err := json.Marshal(toSerialize)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal ConfigStatus: %w", err)
	}

	return serialized, nil
}

func (o ConfigStatus) String() string {
	out := ""
	out += fmt.Sprintf("  indexName=%v\n", o.IndexName)
	out += fmt.Sprintf("  isRunning=%v\n", o.IsRunning)
	out += fmt.Sprintf("  lastBuiltAt=%v\n", o.LastBuiltAt)
	out += fmt.Sprintf("  lastSuccessfulBuiltAt=%v\n", o.LastSuccessfulBuiltAt)
	out += fmt.Sprintf("  lastSuccessfulBuildDuration=%v\n", o.LastSuccessfulBuildDuration)
	return fmt.Sprintf("ConfigStatus {\n%s}", out)
}