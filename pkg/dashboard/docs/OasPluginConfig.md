# OasPluginConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bundle** | Pointer to [**OasPluginBundle**](OasPluginBundle.md) |  | [optional] 
**Data** | Pointer to [**OasPluginConfigData**](OasPluginConfigData.md) |  | [optional] 
**Driver** | Pointer to **string** |  | [optional] 

## Methods

### NewOasPluginConfig

`func NewOasPluginConfig() *OasPluginConfig`

NewOasPluginConfig instantiates a new OasPluginConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOasPluginConfigWithDefaults

`func NewOasPluginConfigWithDefaults() *OasPluginConfig`

NewOasPluginConfigWithDefaults instantiates a new OasPluginConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBundle

`func (o *OasPluginConfig) GetBundle() OasPluginBundle`

GetBundle returns the Bundle field if non-nil, zero value otherwise.

### GetBundleOk

`func (o *OasPluginConfig) GetBundleOk() (*OasPluginBundle, bool)`

GetBundleOk returns a tuple with the Bundle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBundle

`func (o *OasPluginConfig) SetBundle(v OasPluginBundle)`

SetBundle sets Bundle field to given value.

### HasBundle

`func (o *OasPluginConfig) HasBundle() bool`

HasBundle returns a boolean if a field has been set.

### GetData

`func (o *OasPluginConfig) GetData() OasPluginConfigData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *OasPluginConfig) GetDataOk() (*OasPluginConfigData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *OasPluginConfig) SetData(v OasPluginConfigData)`

SetData sets Data field to given value.

### HasData

`func (o *OasPluginConfig) HasData() bool`

HasData returns a boolean if a field has been set.

### GetDriver

`func (o *OasPluginConfig) GetDriver() string`

GetDriver returns the Driver field if non-nil, zero value otherwise.

### GetDriverOk

`func (o *OasPluginConfig) GetDriverOk() (*string, bool)`

GetDriverOk returns a tuple with the Driver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriver

`func (o *OasPluginConfig) SetDriver(v string)`

SetDriver sets Driver field to given value.

### HasDriver

`func (o *OasPluginConfig) HasDriver() bool`

HasDriver returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


