# Scopes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClaimName** | Pointer to **string** |  | [optional] 
**ScopeToPolicyMapping** | Pointer to [**[]ScopeToPolicy**](ScopeToPolicy.md) |  | [optional] 

## Methods

### NewScopes

`func NewScopes() *Scopes`

NewScopes instantiates a new Scopes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScopesWithDefaults

`func NewScopesWithDefaults() *Scopes`

NewScopesWithDefaults instantiates a new Scopes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClaimName

`func (o *Scopes) GetClaimName() string`

GetClaimName returns the ClaimName field if non-nil, zero value otherwise.

### GetClaimNameOk

`func (o *Scopes) GetClaimNameOk() (*string, bool)`

GetClaimNameOk returns a tuple with the ClaimName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaimName

`func (o *Scopes) SetClaimName(v string)`

SetClaimName sets ClaimName field to given value.

### HasClaimName

`func (o *Scopes) HasClaimName() bool`

HasClaimName returns a boolean if a field has been set.

### GetScopeToPolicyMapping

`func (o *Scopes) GetScopeToPolicyMapping() []ScopeToPolicy`

GetScopeToPolicyMapping returns the ScopeToPolicyMapping field if non-nil, zero value otherwise.

### GetScopeToPolicyMappingOk

`func (o *Scopes) GetScopeToPolicyMappingOk() (*[]ScopeToPolicy, bool)`

GetScopeToPolicyMappingOk returns a tuple with the ScopeToPolicyMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeToPolicyMapping

`func (o *Scopes) SetScopeToPolicyMapping(v []ScopeToPolicy)`

SetScopeToPolicyMapping sets ScopeToPolicyMapping field to given value.

### HasScopeToPolicyMapping

`func (o *Scopes) HasScopeToPolicyMapping() bool`

HasScopeToPolicyMapping returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


