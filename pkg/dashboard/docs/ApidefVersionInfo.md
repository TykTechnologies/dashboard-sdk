# ApidefVersionInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Expires** | Pointer to **string** |  | [optional] 
**ExtendedPaths** | Pointer to [**ApidefExtendedPathsSet**](ApidefExtendedPathsSet.md) |  | [optional] 
**GlobalHeaders** | Pointer to **map[string]string** |  | [optional] 
**GlobalHeadersDisabled** | Pointer to **bool** |  | [optional] 
**GlobalHeadersRemove** | Pointer to **[]string** |  | [optional] 
**GlobalResponseHeaders** | Pointer to **map[string]string** |  | [optional] 
**GlobalResponseHeadersDisabled** | Pointer to **bool** |  | [optional] 
**GlobalResponseHeadersRemove** | Pointer to **[]string** |  | [optional] 
**GlobalSizeLimit** | Pointer to **int32** |  | [optional] 
**IgnoreEndpointCase** | Pointer to **bool** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**OverrideTarget** | Pointer to **string** |  | [optional] 
**Paths** | Pointer to [**ApidefVersionInfoPaths**](ApidefVersionInfoPaths.md) |  | [optional] 
**UseExtendedPaths** | Pointer to **bool** |  | [optional] 

## Methods

### NewApidefVersionInfo

`func NewApidefVersionInfo() *ApidefVersionInfo`

NewApidefVersionInfo instantiates a new ApidefVersionInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApidefVersionInfoWithDefaults

`func NewApidefVersionInfoWithDefaults() *ApidefVersionInfo`

NewApidefVersionInfoWithDefaults instantiates a new ApidefVersionInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpires

`func (o *ApidefVersionInfo) GetExpires() string`

GetExpires returns the Expires field if non-nil, zero value otherwise.

### GetExpiresOk

`func (o *ApidefVersionInfo) GetExpiresOk() (*string, bool)`

GetExpiresOk returns a tuple with the Expires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpires

`func (o *ApidefVersionInfo) SetExpires(v string)`

SetExpires sets Expires field to given value.

### HasExpires

`func (o *ApidefVersionInfo) HasExpires() bool`

HasExpires returns a boolean if a field has been set.

### GetExtendedPaths

`func (o *ApidefVersionInfo) GetExtendedPaths() ApidefExtendedPathsSet`

GetExtendedPaths returns the ExtendedPaths field if non-nil, zero value otherwise.

### GetExtendedPathsOk

`func (o *ApidefVersionInfo) GetExtendedPathsOk() (*ApidefExtendedPathsSet, bool)`

GetExtendedPathsOk returns a tuple with the ExtendedPaths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendedPaths

`func (o *ApidefVersionInfo) SetExtendedPaths(v ApidefExtendedPathsSet)`

SetExtendedPaths sets ExtendedPaths field to given value.

### HasExtendedPaths

`func (o *ApidefVersionInfo) HasExtendedPaths() bool`

HasExtendedPaths returns a boolean if a field has been set.

### GetGlobalHeaders

`func (o *ApidefVersionInfo) GetGlobalHeaders() map[string]string`

GetGlobalHeaders returns the GlobalHeaders field if non-nil, zero value otherwise.

### GetGlobalHeadersOk

`func (o *ApidefVersionInfo) GetGlobalHeadersOk() (*map[string]string, bool)`

GetGlobalHeadersOk returns a tuple with the GlobalHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalHeaders

`func (o *ApidefVersionInfo) SetGlobalHeaders(v map[string]string)`

SetGlobalHeaders sets GlobalHeaders field to given value.

### HasGlobalHeaders

`func (o *ApidefVersionInfo) HasGlobalHeaders() bool`

HasGlobalHeaders returns a boolean if a field has been set.

### SetGlobalHeadersNil

`func (o *ApidefVersionInfo) SetGlobalHeadersNil(b bool)`

 SetGlobalHeadersNil sets the value for GlobalHeaders to be an explicit nil

### UnsetGlobalHeaders
`func (o *ApidefVersionInfo) UnsetGlobalHeaders()`

UnsetGlobalHeaders ensures that no value is present for GlobalHeaders, not even an explicit nil
### GetGlobalHeadersDisabled

`func (o *ApidefVersionInfo) GetGlobalHeadersDisabled() bool`

GetGlobalHeadersDisabled returns the GlobalHeadersDisabled field if non-nil, zero value otherwise.

### GetGlobalHeadersDisabledOk

`func (o *ApidefVersionInfo) GetGlobalHeadersDisabledOk() (*bool, bool)`

GetGlobalHeadersDisabledOk returns a tuple with the GlobalHeadersDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalHeadersDisabled

