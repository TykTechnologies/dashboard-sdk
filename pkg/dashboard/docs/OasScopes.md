# OasScopes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClaimName** | Pointer to **string** |  | [optional] 
**ScopeToPolicyMapping** | Pointer to [**[]OasScopeToPolicy**](OasScopeToPolicy.md) |  | [optional] 

## Methods

### NewOasScopes

`func NewOasScopes() *OasScopes`

NewOasScopes instantiates a new OasScopes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOasScopesWithDefaults

`func NewOasScopesWithDefaults() *OasScopes`

NewOasScopesWithDefaults instantiates a new OasScopes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClaimName

`func (o *OasScopes) GetClaimName() string`

GetClaimName returns the ClaimName field if non-nil, zero value otherwise.

### GetClaimNameOk

`func (o *OasScopes) GetClaimNameOk() (*string, bool)`

GetClaimNameOk returns a tuple with the ClaimName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaimName

`func (o *OasScopes) SetClaimName(v string)`

SetClaimName sets ClaimName field to given value.

### HasClaimName

`func (o *OasScopes) HasClaimName() bool`

HasClaimName returns a boolean if a field has been set.

### GetScopeToPolicyMapping

`func (o *OasScopes) GetScopeToPolicyMapping() []OasScopeToPolicy`

GetScopeToPolicyMapping returns the ScopeToPolicyMapping field if non-nil, zero value otherwise.

### GetScopeToPolicyMappingOk

`func (o *OasScopes) GetScopeToPolicyMappingOk() (*[]OasScopeToPolicy, bool)`

GetScopeToPolicyMappingOk returns a tuple with the ScopeToPolicyMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeToPolicyMapping

`func (o *OasScopes) SetScopeToPolicyMapping(v []OasScopeToPolicy)`

SetScopeToPolicyMapping sets ScopeToPolicyMapping field to given value.

### HasScopeToPolicyMapping

`func (o *OasScopes) HasScopeToPolicyMapping() bool`

HasScopeToPolicyMapping returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


