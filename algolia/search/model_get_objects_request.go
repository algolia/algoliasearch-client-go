// File generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation.
package search

import (
	"encoding/json"
	"fmt"
)

// GetObjectsRequest Request body for retrieving records.
type GetObjectsRequest struct {
	// Attributes to retrieve. If not specified, all retrievable attributes are returned.
	AttributesToRetrieve []string `json:"attributesToRetrieve,omitempty"`
	// Object ID for the record to retrieve.
	ObjectID string `json:"objectID"`
	// Index from which to retrieve the records.
	IndexName string `json:"indexName"`
}

type GetObjectsRequestOption func(f *GetObjectsRequest)

func WithGetObjectsRequestAttributesToRetrieve(val []string) GetObjectsRequestOption {
	return func(f *GetObjectsRequest) {
		f.AttributesToRetrieve = val
	}
}

// NewGetObjectsRequest instantiates a new GetObjectsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewGetObjectsRequest(objectID string, indexName string, opts ...GetObjectsRequestOption) *GetObjectsRequest {
	this := &GetObjectsRequest{}
	this.ObjectID = objectID
	this.IndexName = indexName
	for _, opt := range opts {
		opt(this)
	}
	return this
}

// NewEmptyGetObjectsRequest return a pointer to an empty GetObjectsRequest object.
func NewEmptyGetObjectsRequest() *GetObjectsRequest {
	return &GetObjectsRequest{}
}

// GetAttributesToRetrieve returns the AttributesToRetrieve field value if set, zero value otherwise.
func (o *GetObjectsRequest) GetAttributesToRetrieve() []string {
	if o == nil || o.AttributesToRetrieve == nil {
		var ret []string
		return ret
	}
	return o.AttributesToRetrieve
}

// GetAttributesToRetrieveOk returns a tuple with the AttributesToRetrieve field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetObjectsRequest) GetAttributesToRetrieveOk() ([]string, bool) {
	if o == nil || o.AttributesToRetrieve == nil {
		return nil, false
	}
	return o.AttributesToRetrieve, true
}

// HasAttributesToRetrieve returns a boolean if a field has been set.
func (o *GetObjectsRequest) HasAttributesToRetrieve() bool {
	if o != nil && o.AttributesToRetrieve != nil {
		return true
	}

	return false
}

// SetAttributesToRetrieve gets a reference to the given []string and assigns it to the AttributesToRetrieve field.
func (o *GetObjectsRequest) SetAttributesToRetrieve(v []string) *GetObjectsRequest {
	o.AttributesToRetrieve = v
	return o
}

// GetObjectID returns the ObjectID field value.
func (o *GetObjectsRequest) GetObjectID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectID
}

// GetObjectIDOk returns a tuple with the ObjectID field value
// and a boolean to check if the value has been set.
func (o *GetObjectsRequest) GetObjectIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectID, true
}

// SetObjectID sets field value.
func (o *GetObjectsRequest) SetObjectID(v string) *GetObjectsRequest {
	o.ObjectID = v
	return o
}

// GetIndexName returns the IndexName field value.
func (o *GetObjectsRequest) GetIndexName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IndexName
}

// GetIndexNameOk returns a tuple with the IndexName field value
// and a boolean to check if the value has been set.
func (o *GetObjectsRequest) GetIndexNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IndexName, true
}

// SetIndexName sets field value.
func (o *GetObjectsRequest) SetIndexName(v string) *GetObjectsRequest {
	o.IndexName = v
	return o
}

func (o GetObjectsRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.AttributesToRetrieve != nil {
		toSerialize["attributesToRetrieve"] = o.AttributesToRetrieve
	}
	if true {
		toSerialize["objectID"] = o.ObjectID
	}
	if true {
		toSerialize["indexName"] = o.IndexName
	}
	serialized, err := json.Marshal(toSerialize)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal GetObjectsRequest: %w", err)
	}

	return serialized, nil
}

func (o GetObjectsRequest) String() string {
	out := ""
	out += fmt.Sprintf("  attributesToRetrieve=%v\n", o.AttributesToRetrieve)
	out += fmt.Sprintf("  objectID=%v\n", o.ObjectID)
	out += fmt.Sprintf("  indexName=%v\n", o.IndexName)
	return fmt.Sprintf("GetObjectsRequest {\n%s}", out)
}

type NullableGetObjectsRequest struct {
	value *GetObjectsRequest
	isSet bool
}

func (v NullableGetObjectsRequest) Get() *GetObjectsRequest {
	return v.value
}

func (v *NullableGetObjectsRequest) Set(val *GetObjectsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGetObjectsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGetObjectsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetObjectsRequest(val *GetObjectsRequest) *NullableGetObjectsRequest {
	return &NullableGetObjectsRequest{value: val, isSet: true}
}

func (v NullableGetObjectsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value) //nolint:wrapcheck
}

func (v *NullableGetObjectsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value) //nolint:wrapcheck
}
