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

// FundamentalsEarningsHistory struct for FundamentalsEarningsHistory
type FundamentalsEarningsHistory struct {
	ReportDate        *string  `json:"reportDate,omitempty"`
	Date              *string  `json:"date,omitempty"`
	BeforeAfterMarket *string  `json:"beforeAfterMarket,omitempty"`
	Currency          *string  `json:"currency,omitempty"`
	EpsActual         *float64 `json:"epsActual,omitempty"`
	EpsEstimate       *float64 `json:"epsEstimate,omitempty"`
	EpsDifference     *float64 `json:"epsDifference,omitempty"`
	SurprisePercent   *float64 `json:"surprisePercent,omitempty"`
}

// NewFundamentalsEarningsHistory instantiates a new FundamentalsEarningsHistory object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFundamentalsEarningsHistory() *FundamentalsEarningsHistory {
	this := FundamentalsEarningsHistory{}
	return &this
}

// NewFundamentalsEarningsHistoryWithDefaults instantiates a new FundamentalsEarningsHistory object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFundamentalsEarningsHistoryWithDefaults() *FundamentalsEarningsHistory {
	this := FundamentalsEarningsHistory{}
	return &this
}

// GetReportDate returns the ReportDate field value if set, zero value otherwise.
func (o *FundamentalsEarningsHistory) GetReportDate() string {
	if o == nil || o.ReportDate == nil {
		var ret string
		return ret
	}
	return *o.ReportDate
}

// GetReportDateOk returns a tuple with the ReportDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FundamentalsEarningsHistory) GetReportDateOk() (*string, bool) {
	if o == nil || o.ReportDate == nil {
		return nil, false
	}
	return o.ReportDate, true
}

// HasReportDate returns a boolean if a field has been set.
func (o *FundamentalsEarningsHistory) HasReportDate() bool {
	if o != nil && o.ReportDate != nil {
		return true
	}

	return false
}

// SetReportDate gets a reference to the given string and assigns it to the ReportDate field.
func (o *FundamentalsEarningsHistory) SetReportDate(v string) {
	o.ReportDate = &v
}

// GetDate returns the Date field value if set, zero value otherwise.
func (o *FundamentalsEarningsHistory) GetDate() string {
	if o == nil || o.Date == nil {
		var ret string
		return ret
	}
	return *o.Date
}

// GetDateOk returns a tuple with the Date field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FundamentalsEarningsHistory) GetDateOk() (*string, bool) {
	if o == nil || o.Date == nil {
		return nil, false
	}
	return o.Date, true
}

// HasDate returns a boolean if a field has been set.
func (o *FundamentalsEarningsHistory) HasDate() bool {
	if o != nil && o.Date != nil {
		return true
	}

	return false
}

// SetDate gets a reference to the given string and assigns it to the Date field.
func (o *FundamentalsEarningsHistory) SetDate(v string) {
	o.Date = &v
}

// GetBeforeAfterMarket returns the BeforeAfterMarket field value if set, zero value otherwise.
func (o *FundamentalsEarningsHistory) GetBeforeAfterMarket() string {
	if o == nil || o.BeforeAfterMarket == nil {
		var ret string
		return ret
	}
	return *o.BeforeAfterMarket
}

// GetBeforeAfterMarketOk returns a tuple with the BeforeAfterMarket field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FundamentalsEarningsHistory) GetBeforeAfterMarketOk() (*string, bool) {
	if o == nil || o.BeforeAfterMarket == nil {
		return nil, false
	}
	return o.BeforeAfterMarket, true
}

// HasBeforeAfterMarket returns a boolean if a field has been set.
func (o *FundamentalsEarningsHistory) HasBeforeAfterMarket() bool {
	if o != nil && o.BeforeAfterMarket != nil {
		return true
	}

	return false
}

