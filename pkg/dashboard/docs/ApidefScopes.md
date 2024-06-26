# ApidefScopes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Jwt** | Pointer to [**ApidefScopeClaim**](ApidefScopeClaim.md) |  | [optional] 
**Oidc** | Pointer to [**ApidefScopeClaim**](ApidefScopeClaim.md) |  | [optional] 

## Methods

### NewApidefScopes

`func NewApidefScopes() *ApidefScopes`

NewApidefScopes instantiates a new ApidefScopes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApidefScopesWithDefaults

`func NewApidefScopesWithDefaults() *ApidefScopes`

NewApidefScopesWithDefaults instantiates a new ApidefScopes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJwt

`func (o *ApidefScopes) GetJwt() ApidefScopeClaim`

GetJwt returns the Jwt field if non-nil, zero value otherwise.

### GetJwtOk

`func (o *ApidefScopes) GetJwtOk() (*ApidefScopeClaim, bool)`

GetJwtOk returns a tuple with the Jwt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwt

`func (o *ApidefScopes) SetJwt(v ApidefScopeClaim)`

SetJwt sets Jwt field to given value.

### HasJwt

`func (o *ApidefScopes) HasJwt() bool`

HasJwt returns a boolean if a field has been set.

### GetOidc

`func (o *ApidefScopes) GetOidc() ApidefScopeClaim`

GetOidc returns the Oidc field if non-nil, zero value otherwise.

### GetOidcOk

`func (o *ApidefScopes) GetOidcOk() (*ApidefScopeClaim, bool)`

GetOidcOk returns a tuple with the Oidc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOidc

`func (o *ApidefScopes) SetOidc(v ApidefScopeClaim)`

SetOidc sets Oidc field to given value.

### HasOidc

`func (o *ApidefScopes) HasOidc() bool`

HasOidc returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


