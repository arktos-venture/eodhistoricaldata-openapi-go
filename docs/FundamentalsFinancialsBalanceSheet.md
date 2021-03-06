# FundamentalsFinancialsBalanceSheet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrencySymbol** | Pointer to **string** |  | [optional] 
**Quarterly** | Pointer to [**map[string]FundamentalsFinancial**](FundamentalsFinancial.md) |  | [optional] 
**Yearly** | Pointer to [**map[string]FundamentalsFinancial**](FundamentalsFinancial.md) |  | [optional] 

## Methods

### NewFundamentalsFinancialsBalanceSheet

`func NewFundamentalsFinancialsBalanceSheet() *FundamentalsFinancialsBalanceSheet`

NewFundamentalsFinancialsBalanceSheet instantiates a new FundamentalsFinancialsBalanceSheet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFundamentalsFinancialsBalanceSheetWithDefaults

`func NewFundamentalsFinancialsBalanceSheetWithDefaults() *FundamentalsFinancialsBalanceSheet`

NewFundamentalsFinancialsBalanceSheetWithDefaults instantiates a new FundamentalsFinancialsBalanceSheet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrencySymbol

`func (o *FundamentalsFinancialsBalanceSheet) GetCurrencySymbol() string`

GetCurrencySymbol returns the CurrencySymbol field if non-nil, zero value otherwise.

### GetCurrencySymbolOk

`func (o *FundamentalsFinancialsBalanceSheet) GetCurrencySymbolOk() (*string, bool)`

GetCurrencySymbolOk returns a tuple with the CurrencySymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencySymbol

`func (o *FundamentalsFinancialsBalanceSheet) SetCurrencySymbol(v string)`

SetCurrencySymbol sets CurrencySymbol field to given value.

### HasCurrencySymbol

`func (o *FundamentalsFinancialsBalanceSheet) HasCurrencySymbol() bool`

HasCurrencySymbol returns a boolean if a field has been set.

### GetQuarterly

`func (o *FundamentalsFinancialsBalanceSheet) GetQuarterly() map[string]FundamentalsFinancial`

GetQuarterly returns the Quarterly field if non-nil, zero value otherwise.

### GetQuarterlyOk

`func (o *FundamentalsFinancialsBalanceSheet) GetQuarterlyOk() (*map[string]FundamentalsFinancial, bool)`

GetQuarterlyOk returns a tuple with the Quarterly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuarterly

`func (o *FundamentalsFinancialsBalanceSheet) SetQuarterly(v map[string]FundamentalsFinancial)`

SetQuarterly sets Quarterly field to given value.

### HasQuarterly

`func (o *FundamentalsFinancialsBalanceSheet) HasQuarterly() bool`

HasQuarterly returns a boolean if a field has been set.

### GetYearly

`func (o *FundamentalsFinancialsBalanceSheet) GetYearly() map[string]FundamentalsFinancial`

GetYearly returns the Yearly field if non-nil, zero value otherwise.

### GetYearlyOk

`func (o *FundamentalsFinancialsBalanceSheet) GetYearlyOk() (*map[string]FundamentalsFinancial, bool)`

GetYearlyOk returns a tuple with the Yearly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYearly

`func (o *FundamentalsFinancialsBalanceSheet) SetYearly(v map[string]FundamentalsFinancial)`

SetYearly sets Yearly field to given value.

### HasYearly

`func (o *FundamentalsFinancialsBalanceSheet) HasYearly() bool`

HasYearly returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


