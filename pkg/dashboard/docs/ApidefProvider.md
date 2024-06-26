# ApidefProvider

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Introspection** | Pointer to [**ApidefIntrospection**](ApidefIntrospection.md) |  | [optional] 
**Jwt** | Pointer to [**ApidefJWTValidation**](ApidefJWTValidation.md) |  | [optional] 

## Methods

### NewApidefProvider

`func NewApidefProvider() *ApidefProvider`

NewApidefProvider instantiates a new ApidefProvider object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApidefProviderWithDefaults

`func NewApidefProviderWithDefaults() *ApidefProvider`

NewApidefProviderWithDefaults instantiates a new ApidefProvider object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIntrospection

`func (o *ApidefProvider) GetIntrospection() ApidefIntrospection`

GetIntrospection returns the Introspection field if non-nil, zero value otherwise.

### GetIntrospectionOk

`func (o *ApidefProvider) GetIntrospectionOk() (*ApidefIntrospection, bool)`

GetIntrospectionOk returns a tuple with the Introspection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntrospection

`func (o *ApidefProvider) SetIntrospection(v ApidefIntrospection)`

SetIntrospection sets Introspection field to given value.

### HasIntrospection

`func (o *ApidefProvider) HasIntrospection() bool`

HasIntrospection returns a boolean if a field has been set.

### GetJwt

`func (o *ApidefProvider) GetJwt() ApidefJWTValidation`

GetJwt returns the Jwt field if non-nil, zero value otherwise.

### GetJwtOk

`func (o *ApidefProvider) GetJwtOk() (*ApidefJWTValidation, bool)`

GetJwtOk returns a tuple with the Jwt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwt

`func (o *ApidefProvider) SetJwt(v ApidefJWTValidation)`

SetJwt sets Jwt field to given value.

### HasJwt

`func (o *ApidefProvider) HasJwt() bool`

HasJwt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


