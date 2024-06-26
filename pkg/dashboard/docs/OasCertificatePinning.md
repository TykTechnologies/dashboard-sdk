# OasCertificatePinning

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DomainToPublicKeysMapping** | Pointer to [**[]OasPinnedPublicKey**](OasPinnedPublicKey.md) |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 

## Methods

### NewOasCertificatePinning

`func NewOasCertificatePinning() *OasCertificatePinning`

NewOasCertificatePinning instantiates a new OasCertificatePinning object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOasCertificatePinningWithDefaults

`func NewOasCertificatePinningWithDefaults() *OasCertificatePinning`

NewOasCertificatePinningWithDefaults instantiates a new OasCertificatePinning object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomainToPublicKeysMapping

`func (o *OasCertificatePinning) GetDomainToPublicKeysMapping() []OasPinnedPublicKey`

GetDomainToPublicKeysMapping returns the DomainToPublicKeysMapping field if non-nil, zero value otherwise.

### GetDomainToPublicKeysMappingOk

`func (o *OasCertificatePinning) GetDomainToPublicKeysMappingOk() (*[]OasPinnedPublicKey, bool)`

GetDomainToPublicKeysMappingOk returns a tuple with the DomainToPublicKeysMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainToPublicKeysMapping

`func (o *OasCertificatePinning) SetDomainToPublicKeysMapping(v []OasPinnedPublicKey)`

SetDomainToPublicKeysMapping sets DomainToPublicKeysMapping field to given value.

### HasDomainToPublicKeysMapping

`func (o *OasCertificatePinning) HasDomainToPublicKeysMapping() bool`

HasDomainToPublicKeysMapping returns a boolean if a field has been set.

### SetDomainToPublicKeysMappingNil

`func (o *OasCertificatePinning) SetDomainToPublicKeysMappingNil(b bool)`

 SetDomainToPublicKeysMappingNil sets the value for DomainToPublicKeysMapping to be an explicit nil

### UnsetDomainToPublicKeysMapping
`func (o *OasCertificatePinning) UnsetDomainToPublicKeysMapping()`

UnsetDomainToPublicKeysMapping ensures that no value is present for DomainToPublicKeysMapping, not even an explicit nil
### GetEnabled

`func (o *OasCertificatePinning) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *OasCertificatePinning) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *OasCertificatePinning) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *OasCertificatePinning) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


