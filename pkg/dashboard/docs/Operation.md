# Operation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Allow** | Pointer to [**NullableAllowance**](Allowance.md) |  | [optional] 
**Block** | Pointer to [**NullableAllowance**](Allowance.md) |  | [optional] 
**Cache** | Pointer to [**NullableCachePlugin**](CachePlugin.md) |  | [optional] 
**CircuitBreaker** | Pointer to [**NullableCircuitBreaker**](CircuitBreaker.md) |  | [optional] 
**DoNotTrackEndpoint** | Pointer to [**NullableTrackEndpoint**](TrackEndpoint.md) |  | [optional] 
**EnforceTimeout** | Pointer to [**NullableEnforceTimeout**](EnforceTimeout.md) |  | [optional] 
**IgnoreAuthentication** | Pointer to [**NullableAllowance**](Allowance.md) |  | [optional] 
**Internal** | Pointer to [**NullableInternal**](Internal.md) |  | [optional] 
**MockResponse** | Pointer to [**NullableMockResponse**](MockResponse.md) |  | [optional] 
**PostPlugins** | Pointer to [**[]EndpointPostPlugin**](EndpointPostPlugin.md) |  | [optional] 
**RateLimit** | Pointer to [**NullableRateLimitEndpoint**](RateLimitEndpoint.md) |  | [optional] 
**RequestSizeLimit** | Pointer to [**NullableRequestSizeLimit**](RequestSizeLimit.md) |  | [optional] 
**TrackEndpoint** | Pointer to [**NullableTrackEndpoint**](TrackEndpoint.md) |  | [optional] 
**TransformRequestBody** | Pointer to [**NullableTransformBody**](TransformBody.md) |  | [optional] 
**TransformRequestHeaders** | Pointer to [**NullableTransformHeaders**](TransformHeaders.md) |  | [optional] 
**TransformRequestMethod** | Pointer to [**NullableTransformRequestMethod**](TransformRequestMethod.md) |  | [optional] 
**TransformResponseBody** | Pointer to [**NullableTransformBody**](TransformBody.md) |  | [optional] 
**TransformResponseHeaders** | Pointer to [**NullableTransformHeaders**](TransformHeaders.md) |  | [optional] 
**UrlRewrite** | Pointer to [**NullableURLRewrite**](URLRewrite.md) |  | [optional] 
**ValidateRequest** | Pointer to [**NullableValidateRequest**](ValidateRequest.md) |  | [optional] 
**VirtualEndpoint** | Pointer to [**NullableVirtualEndpoint**](VirtualEndpoint.md) |  | [optional] 

## Methods

### NewOperation

`func NewOperation() *Operation`

NewOperation instantiates a new Operation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOperationWithDefaults

`func NewOperationWithDefaults() *Operation`

NewOperationWithDefaults instantiates a new Operation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllow

`func (o *Operation) GetAllow() Allowance`

GetAllow returns the Allow field if non-nil, zero value otherwise.

### GetAllowOk

`func (o *Operation) GetAllowOk() (*Allowance, bool)`

GetAllowOk returns a tuple with the Allow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllow

`func (o *Operation) SetAllow(v Allowance)`

SetAllow sets Allow field to given value.

### HasAllow

`func (o *Operation) HasAllow() bool`

HasAllow returns a boolean if a field has been set.

### SetAllowNil

`func (o *Operation) SetAllowNil(b bool)`

 SetAllowNil sets the value for Allow to be an explicit nil

### UnsetAllow
`func (o *Operation) UnsetAllow()`

UnsetAllow ensures that no value is present for Allow, not even an explicit nil
### GetBlock

`func (o *Operation) GetBlock() Allowance`

GetBlock returns the Block field if non-nil, zero value otherwise.

### GetBlockOk

`func (o *Operation) GetBlockOk() (*Allowance, bool)`

GetBlockOk returns a tuple with the Block field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlock

`func (o *Operation) SetBlock(v Allowance)`

SetBlock sets Block field to given value.

### HasBlock

`func (o *Operation) HasBlock() bool`

HasBlock returns a boolean if a field has been set.

