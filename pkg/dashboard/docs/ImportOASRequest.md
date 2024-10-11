# ImportOASRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | Pointer to **string** |  | [optional] 
**Openapi** | **string** |  | 
**Info** | [**Info1**](Info1.md) |  | 
**ExternalDocs** | Pointer to [**ExternalDocumentation**](ExternalDocumentation.md) |  | [optional] 
**Servers** | Pointer to [**[]Server1**](Server1.md) |  | [optional] 
**Security** | Pointer to [**[]map[string][]string**](map[string][]string.md) |  | [optional] 
**Tags** | Pointer to [**[]Tag1**](Tag1.md) |  | [optional] 
**Paths** | **map[string]interface{}** |  | 
**Components** | Pointer to [**Components1**](Components1.md) |  | [optional] 

## Methods

### NewImportOASRequest

`func NewImportOASRequest(openapi string, info Info1, paths map[string]interface{}, ) *ImportOASRequest`

NewImportOASRequest instantiates a new ImportOASRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImportOASRequestWithDefaults

`func NewImportOASRequestWithDefaults() *ImportOASRequest`

NewImportOASRequestWithDefaults instantiates a new ImportOASRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *ImportOASRequest) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ImportOASRequest) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ImportOASRequest) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *ImportOASRequest) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetOpenapi

`func (o *ImportOASRequest) GetOpenapi() string`

GetOpenapi returns the Openapi field if non-nil, zero value otherwise.

### GetOpenapiOk

`func (o *ImportOASRequest) GetOpenapiOk() (*string, bool)`

GetOpenapiOk returns a tuple with the Openapi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenapi

`func (o *ImportOASRequest) SetOpenapi(v string)`

SetOpenapi sets Openapi field to given value.


### GetInfo

`func (o *ImportOASRequest) GetInfo() Info1`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *ImportOASRequest) GetInfoOk() (*Info1, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *ImportOASRequest) SetInfo(v Info1)`

SetInfo sets Info field to given value.


### GetExternalDocs

`func (o *ImportOASRequest) GetExternalDocs() ExternalDocumentation`

GetExternalDocs returns the ExternalDocs field if non-nil, zero value otherwise.

### GetExternalDocsOk

`func (o *ImportOASRequest) GetExternalDocsOk() (*ExternalDocumentation, bool)`

GetExternalDocsOk returns a tuple with the ExternalDocs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalDocs

`func (o *ImportOASRequest) SetExternalDocs(v ExternalDocumentation)`

SetExternalDocs sets ExternalDocs field to given value.

### HasExternalDocs

`func (o *ImportOASRequest) HasExternalDocs() bool`

HasExternalDocs returns a boolean if a field has been set.

### GetServers

`func (o *ImportOASRequest) GetServers() []Server1`

GetServers returns the Servers field if non-nil, zero value otherwise.

### GetServersOk

`func (o *ImportOASRequest) GetServersOk() (*[]Server1, bool)`

GetServersOk returns a tuple with the Servers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServers

`func (o *ImportOASRequest) SetServers(v []Server1)`

SetServers sets Servers field to given value.

### HasServers

`func (o *ImportOASRequest) HasServers() bool`

HasServers returns a boolean if a field has been set.

### GetSecurity

`func (o *ImportOASRequest) GetSecurity() []map[string][]string`

GetSecurity returns the Security field if non-nil, zero value otherwise.

### GetSecurityOk

`func (o *ImportOASRequest) GetSecurityOk() (*[]map[string][]string, bool)`

GetSecurityOk returns a tuple with the Security field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurity

`func (o *ImportOASRequest) SetSecurity(v []map[string][]string)`

SetSecurity sets Security field to given value.

### HasSecurity

`func (o *ImportOASRequest) HasSecurity() bool`

HasSecurity returns a boolean if a field has been set.

### GetTags

`func (o *ImportOASRequest) GetTags() []Tag1`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ImportOASRequest) GetTagsOk() (*[]Tag1, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ImportOASRequest) SetTags(v []Tag1)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ImportOASRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetPaths

`func (o *ImportOASRequest) GetPaths() map[string]interface{}`

GetPaths returns the Paths field if non-nil, zero value otherwise.

### GetPathsOk

`func (o *ImportOASRequest) GetPathsOk() (*map[string]interface{}, bool)`

GetPathsOk returns a tuple with the Paths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaths

`func (o *ImportOASRequest) SetPaths(v map[string]interface{})`

SetPaths sets Paths field to given value.


### GetComponents

`func (o *ImportOASRequest) GetComponents() Components1`

GetComponents returns the Components field if non-nil, zero value otherwise.

### GetComponentsOk

`func (o *ImportOASRequest) GetComponentsOk() (*Components1, bool)`

GetComponentsOk returns a tuple with the Components field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponents

`func (o *ImportOASRequest) SetComponents(v Components1)`

SetComponents sets Components field to given value.

### HasComponents

`func (o *ImportOASRequest) HasComponents() bool`

HasComponents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


