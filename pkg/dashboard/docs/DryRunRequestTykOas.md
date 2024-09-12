# DryRunRequestTykOas

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Openapi** | **string** |  | 
**Info** | [**Info**](Info.md) |  | 
**ExternalDocs** | Pointer to [**ExternalDocumentation**](ExternalDocumentation.md) |  | [optional] 
**Servers** | Pointer to [**[]Server**](Server.md) |  | [optional] 
**Security** | Pointer to [**[]map[string][]string**](map[string][]string.md) |  | [optional] 
**Tags** | Pointer to [**[]Tag**](Tag.md) |  | [optional] 
**Paths** | **map[string]interface{}** |  | 
**Components** | Pointer to [**Components**](Components.md) |  | [optional] 
**XTykApiGateway** | Pointer to [**OasXTykAPIGateway**](OasXTykAPIGateway.md) |  | [optional] 

## Methods

### NewDryRunRequestTykOas

`func NewDryRunRequestTykOas(openapi string, info Info, paths map[string]interface{}, ) *DryRunRequestTykOas`

NewDryRunRequestTykOas instantiates a new DryRunRequestTykOas object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDryRunRequestTykOasWithDefaults

`func NewDryRunRequestTykOasWithDefaults() *DryRunRequestTykOas`

NewDryRunRequestTykOasWithDefaults instantiates a new DryRunRequestTykOas object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOpenapi

`func (o *DryRunRequestTykOas) GetOpenapi() string`

GetOpenapi returns the Openapi field if non-nil, zero value otherwise.

### GetOpenapiOk

`func (o *DryRunRequestTykOas) GetOpenapiOk() (*string, bool)`

GetOpenapiOk returns a tuple with the Openapi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenapi

`func (o *DryRunRequestTykOas) SetOpenapi(v string)`

SetOpenapi sets Openapi field to given value.


### GetInfo

`func (o *DryRunRequestTykOas) GetInfo() Info`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *DryRunRequestTykOas) GetInfoOk() (*Info, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *DryRunRequestTykOas) SetInfo(v Info)`

SetInfo sets Info field to given value.


### GetExternalDocs

`func (o *DryRunRequestTykOas) GetExternalDocs() ExternalDocumentation`

GetExternalDocs returns the ExternalDocs field if non-nil, zero value otherwise.

### GetExternalDocsOk

`func (o *DryRunRequestTykOas) GetExternalDocsOk() (*ExternalDocumentation, bool)`

GetExternalDocsOk returns a tuple with the ExternalDocs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalDocs

`func (o *DryRunRequestTykOas) SetExternalDocs(v ExternalDocumentation)`

SetExternalDocs sets ExternalDocs field to given value.

### HasExternalDocs

`func (o *DryRunRequestTykOas) HasExternalDocs() bool`

HasExternalDocs returns a boolean if a field has been set.

### GetServers

`func (o *DryRunRequestTykOas) GetServers() []Server`

GetServers returns the Servers field if non-nil, zero value otherwise.

### GetServersOk

`func (o *DryRunRequestTykOas) GetServersOk() (*[]Server, bool)`

GetServersOk returns a tuple with the Servers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServers

`func (o *DryRunRequestTykOas) SetServers(v []Server)`

SetServers sets Servers field to given value.

### HasServers

`func (o *DryRunRequestTykOas) HasServers() bool`

HasServers returns a boolean if a field has been set.

### GetSecurity

`func (o *DryRunRequestTykOas) GetSecurity() []map[string][]string`

GetSecurity returns the Security field if non-nil, zero value otherwise.

### GetSecurityOk

`func (o *DryRunRequestTykOas) GetSecurityOk() (*[]map[string][]string, bool)`

GetSecurityOk returns a tuple with the Security field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurity

`func (o *DryRunRequestTykOas) SetSecurity(v []map[string][]string)`

SetSecurity sets Security field to given value.

### HasSecurity

`func (o *DryRunRequestTykOas) HasSecurity() bool`

HasSecurity returns a boolean if a field has been set.

### GetTags

`func (o *DryRunRequestTykOas) GetTags() []Tag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *DryRunRequestTykOas) GetTagsOk() (*[]Tag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *DryRunRequestTykOas) SetTags(v []Tag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *DryRunRequestTykOas) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetPaths

`func (o *DryRunRequestTykOas) GetPaths() map[string]interface{}`

GetPaths returns the Paths field if non-nil, zero value otherwise.

### GetPathsOk

`func (o *DryRunRequestTykOas) GetPathsOk() (*map[string]interface{}, bool)`

GetPathsOk returns a tuple with the Paths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaths

`func (o *DryRunRequestTykOas) SetPaths(v map[string]interface{})`

SetPaths sets Paths field to given value.


### GetComponents

`func (o *DryRunRequestTykOas) GetComponents() Components`

GetComponents returns the Components field if non-nil, zero value otherwise.

### GetComponentsOk

`func (o *DryRunRequestTykOas) GetComponentsOk() (*Components, bool)`

GetComponentsOk returns a tuple with the Components field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponents

`func (o *DryRunRequestTykOas) SetComponents(v Components)`

SetComponents sets Components field to given value.

### HasComponents

`func (o *DryRunRequestTykOas) HasComponents() bool`

HasComponents returns a boolean if a field has been set.

### GetXTykApiGateway

`func (o *DryRunRequestTykOas) GetXTykApiGateway() OasXTykAPIGateway`

GetXTykApiGateway returns the XTykApiGateway field if non-nil, zero value otherwise.

### GetXTykApiGatewayOk

`func (o *DryRunRequestTykOas) GetXTykApiGatewayOk() (*OasXTykAPIGateway, bool)`

GetXTykApiGatewayOk returns a tuple with the XTykApiGateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXTykApiGateway

`func (o *DryRunRequestTykOas) SetXTykApiGateway(v OasXTykAPIGateway)`

SetXTykApiGateway sets XTykApiGateway field to given value.

### HasXTykApiGateway

`func (o *DryRunRequestTykOas) HasXTykApiGateway() bool`

HasXTykApiGateway returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


