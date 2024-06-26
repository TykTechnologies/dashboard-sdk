# OasAuthentication

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BaseIdentityProvider** | Pointer to **string** |  | [optional] 
**Custom** | Pointer to [**OasCustomPluginAuthentication**](OasCustomPluginAuthentication.md) |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Hmac** | Pointer to [**OasHMAC**](OasHMAC.md) |  | [optional] 
**Oidc** | Pointer to [**OasOIDC**](OasOIDC.md) |  | [optional] 
**SecuritySchemes** | Pointer to **map[string]interface{}** |  | [optional] 
**StripAuthorizationData** | Pointer to **bool** |  | [optional] 

## Methods

### NewOasAuthentication

`func NewOasAuthentication() *OasAuthentication`

NewOasAuthentication instantiates a new OasAuthentication object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOasAuthenticationWithDefaults

`func NewOasAuthenticationWithDefaults() *OasAuthentication`

NewOasAuthenticationWithDefaults instantiates a new OasAuthentication object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBaseIdentityProvider

`func (o *OasAuthentication) GetBaseIdentityProvider() string`

GetBaseIdentityProvider returns the BaseIdentityProvider field if non-nil, zero value otherwise.

### GetBaseIdentityProviderOk

`func (o *OasAuthentication) GetBaseIdentityProviderOk() (*string, bool)`

GetBaseIdentityProviderOk returns a tuple with the BaseIdentityProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseIdentityProvider

`func (o *OasAuthentication) SetBaseIdentityProvider(v string)`

SetBaseIdentityProvider sets BaseIdentityProvider field to given value.

### HasBaseIdentityProvider

`func (o *OasAuthentication) HasBaseIdentityProvider() bool`

HasBaseIdentityProvider returns a boolean if a field has been set.

### GetCustom

`func (o *OasAuthentication) GetCustom() OasCustomPluginAuthentication`

GetCustom returns the Custom field if non-nil, zero value otherwise.

### GetCustomOk

`func (o *OasAuthentication) GetCustomOk() (*OasCustomPluginAuthentication, bool)`

GetCustomOk returns a tuple with the Custom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustom

`func (o *OasAuthentication) SetCustom(v OasCustomPluginAuthentication)`

SetCustom sets Custom field to given value.

### HasCustom

`func (o *OasAuthentication) HasCustom() bool`

HasCustom returns a boolean if a field has been set.

### GetEnabled

`func (o *OasAuthentication) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *OasAuthentication) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *OasAuthentication) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *OasAuthentication) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetHmac

`func (o *OasAuthentication) GetHmac() OasHMAC`

GetHmac returns the Hmac field if non-nil, zero value otherwise.

### GetHmacOk

`func (o *OasAuthentication) GetHmacOk() (*OasHMAC, bool)`

GetHmacOk returns a tuple with the Hmac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHmac

`func (o *OasAuthentication) SetHmac(v OasHMAC)`

SetHmac sets Hmac field to given value.

### HasHmac

`func (o *OasAuthentication) HasHmac() bool`

HasHmac returns a boolean if a field has been set.

### GetOidc

`func (o *OasAuthentication) GetOidc() OasOIDC`

GetOidc returns the Oidc field if non-nil, zero value otherwise.

### GetOidcOk

`func (o *OasAuthentication) GetOidcOk() (*OasOIDC, bool)`

GetOidcOk returns a tuple with the Oidc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOidc

`func (o *OasAuthentication) SetOidc(v OasOIDC)`

SetOidc sets Oidc field to given value.

### HasOidc

`func (o *OasAuthentication) HasOidc() bool`

HasOidc returns a boolean if a field has been set.

### GetSecuritySchemes

`func (o *OasAuthentication) GetSecuritySchemes() map[string]interface{}`

GetSecuritySchemes returns the SecuritySchemes field if non-nil, zero value otherwise.

### GetSecuritySchemesOk

`func (o *OasAuthentication) GetSecuritySchemesOk() (*map[string]interface{}, bool)`

GetSecuritySchemesOk returns a tuple with the SecuritySchemes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecuritySchemes

`func (o *OasAuthentication) SetSecuritySchemes(v map[string]interface{})`

SetSecuritySchemes sets SecuritySchemes field to given value.

### HasSecuritySchemes

`func (o *OasAuthentication) HasSecuritySchemes() bool`

HasSecuritySchemes returns a boolean if a field has been set.

### GetStripAuthorizationData

`func (o *OasAuthentication) GetStripAuthorizationData() bool`

GetStripAuthorizationData returns the StripAuthorizationData field if non-nil, zero value otherwise.

### GetStripAuthorizationDataOk

`func (o *OasAuthentication) GetStripAuthorizationDataOk() (*bool, bool)`

GetStripAuthorizationDataOk returns a tuple with the StripAuthorizationData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStripAuthorizationData

`func (o *OasAuthentication) SetStripAuthorizationData(v bool)`

SetStripAuthorizationData sets StripAuthorizationData field to given value.

### HasStripAuthorizationData

`func (o *OasAuthentication) HasStripAuthorizationData() bool`

HasStripAuthorizationData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


