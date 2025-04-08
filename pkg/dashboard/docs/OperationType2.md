# OperationType2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Callbacks** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**Deprecated** | Pointer to **bool** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**ExternalDocs** | Pointer to [**NullableExternalDocs**](ExternalDocs.md) |  | [optional] 
**OperationId** | Pointer to **string** |  | [optional] 
**Parameters** | Pointer to **[]map[string]interface{}** |  | [optional] 
**RequestBody** | Pointer to **map[string]interface{}** |  | [optional] 
**Responses** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**Security** | Pointer to [**[]map[string][]string**](map[string][]string.md) |  | [optional] 
**Servers** | Pointer to [**[]ServerType2**](ServerType2.md) |  | [optional] 
**Summary** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 

## Methods

### NewOperationType2

`func NewOperationType2() *OperationType2`

NewOperationType2 instantiates a new OperationType2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOperationType2WithDefaults

`func NewOperationType2WithDefaults() *OperationType2`

NewOperationType2WithDefaults instantiates a new OperationType2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCallbacks

`func (o *OperationType2) GetCallbacks() map[string]map[string]interface{}`

GetCallbacks returns the Callbacks field if non-nil, zero value otherwise.

### GetCallbacksOk

`func (o *OperationType2) GetCallbacksOk() (*map[string]map[string]interface{}, bool)`

GetCallbacksOk returns a tuple with the Callbacks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallbacks

`func (o *OperationType2) SetCallbacks(v map[string]map[string]interface{})`

SetCallbacks sets Callbacks field to given value.

### HasCallbacks

`func (o *OperationType2) HasCallbacks() bool`

HasCallbacks returns a boolean if a field has been set.

### GetDeprecated

`func (o *OperationType2) GetDeprecated() bool`

GetDeprecated returns the Deprecated field if non-nil, zero value otherwise.

### GetDeprecatedOk

`func (o *OperationType2) GetDeprecatedOk() (*bool, bool)`

GetDeprecatedOk returns a tuple with the Deprecated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeprecated

`func (o *OperationType2) SetDeprecated(v bool)`

SetDeprecated sets Deprecated field to given value.

### HasDeprecated

`func (o *OperationType2) HasDeprecated() bool`

HasDeprecated returns a boolean if a field has been set.

### GetDescription

`func (o *OperationType2) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *OperationType2) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *OperationType2) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *OperationType2) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExternalDocs

`func (o *OperationType2) GetExternalDocs() ExternalDocs`

GetExternalDocs returns the ExternalDocs field if non-nil, zero value otherwise.

### GetExternalDocsOk

`func (o *OperationType2) GetExternalDocsOk() (*ExternalDocs, bool)`

GetExternalDocsOk returns a tuple with the ExternalDocs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalDocs

`func (o *OperationType2) SetExternalDocs(v ExternalDocs)`

SetExternalDocs sets ExternalDocs field to given value.

### HasExternalDocs

`func (o *OperationType2) HasExternalDocs() bool`

HasExternalDocs returns a boolean if a field has been set.

### SetExternalDocsNil

`func (o *OperationType2) SetExternalDocsNil(b bool)`

 SetExternalDocsNil sets the value for ExternalDocs to be an explicit nil

### UnsetExternalDocs
`func (o *OperationType2) UnsetExternalDocs()`

UnsetExternalDocs ensures that no value is present for ExternalDocs, not even an explicit nil
### GetOperationId

`func (o *OperationType2) GetOperationId() string`

GetOperationId returns the OperationId field if non-nil, zero value otherwise.

### GetOperationIdOk

`func (o *OperationType2) GetOperationIdOk() (*string, bool)`

GetOperationIdOk returns a tuple with the OperationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationId

`func (o *OperationType2) SetOperationId(v string)`

SetOperationId sets OperationId field to given value.

### HasOperationId

`func (o *OperationType2) HasOperationId() bool`

HasOperationId returns a boolean if a field has been set.

### GetParameters

`func (o *OperationType2) GetParameters() []map[string]interface{}`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *OperationType2) GetParametersOk() (*[]map[string]interface{}, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *OperationType2) SetParameters(v []map[string]interface{})`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *OperationType2) HasParameters() bool`

HasParameters returns a boolean if a field has been set.

### GetRequestBody

`func (o *OperationType2) GetRequestBody() map[string]interface{}`

GetRequestBody returns the RequestBody field if non-nil, zero value otherwise.

### GetRequestBodyOk

`func (o *OperationType2) GetRequestBodyOk() (*map[string]interface{}, bool)`

GetRequestBodyOk returns a tuple with the RequestBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestBody

`func (o *OperationType2) SetRequestBody(v map[string]interface{})`

SetRequestBody sets RequestBody field to given value.

### HasRequestBody

`func (o *OperationType2) HasRequestBody() bool`

HasRequestBody returns a boolean if a field has been set.

### GetResponses

`func (o *OperationType2) GetResponses() map[string]map[string]interface{}`

GetResponses returns the Responses field if non-nil, zero value otherwise.

### GetResponsesOk

`func (o *OperationType2) GetResponsesOk() (*map[string]map[string]interface{}, bool)`

GetResponsesOk returns a tuple with the Responses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponses

`func (o *OperationType2) SetResponses(v map[string]map[string]interface{})`

SetResponses sets Responses field to given value.

### HasResponses

`func (o *OperationType2) HasResponses() bool`

HasResponses returns a boolean if a field has been set.

### GetSecurity

`func (o *OperationType2) GetSecurity() []map[string][]string`

GetSecurity returns the Security field if non-nil, zero value otherwise.

### GetSecurityOk

`func (o *OperationType2) GetSecurityOk() (*[]map[string][]string, bool)`

GetSecurityOk returns a tuple with the Security field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurity

`func (o *OperationType2) SetSecurity(v []map[string][]string)`

SetSecurity sets Security field to given value.

### HasSecurity

`func (o *OperationType2) HasSecurity() bool`

HasSecurity returns a boolean if a field has been set.

### SetSecurityNil

`func (o *OperationType2) SetSecurityNil(b bool)`

 SetSecurityNil sets the value for Security to be an explicit nil

### UnsetSecurity
`func (o *OperationType2) UnsetSecurity()`

UnsetSecurity ensures that no value is present for Security, not even an explicit nil
### GetServers

`func (o *OperationType2) GetServers() []ServerType2`

GetServers returns the Servers field if non-nil, zero value otherwise.

### GetServersOk

`func (o *OperationType2) GetServersOk() (*[]ServerType2, bool)`

GetServersOk returns a tuple with the Servers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServers

`func (o *OperationType2) SetServers(v []ServerType2)`

SetServers sets Servers field to given value.

### HasServers

`func (o *OperationType2) HasServers() bool`

HasServers returns a boolean if a field has been set.

### SetServersNil

`func (o *OperationType2) SetServersNil(b bool)`

 SetServersNil sets the value for Servers to be an explicit nil

### UnsetServers
`func (o *OperationType2) UnsetServers()`

UnsetServers ensures that no value is present for Servers, not even an explicit nil
### GetSummary

`func (o *OperationType2) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *OperationType2) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *OperationType2) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *OperationType2) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetTags

`func (o *OperationType2) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *OperationType2) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *OperationType2) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *OperationType2) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


