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

// FundamentalsTechnicals struct for FundamentalsTechnicals
type FundamentalsTechnicals struct {
	Beta                  *float64 `json:"Beta,omitempty"`
	Var52WeekHigh         *float64 `json:"52WeekHigh,omitempty"`
	Var52WeekLow          *float64 `json:"52WeekLow,omitempty"`
	Var50DayMA            *float64 `json:"50DayMA,omitempty"`
	Var200DayMA           *float64 `json:"200DayMA,omitempty"`
	SharesShort           *int64   `json:"SharesShort,omitempty"`
	SharesShortPriorMonth *int64   `json:"SharesShortPriorMonth,omitempty"`
	ShortRatio            *float64 `json:"ShortRatio,omitempty"`
	ShortPercent          *float64 `json:"ShortPercent,omitempty"`
}

// NewFundamentalsTechnicals instantiates a new FundamentalsTechnicals object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFundamentalsTechnicals() *FundamentalsTechnicals {
	this := FundamentalsTechnicals{}
	return &this
}

// NewFundamentalsTechnicalsWithDefaults instantiates a new FundamentalsTechnicals object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFundamentalsTechnicalsWithDefaults() *FundamentalsTechnicals {
	this := FundamentalsTechnicals{}
	return &this
}

// GetBeta returns the Beta field value if set, zero value otherwise.
func (o *FundamentalsTechnicals) GetBeta() float64 {
	if o == nil || o.Beta == nil {
		var ret float64
		return ret
	}
	return *o.Beta
}

// GetBetaOk returns a tuple with the Beta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FundamentalsTechnicals) GetBetaOk() (*float64, bool) {
	if o == nil || o.Beta == nil {
		return nil, false
	}
	return o.Beta, true
}

// HasBeta returns a boolean if a field has been set.
func (o *FundamentalsTechnicals) HasBeta() bool {
	if o != nil && o.Beta != nil {
		return true
	}

	return false
}

// SetBeta gets a reference to the given float64 and assigns it to the Beta field.
func (o *FundamentalsTechnicals) SetBeta(v float64) {
	o.Beta = &v
}

// GetVar52WeekHigh returns the Var52WeekHigh field value if set, zero value otherwise.
func (o *FundamentalsTechnicals) GetVar52WeekHigh() float64 {
	if o == nil || o.Var52WeekHigh == nil {
		var ret float64
		return ret
	}
	return *o.Var52WeekHigh
}

// GetVar52WeekHighOk returns a tuple with the Var52WeekHigh field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FundamentalsTechnicals) GetVar52WeekHighOk() (*float64, bool) {
	if o == nil || o.Var52WeekHigh == nil {
		return nil, false
	}
	return o.Var52WeekHigh, true
}

// HasVar52WeekHigh returns a boolean if a field has been set.
func (o *FundamentalsTechnicals) HasVar52WeekHigh() bool {
	if o != nil && o.Var52WeekHigh != nil {
		return true
	}

	return false
}

// SetVar52WeekHigh gets a reference to the given float64 and assigns it to the Var52WeekHigh field.
func (o *FundamentalsTechnicals) SetVar52WeekHigh(v float64) {
	o.Var52WeekHigh = &v
}

// GetVar52WeekLow returns the Var52WeekLow field value if set, zero value otherwise.
func (o *FundamentalsTechnicals) GetVar52WeekLow() float64 {
	if o == nil || o.Var52WeekLow == nil {
		var ret float64
		return ret
	}
	return *o.Var52WeekLow
}

// GetVar52WeekLowOk returns a tuple with the Var52WeekLow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FundamentalsTechnicals) GetVar52WeekLowOk() (*float64, bool) {
	if o == nil || o.Var52WeekLow == nil {
		return nil, false
	}
	return o.Var52WeekLow, true
}

// HasVar52WeekLow returns a boolean if a field has been set.
func (o *FundamentalsTechnicals) HasVar52WeekLow() bool {
	if o != nil && o.Var52WeekLow != nil {
		return true
	}

	return false
}

