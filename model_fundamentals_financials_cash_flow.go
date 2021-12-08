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

// FundamentalsFinancialsCashFlow struct for FundamentalsFinancialsCashFlow
type FundamentalsFinancialsCashFlow struct {
	CurrencySymbol *string              `json:"currency_symbol,omitempty"`
	Quarterly      *map[string]CashFlow `json:"quarterly,omitempty"`
	Yearly         *map[string]CashFlow `json:"yearly,omitempty"`
}

// NewFundamentalsFinancialsCashFlow instantiates a new FundamentalsFinancialsCashFlow object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFundamentalsFinancialsCashFlow() *FundamentalsFinancialsCashFlow {
	this := FundamentalsFinancialsCashFlow{}
	return &this
}

// NewFundamentalsFinancialsCashFlowWithDefaults instantiates a new FundamentalsFinancialsCashFlow object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFundamentalsFinancialsCashFlowWithDefaults() *FundamentalsFinancialsCashFlow {
	this := FundamentalsFinancialsCashFlow{}
	return &this
}

// GetCurrencySymbol returns the CurrencySymbol field value if set, zero value otherwise.
func (o *FundamentalsFinancialsCashFlow) GetCurrencySymbol() string {
	if o == nil || o.CurrencySymbol == nil {
		var ret string
		return ret
	}
	return *o.CurrencySymbol
}

// GetCurrencySymbolOk returns a tuple with the CurrencySymbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FundamentalsFinancialsCashFlow) GetCurrencySymbolOk() (*string, bool) {
	if o == nil || o.CurrencySymbol == nil {
		return nil, false
	}
	return o.CurrencySymbol, true
}

// HasCurrencySymbol returns a boolean if a field has been set.
func (o *FundamentalsFinancialsCashFlow) HasCurrencySymbol() bool {
	if o != nil && o.CurrencySymbol != nil {
		return true
	}

	return false
}

// SetCurrencySymbol gets a reference to the given string and assigns it to the CurrencySymbol field.
func (o *FundamentalsFinancialsCashFlow) SetCurrencySymbol(v string) {
	o.CurrencySymbol = &v
}

// GetQuarterly returns the Quarterly field value if set, zero value otherwise.
func (o *FundamentalsFinancialsCashFlow) GetQuarterly() map[string]CashFlow {
	if o == nil || o.Quarterly == nil {
		var ret map[string]CashFlow
		return ret
	}
	return *o.Quarterly
}

// GetQuarterlyOk returns a tuple with the Quarterly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FundamentalsFinancialsCashFlow) GetQuarterlyOk() (*map[string]CashFlow, bool) {
	if o == nil || o.Quarterly == nil {
		return nil, false
	}
	return o.Quarterly, true
}

// HasQuarterly returns a boolean if a field has been set.
func (o *FundamentalsFinancialsCashFlow) HasQuarterly() bool {
	if o != nil && o.Quarterly != nil {
		return true
	}

	return false
}

// SetQuarterly gets a reference to the given map[string]CashFlow and assigns it to the Quarterly field.
func (o *FundamentalsFinancialsCashFlow) SetQuarterly(v map[string]CashFlow) {
	o.Quarterly = &v
}

// GetYearly returns the Yearly field value if set, zero value otherwise.
func (o *FundamentalsFinancialsCashFlow) GetYearly() map[string]CashFlow {
	if o == nil || o.Yearly == nil {
		var ret map[string]CashFlow
		return ret
	}
	return *o.Yearly
}

// GetYearlyOk returns a tuple with the Yearly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FundamentalsFinancialsCashFlow) GetYearlyOk() (*map[string]CashFlow, bool) {
	if o == nil || o.Yearly == nil {
		return nil, false
	}
	return o.Yearly, true
}

// HasYearly returns a boolean if a field has been set.
func (o *FundamentalsFinancialsCashFlow) HasYearly() bool {
	if o != nil && o.Yearly != nil {
		return true
	}

	return false
}

// SetYearly gets a reference to the given map[string]CashFlow and assigns it to the Yearly field.
func (o *FundamentalsFinancialsCashFlow) SetYearly(v map[string]CashFlow) {
	o.Yearly = &v
}

func (o FundamentalsFinancialsCashFlow) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CurrencySymbol != nil {
		toSerialize["currency_symbol"] = o.CurrencySymbol
	}
	if o.Quarterly != nil {
		toSerialize["quarterly"] = o.Quarterly
	}
	if o.Yearly != nil {
		toSerialize["yearly"] = o.Yearly
	}
	return json.Marshal(toSerialize)
}

type NullableFundamentalsFinancialsCashFlow struct {
	value *FundamentalsFinancialsCashFlow
	isSet bool
}

func (v NullableFundamentalsFinancialsCashFlow) Get() *FundamentalsFinancialsCashFlow {
	return v.value
}

func (v *NullableFundamentalsFinancialsCashFlow) Set(val *FundamentalsFinancialsCashFlow) {
	v.value = val
	v.isSet = true
}

func (v NullableFundamentalsFinancialsCashFlow) IsSet() bool {
	return v.isSet
}

func (v *NullableFundamentalsFinancialsCashFlow) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFundamentalsFinancialsCashFlow(val *FundamentalsFinancialsCashFlow) *NullableFundamentalsFinancialsCashFlow {
	return &NullableFundamentalsFinancialsCashFlow{value: val, isSet: true}
}

func (v NullableFundamentalsFinancialsCashFlow) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFundamentalsFinancialsCashFlow) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}