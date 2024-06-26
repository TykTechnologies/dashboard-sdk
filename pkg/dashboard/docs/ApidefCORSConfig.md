# ApidefCORSConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowCredentials** | Pointer to **bool** |  | [optional] 
**AllowedHeaders** | Pointer to **[]string** |  | [optional] 
**AllowedMethods** | Pointer to **[]string** |  | [optional] 
**AllowedOrigins** | Pointer to **[]string** |  | [optional] 
**Debug** | Pointer to **bool** |  | [optional] 
**Enable** | Pointer to **bool** |  | [optional] 
**ExposedHeaders** | Pointer to **[]string** |  | [optional] 
**MaxAge** | Pointer to **int32** |  | [optional] 
**OptionsPassthrough** | Pointer to **bool** |  | [optional] 

## Methods

### NewApidefCORSConfig

`func NewApidefCORSConfig() *ApidefCORSConfig`

NewApidefCORSConfig instantiates a new ApidefCORSConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApidefCORSConfigWithDefaults

`func NewApidefCORSConfigWithDefaults() *ApidefCORSConfig`

NewApidefCORSConfigWithDefaults instantiates a new ApidefCORSConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowCredentials

`func (o *ApidefCORSConfig) GetAllowCredentials() bool`

GetAllowCredentials returns the AllowCredentials field if non-nil, zero value otherwise.

### GetAllowCredentialsOk

`func (o *ApidefCORSConfig) GetAllowCredentialsOk() (*bool, bool)`

GetAllowCredentialsOk returns a tuple with the AllowCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowCredentials

`func (o *ApidefCORSConfig) SetAllowCredentials(v bool)`

SetAllowCredentials sets AllowCredentials field to given value.

### HasAllowCredentials

`func (o *ApidefCORSConfig) HasAllowCredentials() bool`

HasAllowCredentials returns a boolean if a field has been set.

### GetAllowedHeaders

`func (o *ApidefCORSConfig) GetAllowedHeaders() []string`

GetAllowedHeaders returns the AllowedHeaders field if non-nil, zero value otherwise.

### GetAllowedHeadersOk

`func (o *ApidefCORSConfig) GetAllowedHeadersOk() (*[]string, bool)`

GetAllowedHeadersOk returns a tuple with the AllowedHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedHeaders

`func (o *ApidefCORSConfig) SetAllowedHeaders(v []string)`

SetAllowedHeaders sets AllowedHeaders field to given value.

### HasAllowedHeaders

`func (o *ApidefCORSConfig) HasAllowedHeaders() bool`

HasAllowedHeaders returns a boolean if a field has been set.

### SetAllowedHeadersNil

`func (o *ApidefCORSConfig) SetAllowedHeadersNil(b bool)`

 SetAllowedHeadersNil sets the value for AllowedHeaders to be an explicit nil

### UnsetAllowedHeaders
`func (o *ApidefCORSConfig) UnsetAllowedHeaders()`

UnsetAllowedHeaders ensures that no value is present for AllowedHeaders, not even an explicit nil
### GetAllowedMethods

`func (o *ApidefCORSConfig) GetAllowedMethods() []string`

GetAllowedMethods returns the AllowedMethods field if non-nil, zero value otherwise.

### GetAllowedMethodsOk

`func (o *ApidefCORSConfig) GetAllowedMethodsOk() (*[]string, bool)`

GetAllowedMethodsOk returns a tuple with the AllowedMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedMethods

`func (o *ApidefCORSConfig) SetAllowedMethods(v []string)`

SetAllowedMethods sets AllowedMethods field to given value.

### HasAllowedMethods

`func (o *ApidefCORSConfig) HasAllowedMethods() bool`

HasAllowedMethods returns a boolean if a field has been set.

### SetAllowedMethodsNil

`func (o *ApidefCORSConfig) SetAllowedMethodsNil(b bool)`

 SetAllowedMethodsNil sets the value for AllowedMethods to be an explicit nil

### UnsetAllowedMethods
`func (o *ApidefCORSConfig) UnsetAllowedMethods()`

UnsetAllowedMethods ensures that no value is present for AllowedMethods, not even an explicit nil
### GetAllowedOrigins

`func (o *ApidefCORSConfig) GetAllowedOrigins() []string`

GetAllowedOrigins returns the AllowedOrigins field if non-nil, zero value otherwise.