### SetBlockNil

`func (o *Operation) SetBlockNil(b bool)`

 SetBlockNil sets the value for Block to be an explicit nil

### UnsetBlock
`func (o *Operation) UnsetBlock()`

UnsetBlock ensures that no value is present for Block, not even an explicit nil
### GetCache

`func (o *Operation) GetCache() CachePlugin`

GetCache returns the Cache field if non-nil, zero value otherwise.

### GetCacheOk

`func (o *Operation) GetCacheOk() (*CachePlugin, bool)`

GetCacheOk returns a tuple with the Cache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCache

`func (o *Operation) SetCache(v CachePlugin)`

SetCache sets Cache field to given value.

### HasCache

`func (o *Operation) HasCache() bool`

HasCache returns a boolean if a field has been set.

### SetCacheNil

`func (o *Operation) SetCacheNil(b bool)`

 SetCacheNil sets the value for Cache to be an explicit nil

### UnsetCache
`func (o *Operation) UnsetCache()`

UnsetCache ensures that no value is present for Cache, not even an explicit nil
### GetCircuitBreaker

`func (o *Operation) GetCircuitBreaker() CircuitBreaker`

GetCircuitBreaker returns the CircuitBreaker field if non-nil, zero value otherwise.

### GetCircuitBreakerOk

`func (o *Operation) GetCircuitBreakerOk() (*CircuitBreaker, bool)`

GetCircuitBreakerOk returns a tuple with the CircuitBreaker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCircuitBreaker

`func (o *Operation) SetCircuitBreaker(v CircuitBreaker)`

SetCircuitBreaker sets CircuitBreaker field to given value.

### HasCircuitBreaker

`func (o *Operation) HasCircuitBreaker() bool`

HasCircuitBreaker returns a boolean if a field has been set.

### SetCircuitBreakerNil

`func (o *Operation) SetCircuitBreakerNil(b bool)`

 SetCircuitBreakerNil sets the value for CircuitBreaker to be an explicit nil

### UnsetCircuitBreaker
`func (o *Operation) UnsetCircuitBreaker()`

UnsetCircuitBreaker ensures that no value is present for CircuitBreaker, not even an explicit nil
### GetDoNotTrackEndpoint

`func (o *Operation) GetDoNotTrackEndpoint() TrackEndpoint`

GetDoNotTrackEndpoint returns the DoNotTrackEndpoint field if non-nil, zero value otherwise.

### GetDoNotTrackEndpointOk

`func (o *Operation) GetDoNotTrackEndpointOk() (*TrackEndpoint, bool)`

GetDoNotTrackEndpointOk returns a tuple with the DoNotTrackEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoNotTrackEndpoint

`func (o *Operation) SetDoNotTrackEndpoint(v TrackEndpoint)`

SetDoNotTrackEndpoint sets DoNotTrackEndpoint field to given value.

### HasDoNotTrackEndpoint

`func (o *Operation) HasDoNotTrackEndpoint() bool`

HasDoNotTrackEndpoint returns a boolean if a field has been set.

### SetDoNotTrackEndpointNil

`func (o *Operation) SetDoNotTrackEndpointNil(b bool)`

 SetDoNotTrackEndpointNil sets the value for DoNotTrackEndpoint to be an explicit nil

### UnsetDoNotTrackEndpoint
`func (o *Operation) UnsetDoNotTrackEndpoint()`

UnsetDoNotTrackEndpoint ensures that no value is present for DoNotTrackEndpoint, not even an explicit nil
### GetEnforceTimeout

`func (o *Operation) GetEnforceTimeout() EnforceTimeout`

GetEnforceTimeout returns the EnforceTimeout field if non-nil, zero value otherwise.

### GetEnforceTimeoutOk

`func (o *Operation) GetEnforceTimeoutOk() (*EnforceTimeout, bool)`

GetEnforceTimeoutOk returns a tuple with the EnforceTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnforceTimeout

`func (o *Operation) SetEnforceTimeout(v EnforceTimeout)`

SetEnforceTimeout sets EnforceTimeout field to given value.

