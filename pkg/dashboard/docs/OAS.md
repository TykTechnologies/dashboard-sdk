# OAS

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Components** | Pointer to [**NullableComponents**](Components.md) |  | [optional] 
**ExternalDocs** | Pointer to [**NullableExternalDocs**](ExternalDocs.md) |  | [optional] 
**Info** | Pointer to [**NullableInfoType2**](InfoType2.md) |  | [optional] 
**Openapi** | Pointer to **string** |  | [optional] 
**Paths** | Pointer to [**map[string]PathItem**](PathItem.md) |  | [optional] 
**Security** | Pointer to [**[]map[string][]string**](map[string][]string.md) |  | [optional] 
**Servers** | Pointer to [**[]ServerType2**](ServerType2.md) |  | [optional] 
**Tags** | Pointer to [**[]Tag**](Tag.md) |  | [optional] 

## Methods

### NewOAS

`func NewOAS() *OAS`

NewOAS instantiates a new OAS object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOASWithDefaults

`func NewOASWithDefaults() *OAS`

NewOASWithDefaults instantiates a new OAS object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComponents

`func (o *OAS) GetComponents() Components`

GetComponents returns the Components field if non-nil, zero value otherwise.

### GetComponentsOk

`func (o *OAS) GetComponentsOk() (*Components, bool)`

GetComponentsOk returns a tuple with the Components field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponents

`func (o *OAS) SetComponents(v Components)`

SetComponents sets Components field to given value.

### HasComponents

`func (o *OAS) HasComponents() bool`

HasComponents returns a boolean if a field has been set.

### SetComponentsNil

`func (o *OAS) SetComponentsNil(b bool)`

 SetComponentsNil sets the value for Components to be an explicit nil

### UnsetComponents
`func (o *OAS) UnsetComponents()`

UnsetComponents ensures that no value is present for Components, not even an explicit nil
### GetExternalDocs

`func (o *OAS) GetExternalDocs() ExternalDocs`

GetExternalDocs returns the ExternalDocs field if non-nil, zero value otherwise.

### GetExternalDocsOk

`func (o *OAS) GetExternalDocsOk() (*ExternalDocs, bool)`

GetExternalDocsOk returns a tuple with the ExternalDocs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalDocs

`func (o *OAS) SetExternalDocs(v ExternalDocs)`

SetExternalDocs sets ExternalDocs field to given value.

### HasExternalDocs

`func (o *OAS) HasExternalDocs() bool`

HasExternalDocs returns a boolean if a field has been set.

### SetExternalDocsNil

`func (o *OAS) SetExternalDocsNil(b bool)`

 SetExternalDocsNil sets the value for ExternalDocs to be an explicit nil

### UnsetExternalDocs
`func (o *OAS) UnsetExternalDocs()`

UnsetExternalDocs ensures that no value is present for ExternalDocs, not even an explicit nil
### GetInfo

`func (o *OAS) GetInfo() InfoType2`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *OAS) GetInfoOk() (*InfoType2, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *OAS) SetInfo(v InfoType2)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *OAS) HasInfo() bool`

HasInfo returns a boolean if a field has been set.

### SetInfoNil

`func (o *OAS) SetInfoNil(b bool)`

 SetInfoNil sets the value for Info to be an explicit nil

### UnsetInfo
`func (o *OAS) UnsetInfo()`

UnsetInfo ensures that no value is present for Info, not even an explicit nil
### GetOpenapi

`func (o *OAS) GetOpenapi() string`

GetOpenapi returns the Openapi field if non-nil, zero value otherwise.

### GetOpenapiOk

`func (o *OAS) GetOpenapiOk() (*string, bool)`

GetOpenapiOk returns a tuple with the Openapi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenapi

`func (o *OAS) SetOpenapi(v string)`

SetOpenapi sets Openapi field to given value.

### HasOpenapi

`func (o *OAS) HasOpenapi() bool`

HasOpenapi returns a boolean if a field has been set.

### GetPaths

`func (o *OAS) GetPaths() map[string]PathItem`

GetPaths returns the Paths field if non-nil, zero value otherwise.

### GetPathsOk

`func (o *OAS) GetPathsOk() (*map[string]PathItem, bool)`

GetPathsOk returns a tuple with the Paths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaths

`func (o *OAS) SetPaths(v map[string]PathItem)`

SetPaths sets Paths field to given value.

### HasPaths

`func (o *OAS) HasPaths() bool`

HasPaths returns a boolean if a field has been set.

### GetSecurity

`func (o *OAS) GetSecurity() []map[string][]string`

GetSecurity returns the Security field if non-nil, zero value otherwise.

### GetSecurityOk

`func (o *OAS) GetSecurityOk() (*[]map[string][]string, bool)`

GetSecurityOk returns a tuple with the Security field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurity

`func (o *OAS) SetSecurity(v []map[string][]string)`

SetSecurity sets Security field to given value.

### HasSecurity

`func (o *OAS) HasSecurity() bool`

HasSecurity returns a boolean if a field has been set.

### SetSecurityNil

`func (o *OAS) SetSecurityNil(b bool)`

 SetSecurityNil sets the value for Security to be an explicit nil

### UnsetSecurity
`func (o *OAS) UnsetSecurity()`

UnsetSecurity ensures that no value is present for Security, not even an explicit nil
### GetServers

`func (o *OAS) GetServers() []ServerType2`

GetServers returns the Servers field if non-nil, zero value otherwise.

### GetServersOk

`func (o *OAS) GetServersOk() (*[]ServerType2, bool)`

GetServersOk returns a tuple with the Servers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServers

`func (o *OAS) SetServers(v []ServerType2)`

SetServers sets Servers field to given value.

### HasServers

`func (o *OAS) HasServers() bool`

HasServers returns a boolean if a field has been set.

### SetServersNil

`func (o *OAS) SetServersNil(b bool)`

 SetServersNil sets the value for Servers to be an explicit nil

### UnsetServers
`func (o *OAS) UnsetServers()`

UnsetServers ensures that no value is present for Servers, not even an explicit nil
### GetTags

`func (o *OAS) GetTags() []Tag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *OAS) GetTagsOk() (*[]Tag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *OAS) SetTags(v []Tag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *OAS) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


