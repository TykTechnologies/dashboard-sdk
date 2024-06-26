# OasOIDC

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthSources** | Pointer to [**OasAuthSources**](OasAuthSources.md) |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Providers** | Pointer to [**[]OasProvider**](OasProvider.md) |  | [optional] 
**Scopes** | Pointer to [**OasScopes**](OasScopes.md) |  | [optional] 
**SegregateByClientId** | Pointer to **bool** |  | [optional] 

## Methods

### NewOasOIDC

`func NewOasOIDC() *OasOIDC`

NewOasOIDC instantiates a new OasOIDC object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOasOIDCWithDefaults

`func NewOasOIDCWithDefaults() *OasOIDC`

NewOasOIDCWithDefaults instantiates a new OasOIDC object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthSources

`func (o *OasOIDC) GetAuthSources() OasAuthSources`

GetAuthSources returns the AuthSources field if non-nil, zero value otherwise.

### GetAuthSourcesOk

`func (o *OasOIDC) GetAuthSourcesOk() (*OasAuthSources, bool)`

GetAuthSourcesOk returns a tuple with the AuthSources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthSources

`func (o *OasOIDC) SetAuthSources(v OasAuthSources)`

SetAuthSources sets AuthSources field to given value.

### HasAuthSources

`func (o *OasOIDC) HasAuthSources() bool`

HasAuthSources returns a boolean if a field has been set.

### GetEnabled

`func (o *OasOIDC) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *OasOIDC) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *OasOIDC) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *OasOIDC) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetProviders

`func (o *OasOIDC) GetProviders() []OasProvider`

GetProviders returns the Providers field if non-nil, zero value otherwise.

### GetProvidersOk

`func (o *OasOIDC) GetProvidersOk() (*[]OasProvider, bool)`

GetProvidersOk returns a tuple with the Providers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviders

`func (o *OasOIDC) SetProviders(v []OasProvider)`

SetProviders sets Providers field to given value.

### HasProviders

`func (o *OasOIDC) HasProviders() bool`

HasProviders returns a boolean if a field has been set.

### GetScopes

`func (o *OasOIDC) GetScopes() OasScopes`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *OasOIDC) GetScopesOk() (*OasScopes, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *OasOIDC) SetScopes(v OasScopes)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *OasOIDC) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### GetSegregateByClientId

`func (o *OasOIDC) GetSegregateByClientId() bool`

GetSegregateByClientId returns the SegregateByClientId field if non-nil, zero value otherwise.

### GetSegregateByClientIdOk

`func (o *OasOIDC) GetSegregateByClientIdOk() (*bool, bool)`

GetSegregateByClientIdOk returns a tuple with the SegregateByClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegregateByClientId

`func (o *OasOIDC) SetSegregateByClientId(v bool)`

SetSegregateByClientId sets SegregateByClientId field to given value.

### HasSegregateByClientId

`func (o *OasOIDC) HasSegregateByClientId() bool`

HasSegregateByClientId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