### HasEnforceTimeout

`func (o *Operation) HasEnforceTimeout() bool`

HasEnforceTimeout returns a boolean if a field has been set.

### SetEnforceTimeoutNil

`func (o *Operation) SetEnforceTimeoutNil(b bool)`

 SetEnforceTimeoutNil sets the value for EnforceTimeout to be an explicit nil

### UnsetEnforceTimeout
`func (o *Operation) UnsetEnforceTimeout()`

UnsetEnforceTimeout ensures that no value is present for EnforceTimeout, not even an explicit nil
### GetIgnoreAuthentication

`func (o *Operation) GetIgnoreAuthentication() Allowance`

GetIgnoreAuthentication returns the IgnoreAuthentication field if non-nil, zero value otherwise.

### GetIgnoreAuthenticationOk

`func (o *Operation) GetIgnoreAuthenticationOk() (*Allowance, bool)`

GetIgnoreAuthenticationOk returns a tuple with the IgnoreAuthentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreAuthentication

`func (o *Operation) SetIgnoreAuthentication(v Allowance)`

SetIgnoreAuthentication sets IgnoreAuthentication field to given value.

### HasIgnoreAuthentication

`func (o *Operation) HasIgnoreAuthentication() bool`

HasIgnoreAuthentication returns a boolean if a field has been set.

### SetIgnoreAuthenticationNil

`func (o *Operation) SetIgnoreAuthenticationNil(b bool)`

 SetIgnoreAuthenticationNil sets the value for IgnoreAuthentication to be an explicit nil

### UnsetIgnoreAuthentication
`func (o *Operation) UnsetIgnoreAuthentication()`

UnsetIgnoreAuthentication ensures that no value is present for IgnoreAuthentication, not even an explicit nil
### GetInternal

`func (o *Operation) GetInternal() Internal`

GetInternal returns the Internal field if non-nil, zero value otherwise.

### GetInternalOk

`func (o *Operation) GetInternalOk() (*Internal, bool)`

GetInternalOk returns a tuple with the Internal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternal

`func (o *Operation) SetInternal(v Internal)`

SetInternal sets Internal field to given value.

### HasInternal

`func (o *Operation) HasInternal() bool`

HasInternal returns a boolean if a field has been set.

### SetInternalNil

`func (o *Operation) SetInternalNil(b bool)`

 SetInternalNil sets the value for Internal to be an explicit nil

### UnsetInternal
`func (o *Operation) UnsetInternal()`

UnsetInternal ensures that no value is present for Internal, not even an explicit nil
### GetMockResponse

`func (o *Operation) GetMockResponse() MockResponse`

GetMockResponse returns the MockResponse field if non-nil, zero value otherwise.

### GetMockResponseOk

`func (o *Operation) GetMockResponseOk() (*MockResponse, bool)`

GetMockResponseOk returns a tuple with the MockResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMockResponse

`func (o *Operation) SetMockResponse(v MockResponse)`

SetMockResponse sets MockResponse field to given value.

### HasMockResponse

`func (o *Operation) HasMockResponse() bool`

HasMockResponse returns a boolean if a field has been set.

### SetMockResponseNil

`func (o *Operation) SetMockResponseNil(b bool)`

 SetMockResponseNil sets the value for MockResponse to be an explicit nil

### UnsetMockResponse
`func (o *Operation) UnsetMockResponse()`

UnsetMockResponse ensures that no value is present for MockResponse, not even an explicit nil
### GetPostPlugins

`func (o *Operation) GetPostPlugins() []EndpointPostPlugin`

GetPostPlugins returns the PostPlugins field if non-nil, zero value otherwise.

### GetPostPluginsOk

`func (o *Operation) GetPostPluginsOk() (*[]EndpointPostPlugin, bool)`

GetPostPluginsOk returns a tuple with the PostPlugins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostPlugins

`func (o *Operation) SetPostPlugins(v []EndpointPostPlugin)`

SetPostPlugins sets PostPlugins field to given value.

### HasPostPlugins

`func (o *Operation) HasPostPlugins() bool`

