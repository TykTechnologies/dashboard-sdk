# ApidefExtendedPathsSet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdvanceCacheConfig** | Pointer to [**[]ApidefCacheMeta**](ApidefCacheMeta.md) |  | [optional] 
**BlackList** | Pointer to [**[]ApidefEndPointMeta**](ApidefEndPointMeta.md) |  | [optional] 
**Cache** | Pointer to **[]string** |  | [optional] 
**CircuitBreakers** | Pointer to [**[]ApidefCircuitBreakerMeta**](ApidefCircuitBreakerMeta.md) |  | [optional] 
**DoNotTrackEndpoints** | Pointer to [**[]ApidefTrackEndpointMeta**](ApidefTrackEndpointMeta.md) |  | [optional] 
**GoPlugin** | Pointer to [**[]ApidefGoPluginMeta**](ApidefGoPluginMeta.md) |  | [optional] 
**HardTimeouts** | Pointer to [**[]ApidefHardTimeoutMeta**](ApidefHardTimeoutMeta.md) |  | [optional] 
**Ignored** | Pointer to [**[]ApidefEndPointMeta**](ApidefEndPointMeta.md) |  | [optional] 
**Internal** | Pointer to [**[]ApidefInternalMeta**](ApidefInternalMeta.md) |  | [optional] 
**MethodTransforms** | Pointer to [**[]ApidefMethodTransformMeta**](ApidefMethodTransformMeta.md) |  | [optional] 
**MockResponse** | Pointer to [**[]ApidefMockResponseMeta**](ApidefMockResponseMeta.md) |  | [optional] 
**PersistGraphql** | Pointer to [**[]ApidefPersistGraphQLMeta**](ApidefPersistGraphQLMeta.md) |  | [optional] 
**SizeLimits** | Pointer to [**[]ApidefRequestSizeMeta**](ApidefRequestSizeMeta.md) |  | [optional] 
**TrackEndpoints** | Pointer to [**[]ApidefTrackEndpointMeta**](ApidefTrackEndpointMeta.md) |  | [optional] 
**Transform** | Pointer to [**[]ApidefTemplateMeta**](ApidefTemplateMeta.md) |  | [optional] 
**TransformHeaders** | Pointer to [**[]ApidefHeaderInjectionMeta**](ApidefHeaderInjectionMeta.md) |  | [optional] 
**TransformJq** | Pointer to [**[]ApidefTransformJQMeta**](ApidefTransformJQMeta.md) |  | [optional] 
**TransformJqResponse** | Pointer to [**[]ApidefTransformJQMeta**](ApidefTransformJQMeta.md) |  | [optional] 
**TransformResponse** | Pointer to [**[]ApidefTemplateMeta**](ApidefTemplateMeta.md) |  | [optional] 
**TransformResponseHeaders** | Pointer to [**[]ApidefHeaderInjectionMeta**](ApidefHeaderInjectionMeta.md) |  | [optional] 
**UrlRewrites** | Pointer to [**[]ApidefURLRewriteMeta**](ApidefURLRewriteMeta.md) |  | [optional] 
**ValidateJson** | Pointer to [**[]ApidefValidatePathMeta**](ApidefValidatePathMeta.md) |  | [optional] 
**ValidateRequest** | Pointer to [**[]ApidefValidateRequestMeta**](ApidefValidateRequestMeta.md) |  | [optional] 
**Virtual** | Pointer to [**[]ApidefVirtualMeta**](ApidefVirtualMeta.md) |  | [optional] 
**WhiteList** | Pointer to [**[]ApidefEndPointMeta**](ApidefEndPointMeta.md) |  | [optional] 

## Methods

### NewApidefExtendedPathsSet

`func NewApidefExtendedPathsSet() *ApidefExtendedPathsSet`

NewApidefExtendedPathsSet instantiates a new ApidefExtendedPathsSet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApidefExtendedPathsSetWithDefaults

`func NewApidefExtendedPathsSetWithDefaults() *ApidefExtendedPathsSet`

NewApidefExtendedPathsSetWithDefaults instantiates a new ApidefExtendedPathsSet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdvanceCacheConfig

