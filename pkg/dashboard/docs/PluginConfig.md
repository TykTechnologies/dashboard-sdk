# PluginConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bundle** | Pointer to [**NullablePluginBundle**](PluginBundle.md) |  | [optional] 
**Data** | Pointer to [**NullablePluginConfigData**](PluginConfigData.md) |  | [optional] 
**Driver** | Pointer to **string** |  | [optional] 

## Methods

### NewPluginConfig

`func NewPluginConfig() *PluginConfig`

NewPluginConfig instantiates a new PluginConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPluginConfigWithDefaults

`func NewPluginConfigWithDefaults() *PluginConfig`

NewPluginConfigWithDefaults instantiates a new PluginConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBundle

`func (o *PluginConfig) GetBundle() PluginBundle`

GetBundle returns the Bundle field if non-nil, zero value otherwise.

### GetBundleOk

`func (o *PluginConfig) GetBundleOk() (*PluginBundle, bool)`

GetBundleOk returns a tuple with the Bundle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBundle

`func (o *PluginConfig) SetBundle(v PluginBundle)`

SetBundle sets Bundle field to given value.

### HasBundle

`func (o *PluginConfig) HasBundle() bool`

HasBundle returns a boolean if a field has been set.

### SetBundleNil

`func (o *PluginConfig) SetBundleNil(b bool)`

 SetBundleNil sets the value for Bundle to be an explicit nil

### UnsetBundle
`func (o *PluginConfig) UnsetBundle()`

UnsetBundle ensures that no value is present for Bundle, not even an explicit nil
### GetData

`func (o *PluginConfig) GetData() PluginConfigData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *PluginConfig) GetDataOk() (*PluginConfigData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *PluginConfig) SetData(v PluginConfigData)`

SetData sets Data field to given value.

### HasData

`func (o *PluginConfig) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *PluginConfig) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *PluginConfig) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetDriver

`func (o *PluginConfig) GetDriver() string`

GetDriver returns the Driver field if non-nil, zero value otherwise.

### GetDriverOk

`func (o *PluginConfig) GetDriverOk() (*string, bool)`

GetDriverOk returns a tuple with the Driver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriver

`func (o *PluginConfig) SetDriver(v string)`

SetDriver sets Driver field to given value.

### HasDriver

`func (o *PluginConfig) HasDriver() bool`

HasDriver returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


