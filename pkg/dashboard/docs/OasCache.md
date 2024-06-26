# OasCache

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CacheAllSafeRequests** | Pointer to **bool** |  | [optional] 
**CacheByHeaders** | Pointer to **[]string** |  | [optional] 
**CacheResponseCodes** | Pointer to **[]int32** |  | [optional] 
**ControlTTLHeaderName** | Pointer to **string** |  | [optional] 
**EnableUpstreamCacheControl** | Pointer to **bool** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Timeout** | Pointer to **int32** |  | [optional] 

## Methods

### NewOasCache

`func NewOasCache() *OasCache`

NewOasCache instantiates a new OasCache object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOasCacheWithDefaults

`func NewOasCacheWithDefaults() *OasCache`

NewOasCacheWithDefaults instantiates a new OasCache object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCacheAllSafeRequests

`func (o *OasCache) GetCacheAllSafeRequests() bool`

GetCacheAllSafeRequests returns the CacheAllSafeRequests field if non-nil, zero value otherwise.

### GetCacheAllSafeRequestsOk

`func (o *OasCache) GetCacheAllSafeRequestsOk() (*bool, bool)`

GetCacheAllSafeRequestsOk returns a tuple with the CacheAllSafeRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheAllSafeRequests

`func (o *OasCache) SetCacheAllSafeRequests(v bool)`

SetCacheAllSafeRequests sets CacheAllSafeRequests field to given value.

### HasCacheAllSafeRequests

`func (o *OasCache) HasCacheAllSafeRequests() bool`

HasCacheAllSafeRequests returns a boolean if a field has been set.

### GetCacheByHeaders

`func (o *OasCache) GetCacheByHeaders() []string`

GetCacheByHeaders returns the CacheByHeaders field if non-nil, zero value otherwise.

### GetCacheByHeadersOk

`func (o *OasCache) GetCacheByHeadersOk() (*[]string, bool)`

GetCacheByHeadersOk returns a tuple with the CacheByHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheByHeaders

`func (o *OasCache) SetCacheByHeaders(v []string)`

SetCacheByHeaders sets CacheByHeaders field to given value.

### HasCacheByHeaders

`func (o *OasCache) HasCacheByHeaders() bool`

HasCacheByHeaders returns a boolean if a field has been set.

### GetCacheResponseCodes

`func (o *OasCache) GetCacheResponseCodes() []int32`

GetCacheResponseCodes returns the CacheResponseCodes field if non-nil, zero value otherwise.

### GetCacheResponseCodesOk

`func (o *OasCache) GetCacheResponseCodesOk() (*[]int32, bool)`

GetCacheResponseCodesOk returns a tuple with the CacheResponseCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheResponseCodes

`func (o *OasCache) SetCacheResponseCodes(v []int32)`

SetCacheResponseCodes sets CacheResponseCodes field to given value.

### HasCacheResponseCodes

`func (o *OasCache) HasCacheResponseCodes() bool`

HasCacheResponseCodes returns a boolean if a field has been set.

### GetControlTTLHeaderName

`func (o *OasCache) GetControlTTLHeaderName() string`

GetControlTTLHeaderName returns the ControlTTLHeaderName field if non-nil, zero value otherwise.

### GetControlTTLHeaderNameOk

`func (o *OasCache) GetControlTTLHeaderNameOk() (*string, bool)`

GetControlTTLHeaderNameOk returns a tuple with the ControlTTLHeaderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlTTLHeaderName

`func (o *OasCache) SetControlTTLHeaderName(v string)`

SetControlTTLHeaderName sets ControlTTLHeaderName field to given value.

### HasControlTTLHeaderName

`func (o *OasCache) HasControlTTLHeaderName() bool`

HasControlTTLHeaderName returns a boolean if a field has been set.

### GetEnableUpstreamCacheControl

`func (o *OasCache) GetEnableUpstreamCacheControl() bool`

GetEnableUpstreamCacheControl returns the EnableUpstreamCacheControl field if non-nil, zero value otherwise.

### GetEnableUpstreamCacheControlOk

`func (o *OasCache) GetEnableUpstreamCacheControlOk() (*bool, bool)`

GetEnableUpstreamCacheControlOk returns a tuple with the EnableUpstreamCacheControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableUpstreamCacheControl

`func (o *OasCache) SetEnableUpstreamCacheControl(v bool)`

SetEnableUpstreamCacheControl sets EnableUpstreamCacheControl field to given value.

### HasEnableUpstreamCacheControl

`func (o *OasCache) HasEnableUpstreamCacheControl() bool`

HasEnableUpstreamCacheControl returns a boolean if a field has been set.

### GetEnabled

`func (o *OasCache) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *OasCache) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *OasCache) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *OasCache) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetTimeout

`func (o *OasCache) GetTimeout() int32`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *OasCache) GetTimeoutOk() (*int32, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *OasCache) SetTimeout(v int32)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *OasCache) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


