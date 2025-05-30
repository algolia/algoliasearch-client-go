// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
package ingestion

import (
	"encoding/json"
	"fmt"

	"github.com/algolia/algoliasearch-client-go/v4/algolia/utils"
)

// Destination Destinations are Algolia resources like indices or event streams.
type Destination struct {
	// Universally unique identifier (UUID) of a destination resource.
	DestinationID string          `json:"destinationID"`
	Type          DestinationType `json:"type"`
	// Descriptive name for the resource.
	Name string `json:"name"`
	// Owner of the resource.
	Owner utils.Nullable[string] `json:"owner,omitempty"`
	Input DestinationInput       `json:"input"`
	// Date of creation in RFC 3339 format.
	CreatedAt string `json:"createdAt"`
	// Date of last update in RFC 3339 format.
	UpdatedAt string `json:"updatedAt"`
	// Universally unique identifier (UUID) of an authentication resource.
	AuthenticationID  *string  `json:"authenticationID,omitempty"`
	TransformationIDs []string `json:"transformationIDs,omitempty"`
}

type DestinationOption func(f *Destination)

func WithDestinationOwner(val utils.Nullable[string]) DestinationOption {
	return func(f *Destination) {
		f.Owner = val
	}
}

func WithDestinationAuthenticationID(val string) DestinationOption {
	return func(f *Destination) {
		f.AuthenticationID = &val
	}
}

func WithDestinationTransformationIDs(val []string) DestinationOption {
	return func(f *Destination) {
		f.TransformationIDs = val
	}
}

// NewDestination instantiates a new Destination object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDestination(destinationID string, type_ DestinationType, name string, input DestinationInput, createdAt string, updatedAt string, opts ...DestinationOption) *Destination {
	this := &Destination{}
	this.DestinationID = destinationID
	this.Type = type_
	this.Name = name
	this.Input = input
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	for _, opt := range opts {
		opt(this)
	}
	return this
}

// NewEmptyDestination return a pointer to an empty Destination object.
func NewEmptyDestination() *Destination {
	return &Destination{}
}

// GetDestinationID returns the DestinationID field value.
func (o *Destination) GetDestinationID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DestinationID
}

// GetDestinationIDOk returns a tuple with the DestinationID field value
// and a boolean to check if the value has been set.
func (o *Destination) GetDestinationIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DestinationID, true
}

// SetDestinationID sets field value.
func (o *Destination) SetDestinationID(v string) *Destination {
	o.DestinationID = v
	return o
}

// GetType returns the Type field value.
func (o *Destination) GetType() DestinationType {
	if o == nil {
		var ret DestinationType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *Destination) GetTypeOk() (*DestinationType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *Destination) SetType(v DestinationType) *Destination {
	o.Type = v
	return o
}

// GetName returns the Name field value.
func (o *Destination) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Destination) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *Destination) SetName(v string) *Destination {
	o.Name = v
	return o
}

// GetOwner returns the Owner field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Destination) GetOwner() string {
	if o == nil || o.Owner.Get() == nil {
		var ret string
		return ret
	}
	return *o.Owner.Get()
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *Destination) GetOwnerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Owner.Get(), o.Owner.IsSet()
}

// HasOwner returns a boolean if a field has been set.
func (o *Destination) HasOwner() bool {
	if o != nil && o.Owner.IsSet() {
		return true
	}

	return false
}

// SetOwner gets a reference to the given utils.Nullable[string] and assigns it to the Owner field.
func (o *Destination) SetOwner(v string) *Destination {
	o.Owner.Set(&v)
	return o
}

// SetOwnerNil sets the value for Owner to be an explicit nil.
func (o *Destination) SetOwnerNil() {
	o.Owner.Set(nil)
}

// UnsetOwner ensures that no value is present for Owner, not even an explicit nil.
func (o *Destination) UnsetOwner() {
	o.Owner.Unset()
}

// GetInput returns the Input field value.
func (o *Destination) GetInput() DestinationInput {
	if o == nil {
		var ret DestinationInput
		return ret
	}

	return o.Input
}

// GetInputOk returns a tuple with the Input field value
// and a boolean to check if the value has been set.
func (o *Destination) GetInputOk() (*DestinationInput, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Input, true
}

