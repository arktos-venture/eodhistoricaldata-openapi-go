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

// ExchangeDetails struct for ExchangeDetails
type ExchangeDetails struct {
	Name             *string                                     `json:"Name,omitempty"`
	Code             *string                                     `json:"Code,omitempty"`
	OperatingMIC     *string                                     `json:"OperatingMIC,omitempty"`
	Country          *string                                     `json:"Country,omitempty"`
	Currency         *string                                     `json:"Currency,omitempty"`
	Timezone         *string                                     `json:"Timezone,omitempty"`
	IsOpen           *bool                                       `json:"isOpen,omitempty"`
	TradingHours     *ExchangeDetailsTradingHours                `json:"TradingHours,omitempty"`
	ExchangeHolidays *map[string]ExchangeDetailsExchangeHolidays `json:"ExchangeHolidays,omitempty"`
	ActiveTickers    *int32                                      `json:"ActiveTickers,omitempty"`
	UpdatedTickers   *int32                                      `json:"UpdatedTickers,omitempty"`
}

// NewExchangeDetails instantiates a new ExchangeDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExchangeDetails() *ExchangeDetails {
	this := ExchangeDetails{}
	return &this
}

// NewExchangeDetailsWithDefaults instantiates a new ExchangeDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExchangeDetailsWithDefaults() *ExchangeDetails {
	this := ExchangeDetails{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ExchangeDetails) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeDetails) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ExchangeDetails) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ExchangeDetails) SetName(v string) {
	o.Name = &v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *ExchangeDetails) GetCode() string {
	if o == nil || o.Code == nil {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeDetails) GetCodeOk() (*string, bool) {
	if o == nil || o.Code == nil {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *ExchangeDetails) HasCode() bool {
	if o != nil && o.Code != nil {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *ExchangeDetails) SetCode(v string) {
	o.Code = &v
}

// GetOperatingMIC returns the OperatingMIC field value if set, zero value otherwise.
func (o *ExchangeDetails) GetOperatingMIC() string {
	if o == nil || o.OperatingMIC == nil {
		var ret string
		return ret
	}
	return *o.OperatingMIC
}

// GetOperatingMICOk returns a tuple with the OperatingMIC field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeDetails) GetOperatingMICOk() (*string, bool) {
	if o == nil || o.OperatingMIC == nil {
		return nil, false
	}
	return o.OperatingMIC, true
}

// HasOperatingMIC returns a boolean if a field has been set.
func (o *ExchangeDetails) HasOperatingMIC() bool {
	if o != nil && o.OperatingMIC != nil {
		return true
	}

	return false
}

// SetOperatingMIC gets a reference to the given string and assigns it to the OperatingMIC field.
func (o *ExchangeDetails) SetOperatingMIC(v string) {
	o.OperatingMIC = &v
}

// GetCountry returns the Country field value if set, zero value otherwise.
func (o *ExchangeDetails) GetCountry() string {
	if o == nil || o.Country == nil {
		var ret string
		return ret
	}
	return *o.Country
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeDetails) GetCountryOk() (*string, bool) {
	if o == nil || o.Country == nil {
		return nil, false
	}
	return o.Country, true
}

// HasCountry returns a boolean if a field has been set.
func (o *ExchangeDetails) HasCountry() bool {
	if o != nil && o.Country != nil {
		return true
	}

	return false
}

// SetCountry gets a reference to the given string and assigns it to the Country field.
func (o *ExchangeDetails) SetCountry(v string) {
	o.Country = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *ExchangeDetails) GetCurrency() string {
	if o == nil || o.Currency == nil {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeDetails) GetCurrencyOk() (*string, bool) {
	if o == nil || o.Currency == nil {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *ExchangeDetails) HasCurrency() bool {
	if o != nil && o.Currency != nil {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *ExchangeDetails) SetCurrency(v string) {
	o.Currency = &v
}

// GetTimezone returns the Timezone field value if set, zero value otherwise.
func (o *ExchangeDetails) GetTimezone() string {
	if o == nil || o.Timezone == nil {
		var ret string
		return ret
	}
	return *o.Timezone
}

// GetTimezoneOk returns a tuple with the Timezone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeDetails) GetTimezoneOk() (*string, bool) {
	if o == nil || o.Timezone == nil {
		return nil, false
	}
	return o.Timezone, true
}

// HasTimezone returns a boolean if a field has been set.
func (o *ExchangeDetails) HasTimezone() bool {
	if o != nil && o.Timezone != nil {
		return true
	}

	return false
}

// SetTimezone gets a reference to the given string and assigns it to the Timezone field.
func (o *ExchangeDetails) SetTimezone(v string) {
	o.Timezone = &v
}

// GetIsOpen returns the IsOpen field value if set, zero value otherwise.
func (o *ExchangeDetails) GetIsOpen() bool {
	if o == nil || o.IsOpen == nil {
		var ret bool
		return ret
	}
	return *o.IsOpen
}

// GetIsOpenOk returns a tuple with the IsOpen field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeDetails) GetIsOpenOk() (*bool, bool) {
	if o == nil || o.IsOpen == nil {
		return nil, false
	}
	return o.IsOpen, true
}

// HasIsOpen returns a boolean if a field has been set.
func (o *ExchangeDetails) HasIsOpen() bool {
	if o != nil && o.IsOpen != nil {
		return true
	}

	return false
}

// SetIsOpen gets a reference to the given bool and assigns it to the IsOpen field.
func (o *ExchangeDetails) SetIsOpen(v bool) {
	o.IsOpen = &v
}

// GetTradingHours returns the TradingHours field value if set, zero value otherwise.
func (o *ExchangeDetails) GetTradingHours() ExchangeDetailsTradingHours {
	if o == nil || o.TradingHours == nil {
		var ret ExchangeDetailsTradingHours
		return ret
	}
	return *o.TradingHours
}

// GetTradingHoursOk returns a tuple with the TradingHours field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeDetails) GetTradingHoursOk() (*ExchangeDetailsTradingHours, bool) {
	if o == nil || o.TradingHours == nil {
		return nil, false
	}
	return o.TradingHours, true
}

// HasTradingHours returns a boolean if a field has been set.
func (o *ExchangeDetails) HasTradingHours() bool {
	if o != nil && o.TradingHours != nil {
		return true
	}

	return false
}

// SetTradingHours gets a reference to the given ExchangeDetailsTradingHours and assigns it to the TradingHours field.
func (o *ExchangeDetails) SetTradingHours(v ExchangeDetailsTradingHours) {
	o.TradingHours = &v
}

// GetExchangeHolidays returns the ExchangeHolidays field value if set, zero value otherwise.
func (o *ExchangeDetails) GetExchangeHolidays() map[string]ExchangeDetailsExchangeHolidays {
	if o == nil || o.ExchangeHolidays == nil {
		var ret map[string]ExchangeDetailsExchangeHolidays
		return ret
	}
	return *o.ExchangeHolidays
}

// GetExchangeHolidaysOk returns a tuple with the ExchangeHolidays field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeDetails) GetExchangeHolidaysOk() (*map[string]ExchangeDetailsExchangeHolidays, bool) {
	if o == nil || o.ExchangeHolidays == nil {
		return nil, false
	}
	return o.ExchangeHolidays, true
}

// HasExchangeHolidays returns a boolean if a field has been set.
func (o *ExchangeDetails) HasExchangeHolidays() bool {
	if o != nil && o.ExchangeHolidays != nil {
		return true
	}

	return false
}

// SetExchangeHolidays gets a reference to the given map[string]ExchangeDetailsExchangeHolidays and assigns it to the ExchangeHolidays field.
func (o *ExchangeDetails) SetExchangeHolidays(v map[string]ExchangeDetailsExchangeHolidays) {
	o.ExchangeHolidays = &v
}

// GetActiveTickers returns the ActiveTickers field value if set, zero value otherwise.
func (o *ExchangeDetails) GetActiveTickers() int32 {
	if o == nil || o.ActiveTickers == nil {
		var ret int32
		return ret
	}
	return *o.ActiveTickers
}

// GetActiveTickersOk returns a tuple with the ActiveTickers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeDetails) GetActiveTickersOk() (*int32, bool) {
	if o == nil || o.ActiveTickers == nil {
		return nil, false
	}
	return o.ActiveTickers, true
}

