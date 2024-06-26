# OasOperation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Allow** | Pointer to [**OasAllowance**](OasAllowance.md) |  | [optional] 
**Block** | Pointer to [**OasAllowance**](OasAllowance.md) |  | [optional] 
**Cache** | Pointer to [**OasCachePlugin**](OasCachePlugin.md) |  | [optional] 
**CircuitBreaker** | Pointer to [**OasCircuitBreaker**](OasCircuitBreaker.md) |  | [optional] 
**DoNotTrackEndpoint** | Pointer to [**OasTrackEndpoint**](OasTrackEndpoint.md) |  | [optional] 
**EnforceTimeout** | Pointer to [**OasEnforceTimeout**](OasEnforceTimeout.md) |  | [optional] 
**IgnoreAuthentication** | Pointer to [**OasAllowance**](OasAllowance.md) |  | [optional] 
**Internal** | Pointer to [**OasInternal**](OasInternal.md) |  | [optional] 
**MockResponse** | Pointer to [**OasMockResponse**](OasMockResponse.md) |  | [optional] 
**PostPlugins** | Pointer to [**[]OasEndpointPostPlugin**](OasEndpointPostPlugin.md) |  | [optional] 
**RequestSizeLimit** | Pointer to [**OasRequestSizeLimit**](OasRequestSizeLimit.md) |  | [optional] 
**TrackEndpoint** | Pointer to [**OasTrackEndpoint**](OasTrackEndpoint.md) |  | [optional] 
**TransformRequestBody** | Pointer to [**OasTransformBody**](OasTransformBody.md) |  | [optional] 
**TransformRequestHeaders** | Pointer to [**OasTransformHeaders**](OasTransformHeaders.md) |  | [optional] 
**TransformRequestMethod** | Pointer to [**OasTransformRequestMethod**](OasTransformRequestMethod.md) |  | [optional] 
**TransformResponseBody** | Pointer to [**OasTransformBody**](OasTransformBody.md) |  | [optional] 
**TransformResponseHeaders** | Pointer to [**OasTransformHeaders**](OasTransformHeaders.md) |  | [optional] 
**UrlRewrite** | Pointer to [**OasURLRewrite**](OasURLRewrite.md) |  | [optional] 
**ValidateRequest** | Pointer to [**OasValidateRequest**](OasValidateRequest.md) |  | [optional] 
**VirtualEndpoint** | Pointer to [**OasVirtualEndpoint**](OasVirtualEndpoint.md) |  | [optional] 

## Methods

### NewOasOperation

`func NewOasOperation() *OasOperation`

NewOasOperation instantiates a new OasOperation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOasOperationWithDefaults

`func NewOasOperationWithDefaults() *OasOperation`

NewOasOperationWithDefaults instantiates a new OasOperation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllow

`func (o *OasOperation) GetAllow() OasAllowance`

GetAllow returns the Allow field if non-nil, zero value otherwise.

### GetAllowOk

`func (o *OasOperation) GetAllowOk() (*OasAllowance, bool)`

GetAllowOk returns a tuple with the Allow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllow

`func (o *OasOperation) SetAllow(v OasAllowance)`

SetAllow sets Allow field to given value.

### HasAllow

`func (o *OasOperation) HasAllow() bool`

HasAllow returns a boolean if a field has been set.

### GetBlock

`func (o *OasOperation) GetBlock() OasAllowance`

GetBlock returns the Block field if non-nil, zero value otherwise.

### GetBlockOk

`func (o *OasOperation) GetBlockOk() (*OasAllowance, bool)`

GetBlockOk returns a tuple with the Block field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlock

`func (o *OasOperation) SetBlock(v OasAllowance)`

SetBlock sets Block field to given value.

### HasBlock

`func (o *OasOperation) HasBlock() bool`

HasBlock returns a boolean if a field has been set.

### GetCache

`func (o *OasOperation) GetCache() OasCachePlugin`

GetCache returns the Cache field if non-nil, zero value otherwise.

### GetCacheOk

`func (o *OasOperation) GetCacheOk() (*OasCachePlugin, bool)`

GetCacheOk returns a tuple with the Cache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCache

`func (o *OasOperation) SetCache(v OasCachePlugin)`

SetCache sets Cache field to given value.

