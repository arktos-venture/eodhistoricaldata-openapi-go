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

// FundamentalsInsiderTransactionsData struct for FundamentalsInsiderTransactionsData
type FundamentalsInsiderTransactionsData struct {
	Date                        *string  `json:"date,omitempty"`
	OwnerCik                    *string  `json:"ownerCik,omitempty"`
	OwnerName                   *string  `json:"ownerName,omitempty"`
	TransactionDate             *string  `json:"transactionDate,omitempty"`
	TransactionCode             *string  `json:"transactionCode,omitempty"`
	TransactionAmount           *int64   `json:"transactionAmount,omitempty"`
	TransactionPrice            *float64 `json:"transactionPrice,omitempty"`
	TransactionAcquiredDisposed *string  `json:"transactionAcquiredDisposed,omitempty"`
	PostTransactionAmount       *int64   `json:"postTransactionAmount,omitempty"`
	SecLink                     *string  `json:"secLink,omitempty"`
}

// NewFundamentalsInsiderTransactionsData instantiates a new FundamentalsInsiderTransactionsData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFundamentalsInsiderTransactionsData() *FundamentalsInsiderTransactionsData {
	this := FundamentalsInsiderTransactionsData{}
	return &this
}

// NewFundamentalsInsiderTransactionsDataWithDefaults instantiates a new FundamentalsInsiderTransactionsData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFundamentalsInsiderTransactionsDataWithDefaults() *FundamentalsInsiderTransactionsData {
	this := FundamentalsInsiderTransactionsData{}
	return &this
}

// GetDate returns the Date field value if set, zero value otherwise.
func (o *FundamentalsInsiderTransactionsData) GetDate() string {
	if o == nil || o.Date == nil {
		var ret string
		return ret
	}
	return *o.Date
}

// GetDateOk returns a tuple with the Date field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FundamentalsInsiderTransactionsData) GetDateOk() (*string, bool) {
	if o == nil || o.Date == nil {
		return nil, false
	}
	return o.Date, true
}

// HasDate returns a boolean if a field has been set.
func (o *FundamentalsInsiderTransactionsData) HasDate() bool {
	if o != nil && o.Date != nil {
		return true
	}

	return false
}

// SetDate gets a reference to the given string and assigns it to the Date field.
func (o *FundamentalsInsiderTransactionsData) SetDate(v string) {
	o.Date = &v
}

// GetOwnerCik returns the OwnerCik field value if set, zero value otherwise.
func (o *FundamentalsInsiderTransactionsData) GetOwnerCik() string {
	if o == nil || o.OwnerCik == nil {
		var ret string
		return ret
	}
	return *o.OwnerCik
}

// GetOwnerCikOk returns a tuple with the OwnerCik field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FundamentalsInsiderTransactionsData) GetOwnerCikOk() (*string, bool) {
	if o == nil || o.OwnerCik == nil {
		return nil, false
	}
	return o.OwnerCik, true
}

// HasOwnerCik returns a boolean if a field has been set.
func (o *FundamentalsInsiderTransactionsData) HasOwnerCik() bool {
	if o != nil && o.OwnerCik != nil {
		return true
	}

	return false
}

// SetOwnerCik gets a reference to the given string and assigns it to the OwnerCik field.
func (o *FundamentalsInsiderTransactionsData) SetOwnerCik(v string) {
	o.OwnerCik = &v
}

// GetOwnerName returns the OwnerName field value if set, zero value otherwise.
func (o *FundamentalsInsiderTransactionsData) GetOwnerName() string {
	if o == nil || o.OwnerName == nil {
		var ret string
		return ret
	}
	return *o.OwnerName
}

// GetOwnerNameOk returns a tuple with the OwnerName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FundamentalsInsiderTransactionsData) GetOwnerNameOk() (*string, bool) {
	if o == nil || o.OwnerName == nil {
		return nil, false
	}
	return o.OwnerName, true
}

// HasOwnerName returns a boolean if a field has been set.
func (o *FundamentalsInsiderTransactionsData) HasOwnerName() bool {
	if o != nil && o.OwnerName != nil {
		return true
	}

	return false
}

// SetOwnerName gets a reference to the given string and assigns it to the OwnerName field.
func (o *FundamentalsInsiderTransactionsData) SetOwnerName(v string) {
	o.OwnerName = &v
}

// GetTransactionDate returns the TransactionDate field value if set, zero value otherwise.
func (o *FundamentalsInsiderTransactionsData) GetTransactionDate() string {
	if o == nil || o.TransactionDate == nil {
		var ret string
		return ret
	}
	return *o.TransactionDate
}

