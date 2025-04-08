# CertificateDependencies

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CertID** | Pointer to **string** |  | [optional] 
**ClientCertApiNames** | Pointer to **[]string** |  | [optional] 
**ClientCertApis** | Pointer to **[]string** |  | [optional] 
**KeyCert** | Pointer to **string** |  | [optional] 
**UpstreamCertApiNames** | Pointer to **[]string** |  | [optional] 
**UpstreamCertApis** | Pointer to **[]string** |  | [optional] 

## Methods

### NewCertificateDependencies

`func NewCertificateDependencies() *CertificateDependencies`

NewCertificateDependencies instantiates a new CertificateDependencies object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertificateDependenciesWithDefaults

`func NewCertificateDependenciesWithDefaults() *CertificateDependencies`

NewCertificateDependenciesWithDefaults instantiates a new CertificateDependencies object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertID

`func (o *CertificateDependencies) GetCertID() string`

GetCertID returns the CertID field if non-nil, zero value otherwise.

### GetCertIDOk

`func (o *CertificateDependencies) GetCertIDOk() (*string, bool)`

GetCertIDOk returns a tuple with the CertID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertID

`func (o *CertificateDependencies) SetCertID(v string)`

SetCertID sets CertID field to given value.

### HasCertID

`func (o *CertificateDependencies) HasCertID() bool`

HasCertID returns a boolean if a field has been set.

### GetClientCertApiNames

`func (o *CertificateDependencies) GetClientCertApiNames() []string`

GetClientCertApiNames returns the ClientCertApiNames field if non-nil, zero value otherwise.

### GetClientCertApiNamesOk

`func (o *CertificateDependencies) GetClientCertApiNamesOk() (*[]string, bool)`

GetClientCertApiNamesOk returns a tuple with the ClientCertApiNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientCertApiNames

`func (o *CertificateDependencies) SetClientCertApiNames(v []string)`

SetClientCertApiNames sets ClientCertApiNames field to given value.

### HasClientCertApiNames

`func (o *CertificateDependencies) HasClientCertApiNames() bool`

HasClientCertApiNames returns a boolean if a field has been set.

### SetClientCertApiNamesNil

`func (o *CertificateDependencies) SetClientCertApiNamesNil(b bool)`

 SetClientCertApiNamesNil sets the value for ClientCertApiNames to be an explicit nil

### UnsetClientCertApiNames
`func (o *CertificateDependencies) UnsetClientCertApiNames()`

UnsetClientCertApiNames ensures that no value is present for ClientCertApiNames, not even an explicit nil
### GetClientCertApis

`func (o *CertificateDependencies) GetClientCertApis() []string`

GetClientCertApis returns the ClientCertApis field if non-nil, zero value otherwise.

### GetClientCertApisOk

`func (o *CertificateDependencies) GetClientCertApisOk() (*[]string, bool)`

GetClientCertApisOk returns a tuple with the ClientCertApis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientCertApis

`func (o *CertificateDependencies) SetClientCertApis(v []string)`

SetClientCertApis sets ClientCertApis field to given value.

### HasClientCertApis

`func (o *CertificateDependencies) HasClientCertApis() bool`

HasClientCertApis returns a boolean if a field has been set.

### SetClientCertApisNil

`func (o *CertificateDependencies) SetClientCertApisNil(b bool)`

 SetClientCertApisNil sets the value for ClientCertApis to be an explicit nil

### UnsetClientCertApis
`func (o *CertificateDependencies) UnsetClientCertApis()`

UnsetClientCertApis ensures that no value is present for ClientCertApis, not even an explicit nil
### GetKeyCert

`func (o *CertificateDependencies) GetKeyCert() string`

GetKeyCert returns the KeyCert field if non-nil, zero value otherwise.

### GetKeyCertOk

`func (o *CertificateDependencies) GetKeyCertOk() (*string, bool)`

GetKeyCertOk returns a tuple with the KeyCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyCert

`func (o *CertificateDependencies) SetKeyCert(v string)`

SetKeyCert sets KeyCert field to given value.

### HasKeyCert

`func (o *CertificateDependencies) HasKeyCert() bool`

HasKeyCert returns a boolean if a field has been set.

### GetUpstreamCertApiNames

`func (o *CertificateDependencies) GetUpstreamCertApiNames() []string`

GetUpstreamCertApiNames returns the UpstreamCertApiNames field if non-nil, zero value otherwise.

### GetUpstreamCertApiNamesOk

`func (o *CertificateDependencies) GetUpstreamCertApiNamesOk() (*[]string, bool)`

GetUpstreamCertApiNamesOk returns a tuple with the UpstreamCertApiNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpstreamCertApiNames

`func (o *CertificateDependencies) SetUpstreamCertApiNames(v []string)`

SetUpstreamCertApiNames sets UpstreamCertApiNames field to given value.

### HasUpstreamCertApiNames

`func (o *CertificateDependencies) HasUpstreamCertApiNames() bool`

HasUpstreamCertApiNames returns a boolean if a field has been set.

### SetUpstreamCertApiNamesNil

`func (o *CertificateDependencies) SetUpstreamCertApiNamesNil(b bool)`

 SetUpstreamCertApiNamesNil sets the value for UpstreamCertApiNames to be an explicit nil

### UnsetUpstreamCertApiNames
`func (o *CertificateDependencies) UnsetUpstreamCertApiNames()`

UnsetUpstreamCertApiNames ensures that no value is present for UpstreamCertApiNames, not even an explicit nil
### GetUpstreamCertApis

`func (o *CertificateDependencies) GetUpstreamCertApis() []string`

GetUpstreamCertApis returns the UpstreamCertApis field if non-nil, zero value otherwise.

### GetUpstreamCertApisOk

`func (o *CertificateDependencies) GetUpstreamCertApisOk() (*[]string, bool)`

GetUpstreamCertApisOk returns a tuple with the UpstreamCertApis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpstreamCertApis

`func (o *CertificateDependencies) SetUpstreamCertApis(v []string)`

SetUpstreamCertApis sets UpstreamCertApis field to given value.

### HasUpstreamCertApis

`func (o *CertificateDependencies) HasUpstreamCertApis() bool`

HasUpstreamCertApis returns a boolean if a field has been set.

### SetUpstreamCertApisNil

`func (o *CertificateDependencies) SetUpstreamCertApisNil(b bool)`

 SetUpstreamCertApisNil sets the value for UpstreamCertApis to be an explicit nil

### UnsetUpstreamCertApis
`func (o *CertificateDependencies) UnsetUpstreamCertApis()`

UnsetUpstreamCertApis ensures that no value is present for UpstreamCertApis, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