### HasCache

`func (o *OasOperation) HasCache() bool`

HasCache returns a boolean if a field has been set.

### GetCircuitBreaker

`func (o *OasOperation) GetCircuitBreaker() OasCircuitBreaker`

GetCircuitBreaker returns the CircuitBreaker field if non-nil, zero value otherwise.

### GetCircuitBreakerOk

`func (o *OasOperation) GetCircuitBreakerOk() (*OasCircuitBreaker, bool)`

GetCircuitBreakerOk returns a tuple with the CircuitBreaker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCircuitBreaker

`func (o *OasOperation) SetCircuitBreaker(v OasCircuitBreaker)`

SetCircuitBreaker sets CircuitBreaker field to given value.

### HasCircuitBreaker

`func (o *OasOperation) HasCircuitBreaker() bool`

HasCircuitBreaker returns a boolean if a field has been set.

### GetDoNotTrackEndpoint

`func (o *OasOperation) GetDoNotTrackEndpoint() OasTrackEndpoint`

GetDoNotTrackEndpoint returns the DoNotTrackEndpoint field if non-nil, zero value otherwise.

### GetDoNotTrackEndpointOk

`func (o *OasOperation) GetDoNotTrackEndpointOk() (*OasTrackEndpoint, bool)`

GetDoNotTrackEndpointOk returns a tuple with the DoNotTrackEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoNotTrackEndpoint

`func (o *OasOperation) SetDoNotTrackEndpoint(v OasTrackEndpoint)`

SetDoNotTrackEndpoint sets DoNotTrackEndpoint field to given value.

### HasDoNotTrackEndpoint

`func (o *OasOperation) HasDoNotTrackEndpoint() bool`

HasDoNotTrackEndpoint returns a boolean if a field has been set.

### GetEnforceTimeout

`func (o *OasOperation) GetEnforceTimeout() OasEnforceTimeout`

GetEnforceTimeout returns the EnforceTimeout field if non-nil, zero value otherwise.

### GetEnforceTimeoutOk

`func (o *OasOperation) GetEnforceTimeoutOk() (*OasEnforceTimeout, bool)`

GetEnforceTimeoutOk returns a tuple with the EnforceTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnforceTimeout

`func (o *OasOperation) SetEnforceTimeout(v OasEnforceTimeout)`

SetEnforceTimeout sets EnforceTimeout field to given value.

### HasEnforceTimeout

`func (o *OasOperation) HasEnforceTimeout() bool`

HasEnforceTimeout returns a boolean if a field has been set.

### GetIgnoreAuthentication

`func (o *OasOperation) GetIgnoreAuthentication() OasAllowance`

GetIgnoreAuthentication returns the IgnoreAuthentication field if non-nil, zero value otherwise.

### GetIgnoreAuthenticationOk

`func (o *OasOperation) GetIgnoreAuthenticationOk() (*OasAllowance, bool)`

GetIgnoreAuthenticationOk returns a tuple with the IgnoreAuthentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreAuthentication

`func (o *OasOperation) SetIgnoreAuthentication(v OasAllowance)`

SetIgnoreAuthentication sets IgnoreAuthentication field to given value.

### HasIgnoreAuthentication

`func (o *OasOperation) HasIgnoreAuthentication() bool`

HasIgnoreAuthentication returns a boolean if a field has been set.

### GetInternal

`func (o *OasOperation) GetInternal() OasInternal`

GetInternal returns the Internal field if non-nil, zero value otherwise.

### GetInternalOk

`func (o *OasOperation) GetInternalOk() (*OasInternal, bool)`

GetInternalOk returns a tuple with the Internal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternal

`func (o *OasOperation) SetInternal(v OasInternal)`

SetInternal sets Internal field to given value.

### HasInternal

`func (o *OasOperation) HasInternal() bool`

HasInternal returns a boolean if a field has been set.

### GetMockResponse

`func (o *OasOperation) GetMockResponse() OasMockResponse`

GetMockResponse returns the MockResponse field if non-nil, zero value otherwise.

### GetMockResponseOk

`func (o *OasOperation) GetMockResponseOk() (*OasMockResponse, bool)`

GetMockResponseOk returns a tuple with the MockResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMockResponse

`func (o *OasOperation) SetMockResponse(v OasMockResponse)`

