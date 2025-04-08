# ListCertificates200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CertBasics** | Pointer to [**[]CertsCertificateBasics**](CertsCertificateBasics.md) |  | [optional] 
**Pages** | Pointer to **int32** |  | [optional] 
**Certs** | Pointer to **[]string** |  | [optional] 

## Methods

### NewListCertificates200Response

`func NewListCertificates200Response() *ListCertificates200Response`

NewListCertificates200Response instantiates a new ListCertificates200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListCertificates200ResponseWithDefaults

`func NewListCertificates200ResponseWithDefaults() *ListCertificates200Response`

NewListCertificates200ResponseWithDefaults instantiates a new ListCertificates200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertBasics

`func (o *ListCertificates200Response) GetCertBasics() []CertsCertificateBasics`

GetCertBasics returns the CertBasics field if non-nil, zero value otherwise.

### GetCertBasicsOk

`func (o *ListCertificates200Response) GetCertBasicsOk() (*[]CertsCertificateBasics, bool)`

GetCertBasicsOk returns a tuple with the CertBasics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertBasics

`func (o *ListCertificates200Response) SetCertBasics(v []CertsCertificateBasics)`

SetCertBasics sets CertBasics field to given value.

### HasCertBasics

`func (o *ListCertificates200Response) HasCertBasics() bool`

HasCertBasics returns a boolean if a field has been set.

### SetCertBasicsNil

`func (o *ListCertificates200Response) SetCertBasicsNil(b bool)`

 SetCertBasicsNil sets the value for CertBasics to be an explicit nil

### UnsetCertBasics
`func (o *ListCertificates200Response) UnsetCertBasics()`

UnsetCertBasics ensures that no value is present for CertBasics, not even an explicit nil
### GetPages

`func (o *ListCertificates200Response) GetPages() int32`

GetPages returns the Pages field if non-nil, zero value otherwise.

### GetPagesOk

`func (o *ListCertificates200Response) GetPagesOk() (*int32, bool)`

GetPagesOk returns a tuple with the Pages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPages

`func (o *ListCertificates200Response) SetPages(v int32)`

SetPages sets Pages field to given value.

### HasPages

`func (o *ListCertificates200Response) HasPages() bool`

HasPages returns a boolean if a field has been set.

### GetCerts

`func (o *ListCertificates200Response) GetCerts() []string`

GetCerts returns the Certs field if non-nil, zero value otherwise.

### GetCertsOk

`func (o *ListCertificates200Response) GetCertsOk() (*[]string, bool)`

GetCertsOk returns a tuple with the Certs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCerts

`func (o *ListCertificates200Response) SetCerts(v []string)`

SetCerts sets Certs field to given value.

### HasCerts

`func (o *ListCertificates200Response) HasCerts() bool`

HasCerts returns a boolean if a field has been set.

### SetCertsNil

`func (o *ListCertificates200Response) SetCertsNil(b bool)`

 SetCertsNil sets the value for Certs to be an explicit nil

### UnsetCerts
`func (o *ListCertificates200Response) UnsetCerts()`

UnsetCerts ensures that no value is present for Certs, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