// SetInput sets field value.
func (o *Destination) SetInput(v *DestinationInput) *Destination {
	o.Input = *v
	return o
}

// GetCreatedAt returns the CreatedAt field value.
func (o *Destination) GetCreatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *Destination) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value.
func (o *Destination) SetCreatedAt(v string) *Destination {
	o.CreatedAt = v
	return o
}

// GetUpdatedAt returns the UpdatedAt field value.
func (o *Destination) GetUpdatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *Destination) GetUpdatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value.
func (o *Destination) SetUpdatedAt(v string) *Destination {
	o.UpdatedAt = v
	return o
}

// GetAuthenticationID returns the AuthenticationID field value if set, zero value otherwise.
func (o *Destination) GetAuthenticationID() string {
	if o == nil || o.AuthenticationID == nil {
		var ret string
		return ret
	}
	return *o.AuthenticationID
}

// GetAuthenticationIDOk returns a tuple with the AuthenticationID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Destination) GetAuthenticationIDOk() (*string, bool) {
	if o == nil || o.AuthenticationID == nil {
		return nil, false
	}
	return o.AuthenticationID, true
}

// HasAuthenticationID returns a boolean if a field has been set.
func (o *Destination) HasAuthenticationID() bool {
	if o != nil && o.AuthenticationID != nil {
		return true
	}

	return false
}

// SetAuthenticationID gets a reference to the given string and assigns it to the AuthenticationID field.
func (o *Destination) SetAuthenticationID(v string) *Destination {
	o.AuthenticationID = &v
	return o
}

// GetTransformationIDs returns the TransformationIDs field value if set, zero value otherwise.
func (o *Destination) GetTransformationIDs() []string {
	if o == nil || o.TransformationIDs == nil {
		var ret []string
		return ret
	}
	return o.TransformationIDs
}

// GetTransformationIDsOk returns a tuple with the TransformationIDs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Destination) GetTransformationIDsOk() ([]string, bool) {
	if o == nil || o.TransformationIDs == nil {
		return nil, false
	}
	return o.TransformationIDs, true
}

// HasTransformationIDs returns a boolean if a field has been set.
func (o *Destination) HasTransformationIDs() bool {
	if o != nil && o.TransformationIDs != nil {
		return true
	}

	return false
}

// SetTransformationIDs gets a reference to the given []string and assigns it to the TransformationIDs field.
func (o *Destination) SetTransformationIDs(v []string) *Destination {
	o.TransformationIDs = v
	return o
}

func (o Destination) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	toSerialize["destinationID"] = o.DestinationID
	toSerialize["type"] = o.Type
	toSerialize["name"] = o.Name
	if o.Owner.IsSet() {
		toSerialize["owner"] = o.Owner.Get()
	}
	toSerialize["input"] = o.Input
	toSerialize["createdAt"] = o.CreatedAt
	toSerialize["updatedAt"] = o.UpdatedAt
	if o.AuthenticationID != nil {
		toSerialize["authenticationID"] = o.AuthenticationID
	}
	if o.TransformationIDs != nil {
		toSerialize["transformationIDs"] = o.TransformationIDs
	}
	serialized, err := json.Marshal(toSerialize)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal Destination: %w", err)
	}

	return serialized, nil
}

func (o Destination) String() string {
	out := ""
	out += fmt.Sprintf("  destinationID=%v\n", o.DestinationID)
	out += fmt.Sprintf("  type=%v\n", o.Type)
	out += fmt.Sprintf("  name=%v\n", o.Name)
	out += fmt.Sprintf("  owner=%v\n", o.Owner)
	out += fmt.Sprintf("  input=%v\n", o.Input)
	out += fmt.Sprintf("  createdAt=%v\n", o.CreatedAt)
	out += fmt.Sprintf("  updatedAt=%v\n", o.UpdatedAt)
	out += fmt.Sprintf("  authenticationID=%v\n", o.AuthenticationID)
	out += fmt.Sprintf("  transformationIDs=%v\n", o.TransformationIDs)
	return fmt.Sprintf("Destination {\n%s}", out)
}
