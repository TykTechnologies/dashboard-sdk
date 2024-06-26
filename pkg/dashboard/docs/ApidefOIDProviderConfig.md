# ApidefOIDProviderConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientIds** | Pointer to **map[string]string** |  | [optional] 
**Issuer** | Pointer to **string** |  | [optional] 

## Methods

### NewApidefOIDProviderConfig

`func NewApidefOIDProviderConfig() *ApidefOIDProviderConfig`

NewApidefOIDProviderConfig instantiates a new ApidefOIDProviderConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApidefOIDProviderConfigWithDefaults

`func NewApidefOIDProviderConfigWithDefaults() *ApidefOIDProviderConfig`

NewApidefOIDProviderConfigWithDefaults instantiates a new ApidefOIDProviderConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientIds

`func (o *ApidefOIDProviderConfig) GetClientIds() map[string]string`

GetClientIds returns the ClientIds field if non-nil, zero value otherwise.

### GetClientIdsOk

`func (o *ApidefOIDProviderConfig) GetClientIdsOk() (*map[string]string, bool)`

GetClientIdsOk returns a tuple with the ClientIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientIds

`func (o *ApidefOIDProviderConfig) SetClientIds(v map[string]string)`

SetClientIds sets ClientIds field to given value.

### HasClientIds

`func (o *ApidefOIDProviderConfig) HasClientIds() bool`

HasClientIds returns a boolean if a field has been set.

### SetClientIdsNil

`func (o *ApidefOIDProviderConfig) SetClientIdsNil(b bool)`

 SetClientIdsNil sets the value for ClientIds to be an explicit nil

### UnsetClientIds
`func (o *ApidefOIDProviderConfig) UnsetClientIds()`

UnsetClientIds ensures that no value is present for ClientIds, not even an explicit nil
### GetIssuer

`func (o *ApidefOIDProviderConfig) GetIssuer() string`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *ApidefOIDProviderConfig) GetIssuerOk() (*string, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *ApidefOIDProviderConfig) SetIssuer(v string)`

SetIssuer sets Issuer field to given value.

### HasIssuer

`func (o *ApidefOIDProviderConfig) HasIssuer() bool`

HasIssuer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


