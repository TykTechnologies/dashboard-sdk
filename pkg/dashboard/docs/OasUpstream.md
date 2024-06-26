# OasUpstream

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CertificatePinning** | Pointer to [**OasCertificatePinning**](OasCertificatePinning.md) |  | [optional] 
**MutualTLS** | Pointer to [**OasMutualTLS**](OasMutualTLS.md) |  | [optional] 
**RateLimit** | Pointer to [**OasRateLimit**](OasRateLimit.md) |  | [optional] 
**ServiceDiscovery** | Pointer to [**OasServiceDiscovery**](OasServiceDiscovery.md) |  | [optional] 
**Test** | Pointer to [**OasTest**](OasTest.md) |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 

## Methods

### NewOasUpstream

`func NewOasUpstream() *OasUpstream`

NewOasUpstream instantiates a new OasUpstream object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOasUpstreamWithDefaults

`func NewOasUpstreamWithDefaults() *OasUpstream`

NewOasUpstreamWithDefaults instantiates a new OasUpstream object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertificatePinning

`func (o *OasUpstream) GetCertificatePinning() OasCertificatePinning`

GetCertificatePinning returns the CertificatePinning field if non-nil, zero value otherwise.

### GetCertificatePinningOk

`func (o *OasUpstream) GetCertificatePinningOk() (*OasCertificatePinning, bool)`

GetCertificatePinningOk returns a tuple with the CertificatePinning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificatePinning

`func (o *OasUpstream) SetCertificatePinning(v OasCertificatePinning)`

SetCertificatePinning sets CertificatePinning field to given value.

### HasCertificatePinning

`func (o *OasUpstream) HasCertificatePinning() bool`

HasCertificatePinning returns a boolean if a field has been set.

### GetMutualTLS

`func (o *OasUpstream) GetMutualTLS() OasMutualTLS`

GetMutualTLS returns the MutualTLS field if non-nil, zero value otherwise.

### GetMutualTLSOk

`func (o *OasUpstream) GetMutualTLSOk() (*OasMutualTLS, bool)`

GetMutualTLSOk returns a tuple with the MutualTLS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMutualTLS

`func (o *OasUpstream) SetMutualTLS(v OasMutualTLS)`

SetMutualTLS sets MutualTLS field to given value.

### HasMutualTLS

`func (o *OasUpstream) HasMutualTLS() bool`

HasMutualTLS returns a boolean if a field has been set.

### GetRateLimit

`func (o *OasUpstream) GetRateLimit() OasRateLimit`

GetRateLimit returns the RateLimit field if non-nil, zero value otherwise.

### GetRateLimitOk

`func (o *OasUpstream) GetRateLimitOk() (*OasRateLimit, bool)`

GetRateLimitOk returns a tuple with the RateLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateLimit

`func (o *OasUpstream) SetRateLimit(v OasRateLimit)`

SetRateLimit sets RateLimit field to given value.

### HasRateLimit

`func (o *OasUpstream) HasRateLimit() bool`

HasRateLimit returns a boolean if a field has been set.

### GetServiceDiscovery

`func (o *OasUpstream) GetServiceDiscovery() OasServiceDiscovery`

GetServiceDiscovery returns the ServiceDiscovery field if non-nil, zero value otherwise.

### GetServiceDiscoveryOk

`func (o *OasUpstream) GetServiceDiscoveryOk() (*OasServiceDiscovery, bool)`

GetServiceDiscoveryOk returns a tuple with the ServiceDiscovery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceDiscovery

`func (o *OasUpstream) SetServiceDiscovery(v OasServiceDiscovery)`

SetServiceDiscovery sets ServiceDiscovery field to given value.

### HasServiceDiscovery

`func (o *OasUpstream) HasServiceDiscovery() bool`

HasServiceDiscovery returns a boolean if a field has been set.

### GetTest

`func (o *OasUpstream) GetTest() OasTest`

GetTest returns the Test field if non-nil, zero value otherwise.

### GetTestOk

`func (o *OasUpstream) GetTestOk() (*OasTest, bool)`

GetTestOk returns a tuple with the Test field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTest

`func (o *OasUpstream) SetTest(v OasTest)`

SetTest sets Test field to given value.

### HasTest

`func (o *OasUpstream) HasTest() bool`

HasTest returns a boolean if a field has been set.

### GetUrl

`func (o *OasUpstream) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *OasUpstream) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *OasUpstream) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *OasUpstream) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


