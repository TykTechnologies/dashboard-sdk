# ApidefGraphQLEngineConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataSources** | Pointer to [**[]ApidefGraphQLEngineDataSource**](ApidefGraphQLEngineDataSource.md) |  | [optional] 
**FieldConfigs** | Pointer to [**[]ApidefGraphQLFieldConfig**](ApidefGraphQLFieldConfig.md) |  | [optional] 
**GlobalHeaders** | Pointer to [**[]ApidefUDGGlobalHeader**](ApidefUDGGlobalHeader.md) |  | [optional] 

## Methods

### NewApidefGraphQLEngineConfig

`func NewApidefGraphQLEngineConfig() *ApidefGraphQLEngineConfig`

NewApidefGraphQLEngineConfig instantiates a new ApidefGraphQLEngineConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApidefGraphQLEngineConfigWithDefaults

`func NewApidefGraphQLEngineConfigWithDefaults() *ApidefGraphQLEngineConfig`

NewApidefGraphQLEngineConfigWithDefaults instantiates a new ApidefGraphQLEngineConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataSources

`func (o *ApidefGraphQLEngineConfig) GetDataSources() []ApidefGraphQLEngineDataSource`

GetDataSources returns the DataSources field if non-nil, zero value otherwise.

### GetDataSourcesOk

`func (o *ApidefGraphQLEngineConfig) GetDataSourcesOk() (*[]ApidefGraphQLEngineDataSource, bool)`

GetDataSourcesOk returns a tuple with the DataSources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSources

`func (o *ApidefGraphQLEngineConfig) SetDataSources(v []ApidefGraphQLEngineDataSource)`

SetDataSources sets DataSources field to given value.

### HasDataSources

`func (o *ApidefGraphQLEngineConfig) HasDataSources() bool`

HasDataSources returns a boolean if a field has been set.

### SetDataSourcesNil

`func (o *ApidefGraphQLEngineConfig) SetDataSourcesNil(b bool)`

 SetDataSourcesNil sets the value for DataSources to be an explicit nil

### UnsetDataSources
`func (o *ApidefGraphQLEngineConfig) UnsetDataSources()`

UnsetDataSources ensures that no value is present for DataSources, not even an explicit nil
### GetFieldConfigs

`func (o *ApidefGraphQLEngineConfig) GetFieldConfigs() []ApidefGraphQLFieldConfig`

GetFieldConfigs returns the FieldConfigs field if non-nil, zero value otherwise.

### GetFieldConfigsOk

`func (o *ApidefGraphQLEngineConfig) GetFieldConfigsOk() (*[]ApidefGraphQLFieldConfig, bool)`

GetFieldConfigsOk returns a tuple with the FieldConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldConfigs

`func (o *ApidefGraphQLEngineConfig) SetFieldConfigs(v []ApidefGraphQLFieldConfig)`

SetFieldConfigs sets FieldConfigs field to given value.

### HasFieldConfigs

`func (o *ApidefGraphQLEngineConfig) HasFieldConfigs() bool`

HasFieldConfigs returns a boolean if a field has been set.

### SetFieldConfigsNil

`func (o *ApidefGraphQLEngineConfig) SetFieldConfigsNil(b bool)`

 SetFieldConfigsNil sets the value for FieldConfigs to be an explicit nil

### UnsetFieldConfigs
`func (o *ApidefGraphQLEngineConfig) UnsetFieldConfigs()`

UnsetFieldConfigs ensures that no value is present for FieldConfigs, not even an explicit nil
### GetGlobalHeaders

`func (o *ApidefGraphQLEngineConfig) GetGlobalHeaders() []ApidefUDGGlobalHeader`

GetGlobalHeaders returns the GlobalHeaders field if non-nil, zero value otherwise.

### GetGlobalHeadersOk

`func (o *ApidefGraphQLEngineConfig) GetGlobalHeadersOk() (*[]ApidefUDGGlobalHeader, bool)`

GetGlobalHeadersOk returns a tuple with the GlobalHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalHeaders

`func (o *ApidefGraphQLEngineConfig) SetGlobalHeaders(v []ApidefUDGGlobalHeader)`

SetGlobalHeaders sets GlobalHeaders field to given value.

### HasGlobalHeaders

`func (o *ApidefGraphQLEngineConfig) HasGlobalHeaders() bool`

HasGlobalHeaders returns a boolean if a field has been set.

### SetGlobalHeadersNil

`func (o *ApidefGraphQLEngineConfig) SetGlobalHeadersNil(b bool)`

 SetGlobalHeadersNil sets the value for GlobalHeaders to be an explicit nil

### UnsetGlobalHeaders
`func (o *ApidefGraphQLEngineConfig) UnsetGlobalHeaders()`

UnsetGlobalHeaders ensures that no value is present for GlobalHeaders, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


