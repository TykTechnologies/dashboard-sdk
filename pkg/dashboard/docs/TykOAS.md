# TykOAS

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
**XTykApiGateway** | Pointer to [**OasXTykAPIGateway**](OasXTykAPIGateway.md) |  | [optional] 

## Methods

### NewTykOAS

`func NewTykOAS() *TykOAS`

NewTykOAS instantiates a new TykOAS object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTykOASWithDefaults

`func NewTykOASWithDefaults() *TykOAS`

NewTykOASWithDefaults instantiates a new TykOAS object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComponents

`func (o *TykOAS) GetComponents() Openapi3Components`

GetComponents returns the Components field if non-nil, zero value otherwise.

### GetComponentsOk

`func (o *TykOAS) GetComponentsOk() (*Openapi3Components, bool)`

GetComponentsOk returns a tuple with the Components field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponents

`func (o *TykOAS) SetComponents(v Openapi3Components)`

SetComponents sets Components field to given value.

### HasComponents

`func (o *TykOAS) HasComponents() bool`

HasComponents returns a boolean if a field has been set.

### GetExternalDocs

`func (o *TykOAS) GetExternalDocs() Openapi3ExternalDocs`

GetExternalDocs returns the ExternalDocs field if non-nil, zero value otherwise.

### GetExternalDocsOk

`func (o *TykOAS) GetExternalDocsOk() (*Openapi3ExternalDocs, bool)`

GetExternalDocsOk returns a tuple with the ExternalDocs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalDocs

`func (o *TykOAS) SetExternalDocs(v Openapi3ExternalDocs)`

SetExternalDocs sets ExternalDocs field to given value.

### HasExternalDocs

`func (o *TykOAS) HasExternalDocs() bool`

HasExternalDocs returns a boolean if a field has been set.

### GetInfo

`func (o *TykOAS) GetInfo() Openapi3Info`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *TykOAS) GetInfoOk() (*Openapi3Info, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *TykOAS) SetInfo(v Openapi3Info)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *TykOAS) HasInfo() bool`

HasInfo returns a boolean if a field has been set.

### GetOpenapi

`func (o *TykOAS) GetOpenapi() string`

GetOpenapi returns the Openapi field if non-nil, zero value otherwise.

### GetOpenapiOk

`func (o *TykOAS) GetOpenapiOk() (*string, bool)`

GetOpenapiOk returns a tuple with the Openapi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenapi

`func (o *TykOAS) SetOpenapi(v string)`

SetOpenapi sets Openapi field to given value.

### HasOpenapi

`func (o *TykOAS) HasOpenapi() bool`

HasOpenapi returns a boolean if a field has been set.

### GetPaths

`func (o *TykOAS) GetPaths() map[string]Openapi3PathItem`

GetPaths returns the Paths field if non-nil, zero value otherwise.

### GetPathsOk

`func (o *TykOAS) GetPathsOk() (*map[string]Openapi3PathItem, bool)`

GetPathsOk returns a tuple with the Paths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaths

`func (o *TykOAS) SetPaths(v map[string]Openapi3PathItem)`

SetPaths sets Paths field to given value.

### HasPaths

`func (o *TykOAS) HasPaths() bool`

HasPaths returns a boolean if a field has been set.

### GetSecurity

`func (o *TykOAS) GetSecurity() []map[string][]string`

GetSecurity returns the Security field if non-nil, zero value otherwise.

### GetSecurityOk

`func (o *TykOAS) GetSecurityOk() (*[]map[string][]string, bool)`

GetSecurityOk returns a tuple with the Security field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurity

`func (o *TykOAS) SetSecurity(v []map[string][]string)`

SetSecurity sets Security field to given value.

### HasSecurity

`func (o *TykOAS) HasSecurity() bool`

HasSecurity returns a boolean if a field has been set.

### SetSecurityNil

`func (o *TykOAS) SetSecurityNil(b bool)`

 SetSecurityNil sets the value for Security to be an explicit nil

### UnsetSecurity
`func (o *TykOAS) UnsetSecurity()`

UnsetSecurity ensures that no value is present for Security, not even an explicit nil
### GetServers

`func (o *TykOAS) GetServers() []Openapi3Server`

GetServers returns the Servers field if non-nil, zero value otherwise.

### GetServersOk

`func (o *TykOAS) GetServersOk() (*[]Openapi3Server, bool)`

GetServersOk returns a tuple with the Servers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServers

`func (o *TykOAS) SetServers(v []Openapi3Server)`

SetServers sets Servers field to given value.

### HasServers

`func (o *TykOAS) HasServers() bool`

HasServers returns a boolean if a field has been set.

### SetServersNil

`func (o *TykOAS) SetServersNil(b bool)`

 SetServersNil sets the value for Servers to be an explicit nil

### UnsetServers
`func (o *TykOAS) UnsetServers()`

UnsetServers ensures that no value is present for Servers, not even an explicit nil
### GetTags

`func (o *TykOAS) GetTags() []Openapi3Tag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *TykOAS) GetTagsOk() (*[]Openapi3Tag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *TykOAS) SetTags(v []Openapi3Tag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *TykOAS) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetXTykApiGateway

`func (o *TykOAS) GetXTykApiGateway() OasXTykAPIGateway`

GetXTykApiGateway returns the XTykApiGateway field if non-nil, zero value otherwise.

### GetXTykApiGatewayOk

`func (o *TykOAS) GetXTykApiGatewayOk() (*OasXTykAPIGateway, bool)`

GetXTykApiGatewayOk returns a tuple with the XTykApiGateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXTykApiGateway

`func (o *TykOAS) SetXTykApiGateway(v OasXTykAPIGateway)`

SetXTykApiGateway sets XTykApiGateway field to given value.

### HasXTykApiGateway

`func (o *TykOAS) HasXTykApiGateway() bool`

HasXTykApiGateway returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


