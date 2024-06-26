# OasPluginConfigData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** |  | [optional] 
**Value** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewOasPluginConfigData

`func NewOasPluginConfigData() *OasPluginConfigData`

NewOasPluginConfigData instantiates a new OasPluginConfigData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOasPluginConfigDataWithDefaults

`func NewOasPluginConfigDataWithDefaults() *OasPluginConfigData`

NewOasPluginConfigDataWithDefaults instantiates a new OasPluginConfigData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *OasPluginConfigData) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *OasPluginConfigData) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *OasPluginConfigData) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *OasPluginConfigData) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetValue

`func (o *OasPluginConfigData) GetValue() map[string]interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *OasPluginConfigData) GetValueOk() (*map[string]interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *OasPluginConfigData) SetValue(v map[string]interface{})`

SetValue sets Value field to given value.

### HasValue

`func (o *OasPluginConfigData) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *OasPluginConfigData) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *OasPluginConfigData) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