### GetAllowedOriginsOk

`func (o *ApidefCORSConfig) GetAllowedOriginsOk() (*[]string, bool)`

GetAllowedOriginsOk returns a tuple with the AllowedOrigins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedOrigins

`func (o *ApidefCORSConfig) SetAllowedOrigins(v []string)`

SetAllowedOrigins sets AllowedOrigins field to given value.

### HasAllowedOrigins

`func (o *ApidefCORSConfig) HasAllowedOrigins() bool`

HasAllowedOrigins returns a boolean if a field has been set.

### SetAllowedOriginsNil

`func (o *ApidefCORSConfig) SetAllowedOriginsNil(b bool)`

 SetAllowedOriginsNil sets the value for AllowedOrigins to be an explicit nil

### UnsetAllowedOrigins
`func (o *ApidefCORSConfig) UnsetAllowedOrigins()`

UnsetAllowedOrigins ensures that no value is present for AllowedOrigins, not even an explicit nil
### GetDebug

`func (o *ApidefCORSConfig) GetDebug() bool`

GetDebug returns the Debug field if non-nil, zero value otherwise.

### GetDebugOk

`func (o *ApidefCORSConfig) GetDebugOk() (*bool, bool)`

GetDebugOk returns a tuple with the Debug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebug

`func (o *ApidefCORSConfig) SetDebug(v bool)`

SetDebug sets Debug field to given value.

### HasDebug

`func (o *ApidefCORSConfig) HasDebug() bool`

HasDebug returns a boolean if a field has been set.

### GetEnable

`func (o *ApidefCORSConfig) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *ApidefCORSConfig) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *ApidefCORSConfig) SetEnable(v bool)`

SetEnable sets Enable field to given value.

### HasEnable

`func (o *ApidefCORSConfig) HasEnable() bool`

HasEnable returns a boolean if a field has been set.

### GetExposedHeaders

`func (o *ApidefCORSConfig) GetExposedHeaders() []string`

GetExposedHeaders returns the ExposedHeaders field if non-nil, zero value otherwise.

### GetExposedHeadersOk

`func (o *ApidefCORSConfig) GetExposedHeadersOk() (*[]string, bool)`

GetExposedHeadersOk returns a tuple with the ExposedHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExposedHeaders

`func (o *ApidefCORSConfig) SetExposedHeaders(v []string)`

SetExposedHeaders sets ExposedHeaders field to given value.

### HasExposedHeaders

`func (o *ApidefCORSConfig) HasExposedHeaders() bool`

HasExposedHeaders returns a boolean if a field has been set.

### SetExposedHeadersNil

`func (o *ApidefCORSConfig) SetExposedHeadersNil(b bool)`

 SetExposedHeadersNil sets the value for ExposedHeaders to be an explicit nil

### UnsetExposedHeaders
`func (o *ApidefCORSConfig) UnsetExposedHeaders()`

UnsetExposedHeaders ensures that no value is present for ExposedHeaders, not even an explicit nil
### GetMaxAge

`func (o *ApidefCORSConfig) GetMaxAge() int32`

GetMaxAge returns the MaxAge field if non-nil, zero value otherwise.

### GetMaxAgeOk

`func (o *ApidefCORSConfig) GetMaxAgeOk() (*int32, bool)`

GetMaxAgeOk returns a tuple with the MaxAge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAge

`func (o *ApidefCORSConfig) SetMaxAge(v int32)`

SetMaxAge sets MaxAge field to given value.

### HasMaxAge

`func (o *ApidefCORSConfig) HasMaxAge() bool`

HasMaxAge returns a boolean if a field has been set.

### GetOptionsPassthrough

`func (o *ApidefCORSConfig) GetOptionsPassthrough() bool`

GetOptionsPassthrough returns the OptionsPassthrough field if non-nil, zero value otherwise.

### GetOptionsPassthroughOk

`func (o *ApidefCORSConfig) GetOptionsPassthroughOk() (*bool, bool)`

GetOptionsPassthroughOk returns a tuple with the OptionsPassthrough field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionsPassthrough

`func (o *ApidefCORSConfig) SetOptionsPassthrough(v bool)`

SetOptionsPassthrough sets OptionsPassthrough field to given value.

### HasOptionsPassthrough

`func (o *ApidefCORSConfig) HasOptionsPassthrough() bool`

HasOptionsPassthrough returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