// GetTransactionDateOk returns a tuple with the TransactionDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FundamentalsInsiderTransactionsData) GetTransactionDateOk() (*string, bool) {
	if o == nil || o.TransactionDate == nil {
		return nil, false
	}
	return o.TransactionDate, true
}

// HasTransactionDate returns a boolean if a field has been set.
func (o *FundamentalsInsiderTransactionsData) HasTransactionDate() bool {
	if o != nil && o.TransactionDate != nil {
		return true
	}

	return false
}

// SetTransactionDate gets a reference to the given string and assigns it to the TransactionDate field.
func (o *FundamentalsInsiderTransactionsData) SetTransactionDate(v string) {
	o.TransactionDate = &v
}

// GetTransactionCode returns the TransactionCode field value if set, zero value otherwise.
func (o *FundamentalsInsiderTransactionsData) GetTransactionCode() string {
	if o == nil || o.TransactionCode == nil {
		var ret string
		return ret
	}
	return *o.TransactionCode
}

// GetTransactionCodeOk returns a tuple with the TransactionCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FundamentalsInsiderTransactionsData) GetTransactionCodeOk() (*string, bool) {
	if o == nil || o.TransactionCode == nil {
		return nil, false
	}
	return o.TransactionCode, true
}

// HasTransactionCode returns a boolean if a field has been set.
func (o *FundamentalsInsiderTransactionsData) HasTransactionCode() bool {
	if o != nil && o.TransactionCode != nil {
		return true
	}

	return false
}

// SetTransactionCode gets a reference to the given string and assigns it to the TransactionCode field.
func (o *FundamentalsInsiderTransactionsData) SetTransactionCode(v string) {
	o.TransactionCode = &v
}

// GetTransactionAmount returns the TransactionAmount field value if set, zero value otherwise.
func (o *FundamentalsInsiderTransactionsData) GetTransactionAmount() int64 {
	if o == nil || o.TransactionAmount == nil {
		var ret int64
		return ret
	}
	return *o.TransactionAmount
}

// GetTransactionAmountOk returns a tuple with the TransactionAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FundamentalsInsiderTransactionsData) GetTransactionAmountOk() (*int64, bool) {
	if o == nil || o.TransactionAmount == nil {
		return nil, false
	}
	return o.TransactionAmount, true
}

// HasTransactionAmount returns a boolean if a field has been set.
func (o *FundamentalsInsiderTransactionsData) HasTransactionAmount() bool {
	if o != nil && o.TransactionAmount != nil {
		return true
	}

	return false
}

// SetTransactionAmount gets a reference to the given int64 and assigns it to the TransactionAmount field.
func (o *FundamentalsInsiderTransactionsData) SetTransactionAmount(v int64) {
	o.TransactionAmount = &v
}

// GetTransactionPrice returns the TransactionPrice field value if set, zero value otherwise.
func (o *FundamentalsInsiderTransactionsData) GetTransactionPrice() float64 {
	if o == nil || o.TransactionPrice == nil {
		var ret float64
		return ret
	}
	return *o.TransactionPrice
}

// GetTransactionPriceOk returns a tuple with the TransactionPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FundamentalsInsiderTransactionsData) GetTransactionPriceOk() (*float64, bool) {
	if o == nil || o.TransactionPrice == nil {
		return nil, false
	}
	return o.TransactionPrice, true
}

// HasTransactionPrice returns a boolean if a field has been set.
func (o *FundamentalsInsiderTransactionsData) HasTransactionPrice() bool {
	if o != nil && o.TransactionPrice != nil {
		return true
	}

	return false
}

// SetTransactionPrice gets a reference to the given float64 and assigns it to the TransactionPrice field.
func (o *FundamentalsInsiderTransactionsData) SetTransactionPrice(v float64) {
	o.TransactionPrice = &v
}

// GetTransactionAcquiredDisposed returns the TransactionAcquiredDisposed field value if set, zero value otherwise.
func (o *FundamentalsInsiderTransactionsData) GetTransactionAcquiredDisposed() string {
	if o == nil || o.TransactionAcquiredDisposed == nil {
		var ret string
		return ret
	}
	return *o.TransactionAcquiredDisposed
}

// GetTransactionAcquiredDisposedOk returns a tuple with the TransactionAcquiredDisposed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FundamentalsInsiderTransactionsData) GetTransactionAcquiredDisposedOk() (*string, bool) {
	if o == nil || o.TransactionAcquiredDisposed == nil {
		return nil, false
	}
	return o.TransactionAcquiredDisposed, true
}

