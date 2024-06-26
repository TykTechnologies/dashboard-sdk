# ApidefOpenIDOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Providers** | Pointer to [**[]ApidefOIDProviderConfig**](ApidefOIDProviderConfig.md) |  | [optional] 
**SegregateByClient** | Pointer to **bool** |  | [optional] 

## Methods

### NewApidefOpenIDOptions

`func NewApidefOpenIDOptions() *ApidefOpenIDOptions`

NewApidefOpenIDOptions instantiates a new ApidefOpenIDOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApidefOpenIDOptionsWithDefaults

`func NewApidefOpenIDOptionsWithDefaults() *ApidefOpenIDOptions`

NewApidefOpenIDOptionsWithDefaults instantiates a new ApidefOpenIDOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProviders

`func (o *ApidefOpenIDOptions) GetProviders() []ApidefOIDProviderConfig`

GetProviders returns the Providers field if non-nil, zero value otherwise.

### GetProvidersOk

`func (o *ApidefOpenIDOptions) GetProvidersOk() (*[]ApidefOIDProviderConfig, bool)`

GetProvidersOk returns a tuple with the Providers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviders

`func (o *ApidefOpenIDOptions) SetProviders(v []ApidefOIDProviderConfig)`

SetProviders sets Providers field to given value.

### HasProviders

`func (o *ApidefOpenIDOptions) HasProviders() bool`

HasProviders returns a boolean if a field has been set.

### SetProvidersNil

`func (o *ApidefOpenIDOptions) SetProvidersNil(b bool)`

 SetProvidersNil sets the value for Providers to be an explicit nil

### UnsetProviders
`func (o *ApidefOpenIDOptions) UnsetProviders()`

UnsetProviders ensures that no value is present for Providers, not even an explicit nil
### GetSegregateByClient

`func (o *ApidefOpenIDOptions) GetSegregateByClient() bool`

GetSegregateByClient returns the SegregateByClient field if non-nil, zero value otherwise.

### GetSegregateByClientOk

`func (o *ApidefOpenIDOptions) GetSegregateByClientOk() (*bool, bool)`

GetSegregateByClientOk returns a tuple with the SegregateByClient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegregateByClient

`func (o *ApidefOpenIDOptions) SetSegregateByClient(v bool)`

SetSegregateByClient sets SegregateByClient field to given value.

### HasSegregateByClient

`func (o *ApidefOpenIDOptions) HasSegregateByClient() bool`

HasSegregateByClient returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