SetMockResponse sets MockResponse field to given value.

### HasMockResponse

`func (o *OasOperation) HasMockResponse() bool`

HasMockResponse returns a boolean if a field has been set.

### GetPostPlugins

`func (o *OasOperation) GetPostPlugins() []OasEndpointPostPlugin`

GetPostPlugins returns the PostPlugins field if non-nil, zero value otherwise.

### GetPostPluginsOk

`func (o *OasOperation) GetPostPluginsOk() (*[]OasEndpointPostPlugin, bool)`

GetPostPluginsOk returns a tuple with the PostPlugins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostPlugins

`func (o *OasOperation) SetPostPlugins(v []OasEndpointPostPlugin)`

SetPostPlugins sets PostPlugins field to given value.

### HasPostPlugins

`func (o *OasOperation) HasPostPlugins() bool`

HasPostPlugins returns a boolean if a field has been set.

### GetRequestSizeLimit

`func (o *OasOperation) GetRequestSizeLimit() OasRequestSizeLimit`

GetRequestSizeLimit returns the RequestSizeLimit field if non-nil, zero value otherwise.

### GetRequestSizeLimitOk

`func (o *OasOperation) GetRequestSizeLimitOk() (*OasRequestSizeLimit, bool)`

GetRequestSizeLimitOk returns a tuple with the RequestSizeLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestSizeLimit

`func (o *OasOperation) SetRequestSizeLimit(v OasRequestSizeLimit)`

SetRequestSizeLimit sets RequestSizeLimit field to given value.

### HasRequestSizeLimit

`func (o *OasOperation) HasRequestSizeLimit() bool`

HasRequestSizeLimit returns a boolean if a field has been set.

### GetTrackEndpoint

`func (o *OasOperation) GetTrackEndpoint() OasTrackEndpoint`

GetTrackEndpoint returns the TrackEndpoint field if non-nil, zero value otherwise.

### GetTrackEndpointOk

`func (o *OasOperation) GetTrackEndpointOk() (*OasTrackEndpoint, bool)`

GetTrackEndpointOk returns a tuple with the TrackEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackEndpoint

`func (o *OasOperation) SetTrackEndpoint(v OasTrackEndpoint)`

SetTrackEndpoint sets TrackEndpoint field to given value.

### HasTrackEndpoint

`func (o *OasOperation) HasTrackEndpoint() bool`

HasTrackEndpoint returns a boolean if a field has been set.

### GetTransformRequestBody

`func (o *OasOperation) GetTransformRequestBody() OasTransformBody`

GetTransformRequestBody returns the TransformRequestBody field if non-nil, zero value otherwise.

### GetTransformRequestBodyOk

`func (o *OasOperation) GetTransformRequestBodyOk() (*OasTransformBody, bool)`

GetTransformRequestBodyOk returns a tuple with the TransformRequestBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransformRequestBody

`func (o *OasOperation) SetTransformRequestBody(v OasTransformBody)`

SetTransformRequestBody sets TransformRequestBody field to given value.

### HasTransformRequestBody

`func (o *OasOperation) HasTransformRequestBody() bool`

HasTransformRequestBody returns a boolean if a field has been set.

### GetTransformRequestHeaders

`func (o *OasOperation) GetTransformRequestHeaders() OasTransformHeaders`

GetTransformRequestHeaders returns the TransformRequestHeaders field if non-nil, zero value otherwise.

### GetTransformRequestHeadersOk

`func (o *OasOperation) GetTransformRequestHeadersOk() (*OasTransformHeaders, bool)`

GetTransformRequestHeadersOk returns a tuple with the TransformRequestHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransformRequestHeaders

`func (o *OasOperation) SetTransformRequestHeaders(v OasTransformHeaders)`

SetTransformRequestHeaders sets TransformRequestHeaders field to given value.

### HasTransformRequestHeaders

`func (o *OasOperation) HasTransformRequestHeaders() bool`

HasTransformRequestHeaders returns a boolean if a field has been set.

### GetTransformRequestMethod

`func (o *OasOperation) GetTransformRequestMethod() OasTransformRequestMethod`

GetTransformRequestMethod returns the TransformRequestMethod field if non-nil, zero value otherwise.