`func (o *ApidefExtendedPathsSet) GetAdvanceCacheConfig() []ApidefCacheMeta`

GetAdvanceCacheConfig returns the AdvanceCacheConfig field if non-nil, zero value otherwise.

### GetAdvanceCacheConfigOk

`func (o *ApidefExtendedPathsSet) GetAdvanceCacheConfigOk() (*[]ApidefCacheMeta, bool)`

GetAdvanceCacheConfigOk returns a tuple with the AdvanceCacheConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvanceCacheConfig

`func (o *ApidefExtendedPathsSet) SetAdvanceCacheConfig(v []ApidefCacheMeta)`

SetAdvanceCacheConfig sets AdvanceCacheConfig field to given value.

### HasAdvanceCacheConfig

`func (o *ApidefExtendedPathsSet) HasAdvanceCacheConfig() bool`

HasAdvanceCacheConfig returns a boolean if a field has been set.

### GetBlackList

`func (o *ApidefExtendedPathsSet) GetBlackList() []ApidefEndPointMeta`

GetBlackList returns the BlackList field if non-nil, zero value otherwise.

### GetBlackListOk

`func (o *ApidefExtendedPathsSet) GetBlackListOk() (*[]ApidefEndPointMeta, bool)`

GetBlackListOk returns a tuple with the BlackList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlackList

`func (o *ApidefExtendedPathsSet) SetBlackList(v []ApidefEndPointMeta)`

SetBlackList sets BlackList field to given value.

### HasBlackList

`func (o *ApidefExtendedPathsSet) HasBlackList() bool`

HasBlackList returns a boolean if a field has been set.

### GetCache

`func (o *ApidefExtendedPathsSet) GetCache() []string`

GetCache returns the Cache field if non-nil, zero value otherwise.

### GetCacheOk

`func (o *ApidefExtendedPathsSet) GetCacheOk() (*[]string, bool)`

GetCacheOk returns a tuple with the Cache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCache

`func (o *ApidefExtendedPathsSet) SetCache(v []string)`

SetCache sets Cache field to given value.

### HasCache

`func (o *ApidefExtendedPathsSet) HasCache() bool`

HasCache returns a boolean if a field has been set.

### GetCircuitBreakers

`func (o *ApidefExtendedPathsSet) GetCircuitBreakers() []ApidefCircuitBreakerMeta`

GetCircuitBreakers returns the CircuitBreakers field if non-nil, zero value otherwise.

### GetCircuitBreakersOk

`func (o *ApidefExtendedPathsSet) GetCircuitBreakersOk() (*[]ApidefCircuitBreakerMeta, bool)`

GetCircuitBreakersOk returns a tuple with the CircuitBreakers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCircuitBreakers

`func (o *ApidefExtendedPathsSet) SetCircuitBreakers(v []ApidefCircuitBreakerMeta)`

SetCircuitBreakers sets CircuitBreakers field to given value.

### HasCircuitBreakers

`func (o *ApidefExtendedPathsSet) HasCircuitBreakers() bool`

HasCircuitBreakers returns a boolean if a field has been set.

### GetDoNotTrackEndpoints

`func (o *ApidefExtendedPathsSet) GetDoNotTrackEndpoints() []ApidefTrackEndpointMeta`

GetDoNotTrackEndpoints returns the DoNotTrackEndpoints field if non-nil, zero value otherwise.

### GetDoNotTrackEndpointsOk

`func (o *ApidefExtendedPathsSet) GetDoNotTrackEndpointsOk() (*[]ApidefTrackEndpointMeta, bool)`

GetDoNotTrackEndpointsOk returns a tuple with the DoNotTrackEndpoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoNotTrackEndpoints

`func (o *ApidefExtendedPathsSet) SetDoNotTrackEndpoints(v []ApidefTrackEndpointMeta)`

SetDoNotTrackEndpoints sets DoNotTrackEndpoints field to given value.

### HasDoNotTrackEndpoints

`func (o *ApidefExtendedPathsSet) HasDoNotTrackEndpoints() bool`

