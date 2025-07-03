# CertificateDetailedList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Certificates** | Pointer to [**[]CertsCertificateMeta**](CertsCertificateMeta.md) |  | [optional] 
**Pages** | Pointer to **int32** |  | [optional] 

## Methods

### NewCertificateDetailedList

`func NewCertificateDetailedList() *CertificateDetailedList`

NewCertificateDetailedList instantiates a new CertificateDetailedList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertificateDetailedListWithDefaults

`func NewCertificateDetailedListWithDefaults() *CertificateDetailedList`

NewCertificateDetailedListWithDefaults instantiates a new CertificateDetailedList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertificates

`func (o *CertificateDetailedList) GetCertificates() []CertsCertificateMeta`

GetCertificates returns the Certificates field if non-nil, zero value otherwise.

### GetCertificatesOk

`func (o *CertificateDetailedList) GetCertificatesOk() (*[]CertsCertificateMeta, bool)`

GetCertificatesOk returns a tuple with the Certificates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificates

`func (o *CertificateDetailedList) SetCertificates(v []CertsCertificateMeta)`

SetCertificates sets Certificates field to given value.

### HasCertificates

`func (o *CertificateDetailedList) HasCertificates() bool`

HasCertificates returns a boolean if a field has been set.

### SetCertificatesNil

`func (o *CertificateDetailedList) SetCertificatesNil(b bool)`

 SetCertificatesNil sets the value for Certificates to be an explicit nil

### UnsetCertificates
`func (o *CertificateDetailedList) UnsetCertificates()`

UnsetCertificates ensures that no value is present for Certificates, not even an explicit nil
### GetPages

`func (o *CertificateDetailedList) GetPages() int32`

GetPages returns the Pages field if non-nil, zero value otherwise.

### GetPagesOk

`func (o *CertificateDetailedList) GetPagesOk() (*int32, bool)`

GetPagesOk returns a tuple with the Pages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPages

`func (o *CertificateDetailedList) SetPages(v int32)`

SetPages sets Pages field to given value.

### HasPages

`func (o *CertificateDetailedList) HasPages() bool`

HasPages returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