HasPostPlugins returns a boolean if a field has been set.

### GetRateLimit

`func (o *Operation) GetRateLimit() RateLimitEndpoint`

GetRateLimit returns the RateLimit field if non-nil, zero value otherwise.

### GetRateLimitOk

`func (o *Operation) GetRateLimitOk() (*RateLimitEndpoint, bool)`

GetRateLimitOk returns a tuple with the RateLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateLimit

`func (o *Operation) SetRateLimit(v RateLimitEndpoint)`

SetRateLimit sets RateLimit field to given value.

### HasRateLimit

`func (o *Operation) HasRateLimit() bool`

HasRateLimit returns a boolean if a field has been set.

### SetRateLimitNil

`func (o *Operation) SetRateLimitNil(b bool)`

 SetRateLimitNil sets the value for RateLimit to be an explicit nil

### UnsetRateLimit
`func (o *Operation) UnsetRateLimit()`

UnsetRateLimit ensures that no value is present for RateLimit, not even an explicit nil
### GetRequestSizeLimit

`func (o *Operation) GetRequestSizeLimit() RequestSizeLimit`

GetRequestSizeLimit returns the RequestSizeLimit field if non-nil, zero value otherwise.

### GetRequestSizeLimitOk

`func (o *Operation) GetRequestSizeLimitOk() (*RequestSizeLimit, bool)`

GetRequestSizeLimitOk returns a tuple with the RequestSizeLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestSizeLimit

`func (o *Operation) SetRequestSizeLimit(v RequestSizeLimit)`

SetRequestSizeLimit sets RequestSizeLimit field to given value.

### HasRequestSizeLimit

`func (o *Operation) HasRequestSizeLimit() bool`

HasRequestSizeLimit returns a boolean if a field has been set.

### SetRequestSizeLimitNil

`func (o *Operation) SetRequestSizeLimitNil(b bool)`

 SetRequestSizeLimitNil sets the value for RequestSizeLimit to be an explicit nil

### UnsetRequestSizeLimit
`func (o *Operation) UnsetRequestSizeLimit()`

UnsetRequestSizeLimit ensures that no value is present for RequestSizeLimit, not even an explicit nil
### GetTrackEndpoint

`func (o *Operation) GetTrackEndpoint() TrackEndpoint`

GetTrackEndpoint returns the TrackEndpoint field if non-nil, zero value otherwise.

### GetTrackEndpointOk

`func (o *Operation) GetTrackEndpointOk() (*TrackEndpoint, bool)`

GetTrackEndpointOk returns a tuple with the TrackEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackEndpoint

`func (o *Operation) SetTrackEndpoint(v TrackEndpoint)`

SetTrackEndpoint sets TrackEndpoint field to given value.

### HasTrackEndpoint

`func (o *Operation) HasTrackEndpoint() bool`

HasTrackEndpoint returns a boolean if a field has been set.

### SetTrackEndpointNil

`func (o *Operation) SetTrackEndpointNil(b bool)`

 SetTrackEndpointNil sets the value for TrackEndpoint to be an explicit nil

### UnsetTrackEndpoint
`func (o *Operation) UnsetTrackEndpoint()`

UnsetTrackEndpoint ensures that no value is present for TrackEndpoint, not even an explicit nil
### GetTransformRequestBody

`func (o *Operation) GetTransformRequestBody() TransformBody`

GetTransformRequestBody returns the TransformRequestBody field if non-nil, zero value otherwise.

### GetTransformRequestBodyOk

`func (o *Operation) GetTransformRequestBodyOk() (*TransformBody, bool)`

GetTransformRequestBodyOk returns a tuple with the TransformRequestBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransformRequestBody

`func (o *Operation) SetTransformRequestBody(v TransformBody)`

SetTransformRequestBody sets TransformRequestBody field to given value.

### HasTransformRequestBody

`func (o *Operation) HasTransformRequestBody() bool`

HasTransformRequestBody returns a boolean if a field has been set.

### SetTransformRequestBodyNil