// HasTransactionAcquiredDisposed returns a boolean if a field has been set.
func (o *FundamentalsInsiderTransactionsData) HasTransactionAcquiredDisposed() bool {
	if o != nil && o.TransactionAcquiredDisposed != nil {
		return true
	}

	return false
}

// SetTransactionAcquiredDisposed gets a reference to the given string and assigns it to the TransactionAcquiredDisposed field.
func (o *FundamentalsInsiderTransactionsData) SetTransactionAcquiredDisposed(v string) {
	o.TransactionAcquiredDisposed = &v
}

// GetPostTransactionAmount returns the PostTransactionAmount field value if set, zero value otherwise.
func (o *FundamentalsInsiderTransactionsData) GetPostTransactionAmount() int64 {
	if o == nil || o.PostTransactionAmount == nil {
		var ret int64
		return ret
	}
	return *o.PostTransactionAmount
}

// GetPostTransactionAmountOk returns a tuple with the PostTransactionAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FundamentalsInsiderTransactionsData) GetPostTransactionAmountOk() (*int64, bool) {
	if o == nil || o.PostTransactionAmount == nil {
		return nil, false
	}
	return o.PostTransactionAmount, true
}

// HasPostTransactionAmount returns a boolean if a field has been set.
func (o *FundamentalsInsiderTransactionsData) HasPostTransactionAmount() bool {
	if o != nil && o.PostTransactionAmount != nil {
		return true
	}

	return false
}

// SetPostTransactionAmount gets a reference to the given int64 and assigns it to the PostTransactionAmount field.
func (o *FundamentalsInsiderTransactionsData) SetPostTransactionAmount(v int64) {
	o.PostTransactionAmount = &v
}

// GetSecLink returns the SecLink field value if set, zero value otherwise.
func (o *FundamentalsInsiderTransactionsData) GetSecLink() string {
	if o == nil || o.SecLink == nil {
		var ret string
		return ret
	}
	return *o.SecLink
}

// GetSecLinkOk returns a tuple with the SecLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FundamentalsInsiderTransactionsData) GetSecLinkOk() (*string, bool) {
	if o == nil || o.SecLink == nil {
		return nil, false
	}
	return o.SecLink, true
}

// HasSecLink returns a boolean if a field has been set.
func (o *FundamentalsInsiderTransactionsData) HasSecLink() bool {
	if o != nil && o.SecLink != nil {
		return true
	}

	return false
}

// SetSecLink gets a reference to the given string and assigns it to the SecLink field.
func (o *FundamentalsInsiderTransactionsData) SetSecLink(v string) {
	o.SecLink = &v
}

func (o FundamentalsInsiderTransactionsData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Date != nil {
		toSerialize["date"] = o.Date
	}
	if o.OwnerCik != nil {
		toSerialize["ownerCik"] = o.OwnerCik
	}
	if o.OwnerName != nil {
		toSerialize["ownerName"] = o.OwnerName
	}
	if o.TransactionDate != nil {
		toSerialize["transactionDate"] = o.TransactionDate
	}
	if o.TransactionCode != nil {
		toSerialize["transactionCode"] = o.TransactionCode
	}
	if o.TransactionAmount != nil {
		toSerialize["transactionAmount"] = o.TransactionAmount
	}
	if o.TransactionPrice != nil {
		toSerialize["transactionPrice"] = o.TransactionPrice
	}
	if o.TransactionAcquiredDisposed != nil {
		toSerialize["transactionAcquiredDisposed"] = o.TransactionAcquiredDisposed
	}
	if o.PostTransactionAmount != nil {
		toSerialize["postTransactionAmount"] = o.PostTransactionAmount
	}
	if o.SecLink != nil {
		toSerialize["secLink"] = o.SecLink
	}
	return json.Marshal(toSerialize)
}

type NullableFundamentalsInsiderTransactionsData struct {
	value *FundamentalsInsiderTransactionsData
	isSet bool
}

func (v NullableFundamentalsInsiderTransactionsData) Get() *FundamentalsInsiderTransactionsData {
	return v.value
}

func (v *NullableFundamentalsInsiderTransactionsData) Set(val *FundamentalsInsiderTransactionsData) {
	v.value = val
	v.isSet = true
}

func (v NullableFundamentalsInsiderTransactionsData) IsSet() bool {
	return v.isSet
}

func (v *NullableFundamentalsInsiderTransactionsData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFundamentalsInsiderTransactionsData(val *FundamentalsInsiderTransactionsData) *NullableFundamentalsInsiderTransactionsData {
	return &NullableFundamentalsInsiderTransactionsData{value: val, isSet: true}
}

func (v NullableFundamentalsInsiderTransactionsData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFundamentalsInsiderTransactionsData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