// SetVar52WeekLow gets a reference to the given float64 and assigns it to the Var52WeekLow field.
func (o *FundamentalsTechnicals) SetVar52WeekLow(v float64) {
	o.Var52WeekLow = &v
}

// GetVar50DayMA returns the Var50DayMA field value if set, zero value otherwise.
func (o *FundamentalsTechnicals) GetVar50DayMA() float64 {
	if o == nil || o.Var50DayMA == nil {
		var ret float64
		return ret
	}
	return *o.Var50DayMA
}

// GetVar50DayMAOk returns a tuple with the Var50DayMA field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FundamentalsTechnicals) GetVar50DayMAOk() (*float64, bool) {
	if o == nil || o.Var50DayMA == nil {
		return nil, false
	}
	return o.Var50DayMA, true
}

// HasVar50DayMA returns a boolean if a field has been set.
func (o *FundamentalsTechnicals) HasVar50DayMA() bool {
	if o != nil && o.Var50DayMA != nil {
		return true
	}

	return false
}

// SetVar50DayMA gets a reference to the given float64 and assigns it to the Var50DayMA field.
func (o *FundamentalsTechnicals) SetVar50DayMA(v float64) {
	o.Var50DayMA = &v
}

// GetVar200DayMA returns the Var200DayMA field value if set, zero value otherwise.
func (o *FundamentalsTechnicals) GetVar200DayMA() float64 {
	if o == nil || o.Var200DayMA == nil {
		var ret float64
		return ret
	}
	return *o.Var200DayMA
}

// GetVar200DayMAOk returns a tuple with the Var200DayMA field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FundamentalsTechnicals) GetVar200DayMAOk() (*float64, bool) {
	if o == nil || o.Var200DayMA == nil {
		return nil, false
	}
	return o.Var200DayMA, true
}

// HasVar200DayMA returns a boolean if a field has been set.
func (o *FundamentalsTechnicals) HasVar200DayMA() bool {
	if o != nil && o.Var200DayMA != nil {
		return true
	}

	return false
}

// SetVar200DayMA gets a reference to the given float64 and assigns it to the Var200DayMA field.
func (o *FundamentalsTechnicals) SetVar200DayMA(v float64) {
	o.Var200DayMA = &v
}

// GetSharesShort returns the SharesShort field value if set, zero value otherwise.
func (o *FundamentalsTechnicals) GetSharesShort() int64 {
	if o == nil || o.SharesShort == nil {
		var ret int64
		return ret
	}
	return *o.SharesShort
}

// GetSharesShortOk returns a tuple with the SharesShort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FundamentalsTechnicals) GetSharesShortOk() (*int64, bool) {
	if o == nil || o.SharesShort == nil {
		return nil, false
	}
	return o.SharesShort, true
}

// HasSharesShort returns a boolean if a field has been set.
func (o *FundamentalsTechnicals) HasSharesShort() bool {
	if o != nil && o.SharesShort != nil {
		return true
	}

	return false
}

// SetSharesShort gets a reference to the given int64 and assigns it to the SharesShort field.
func (o *FundamentalsTechnicals) SetSharesShort(v int64) {
	o.SharesShort = &v
}

// GetSharesShortPriorMonth returns the SharesShortPriorMonth field value if set, zero value otherwise.
func (o *FundamentalsTechnicals) GetSharesShortPriorMonth() int64 {
	if o == nil || o.SharesShortPriorMonth == nil {
		var ret int64
		return ret
	}
	return *o.SharesShortPriorMonth
}

// GetSharesShortPriorMonthOk returns a tuple with the SharesShortPriorMonth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FundamentalsTechnicals) GetSharesShortPriorMonthOk() (*int64, bool) {
	if o == nil || o.SharesShortPriorMonth == nil {
		return nil, false
	}
	return o.SharesShortPriorMonth, true
}

// HasSharesShortPriorMonth returns a boolean if a field has been set.
func (o *FundamentalsTechnicals) HasSharesShortPriorMonth() bool {
	if o != nil && o.SharesShortPriorMonth != nil {
		return true
	}

	return false
}

