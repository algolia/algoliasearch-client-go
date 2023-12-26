// File generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation.
package insights

import (
	"encoding/json"
	"fmt"
)

// ClickedObjectIDsAfterSearch Click event after an Algolia request.  Use this event to track when users click items in the search results. If you're building your category pages with Algolia, you'll also use this event.
type ClickedObjectIDsAfterSearch struct {
	// Can contain up to 64 ASCII characters.   Consider naming events consistently—for example, by adopting Segment's [object-action](https://segment.com/academy/collecting-data/naming-conventions-for-clean-data/#the-object-action-framework) framework.
	EventName string     `json:"eventName"`
	EventType ClickEvent `json:"eventType"`
	// Name of the Algolia index.
	Index string `json:"index"`
	// List of object identifiers for items of an Algolia index.
	ObjectIDs []string `json:"objectIDs"`
	// Position of the clicked objects in the search results.  The first search result has a position of 1 (not 0). You must provide 1 `position` for each `objectID`.
	Positions []int32 `json:"positions"`
	// Unique identifier for a search query.  The query ID is required for events related to search or browse requests. If you add `clickAnalytics: true` as a search request parameter, the query ID is included in the API response.
	QueryID string `json:"queryID"`
	// Anonymous or pseudonymous user identifier.   > **Note**: Never include personally identifiable information in user tokens.
	UserToken string `json:"userToken"`
	// Time of the event in milliseconds in [Unix epoch time](https://wikipedia.org/wiki/Unix_time). By default, the Insights API uses the time it receives an event as its timestamp.
	Timestamp *int64 `json:"timestamp,omitempty"`
	// User token for authenticated users.
	AuthenticatedUserToken *string `json:"authenticatedUserToken,omitempty"`
}

type ClickedObjectIDsAfterSearchOption func(f *ClickedObjectIDsAfterSearch)

func WithClickedObjectIDsAfterSearchTimestamp(val int64) ClickedObjectIDsAfterSearchOption {
	return func(f *ClickedObjectIDsAfterSearch) {
		f.Timestamp = &val
	}
}

func WithClickedObjectIDsAfterSearchAuthenticatedUserToken(val string) ClickedObjectIDsAfterSearchOption {
	return func(f *ClickedObjectIDsAfterSearch) {
		f.AuthenticatedUserToken = &val
	}
}

// NewClickedObjectIDsAfterSearch instantiates a new ClickedObjectIDsAfterSearch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewClickedObjectIDsAfterSearch(eventName string, eventType ClickEvent, index string, objectIDs []string, positions []int32, queryID string, userToken string, opts ...ClickedObjectIDsAfterSearchOption) *ClickedObjectIDsAfterSearch {
	this := &ClickedObjectIDsAfterSearch{}
	this.EventName = eventName
	this.EventType = eventType
	this.Index = index
	this.ObjectIDs = objectIDs
	this.Positions = positions
	this.QueryID = queryID
	this.UserToken = userToken
	for _, opt := range opts {
		opt(this)
	}
	return this
}

// NewClickedObjectIDsAfterSearchWithDefaults instantiates a new ClickedObjectIDsAfterSearch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewClickedObjectIDsAfterSearchWithDefaults() *ClickedObjectIDsAfterSearch {
	this := &ClickedObjectIDsAfterSearch{}
	return this
}

// GetEventName returns the EventName field value.
func (o *ClickedObjectIDsAfterSearch) GetEventName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EventName
}

// GetEventNameOk returns a tuple with the EventName field value
// and a boolean to check if the value has been set.
func (o *ClickedObjectIDsAfterSearch) GetEventNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventName, true
}

// SetEventName sets field value.
func (o *ClickedObjectIDsAfterSearch) SetEventName(v string) {
	o.EventName = v
}

// GetEventType returns the EventType field value.
func (o *ClickedObjectIDsAfterSearch) GetEventType() ClickEvent {
	if o == nil {
		var ret ClickEvent
		return ret
	}

	return o.EventType
}

// GetEventTypeOk returns a tuple with the EventType field value
// and a boolean to check if the value has been set.
func (o *ClickedObjectIDsAfterSearch) GetEventTypeOk() (*ClickEvent, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventType, true
}

// SetEventType sets field value.
func (o *ClickedObjectIDsAfterSearch) SetEventType(v ClickEvent) {
	o.EventType = v
}

// GetIndex returns the Index field value.
func (o *ClickedObjectIDsAfterSearch) GetIndex() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Index
}

// GetIndexOk returns a tuple with the Index field value
// and a boolean to check if the value has been set.
func (o *ClickedObjectIDsAfterSearch) GetIndexOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Index, true
}

// SetIndex sets field value.
func (o *ClickedObjectIDsAfterSearch) SetIndex(v string) {
	o.Index = v
}

// GetObjectIDs returns the ObjectIDs field value.
func (o *ClickedObjectIDsAfterSearch) GetObjectIDs() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.ObjectIDs
}

// GetObjectIDsOk returns a tuple with the ObjectIDs field value
// and a boolean to check if the value has been set.
func (o *ClickedObjectIDsAfterSearch) GetObjectIDsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ObjectIDs, true
}

// SetObjectIDs sets field value.
func (o *ClickedObjectIDsAfterSearch) SetObjectIDs(v []string) {
	o.ObjectIDs = v
}

// GetPositions returns the Positions field value.
func (o *ClickedObjectIDsAfterSearch) GetPositions() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}

	return o.Positions
}

// GetPositionsOk returns a tuple with the Positions field value
// and a boolean to check if the value has been set.
func (o *ClickedObjectIDsAfterSearch) GetPositionsOk() ([]int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Positions, true
}