// SetBeforeAfterMarket gets a reference to the given string and assigns it to the BeforeAfterMarket field.
func (o *FundamentalsEarningsHistory) SetBeforeAfterMarket(v string) {
	o.BeforeAfterMarket = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *FundamentalsEarningsHistory) GetCurrency() string {
	if o == nil || o.Currency == nil {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FundamentalsEarningsHistory) GetCurrencyOk() (*string, bool) {
	if o == nil || o.Currency == nil {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *FundamentalsEarningsHistory) HasCurrency() bool {
	if o != nil && o.Currency != nil {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *FundamentalsEarningsHistory) SetCurrency(v string) {
	o.Currency = &v
}

// GetEpsActual returns the EpsActual field value if set, zero value otherwise.
func (o *FundamentalsEarningsHistory) GetEpsActual() float64 {
	if o == nil || o.EpsActual == nil {
		var ret float64
		return ret
	}
	return *o.EpsActual
}

// GetEpsActualOk returns a tuple with the EpsActual field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FundamentalsEarningsHistory) GetEpsActualOk() (*float64, bool) {
	if o == nil || o.EpsActual == nil {
		return nil, false
	}
	return o.EpsActual, true
}

// HasEpsActual returns a boolean if a field has been set.
func (o *FundamentalsEarningsHistory) HasEpsActual() bool {
	if o != nil && o.EpsActual != nil {
		return true
	}

	return false
}

// SetEpsActual gets a reference to the given float64 and assigns it to the EpsActual field.
func (o *FundamentalsEarningsHistory) SetEpsActual(v float64) {
	o.EpsActual = &v
}

// GetEpsEstimate returns the EpsEstimate field value if set, zero value otherwise.
func (o *FundamentalsEarningsHistory) GetEpsEstimate() float64 {
	if o == nil || o.EpsEstimate == nil {
		var ret float64
		return ret
	}
	return *o.EpsEstimate
}

// GetEpsEstimateOk returns a tuple with the EpsEstimate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FundamentalsEarningsHistory) GetEpsEstimateOk() (*float64, bool) {
	if o == nil || o.EpsEstimate == nil {
		return nil, false
	}
	return o.EpsEstimate, true
}

// HasEpsEstimate returns a boolean if a field has been set.
func (o *FundamentalsEarningsHistory) HasEpsEstimate() bool {
	if o != nil && o.EpsEstimate != nil {
		return true
	}

	return false
}

// SetEpsEstimate gets a reference to the given float64 and assigns it to the EpsEstimate field.
func (o *FundamentalsEarningsHistory) SetEpsEstimate(v float64) {
	o.EpsEstimate = &v
}

// GetEpsDifference returns the EpsDifference field value if set, zero value otherwise.
func (o *FundamentalsEarningsHistory) GetEpsDifference() float64 {
	if o == nil || o.EpsDifference == nil {
		var ret float64
		return ret
	}
	return *o.EpsDifference
}

// GetEpsDifferenceOk returns a tuple with the EpsDifference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FundamentalsEarningsHistory) GetEpsDifferenceOk() (*float64, bool) {
	if o == nil || o.EpsDifference == nil {
		return nil, false
	}
	return o.EpsDifference, true
}

// HasEpsDifference returns a boolean if a field has been set.
func (o *FundamentalsEarningsHistory) HasEpsDifference() bool {
	if o != nil && o.EpsDifference != nil {
		return true
	}

	return false
}

// SetEpsDifference gets a reference to the given float64 and assigns it to the EpsDifference field.
func (o *FundamentalsEarningsHistory) SetEpsDifference(v float64) {
	o.EpsDifference = &v
}

// GetSurprisePercent returns the SurprisePercent field value if set, zero value otherwise.
func (o *FundamentalsEarningsHistory) GetSurprisePercent() float64 {
	if o == nil || o.SurprisePercent == nil {
		var ret float64
		return ret
	}
	return *o.SurprisePercent
}

// GetSurprisePercentOk returns a tuple with the SurprisePercent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FundamentalsEarningsHistory) GetSurprisePercentOk() (*float64, bool) {
	if o == nil || o.SurprisePercent == nil {
		return nil, false
	}
	return o.SurprisePercent, true
}

// HasSurprisePercent returns a boolean if a field has been set.
func (o *FundamentalsEarningsHistory) HasSurprisePercent() bool {
	if o != nil && o.SurprisePercent != nil {
		return true
	}

	return false
}

// SetSurprisePercent gets a reference to the given float64 and assigns it to the SurprisePercent field.
func (o *FundamentalsEarningsHistory) SetSurprisePercent(v float64) {
	o.SurprisePercent = &v
}

func (o FundamentalsEarningsHistory) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ReportDate != nil {
		toSerialize["reportDate"] = o.ReportDate
	}
	if o.Date != nil {
		toSerialize["date"] = o.Date
	}
	if o.BeforeAfterMarket != nil {
		toSerialize["beforeAfterMarket"] = o.BeforeAfterMarket
	}
	if o.Currency != nil {
		toSerialize["currency"] = o.Currency
	}
	if o.EpsActual != nil {
		toSerialize["epsActual"] = o.EpsActual
	}
	if o.EpsEstimate != nil {
		toSerialize["epsEstimate"] = o.EpsEstimate
	}
	if o.EpsDifference != nil {
		toSerialize["epsDifference"] = o.EpsDifference
	}
	if o.SurprisePercent != nil {
		toSerialize["surprisePercent"] = o.SurprisePercent
	}
	return json.Marshal(toSerialize)
}

type NullableFundamentalsEarningsHistory struct {
	value *FundamentalsEarningsHistory
	isSet bool
}

func (v NullableFundamentalsEarningsHistory) Get() *FundamentalsEarningsHistory {
	return v.value
}

func (v *NullableFundamentalsEarningsHistory) Set(val *FundamentalsEarningsHistory) {
	v.value = val
	v.isSet = true
}

func (v NullableFundamentalsEarningsHistory) IsSet() bool {
	return v.isSet
}

func (v *NullableFundamentalsEarningsHistory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFundamentalsEarningsHistory(val *FundamentalsEarningsHistory) *NullableFundamentalsEarningsHistory {
	return &NullableFundamentalsEarningsHistory{value: val, isSet: true}
}

func (v NullableFundamentalsEarningsHistory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFundamentalsEarningsHistory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
