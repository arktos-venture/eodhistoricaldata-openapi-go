# EodBulkLastDayItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Date** | Pointer to **string** |  | [optional] 
**Code** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**ExchangeShortName** | Pointer to **string** |  | [optional] 
**Exchange** | Pointer to **string** |  | [optional] 
**Dividend** | Pointer to **string** |  | [optional] 
**Currency** | Pointer to **string** |  | [optional] 
**DeclarationDate** | Pointer to **string** |  | [optional] 
**RecordDate** | Pointer to **string** |  | [optional] 
**PaymentDate** | Pointer to **string** |  | [optional] 
**Period** | Pointer to **string** |  | [optional] 
**UnadjustedValue** | Pointer to **string** |  | [optional] 
**MarketCapitalization** | Pointer to **float64** |  | [optional] 
**Beta** | Pointer to **float64** |  | [optional] 
**Open** | Pointer to **float64** |  | [optional] 
**High** | Pointer to **float64** |  | [optional] 
**Low** | Pointer to **float64** |  | [optional] 
**Close** | Pointer to **float64** |  | [optional] 
**AdjustedClose** | Pointer to **float64** |  | [optional] 
**Volume** | Pointer to **float64** |  | [optional] 
**Ema50d** | Pointer to **float64** |  | [optional] 
**Ema200d** | Pointer to **float64** |  | [optional] 
**Hi250d** | Pointer to **float64** |  | [optional] 
**Lo250d** | Pointer to **float64** |  | [optional] 
**Avgvol14d** | Pointer to **float64** |  | [optional] 
**Avgvol50d** | Pointer to **float64** |  | [optional] 
**Avgvol200d** | Pointer to **float64** |  | [optional] 
**Split** | Pointer to **string** |  | [optional] 

## Methods

### NewEodBulkLastDayItem

`func NewEodBulkLastDayItem() *EodBulkLastDayItem`

NewEodBulkLastDayItem instantiates a new EodBulkLastDayItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEodBulkLastDayItemWithDefaults

`func NewEodBulkLastDayItemWithDefaults() *EodBulkLastDayItem`

NewEodBulkLastDayItemWithDefaults instantiates a new EodBulkLastDayItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDate

`func (o *EodBulkLastDayItem) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *EodBulkLastDayItem) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *EodBulkLastDayItem) SetDate(v string)`

SetDate sets Date field to given value.

### HasDate

`func (o *EodBulkLastDayItem) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetCode

`func (o *EodBulkLastDayItem) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *EodBulkLastDayItem) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *EodBulkLastDayItem) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *EodBulkLastDayItem) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetType

`func (o *EodBulkLastDayItem) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EodBulkLastDayItem) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EodBulkLastDayItem) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *EodBulkLastDayItem) HasType() bool`

HasType returns a boolean if a field has been set.

### GetExchangeShortName

`func (o *EodBulkLastDayItem) GetExchangeShortName() string`

GetExchangeShortName returns the ExchangeShortName field if non-nil, zero value otherwise.

### GetExchangeShortNameOk

`func (o *EodBulkLastDayItem) GetExchangeShortNameOk() (*string, bool)`

GetExchangeShortNameOk returns a tuple with the ExchangeShortName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeShortName

`func (o *EodBulkLastDayItem) SetExchangeShortName(v string)`

SetExchangeShortName sets ExchangeShortName field to given value.

### HasExchangeShortName

`func (o *EodBulkLastDayItem) HasExchangeShortName() bool`

HasExchangeShortName returns a boolean if a field has been set.

### GetExchange

`func (o *EodBulkLastDayItem) GetExchange() string`

GetExchange returns the Exchange field if non-nil, zero value otherwise.

### GetExchangeOk

`func (o *EodBulkLastDayItem) GetExchangeOk() (*string, bool)`

GetExchangeOk returns a tuple with the Exchange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchange

`func (o *EodBulkLastDayItem) SetExchange(v string)`

SetExchange sets Exchange field to given value.

### HasExchange

`func (o *EodBulkLastDayItem) HasExchange() bool`

HasExchange returns a boolean if a field has been set.

### GetDividend

`func (o *EodBulkLastDayItem) GetDividend() string`

GetDividend returns the Dividend field if non-nil, zero value otherwise.

### GetDividendOk

`func (o *EodBulkLastDayItem) GetDividendOk() (*string, bool)`

GetDividendOk returns a tuple with the Dividend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDividend

`func (o *EodBulkLastDayItem) SetDividend(v string)`

