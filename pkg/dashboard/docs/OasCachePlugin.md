# OasCachePlugin

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CacheByRegex** | Pointer to **string** |  | [optional] 
**CacheResponseCodes** | Pointer to **[]int32** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Timeout** | Pointer to **int32** |  | [optional] 

## Methods

### NewOasCachePlugin

`func NewOasCachePlugin() *OasCachePlugin`

NewOasCachePlugin instantiates a new OasCachePlugin object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOasCachePluginWithDefaults

`func NewOasCachePluginWithDefaults() *OasCachePlugin`

NewOasCachePluginWithDefaults instantiates a new OasCachePlugin object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCacheByRegex

`func (o *OasCachePlugin) GetCacheByRegex() string`

GetCacheByRegex returns the CacheByRegex field if non-nil, zero value otherwise.

### GetCacheByRegexOk

`func (o *OasCachePlugin) GetCacheByRegexOk() (*string, bool)`

GetCacheByRegexOk returns a tuple with the CacheByRegex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheByRegex

`func (o *OasCachePlugin) SetCacheByRegex(v string)`

SetCacheByRegex sets CacheByRegex field to given value.

### HasCacheByRegex

`func (o *OasCachePlugin) HasCacheByRegex() bool`

HasCacheByRegex returns a boolean if a field has been set.

### GetCacheResponseCodes

`func (o *OasCachePlugin) GetCacheResponseCodes() []int32`

GetCacheResponseCodes returns the CacheResponseCodes field if non-nil, zero value otherwise.

### GetCacheResponseCodesOk

`func (o *OasCachePlugin) GetCacheResponseCodesOk() (*[]int32, bool)`

GetCacheResponseCodesOk returns a tuple with the CacheResponseCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheResponseCodes

`func (o *OasCachePlugin) SetCacheResponseCodes(v []int32)`

SetCacheResponseCodes sets CacheResponseCodes field to given value.

### HasCacheResponseCodes

`func (o *OasCachePlugin) HasCacheResponseCodes() bool`

HasCacheResponseCodes returns a boolean if a field has been set.

### GetEnabled

`func (o *OasCachePlugin) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *OasCachePlugin) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *OasCachePlugin) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *OasCachePlugin) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetTimeout

`func (o *OasCachePlugin) GetTimeout() int32`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *OasCachePlugin) GetTimeoutOk() (*int32, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *OasCachePlugin) SetTimeout(v int32)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *OasCachePlugin) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


