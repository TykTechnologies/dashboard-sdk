# ApidefCacheMeta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CacheKeyRegex** | Pointer to **string** |  | [optional] 
**CacheResponseCodes** | Pointer to **[]int32** |  | [optional] 
**Disabled** | Pointer to **bool** |  | [optional] 
**Method** | Pointer to **string** |  | [optional] 
**Path** | Pointer to **string** |  | [optional] 
**Timeout** | Pointer to **int32** |  | [optional] 

## Methods

### NewApidefCacheMeta

`func NewApidefCacheMeta() *ApidefCacheMeta`

NewApidefCacheMeta instantiates a new ApidefCacheMeta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApidefCacheMetaWithDefaults

`func NewApidefCacheMetaWithDefaults() *ApidefCacheMeta`

NewApidefCacheMetaWithDefaults instantiates a new ApidefCacheMeta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCacheKeyRegex

`func (o *ApidefCacheMeta) GetCacheKeyRegex() string`

GetCacheKeyRegex returns the CacheKeyRegex field if non-nil, zero value otherwise.

### GetCacheKeyRegexOk

`func (o *ApidefCacheMeta) GetCacheKeyRegexOk() (*string, bool)`

GetCacheKeyRegexOk returns a tuple with the CacheKeyRegex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheKeyRegex

`func (o *ApidefCacheMeta) SetCacheKeyRegex(v string)`

SetCacheKeyRegex sets CacheKeyRegex field to given value.

### HasCacheKeyRegex

`func (o *ApidefCacheMeta) HasCacheKeyRegex() bool`

HasCacheKeyRegex returns a boolean if a field has been set.

### GetCacheResponseCodes

`func (o *ApidefCacheMeta) GetCacheResponseCodes() []int32`

GetCacheResponseCodes returns the CacheResponseCodes field if non-nil, zero value otherwise.

### GetCacheResponseCodesOk

`func (o *ApidefCacheMeta) GetCacheResponseCodesOk() (*[]int32, bool)`

GetCacheResponseCodesOk returns a tuple with the CacheResponseCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheResponseCodes

`func (o *ApidefCacheMeta) SetCacheResponseCodes(v []int32)`

SetCacheResponseCodes sets CacheResponseCodes field to given value.

### HasCacheResponseCodes

`func (o *ApidefCacheMeta) HasCacheResponseCodes() bool`

HasCacheResponseCodes returns a boolean if a field has been set.

### SetCacheResponseCodesNil

`func (o *ApidefCacheMeta) SetCacheResponseCodesNil(b bool)`

 SetCacheResponseCodesNil sets the value for CacheResponseCodes to be an explicit nil

### UnsetCacheResponseCodes
`func (o *ApidefCacheMeta) UnsetCacheResponseCodes()`

UnsetCacheResponseCodes ensures that no value is present for CacheResponseCodes, not even an explicit nil
### GetDisabled

`func (o *ApidefCacheMeta) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *ApidefCacheMeta) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *ApidefCacheMeta) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *ApidefCacheMeta) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetMethod

`func (o *ApidefCacheMeta) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *ApidefCacheMeta) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *ApidefCacheMeta) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *ApidefCacheMeta) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetPath

`func (o *ApidefCacheMeta) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *ApidefCacheMeta) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *ApidefCacheMeta) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *ApidefCacheMeta) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetTimeout

`func (o *ApidefCacheMeta) GetTimeout() int32`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *ApidefCacheMeta) GetTimeoutOk() (*int32, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *ApidefCacheMeta) SetTimeout(v int32)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *ApidefCacheMeta) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