HasDoNotTrackEndpoints returns a boolean if a field has been set.

### GetGoPlugin

`func (o *ApidefExtendedPathsSet) GetGoPlugin() []ApidefGoPluginMeta`

GetGoPlugin returns the GoPlugin field if non-nil, zero value otherwise.

### GetGoPluginOk

`func (o *ApidefExtendedPathsSet) GetGoPluginOk() (*[]ApidefGoPluginMeta, bool)`

GetGoPluginOk returns a tuple with the GoPlugin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoPlugin

`func (o *ApidefExtendedPathsSet) SetGoPlugin(v []ApidefGoPluginMeta)`

SetGoPlugin sets GoPlugin field to given value.

### HasGoPlugin

`func (o *ApidefExtendedPathsSet) HasGoPlugin() bool`

HasGoPlugin returns a boolean if a field has been set.

### GetHardTimeouts

`func (o *ApidefExtendedPathsSet) GetHardTimeouts() []ApidefHardTimeoutMeta`

GetHardTimeouts returns the HardTimeouts field if non-nil, zero value otherwise.

### GetHardTimeoutsOk

`func (o *ApidefExtendedPathsSet) GetHardTimeoutsOk() (*[]ApidefHardTimeoutMeta, bool)`

GetHardTimeoutsOk returns a tuple with the HardTimeouts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHardTimeouts

`func (o *ApidefExtendedPathsSet) SetHardTimeouts(v []ApidefHardTimeoutMeta)`

SetHardTimeouts sets HardTimeouts field to given value.

### HasHardTimeouts

`func (o *ApidefExtendedPathsSet) HasHardTimeouts() bool`

HasHardTimeouts returns a boolean if a field has been set.

### GetIgnored

`func (o *ApidefExtendedPathsSet) GetIgnored() []ApidefEndPointMeta`

GetIgnored returns the Ignored field if non-nil, zero value otherwise.

### GetIgnoredOk

`func (o *ApidefExtendedPathsSet) GetIgnoredOk() (*[]ApidefEndPointMeta, bool)`

GetIgnoredOk returns a tuple with the Ignored field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnored

`func (o *ApidefExtendedPathsSet) SetIgnored(v []ApidefEndPointMeta)`

SetIgnored sets Ignored field to given value.

### HasIgnored

`func (o *ApidefExtendedPathsSet) HasIgnored() bool`

HasIgnored returns a boolean if a field has been set.

### GetInternal

`func (o *ApidefExtendedPathsSet) GetInternal() []ApidefInternalMeta`

GetInternal returns the Internal field if non-nil, zero value otherwise.

### GetInternalOk

`func (o *ApidefExtendedPathsSet) GetInternalOk() (*[]ApidefInternalMeta, bool)`

GetInternalOk returns a tuple with the Internal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternal

`func (o *ApidefExtendedPathsSet) SetInternal(v []ApidefInternalMeta)`

SetInternal sets Internal field to given value.

### HasInternal

`func (o *ApidefExtendedPathsSet) HasInternal() bool`

HasInternal returns a boolean if a field has been set.

### GetMethodTransforms

`func (o *ApidefExtendedPathsSet) GetMethodTransforms() []ApidefMethodTransformMeta`

GetMethodTransforms returns the MethodTransforms field if non-nil, zero value otherwise.

### GetMethodTransformsOk

`func (o *ApidefExtendedPathsSet) GetMethodTransformsOk() (*[]ApidefMethodTransformMeta, bool)`

GetMethodTransformsOk returns a tuple with the MethodTransforms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethodTransforms

`func (o *ApidefExtendedPathsSet) SetMethodTransforms(v []ApidefMethodTransformMeta)`

SetMethodTransforms sets MethodTransforms field to given value.

### HasMethodTransforms

`func (o *ApidefExtendedPathsSet) HasMethodTransforms() bool`

HasMethodTransforms returns a boolean if a field has been set.

### GetMockResponse

`func (o *ApidefExtendedPathsSet) GetMockResponse() []ApidefMockResponseMeta`

