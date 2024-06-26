# OasCORS

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowCredentials** | Pointer to **bool** |  | [optional] 
**AllowedHeaders** | Pointer to **[]string** |  | [optional] 
**AllowedMethods** | Pointer to **[]string** |  | [optional] 
**AllowedOrigins** | Pointer to **[]string** |  | [optional] 
**Debug** | Pointer to **bool** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**ExposedHeaders** | Pointer to **[]string** |  | [optional] 
**MaxAge** | Pointer to **int32** |  | [optional] 
**OptionsPassthrough** | Pointer to **bool** |  | [optional] 

## Methods

### NewOasCORS

`func NewOasCORS() *OasCORS`

NewOasCORS instantiates a new OasCORS object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOasCORSWithDefaults

`func NewOasCORSWithDefaults() *OasCORS`

NewOasCORSWithDefaults instantiates a new OasCORS object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowCredentials

`func (o *OasCORS) GetAllowCredentials() bool`

GetAllowCredentials returns the AllowCredentials field if non-nil, zero value otherwise.

### GetAllowCredentialsOk

`func (o *OasCORS) GetAllowCredentialsOk() (*bool, bool)`

GetAllowCredentialsOk returns a tuple with the AllowCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowCredentials

`func (o *OasCORS) SetAllowCredentials(v bool)`

SetAllowCredentials sets AllowCredentials field to given value.

### HasAllowCredentials

`func (o *OasCORS) HasAllowCredentials() bool`

HasAllowCredentials returns a boolean if a field has been set.

### GetAllowedHeaders

`func (o *OasCORS) GetAllowedHeaders() []string`

GetAllowedHeaders returns the AllowedHeaders field if non-nil, zero value otherwise.

### GetAllowedHeadersOk

`func (o *OasCORS) GetAllowedHeadersOk() (*[]string, bool)`

GetAllowedHeadersOk returns a tuple with the AllowedHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedHeaders

`func (o *OasCORS) SetAllowedHeaders(v []string)`

SetAllowedHeaders sets AllowedHeaders field to given value.

### HasAllowedHeaders

`func (o *OasCORS) HasAllowedHeaders() bool`

HasAllowedHeaders returns a boolean if a field has been set.

### GetAllowedMethods

`func (o *OasCORS) GetAllowedMethods() []string`

GetAllowedMethods returns the AllowedMethods field if non-nil, zero value otherwise.

### GetAllowedMethodsOk

`func (o *OasCORS) GetAllowedMethodsOk() (*[]string, bool)`

GetAllowedMethodsOk returns a tuple with the AllowedMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedMethods

`func (o *OasCORS) SetAllowedMethods(v []string)`

SetAllowedMethods sets AllowedMethods field to given value.

### HasAllowedMethods

`func (o *OasCORS) HasAllowedMethods() bool`

HasAllowedMethods returns a boolean if a field has been set.

### GetAllowedOrigins

`func (o *OasCORS) GetAllowedOrigins() []string`

GetAllowedOrigins returns the AllowedOrigins field if non-nil, zero value otherwise.

### GetAllowedOriginsOk

`func (o *OasCORS) GetAllowedOriginsOk() (*[]string, bool)`

GetAllowedOriginsOk returns a tuple with the AllowedOrigins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedOrigins

`func (o *OasCORS) SetAllowedOrigins(v []string)`

SetAllowedOrigins sets AllowedOrigins field to given value.

### HasAllowedOrigins

`func (o *OasCORS) HasAllowedOrigins() bool`

HasAllowedOrigins returns a boolean if a field has been set.

### GetDebug

`func (o *OasCORS) GetDebug() bool`

GetDebug returns the Debug field if non-nil, zero value otherwise.

### GetDebugOk

`func (o *OasCORS) GetDebugOk() (*bool, bool)`

GetDebugOk returns a tuple with the Debug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebug

`func (o *OasCORS) SetDebug(v bool)`

SetDebug sets Debug field to given value.

### HasDebug

`func (o *OasCORS) HasDebug() bool`

HasDebug returns a boolean if a field has been set.

### GetEnabled

`func (o *OasCORS) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *OasCORS) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *OasCORS) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *OasCORS) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetExposedHeaders

`func (o *OasCORS) GetExposedHeaders() []string`

GetExposedHeaders returns the ExposedHeaders field if non-nil, zero value otherwise.

### GetExposedHeadersOk

`func (o *OasCORS) GetExposedHeadersOk() (*[]string, bool)`

GetExposedHeadersOk returns a tuple with the ExposedHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExposedHeaders

`func (o *OasCORS) SetExposedHeaders(v []string)`

SetExposedHeaders sets ExposedHeaders field to given value.

### HasExposedHeaders

`func (o *OasCORS) HasExposedHeaders() bool`

HasExposedHeaders returns a boolean if a field has been set.

### GetMaxAge

`func (o *OasCORS) GetMaxAge() int32`

GetMaxAge returns the MaxAge field if non-nil, zero value otherwise.

### GetMaxAgeOk

`func (o *OasCORS) GetMaxAgeOk() (*int32, bool)`

GetMaxAgeOk returns a tuple with the MaxAge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAge

`func (o *OasCORS) SetMaxAge(v int32)`

SetMaxAge sets MaxAge field to given value.

### HasMaxAge

`func (o *OasCORS) HasMaxAge() bool`

HasMaxAge returns a boolean if a field has been set.

### GetOptionsPassthrough

`func (o *OasCORS) GetOptionsPassthrough() bool`

GetOptionsPassthrough returns the OptionsPassthrough field if non-nil, zero value otherwise.

### GetOptionsPassthroughOk

`func (o *OasCORS) GetOptionsPassthroughOk() (*bool, bool)`

GetOptionsPassthroughOk returns a tuple with the OptionsPassthrough field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionsPassthrough

`func (o *OasCORS) SetOptionsPassthrough(v bool)`

SetOptionsPassthrough sets OptionsPassthrough field to given value.

### HasOptionsPassthrough

`func (o *OasCORS) HasOptionsPassthrough() bool`

HasOptionsPassthrough returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