`func (o *ApidefVersionInfo) SetGlobalHeadersDisabled(v bool)`

SetGlobalHeadersDisabled sets GlobalHeadersDisabled field to given value.

### HasGlobalHeadersDisabled

`func (o *ApidefVersionInfo) HasGlobalHeadersDisabled() bool`

HasGlobalHeadersDisabled returns a boolean if a field has been set.

### GetGlobalHeadersRemove

`func (o *ApidefVersionInfo) GetGlobalHeadersRemove() []string`

GetGlobalHeadersRemove returns the GlobalHeadersRemove field if non-nil, zero value otherwise.

### GetGlobalHeadersRemoveOk

`func (o *ApidefVersionInfo) GetGlobalHeadersRemoveOk() (*[]string, bool)`

GetGlobalHeadersRemoveOk returns a tuple with the GlobalHeadersRemove field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalHeadersRemove

`func (o *ApidefVersionInfo) SetGlobalHeadersRemove(v []string)`

SetGlobalHeadersRemove sets GlobalHeadersRemove field to given value.

### HasGlobalHeadersRemove

`func (o *ApidefVersionInfo) HasGlobalHeadersRemove() bool`

HasGlobalHeadersRemove returns a boolean if a field has been set.

### SetGlobalHeadersRemoveNil

`func (o *ApidefVersionInfo) SetGlobalHeadersRemoveNil(b bool)`

 SetGlobalHeadersRemoveNil sets the value for GlobalHeadersRemove to be an explicit nil

### UnsetGlobalHeadersRemove
`func (o *ApidefVersionInfo) UnsetGlobalHeadersRemove()`

UnsetGlobalHeadersRemove ensures that no value is present for GlobalHeadersRemove, not even an explicit nil
### GetGlobalResponseHeaders

`func (o *ApidefVersionInfo) GetGlobalResponseHeaders() map[string]string`

GetGlobalResponseHeaders returns the GlobalResponseHeaders field if non-nil, zero value otherwise.

### GetGlobalResponseHeadersOk

`func (o *ApidefVersionInfo) GetGlobalResponseHeadersOk() (*map[string]string, bool)`

GetGlobalResponseHeadersOk returns a tuple with the GlobalResponseHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalResponseHeaders

`func (o *ApidefVersionInfo) SetGlobalResponseHeaders(v map[string]string)`

SetGlobalResponseHeaders sets GlobalResponseHeaders field to given value.

### HasGlobalResponseHeaders

`func (o *ApidefVersionInfo) HasGlobalResponseHeaders() bool`

HasGlobalResponseHeaders returns a boolean if a field has been set.

### SetGlobalResponseHeadersNil

`func (o *ApidefVersionInfo) SetGlobalResponseHeadersNil(b bool)`

 SetGlobalResponseHeadersNil sets the value for GlobalResponseHeaders to be an explicit nil

### UnsetGlobalResponseHeaders
`func (o *ApidefVersionInfo) UnsetGlobalResponseHeaders()`

UnsetGlobalResponseHeaders ensures that no value is present for GlobalResponseHeaders, not even an explicit nil
### GetGlobalResponseHeadersDisabled

`func (o *ApidefVersionInfo) GetGlobalResponseHeadersDisabled() bool`

GetGlobalResponseHeadersDisabled returns the GlobalResponseHeadersDisabled field if non-nil, zero value otherwise.

### GetGlobalResponseHeadersDisabledOk

`func (o *ApidefVersionInfo) GetGlobalResponseHeadersDisabledOk() (*bool, bool)`

GetGlobalResponseHeadersDisabledOk returns a tuple with the GlobalResponseHeadersDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalResponseHeadersDisabled

`func (o *ApidefVersionInfo) SetGlobalResponseHeadersDisabled(v bool)`

SetGlobalResponseHeadersDisabled sets GlobalResponseHeadersDisabled field to given value.

### HasGlobalResponseHeadersDisabled

`func (o *ApidefVersionInfo) HasGlobalResponseHeadersDisabled() bool`

HasGlobalResponseHeadersDisabled returns a boolean if a field has been set.

### GetGlobalResponseHeadersRemove

`func (o *ApidefVersionInfo) GetGlobalResponseHeadersRemove() []string`

GetGlobalResponseHeadersRemove returns the GlobalResponseHeadersRemove field if non-nil, zero value otherwise.

### GetGlobalResponseHeadersRemoveOk

`func (o *ApidefVersionInfo) GetGlobalResponseHeadersRemoveOk() (*[]string, bool)`

GetGlobalResponseHeadersRemoveOk returns a tuple with the GlobalResponseHeadersRemove field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalResponseHeadersRemove

`func (o *ApidefVersionInfo) SetGlobalResponseHeadersRemove(v []string)`