GetMockResponse returns the MockResponse field if non-nil, zero value otherwise.

### GetMockResponseOk

`func (o *ApidefExtendedPathsSet) GetMockResponseOk() (*[]ApidefMockResponseMeta, bool)`

GetMockResponseOk returns a tuple with the MockResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMockResponse

`func (o *ApidefExtendedPathsSet) SetMockResponse(v []ApidefMockResponseMeta)`

SetMockResponse sets MockResponse field to given value.

### HasMockResponse

`func (o *ApidefExtendedPathsSet) HasMockResponse() bool`

HasMockResponse returns a boolean if a field has been set.

### GetPersistGraphql

`func (o *ApidefExtendedPathsSet) GetPersistGraphql() []ApidefPersistGraphQLMeta`

GetPersistGraphql returns the PersistGraphql field if non-nil, zero value otherwise.

### GetPersistGraphqlOk

`func (o *ApidefExtendedPathsSet) GetPersistGraphqlOk() (*[]ApidefPersistGraphQLMeta, bool)`

GetPersistGraphqlOk returns a tuple with the PersistGraphql field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersistGraphql

`func (o *ApidefExtendedPathsSet) SetPersistGraphql(v []ApidefPersistGraphQLMeta)`

SetPersistGraphql sets PersistGraphql field to given value.

### HasPersistGraphql

`func (o *ApidefExtendedPathsSet) HasPersistGraphql() bool`

HasPersistGraphql returns a boolean if a field has been set.

### SetPersistGraphqlNil

`func (o *ApidefExtendedPathsSet) SetPersistGraphqlNil(b bool)`

 SetPersistGraphqlNil sets the value for PersistGraphql to be an explicit nil

### UnsetPersistGraphql
`func (o *ApidefExtendedPathsSet) UnsetPersistGraphql()`

UnsetPersistGraphql ensures that no value is present for PersistGraphql, not even an explicit nil
### GetSizeLimits

`func (o *ApidefExtendedPathsSet) GetSizeLimits() []ApidefRequestSizeMeta`

GetSizeLimits returns the SizeLimits field if non-nil, zero value otherwise.

### GetSizeLimitsOk

`func (o *ApidefExtendedPathsSet) GetSizeLimitsOk() (*[]ApidefRequestSizeMeta, bool)`

GetSizeLimitsOk returns a tuple with the SizeLimits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeLimits

`func (o *ApidefExtendedPathsSet) SetSizeLimits(v []ApidefRequestSizeMeta)`

SetSizeLimits sets SizeLimits field to given value.

### HasSizeLimits

`func (o *ApidefExtendedPathsSet) HasSizeLimits() bool`

HasSizeLimits returns a boolean if a field has been set.

### GetTrackEndpoints

`func (o *ApidefExtendedPathsSet) GetTrackEndpoints() []ApidefTrackEndpointMeta`

GetTrackEndpoints returns the TrackEndpoints field if non-nil, zero value otherwise.

### GetTrackEndpointsOk

`func (o *ApidefExtendedPathsSet) GetTrackEndpointsOk() (*[]ApidefTrackEndpointMeta, bool)`

GetTrackEndpointsOk returns a tuple with the TrackEndpoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackEndpoints

`func (o *ApidefExtendedPathsSet) SetTrackEndpoints(v []ApidefTrackEndpointMeta)`

SetTrackEndpoints sets TrackEndpoints field to given value.

### HasTrackEndpoints

`func (o *ApidefExtendedPathsSet) HasTrackEndpoints() bool`

HasTrackEndpoints returns a boolean if a field has been set.

### GetTransform

`func (o *ApidefExtendedPathsSet) GetTransform() []ApidefTemplateMeta`

GetTransform returns the Transform field if non-nil, zero value otherwise.

### GetTransformOk

`func (o *ApidefExtendedPathsSet) GetTransformOk() (*[]ApidefTemplateMeta, bool)`

GetTransformOk returns a tuple with the Transform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransform

`func (o *ApidefExtendedPathsSet) SetTransform(v []ApidefTemplateMeta)`

SetTransform sets Transform field to given value.

