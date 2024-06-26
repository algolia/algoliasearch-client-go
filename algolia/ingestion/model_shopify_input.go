// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
package ingestion

import (
	"encoding/json"
	"fmt"
)

// ShopifyInput Represents the required elements of the task input when using a `shopify` source.
type ShopifyInput struct {
	Metafields []ShopifyMetafield `json:"metafields"`
	Market     ShopifyMarket      `json:"market"`
}

// NewShopifyInput instantiates a new ShopifyInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewShopifyInput(metafields []ShopifyMetafield, market ShopifyMarket) *ShopifyInput {
	this := &ShopifyInput{}
	this.Metafields = metafields
	this.Market = market
	return this
}

// NewEmptyShopifyInput return a pointer to an empty ShopifyInput object.
func NewEmptyShopifyInput() *ShopifyInput {
	return &ShopifyInput{}
}

// GetMetafields returns the Metafields field value.
func (o *ShopifyInput) GetMetafields() []ShopifyMetafield {
	if o == nil {
		var ret []ShopifyMetafield
		return ret
	}

	return o.Metafields
}

// GetMetafieldsOk returns a tuple with the Metafields field value
// and a boolean to check if the value has been set.
func (o *ShopifyInput) GetMetafieldsOk() ([]ShopifyMetafield, bool) {
	if o == nil {
		return nil, false
	}
	return o.Metafields, true
}

// SetMetafields sets field value.
func (o *ShopifyInput) SetMetafields(v []ShopifyMetafield) *ShopifyInput {
	o.Metafields = v
	return o
}

// GetMarket returns the Market field value.
func (o *ShopifyInput) GetMarket() ShopifyMarket {
	if o == nil {
		var ret ShopifyMarket
		return ret
	}

	return o.Market
}

// GetMarketOk returns a tuple with the Market field value
// and a boolean to check if the value has been set.
func (o *ShopifyInput) GetMarketOk() (*ShopifyMarket, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Market, true
}

// SetMarket sets field value.
func (o *ShopifyInput) SetMarket(v *ShopifyMarket) *ShopifyInput {
	o.Market = *v
	return o
}

func (o ShopifyInput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if true {
		toSerialize["metafields"] = o.Metafields
	}
	if true {
		toSerialize["market"] = o.Market
	}
	serialized, err := json.Marshal(toSerialize)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal ShopifyInput: %w", err)
	}

	return serialized, nil
}

func (o ShopifyInput) String() string {
	out := ""
	out += fmt.Sprintf("  metafields=%v\n", o.Metafields)
	out += fmt.Sprintf("  market=%v\n", o.Market)
	return fmt.Sprintf("ShopifyInput {\n%s}", out)
}