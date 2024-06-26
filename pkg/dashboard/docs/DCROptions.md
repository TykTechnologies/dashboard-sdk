# DCROptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessToken** | Pointer to **string** |  | [optional] 
**GrantTypes** | Pointer to **[]string** |  | [optional] 
**IdpHost** | Pointer to **string** |  | [optional] 
**Provider** | Pointer to **string** |  | [optional] 
**RegistrationEndpoint** | Pointer to **string** |  | [optional] 
**ResponseTypes** | Pointer to **[]string** |  | [optional] 
**TokenEndpointAuthMethod** | Pointer to **string** |  | [optional] 

## Methods

### NewDCROptions

`func NewDCROptions() *DCROptions`

NewDCROptions instantiates a new DCROptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDCROptionsWithDefaults

`func NewDCROptionsWithDefaults() *DCROptions`

NewDCROptionsWithDefaults instantiates a new DCROptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessToken

`func (o *DCROptions) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *DCROptions) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *DCROptions) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.

### HasAccessToken

`func (o *DCROptions) HasAccessToken() bool`

HasAccessToken returns a boolean if a field has been set.

### GetGrantTypes

`func (o *DCROptions) GetGrantTypes() []string`

GetGrantTypes returns the GrantTypes field if non-nil, zero value otherwise.

### GetGrantTypesOk

`func (o *DCROptions) GetGrantTypesOk() (*[]string, bool)`

GetGrantTypesOk returns a tuple with the GrantTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantTypes

`func (o *DCROptions) SetGrantTypes(v []string)`

SetGrantTypes sets GrantTypes field to given value.

### HasGrantTypes

`func (o *DCROptions) HasGrantTypes() bool`

HasGrantTypes returns a boolean if a field has been set.

### SetGrantTypesNil

`func (o *DCROptions) SetGrantTypesNil(b bool)`

 SetGrantTypesNil sets the value for GrantTypes to be an explicit nil

### UnsetGrantTypes
`func (o *DCROptions) UnsetGrantTypes()`

UnsetGrantTypes ensures that no value is present for GrantTypes, not even an explicit nil
### GetIdpHost

`func (o *DCROptions) GetIdpHost() string`

GetIdpHost returns the IdpHost field if non-nil, zero value otherwise.

### GetIdpHostOk

`func (o *DCROptions) GetIdpHostOk() (*string, bool)`

GetIdpHostOk returns a tuple with the IdpHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpHost

`func (o *DCROptions) SetIdpHost(v string)`

SetIdpHost sets IdpHost field to given value.

### HasIdpHost

`func (o *DCROptions) HasIdpHost() bool`

HasIdpHost returns a boolean if a field has been set.

### GetProvider

`func (o *DCROptions) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *DCROptions) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *DCROptions) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *DCROptions) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetRegistrationEndpoint

`func (o *DCROptions) GetRegistrationEndpoint() string`

GetRegistrationEndpoint returns the RegistrationEndpoint field if non-nil, zero value otherwise.

### GetRegistrationEndpointOk

`func (o *DCROptions) GetRegistrationEndpointOk() (*string, bool)`

GetRegistrationEndpointOk returns a tuple with the RegistrationEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrationEndpoint

`func (o *DCROptions) SetRegistrationEndpoint(v string)`

SetRegistrationEndpoint sets RegistrationEndpoint field to given value.

### HasRegistrationEndpoint

`func (o *DCROptions) HasRegistrationEndpoint() bool`

HasRegistrationEndpoint returns a boolean if a field has been set.

### GetResponseTypes

`func (o *DCROptions) GetResponseTypes() []string`

GetResponseTypes returns the ResponseTypes field if non-nil, zero value otherwise.

### GetResponseTypesOk

`func (o *DCROptions) GetResponseTypesOk() (*[]string, bool)`

GetResponseTypesOk returns a tuple with the ResponseTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseTypes

`func (o *DCROptions) SetResponseTypes(v []string)`

SetResponseTypes sets ResponseTypes field to given value.

### HasResponseTypes

`func (o *DCROptions) HasResponseTypes() bool`

HasResponseTypes returns a boolean if a field has been set.

### SetResponseTypesNil

`func (o *DCROptions) SetResponseTypesNil(b bool)`

 SetResponseTypesNil sets the value for ResponseTypes to be an explicit nil

### UnsetResponseTypes
`func (o *DCROptions) UnsetResponseTypes()`

UnsetResponseTypes ensures that no value is present for ResponseTypes, not even an explicit nil
### GetTokenEndpointAuthMethod

`func (o *DCROptions) GetTokenEndpointAuthMethod() string`

GetTokenEndpointAuthMethod returns the TokenEndpointAuthMethod field if non-nil, zero value otherwise.

### GetTokenEndpointAuthMethodOk

`func (o *DCROptions) GetTokenEndpointAuthMethodOk() (*string, bool)`

GetTokenEndpointAuthMethodOk returns a tuple with the TokenEndpointAuthMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenEndpointAuthMethod

`func (o *DCROptions) SetTokenEndpointAuthMethod(v string)`

SetTokenEndpointAuthMethod sets TokenEndpointAuthMethod field to given value.

### HasTokenEndpointAuthMethod

`func (o *DCROptions) HasTokenEndpointAuthMethod() bool`

HasTokenEndpointAuthMethod returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