### HasTransform

`func (o *ApidefExtendedPathsSet) HasTransform() bool`

HasTransform returns a boolean if a field has been set.

### GetTransformHeaders

`func (o *ApidefExtendedPathsSet) GetTransformHeaders() []ApidefHeaderInjectionMeta`

GetTransformHeaders returns the TransformHeaders field if non-nil, zero value otherwise.

### GetTransformHeadersOk

`func (o *ApidefExtendedPathsSet) GetTransformHeadersOk() (*[]ApidefHeaderInjectionMeta, bool)`

GetTransformHeadersOk returns a tuple with the TransformHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransformHeaders

`func (o *ApidefExtendedPathsSet) SetTransformHeaders(v []ApidefHeaderInjectionMeta)`

SetTransformHeaders sets TransformHeaders field to given value.

### HasTransformHeaders

`func (o *ApidefExtendedPathsSet) HasTransformHeaders() bool`

HasTransformHeaders returns a boolean if a field has been set.

### GetTransformJq

`func (o *ApidefExtendedPathsSet) GetTransformJq() []ApidefTransformJQMeta`

GetTransformJq returns the TransformJq field if non-nil, zero value otherwise.

### GetTransformJqOk

`func (o *ApidefExtendedPathsSet) GetTransformJqOk() (*[]ApidefTransformJQMeta, bool)`

GetTransformJqOk returns a tuple with the TransformJq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransformJq

`func (o *ApidefExtendedPathsSet) SetTransformJq(v []ApidefTransformJQMeta)`

SetTransformJq sets TransformJq field to given value.

### HasTransformJq

`func (o *ApidefExtendedPathsSet) HasTransformJq() bool`

HasTransformJq returns a boolean if a field has been set.

### GetTransformJqResponse

`func (o *ApidefExtendedPathsSet) GetTransformJqResponse() []ApidefTransformJQMeta`

GetTransformJqResponse returns the TransformJqResponse field if non-nil, zero value otherwise.

### GetTransformJqResponseOk

`func (o *ApidefExtendedPathsSet) GetTransformJqResponseOk() (*[]ApidefTransformJQMeta, bool)`

GetTransformJqResponseOk returns a tuple with the TransformJqResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransformJqResponse

`func (o *ApidefExtendedPathsSet) SetTransformJqResponse(v []ApidefTransformJQMeta)`

SetTransformJqResponse sets TransformJqResponse field to given value.

### HasTransformJqResponse

`func (o *ApidefExtendedPathsSet) HasTransformJqResponse() bool`

HasTransformJqResponse returns a boolean if a field has been set.

### GetTransformResponse

`func (o *ApidefExtendedPathsSet) GetTransformResponse() []ApidefTemplateMeta`

GetTransformResponse returns the TransformResponse field if non-nil, zero value otherwise.

### GetTransformResponseOk

`func (o *ApidefExtendedPathsSet) GetTransformResponseOk() (*[]ApidefTemplateMeta, bool)`

GetTransformResponseOk returns a tuple with the TransformResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransformResponse

`func (o *ApidefExtendedPathsSet) SetTransformResponse(v []ApidefTemplateMeta)`

SetTransformResponse sets TransformResponse field to given value.

### HasTransformResponse

`func (o *ApidefExtendedPathsSet) HasTransformResponse() bool`

HasTransformResponse returns a boolean if a field has been set.

### GetTransformResponseHeaders

`func (o *ApidefExtendedPathsSet) GetTransformResponseHeaders() []ApidefHeaderInjectionMeta`

GetTransformResponseHeaders returns the TransformResponseHeaders field if non-nil, zero value otherwise.

### GetTransformResponseHeadersOk

`func (o *ApidefExtendedPathsSet) GetTransformResponseHeadersOk() (*[]ApidefHeaderInjectionMeta, bool)`

GetTransformResponseHeadersOk returns a tuple with the TransformResponseHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransformResponseHeaders

`func (o *ApidefExtendedPathsSet) SetTransformResponseHeaders(v []ApidefHeaderInjectionMeta)`

