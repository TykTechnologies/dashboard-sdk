/*
Tyk Dashboard API

 ## <a name=\"introduction\"></a> Introduction  The Tyk Dashboard API offers granular, programmatic access to a centralised database of resources that your Tyk nodes can pull from. This API has a dynamic user administrative structure which means the secret key that is used to communicate with your Tyk nodes can be kept secret and access to the wider management functions can be handled on a user-by-user and organisation-by-organisation basis.  A common question around using a database-backed configuration is how to programmatically add API definitions to your Tyk nodes, the Dashboard API allows much more fine-grained, secure and multi-user access to your Tyk cluster, and should be used to manage a database-backed Tyk node.  The Tyk Dashboard API works seamlessly with the Tyk Dashboard (and the two come bundled together).  ## <a name=\"security-hierarchy\"></a> Security Hierarchy  The Dashboard API provides a more structured security layer to managing Tyk nodes.  ### Organisations, APIs and Users  With the Dashboard API and a database-backed Tyk setup, (and to an extent with file-based API setups - if diligence is used in naming and creating definitions), the following security model is applied to the management of Upstream APIs:  * **Organisations**: All APIs are *owned* by an organisation, this is designated by the 'OrgID' parameter in the API Definition. * **Users**: All users created in the Dashboard belong to an organisation (unless an exception is made for super-administrative access). * **APIs**: All APIs belong to an Organisation and only Users that belong to that organisation can see the analytics for those APIs and manage their configurations. * **API Keys**: API Keys are designated by organisation, this means an API key that has full access rights will not be allowed to access the APIs of another organisation on the same system, but can have full access to all APIs within the organisation. * **Access Rights**: Access rights are stored with the key, this enables a key to give access to multiple APIs, this is defined by the session object in the core Tyk API.  In order to use the Dashboard API, you'll need to get the 'Tyk Dashboard API Access Credentials' secret from your user profile on the Dashboard UI.  The secret you set should then be sent along as a header with each Dashboard API Request in order for it to be successful:   authorization: <your-secret>

API version: 5.6.0
Contact: support@tyk.io
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

type AnalyticsAPI interface {

	/*
		GetAnalyticsOfApiKey Analytics of API Key.

		It returns analytics of the endpoints of all APIs called a key between start and end date.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param startDay Day to start querying the analytics from.
		@param startMonth Month to start querying the analytics from.
		@param startYear Year to start querying the analytics from.
		@param endDay End date of analytics to query.
		@param endMonth End month of analytics to query.
		@param endYear End year of analytics to query.
		@param keyHash Hash of your API key.
		@return ApiGetAnalyticsOfApiKeyRequest
	*/
	GetAnalyticsOfApiKey(ctx context.Context, startDay string, startMonth string, startYear string, endDay string, endMonth string, endYear string, keyHash string) ApiGetAnalyticsOfApiKeyRequest

	// GetAnalyticsOfApiKeyExecute executes the request
	//  @return AggregateAnalyticsData
	GetAnalyticsOfApiKeyExecute(r ApiGetAnalyticsOfApiKeyRequest) (*AggregateAnalyticsData, *http.Response, error)

	/*
		GetAnalyticsOfOauthClientId Analytics of Oauth Client ID.

		Returns activity of all endpoints which used the given OAuth client between the given time range.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param startDay Day to start querying the analytics from.
		@param startMonth Month to start querying the analytics from.
		@param startYear Year to start querying the analytics from.
		@param endDay End date of analytics to query.
		@param endMonth End month of analytics to query.
		@param endYear End year of analytics to query.
		@param oAuthClientID OAuthClientID
		@return ApiGetAnalyticsOfOauthClientIdRequest
	*/
	GetAnalyticsOfOauthClientId(ctx context.Context, startDay string, startMonth string, startYear string, endDay string, endMonth string, endYear string, oAuthClientID string) ApiGetAnalyticsOfOauthClientIdRequest

	// GetAnalyticsOfOauthClientIdExecute executes the request
	//  @return AggregateAnalyticsData
	GetAnalyticsOfOauthClientIdExecute(r ApiGetAnalyticsOfOauthClientIdRequest) (*AggregateAnalyticsData, *http.Response, error)
}

// AnalyticsAPIService AnalyticsAPI service
type AnalyticsAPIService service

type ApiGetAnalyticsOfApiKeyRequest struct {
	ctx        context.Context
	ApiService AnalyticsAPI
	startDay   string
	startMonth string
	startYear  string
	endDay     string
	endMonth   string
	endYear    string
	keyHash    string
}

func (r ApiGetAnalyticsOfApiKeyRequest) Execute() (*AggregateAnalyticsData, *http.Response, error) {
	return r.ApiService.GetAnalyticsOfApiKeyExecute(r)
}

