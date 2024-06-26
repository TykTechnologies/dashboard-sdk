# Openapi3Operation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Callbacks** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**Deprecated** | Pointer to **bool** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**ExternalDocs** | Pointer to [**Openapi3ExternalDocs**](Openapi3ExternalDocs.md) |  | [optional] 
**OperationId** | Pointer to **string** |  | [optional] 
**Parameters** | Pointer to **[]map[string]interface{}** |  | [optional] 
**RequestBody** | Pointer to **map[string]interface{}** |  | [optional] 
**Responses** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**Security** | Pointer to [**[]map[string][]string**](map[string][]string.md) |  | [optional] 
**Servers** | Pointer to [**[]Openapi3Server**](Openapi3Server.md) |  | [optional] 
**Summary** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 

## Methods

### NewOpenapi3Operation

`func NewOpenapi3Operation() *Openapi3Operation`

NewOpenapi3Operation instantiates a new Openapi3Operation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpenapi3OperationWithDefaults

`func NewOpenapi3OperationWithDefaults() *Openapi3Operation`

NewOpenapi3OperationWithDefaults instantiates a new Openapi3Operation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCallbacks

`func (o *Openapi3Operation) GetCallbacks() map[string]map[string]interface{}`

GetCallbacks returns the Callbacks field if non-nil, zero value otherwise.

### GetCallbacksOk

`func (o *Openapi3Operation) GetCallbacksOk() (*map[string]map[string]interface{}, bool)`

GetCallbacksOk returns a tuple with the Callbacks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallbacks

`func (o *Openapi3Operation) SetCallbacks(v map[string]map[string]interface{})`

SetCallbacks sets Callbacks field to given value.

### HasCallbacks

`func (o *Openapi3Operation) HasCallbacks() bool`

HasCallbacks returns a boolean if a field has been set.

### GetDeprecated

`func (o *Openapi3Operation) GetDeprecated() bool`

GetDeprecated returns the Deprecated field if non-nil, zero value otherwise.

### GetDeprecatedOk

`func (o *Openapi3Operation) GetDeprecatedOk() (*bool, bool)`

GetDeprecatedOk returns a tuple with the Deprecated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeprecated

`func (o *Openapi3Operation) SetDeprecated(v bool)`

SetDeprecated sets Deprecated field to given value.

### HasDeprecated

`func (o *Openapi3Operation) HasDeprecated() bool`

HasDeprecated returns a boolean if a field has been set.

### GetDescription

`func (o *Openapi3Operation) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Openapi3Operation) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Openapi3Operation) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Openapi3Operation) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExternalDocs

`func (o *Openapi3Operation) GetExternalDocs() Openapi3ExternalDocs`

GetExternalDocs returns the ExternalDocs field if non-nil, zero value otherwise.

### GetExternalDocsOk

`func (o *Openapi3Operation) GetExternalDocsOk() (*Openapi3ExternalDocs, bool)`

GetExternalDocsOk returns a tuple with the ExternalDocs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalDocs

`func (o *Openapi3Operation) SetExternalDocs(v Openapi3ExternalDocs)`

SetExternalDocs sets ExternalDocs field to given value.

### HasExternalDocs

`func (o *Openapi3Operation) HasExternalDocs() bool`

HasExternalDocs returns a boolean if a field has been set.

### GetOperationId

`func (o *Openapi3Operation) GetOperationId() string`

GetOperationId returns the OperationId field if non-nil, zero value otherwise.

### GetOperationIdOk

`func (o *Openapi3Operation) GetOperationIdOk() (*string, bool)`

GetOperationIdOk returns a tuple with the OperationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationId

`func (o *Openapi3Operation) SetOperationId(v string)`

SetOperationId sets OperationId field to given value.

### HasOperationId

`func (o *Openapi3Operation) HasOperationId() bool`

HasOperationId returns a boolean if a field has been set.

### GetParameters

`func (o *Openapi3Operation) GetParameters() []map[string]interface{}`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *Openapi3Operation) GetParametersOk() (*[]map[string]interface{}, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *Openapi3Operation) SetParameters(v []map[string]interface{})`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *Openapi3Operation) HasParameters() bool`

HasParameters returns a boolean if a field has been set.

### GetRequestBody

`func (o *Openapi3Operation) GetRequestBody() map[string]interface{}`

GetRequestBody returns the RequestBody field if non-nil, zero value otherwise.

### GetRequestBodyOk

`func (o *Openapi3Operation) GetRequestBodyOk() (*map[string]interface{}, bool)`

GetRequestBodyOk returns a tuple with the RequestBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestBody

`func (o *Openapi3Operation) SetRequestBody(v map[string]interface{})`

SetRequestBody sets RequestBody field to given value.

### HasRequestBody

`func (o *Openapi3Operation) HasRequestBody() bool`

HasRequestBody returns a boolean if a field has been set.

### GetResponses

`func (o *Openapi3Operation) GetResponses() map[string]map[string]interface{}`

GetResponses returns the Responses field if non-nil, zero value otherwise.

### GetResponsesOk

`func (o *Openapi3Operation) GetResponsesOk() (*map[string]map[string]interface{}, bool)`

GetResponsesOk returns a tuple with the Responses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponses

`func (o *Openapi3Operation) SetResponses(v map[string]map[string]interface{})`

SetResponses sets Responses field to given value.

### HasResponses

`func (o *Openapi3Operation) HasResponses() bool`

HasResponses returns a boolean if a field has been set.

### GetSecurity

`func (o *Openapi3Operation) GetSecurity() []map[string][]string`

GetSecurity returns the Security field if non-nil, zero value otherwise.

### GetSecurityOk

`func (o *Openapi3Operation) GetSecurityOk() (*[]map[string][]string, bool)`

GetSecurityOk returns a tuple with the Security field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurity

`func (o *Openapi3Operation) SetSecurity(v []map[string][]string)`

SetSecurity sets Security field to given value.

### HasSecurity

`func (o *Openapi3Operation) HasSecurity() bool`

HasSecurity returns a boolean if a field has been set.

### SetSecurityNil

`func (o *Openapi3Operation) SetSecurityNil(b bool)`

 SetSecurityNil sets the value for Security to be an explicit nil

### UnsetSecurity
`func (o *Openapi3Operation) UnsetSecurity()`

UnsetSecurity ensures that no value is present for Security, not even an explicit nil
### GetServers

`func (o *Openapi3Operation) GetServers() []Openapi3Server`

GetServers returns the Servers field if non-nil, zero value otherwise.

### GetServersOk

`func (o *Openapi3Operation) GetServersOk() (*[]Openapi3Server, bool)`

GetServersOk returns a tuple with the Servers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServers

`func (o *Openapi3Operation) SetServers(v []Openapi3Server)`

SetServers sets Servers field to given value.

### HasServers

`func (o *Openapi3Operation) HasServers() bool`

HasServers returns a boolean if a field has been set.

### SetServersNil

`func (o *Openapi3Operation) SetServersNil(b bool)`

 SetServersNil sets the value for Servers to be an explicit nil

### UnsetServers
`func (o *Openapi3Operation) UnsetServers()`

UnsetServers ensures that no value is present for Servers, not even an explicit nil
### GetSummary

`func (o *Openapi3Operation) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *Openapi3Operation) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *Openapi3Operation) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *Openapi3Operation) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetTags

`func (o *Openapi3Operation) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Openapi3Operation) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Openapi3Operation) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Openapi3Operation) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