SetDividend sets Dividend field to given value.

### HasDividend

`func (o *EodBulkLastDayItem) HasDividend() bool`

HasDividend returns a boolean if a field has been set.

### GetCurrency

`func (o *EodBulkLastDayItem) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *EodBulkLastDayItem) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *EodBulkLastDayItem) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *EodBulkLastDayItem) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetDeclarationDate

`func (o *EodBulkLastDayItem) GetDeclarationDate() string`

GetDeclarationDate returns the DeclarationDate field if non-nil, zero value otherwise.

### GetDeclarationDateOk

`func (o *EodBulkLastDayItem) GetDeclarationDateOk() (*string, bool)`

GetDeclarationDateOk returns a tuple with the DeclarationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeclarationDate

`func (o *EodBulkLastDayItem) SetDeclarationDate(v string)`

SetDeclarationDate sets DeclarationDate field to given value.

### HasDeclarationDate

`func (o *EodBulkLastDayItem) HasDeclarationDate() bool`

HasDeclarationDate returns a boolean if a field has been set.

### GetRecordDate

`func (o *EodBulkLastDayItem) GetRecordDate() string`

GetRecordDate returns the RecordDate field if non-nil, zero value otherwise.

### GetRecordDateOk

`func (o *EodBulkLastDayItem) GetRecordDateOk() (*string, bool)`

GetRecordDateOk returns a tuple with the RecordDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordDate

`func (o *EodBulkLastDayItem) SetRecordDate(v string)`

SetRecordDate sets RecordDate field to given value.

### HasRecordDate

`func (o *EodBulkLastDayItem) HasRecordDate() bool`

HasRecordDate returns a boolean if a field has been set.

### GetPaymentDate

`func (o *EodBulkLastDayItem) GetPaymentDate() string`

GetPaymentDate returns the PaymentDate field if non-nil, zero value otherwise.

### GetPaymentDateOk

`func (o *EodBulkLastDayItem) GetPaymentDateOk() (*string, bool)`

GetPaymentDateOk returns a tuple with the PaymentDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentDate

`func (o *EodBulkLastDayItem) SetPaymentDate(v string)`

SetPaymentDate sets PaymentDate field to given value.

### HasPaymentDate

`func (o *EodBulkLastDayItem) HasPaymentDate() bool`

HasPaymentDate returns a boolean if a field has been set.

### GetPeriod

`func (o *EodBulkLastDayItem) GetPeriod() string`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *EodBulkLastDayItem) GetPeriodOk() (*string, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *EodBulkLastDayItem) SetPeriod(v string)`

SetPeriod sets Period field to given value.

### HasPeriod

`func (o *EodBulkLastDayItem) HasPeriod() bool`

HasPeriod returns a boolean if a field has been set.

### GetUnadjustedValue

`func (o *EodBulkLastDayItem) GetUnadjustedValue() string`

GetUnadjustedValue returns the UnadjustedValue field if non-nil, zero value otherwise.

### GetUnadjustedValueOk

`func (o *EodBulkLastDayItem) GetUnadjustedValueOk() (*string, bool)`

GetUnadjustedValueOk returns a tuple with the UnadjustedValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnadjustedValue

`func (o *EodBulkLastDayItem) SetUnadjustedValue(v string)`

SetUnadjustedValue sets UnadjustedValue field to given value.

### HasUnadjustedValue

`func (o *EodBulkLastDayItem) HasUnadjustedValue() bool`

HasUnadjustedValue returns a boolean if a field has been set.

### GetMarketCapitalization

`func (o *EodBulkLastDayItem) GetMarketCapitalization() float64`

GetMarketCapitalization returns the MarketCapitalization field if non-nil, zero value otherwise.

### GetMarketCapitalizationOk

`func (o *EodBulkLastDayItem) GetMarketCapitalizationOk() (*float64, bool)`

GetMarketCapitalizationOk returns a tuple with the MarketCapitalization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketCapitalization

`func (o *EodBulkLastDayItem) SetMarketCapitalization(v float64)`

SetMarketCapitalization sets MarketCapitalization field to given value.

### HasMarketCapitalization

`func (o *EodBulkLastDayItem) HasMarketCapitalization() bool`

HasMarketCapitalization returns a boolean if a field has been set.

### GetBeta

`func (o *EodBulkLastDayItem) GetBeta() float64`

GetBeta returns the Beta field if non-nil, zero value otherwise.

### GetBetaOk

`func (o *EodBulkLastDayItem) GetBetaOk() (*float64, bool)`