// HasActiveTickers returns a boolean if a field has been set.
func (o *ExchangeDetails) HasActiveTickers() bool {
	if o != nil && o.ActiveTickers != nil {
		return true
	}

	return false
}

// SetActiveTickers gets a reference to the given int32 and assigns it to the ActiveTickers field.
func (o *ExchangeDetails) SetActiveTickers(v int32) {
	o.ActiveTickers = &v
}

// GetUpdatedTickers returns the UpdatedTickers field value if set, zero value otherwise.
func (o *ExchangeDetails) GetUpdatedTickers() int32 {
	if o == nil || o.UpdatedTickers == nil {
		var ret int32
		return ret
	}
	return *o.UpdatedTickers
}

// GetUpdatedTickersOk returns a tuple with the UpdatedTickers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeDetails) GetUpdatedTickersOk() (*int32, bool) {
	if o == nil || o.UpdatedTickers == nil {
		return nil, false
	}
	return o.UpdatedTickers, true
}

// HasUpdatedTickers returns a boolean if a field has been set.
func (o *ExchangeDetails) HasUpdatedTickers() bool {
	if o != nil && o.UpdatedTickers != nil {
		return true
	}

	return false
}

// SetUpdatedTickers gets a reference to the given int32 and assigns it to the UpdatedTickers field.
func (o *ExchangeDetails) SetUpdatedTickers(v int32) {
	o.UpdatedTickers = &v
}

func (o ExchangeDetails) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.Code != nil {
		toSerialize["Code"] = o.Code
	}
	if o.OperatingMIC != nil {
		toSerialize["OperatingMIC"] = o.OperatingMIC
	}
	if o.Country != nil {
		toSerialize["Country"] = o.Country
	}
	if o.Currency != nil {
		toSerialize["Currency"] = o.Currency
	}
	if o.Timezone != nil {
		toSerialize["Timezone"] = o.Timezone
	}
	if o.IsOpen != nil {
		toSerialize["isOpen"] = o.IsOpen
	}
	if o.TradingHours != nil {
		toSerialize["TradingHours"] = o.TradingHours
	}
	if o.ExchangeHolidays != nil {
		toSerialize["ExchangeHolidays"] = o.ExchangeHolidays
	}
	if o.ActiveTickers != nil {
		toSerialize["ActiveTickers"] = o.ActiveTickers
	}
	if o.UpdatedTickers != nil {
		toSerialize["UpdatedTickers"] = o.UpdatedTickers
	}
	return json.Marshal(toSerialize)
}

type NullableExchangeDetails struct {
	value *ExchangeDetails
	isSet bool
}

func (v NullableExchangeDetails) Get() *ExchangeDetails {
	return v.value
}

func (v *NullableExchangeDetails) Set(val *ExchangeDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableExchangeDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableExchangeDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExchangeDetails(val *ExchangeDetails) *NullableExchangeDetails {
	return &NullableExchangeDetails{value: val, isSet: true}
}

func (v NullableExchangeDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExchangeDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
