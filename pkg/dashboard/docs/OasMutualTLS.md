# OasMutualTLS

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DomainToCertificateMapping** | Pointer to [**[]OasDomainToCertificate**](OasDomainToCertificate.md) |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 

## Methods

### NewOasMutualTLS

`func NewOasMutualTLS() *OasMutualTLS`

NewOasMutualTLS instantiates a new OasMutualTLS object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOasMutualTLSWithDefaults

`func NewOasMutualTLSWithDefaults() *OasMutualTLS`

NewOasMutualTLSWithDefaults instantiates a new OasMutualTLS object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomainToCertificateMapping

`func (o *OasMutualTLS) GetDomainToCertificateMapping() []OasDomainToCertificate`

GetDomainToCertificateMapping returns the DomainToCertificateMapping field if non-nil, zero value otherwise.

### GetDomainToCertificateMappingOk

`func (o *OasMutualTLS) GetDomainToCertificateMappingOk() (*[]OasDomainToCertificate, bool)`

GetDomainToCertificateMappingOk returns a tuple with the DomainToCertificateMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainToCertificateMapping

`func (o *OasMutualTLS) SetDomainToCertificateMapping(v []OasDomainToCertificate)`

SetDomainToCertificateMapping sets DomainToCertificateMapping field to given value.

### HasDomainToCertificateMapping

`func (o *OasMutualTLS) HasDomainToCertificateMapping() bool`

HasDomainToCertificateMapping returns a boolean if a field has been set.

### SetDomainToCertificateMappingNil

`func (o *OasMutualTLS) SetDomainToCertificateMappingNil(b bool)`

 SetDomainToCertificateMappingNil sets the value for DomainToCertificateMapping to be an explicit nil

### UnsetDomainToCertificateMapping
`func (o *OasMutualTLS) UnsetDomainToCertificateMapping()`

UnsetDomainToCertificateMapping ensures that no value is present for DomainToCertificateMapping, not even an explicit nil
### GetEnabled

`func (o *OasMutualTLS) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *OasMutualTLS) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *OasMutualTLS) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *OasMutualTLS) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