GetBetaOk returns a tuple with the Beta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeta

`func (o *EodBulkLastDayItem) SetBeta(v float64)`

SetBeta sets Beta field to given value.

### HasBeta

`func (o *EodBulkLastDayItem) HasBeta() bool`

HasBeta returns a boolean if a field has been set.

### GetOpen

`func (o *EodBulkLastDayItem) GetOpen() float64`

GetOpen returns the Open field if non-nil, zero value otherwise.

### GetOpenOk

`func (o *EodBulkLastDayItem) GetOpenOk() (*float64, bool)`

GetOpenOk returns a tuple with the Open field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpen

`func (o *EodBulkLastDayItem) SetOpen(v float64)`

SetOpen sets Open field to given value.

### HasOpen

`func (o *EodBulkLastDayItem) HasOpen() bool`

HasOpen returns a boolean if a field has been set.

### GetHigh

`func (o *EodBulkLastDayItem) GetHigh() float64`

GetHigh returns the High field if non-nil, zero value otherwise.

### GetHighOk

`func (o *EodBulkLastDayItem) GetHighOk() (*float64, bool)`

GetHighOk returns a tuple with the High field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHigh

`func (o *EodBulkLastDayItem) SetHigh(v float64)`

SetHigh sets High field to given value.

### HasHigh

`func (o *EodBulkLastDayItem) HasHigh() bool`

HasHigh returns a boolean if a field has been set.

### GetLow

`func (o *EodBulkLastDayItem) GetLow() float64`

GetLow returns the Low field if non-nil, zero value otherwise.

### GetLowOk

`func (o *EodBulkLastDayItem) GetLowOk() (*float64, bool)`

GetLowOk returns a tuple with the Low field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLow

`func (o *EodBulkLastDayItem) SetLow(v float64)`

SetLow sets Low field to given value.

### HasLow

`func (o *EodBulkLastDayItem) HasLow() bool`

HasLow returns a boolean if a field has been set.

### GetClose

`func (o *EodBulkLastDayItem) GetClose() float64`

GetClose returns the Close field if non-nil, zero value otherwise.

### GetCloseOk

`func (o *EodBulkLastDayItem) GetCloseOk() (*float64, bool)`

GetCloseOk returns a tuple with the Close field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClose

`func (o *EodBulkLastDayItem) SetClose(v float64)`

SetClose sets Close field to given value.

### HasClose

`func (o *EodBulkLastDayItem) HasClose() bool`

HasClose returns a boolean if a field has been set.

### GetAdjustedClose

`func (o *EodBulkLastDayItem) GetAdjustedClose() float64`

GetAdjustedClose returns the AdjustedClose field if non-nil, zero value otherwise.

### GetAdjustedCloseOk

`func (o *EodBulkLastDayItem) GetAdjustedCloseOk() (*float64, bool)`

GetAdjustedCloseOk returns a tuple with the AdjustedClose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdjustedClose

`func (o *EodBulkLastDayItem) SetAdjustedClose(v float64)`

SetAdjustedClose sets AdjustedClose field to given value.

### HasAdjustedClose

`func (o *EodBulkLastDayItem) HasAdjustedClose() bool`

HasAdjustedClose returns a boolean if a field has been set.

### GetVolume

`func (o *EodBulkLastDayItem) GetVolume() float64`

GetVolume returns the Volume field if non-nil, zero value otherwise.

### GetVolumeOk

`func (o *EodBulkLastDayItem) GetVolumeOk() (*float64, bool)`

GetVolumeOk returns a tuple with the Volume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolume

`func (o *EodBulkLastDayItem) SetVolume(v float64)`

SetVolume sets Volume field to given value.

### HasVolume

`func (o *EodBulkLastDayItem) HasVolume() bool`

HasVolume returns a boolean if a field has been set.

### GetEma50d

`func (o *EodBulkLastDayItem) GetEma50d() float64`

GetEma50d returns the Ema50d field if non-nil, zero value otherwise.

### GetEma50dOk

`func (o *EodBulkLastDayItem) GetEma50dOk() (*float64, bool)`

GetEma50dOk returns a tuple with the Ema50d field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEma50d

`func (o *EodBulkLastDayItem) SetEma50d(v float64)`

SetEma50d sets Ema50d field to given value.

### HasEma50d

`func (o *EodBulkLastDayItem) HasEma50d() bool`

HasEma50d returns a boolean if a field has been set.

### GetEma200d

