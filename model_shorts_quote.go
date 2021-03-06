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

// ShortsQuote struct for ShortsQuote
type ShortsQuote struct {
	Date   *string  `json:"date,omitempty"`
	Short  *int64   `json:"short,omitempty"`
	Volume *float64 `json:"volume,omitempty"`
}

// NewShortsQuote instantiates a new ShortsQuote object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShortsQuote() *ShortsQuote {
	this := ShortsQuote{}
	return &this
}

// NewShortsQuoteWithDefaults instantiates a new ShortsQuote object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShortsQuoteWithDefaults() *ShortsQuote {
	this := ShortsQuote{}
	return &this
}

// GetDate returns the Date field value if set, zero value otherwise.
func (o *ShortsQuote) GetDate() string {
	if o == nil || o.Date == nil {
		var ret string
		return ret
	}
	return *o.Date
}

// GetDateOk returns a tuple with the Date field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShortsQuote) GetDateOk() (*string, bool) {
	if o == nil || o.Date == nil {
		return nil, false
	}
	return o.Date, true
}

// HasDate returns a boolean if a field has been set.
func (o *ShortsQuote) HasDate() bool {
	if o != nil && o.Date != nil {
		return true
	}

	return false
}

// SetDate gets a reference to the given string and assigns it to the Date field.
func (o *ShortsQuote) SetDate(v string) {
	o.Date = &v
}

// GetShort returns the Short field value if set, zero value otherwise.
func (o *ShortsQuote) GetShort() int64 {
	if o == nil || o.Short == nil {
		var ret int64
		return ret
	}
	return *o.Short
}

// GetShortOk returns a tuple with the Short field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShortsQuote) GetShortOk() (*int64, bool) {
	if o == nil || o.Short == nil {
		return nil, false
	}
	return o.Short, true
}

// HasShort returns a boolean if a field has been set.
func (o *ShortsQuote) HasShort() bool {
	if o != nil && o.Short != nil {
		return true
	}

	return false
}

// SetShort gets a reference to the given int64 and assigns it to the Short field.
func (o *ShortsQuote) SetShort(v int64) {
	o.Short = &v
}

// GetVolume returns the Volume field value if set, zero value otherwise.
func (o *ShortsQuote) GetVolume() float64 {
	if o == nil || o.Volume == nil {
		var ret float64
		return ret
	}
	return *o.Volume
}

// GetVolumeOk returns a tuple with the Volume field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShortsQuote) GetVolumeOk() (*float64, bool) {
	if o == nil || o.Volume == nil {
		return nil, false
	}
	return o.Volume, true
}

// HasVolume returns a boolean if a field has been set.
func (o *ShortsQuote) HasVolume() bool {
	if o != nil && o.Volume != nil {
		return true
	}

	return false
}

// SetVolume gets a reference to the given float64 and assigns it to the Volume field.
func (o *ShortsQuote) SetVolume(v float64) {
	o.Volume = &v
}

func (o ShortsQuote) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Date != nil {
		toSerialize["date"] = o.Date
	}
	if o.Short != nil {
		toSerialize["short"] = o.Short
	}
	if o.Volume != nil {
		toSerialize["volume"] = o.Volume
	}
	return json.Marshal(toSerialize)
}

type NullableShortsQuote struct {
	value *ShortsQuote
	isSet bool
}

func (v NullableShortsQuote) Get() *ShortsQuote {
	return v.value
}

func (v *NullableShortsQuote) Set(val *ShortsQuote) {
	v.value = val
	v.isSet = true
}

func (v NullableShortsQuote) IsSet() bool {
	return v.isSet
}

func (v *NullableShortsQuote) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShortsQuote(val *ShortsQuote) *NullableShortsQuote {
	return &NullableShortsQuote{value: val, isSet: true}
}

func (v NullableShortsQuote) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableShortsQuote) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
