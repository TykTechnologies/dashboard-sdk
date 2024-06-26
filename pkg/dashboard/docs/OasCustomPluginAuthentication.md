# OasCustomPluginAuthentication

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthSources** | Pointer to [**OasAuthSources**](OasAuthSources.md) |  | [optional] 
**Config** | Pointer to [**OasAuthenticationPlugin**](OasAuthenticationPlugin.md) |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 

## Methods

### NewOasCustomPluginAuthentication

`func NewOasCustomPluginAuthentication() *OasCustomPluginAuthentication`

NewOasCustomPluginAuthentication instantiates a new OasCustomPluginAuthentication object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOasCustomPluginAuthenticationWithDefaults

`func NewOasCustomPluginAuthenticationWithDefaults() *OasCustomPluginAuthentication`

NewOasCustomPluginAuthenticationWithDefaults instantiates a new OasCustomPluginAuthentication object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthSources

`func (o *OasCustomPluginAuthentication) GetAuthSources() OasAuthSources`

GetAuthSources returns the AuthSources field if non-nil, zero value otherwise.

### GetAuthSourcesOk

`func (o *OasCustomPluginAuthentication) GetAuthSourcesOk() (*OasAuthSources, bool)`

GetAuthSourcesOk returns a tuple with the AuthSources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthSources

`func (o *OasCustomPluginAuthentication) SetAuthSources(v OasAuthSources)`

SetAuthSources sets AuthSources field to given value.

### HasAuthSources

`func (o *OasCustomPluginAuthentication) HasAuthSources() bool`

HasAuthSources returns a boolean if a field has been set.

### GetConfig

`func (o *OasCustomPluginAuthentication) GetConfig() OasAuthenticationPlugin`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *OasCustomPluginAuthentication) GetConfigOk() (*OasAuthenticationPlugin, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *OasCustomPluginAuthentication) SetConfig(v OasAuthenticationPlugin)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *OasCustomPluginAuthentication) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetEnabled

`func (o *OasCustomPluginAuthentication) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *OasCustomPluginAuthentication) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *OasCustomPluginAuthentication) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *OasCustomPluginAuthentication) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