`func (o *EodBulkLastDayItem) GetEma200d() float64`

GetEma200d returns the Ema200d field if non-nil, zero value otherwise.

### GetEma200dOk

`func (o *EodBulkLastDayItem) GetEma200dOk() (*float64, bool)`

GetEma200dOk returns a tuple with the Ema200d field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEma200d

`func (o *EodBulkLastDayItem) SetEma200d(v float64)`

SetEma200d sets Ema200d field to given value.

### HasEma200d

`func (o *EodBulkLastDayItem) HasEma200d() bool`

HasEma200d returns a boolean if a field has been set.

### GetHi250d

`func (o *EodBulkLastDayItem) GetHi250d() float64`

GetHi250d returns the Hi250d field if non-nil, zero value otherwise.

### GetHi250dOk

`func (o *EodBulkLastDayItem) GetHi250dOk() (*float64, bool)`

GetHi250dOk returns a tuple with the Hi250d field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHi250d

`func (o *EodBulkLastDayItem) SetHi250d(v float64)`

SetHi250d sets Hi250d field to given value.

### HasHi250d

`func (o *EodBulkLastDayItem) HasHi250d() bool`

HasHi250d returns a boolean if a field has been set.

### GetLo250d

`func (o *EodBulkLastDayItem) GetLo250d() float64`

GetLo250d returns the Lo250d field if non-nil, zero value otherwise.

### GetLo250dOk

`func (o *EodBulkLastDayItem) GetLo250dOk() (*float64, bool)`

GetLo250dOk returns a tuple with the Lo250d field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLo250d

`func (o *EodBulkLastDayItem) SetLo250d(v float64)`

SetLo250d sets Lo250d field to given value.

### HasLo250d

`func (o *EodBulkLastDayItem) HasLo250d() bool`

HasLo250d returns a boolean if a field has been set.

### GetAvgvol14d

`func (o *EodBulkLastDayItem) GetAvgvol14d() float64`

GetAvgvol14d returns the Avgvol14d field if non-nil, zero value otherwise.

### GetAvgvol14dOk

`func (o *EodBulkLastDayItem) GetAvgvol14dOk() (*float64, bool)`

GetAvgvol14dOk returns a tuple with the Avgvol14d field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgvol14d

`func (o *EodBulkLastDayItem) SetAvgvol14d(v float64)`

SetAvgvol14d sets Avgvol14d field to given value.

### HasAvgvol14d

`func (o *EodBulkLastDayItem) HasAvgvol14d() bool`

HasAvgvol14d returns a boolean if a field has been set.

### GetAvgvol50d

`func (o *EodBulkLastDayItem) GetAvgvol50d() float64`

GetAvgvol50d returns the Avgvol50d field if non-nil, zero value otherwise.

### GetAvgvol50dOk

`func (o *EodBulkLastDayItem) GetAvgvol50dOk() (*float64, bool)`

GetAvgvol50dOk returns a tuple with the Avgvol50d field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgvol50d

`func (o *EodBulkLastDayItem) SetAvgvol50d(v float64)`

SetAvgvol50d sets Avgvol50d field to given value.

### HasAvgvol50d

`func (o *EodBulkLastDayItem) HasAvgvol50d() bool`

HasAvgvol50d returns a boolean if a field has been set.

### GetAvgvol200d

`func (o *EodBulkLastDayItem) GetAvgvol200d() float64`

GetAvgvol200d returns the Avgvol200d field if non-nil, zero value otherwise.

### GetAvgvol200dOk

`func (o *EodBulkLastDayItem) GetAvgvol200dOk() (*float64, bool)`

GetAvgvol200dOk returns a tuple with the Avgvol200d field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgvol200d

`func (o *EodBulkLastDayItem) SetAvgvol200d(v float64)`

SetAvgvol200d sets Avgvol200d field to given value.

### HasAvgvol200d

`func (o *EodBulkLastDayItem) HasAvgvol200d() bool`

HasAvgvol200d returns a boolean if a field has been set.

### GetSplit

`func (o *EodBulkLastDayItem) GetSplit() string`

GetSplit returns the Split field if non-nil, zero value otherwise.

### GetSplitOk

`func (o *EodBulkLastDayItem) GetSplitOk() (*string, bool)`

GetSplitOk returns a tuple with the Split field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSplit

`func (o *EodBulkLastDayItem) SetSplit(v string)`

SetSplit sets Split field to given value.

### HasSplit

`func (o *EodBulkLastDayItem) HasSplit() bool`

HasSplit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


