# Openapi3Server

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 
**Variables** | Pointer to [**map[string]Openapi3ServerVariable**](Openapi3ServerVariable.md) |  | [optional] 

## Methods

### NewOpenapi3Server

`func NewOpenapi3Server() *Openapi3Server`

NewOpenapi3Server instantiates a new Openapi3Server object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpenapi3ServerWithDefaults

`func NewOpenapi3ServerWithDefaults() *Openapi3Server`

NewOpenapi3ServerWithDefaults instantiates a new Openapi3Server object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *Openapi3Server) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Openapi3Server) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Openapi3Server) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Openapi3Server) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetUrl

`func (o *Openapi3Server) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Openapi3Server) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Openapi3Server) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *Openapi3Server) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetVariables

`func (o *Openapi3Server) GetVariables() map[string]Openapi3ServerVariable`

GetVariables returns the Variables field if non-nil, zero value otherwise.

### GetVariablesOk

`func (o *Openapi3Server) GetVariablesOk() (*map[string]Openapi3ServerVariable, bool)`

GetVariablesOk returns a tuple with the Variables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariables

`func (o *Openapi3Server) SetVariables(v map[string]Openapi3ServerVariable)`

SetVariables sets Variables field to given value.

### HasVariables

`func (o *Openapi3Server) HasVariables() bool`

HasVariables returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


