
/*
 * NiFi Rest Api
 *
 * The Rest Api provides programmatic access to command and control a NiFi instance in real time. Start and                                              stop processors, monitor queues, query provenance data, and more. Each endpoint below includes a description,                                             definitions of the expected input and output, potential response codes, and the authorizations required                                             to invoke each service.
 *
 * API version: 1.12.0
 * Contact: dev@nifi.apache.org
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nifi

import (
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"fmt"
)

// Linger please
var (
	_ context.Context
)

type DataTransferApiService service

/* 
DataTransferApiService Commit or cancel the specified transaction

 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param responseCode The response code. Available values are BAD_CHECKSUM(19), CONFIRM_TRANSACTION(12) or CANCEL_TRANSACTION(15).
 * @param portId The input port id.
 * @param transactionId The transaction id.

@return TransactionResultEntity
*/
func (a *DataTransferApiService) CommitInputPortTransaction(ctx context.Context, responseCode int32, portId string, transactionId string) (TransactionResultEntity, *http.Response, *string, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue TransactionResultEntity
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/data-transfer/input-ports/{portId}/transactions/{transactionId}"
	localVarPath = strings.Replace(localVarPath, "{"+"portId"+"}", fmt.Sprintf("%v", portId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"transactionId"+"}", fmt.Sprintf("%v", transactionId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarQueryParams.Add("responseCode", parameterToString(responseCode, ""))
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/octet-stream"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, nil, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, nil, err
	}

	localStringBody := string(localVarBody)

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
		if err == nil { 
			return localVarReturnValue, localVarHttpResponse, &localStringBody, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		
		if localVarHttpResponse.StatusCode == 200 {
			var v TransactionResultEntity
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, &localStringBody, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, &localStringBody, newErr
		}
		
		return localVarReturnValue, localVarHttpResponse, &localStringBody, newErr
	}

	return localVarReturnValue, localVarHttpResponse, &localStringBody, nil
}

/* 
DataTransferApiService Commit or cancel the specified transaction

 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param responseCode The response code. Available values are CONFIRM_TRANSACTION(12) or CANCEL_TRANSACTION(15).
 * @param checksum A checksum calculated at client side using CRC32 to check flow file content integrity. It must match with the value calculated at server side.
 * @param portId The output port id.
 * @param transactionId The transaction id.

@return TransactionResultEntity
*/
func (a *DataTransferApiService) CommitOutputPortTransaction(ctx context.Context, responseCode int32, checksum string, portId string, transactionId string) (TransactionResultEntity, *http.Response, *string, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue TransactionResultEntity
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/data-transfer/output-ports/{portId}/transactions/{transactionId}"
	localVarPath = strings.Replace(localVarPath, "{"+"portId"+"}", fmt.Sprintf("%v", portId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"transactionId"+"}", fmt.Sprintf("%v", transactionId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarQueryParams.Add("responseCode", parameterToString(responseCode, ""))
	localVarQueryParams.Add("checksum", parameterToString(checksum, ""))
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/octet-stream"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, nil, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, nil, err
	}

	localStringBody := string(localVarBody)

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
		if err == nil { 
			return localVarReturnValue, localVarHttpResponse, &localStringBody, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		
		if localVarHttpResponse.StatusCode == 200 {
			var v TransactionResultEntity
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, &localStringBody, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, &localStringBody, newErr
		}
		
		return localVarReturnValue, localVarHttpResponse, &localStringBody, newErr
	}

	return localVarReturnValue, localVarHttpResponse, &localStringBody, nil
}

/* 
DataTransferApiService Create a transaction to the specified output port or input port

 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param portType The port type.
 * @param portId

@return TransactionResultEntity
*/
func (a *DataTransferApiService) CreatePortTransaction(ctx context.Context, portType string, portId string) (TransactionResultEntity, *http.Response, *string, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue TransactionResultEntity
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/data-transfer/{portType}/{portId}/transactions"
	localVarPath = strings.Replace(localVarPath, "{"+"portType"+"}", fmt.Sprintf("%v", portType), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"portId"+"}", fmt.Sprintf("%v", portId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, nil, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, nil, err
	}

	localStringBody := string(localVarBody)

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
		if err == nil { 
			return localVarReturnValue, localVarHttpResponse, &localStringBody, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		
		if localVarHttpResponse.StatusCode == 200 {
			var v TransactionResultEntity
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, &localStringBody, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, &localStringBody, newErr
		}
		
		return localVarReturnValue, localVarHttpResponse, &localStringBody, newErr
	}

	return localVarReturnValue, localVarHttpResponse, &localStringBody, nil
}

/* 
DataTransferApiService Extend transaction TTL

 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param portId
 * @param transactionId

@return TransactionResultEntity
*/
func (a *DataTransferApiService) ExtendInputPortTransactionTTL(ctx context.Context, portId string, transactionId string) (TransactionResultEntity, *http.Response, *string, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue TransactionResultEntity
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/data-transfer/input-ports/{portId}/transactions/{transactionId}"
	localVarPath = strings.Replace(localVarPath, "{"+"portId"+"}", fmt.Sprintf("%v", portId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"transactionId"+"}", fmt.Sprintf("%v", transactionId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"*/*"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, nil, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, nil, err
	}

	localStringBody := string(localVarBody)

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
		if err == nil { 
			return localVarReturnValue, localVarHttpResponse, &localStringBody, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		
		if localVarHttpResponse.StatusCode == 200 {
			var v TransactionResultEntity
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, &localStringBody, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, &localStringBody, newErr
		}
		
		return localVarReturnValue, localVarHttpResponse, &localStringBody, newErr
	}

	return localVarReturnValue, localVarHttpResponse, &localStringBody, nil
}

/* 
DataTransferApiService Extend transaction TTL

 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param portId
 * @param transactionId

@return TransactionResultEntity
*/
func (a *DataTransferApiService) ExtendOutputPortTransactionTTL(ctx context.Context, portId string, transactionId string) (TransactionResultEntity, *http.Response, *string, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue TransactionResultEntity
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/data-transfer/output-ports/{portId}/transactions/{transactionId}"
	localVarPath = strings.Replace(localVarPath, "{"+"portId"+"}", fmt.Sprintf("%v", portId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"transactionId"+"}", fmt.Sprintf("%v", transactionId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"*/*"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, nil, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, nil, err
	}

	localStringBody := string(localVarBody)

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
		if err == nil { 
			return localVarReturnValue, localVarHttpResponse, &localStringBody, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		
		if localVarHttpResponse.StatusCode == 200 {
			var v TransactionResultEntity
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, &localStringBody, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, &localStringBody, newErr
		}
		
		return localVarReturnValue, localVarHttpResponse, &localStringBody, newErr
	}

	return localVarReturnValue, localVarHttpResponse, &localStringBody, nil
}

/* 
DataTransferApiService Transfer flow files to the input port

 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param portId The input port id.
 * @param transactionId

@return string
*/
func (a *DataTransferApiService) ReceiveFlowFiles(ctx context.Context, portId string, transactionId string) (string, *http.Response, *string, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue string
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/data-transfer/input-ports/{portId}/transactions/{transactionId}/flow-files"
	localVarPath = strings.Replace(localVarPath, "{"+"portId"+"}", fmt.Sprintf("%v", portId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"transactionId"+"}", fmt.Sprintf("%v", transactionId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/octet-stream"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"text/plain"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, nil, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, nil, err
	}

	localStringBody := string(localVarBody)

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
		if err == nil { 
			return localVarReturnValue, localVarHttpResponse, &localStringBody, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		
		if localVarHttpResponse.StatusCode == 200 {
			var v string
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, &localStringBody, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, &localStringBody, newErr
		}
		
		return localVarReturnValue, localVarHttpResponse, &localStringBody, newErr
	}

	return localVarReturnValue, localVarHttpResponse, &localStringBody, nil
}

/* 
DataTransferApiService Transfer flow files from the output port

 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param portId The output port id.
 * @param transactionId

@return StreamingOutput
*/
func (a *DataTransferApiService) TransferFlowFiles(ctx context.Context, portId string, transactionId string) (StreamingOutput, *http.Response, *string, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue StreamingOutput
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/data-transfer/output-ports/{portId}/transactions/{transactionId}/flow-files"
	localVarPath = strings.Replace(localVarPath, "{"+"portId"+"}", fmt.Sprintf("%v", portId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"transactionId"+"}", fmt.Sprintf("%v", transactionId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"*/*"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/octet-stream"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, nil, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, nil, err
	}

	localStringBody := string(localVarBody)

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
		if err == nil { 
			return localVarReturnValue, localVarHttpResponse, &localStringBody, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		
		if localVarHttpResponse.StatusCode == 200 {
			var v StreamingOutput
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, &localStringBody, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, &localStringBody, newErr
		}
		
		return localVarReturnValue, localVarHttpResponse, &localStringBody, newErr
	}

	return localVarReturnValue, localVarHttpResponse, &localStringBody, nil
}