// SetPositions sets field value.
func (o *ClickedObjectIDsAfterSearch) SetPositions(v []int32) {
	o.Positions = v
}

// GetQueryID returns the QueryID field value.
func (o *ClickedObjectIDsAfterSearch) GetQueryID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.QueryID
}

// GetQueryIDOk returns a tuple with the QueryID field value
// and a boolean to check if the value has been set.
func (o *ClickedObjectIDsAfterSearch) GetQueryIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.QueryID, true
}

// SetQueryID sets field value.
func (o *ClickedObjectIDsAfterSearch) SetQueryID(v string) {
	o.QueryID = v
}

// GetUserToken returns the UserToken field value.
func (o *ClickedObjectIDsAfterSearch) GetUserToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserToken
}

// GetUserTokenOk returns a tuple with the UserToken field value
// and a boolean to check if the value has been set.
func (o *ClickedObjectIDsAfterSearch) GetUserTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserToken, true
}

// SetUserToken sets field value.
func (o *ClickedObjectIDsAfterSearch) SetUserToken(v string) {
	o.UserToken = v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *ClickedObjectIDsAfterSearch) GetTimestamp() int64 {
	if o == nil || o.Timestamp == nil {
		var ret int64
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickedObjectIDsAfterSearch) GetTimestampOk() (*int64, bool) {
	if o == nil || o.Timestamp == nil {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *ClickedObjectIDsAfterSearch) HasTimestamp() bool {
	if o != nil && o.Timestamp != nil {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given int64 and assigns it to the Timestamp field.
func (o *ClickedObjectIDsAfterSearch) SetTimestamp(v int64) {
	o.Timestamp = &v
}

// GetAuthenticatedUserToken returns the AuthenticatedUserToken field value if set, zero value otherwise.
func (o *ClickedObjectIDsAfterSearch) GetAuthenticatedUserToken() string {
	if o == nil || o.AuthenticatedUserToken == nil {
		var ret string
		return ret
	}
	return *o.AuthenticatedUserToken
}

// GetAuthenticatedUserTokenOk returns a tuple with the AuthenticatedUserToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickedObjectIDsAfterSearch) GetAuthenticatedUserTokenOk() (*string, bool) {
	if o == nil || o.AuthenticatedUserToken == nil {
		return nil, false
	}
	return o.AuthenticatedUserToken, true
}

// HasAuthenticatedUserToken returns a boolean if a field has been set.
func (o *ClickedObjectIDsAfterSearch) HasAuthenticatedUserToken() bool {
	if o != nil && o.AuthenticatedUserToken != nil {
		return true
	}

	return false
}

// SetAuthenticatedUserToken gets a reference to the given string and assigns it to the AuthenticatedUserToken field.
func (o *ClickedObjectIDsAfterSearch) SetAuthenticatedUserToken(v string) {
	o.AuthenticatedUserToken = &v
}

func (o ClickedObjectIDsAfterSearch) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if true {
		toSerialize["eventName"] = o.EventName
	}
	if true {
		toSerialize["eventType"] = o.EventType
	}
	if true {
		toSerialize["index"] = o.Index
	}
	if true {
		toSerialize["objectIDs"] = o.ObjectIDs
	}
	if true {
		toSerialize["positions"] = o.Positions
	}
	if true {
		toSerialize["queryID"] = o.QueryID
	}
	if true {
		toSerialize["userToken"] = o.UserToken
	}
	if o.Timestamp != nil {
		toSerialize["timestamp"] = o.Timestamp
	}
	if o.AuthenticatedUserToken != nil {
		toSerialize["authenticatedUserToken"] = o.AuthenticatedUserToken
	}
	return json.Marshal(toSerialize)
}

func (o ClickedObjectIDsAfterSearch) String() string {
	out := ""
	out += fmt.Sprintf("  eventName=%v\n", o.EventName)
	out += fmt.Sprintf("  eventType=%v\n", o.EventType)
	out += fmt.Sprintf("  index=%v\n", o.Index)
	out += fmt.Sprintf("  objectIDs=%v\n", o.ObjectIDs)
	out += fmt.Sprintf("  positions=%v\n", o.Positions)
	out += fmt.Sprintf("  queryID=%v\n", o.QueryID)
	out += fmt.Sprintf("  userToken=%v\n", o.UserToken)
	out += fmt.Sprintf("  timestamp=%v\n", o.Timestamp)
	out += fmt.Sprintf("  authenticatedUserToken=%v\n", o.AuthenticatedUserToken)
	return fmt.Sprintf("ClickedObjectIDsAfterSearch {\n%s}", out)
}

type NullableClickedObjectIDsAfterSearch struct {
	value *ClickedObjectIDsAfterSearch
	isSet bool
}

func (v NullableClickedObjectIDsAfterSearch) Get() *ClickedObjectIDsAfterSearch {
	return v.value
}

func (v *NullableClickedObjectIDsAfterSearch) Set(val *ClickedObjectIDsAfterSearch) {
	v.value = val
	v.isSet = true
}

func (v NullableClickedObjectIDsAfterSearch) IsSet() bool {
	return v.isSet
}

func (v *NullableClickedObjectIDsAfterSearch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClickedObjectIDsAfterSearch(val *ClickedObjectIDsAfterSearch) *NullableClickedObjectIDsAfterSearch {
	return &NullableClickedObjectIDsAfterSearch{value: val, isSet: true}
}

func (v NullableClickedObjectIDsAfterSearch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClickedObjectIDsAfterSearch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
