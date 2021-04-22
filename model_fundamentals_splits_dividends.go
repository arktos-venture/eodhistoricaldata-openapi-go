/*
 * eodhistoricaldata
 *
 * eodhistoricaldata API
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// FundamentalsSplitsDividends struct for FundamentalsSplitsDividends
type FundamentalsSplitsDividends struct {
	ForwardAnnualDividendRate  *float64                                            `json:"ForwardAnnualDividendRate,omitempty"`
	ForwardAnnualDividendYield *float64                                            `json:"ForwardAnnualDividendYield,omitempty"`
	PayoutRatio                *float64                                            `json:"PayoutRatio,omitempty"`
	DividendDate               *string                                             `json:"DividendDate,omitempty"`
	ExDividendDate             *string                                             `json:"ExDividendDate,omitempty"`
	LastSplitFactor            *string                                             `json:"LastSplitFactor,omitempty"`
	LastSplitDate              *string                                             `json:"LastSplitDate,omitempty"`
	NumberDividendsByYear      *[]FundamentalsSplitsDividendsNumberDividendsByYear `json:"NumberDividendsByYear,omitempty"`
}

// NewFundamentalsSplitsDividends instantiates a new FundamentalsSplitsDividends object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFundamentalsSplitsDividends() *FundamentalsSplitsDividends {
	this := FundamentalsSplitsDividends{}
	return &this
}

// NewFundamentalsSplitsDividendsWithDefaults instantiates a new FundamentalsSplitsDividends object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFundamentalsSplitsDividendsWithDefaults() *FundamentalsSplitsDividends {
	this := FundamentalsSplitsDividends{}
	return &this
}

// GetForwardAnnualDividendRate returns the ForwardAnnualDividendRate field value if set, zero value otherwise.
func (o *FundamentalsSplitsDividends) GetForwardAnnualDividendRate() float64 {
	if o == nil || o.ForwardAnnualDividendRate == nil {
		var ret float64
		return ret
	}
	return *o.ForwardAnnualDividendRate
}

// GetForwardAnnualDividendRateOk returns a tuple with the ForwardAnnualDividendRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FundamentalsSplitsDividends) GetForwardAnnualDividendRateOk() (*float64, bool) {
	if o == nil || o.ForwardAnnualDividendRate == nil {
		return nil, false
	}
	return o.ForwardAnnualDividendRate, true
}

// HasForwardAnnualDividendRate returns a boolean if a field has been set.
func (o *FundamentalsSplitsDividends) HasForwardAnnualDividendRate() bool {
	if o != nil && o.ForwardAnnualDividendRate != nil {
		return true
	}

	return false
}

// SetForwardAnnualDividendRate gets a reference to the given float64 and assigns it to the ForwardAnnualDividendRate field.
func (o *FundamentalsSplitsDividends) SetForwardAnnualDividendRate(v float64) {
	o.ForwardAnnualDividendRate = &v
}

// GetForwardAnnualDividendYield returns the ForwardAnnualDividendYield field value if set, zero value otherwise.
func (o *FundamentalsSplitsDividends) GetForwardAnnualDividendYield() float64 {
	if o == nil || o.ForwardAnnualDividendYield == nil {
		var ret float64
		return ret
	}
	return *o.ForwardAnnualDividendYield
}

// GetForwardAnnualDividendYieldOk returns a tuple with the ForwardAnnualDividendYield field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FundamentalsSplitsDividends) GetForwardAnnualDividendYieldOk() (*float64, bool) {
	if o == nil || o.ForwardAnnualDividendYield == nil {
		return nil, false
	}
	return o.ForwardAnnualDividendYield, true
}

// HasForwardAnnualDividendYield returns a boolean if a field has been set.
func (o *FundamentalsSplitsDividends) HasForwardAnnualDividendYield() bool {
	if o != nil && o.ForwardAnnualDividendYield != nil {
		return true
	}

	return false
}

// SetForwardAnnualDividendYield gets a reference to the given float64 and assigns it to the ForwardAnnualDividendYield field.
func (o *FundamentalsSplitsDividends) SetForwardAnnualDividendYield(v float64) {
	o.ForwardAnnualDividendYield = &v
}

// GetPayoutRatio returns the PayoutRatio field value if set, zero value otherwise.
func (o *FundamentalsSplitsDividends) GetPayoutRatio() float64 {
	if o == nil || o.PayoutRatio == nil {
		var ret float64
		return ret
	}
	return *o.PayoutRatio
}

// GetPayoutRatioOk returns a tuple with the PayoutRatio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FundamentalsSplitsDividends) GetPayoutRatioOk() (*float64, bool) {
	if o == nil || o.PayoutRatio == nil {
		return nil, false
	}
	return o.PayoutRatio, true
}

// HasPayoutRatio returns a boolean if a field has been set.
func (o *FundamentalsSplitsDividends) HasPayoutRatio() bool {
	if o != nil && o.PayoutRatio != nil {
		return true
	}

	return false
}

// SetPayoutRatio gets a reference to the given float64 and assigns it to the PayoutRatio field.
func (o *FundamentalsSplitsDividends) SetPayoutRatio(v float64) {
	o.PayoutRatio = &v
}

// GetDividendDate returns the DividendDate field value if set, zero value otherwise.
func (o *FundamentalsSplitsDividends) GetDividendDate() string {
	if o == nil || o.DividendDate == nil {
		var ret string
		return ret
	}
	return *o.DividendDate
}

// GetDividendDateOk returns a tuple with the DividendDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FundamentalsSplitsDividends) GetDividendDateOk() (*string, bool) {
	if o == nil || o.DividendDate == nil {
		return nil, false
	}
	return o.DividendDate, true
}

// HasDividendDate returns a boolean if a field has been set.
func (o *FundamentalsSplitsDividends) HasDividendDate() bool {
	if o != nil && o.DividendDate != nil {
		return true
	}

	return false
}

// SetDividendDate gets a reference to the given string and assigns it to the DividendDate field.
func (o *FundamentalsSplitsDividends) SetDividendDate(v string) {
	o.DividendDate = &v
}

// GetExDividendDate returns the ExDividendDate field value if set, zero value otherwise.
func (o *FundamentalsSplitsDividends) GetExDividendDate() string {
	if o == nil || o.ExDividendDate == nil {
		var ret string
		return ret
	}
	return *o.ExDividendDate
}

// GetExDividendDateOk returns a tuple with the ExDividendDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FundamentalsSplitsDividends) GetExDividendDateOk() (*string, bool) {
	if o == nil || o.ExDividendDate == nil {
		return nil, false
	}
	return o.ExDividendDate, true
}

// HasExDividendDate returns a boolean if a field has been set.
func (o *FundamentalsSplitsDividends) HasExDividendDate() bool {
	if o != nil && o.ExDividendDate != nil {
		return true
	}

	return false
}

// SetExDividendDate gets a reference to the given string and assigns it to the ExDividendDate field.
func (o *FundamentalsSplitsDividends) SetExDividendDate(v string) {
	o.ExDividendDate = &v
}

// GetLastSplitFactor returns the LastSplitFactor field value if set, zero value otherwise.
func (o *FundamentalsSplitsDividends) GetLastSplitFactor() string {
	if o == nil || o.LastSplitFactor == nil {
		var ret string
		return ret
	}
	return *o.LastSplitFactor
}

// GetLastSplitFactorOk returns a tuple with the LastSplitFactor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FundamentalsSplitsDividends) GetLastSplitFactorOk() (*string, bool) {
	if o == nil || o.LastSplitFactor == nil {
		return nil, false
	}
	return o.LastSplitFactor, true
}

// HasLastSplitFactor returns a boolean if a field has been set.
func (o *FundamentalsSplitsDividends) HasLastSplitFactor() bool {
	if o != nil && o.LastSplitFactor != nil {
		return true
	}

	return false
}

// SetLastSplitFactor gets a reference to the given string and assigns it to the LastSplitFactor field.
func (o *FundamentalsSplitsDividends) SetLastSplitFactor(v string) {
	o.LastSplitFactor = &v
}

// GetLastSplitDate returns the LastSplitDate field value if set, zero value otherwise.
func (o *FundamentalsSplitsDividends) GetLastSplitDate() string {
	if o == nil || o.LastSplitDate == nil {
		var ret string
		return ret
	}
	return *o.LastSplitDate
}

// GetLastSplitDateOk returns a tuple with the LastSplitDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FundamentalsSplitsDividends) GetLastSplitDateOk() (*string, bool) {
	if o == nil || o.LastSplitDate == nil {
		return nil, false
	}
	return o.LastSplitDate, true
}

// HasLastSplitDate returns a boolean if a field has been set.
func (o *FundamentalsSplitsDividends) HasLastSplitDate() bool {
	if o != nil && o.LastSplitDate != nil {
		return true
	}

	return false
}

// SetLastSplitDate gets a reference to the given string and assigns it to the LastSplitDate field.
func (o *FundamentalsSplitsDividends) SetLastSplitDate(v string) {
	o.LastSplitDate = &v
}

// GetNumberDividendsByYear returns the NumberDividendsByYear field value if set, zero value otherwise.
func (o *FundamentalsSplitsDividends) GetNumberDividendsByYear() []FundamentalsSplitsDividendsNumberDividendsByYear {
	if o == nil || o.NumberDividendsByYear == nil {
		var ret []FundamentalsSplitsDividendsNumberDividendsByYear
		return ret
	}
	return *o.NumberDividendsByYear
}

// GetNumberDividendsByYearOk returns a tuple with the NumberDividendsByYear field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FundamentalsSplitsDividends) GetNumberDividendsByYearOk() (*[]FundamentalsSplitsDividendsNumberDividendsByYear, bool) {
	if o == nil || o.NumberDividendsByYear == nil {
		return nil, false
	}
	return o.NumberDividendsByYear, true
}

// HasNumberDividendsByYear returns a boolean if a field has been set.
func (o *FundamentalsSplitsDividends) HasNumberDividendsByYear() bool {
	if o != nil && o.NumberDividendsByYear != nil {
		return true
	}

	return false
}

// SetNumberDividendsByYear gets a reference to the given []FundamentalsSplitsDividendsNumberDividendsByYear and assigns it to the NumberDividendsByYear field.
func (o *FundamentalsSplitsDividends) SetNumberDividendsByYear(v []FundamentalsSplitsDividendsNumberDividendsByYear) {
	o.NumberDividendsByYear = &v
}

func (o FundamentalsSplitsDividends) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ForwardAnnualDividendRate != nil {
		toSerialize["ForwardAnnualDividendRate"] = o.ForwardAnnualDividendRate
	}
	if o.ForwardAnnualDividendYield != nil {
		toSerialize["ForwardAnnualDividendYield"] = o.ForwardAnnualDividendYield
	}
	if o.PayoutRatio != nil {
		toSerialize["PayoutRatio"] = o.PayoutRatio
	}
	if o.DividendDate != nil {
		toSerialize["DividendDate"] = o.DividendDate
	}
	if o.ExDividendDate != nil {
		toSerialize["ExDividendDate"] = o.ExDividendDate
	}
	if o.LastSplitFactor != nil {
		toSerialize["LastSplitFactor"] = o.LastSplitFactor
	}
	if o.LastSplitDate != nil {
		toSerialize["LastSplitDate"] = o.LastSplitDate
	}
	if o.NumberDividendsByYear != nil {
		toSerialize["NumberDividendsByYear"] = o.NumberDividendsByYear
	}
	return json.Marshal(toSerialize)
}

type NullableFundamentalsSplitsDividends struct {
	value *FundamentalsSplitsDividends
	isSet bool
}

func (v NullableFundamentalsSplitsDividends) Get() *FundamentalsSplitsDividends {
	return v.value
}

func (v *NullableFundamentalsSplitsDividends) Set(val *FundamentalsSplitsDividends) {
	v.value = val
	v.isSet = true
}

func (v NullableFundamentalsSplitsDividends) IsSet() bool {
	return v.isSet
}

func (v *NullableFundamentalsSplitsDividends) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFundamentalsSplitsDividends(val *FundamentalsSplitsDividends) *NullableFundamentalsSplitsDividends {
	return &NullableFundamentalsSplitsDividends{value: val, isSet: true}
}

func (v NullableFundamentalsSplitsDividends) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFundamentalsSplitsDividends) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}