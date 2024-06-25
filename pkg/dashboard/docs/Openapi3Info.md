# Openapi3Info

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Contact** | Pointer to [**Openapi3Contact**](Openapi3Contact.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**License** | Pointer to [**Openapi3License**](Openapi3License.md) |  | [optional] 
**TermsOfService** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 

## Methods

### NewOpenapi3Info

`func NewOpenapi3Info() *Openapi3Info`

NewOpenapi3Info instantiates a new Openapi3Info object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpenapi3InfoWithDefaults

`func NewOpenapi3InfoWithDefaults() *Openapi3Info`

NewOpenapi3InfoWithDefaults instantiates a new Openapi3Info object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContact

`func (o *Openapi3Info) GetContact() Openapi3Contact`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *Openapi3Info) GetContactOk() (*Openapi3Contact, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *Openapi3Info) SetContact(v Openapi3Contact)`

SetContact sets Contact field to given value.

### HasContact

`func (o *Openapi3Info) HasContact() bool`

HasContact returns a boolean if a field has been set.

### GetDescription

`func (o *Openapi3Info) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Openapi3Info) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Openapi3Info) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Openapi3Info) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLicense

`func (o *Openapi3Info) GetLicense() Openapi3License`

GetLicense returns the License field if non-nil, zero value otherwise.

### GetLicenseOk

`func (o *Openapi3Info) GetLicenseOk() (*Openapi3License, bool)`

GetLicenseOk returns a tuple with the License field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicense

`func (o *Openapi3Info) SetLicense(v Openapi3License)`

SetLicense sets License field to given value.

### HasLicense

`func (o *Openapi3Info) HasLicense() bool`

HasLicense returns a boolean if a field has been set.

### GetTermsOfService

`func (o *Openapi3Info) GetTermsOfService() string`

GetTermsOfService returns the TermsOfService field if non-nil, zero value otherwise.

### GetTermsOfServiceOk

`func (o *Openapi3Info) GetTermsOfServiceOk() (*string, bool)`

GetTermsOfServiceOk returns a tuple with the TermsOfService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTermsOfService

`func (o *Openapi3Info) SetTermsOfService(v string)`

SetTermsOfService sets TermsOfService field to given value.

### HasTermsOfService

`func (o *Openapi3Info) HasTermsOfService() bool`

HasTermsOfService returns a boolean if a field has been set.

### GetTitle

`func (o *Openapi3Info) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Openapi3Info) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Openapi3Info) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *Openapi3Info) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetVersion

`func (o *Openapi3Info) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *Openapi3Info) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *Openapi3Info) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *Openapi3Info) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