`func (o *Operation) SetTransformRequestBodyNil(b bool)`

 SetTransformRequestBodyNil sets the value for TransformRequestBody to be an explicit nil

### UnsetTransformRequestBody
`func (o *Operation) UnsetTransformRequestBody()`

UnsetTransformRequestBody ensures that no value is present for TransformRequestBody, not even an explicit nil
### GetTransformRequestHeaders

`func (o *Operation) GetTransformRequestHeaders() TransformHeaders`

GetTransformRequestHeaders returns the TransformRequestHeaders field if non-nil, zero value otherwise.

### GetTransformRequestHeadersOk

`func (o *Operation) GetTransformRequestHeadersOk() (*TransformHeaders, bool)`

GetTransformRequestHeadersOk returns a tuple with the TransformRequestHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransformRequestHeaders

`func (o *Operation) SetTransformRequestHeaders(v TransformHeaders)`

SetTransformRequestHeaders sets TransformRequestHeaders field to given value.

### HasTransformRequestHeaders

`func (o *Operation) HasTransformRequestHeaders() bool`

HasTransformRequestHeaders returns a boolean if a field has been set.

### SetTransformRequestHeadersNil

`func (o *Operation) SetTransformRequestHeadersNil(b bool)`

 SetTransformRequestHeadersNil sets the value for TransformRequestHeaders to be an explicit nil

### UnsetTransformRequestHeaders
`func (o *Operation) UnsetTransformRequestHeaders()`

UnsetTransformRequestHeaders ensures that no value is present for TransformRequestHeaders, not even an explicit nil
### GetTransformRequestMethod

`func (o *Operation) GetTransformRequestMethod() TransformRequestMethod`

GetTransformRequestMethod returns the TransformRequestMethod field if non-nil, zero value otherwise.

### GetTransformRequestMethodOk

`func (o *Operation) GetTransformRequestMethodOk() (*TransformRequestMethod, bool)`

GetTransformRequestMethodOk returns a tuple with the TransformRequestMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransformRequestMethod

`func (o *Operation) SetTransformRequestMethod(v TransformRequestMethod)`

SetTransformRequestMethod sets TransformRequestMethod field to given value.

### HasTransformRequestMethod

`func (o *Operation) HasTransformRequestMethod() bool`

HasTransformRequestMethod returns a boolean if a field has been set.

### SetTransformRequestMethodNil

`func (o *Operation) SetTransformRequestMethodNil(b bool)`

 SetTransformRequestMethodNil sets the value for TransformRequestMethod to be an explicit nil

### UnsetTransformRequestMethod
`func (o *Operation) UnsetTransformRequestMethod()`

UnsetTransformRequestMethod ensures that no value is present for TransformRequestMethod, not even an explicit nil
### GetTransformResponseBody

`func (o *Operation) GetTransformResponseBody() TransformBody`

GetTransformResponseBody returns the TransformResponseBody field if non-nil, zero value otherwise.

### GetTransformResponseBodyOk

`func (o *Operation) GetTransformResponseBodyOk() (*TransformBody, bool)`

GetTransformResponseBodyOk returns a tuple with the TransformResponseBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransformResponseBody

`func (o *Operation) SetTransformResponseBody(v TransformBody)`

SetTransformResponseBody sets TransformResponseBody field to given value.

### HasTransformResponseBody

`func (o *Operation) HasTransformResponseBody() bool`

HasTransformResponseBody returns a boolean if a field has been set.

### SetTransformResponseBodyNil

`func (o *Operation) SetTransformResponseBodyNil(b bool)`

 SetTransformResponseBodyNil sets the value for TransformResponseBody to be an explicit nil

### UnsetTransformResponseBody
`func (o *Operation) UnsetTransformResponseBody()`

UnsetTransformResponseBody ensures that no value is present for TransformResponseBody, not even an explicit nil
### GetTransformResponseHeaders

`func (o *Operation) GetTransformResponseHeaders() TransformHeaders`

GetTransformResponseHeaders returns the TransformResponseHeaders field if non-nil, zero value otherwise.

### GetTransformResponseHeadersOk

