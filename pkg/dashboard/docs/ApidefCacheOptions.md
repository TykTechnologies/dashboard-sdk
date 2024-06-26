# ApidefCacheOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CacheAllSafeRequests** | Pointer to **bool** |  | [optional] 
**CacheByHeaders** | Pointer to **[]string** |  | [optional] 
**CacheControlTtlHeader** | Pointer to **string** |  | [optional] 
**CacheResponseCodes** | Pointer to **[]int32** |  | [optional] 
**CacheTimeout** | Pointer to **int32** |  | [optional] 
**EnableCache** | Pointer to **bool** |  | [optional] 
**EnableUpstreamCacheControl** | Pointer to **bool** |  | [optional] 

## Methods

### NewApidefCacheOptions

`func NewApidefCacheOptions() *ApidefCacheOptions`

NewApidefCacheOptions instantiates a new ApidefCacheOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApidefCacheOptionsWithDefaults

`func NewApidefCacheOptionsWithDefaults() *ApidefCacheOptions`

NewApidefCacheOptionsWithDefaults instantiates a new ApidefCacheOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCacheAllSafeRequests

`func (o *ApidefCacheOptions) GetCacheAllSafeRequests() bool`

GetCacheAllSafeRequests returns the CacheAllSafeRequests field if non-nil, zero value otherwise.

### GetCacheAllSafeRequestsOk

`func (o *ApidefCacheOptions) GetCacheAllSafeRequestsOk() (*bool, bool)`

GetCacheAllSafeRequestsOk returns a tuple with the CacheAllSafeRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheAllSafeRequests

`func (o *ApidefCacheOptions) SetCacheAllSafeRequests(v bool)`

SetCacheAllSafeRequests sets CacheAllSafeRequests field to given value.

### HasCacheAllSafeRequests

`func (o *ApidefCacheOptions) HasCacheAllSafeRequests() bool`

HasCacheAllSafeRequests returns a boolean if a field has been set.

### GetCacheByHeaders

`func (o *ApidefCacheOptions) GetCacheByHeaders() []string`

GetCacheByHeaders returns the CacheByHeaders field if non-nil, zero value otherwise.

### GetCacheByHeadersOk

`func (o *ApidefCacheOptions) GetCacheByHeadersOk() (*[]string, bool)`

GetCacheByHeadersOk returns a tuple with the CacheByHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheByHeaders

`func (o *ApidefCacheOptions) SetCacheByHeaders(v []string)`

SetCacheByHeaders sets CacheByHeaders field to given value.

### HasCacheByHeaders

`func (o *ApidefCacheOptions) HasCacheByHeaders() bool`

HasCacheByHeaders returns a boolean if a field has been set.

### SetCacheByHeadersNil

`func (o *ApidefCacheOptions) SetCacheByHeadersNil(b bool)`

 SetCacheByHeadersNil sets the value for CacheByHeaders to be an explicit nil

### UnsetCacheByHeaders
`func (o *ApidefCacheOptions) UnsetCacheByHeaders()`

UnsetCacheByHeaders ensures that no value is present for CacheByHeaders, not even an explicit nil
### GetCacheControlTtlHeader

`func (o *ApidefCacheOptions) GetCacheControlTtlHeader() string`

GetCacheControlTtlHeader returns the CacheControlTtlHeader field if non-nil, zero value otherwise.

### GetCacheControlTtlHeaderOk

`func (o *ApidefCacheOptions) GetCacheControlTtlHeaderOk() (*string, bool)`

GetCacheControlTtlHeaderOk returns a tuple with the CacheControlTtlHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheControlTtlHeader

`func (o *ApidefCacheOptions) SetCacheControlTtlHeader(v string)`

SetCacheControlTtlHeader sets CacheControlTtlHeader field to given value.

### HasCacheControlTtlHeader

`func (o *ApidefCacheOptions) HasCacheControlTtlHeader() bool`

HasCacheControlTtlHeader returns a boolean if a field has been set.

### GetCacheResponseCodes

`func (o *ApidefCacheOptions) GetCacheResponseCodes() []int32`

GetCacheResponseCodes returns the CacheResponseCodes field if non-nil, zero value otherwise.

### GetCacheResponseCodesOk

`func (o *ApidefCacheOptions) GetCacheResponseCodesOk() (*[]int32, bool)`

GetCacheResponseCodesOk returns a tuple with the CacheResponseCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheResponseCodes

`func (o *ApidefCacheOptions) SetCacheResponseCodes(v []int32)`

SetCacheResponseCodes sets CacheResponseCodes field to given value.

### HasCacheResponseCodes

`func (o *ApidefCacheOptions) HasCacheResponseCodes() bool`

HasCacheResponseCodes returns a boolean if a field has been set.

### SetCacheResponseCodesNil

`func (o *ApidefCacheOptions) SetCacheResponseCodesNil(b bool)`

 SetCacheResponseCodesNil sets the value for CacheResponseCodes to be an explicit nil

### UnsetCacheResponseCodes
`func (o *ApidefCacheOptions) UnsetCacheResponseCodes()`

UnsetCacheResponseCodes ensures that no value is present for CacheResponseCodes, not even an explicit nil
### GetCacheTimeout

`func (o *ApidefCacheOptions) GetCacheTimeout() int32`

GetCacheTimeout returns the CacheTimeout field if non-nil, zero value otherwise.

### GetCacheTimeoutOk

`func (o *ApidefCacheOptions) GetCacheTimeoutOk() (*int32, bool)`

GetCacheTimeoutOk returns a tuple with the CacheTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheTimeout

`func (o *ApidefCacheOptions) SetCacheTimeout(v int32)`

SetCacheTimeout sets CacheTimeout field to given value.

### HasCacheTimeout

`func (o *ApidefCacheOptions) HasCacheTimeout() bool`

HasCacheTimeout returns a boolean if a field has been set.

### GetEnableCache

`func (o *ApidefCacheOptions) GetEnableCache() bool`

GetEnableCache returns the EnableCache field if non-nil, zero value otherwise.

### GetEnableCacheOk

`func (o *ApidefCacheOptions) GetEnableCacheOk() (*bool, bool)`

GetEnableCacheOk returns a tuple with the EnableCache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableCache

`func (o *ApidefCacheOptions) SetEnableCache(v bool)`

SetEnableCache sets EnableCache field to given value.

### HasEnableCache

`func (o *ApidefCacheOptions) HasEnableCache() bool`

HasEnableCache returns a boolean if a field has been set.

### GetEnableUpstreamCacheControl

`func (o *ApidefCacheOptions) GetEnableUpstreamCacheControl() bool`

GetEnableUpstreamCacheControl returns the EnableUpstreamCacheControl field if non-nil, zero value otherwise.

### GetEnableUpstreamCacheControlOk

`func (o *ApidefCacheOptions) GetEnableUpstreamCacheControlOk() (*bool, bool)`

GetEnableUpstreamCacheControlOk returns a tuple with the EnableUpstreamCacheControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableUpstreamCacheControl

`func (o *ApidefCacheOptions) SetEnableUpstreamCacheControl(v bool)`

SetEnableUpstreamCacheControl sets EnableUpstreamCacheControl field to given value.

### HasEnableUpstreamCacheControl

`func (o *ApidefCacheOptions) HasEnableUpstreamCacheControl() bool`

HasEnableUpstreamCacheControl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