// SetSharesShortPriorMonth gets a reference to the given int64 and assigns it to the SharesShortPriorMonth field.
func (o *FundamentalsTechnicals) SetSharesShortPriorMonth(v int64) {
	o.SharesShortPriorMonth = &v
}

// GetShortRatio returns the ShortRatio field value if set, zero value otherwise.
func (o *FundamentalsTechnicals) GetShortRatio() float64 {
	if o == nil || o.ShortRatio == nil {
		var ret float64
		return ret
	}
	return *o.ShortRatio
}

// GetShortRatioOk returns a tuple with the ShortRatio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FundamentalsTechnicals) GetShortRatioOk() (*float64, bool) {
	if o == nil || o.ShortRatio == nil {
		return nil, false
	}
	return o.ShortRatio, true
}

// HasShortRatio returns a boolean if a field has been set.
func (o *FundamentalsTechnicals) HasShortRatio() bool {
	if o != nil && o.ShortRatio != nil {
		return true
	}

	return false
}

// SetShortRatio gets a reference to the given float64 and assigns it to the ShortRatio field.
func (o *FundamentalsTechnicals) SetShortRatio(v float64) {
	o.ShortRatio = &v
}

// GetShortPercent returns the ShortPercent field value if set, zero value otherwise.
func (o *FundamentalsTechnicals) GetShortPercent() float64 {
	if o == nil || o.ShortPercent == nil {
		var ret float64
		return ret
	}
	return *o.ShortPercent
}

// GetShortPercentOk returns a tuple with the ShortPercent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FundamentalsTechnicals) GetShortPercentOk() (*float64, bool) {
	if o == nil || o.ShortPercent == nil {
		return nil, false
	}
	return o.ShortPercent, true
}

// HasShortPercent returns a boolean if a field has been set.
func (o *FundamentalsTechnicals) HasShortPercent() bool {
	if o != nil && o.ShortPercent != nil {
		return true
	}

	return false
}

// SetShortPercent gets a reference to the given float64 and assigns it to the ShortPercent field.
func (o *FundamentalsTechnicals) SetShortPercent(v float64) {
	o.ShortPercent = &v
}

func (o FundamentalsTechnicals) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Beta != nil {
		toSerialize["Beta"] = o.Beta
	}
	if o.Var52WeekHigh != nil {
		toSerialize["52WeekHigh"] = o.Var52WeekHigh
	}
	if o.Var52WeekLow != nil {
		toSerialize["52WeekLow"] = o.Var52WeekLow
	}
	if o.Var50DayMA != nil {
		toSerialize["50DayMA"] = o.Var50DayMA
	}
	if o.Var200DayMA != nil {
		toSerialize["200DayMA"] = o.Var200DayMA
	}
	if o.SharesShort != nil {
		toSerialize["SharesShort"] = o.SharesShort
	}
	if o.SharesShortPriorMonth != nil {
		toSerialize["SharesShortPriorMonth"] = o.SharesShortPriorMonth
	}
	if o.ShortRatio != nil {
		toSerialize["ShortRatio"] = o.ShortRatio
	}
	if o.ShortPercent != nil {
		toSerialize["ShortPercent"] = o.ShortPercent
	}
	return json.Marshal(toSerialize)
}

type NullableFundamentalsTechnicals struct {
	value *FundamentalsTechnicals
	isSet bool
}

func (v NullableFundamentalsTechnicals) Get() *FundamentalsTechnicals {
	return v.value
}

func (v *NullableFundamentalsTechnicals) Set(val *FundamentalsTechnicals) {
	v.value = val
	v.isSet = true
}

func (v NullableFundamentalsTechnicals) IsSet() bool {
	return v.isSet
}

func (v *NullableFundamentalsTechnicals) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFundamentalsTechnicals(val *FundamentalsTechnicals) *NullableFundamentalsTechnicals {
	return &NullableFundamentalsTechnicals{value: val, isSet: true}
}

func (v NullableFundamentalsTechnicals) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFundamentalsTechnicals) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
