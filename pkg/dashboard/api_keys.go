/*
Tyk Dashboard API

## <a name=\"introduction\"></a> Introduction  The Tyk Dashboard API offers granular, programmatic access to a centralised database of resources that your Tyk nodes can pull from. This API has a dynamic user administrative structure which means the secret key that is used to communicate with your Tyk nodes can be kept secret and access to the wider management functions can be handled on a user-by-user and organisation-by-organisation basis.  A common question around using a database-backed configuration is how to programatically add API definitions to your Tyk nodes, the Dashboard API allows much more fine-grained, secure and multi-user access to your Tyk cluster, and should be used to manage a database-backed Tyk node.  The Tyk Dashboard API works seamlessly with the Tyk Dashboard (and the two come bundled together).  ## <a name=\"security-hierarchy\"></a> Security Hierarchy  The Dashboard API provides a more structured security layer to managing Tyk nodes.  ### Organisations, APIs and Users  With the Dashboard API and a database-backed Tyk setup, (and to an extent with file-based API setups - if diligence is used in naming an creating definitions), the following security model is applied to the management of Upstream APIs:  * **Organisations**: All APIs are *owned* by an organisation, this is designated by the `OrgID` parameter in the API Definition. * **Users**: All users created in the Dashboard belong to an organisation (unless an exception is made for super-administrative access). * **APIs**: All APIs belong to an Organisation and only Users that belong to that organisation can see the analytics for those APIs and manage their configurations. * **API Keys**: API Keys are designated by organisation, this means an API key that has full access rights will not be allowed to access the APIs of another organisation on the same system, but can have full access to all APIs within the organisation. * **Access Rights**: Access rights are stored with the key, this enables a key to give access to multiple APIs, this is defined by the session object in the core Tyk API.  In order to use the Dashboard API, you'll need to get the `Tyk Dashboard API Access Credentials` secret from your user profile on the Dashboard UI.  The secret you set should then be sent along as a header with each Dashboard API Request in order for it to be successful:  ``` authorization: <your-secret> ```

API version: 5.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dashboard

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)


type KeysAPI interface {

	/*
	AddKey Create a key

	Tyk will generate the access token based on the OrgID specified in the API Definition and a random UUID. This ensures that keys can be "owned" by different API Owners should segmentation be needed at an organisational level.
<br/><br/>
API keys without access_rights data will be written to all APIs on the system (this also means that they will be created across all SessionHandlers and StorageHandlers, it is recommended to always embed access_rights data in a key to ensure that only targeted APIs and their back-ends are written to.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiAddKeyRequest
	*/
	AddKey(ctx context.Context) ApiAddKeyRequest

	// AddKeyExecute executes the request
	//  @return SessionState
	AddKeyExecute(r ApiAddKeyRequest) (*SessionState, *http.Response, error)

	/*
	CreateCustomKey Create custom key

	Creates a key with a custom key ID.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param keyID The Key ID.
	@return ApiCreateCustomKeyRequest
	*/
	CreateCustomKey(ctx context.Context, keyID string) ApiCreateCustomKeyRequest

	// CreateCustomKeyExecute executes the request
	//  @return SessionState
	CreateCustomKeyExecute(r ApiCreateCustomKeyRequest) (*SessionState, *http.Response, error)

	/*
	DeleteKey Delete key

	Deleting a key will remove it permanently from the system, however analytics relating to that key will still be available.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param apiID ID of API the keys grant access to. Can either be the internal or external API ID.
	@param keyID The Key ID.
	@return ApiDeleteKeyRequest
	*/
	DeleteKey(ctx context.Context, apiID string, keyID string) ApiDeleteKeyRequest

	// DeleteKeyExecute executes the request
	//  @return ApiStatusMessage
	DeleteKeyExecute(r ApiDeleteKeyRequest) (*ApiStatusMessage, *http.Response, error)

	/*
	GetKey Get key

	Fetches the key that grant access to the API with the ID {apiID} and key ID {keyID}.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param apiID ID of API the keys grant access to. Can either be the internal or external API ID.
	@param keyID The Key ID.
	@return ApiGetKeyRequest
	*/
	GetKey(ctx context.Context, apiID string, keyID string) ApiGetKeyRequest

	// GetKeyExecute executes the request
	//  @return SessionState
	GetKeyExecute(r ApiGetKeyRequest) (*SessionState, *http.Response, error)

	/*
	ListKeys List keys

	Lists keys that grant access to the API with the ID {apiID}.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param apiID ID of API the keys grant access to. Can either be the internal or external API ID.
	@return ApiListKeysRequest
	*/
	ListKeys(ctx context.Context, apiID string) ApiListKeysRequest

	// ListKeysExecute executes the request
	//  @return ApiAllKeys
	ListKeysExecute(r ApiListKeysRequest) (*ApiAllKeys, *http.Response, error)

	/*
	UpdateKey Update key

	You can also manually add keys to Tyk using your own key-generation algorithm. It is recommended if using this approach to ensure that the OrgID being used in the API Definition and the key data is blank so that Tyk does not try to prepend or manage the key in any way.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param apiID ID of API the keys grant access to. Can either be the internal or external API ID.
	@param keyID The Key ID.
	@return ApiUpdateKeyRequest
	*/
	UpdateKey(ctx context.Context, apiID string, keyID string) ApiUpdateKeyRequest

	// UpdateKeyExecute executes the request
	//  @return SessionState
	UpdateKeyExecute(r ApiUpdateKeyRequest) (*SessionState, *http.Response, error)
}

