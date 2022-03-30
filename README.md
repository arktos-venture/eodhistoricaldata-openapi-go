# Go API client for openapi

eodhistoricaldata API

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 1.0.0
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import openapi "github.com/GIT_USER_ID/GIT_REPO_ID"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sw.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), openapi.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), openapi.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```
ctx := context.WithValue(context.Background(), openapi.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), openapi.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://eodhistoricaldata.com/api*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*BondsApi* | [**ReadBondFundamentals**](docs/BondsApi.md#readbondfundamentals) | **Get** /bond-fundamentals/{bond} | 
*CompaniesApi* | [**ListBulkFundamentals**](docs/CompaniesApi.md#listbulkfundamentals) | **Get** /bulk-fundamentals/{exchange} | 
*CompaniesApi* | [**ReadFundamentals**](docs/CompaniesApi.md#readfundamentals) | **Get** /fundamentals/{ticker} | 
*DividendsApi* | [**ListDividends**](docs/DividendsApi.md#listdividends) | **Get** /div/{ticker} | 
*ExchangesApi* | [**ListExchangeDetails**](docs/ExchangesApi.md#listexchangedetails) | **Get** /exchange-details/{exchange} | 
*ExchangesApi* | [**ListExchangeTickers**](docs/ExchangesApi.md#listexchangetickers) | **Get** /exchange-symbol-list/{exchange} | 
*ExchangesApi* | [**ListExchanges**](docs/ExchangesApi.md#listexchanges) | **Get** /exchanges-list | 
*MacroIndicatorApi* | [**ListMacroIndicator**](docs/MacroIndicatorApi.md#listmacroindicator) | **Get** /macro-indicator/{country} | 
*NewsApi* | [**ListNews**](docs/NewsApi.md#listnews) | **Get** /news | 
*QuotesApi* | [**ListEodBulkLastDay**](docs/QuotesApi.md#listeodbulklastday) | **Get** /eod-bulk-last-day/{exchange} | 
*QuotesApi* | [**ListHistoryIntradayQuotes**](docs/QuotesApi.md#listhistoryintradayquotes) | **Get** /intraday/{ticker} | 
*QuotesApi* | [**ListHistoryQuotes**](docs/QuotesApi.md#listhistoryquotes) | **Get** /eod/{ticker} | 
*QuotesApi* | [**ListShortsQuotes**](docs/QuotesApi.md#listshortsquotes) | **Get** /shorts/{ticker} | 
*QuotesApi* | [**ListSplitsQuotes**](docs/QuotesApi.md#listsplitsquotes) | **Get** /splits/{ticker} | 
*QuotesApi* | [**ReadRealtimeQuotes**](docs/QuotesApi.md#readrealtimequotes) | **Get** /real-time/{ticker} | 


## Documentation For Models

 - [BondFundamentals](docs/BondFundamentals.md)
 - [BondFundamentalsClassification](docs/BondFundamentalsClassification.md)
 - [BondFundamentalsIssueData](docs/BondFundamentalsIssueData.md)
 - [BondFundamentalsRating](docs/BondFundamentalsRating.md)
 - [CashFlow](docs/CashFlow.md)
 - [Dividend](docs/Dividend.md)
 - [EodBulkLastDayItem](docs/EodBulkLastDayItem.md)
 - [Exchange](docs/Exchange.md)
 - [ExchangeDetails](docs/ExchangeDetails.md)
 - [ExchangeDetailsExchangeHolidays](docs/ExchangeDetailsExchangeHolidays.md)
 - [ExchangeDetailsTradingHours](docs/ExchangeDetailsTradingHours.md)
 - [ExchangeTicker](docs/ExchangeTicker.md)
 - [Fundamentals](docs/Fundamentals.md)
 - [FundamentalsAnalystRatings](docs/FundamentalsAnalystRatings.md)
 - [FundamentalsComponents](docs/FundamentalsComponents.md)
 - [FundamentalsESGScores](docs/FundamentalsESGScores.md)
 - [FundamentalsESGScoresActivitiesInvolvement](docs/FundamentalsESGScoresActivitiesInvolvement.md)
 - [FundamentalsEarnings](docs/FundamentalsEarnings.md)
 - [FundamentalsEarningsHistory](docs/FundamentalsEarningsHistory.md)
 - [FundamentalsEarningsTrend](docs/FundamentalsEarningsTrend.md)
 - [FundamentalsFinancial](docs/FundamentalsFinancial.md)
 - [FundamentalsFinancials](docs/FundamentalsFinancials.md)
 - [FundamentalsFinancialsBalanceSheet](docs/FundamentalsFinancialsBalanceSheet.md)
 - [FundamentalsFinancialsCashFlow](docs/FundamentalsFinancialsCashFlow.md)
 - [FundamentalsFinancialsIncomeStatement](docs/FundamentalsFinancialsIncomeStatement.md)
 - [FundamentalsGeneral](docs/FundamentalsGeneral.md)
 - [FundamentalsGeneralAddressData](docs/FundamentalsGeneralAddressData.md)
 - [FundamentalsGeneralListings](docs/FundamentalsGeneralListings.md)
 - [FundamentalsGeneralOfficers](docs/FundamentalsGeneralOfficers.md)
 - [FundamentalsHighlights](docs/FundamentalsHighlights.md)
 - [FundamentalsHistoricalTickerComponentsData](docs/FundamentalsHistoricalTickerComponentsData.md)
 - [FundamentalsHolders](docs/FundamentalsHolders.md)
 - [FundamentalsHoldersData](docs/FundamentalsHoldersData.md)
 - [FundamentalsInsiderTransactionsData](docs/FundamentalsInsiderTransactionsData.md)
 - [FundamentalsOutstandingShares](docs/FundamentalsOutstandingShares.md)
 - [FundamentalsOutstandingSharesData](docs/FundamentalsOutstandingSharesData.md)
 - [FundamentalsSharesStats](docs/FundamentalsSharesStats.md)
 - [FundamentalsSplitsDividends](docs/FundamentalsSplitsDividends.md)
 - [FundamentalsSplitsDividendsNumberDividendsByYear](docs/FundamentalsSplitsDividendsNumberDividendsByYear.md)
 - [FundamentalsTechnicals](docs/FundamentalsTechnicals.md)
 - [FundamentalsValuation](docs/FundamentalsValuation.md)
 - [IncomeStatement](docs/IncomeStatement.md)
 - [MacroIndicator](docs/MacroIndicator.md)
 - [New](docs/New.md)
 - [Quote](docs/Quote.md)
 - [ShortsQuote](docs/ShortsQuote.md)
 - [SplitsQuote](docs/SplitsQuote.md)


## Documentation For Authorization



### ApiKeyAuth

- **Type**: API key
- **API key parameter name**: api_token
- **Location**: URL query string

Note, each API key must be added to a map of `map[string]APIKey` where the key is: api_token and passed in as the auth context for each request.


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author



