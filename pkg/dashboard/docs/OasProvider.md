# OasProvider

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientToPolicyMapping** | Pointer to [**[]OasClientToPolicy**](OasClientToPolicy.md) |  | [optional] 
**Issuer** | Pointer to **string** |  | [optional] 

## Methods

### NewOasProvider

`func NewOasProvider() *OasProvider`

NewOasProvider instantiates a new OasProvider object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOasProviderWithDefaults

`func NewOasProviderWithDefaults() *OasProvider`

NewOasProviderWithDefaults instantiates a new OasProvider object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientToPolicyMapping

`func (o *OasProvider) GetClientToPolicyMapping() []OasClientToPolicy`

GetClientToPolicyMapping returns the ClientToPolicyMapping field if non-nil, zero value otherwise.

### GetClientToPolicyMappingOk

`func (o *OasProvider) GetClientToPolicyMappingOk() (*[]OasClientToPolicy, bool)`

GetClientToPolicyMappingOk returns a tuple with the ClientToPolicyMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientToPolicyMapping

`func (o *OasProvider) SetClientToPolicyMapping(v []OasClientToPolicy)`

SetClientToPolicyMapping sets ClientToPolicyMapping field to given value.

### HasClientToPolicyMapping

`func (o *OasProvider) HasClientToPolicyMapping() bool`

HasClientToPolicyMapping returns a boolean if a field has been set.

### GetIssuer

`func (o *OasProvider) GetIssuer() string`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *OasProvider) GetIssuerOk() (*string, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *OasProvider) SetIssuer(v string)`

SetIssuer sets Issuer field to given value.

### HasIssuer

`func (o *OasProvider) HasIssuer() bool`

HasIssuer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