// KeysAPIService KeysAPI service
type KeysAPIService service

type ApiAddKeyRequest struct {
	ctx context.Context
	ApiService KeysAPI
	sessionState *SessionState
}

func (r ApiAddKeyRequest) SessionState(sessionState SessionState) ApiAddKeyRequest {
	r.sessionState = &sessionState
	return r
}

func (r ApiAddKeyRequest) Execute() (*SessionState, *http.Response, error) {
	return r.ApiService.AddKeyExecute(r)
}

/*
AddKey Create a key

Tyk will generate the access token based on the OrgID specified in the API Definition and a random UUID. This ensures that keys can be "owned" by different API Owners should segmentation be needed at an organisational level.
<br/><br/>
API keys without access_rights data will be written to all APIs on the system (this also means that they will be created across all SessionHandlers and StorageHandlers, it is recommended to always embed access_rights data in a key to ensure that only targeted APIs and their back-ends are written to.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiAddKeyRequest
*/
func (a *KeysAPIService) AddKey(ctx context.Context) ApiAddKeyRequest {
	return ApiAddKeyRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return SessionState
func (a *KeysAPIService) AddKeyExecute(r ApiAddKeyRequest) (*SessionState, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SessionState
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "KeysAPIService.AddKey")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/keys"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = r.sessionState
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v ApiStatusMessage
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiCreateCustomKeyRequest struct {
	ctx context.Context
	ApiService KeysAPI
	keyID string
	sessionState *SessionState
}

func (r ApiCreateCustomKeyRequest) SessionState(sessionState SessionState) ApiCreateCustomKeyRequest {
	r.sessionState = &sessionState
	return r
}

func (r ApiCreateCustomKeyRequest) Execute() (*SessionState, *http.Response, error) {
	return r.ApiService.CreateCustomKeyExecute(r)
}

/*
CreateCustomKey Create custom key

Creates a key with a custom key ID.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param keyID The Key ID.
 @return ApiCreateCustomKeyRequest
*/
func (a *KeysAPIService) CreateCustomKey(ctx context.Context, keyID string) ApiCreateCustomKeyRequest {
	return ApiCreateCustomKeyRequest{
		ApiService: a,
		ctx: ctx,
		keyID: keyID,
	}
}

// Execute executes the request
//  @return SessionState
func (a *KeysAPIService) CreateCustomKeyExecute(r ApiCreateCustomKeyRequest) (*SessionState, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SessionState
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "KeysAPIService.CreateCustomKey")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/keys/{keyID}"
	localVarPath = strings.Replace(localVarPath, "{"+"keyID"+"}", url.PathEscape(parameterValueToString(r.keyID, "keyID")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = r.sessionState
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v ApiStatusMessage
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiDeleteKeyRequest struct {
	ctx context.Context
	ApiService KeysAPI
	apiID string
	keyID string
}

func (r ApiDeleteKeyRequest) Execute() (*ApiStatusMessage, *http.Response, error) {
	return r.ApiService.DeleteKeyExecute(r)
}

/*
DeleteKey Delete key

Deleting a key will remove it permanently from the system, however analytics relating to that key will still be available.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param apiID ID of API the keys grant access to. Can either be the internal or external API ID.
 @param keyID The Key ID.
 @return ApiDeleteKeyRequest
*/
func (a *KeysAPIService) DeleteKey(ctx context.Context, apiID string, keyID string) ApiDeleteKeyRequest {
	return ApiDeleteKeyRequest{
		ApiService: a,
		ctx: ctx,
		apiID: apiID,
		keyID: keyID,
	}
}

// Execute executes the request
//  @return ApiStatusMessage
func (a *KeysAPIService) DeleteKeyExecute(r ApiDeleteKeyRequest) (*ApiStatusMessage, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ApiStatusMessage
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "KeysAPIService.DeleteKey")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/apis/{apiID}/keys/{keyID}"
	localVarPath = strings.Replace(localVarPath, "{"+"apiID"+"}", url.PathEscape(parameterValueToString(r.apiID, "apiID")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"keyID"+"}", url.PathEscape(parameterValueToString(r.keyID, "keyID")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v ApiStatusMessage
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetKeyRequest struct {
	ctx context.Context
	ApiService KeysAPI
	apiID string
	keyID string
}

func (r ApiGetKeyRequest) Execute() (*SessionState, *http.Response, error) {
	return r.ApiService.GetKeyExecute(r)
}

/*
GetKey Get key

Fetches the key that grant access to the API with the ID {apiID} and key ID {keyID}.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param apiID ID of API the keys grant access to. Can either be the internal or external API ID.
 @param keyID The Key ID.
 @return ApiGetKeyRequest
*/
func (a *KeysAPIService) GetKey(ctx context.Context, apiID string, keyID string) ApiGetKeyRequest {
	return ApiGetKeyRequest{
		ApiService: a,
		ctx: ctx,
		apiID: apiID,
		keyID: keyID,
	}
}

// Execute executes the request
//  @return SessionState
func (a *KeysAPIService) GetKeyExecute(r ApiGetKeyRequest) (*SessionState, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SessionState
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "KeysAPIService.GetKey")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/apis/{apiID}/keys/{keyID}"
	localVarPath = strings.Replace(localVarPath, "{"+"apiID"+"}", url.PathEscape(parameterValueToString(r.apiID, "apiID")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"keyID"+"}", url.PathEscape(parameterValueToString(r.keyID, "keyID")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiListKeysRequest struct {
	ctx context.Context
	ApiService KeysAPI
	apiID string
}

func (r ApiListKeysRequest) Execute() (*ApiAllKeys, *http.Response, error) {
	return r.ApiService.ListKeysExecute(r)
}

/*
ListKeys List keys

Lists keys that grant access to the API with the ID {apiID}.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param apiID ID of API the keys grant access to. Can either be the internal or external API ID.
 @return ApiListKeysRequest
*/
func (a *KeysAPIService) ListKeys(ctx context.Context, apiID string) ApiListKeysRequest {
	return ApiListKeysRequest{
		ApiService: a,
		ctx: ctx,
		apiID: apiID,
	}
}

// Execute executes the request
//  @return ApiAllKeys
func (a *KeysAPIService) ListKeysExecute(r ApiListKeysRequest) (*ApiAllKeys, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ApiAllKeys
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "KeysAPIService.ListKeys")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/apis/{apiID}/keys"
	localVarPath = strings.Replace(localVarPath, "{"+"apiID"+"}", url.PathEscape(parameterValueToString(r.apiID, "apiID")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiUpdateKeyRequest struct {
	ctx context.Context
	ApiService KeysAPI
	apiID string
	keyID string
	sessionState *SessionState
}

func (r ApiUpdateKeyRequest) SessionState(sessionState SessionState) ApiUpdateKeyRequest {
	r.sessionState = &sessionState
	return r
}

func (r ApiUpdateKeyRequest) Execute() (*SessionState, *http.Response, error) {
	return r.ApiService.UpdateKeyExecute(r)
}

/*
UpdateKey Update key

You can also manually add keys to Tyk using your own key-generation algorithm. It is recommended if using this approach to ensure that the OrgID being used in the API Definition and the key data is blank so that Tyk does not try to prepend or manage the key in any way.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param apiID ID of API the keys grant access to. Can either be the internal or external API ID.
 @param keyID The Key ID.
 @return ApiUpdateKeyRequest
*/
func (a *KeysAPIService) UpdateKey(ctx context.Context, apiID string, keyID string) ApiUpdateKeyRequest {
	return ApiUpdateKeyRequest{
		ApiService: a,
		ctx: ctx,
		apiID: apiID,
		keyID: keyID,
	}
}

// Execute executes the request
//  @return SessionState
func (a *KeysAPIService) UpdateKeyExecute(r ApiUpdateKeyRequest) (*SessionState, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SessionState
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "KeysAPIService.UpdateKey")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/apis/{apiID}/keys/{keyID}"
	localVarPath = strings.Replace(localVarPath, "{"+"apiID"+"}", url.PathEscape(parameterValueToString(r.apiID, "apiID")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"keyID"+"}", url.PathEscape(parameterValueToString(r.keyID, "keyID")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = r.sessionState
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v ApiStatusMessage
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
