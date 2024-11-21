# DryRunApiOAS200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**XTykApiGateway** | Pointer to [**XTykAPIGateway**](XTykAPIGateway.md) |  | [optional] 
**Openapi** | **string** |  | 
**Info** | [**Info1**](Info1.md) |  | 
**ExternalDocs** | Pointer to [**ExternalDocumentation**](ExternalDocumentation.md) |  | [optional] 
**Servers** | Pointer to [**[]Server1**](Server1.md) |  | [optional] 
**Security** | Pointer to [**[]map[string][]string**](map[string][]string.md) |  | [optional] 
**Tags** | Pointer to [**[]Tag1**](Tag1.md) |  | [optional] 
**Paths** | **map[string]interface{}** |  | 
**Components** | Pointer to [**Components1**](Components1.md) |  | [optional] 

## Methods

### NewDryRunApiOAS200Response

`func NewDryRunApiOAS200Response(openapi string, info Info1, paths map[string]interface{}, ) *DryRunApiOAS200Response`

NewDryRunApiOAS200Response instantiates a new DryRunApiOAS200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDryRunApiOAS200ResponseWithDefaults

`func NewDryRunApiOAS200ResponseWithDefaults() *DryRunApiOAS200Response`

NewDryRunApiOAS200ResponseWithDefaults instantiates a new DryRunApiOAS200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetXTykApiGateway

`func (o *DryRunApiOAS200Response) GetXTykApiGateway() XTykAPIGateway`

GetXTykApiGateway returns the XTykApiGateway field if non-nil, zero value otherwise.

### GetXTykApiGatewayOk

`func (o *DryRunApiOAS200Response) GetXTykApiGatewayOk() (*XTykAPIGateway, bool)`

GetXTykApiGatewayOk returns a tuple with the XTykApiGateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXTykApiGateway

`func (o *DryRunApiOAS200Response) SetXTykApiGateway(v XTykAPIGateway)`

SetXTykApiGateway sets XTykApiGateway field to given value.

### HasXTykApiGateway

`func (o *DryRunApiOAS200Response) HasXTykApiGateway() bool`

HasXTykApiGateway returns a boolean if a field has been set.

### GetOpenapi

`func (o *DryRunApiOAS200Response) GetOpenapi() string`

GetOpenapi returns the Openapi field if non-nil, zero value otherwise.

### GetOpenapiOk

`func (o *DryRunApiOAS200Response) GetOpenapiOk() (*string, bool)`

GetOpenapiOk returns a tuple with the Openapi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenapi

`func (o *DryRunApiOAS200Response) SetOpenapi(v string)`

SetOpenapi sets Openapi field to given value.


### GetInfo

`func (o *DryRunApiOAS200Response) GetInfo() Info1`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *DryRunApiOAS200Response) GetInfoOk() (*Info1, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *DryRunApiOAS200Response) SetInfo(v Info1)`

SetInfo sets Info field to given value.


### GetExternalDocs

`func (o *DryRunApiOAS200Response) GetExternalDocs() ExternalDocumentation`

GetExternalDocs returns the ExternalDocs field if non-nil, zero value otherwise.

### GetExternalDocsOk

`func (o *DryRunApiOAS200Response) GetExternalDocsOk() (*ExternalDocumentation, bool)`

GetExternalDocsOk returns a tuple with the ExternalDocs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalDocs

`func (o *DryRunApiOAS200Response) SetExternalDocs(v ExternalDocumentation)`

SetExternalDocs sets ExternalDocs field to given value.

### HasExternalDocs

`func (o *DryRunApiOAS200Response) HasExternalDocs() bool`

HasExternalDocs returns a boolean if a field has been set.

### GetServers

`func (o *DryRunApiOAS200Response) GetServers() []Server1`

GetServers returns the Servers field if non-nil, zero value otherwise.

### GetServersOk

`func (o *DryRunApiOAS200Response) GetServersOk() (*[]Server1, bool)`

GetServersOk returns a tuple with the Servers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServers

`func (o *DryRunApiOAS200Response) SetServers(v []Server1)`

SetServers sets Servers field to given value.

### HasServers

`func (o *DryRunApiOAS200Response) HasServers() bool`

HasServers returns a boolean if a field has been set.

### GetSecurity

`func (o *DryRunApiOAS200Response) GetSecurity() []map[string][]string`

GetSecurity returns the Security field if non-nil, zero value otherwise.

### GetSecurityOk

`func (o *DryRunApiOAS200Response) GetSecurityOk() (*[]map[string][]string, bool)`

GetSecurityOk returns a tuple with the Security field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurity

`func (o *DryRunApiOAS200Response) SetSecurity(v []map[string][]string)`

SetSecurity sets Security field to given value.

### HasSecurity

`func (o *DryRunApiOAS200Response) HasSecurity() bool`

HasSecurity returns a boolean if a field has been set.

### GetTags

`func (o *DryRunApiOAS200Response) GetTags() []Tag1`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *DryRunApiOAS200Response) GetTagsOk() (*[]Tag1, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *DryRunApiOAS200Response) SetTags(v []Tag1)`

SetTags sets Tags field to given value.

### HasTags

`func (o *DryRunApiOAS200Response) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetPaths

`func (o *DryRunApiOAS200Response) GetPaths() map[string]interface{}`

GetPaths returns the Paths field if non-nil, zero value otherwise.

### GetPathsOk

`func (o *DryRunApiOAS200Response) GetPathsOk() (*map[string]interface{}, bool)`

GetPathsOk returns a tuple with the Paths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaths

`func (o *DryRunApiOAS200Response) SetPaths(v map[string]interface{})`

SetPaths sets Paths field to given value.


### GetComponents

`func (o *DryRunApiOAS200Response) GetComponents() Components1`

GetComponents returns the Components field if non-nil, zero value otherwise.

### GetComponentsOk

`func (o *DryRunApiOAS200Response) GetComponentsOk() (*Components1, bool)`

GetComponentsOk returns a tuple with the Components field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponents

`func (o *DryRunApiOAS200Response) SetComponents(v Components1)`

SetComponents sets Components field to given value.

### HasComponents

`func (o *DryRunApiOAS200Response) HasComponents() bool`

HasComponents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


