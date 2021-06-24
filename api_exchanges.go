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
	"bytes"
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
	"strings"
)

// Linger please
var (
	_ _context.Context
)

// ExchangesApiService ExchangesApi service
type ExchangesApiService service

type ApiListExchangeDetailsRequest struct {
	ctx        _context.Context
	ApiService *ExchangesApiService
	exchange   string
	fmt        *string
	from       *string
	to         *string
}

func (r ApiListExchangeDetailsRequest) Fmt(fmt string) ApiListExchangeDetailsRequest {
	r.fmt = &fmt
	return r
}
func (r ApiListExchangeDetailsRequest) From(from string) ApiListExchangeDetailsRequest {
	r.from = &from
	return r
}
func (r ApiListExchangeDetailsRequest) To(to string) ApiListExchangeDetailsRequest {
	r.to = &to
	return r
}

func (r ApiListExchangeDetailsRequest) Execute() (ExchangeDetails, *_nethttp.Response, error) {
	return r.ApiService.ListExchangeDetailsExecute(r)
}

/*
 * ListExchangeDetails Method for ListExchangeDetails
 * List properties of exchangedetails
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param exchange string exchange (name or id) of the exchangedetails
 * @return ApiListExchangeDetailsRequest
 */
func (a *ExchangesApiService) ListExchangeDetails(ctx _context.Context, exchange string) ApiListExchangeDetailsRequest {
	return ApiListExchangeDetailsRequest{
		ApiService: a,
		ctx:        ctx,
		exchange:   exchange,
	}
}

/*
 * Execute executes the request
 * @return ExchangeDetails
 */
func (a *ExchangesApiService) ListExchangeDetailsExecute(r ApiListExchangeDetailsRequest) (ExchangeDetails, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  ExchangeDetails
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ExchangesApiService.ListExchangeDetails")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/exchange-details/{exchange}"
	localVarPath = strings.Replace(localVarPath, "{"+"exchange"+"}", _neturl.PathEscape(parameterToString(r.exchange, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.fmt != nil {
		localVarQueryParams.Add("fmt", parameterToString(*r.fmt, ""))
	}
	if r.from != nil {
		localVarQueryParams.Add("from", parameterToString(*r.from, ""))
	}
	if r.to != nil {
		localVarQueryParams.Add("to", parameterToString(*r.to, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ApiKeyAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarQueryParams.Add("api_token", key)
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiListExchangeTickersRequest struct {
	ctx        _context.Context
	ApiService *ExchangesApiService
	exchange   string
	fmt        *string
}

func (r ApiListExchangeTickersRequest) Fmt(fmt string) ApiListExchangeTickersRequest {
	r.fmt = &fmt
	return r
}

func (r ApiListExchangeTickersRequest) Execute() ([]ExchangeTicker, *_nethttp.Response, error) {
	return r.ApiService.ListExchangeTickersExecute(r)
}

/*
 * ListExchangeTickers Method for ListExchangeTickers
 * List properties of exchangetickers
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param exchange string exchange (name or id) of the exchangetickers
 * @return ApiListExchangeTickersRequest
 */
func (a *ExchangesApiService) ListExchangeTickers(ctx _context.Context, exchange string) ApiListExchangeTickersRequest {
	return ApiListExchangeTickersRequest{
		ApiService: a,
		ctx:        ctx,
		exchange:   exchange,
	}
}

/*
 * Execute executes the request
 * @return []ExchangeTicker
 */
func (a *ExchangesApiService) ListExchangeTickersExecute(r ApiListExchangeTickersRequest) ([]ExchangeTicker, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []ExchangeTicker
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ExchangesApiService.ListExchangeTickers")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/exchange-symbol-list/{exchange}"
	localVarPath = strings.Replace(localVarPath, "{"+"exchange"+"}", _neturl.PathEscape(parameterToString(r.exchange, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.fmt == nil {
		return localVarReturnValue, nil, reportError("fmt is required and must be specified")
	}

	localVarQueryParams.Add("fmt", parameterToString(*r.fmt, ""))
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ApiKeyAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarQueryParams.Add("api_token", key)
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiListExchangesRequest struct {
	ctx        _context.Context
	ApiService *ExchangesApiService
	fmt        *string
}

func (r ApiListExchangesRequest) Fmt(fmt string) ApiListExchangesRequest {
	r.fmt = &fmt
	return r
}

func (r ApiListExchangesRequest) Execute() ([]Exchange, *_nethttp.Response, error) {
	return r.ApiService.ListExchangesExecute(r)
}

/*
 * ListExchanges Method for ListExchanges
 * List properties of exchanges
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiListExchangesRequest
 */
func (a *ExchangesApiService) ListExchanges(ctx _context.Context) ApiListExchangesRequest {
	return ApiListExchangesRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

/*
 * Execute executes the request
 * @return []Exchange
 */
func (a *ExchangesApiService) ListExchangesExecute(r ApiListExchangesRequest) ([]Exchange, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []Exchange
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ExchangesApiService.ListExchanges")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/exchanges-list"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.fmt == nil {
		return localVarReturnValue, nil, reportError("fmt is required and must be specified")
	}

	localVarQueryParams.Add("fmt", parameterToString(*r.fmt, ""))
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ApiKeyAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarQueryParams.Add("api_token", key)
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
