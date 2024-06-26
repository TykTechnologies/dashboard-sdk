# ApidefExternalOAuth

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** |  | [optional] 
**Providers** | Pointer to [**[]ApidefProvider**](ApidefProvider.md) |  | [optional] 

## Methods

### NewApidefExternalOAuth

`func NewApidefExternalOAuth() *ApidefExternalOAuth`

NewApidefExternalOAuth instantiates a new ApidefExternalOAuth object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApidefExternalOAuthWithDefaults

`func NewApidefExternalOAuthWithDefaults() *ApidefExternalOAuth`

NewApidefExternalOAuthWithDefaults instantiates a new ApidefExternalOAuth object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *ApidefExternalOAuth) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ApidefExternalOAuth) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ApidefExternalOAuth) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ApidefExternalOAuth) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetProviders

`func (o *ApidefExternalOAuth) GetProviders() []ApidefProvider`

GetProviders returns the Providers field if non-nil, zero value otherwise.

### GetProvidersOk

`func (o *ApidefExternalOAuth) GetProvidersOk() (*[]ApidefProvider, bool)`

GetProvidersOk returns a tuple with the Providers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviders

`func (o *ApidefExternalOAuth) SetProviders(v []ApidefProvider)`

SetProviders sets Providers field to given value.

### HasProviders

`func (o *ApidefExternalOAuth) HasProviders() bool`

HasProviders returns a boolean if a field has been set.

### SetProvidersNil

`func (o *ApidefExternalOAuth) SetProvidersNil(b bool)`

 SetProvidersNil sets the value for Providers to be an explicit nil

### UnsetProviders
`func (o *ApidefExternalOAuth) UnsetProviders()`

UnsetProviders ensures that no value is present for Providers, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