### GetTransformRequestMethodOk

`func (o *OasOperation) GetTransformRequestMethodOk() (*OasTransformRequestMethod, bool)`

GetTransformRequestMethodOk returns a tuple with the TransformRequestMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransformRequestMethod

`func (o *OasOperation) SetTransformRequestMethod(v OasTransformRequestMethod)`

SetTransformRequestMethod sets TransformRequestMethod field to given value.

### HasTransformRequestMethod

`func (o *OasOperation) HasTransformRequestMethod() bool`

HasTransformRequestMethod returns a boolean if a field has been set.

### GetTransformResponseBody

`func (o *OasOperation) GetTransformResponseBody() OasTransformBody`

GetTransformResponseBody returns the TransformResponseBody field if non-nil, zero value otherwise.

### GetTransformResponseBodyOk

`func (o *OasOperation) GetTransformResponseBodyOk() (*OasTransformBody, bool)`

GetTransformResponseBodyOk returns a tuple with the TransformResponseBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransformResponseBody

`func (o *OasOperation) SetTransformResponseBody(v OasTransformBody)`

SetTransformResponseBody sets TransformResponseBody field to given value.

### HasTransformResponseBody

`func (o *OasOperation) HasTransformResponseBody() bool`

HasTransformResponseBody returns a boolean if a field has been set.

### GetTransformResponseHeaders

`func (o *OasOperation) GetTransformResponseHeaders() OasTransformHeaders`

GetTransformResponseHeaders returns the TransformResponseHeaders field if non-nil, zero value otherwise.

### GetTransformResponseHeadersOk

`func (o *OasOperation) GetTransformResponseHeadersOk() (*OasTransformHeaders, bool)`

GetTransformResponseHeadersOk returns a tuple with the TransformResponseHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransformResponseHeaders

`func (o *OasOperation) SetTransformResponseHeaders(v OasTransformHeaders)`

SetTransformResponseHeaders sets TransformResponseHeaders field to given value.

### HasTransformResponseHeaders

`func (o *OasOperation) HasTransformResponseHeaders() bool`

HasTransformResponseHeaders returns a boolean if a field has been set.

### GetUrlRewrite

`func (o *OasOperation) GetUrlRewrite() OasURLRewrite`

GetUrlRewrite returns the UrlRewrite field if non-nil, zero value otherwise.

### GetUrlRewriteOk

`func (o *OasOperation) GetUrlRewriteOk() (*OasURLRewrite, bool)`

GetUrlRewriteOk returns a tuple with the UrlRewrite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlRewrite

`func (o *OasOperation) SetUrlRewrite(v OasURLRewrite)`

SetUrlRewrite sets UrlRewrite field to given value.

### HasUrlRewrite

`func (o *OasOperation) HasUrlRewrite() bool`

HasUrlRewrite returns a boolean if a field has been set.

### GetValidateRequest

`func (o *OasOperation) GetValidateRequest() OasValidateRequest`

GetValidateRequest returns the ValidateRequest field if non-nil, zero value otherwise.

### GetValidateRequestOk

`func (o *OasOperation) GetValidateRequestOk() (*OasValidateRequest, bool)`

GetValidateRequestOk returns a tuple with the ValidateRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidateRequest

`func (o *OasOperation) SetValidateRequest(v OasValidateRequest)`

SetValidateRequest sets ValidateRequest field to given value.

### HasValidateRequest

`func (o *OasOperation) HasValidateRequest() bool`

HasValidateRequest returns a boolean if a field has been set.

### GetVirtualEndpoint

`func (o *OasOperation) GetVirtualEndpoint() OasVirtualEndpoint`

GetVirtualEndpoint returns the VirtualEndpoint field if non-nil, zero value otherwise.

### GetVirtualEndpointOk

`func (o *OasOperation) GetVirtualEndpointOk() (*OasVirtualEndpoint, bool)`

GetVirtualEndpointOk returns a tuple with the VirtualEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualEndpoint

`func (o *OasOperation) SetVirtualEndpoint(v OasVirtualEndpoint)`

SetVirtualEndpoint sets VirtualEndpoint field to given value.

### HasVirtualEndpoint

`func (o *OasOperation) HasVirtualEndpoint() bool`

HasVirtualEndpoint returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