`func (o *Operation) GetTransformResponseHeadersOk() (*TransformHeaders, bool)`

GetTransformResponseHeadersOk returns a tuple with the TransformResponseHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransformResponseHeaders

`func (o *Operation) SetTransformResponseHeaders(v TransformHeaders)`

SetTransformResponseHeaders sets TransformResponseHeaders field to given value.

### HasTransformResponseHeaders

`func (o *Operation) HasTransformResponseHeaders() bool`

HasTransformResponseHeaders returns a boolean if a field has been set.

### SetTransformResponseHeadersNil

`func (o *Operation) SetTransformResponseHeadersNil(b bool)`

 SetTransformResponseHeadersNil sets the value for TransformResponseHeaders to be an explicit nil

### UnsetTransformResponseHeaders
`func (o *Operation) UnsetTransformResponseHeaders()`

UnsetTransformResponseHeaders ensures that no value is present for TransformResponseHeaders, not even an explicit nil
### GetUrlRewrite

`func (o *Operation) GetUrlRewrite() URLRewrite`

GetUrlRewrite returns the UrlRewrite field if non-nil, zero value otherwise.

### GetUrlRewriteOk

`func (o *Operation) GetUrlRewriteOk() (*URLRewrite, bool)`

GetUrlRewriteOk returns a tuple with the UrlRewrite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlRewrite

`func (o *Operation) SetUrlRewrite(v URLRewrite)`

SetUrlRewrite sets UrlRewrite field to given value.

### HasUrlRewrite

`func (o *Operation) HasUrlRewrite() bool`

HasUrlRewrite returns a boolean if a field has been set.

### SetUrlRewriteNil

`func (o *Operation) SetUrlRewriteNil(b bool)`

 SetUrlRewriteNil sets the value for UrlRewrite to be an explicit nil

### UnsetUrlRewrite
`func (o *Operation) UnsetUrlRewrite()`

UnsetUrlRewrite ensures that no value is present for UrlRewrite, not even an explicit nil
### GetValidateRequest

`func (o *Operation) GetValidateRequest() ValidateRequest`

GetValidateRequest returns the ValidateRequest field if non-nil, zero value otherwise.

### GetValidateRequestOk

`func (o *Operation) GetValidateRequestOk() (*ValidateRequest, bool)`

GetValidateRequestOk returns a tuple with the ValidateRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidateRequest

`func (o *Operation) SetValidateRequest(v ValidateRequest)`

SetValidateRequest sets ValidateRequest field to given value.

### HasValidateRequest

`func (o *Operation) HasValidateRequest() bool`

HasValidateRequest returns a boolean if a field has been set.

### SetValidateRequestNil

`func (o *Operation) SetValidateRequestNil(b bool)`

 SetValidateRequestNil sets the value for ValidateRequest to be an explicit nil

### UnsetValidateRequest
`func (o *Operation) UnsetValidateRequest()`

UnsetValidateRequest ensures that no value is present for ValidateRequest, not even an explicit nil
### GetVirtualEndpoint

`func (o *Operation) GetVirtualEndpoint() VirtualEndpoint`

GetVirtualEndpoint returns the VirtualEndpoint field if non-nil, zero value otherwise.

### GetVirtualEndpointOk

`func (o *Operation) GetVirtualEndpointOk() (*VirtualEndpoint, bool)`

GetVirtualEndpointOk returns a tuple with the VirtualEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualEndpoint

`func (o *Operation) SetVirtualEndpoint(v VirtualEndpoint)`

SetVirtualEndpoint sets VirtualEndpoint field to given value.

### HasVirtualEndpoint

`func (o *Operation) HasVirtualEndpoint() bool`

HasVirtualEndpoint returns a boolean if a field has been set.

### SetVirtualEndpointNil

`func (o *Operation) SetVirtualEndpointNil(b bool)`

 SetVirtualEndpointNil sets the value for VirtualEndpoint to be an explicit nil

### UnsetVirtualEndpoint
`func (o *Operation) UnsetVirtualEndpoint()`

UnsetVirtualEndpoint ensures that no value is present for VirtualEndpoint, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


