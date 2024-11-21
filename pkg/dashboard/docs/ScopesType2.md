# ScopesType2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Jwt** | Pointer to [**ScopeClaim**](ScopeClaim.md) |  | [optional] 
**Oidc** | Pointer to [**ScopeClaim**](ScopeClaim.md) |  | [optional] 

## Methods

### NewScopesType2

`func NewScopesType2() *ScopesType2`

NewScopesType2 instantiates a new ScopesType2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScopesType2WithDefaults

`func NewScopesType2WithDefaults() *ScopesType2`

NewScopesType2WithDefaults instantiates a new ScopesType2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJwt

`func (o *ScopesType2) GetJwt() ScopeClaim`

GetJwt returns the Jwt field if non-nil, zero value otherwise.

### GetJwtOk

`func (o *ScopesType2) GetJwtOk() (*ScopeClaim, bool)`

GetJwtOk returns a tuple with the Jwt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwt

`func (o *ScopesType2) SetJwt(v ScopeClaim)`

SetJwt sets Jwt field to given value.

### HasJwt

`func (o *ScopesType2) HasJwt() bool`

HasJwt returns a boolean if a field has been set.

### GetOidc

`func (o *ScopesType2) GetOidc() ScopeClaim`

GetOidc returns the Oidc field if non-nil, zero value otherwise.

### GetOidcOk

`func (o *ScopesType2) GetOidcOk() (*ScopeClaim, bool)`

GetOidcOk returns a tuple with the Oidc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOidc

`func (o *ScopesType2) SetOidc(v ScopeClaim)`

SetOidc sets Oidc field to given value.

### HasOidc

`func (o *ScopesType2) HasOidc() bool`

HasOidc returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