/*
GetAnalyticsOfApiKey Analytics of API Key.

It returns analytics of the endpoints of all APIs called a key between start and end date.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param startDay Day to start querying the analytics from.
	@param startMonth Month to start querying the analytics from.
	@param startYear Year to start querying the analytics from.
	@param endDay End date of analytics to query.
	@param endMonth End month of analytics to query.
	@param endYear End year of analytics to query.
	@param keyHash Hash of your API key.
	@return ApiGetAnalyticsOfApiKeyRequest
*/
func (a *AnalyticsAPIService) GetAnalyticsOfApiKey(ctx context.Context, startDay string, startMonth string, startYear string, endDay string, endMonth string, endYear string, keyHash string) ApiGetAnalyticsOfApiKeyRequest {
	return ApiGetAnalyticsOfApiKeyRequest{
		ApiService: a,
		ctx:        ctx,
		startDay:   startDay,
		startMonth: startMonth,
		startYear:  startYear,
		endDay:     endDay,
		endMonth:   endMonth,
		endYear:    endYear,
		keyHash:    keyHash,
	}
}

// Execute executes the request
//
//	@return AggregateAnalyticsData
func (a *AnalyticsAPIService) GetAnalyticsOfApiKeyExecute(r ApiGetAnalyticsOfApiKeyRequest) (*AggregateAnalyticsData, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AggregateAnalyticsData
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AnalyticsAPIService.GetAnalyticsOfApiKey")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/activity/keys/{keyHash}/{startDay}/{startMonth}/{startYear}/{EndDay}/{EndMonth}/{EndYear}"
	localVarPath = strings.Replace(localVarPath, "{"+"startDay"+"}", url.PathEscape(parameterValueToString(r.startDay, "startDay")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"startMonth"+"}", url.PathEscape(parameterValueToString(r.startMonth, "startMonth")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"startYear"+"}", url.PathEscape(parameterValueToString(r.startYear, "startYear")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"EndDay"+"}", url.PathEscape(parameterValueToString(r.endDay, "endDay")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"EndMonth"+"}", url.PathEscape(parameterValueToString(r.endMonth, "endMonth")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"EndYear"+"}", url.PathEscape(parameterValueToString(r.endYear, "endYear")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"keyHash"+"}", url.PathEscape(parameterValueToString(r.keyHash, "keyHash")), -1)

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
		if localVarHTTPResponse.StatusCode == 400 {
			var v ApiResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v ApiResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v ApiResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v ApiResponse
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

type ApiGetAnalyticsOfOauthClientIdRequest struct {
	ctx           context.Context
	ApiService    AnalyticsAPI
	startDay      string
	startMonth    string
	startYear     string
	endDay        string
	endMonth      string
	endYear       string
	oAuthClientID string
}

func (r ApiGetAnalyticsOfOauthClientIdRequest) Execute() (*AggregateAnalyticsData, *http.Response, error) {
	return r.ApiService.GetAnalyticsOfOauthClientIdExecute(r)
}

/*
GetAnalyticsOfOauthClientId Analytics of Oauth Client ID.

Returns activity of all endpoints which used the given OAuth client between the given time range.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param startDay Day to start querying the analytics from.
	@param startMonth Month to start querying the analytics from.
	@param startYear Year to start querying the analytics from.
	@param endDay End date of analytics to query.
	@param endMonth End month of analytics to query.
	@param endYear End year of analytics to query.
	@param oAuthClientID OAuthClientID
	@return ApiGetAnalyticsOfOauthClientIdRequest
*/
func (a *AnalyticsAPIService) GetAnalyticsOfOauthClientId(ctx context.Context, startDay string, startMonth string, startYear string, endDay string, endMonth string, endYear string, oAuthClientID string) ApiGetAnalyticsOfOauthClientIdRequest {
	return ApiGetAnalyticsOfOauthClientIdRequest{
		ApiService:    a,
		ctx:           ctx,
		startDay:      startDay,
		startMonth:    startMonth,
		startYear:     startYear,
		endDay:        endDay,
		endMonth:      endMonth,
		endYear:       endYear,
		oAuthClientID: oAuthClientID,
	}
}

// Execute executes the request
//
//	@return AggregateAnalyticsData
func (a *AnalyticsAPIService) GetAnalyticsOfOauthClientIdExecute(r ApiGetAnalyticsOfOauthClientIdRequest) (*AggregateAnalyticsData, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AggregateAnalyticsData
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AnalyticsAPIService.GetAnalyticsOfOauthClientId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/activity/oauthid/{OAuthClientID}/{startDay}/{startMonth}/{startYear}/{EndDay}/{EndMonth}/{EndYear}"
	localVarPath = strings.Replace(localVarPath, "{"+"startDay"+"}", url.PathEscape(parameterValueToString(r.startDay, "startDay")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"startMonth"+"}", url.PathEscape(parameterValueToString(r.startMonth, "startMonth")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"startYear"+"}", url.PathEscape(parameterValueToString(r.startYear, "startYear")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"EndDay"+"}", url.PathEscape(parameterValueToString(r.endDay, "endDay")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"EndMonth"+"}", url.PathEscape(parameterValueToString(r.endMonth, "endMonth")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"EndYear"+"}", url.PathEscape(parameterValueToString(r.endYear, "endYear")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"OAuthClientID"+"}", url.PathEscape(parameterValueToString(r.oAuthClientID, "oAuthClientID")), -1)

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
			var v ApiResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v ApiResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v ApiResponse
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
