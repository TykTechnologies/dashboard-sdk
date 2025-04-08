# CreateStreamsApiRequest

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

### NewCreateStreamsApiRequest

`func NewCreateStreamsApiRequest(openapi string, info Info1, paths map[string]interface{}, ) *CreateStreamsApiRequest`

NewCreateStreamsApiRequest instantiates a new CreateStreamsApiRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateStreamsApiRequestWithDefaults

`func NewCreateStreamsApiRequestWithDefaults() *CreateStreamsApiRequest`

NewCreateStreamsApiRequestWithDefaults instantiates a new CreateStreamsApiRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOpenapi

`func (o *CreateStreamsApiRequest) GetOpenapi() string`

GetOpenapi returns the Openapi field if non-nil, zero value otherwise.

### GetOpenapiOk

`func (o *CreateStreamsApiRequest) GetOpenapiOk() (*string, bool)`

GetOpenapiOk returns a tuple with the Openapi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenapi

`func (o *CreateStreamsApiRequest) SetOpenapi(v string)`

SetOpenapi sets Openapi field to given value.


### GetInfo

`func (o *CreateStreamsApiRequest) GetInfo() Info1`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *CreateStreamsApiRequest) GetInfoOk() (*Info1, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *CreateStreamsApiRequest) SetInfo(v Info1)`

SetInfo sets Info field to given value.


### GetExternalDocs

`func (o *CreateStreamsApiRequest) GetExternalDocs() ExternalDocumentation`

GetExternalDocs returns the ExternalDocs field if non-nil, zero value otherwise.

### GetExternalDocsOk

`func (o *CreateStreamsApiRequest) GetExternalDocsOk() (*ExternalDocumentation, bool)`

GetExternalDocsOk returns a tuple with the ExternalDocs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalDocs

`func (o *CreateStreamsApiRequest) SetExternalDocs(v ExternalDocumentation)`

SetExternalDocs sets ExternalDocs field to given value.

### HasExternalDocs

`func (o *CreateStreamsApiRequest) HasExternalDocs() bool`

HasExternalDocs returns a boolean if a field has been set.

### GetServers

`func (o *CreateStreamsApiRequest) GetServers() []Server1`

GetServers returns the Servers field if non-nil, zero value otherwise.

### GetServersOk

`func (o *CreateStreamsApiRequest) GetServersOk() (*[]Server1, bool)`

GetServersOk returns a tuple with the Servers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServers

`func (o *CreateStreamsApiRequest) SetServers(v []Server1)`

SetServers sets Servers field to given value.

### HasServers

`func (o *CreateStreamsApiRequest) HasServers() bool`

HasServers returns a boolean if a field has been set.

### GetSecurity

`func (o *CreateStreamsApiRequest) GetSecurity() []map[string][]string`

GetSecurity returns the Security field if non-nil, zero value otherwise.

### GetSecurityOk

`func (o *CreateStreamsApiRequest) GetSecurityOk() (*[]map[string][]string, bool)`

GetSecurityOk returns a tuple with the Security field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurity

`func (o *CreateStreamsApiRequest) SetSecurity(v []map[string][]string)`

SetSecurity sets Security field to given value.

### HasSecurity

`func (o *CreateStreamsApiRequest) HasSecurity() bool`

HasSecurity returns a boolean if a field has been set.

### GetTags

`func (o *CreateStreamsApiRequest) GetTags() []Tag1`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CreateStreamsApiRequest) GetTagsOk() (*[]Tag1, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CreateStreamsApiRequest) SetTags(v []Tag1)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CreateStreamsApiRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetPaths

`func (o *CreateStreamsApiRequest) GetPaths() map[string]interface{}`

GetPaths returns the Paths field if non-nil, zero value otherwise.

### GetPathsOk

`func (o *CreateStreamsApiRequest) GetPathsOk() (*map[string]interface{}, bool)`

GetPathsOk returns a tuple with the Paths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaths

`func (o *CreateStreamsApiRequest) SetPaths(v map[string]interface{})`

SetPaths sets Paths field to given value.


### GetComponents

`func (o *CreateStreamsApiRequest) GetComponents() Components1`

GetComponents returns the Components field if non-nil, zero value otherwise.

### GetComponentsOk

`func (o *CreateStreamsApiRequest) GetComponentsOk() (*Components1, bool)`

GetComponentsOk returns a tuple with the Components field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponents

`func (o *CreateStreamsApiRequest) SetComponents(v Components1)`

SetComponents sets Components field to given value.

### HasComponents

`func (o *CreateStreamsApiRequest) HasComponents() bool`

HasComponents returns a boolean if a field has been set.

### GetXTykApiGateway

`func (o *CreateStreamsApiRequest) GetXTykApiGateway() XTykAPIGateway`

GetXTykApiGateway returns the XTykApiGateway field if non-nil, zero value otherwise.

### GetXTykApiGatewayOk

`func (o *CreateStreamsApiRequest) GetXTykApiGatewayOk() (*XTykAPIGateway, bool)`

GetXTykApiGatewayOk returns a tuple with the XTykApiGateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXTykApiGateway

`func (o *CreateStreamsApiRequest) SetXTykApiGateway(v XTykAPIGateway)`

SetXTykApiGateway sets XTykApiGateway field to given value.

### HasXTykApiGateway

`func (o *CreateStreamsApiRequest) HasXTykApiGateway() bool`

HasXTykApiGateway returns a boolean if a field has been set.

### GetXTykStreaming

`func (o *CreateStreamsApiRequest) GetXTykStreaming() XTykStreamingXTykStreaming`

GetXTykStreaming returns the XTykStreaming field if non-nil, zero value otherwise.

### GetXTykStreamingOk

`func (o *CreateStreamsApiRequest) GetXTykStreamingOk() (*XTykStreamingXTykStreaming, bool)`

GetXTykStreamingOk returns a tuple with the XTykStreaming field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXTykStreaming

`func (o *CreateStreamsApiRequest) SetXTykStreaming(v XTykStreamingXTykStreaming)`

SetXTykStreaming sets XTykStreaming field to given value.

### HasXTykStreaming

`func (o *CreateStreamsApiRequest) HasXTykStreaming() bool`

HasXTykStreaming returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


