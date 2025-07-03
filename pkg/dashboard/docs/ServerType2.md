# ServerType2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 
**Variables** | Pointer to [**map[string]ServerVariable**](ServerVariable.md) |  | [optional] 

## Methods

### NewServerType2

`func NewServerType2() *ServerType2`

NewServerType2 instantiates a new ServerType2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerType2WithDefaults

`func NewServerType2WithDefaults() *ServerType2`

NewServerType2WithDefaults instantiates a new ServerType2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *ServerType2) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ServerType2) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ServerType2) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ServerType2) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetUrl

`func (o *ServerType2) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ServerType2) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ServerType2) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *ServerType2) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetVariables

`func (o *ServerType2) GetVariables() map[string]ServerVariable`

GetVariables returns the Variables field if non-nil, zero value otherwise.

### GetVariablesOk

`func (o *ServerType2) GetVariablesOk() (*map[string]ServerVariable, bool)`

GetVariablesOk returns a tuple with the Variables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariables

`func (o *ServerType2) SetVariables(v map[string]ServerVariable)`

SetVariables sets Variables field to given value.

### HasVariables

`func (o *ServerType2) HasVariables() bool`

HasVariables returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


