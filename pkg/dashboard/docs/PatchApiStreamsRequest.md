# PatchApiStreamsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Openapi** | **string** |  | 
**Info** | [**Info1**](Info1.md) |  | 
**ExternalDocs** | Pointer to [**ExternalDocumentation**](ExternalDocumentation.md) |  | [optional] 
**Servers** | Pointer to [**[]Server1**](Server1.md) |  | [optional] 
**Security** | Pointer to [**[]map[string][]string**](map[string][]string.md) |  | [optional] 
**Tags** | Pointer to [**[]Tag1**](Tag1.md) |  | [optional] 
**Paths** | **map[string]interface{}** |  | 
**Components** | Pointer to [**Components1**](Components1.md) |  | [optional] 
**XTykApiGateway** | Pointer to [**XTykAPIGateway**](XTykAPIGateway.md) |  | [optional] 
**XTykStreaming** | Pointer to [**XTykStreamingXTykStreaming**](XTykStreamingXTykStreaming.md) |  | [optional] 

## Methods

### NewPatchApiStreamsRequest

`func NewPatchApiStreamsRequest(openapi string, info Info1, paths map[string]interface{}, ) *PatchApiStreamsRequest`

NewPatchApiStreamsRequest instantiates a new PatchApiStreamsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchApiStreamsRequestWithDefaults

`func NewPatchApiStreamsRequestWithDefaults() *PatchApiStreamsRequest`

NewPatchApiStreamsRequestWithDefaults instantiates a new PatchApiStreamsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOpenapi

`func (o *PatchApiStreamsRequest) GetOpenapi() string`

GetOpenapi returns the Openapi field if non-nil, zero value otherwise.

### GetOpenapiOk

`func (o *PatchApiStreamsRequest) GetOpenapiOk() (*string, bool)`

GetOpenapiOk returns a tuple with the Openapi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenapi

`func (o *PatchApiStreamsRequest) SetOpenapi(v string)`

SetOpenapi sets Openapi field to given value.


### GetInfo

`func (o *PatchApiStreamsRequest) GetInfo() Info1`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *PatchApiStreamsRequest) GetInfoOk() (*Info1, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *PatchApiStreamsRequest) SetInfo(v Info1)`

SetInfo sets Info field to given value.


### GetExternalDocs

`func (o *PatchApiStreamsRequest) GetExternalDocs() ExternalDocumentation`

GetExternalDocs returns the ExternalDocs field if non-nil, zero value otherwise.

### GetExternalDocsOk

`func (o *PatchApiStreamsRequest) GetExternalDocsOk() (*ExternalDocumentation, bool)`

GetExternalDocsOk returns a tuple with the ExternalDocs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalDocs

`func (o *PatchApiStreamsRequest) SetExternalDocs(v ExternalDocumentation)`

SetExternalDocs sets ExternalDocs field to given value.

### HasExternalDocs

`func (o *PatchApiStreamsRequest) HasExternalDocs() bool`

HasExternalDocs returns a boolean if a field has been set.

### GetServers

`func (o *PatchApiStreamsRequest) GetServers() []Server1`

GetServers returns the Servers field if non-nil, zero value otherwise.

### GetServersOk

`func (o *PatchApiStreamsRequest) GetServersOk() (*[]Server1, bool)`

GetServersOk returns a tuple with the Servers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServers

`func (o *PatchApiStreamsRequest) SetServers(v []Server1)`

SetServers sets Servers field to given value.

### HasServers

`func (o *PatchApiStreamsRequest) HasServers() bool`

HasServers returns a boolean if a field has been set.

### GetSecurity

`func (o *PatchApiStreamsRequest) GetSecurity() []map[string][]string`

GetSecurity returns the Security field if non-nil, zero value otherwise.

### GetSecurityOk

`func (o *PatchApiStreamsRequest) GetSecurityOk() (*[]map[string][]string, bool)`

GetSecurityOk returns a tuple with the Security field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurity

`func (o *PatchApiStreamsRequest) SetSecurity(v []map[string][]string)`

SetSecurity sets Security field to given value.

### HasSecurity

`func (o *PatchApiStreamsRequest) HasSecurity() bool`

HasSecurity returns a boolean if a field has been set.

### GetTags

`func (o *PatchApiStreamsRequest) GetTags() []Tag1`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PatchApiStreamsRequest) GetTagsOk() (*[]Tag1, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PatchApiStreamsRequest) SetTags(v []Tag1)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PatchApiStreamsRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetPaths

`func (o *PatchApiStreamsRequest) GetPaths() map[string]interface{}`

GetPaths returns the Paths field if non-nil, zero value otherwise.

### GetPathsOk

`func (o *PatchApiStreamsRequest) GetPathsOk() (*map[string]interface{}, bool)`

GetPathsOk returns a tuple with the Paths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaths

`func (o *PatchApiStreamsRequest) SetPaths(v map[string]interface{})`

SetPaths sets Paths field to given value.


### GetComponents

`func (o *PatchApiStreamsRequest) GetComponents() Components1`

GetComponents returns the Components field if non-nil, zero value otherwise.

### GetComponentsOk

`func (o *PatchApiStreamsRequest) GetComponentsOk() (*Components1, bool)`

GetComponentsOk returns a tuple with the Components field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponents

`func (o *PatchApiStreamsRequest) SetComponents(v Components1)`

SetComponents sets Components field to given value.

### HasComponents

`func (o *PatchApiStreamsRequest) HasComponents() bool`

HasComponents returns a boolean if a field has been set.

### GetXTykApiGateway

`func (o *PatchApiStreamsRequest) GetXTykApiGateway() XTykAPIGateway`

GetXTykApiGateway returns the XTykApiGateway field if non-nil, zero value otherwise.

### GetXTykApiGatewayOk

`func (o *PatchApiStreamsRequest) GetXTykApiGatewayOk() (*XTykAPIGateway, bool)`

GetXTykApiGatewayOk returns a tuple with the XTykApiGateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXTykApiGateway

`func (o *PatchApiStreamsRequest) SetXTykApiGateway(v XTykAPIGateway)`

SetXTykApiGateway sets XTykApiGateway field to given value.

### HasXTykApiGateway

`func (o *PatchApiStreamsRequest) HasXTykApiGateway() bool`

HasXTykApiGateway returns a boolean if a field has been set.

### GetXTykStreaming

`func (o *PatchApiStreamsRequest) GetXTykStreaming() XTykStreamingXTykStreaming`

GetXTykStreaming returns the XTykStreaming field if non-nil, zero value otherwise.

### GetXTykStreamingOk

`func (o *PatchApiStreamsRequest) GetXTykStreamingOk() (*XTykStreamingXTykStreaming, bool)`

GetXTykStreamingOk returns a tuple with the XTykStreaming field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXTykStreaming

`func (o *PatchApiStreamsRequest) SetXTykStreaming(v XTykStreamingXTykStreaming)`

SetXTykStreaming sets XTykStreaming field to given value.

### HasXTykStreaming

`func (o *PatchApiStreamsRequest) HasXTykStreaming() bool`

HasXTykStreaming returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


