// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
package search

import (
	"encoding/json"
	"fmt"
)

// UserHighlightResult struct for UserHighlightResult.
type UserHighlightResult struct {
	UserID      HighlightResult `json:"userID"`
	ClusterName HighlightResult `json:"clusterName"`
}

// NewUserHighlightResult instantiates a new UserHighlightResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewUserHighlightResult(userID HighlightResult, clusterName HighlightResult) *UserHighlightResult {
	this := &UserHighlightResult{}
	this.UserID = userID
	this.ClusterName = clusterName
	return this
}

// NewEmptyUserHighlightResult return a pointer to an empty UserHighlightResult object.
func NewEmptyUserHighlightResult() *UserHighlightResult {
	return &UserHighlightResult{}
}

// GetUserID returns the UserID field value.
func (o *UserHighlightResult) GetUserID() HighlightResult {
	if o == nil {
		var ret HighlightResult
		return ret
	}

	return o.UserID
}

// GetUserIDOk returns a tuple with the UserID field value
// and a boolean to check if the value has been set.
func (o *UserHighlightResult) GetUserIDOk() (*HighlightResult, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserID, true
}

// SetUserID sets field value.
func (o *UserHighlightResult) SetUserID(v *HighlightResult) *UserHighlightResult {
	o.UserID = *v
	return o
}

// GetClusterName returns the ClusterName field value.
func (o *UserHighlightResult) GetClusterName() HighlightResult {
	if o == nil {
		var ret HighlightResult
		return ret
	}

	return o.ClusterName
}

// GetClusterNameOk returns a tuple with the ClusterName field value
// and a boolean to check if the value has been set.
func (o *UserHighlightResult) GetClusterNameOk() (*HighlightResult, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClusterName, true
}

// SetClusterName sets field value.
func (o *UserHighlightResult) SetClusterName(v *HighlightResult) *UserHighlightResult {
	o.ClusterName = *v
	return o
}

func (o UserHighlightResult) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	toSerialize["userID"] = o.UserID
	toSerialize["clusterName"] = o.ClusterName
	serialized, err := json.Marshal(toSerialize)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal UserHighlightResult: %w", err)
	}

	return serialized, nil
}

func (o UserHighlightResult) String() string {
	out := ""
	out += fmt.Sprintf("  userID=%v\n", o.UserID)
	out += fmt.Sprintf("  clusterName=%v\n", o.ClusterName)
	return fmt.Sprintf("UserHighlightResult {\n%s}", out)
}
