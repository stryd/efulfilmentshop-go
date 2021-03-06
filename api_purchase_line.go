/*
 * E-fulfilment Shop API
 *
 * The E-fulfilment Shop API allows you to integrate your service with our warehouse.
 *
 * API version: 1.3.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package efulfilmentshop

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

// PurchaseLineApiService PurchaseLineApi service
type PurchaseLineApiService service

type ApiGetPurchaseLineCollectionRequest struct {
	ctx _context.Context
	ApiService *PurchaseLineApiService
	purchaseId *int32
	page *int32
	items *int32
}

func (r ApiGetPurchaseLineCollectionRequest) PurchaseId(purchaseId int32) ApiGetPurchaseLineCollectionRequest {
	r.purchaseId = &purchaseId
	return r
}
func (r ApiGetPurchaseLineCollectionRequest) Page(page int32) ApiGetPurchaseLineCollectionRequest {
	r.page = &page
	return r
}
func (r ApiGetPurchaseLineCollectionRequest) Items(items int32) ApiGetPurchaseLineCollectionRequest {
	r.items = &items
	return r
}

func (r ApiGetPurchaseLineCollectionRequest) Execute() ([]PurchaseLineRead, *_nethttp.Response, error) {
	return r.ApiService.GetPurchaseLineCollectionExecute(r)
}

/*
 * GetPurchaseLineCollection Retrieves the collection of PurchaseLine resources.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiGetPurchaseLineCollectionRequest
 */
