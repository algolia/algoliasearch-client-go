// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
package recommend

import (
	"encoding/json"
	"fmt"
)

// MatchedGeoLocation struct for MatchedGeoLocation.
type MatchedGeoLocation struct {
	// Latitude of the matched location.
	Lat *float64 `json:"lat,omitempty"`
	// Longitude of the matched location.
	Lng *float64 `json:"lng,omitempty"`
	// Distance between the matched location and the search location (in meters).
	Distance *int32 `json:"distance,omitempty"`
}

type MatchedGeoLocationOption func(f *MatchedGeoLocation)

func WithMatchedGeoLocationLat(val float64) MatchedGeoLocationOption {
	return func(f *MatchedGeoLocation) {
		f.Lat = &val
	}
}

func WithMatchedGeoLocationLng(val float64) MatchedGeoLocationOption {
	return func(f *MatchedGeoLocation) {
		f.Lng = &val
	}
}

func WithMatchedGeoLocationDistance(val int32) MatchedGeoLocationOption {
	return func(f *MatchedGeoLocation) {
		f.Distance = &val
	}
}

// NewMatchedGeoLocation instantiates a new MatchedGeoLocation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewMatchedGeoLocation(opts ...MatchedGeoLocationOption) *MatchedGeoLocation {
	this := &MatchedGeoLocation{}
	for _, opt := range opts {
		opt(this)
	}
	return this
}

// NewEmptyMatchedGeoLocation return a pointer to an empty MatchedGeoLocation object.
func NewEmptyMatchedGeoLocation() *MatchedGeoLocation {
	return &MatchedGeoLocation{}
}

// GetLat returns the Lat field value if set, zero value otherwise.
func (o *MatchedGeoLocation) GetLat() float64 {
	if o == nil || o.Lat == nil {
		var ret float64
		return ret
	}
	return *o.Lat
}

// GetLatOk returns a tuple with the Lat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MatchedGeoLocation) GetLatOk() (*float64, bool) {
	if o == nil || o.Lat == nil {
		return nil, false
	}
	return o.Lat, true
}

// HasLat returns a boolean if a field has been set.
func (o *MatchedGeoLocation) HasLat() bool {
	if o != nil && o.Lat != nil {
		return true
	}

	return false
}

// SetLat gets a reference to the given float64 and assigns it to the Lat field.
func (o *MatchedGeoLocation) SetLat(v float64) *MatchedGeoLocation {
	o.Lat = &v
	return o
}

// GetLng returns the Lng field value if set, zero value otherwise.
func (o *MatchedGeoLocation) GetLng() float64 {
	if o == nil || o.Lng == nil {
		var ret float64
		return ret
	}
	return *o.Lng
}

// GetLngOk returns a tuple with the Lng field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MatchedGeoLocation) GetLngOk() (*float64, bool) {
	if o == nil || o.Lng == nil {
		return nil, false
	}
	return o.Lng, true
}

// HasLng returns a boolean if a field has been set.
func (o *MatchedGeoLocation) HasLng() bool {
	if o != nil && o.Lng != nil {
		return true
	}

	return false
}

// SetLng gets a reference to the given float64 and assigns it to the Lng field.
func (o *MatchedGeoLocation) SetLng(v float64) *MatchedGeoLocation {
	o.Lng = &v
	return o
}

// GetDistance returns the Distance field value if set, zero value otherwise.
func (o *MatchedGeoLocation) GetDistance() int32 {
	if o == nil || o.Distance == nil {
		var ret int32
		return ret
	}
	return *o.Distance
}

// GetDistanceOk returns a tuple with the Distance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MatchedGeoLocation) GetDistanceOk() (*int32, bool) {
	if o == nil || o.Distance == nil {
		return nil, false
	}
	return o.Distance, true
}

// HasDistance returns a boolean if a field has been set.
func (o *MatchedGeoLocation) HasDistance() bool {
	if o != nil && o.Distance != nil {
		return true
	}

	return false
}

// SetDistance gets a reference to the given int32 and assigns it to the Distance field.
func (o *MatchedGeoLocation) SetDistance(v int32) *MatchedGeoLocation {
	o.Distance = &v
	return o
}

func (o MatchedGeoLocation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Lat != nil {
		toSerialize["lat"] = o.Lat
	}
	if o.Lng != nil {
		toSerialize["lng"] = o.Lng
	}
	if o.Distance != nil {
		toSerialize["distance"] = o.Distance
	}
	serialized, err := json.Marshal(toSerialize)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal MatchedGeoLocation: %w", err)
	}

	return serialized, nil
}

func (o MatchedGeoLocation) String() string {
	out := ""
	out += fmt.Sprintf("  lat=%v\n", o.Lat)
	out += fmt.Sprintf("  lng=%v\n", o.Lng)
	out += fmt.Sprintf("  distance=%v\n", o.Distance)
	return fmt.Sprintf("MatchedGeoLocation {\n%s}", out)
}