SetTransformResponseHeaders sets TransformResponseHeaders field to given value.

### HasTransformResponseHeaders

`func (o *ApidefExtendedPathsSet) HasTransformResponseHeaders() bool`

HasTransformResponseHeaders returns a boolean if a field has been set.

### GetUrlRewrites

`func (o *ApidefExtendedPathsSet) GetUrlRewrites() []ApidefURLRewriteMeta`

GetUrlRewrites returns the UrlRewrites field if non-nil, zero value otherwise.

### GetUrlRewritesOk

`func (o *ApidefExtendedPathsSet) GetUrlRewritesOk() (*[]ApidefURLRewriteMeta, bool)`

GetUrlRewritesOk returns a tuple with the UrlRewrites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlRewrites

`func (o *ApidefExtendedPathsSet) SetUrlRewrites(v []ApidefURLRewriteMeta)`

SetUrlRewrites sets UrlRewrites field to given value.

### HasUrlRewrites

`func (o *ApidefExtendedPathsSet) HasUrlRewrites() bool`

HasUrlRewrites returns a boolean if a field has been set.

### GetValidateJson

`func (o *ApidefExtendedPathsSet) GetValidateJson() []ApidefValidatePathMeta`

GetValidateJson returns the ValidateJson field if non-nil, zero value otherwise.

### GetValidateJsonOk

`func (o *ApidefExtendedPathsSet) GetValidateJsonOk() (*[]ApidefValidatePathMeta, bool)`

GetValidateJsonOk returns a tuple with the ValidateJson field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidateJson

`func (o *ApidefExtendedPathsSet) SetValidateJson(v []ApidefValidatePathMeta)`

SetValidateJson sets ValidateJson field to given value.

### HasValidateJson

`func (o *ApidefExtendedPathsSet) HasValidateJson() bool`

HasValidateJson returns a boolean if a field has been set.

### GetValidateRequest

`func (o *ApidefExtendedPathsSet) GetValidateRequest() []ApidefValidateRequestMeta`

GetValidateRequest returns the ValidateRequest field if non-nil, zero value otherwise.

### GetValidateRequestOk

`func (o *ApidefExtendedPathsSet) GetValidateRequestOk() (*[]ApidefValidateRequestMeta, bool)`

GetValidateRequestOk returns a tuple with the ValidateRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidateRequest

`func (o *ApidefExtendedPathsSet) SetValidateRequest(v []ApidefValidateRequestMeta)`

SetValidateRequest sets ValidateRequest field to given value.

### HasValidateRequest

`func (o *ApidefExtendedPathsSet) HasValidateRequest() bool`

HasValidateRequest returns a boolean if a field has been set.

### GetVirtual

`func (o *ApidefExtendedPathsSet) GetVirtual() []ApidefVirtualMeta`

GetVirtual returns the Virtual field if non-nil, zero value otherwise.

### GetVirtualOk

`func (o *ApidefExtendedPathsSet) GetVirtualOk() (*[]ApidefVirtualMeta, bool)`

GetVirtualOk returns a tuple with the Virtual field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtual

`func (o *ApidefExtendedPathsSet) SetVirtual(v []ApidefVirtualMeta)`

SetVirtual sets Virtual field to given value.

### HasVirtual

`func (o *ApidefExtendedPathsSet) HasVirtual() bool`

HasVirtual returns a boolean if a field has been set.

### GetWhiteList

`func (o *ApidefExtendedPathsSet) GetWhiteList() []ApidefEndPointMeta`

GetWhiteList returns the WhiteList field if non-nil, zero value otherwise.

### GetWhiteListOk

`func (o *ApidefExtendedPathsSet) GetWhiteListOk() (*[]ApidefEndPointMeta, bool)`

GetWhiteListOk returns a tuple with the WhiteList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhiteList

`func (o *ApidefExtendedPathsSet) SetWhiteList(v []ApidefEndPointMeta)`

SetWhiteList sets WhiteList field to given value.

### HasWhiteList

`func (o *ApidefExtendedPathsSet) HasWhiteList() bool`

HasWhiteList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