SetGlobalResponseHeadersRemove sets GlobalResponseHeadersRemove field to given value.

### HasGlobalResponseHeadersRemove

`func (o *ApidefVersionInfo) HasGlobalResponseHeadersRemove() bool`

HasGlobalResponseHeadersRemove returns a boolean if a field has been set.

### SetGlobalResponseHeadersRemoveNil

`func (o *ApidefVersionInfo) SetGlobalResponseHeadersRemoveNil(b bool)`

 SetGlobalResponseHeadersRemoveNil sets the value for GlobalResponseHeadersRemove to be an explicit nil

### UnsetGlobalResponseHeadersRemove
`func (o *ApidefVersionInfo) UnsetGlobalResponseHeadersRemove()`

UnsetGlobalResponseHeadersRemove ensures that no value is present for GlobalResponseHeadersRemove, not even an explicit nil
### GetGlobalSizeLimit

`func (o *ApidefVersionInfo) GetGlobalSizeLimit() int32`

GetGlobalSizeLimit returns the GlobalSizeLimit field if non-nil, zero value otherwise.

### GetGlobalSizeLimitOk

`func (o *ApidefVersionInfo) GetGlobalSizeLimitOk() (*int32, bool)`

GetGlobalSizeLimitOk returns a tuple with the GlobalSizeLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalSizeLimit

`func (o *ApidefVersionInfo) SetGlobalSizeLimit(v int32)`

SetGlobalSizeLimit sets GlobalSizeLimit field to given value.

### HasGlobalSizeLimit

`func (o *ApidefVersionInfo) HasGlobalSizeLimit() bool`

HasGlobalSizeLimit returns a boolean if a field has been set.

### GetIgnoreEndpointCase

`func (o *ApidefVersionInfo) GetIgnoreEndpointCase() bool`

GetIgnoreEndpointCase returns the IgnoreEndpointCase field if non-nil, zero value otherwise.

### GetIgnoreEndpointCaseOk

`func (o *ApidefVersionInfo) GetIgnoreEndpointCaseOk() (*bool, bool)`

GetIgnoreEndpointCaseOk returns a tuple with the IgnoreEndpointCase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreEndpointCase

`func (o *ApidefVersionInfo) SetIgnoreEndpointCase(v bool)`

SetIgnoreEndpointCase sets IgnoreEndpointCase field to given value.

### HasIgnoreEndpointCase

`func (o *ApidefVersionInfo) HasIgnoreEndpointCase() bool`

HasIgnoreEndpointCase returns a boolean if a field has been set.

### GetName

`func (o *ApidefVersionInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApidefVersionInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApidefVersionInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ApidefVersionInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOverrideTarget

`func (o *ApidefVersionInfo) GetOverrideTarget() string`

GetOverrideTarget returns the OverrideTarget field if non-nil, zero value otherwise.

### GetOverrideTargetOk

`func (o *ApidefVersionInfo) GetOverrideTargetOk() (*string, bool)`

GetOverrideTargetOk returns a tuple with the OverrideTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideTarget

`func (o *ApidefVersionInfo) SetOverrideTarget(v string)`

SetOverrideTarget sets OverrideTarget field to given value.

### HasOverrideTarget

`func (o *ApidefVersionInfo) HasOverrideTarget() bool`

HasOverrideTarget returns a boolean if a field has been set.

### GetPaths

`func (o *ApidefVersionInfo) GetPaths() ApidefVersionInfoPaths`

GetPaths returns the Paths field if non-nil, zero value otherwise.

### GetPathsOk

`func (o *ApidefVersionInfo) GetPathsOk() (*ApidefVersionInfoPaths, bool)`

GetPathsOk returns a tuple with the Paths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaths

`func (o *ApidefVersionInfo) SetPaths(v ApidefVersionInfoPaths)`

SetPaths sets Paths field to given value.

### HasPaths

`func (o *ApidefVersionInfo) HasPaths() bool`

HasPaths returns a boolean if a field has been set.

### GetUseExtendedPaths

`func (o *ApidefVersionInfo) GetUseExtendedPaths() bool`

GetUseExtendedPaths returns the UseExtendedPaths field if non-nil, zero value otherwise.

### GetUseExtendedPathsOk

`func (o *ApidefVersionInfo) GetUseExtendedPathsOk() (*bool, bool)`

GetUseExtendedPathsOk returns a tuple with the UseExtendedPaths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseExtendedPaths

`func (o *ApidefVersionInfo) SetUseExtendedPaths(v bool)`

SetUseExtendedPaths sets UseExtendedPaths field to given value.

### HasUseExtendedPaths

`func (o *ApidefVersionInfo) HasUseExtendedPaths() bool`

HasUseExtendedPaths returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