func (a *PurchaseLineApiService) GetPurchaseLineCollection(ctx _context.Context) ApiGetPurchaseLineCollectionRequest {
	return ApiGetPurchaseLineCollectionRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return []PurchaseLineRead
 */
func (a *PurchaseLineApiService) GetPurchaseLineCollectionExecute(r ApiGetPurchaseLineCollectionRequest) ([]PurchaseLineRead, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []PurchaseLineRead
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PurchaseLineApiService.GetPurchaseLineCollection")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/purchase_lines"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.purchaseId != nil {
		localVarQueryParams.Add("purchaseId", parameterToString(*r.purchaseId, ""))
	}
	if r.page != nil {
		localVarQueryParams.Add("page", parameterToString(*r.page, ""))
	}
	if r.items != nil {
		localVarQueryParams.Add("items", parameterToString(*r.items, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "application/ld+json", "text/html"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
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

type ApiGetPurchaseLineItemRequest struct {
	ctx _context.Context
	ApiService *PurchaseLineApiService
	id string
}


func (r ApiGetPurchaseLineItemRequest) Execute() (PurchaseLineRead, *_nethttp.Response, error) {
	return r.ApiService.GetPurchaseLineItemExecute(r)
}

/*
 * GetPurchaseLineItem Retrieves a PurchaseLine resource.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param id
 * @return ApiGetPurchaseLineItemRequest
 */
func (a *PurchaseLineApiService) GetPurchaseLineItem(ctx _context.Context, id string) ApiGetPurchaseLineItemRequest {
	return ApiGetPurchaseLineItemRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

/*
 * Execute executes the request
 * @return PurchaseLineRead
 */
func (a *PurchaseLineApiService) GetPurchaseLineItemExecute(r ApiGetPurchaseLineItemRequest) (PurchaseLineRead, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  PurchaseLineRead
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PurchaseLineApiService.GetPurchaseLineItem")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/purchase_lines/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", _neturl.PathEscape(parameterToString(r.id, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "application/ld+json", "text/html"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
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

type ApiPatchPurchaseLineItemRequest struct {
	ctx _context.Context
	ApiService *PurchaseLineApiService
	id string
	purchaseLineWrite *PurchaseLineWrite
}

func (r ApiPatchPurchaseLineItemRequest) PurchaseLineWrite(purchaseLineWrite PurchaseLineWrite) ApiPatchPurchaseLineItemRequest {
	r.purchaseLineWrite = &purchaseLineWrite
	return r
}

func (r ApiPatchPurchaseLineItemRequest) Execute() (PurchaseLineRead, *_nethttp.Response, error) {
	return r.ApiService.PatchPurchaseLineItemExecute(r)
}

/*
 * PatchPurchaseLineItem Updates the PurchaseLine resource.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param id
 * @return ApiPatchPurchaseLineItemRequest
 */
func (a *PurchaseLineApiService) PatchPurchaseLineItem(ctx _context.Context, id string) ApiPatchPurchaseLineItemRequest {
	return ApiPatchPurchaseLineItemRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

/*
 * Execute executes the request
 * @return PurchaseLineRead
 */
func (a *PurchaseLineApiService) PatchPurchaseLineItemExecute(r ApiPatchPurchaseLineItemRequest) (PurchaseLineRead, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPatch
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  PurchaseLineRead
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PurchaseLineApiService.PatchPurchaseLineItem")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/purchase_lines/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", _neturl.PathEscape(parameterToString(r.id, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/merge-patch+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "application/ld+json", "text/html"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.purchaseLineWrite
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

type ApiPostPurchaseLineCollectionRequest struct {
	ctx _context.Context
	ApiService *PurchaseLineApiService
	purchaseLineWrite *PurchaseLineWrite
}

func (r ApiPostPurchaseLineCollectionRequest) PurchaseLineWrite(purchaseLineWrite PurchaseLineWrite) ApiPostPurchaseLineCollectionRequest {
	r.purchaseLineWrite = &purchaseLineWrite
	return r
}

func (r ApiPostPurchaseLineCollectionRequest) Execute() (PurchaseLineRead, *_nethttp.Response, error) {
	return r.ApiService.PostPurchaseLineCollectionExecute(r)
}

/*
 * PostPurchaseLineCollection Creates a PurchaseLine resource.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiPostPurchaseLineCollectionRequest
 */
func (a *PurchaseLineApiService) PostPurchaseLineCollection(ctx _context.Context) ApiPostPurchaseLineCollectionRequest {
	return ApiPostPurchaseLineCollectionRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return PurchaseLineRead
 */
func (a *PurchaseLineApiService) PostPurchaseLineCollectionExecute(r ApiPostPurchaseLineCollectionRequest) (PurchaseLineRead, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  PurchaseLineRead
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PurchaseLineApiService.PostPurchaseLineCollection")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/purchase_lines"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json", "application/ld+json", "text/html"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "application/ld+json", "text/html"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.purchaseLineWrite
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

type ApiPutPurchaseLineItemRequest struct {
	ctx _context.Context
	ApiService *PurchaseLineApiService
	id string
	purchaseLineWrite *PurchaseLineWrite
}

func (r ApiPutPurchaseLineItemRequest) PurchaseLineWrite(purchaseLineWrite PurchaseLineWrite) ApiPutPurchaseLineItemRequest {
	r.purchaseLineWrite = &purchaseLineWrite
	return r
}

func (r ApiPutPurchaseLineItemRequest) Execute() (PurchaseLineRead, *_nethttp.Response, error) {
	return r.ApiService.PutPurchaseLineItemExecute(r)
}

/*
 * PutPurchaseLineItem Replaces the PurchaseLine resource.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param id
 * @return ApiPutPurchaseLineItemRequest
 */
func (a *PurchaseLineApiService) PutPurchaseLineItem(ctx _context.Context, id string) ApiPutPurchaseLineItemRequest {
	return ApiPutPurchaseLineItemRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

/*
 * Execute executes the request
 * @return PurchaseLineRead
 */
func (a *PurchaseLineApiService) PutPurchaseLineItemExecute(r ApiPutPurchaseLineItemRequest) (PurchaseLineRead, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPut
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  PurchaseLineRead
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PurchaseLineApiService.PutPurchaseLineItem")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/purchase_lines/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", _neturl.PathEscape(parameterToString(r.id, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json", "application/ld+json", "text/html"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "application/ld+json", "text/html"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.purchaseLineWrite
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
