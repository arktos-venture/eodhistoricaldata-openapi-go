/*
eodhistoricaldata

eodhistoricaldata API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// FundamentalsAnalystRatings struct for FundamentalsAnalystRatings
type FundamentalsAnalystRatings struct {
	Rating      *float64 `json:"Rating,omitempty"`
	TargetPrice *float64 `json:"TargetPrice,omitempty"`
	StrongBuy   *int64   `json:"StrongBuy,omitempty"`
	Buy         *int64   `json:"Buy,omitempty"`
	Hold        *int64   `json:"Hold,omitempty"`
	Sell        *int64   `json:"Sell,omitempty"`
	StrongSell  *int64   `json:"StrongSell,omitempty"`
}

// NewFundamentalsAnalystRatings instantiates a new FundamentalsAnalystRatings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFundamentalsAnalystRatings() *FundamentalsAnalystRatings {
	this := FundamentalsAnalystRatings{}
	return &this
}

// NewFundamentalsAnalystRatingsWithDefaults instantiates a new FundamentalsAnalystRatings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFundamentalsAnalystRatingsWithDefaults() *FundamentalsAnalystRatings {
	this := FundamentalsAnalystRatings{}
	return &this
}

// GetRating returns the Rating field value if set, zero value otherwise.
func (o *FundamentalsAnalystRatings) GetRating() float64 {
	if o == nil || o.Rating == nil {
		var ret float64
		return ret
	}
	return *o.Rating
}

// GetRatingOk returns a tuple with the Rating field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FundamentalsAnalystRatings) GetRatingOk() (*float64, bool) {
	if o == nil || o.Rating == nil {
		return nil, false
	}
	return o.Rating, true
}

// HasRating returns a boolean if a field has been set.
func (o *FundamentalsAnalystRatings) HasRating() bool {
	if o != nil && o.Rating != nil {
		return true
	}

	return false
}

// SetRating gets a reference to the given float64 and assigns it to the Rating field.
func (o *FundamentalsAnalystRatings) SetRating(v float64) {
	o.Rating = &v
}

// GetTargetPrice returns the TargetPrice field value if set, zero value otherwise.
func (o *FundamentalsAnalystRatings) GetTargetPrice() float64 {
	if o == nil || o.TargetPrice == nil {
		var ret float64
		return ret
	}
	return *o.TargetPrice
}

// GetTargetPriceOk returns a tuple with the TargetPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FundamentalsAnalystRatings) GetTargetPriceOk() (*float64, bool) {
	if o == nil || o.TargetPrice == nil {
		return nil, false
	}
	return o.TargetPrice, true
}

// HasTargetPrice returns a boolean if a field has been set.
func (o *FundamentalsAnalystRatings) HasTargetPrice() bool {
	if o != nil && o.TargetPrice != nil {
		return true
	}

	return false
}

// SetTargetPrice gets a reference to the given float64 and assigns it to the TargetPrice field.
func (o *FundamentalsAnalystRatings) SetTargetPrice(v float64) {
	o.TargetPrice = &v
}

// GetStrongBuy returns the StrongBuy field value if set, zero value otherwise.
func (o *FundamentalsAnalystRatings) GetStrongBuy() int64 {
	if o == nil || o.StrongBuy == nil {
		var ret int64
		return ret
	}
	return *o.StrongBuy
}

// GetStrongBuyOk returns a tuple with the StrongBuy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FundamentalsAnalystRatings) GetStrongBuyOk() (*int64, bool) {
	if o == nil || o.StrongBuy == nil {
		return nil, false
	}
	return o.StrongBuy, true
}

// HasStrongBuy returns a boolean if a field has been set.
func (o *FundamentalsAnalystRatings) HasStrongBuy() bool {
	if o != nil && o.StrongBuy != nil {
		return true
	}

	return false
}

// SetStrongBuy gets a reference to the given int64 and assigns it to the StrongBuy field.
func (o *FundamentalsAnalystRatings) SetStrongBuy(v int64) {
	o.StrongBuy = &v
}

// GetBuy returns the Buy field value if set, zero value otherwise.
func (o *FundamentalsAnalystRatings) GetBuy() int64 {
	if o == nil || o.Buy == nil {
		var ret int64
		return ret
	}
	return *o.Buy
}

// GetBuyOk returns a tuple with the Buy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FundamentalsAnalystRatings) GetBuyOk() (*int64, bool) {
	if o == nil || o.Buy == nil {
		return nil, false
	}
	return o.Buy, true
}

// HasBuy returns a boolean if a field has been set.
func (o *FundamentalsAnalystRatings) HasBuy() bool {
	if o != nil && o.Buy != nil {
		return true
	}

	return false
}

// SetBuy gets a reference to the given int64 and assigns it to the Buy field.
func (o *FundamentalsAnalystRatings) SetBuy(v int64) {
	o.Buy = &v
}

// GetHold returns the Hold field value if set, zero value otherwise.
func (o *FundamentalsAnalystRatings) GetHold() int64 {
	if o == nil || o.Hold == nil {
		var ret int64
		return ret
	}
	return *o.Hold
}

// GetHoldOk returns a tuple with the Hold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FundamentalsAnalystRatings) GetHoldOk() (*int64, bool) {
	if o == nil || o.Hold == nil {
		return nil, false
	}
	return o.Hold, true
}

// HasHold returns a boolean if a field has been set.
func (o *FundamentalsAnalystRatings) HasHold() bool {
	if o != nil && o.Hold != nil {
		return true
	}

	return false
}

// SetHold gets a reference to the given int64 and assigns it to the Hold field.
func (o *FundamentalsAnalystRatings) SetHold(v int64) {
	o.Hold = &v
}

// GetSell returns the Sell field value if set, zero value otherwise.
func (o *FundamentalsAnalystRatings) GetSell() int64 {
	if o == nil || o.Sell == nil {
		var ret int64
		return ret
	}
	return *o.Sell
}

// GetSellOk returns a tuple with the Sell field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FundamentalsAnalystRatings) GetSellOk() (*int64, bool) {
	if o == nil || o.Sell == nil {
		return nil, false
	}
	return o.Sell, true
}

// HasSell returns a boolean if a field has been set.
func (o *FundamentalsAnalystRatings) HasSell() bool {
	if o != nil && o.Sell != nil {
		return true
	}

	return false
}

// SetSell gets a reference to the given int64 and assigns it to the Sell field.
func (o *FundamentalsAnalystRatings) SetSell(v int64) {
	o.Sell = &v
}

// GetStrongSell returns the StrongSell field value if set, zero value otherwise.
func (o *FundamentalsAnalystRatings) GetStrongSell() int64 {
	if o == nil || o.StrongSell == nil {
		var ret int64
		return ret
	}
	return *o.StrongSell
}

// GetStrongSellOk returns a tuple with the StrongSell field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FundamentalsAnalystRatings) GetStrongSellOk() (*int64, bool) {
	if o == nil || o.StrongSell == nil {
		return nil, false
	}
	return o.StrongSell, true
}

// HasStrongSell returns a boolean if a field has been set.
func (o *FundamentalsAnalystRatings) HasStrongSell() bool {
	if o != nil && o.StrongSell != nil {
		return true
	}

	return false
}

// SetStrongSell gets a reference to the given int64 and assigns it to the StrongSell field.
func (o *FundamentalsAnalystRatings) SetStrongSell(v int64) {
	o.StrongSell = &v
}

func (o FundamentalsAnalystRatings) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Rating != nil {
		toSerialize["Rating"] = o.Rating
	}
	if o.TargetPrice != nil {
		toSerialize["TargetPrice"] = o.TargetPrice
	}
	if o.StrongBuy != nil {
		toSerialize["StrongBuy"] = o.StrongBuy
	}
	if o.Buy != nil {
		toSerialize["Buy"] = o.Buy
	}
	if o.Hold != nil {
		toSerialize["Hold"] = o.Hold
	}
	if o.Sell != nil {
		toSerialize["Sell"] = o.Sell
	}
	if o.StrongSell != nil {
		toSerialize["StrongSell"] = o.StrongSell
	}
	return json.Marshal(toSerialize)
}

type NullableFundamentalsAnalystRatings struct {
	value *FundamentalsAnalystRatings
	isSet bool
}

func (v NullableFundamentalsAnalystRatings) Get() *FundamentalsAnalystRatings {
	return v.value
}

func (v *NullableFundamentalsAnalystRatings) Set(val *FundamentalsAnalystRatings) {
	v.value = val
	v.isSet = true
}

func (v NullableFundamentalsAnalystRatings) IsSet() bool {
	return v.isSet
}

func (v *NullableFundamentalsAnalystRatings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFundamentalsAnalystRatings(val *FundamentalsAnalystRatings) *NullableFundamentalsAnalystRatings {
	return &NullableFundamentalsAnalystRatings{value: val, isSet: true}
}

func (v NullableFundamentalsAnalystRatings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFundamentalsAnalystRatings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
