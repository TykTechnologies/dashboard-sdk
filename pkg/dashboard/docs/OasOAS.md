# OasOAS

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Components** | Pointer to [**Openapi3Components**](Openapi3Components.md) |  | [optional] 
**ExternalDocs** | Pointer to [**Openapi3ExternalDocs**](Openapi3ExternalDocs.md) |  | [optional] 
**Info** | Pointer to [**Openapi3Info**](Openapi3Info.md) |  | [optional] 
**Openapi** | Pointer to **string** |  | [optional] 
**Paths** | Pointer to [**map[string]Openapi3PathItem**](Openapi3PathItem.md) |  | [optional] 
**Security** | Pointer to [**[]map[string][]string**](map[string][]string.md) |  | [optional] 
**Servers** | Pointer to [**[]Openapi3Server**](Openapi3Server.md) |  | [optional] 
**Tags** | Pointer to [**[]Openapi3Tag**](Openapi3Tag.md) |  | [optional] 

## Methods

### NewOasOAS

`func NewOasOAS() *OasOAS`

NewOasOAS instantiates a new OasOAS object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOasOASWithDefaults

`func NewOasOASWithDefaults() *OasOAS`

NewOasOASWithDefaults instantiates a new OasOAS object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComponents

`func (o *OasOAS) GetComponents() Openapi3Components`

GetComponents returns the Components field if non-nil, zero value otherwise.

### GetComponentsOk

`func (o *OasOAS) GetComponentsOk() (*Openapi3Components, bool)`

GetComponentsOk returns a tuple with the Components field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponents

`func (o *OasOAS) SetComponents(v Openapi3Components)`

SetComponents sets Components field to given value.

### HasComponents

`func (o *OasOAS) HasComponents() bool`

HasComponents returns a boolean if a field has been set.

### GetExternalDocs

`func (o *OasOAS) GetExternalDocs() Openapi3ExternalDocs`

GetExternalDocs returns the ExternalDocs field if non-nil, zero value otherwise.

### GetExternalDocsOk

`func (o *OasOAS) GetExternalDocsOk() (*Openapi3ExternalDocs, bool)`

GetExternalDocsOk returns a tuple with the ExternalDocs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalDocs

`func (o *OasOAS) SetExternalDocs(v Openapi3ExternalDocs)`

SetExternalDocs sets ExternalDocs field to given value.

### HasExternalDocs

`func (o *OasOAS) HasExternalDocs() bool`

HasExternalDocs returns a boolean if a field has been set.

### GetInfo

`func (o *OasOAS) GetInfo() Openapi3Info`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *OasOAS) GetInfoOk() (*Openapi3Info, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *OasOAS) SetInfo(v Openapi3Info)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *OasOAS) HasInfo() bool`

HasInfo returns a boolean if a field has been set.

### GetOpenapi

`func (o *OasOAS) GetOpenapi() string`

GetOpenapi returns the Openapi field if non-nil, zero value otherwise.

### GetOpenapiOk

`func (o *OasOAS) GetOpenapiOk() (*string, bool)`

GetOpenapiOk returns a tuple with the Openapi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenapi

`func (o *OasOAS) SetOpenapi(v string)`

SetOpenapi sets Openapi field to given value.

### HasOpenapi

`func (o *OasOAS) HasOpenapi() bool`

HasOpenapi returns a boolean if a field has been set.

### GetPaths

`func (o *OasOAS) GetPaths() map[string]Openapi3PathItem`

GetPaths returns the Paths field if non-nil, zero value otherwise.

### GetPathsOk

`func (o *OasOAS) GetPathsOk() (*map[string]Openapi3PathItem, bool)`

GetPathsOk returns a tuple with the Paths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaths

`func (o *OasOAS) SetPaths(v map[string]Openapi3PathItem)`

SetPaths sets Paths field to given value.

### HasPaths

`func (o *OasOAS) HasPaths() bool`

HasPaths returns a boolean if a field has been set.

### GetSecurity

`func (o *OasOAS) GetSecurity() []map[string][]string`

GetSecurity returns the Security field if non-nil, zero value otherwise.

### GetSecurityOk

`func (o *OasOAS) GetSecurityOk() (*[]map[string][]string, bool)`

GetSecurityOk returns a tuple with the Security field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurity

`func (o *OasOAS) SetSecurity(v []map[string][]string)`

SetSecurity sets Security field to given value.

### HasSecurity

`func (o *OasOAS) HasSecurity() bool`

HasSecurity returns a boolean if a field has been set.

### SetSecurityNil

`func (o *OasOAS) SetSecurityNil(b bool)`

 SetSecurityNil sets the value for Security to be an explicit nil

### UnsetSecurity
`func (o *OasOAS) UnsetSecurity()`

UnsetSecurity ensures that no value is present for Security, not even an explicit nil
### GetServers

`func (o *OasOAS) GetServers() []Openapi3Server`

GetServers returns the Servers field if non-nil, zero value otherwise.

### GetServersOk

`func (o *OasOAS) GetServersOk() (*[]Openapi3Server, bool)`

GetServersOk returns a tuple with the Servers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServers

`func (o *OasOAS) SetServers(v []Openapi3Server)`

SetServers sets Servers field to given value.

### HasServers

`func (o *OasOAS) HasServers() bool`

HasServers returns a boolean if a field has been set.

### SetServersNil

`func (o *OasOAS) SetServersNil(b bool)`

 SetServersNil sets the value for Servers to be an explicit nil

### UnsetServers
`func (o *OasOAS) UnsetServers()`

UnsetServers ensures that no value is present for Servers, not even an explicit nil
### GetTags

`func (o *OasOAS) GetTags() []Openapi3Tag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *OasOAS) GetTagsOk() (*[]Openapi3Tag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *OasOAS) SetTags(v []Openapi3Tag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *OasOAS) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


