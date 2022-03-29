# \QuotesApi

All URIs are relative to *https://eodhistoricaldata.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListEodBulkLastDay**](QuotesApi.md#ListEodBulkLastDay) | **Get** /eod-bulk-last-day/{exchange} | 
[**ListHistoryIntradayQuotes**](QuotesApi.md#ListHistoryIntradayQuotes) | **Get** /intraday/{ticker} | 
[**ListHistoryQuotes**](QuotesApi.md#ListHistoryQuotes) | **Get** /eod/{ticker} | 
[**ListShortsQuotes**](QuotesApi.md#ListShortsQuotes) | **Get** /shorts/{ticker} | 
[**ListSplitsQuotes**](QuotesApi.md#ListSplitsQuotes) | **Get** /splits/{ticker} | 
[**ReadRealtimeQuotes**](QuotesApi.md#ReadRealtimeQuotes) | **Get** /real-time/{ticker} | 



## ListEodBulkLastDay

> []EodBulkLastDayItem ListEodBulkLastDay(ctx, exchange).Fmt(fmt).Symbols(symbols).Type_(type_).Date(date).Filter(filter).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    exchange := "exchange_example" // string | string exchange (name or id) of the eodbulklastday
    fmt := "["json","csv"]" // string | string fmt (name or id) of the eodbulklastday
    symbols := "symbols_example" // string | string symbols (name or id) of the eodbulklastday (optional)
    type_ := "type__example" // string | string type (name or id) of the eodbulklastday (optional)
    date := "date_example" // string | string date (name or id) of the eodbulklastday (optional)
    filter := "filter_example" // string | string filter (name or id) of the eodbulklastday (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.QuotesApi.ListEodBulkLastDay(context.Background(), exchange).Fmt(fmt).Symbols(symbols).Type_(type_).Date(date).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `QuotesApi.ListEodBulkLastDay``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListEodBulkLastDay`: []EodBulkLastDayItem
    fmt.Fprintf(os.Stdout, "Response from `QuotesApi.ListEodBulkLastDay`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**exchange** | **string** | string exchange (name or id) of the eodbulklastday | 

### Other Parameters

Other parameters are passed through a pointer to a apiListEodBulkLastDayRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fmt** | **string** | string fmt (name or id) of the eodbulklastday | 
 **symbols** | **string** | string symbols (name or id) of the eodbulklastday | 
 **type_** | **string** | string type (name or id) of the eodbulklastday | 
 **date** | **string** | string date (name or id) of the eodbulklastday | 
 **filter** | **string** | string filter (name or id) of the eodbulklastday | 

### Return type

[**[]EodBulkLastDayItem**](EodBulkLastDayItem.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListHistoryIntradayQuotes

> []Quote ListHistoryIntradayQuotes(ctx, ticker).Fmt(fmt).Interval(interval).From(from).To(to).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    ticker := "ticker_example" // string | string ticker (name or id) of the historyintradayquotes
    fmt := "["json","csv"]" // string | string fmt (name or id) of the historyintradayquotes
    interval := "["1m","5m"]" // string | string interval (name or id) of the historyintradayquotes
    from := "[10000,40000]" // string | string from (name or id) of the historyintradayquotes
    to := "[10000,40000]" // string | string to (name or id) of the historyintradayquotes

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.QuotesApi.ListHistoryIntradayQuotes(context.Background(), ticker).Fmt(fmt).Interval(interval).From(from).To(to).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `QuotesApi.ListHistoryIntradayQuotes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListHistoryIntradayQuotes`: []Quote
    fmt.Fprintf(os.Stdout, "Response from `QuotesApi.ListHistoryIntradayQuotes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ticker** | **string** | string ticker (name or id) of the historyintradayquotes | 

### Other Parameters

Other parameters are passed through a pointer to a apiListHistoryIntradayQuotesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fmt** | **string** | string fmt (name or id) of the historyintradayquotes | 
 **interval** | **string** | string interval (name or id) of the historyintradayquotes | 
 **from** | **string** | string from (name or id) of the historyintradayquotes | 
 **to** | **string** | string to (name or id) of the historyintradayquotes | 

### Return type

[**[]Quote**](Quote.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListHistoryQuotes

> []Quote ListHistoryQuotes(ctx, ticker).Period(period).Fmt(fmt).Filter(filter).Order(order).From(from).To(to).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    ticker := "["EUR.FOREX","BZ","CL","CT","GC","HG","KC","MALTR","MZN","NG","NICKEL","PA","PL","SB","SI","ZC"]" // string | string ticker (name or id) of the historyquotes
    period := "["d","w","m"]" // string | string period (name or id) of the historyquotes (optional)
    fmt := "["json","csv"]" // string | string fmt (name or id) of the historyquotes (optional)
    filter := "["last_close"]" // string | string filter (name or id) of the historyquotes (optional)
    order := "["d"]" // string | string order (name or id) of the historyquotes (optional)
    from := "["2021-03-01"]" // string | string from (name or id) of the historyquotes (optional)
    to := "["2021-03-10"]" // string | string to (name or id) of the historyquotes (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.QuotesApi.ListHistoryQuotes(context.Background(), ticker).Period(period).Fmt(fmt).Filter(filter).Order(order).From(from).To(to).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `QuotesApi.ListHistoryQuotes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListHistoryQuotes`: []Quote
    fmt.Fprintf(os.Stdout, "Response from `QuotesApi.ListHistoryQuotes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ticker** | **string** | string ticker (name or id) of the historyquotes | 

### Other Parameters

Other parameters are passed through a pointer to a apiListHistoryQuotesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **period** | **string** | string period (name or id) of the historyquotes | 
 **fmt** | **string** | string fmt (name or id) of the historyquotes | 
 **filter** | **string** | string filter (name or id) of the historyquotes | 
 **order** | **string** | string order (name or id) of the historyquotes | 
 **from** | **string** | string from (name or id) of the historyquotes | 
 **to** | **string** | string to (name or id) of the historyquotes | 

### Return type

[**[]Quote**](Quote.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListShortsQuotes

> []SplitsQuote ListShortsQuotes(ctx, ticker).Fmt(fmt).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    ticker := "ticker_example" // string | string ticker (name or id) of the shortsquotes
    fmt := "["json","csv"]" // string | string fmt (name or id) of the shortsquotes

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.QuotesApi.ListShortsQuotes(context.Background(), ticker).Fmt(fmt).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `QuotesApi.ListShortsQuotes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListShortsQuotes`: []SplitsQuote
    fmt.Fprintf(os.Stdout, "Response from `QuotesApi.ListShortsQuotes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ticker** | **string** | string ticker (name or id) of the shortsquotes | 

### Other Parameters

Other parameters are passed through a pointer to a apiListShortsQuotesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fmt** | **string** | string fmt (name or id) of the shortsquotes | 

### Return type

[**[]SplitsQuote**](SplitsQuote.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSplitsQuotes

> []SplitsQuote ListSplitsQuotes(ctx, ticker).Fmt(fmt).From(from).To(to).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    ticker := "ticker_example" // string | string ticker (name or id) of the splitsquotes
    fmt := "["json","csv"]" // string | string fmt (name or id) of the splitsquotes
    from := "["2021-03-01"]" // string | string from (name or id) of the splitsquotes
    to := "["2021-03-10"]" // string | string to (name or id) of the splitsquotes

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.QuotesApi.ListSplitsQuotes(context.Background(), ticker).Fmt(fmt).From(from).To(to).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `QuotesApi.ListSplitsQuotes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSplitsQuotes`: []SplitsQuote
    fmt.Fprintf(os.Stdout, "Response from `QuotesApi.ListSplitsQuotes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ticker** | **string** | string ticker (name or id) of the splitsquotes | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSplitsQuotesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fmt** | **string** | string fmt (name or id) of the splitsquotes | 
 **from** | **string** | string from (name or id) of the splitsquotes | 
 **to** | **string** | string to (name or id) of the splitsquotes | 

### Return type

[**[]SplitsQuote**](SplitsQuote.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadRealtimeQuotes

> Quote ReadRealtimeQuotes(ctx, ticker).Fmt(fmt).Interval(interval).Filter(filter).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    ticker := "ticker_example" // string | string ticker (name or id) of the realtimequotes
    fmt := "["json","csv"]" // string | string fmt (name or id) of the realtimequotes
    interval := "["1m","5m"]" // string | string interval (name or id) of the realtimequotes (optional)
    filter := "["close"]" // string | string filter (name or id) of the realtimequotes (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.QuotesApi.ReadRealtimeQuotes(context.Background(), ticker).Fmt(fmt).Interval(interval).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `QuotesApi.ReadRealtimeQuotes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadRealtimeQuotes`: Quote
    fmt.Fprintf(os.Stdout, "Response from `QuotesApi.ReadRealtimeQuotes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ticker** | **string** | string ticker (name or id) of the realtimequotes | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadRealtimeQuotesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fmt** | **string** | string fmt (name or id) of the realtimequotes | 
 **interval** | **string** | string interval (name or id) of the realtimequotes | 
 **filter** | **string** | string filter (name or id) of the realtimequotes | 

### Return type

[**Quote**](Quote.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

