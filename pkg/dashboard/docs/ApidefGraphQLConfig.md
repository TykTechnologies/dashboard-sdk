# ApidefGraphQLConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** |  | [optional] 
**Engine** | Pointer to [**ApidefGraphQLEngineConfig**](ApidefGraphQLEngineConfig.md) |  | [optional] 
**ExecutionMode** | Pointer to **string** |  | [optional] 
**Introspection** | Pointer to [**ApidefGraphQLIntrospectionConfig**](ApidefGraphQLIntrospectionConfig.md) |  | [optional] 
**LastSchemaUpdate** | Pointer to **NullableTime** |  | [optional] 
**Playground** | Pointer to [**ApidefGraphQLPlayground**](ApidefGraphQLPlayground.md) |  | [optional] 
**Proxy** | Pointer to [**ApidefGraphQLProxyConfig**](ApidefGraphQLProxyConfig.md) |  | [optional] 
**Schema** | Pointer to **string** |  | [optional] 
**Subgraph** | Pointer to [**ApidefGraphQLSubgraphConfig**](ApidefGraphQLSubgraphConfig.md) |  | [optional] 
**Supergraph** | Pointer to [**ApidefGraphQLSupergraphConfig**](ApidefGraphQLSupergraphConfig.md) |  | [optional] 
**TypeFieldConfigurations** | Pointer to [**[]DatasourceTypeFieldConfiguration**](DatasourceTypeFieldConfiguration.md) |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 

## Methods

### NewApidefGraphQLConfig

`func NewApidefGraphQLConfig() *ApidefGraphQLConfig`

NewApidefGraphQLConfig instantiates a new ApidefGraphQLConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApidefGraphQLConfigWithDefaults

`func NewApidefGraphQLConfigWithDefaults() *ApidefGraphQLConfig`

NewApidefGraphQLConfigWithDefaults instantiates a new ApidefGraphQLConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *ApidefGraphQLConfig) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ApidefGraphQLConfig) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ApidefGraphQLConfig) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ApidefGraphQLConfig) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetEngine

`func (o *ApidefGraphQLConfig) GetEngine() ApidefGraphQLEngineConfig`

GetEngine returns the Engine field if non-nil, zero value otherwise.

### GetEngineOk

`func (o *ApidefGraphQLConfig) GetEngineOk() (*ApidefGraphQLEngineConfig, bool)`

GetEngineOk returns a tuple with the Engine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngine

`func (o *ApidefGraphQLConfig) SetEngine(v ApidefGraphQLEngineConfig)`

SetEngine sets Engine field to given value.

### HasEngine

`func (o *ApidefGraphQLConfig) HasEngine() bool`

HasEngine returns a boolean if a field has been set.

### GetExecutionMode

`func (o *ApidefGraphQLConfig) GetExecutionMode() string`

GetExecutionMode returns the ExecutionMode field if non-nil, zero value otherwise.

### GetExecutionModeOk

`func (o *ApidefGraphQLConfig) GetExecutionModeOk() (*string, bool)`

GetExecutionModeOk returns a tuple with the ExecutionMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionMode

`func (o *ApidefGraphQLConfig) SetExecutionMode(v string)`

SetExecutionMode sets ExecutionMode field to given value.

### HasExecutionMode

`func (o *ApidefGraphQLConfig) HasExecutionMode() bool`

HasExecutionMode returns a boolean if a field has been set.

### GetIntrospection

`func (o *ApidefGraphQLConfig) GetIntrospection() ApidefGraphQLIntrospectionConfig`

GetIntrospection returns the Introspection field if non-nil, zero value otherwise.

### GetIntrospectionOk

`func (o *ApidefGraphQLConfig) GetIntrospectionOk() (*ApidefGraphQLIntrospectionConfig, bool)`

GetIntrospectionOk returns a tuple with the Introspection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntrospection

`func (o *ApidefGraphQLConfig) SetIntrospection(v ApidefGraphQLIntrospectionConfig)`

SetIntrospection sets Introspection field to given value.

### HasIntrospection

`func (o *ApidefGraphQLConfig) HasIntrospection() bool`

HasIntrospection returns a boolean if a field has been set.

### GetLastSchemaUpdate

`func (o *ApidefGraphQLConfig) GetLastSchemaUpdate() time.Time`

GetLastSchemaUpdate returns the LastSchemaUpdate field if non-nil, zero value otherwise.

### GetLastSchemaUpdateOk

`func (o *ApidefGraphQLConfig) GetLastSchemaUpdateOk() (*time.Time, bool)`

GetLastSchemaUpdateOk returns a tuple with the LastSchemaUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSchemaUpdate

`func (o *ApidefGraphQLConfig) SetLastSchemaUpdate(v time.Time)`

SetLastSchemaUpdate sets LastSchemaUpdate field to given value.

### HasLastSchemaUpdate

`func (o *ApidefGraphQLConfig) HasLastSchemaUpdate() bool`

HasLastSchemaUpdate returns a boolean if a field has been set.

### SetLastSchemaUpdateNil

`func (o *ApidefGraphQLConfig) SetLastSchemaUpdateNil(b bool)`

 SetLastSchemaUpdateNil sets the value for LastSchemaUpdate to be an explicit nil

### UnsetLastSchemaUpdate
`func (o *ApidefGraphQLConfig) UnsetLastSchemaUpdate()`

UnsetLastSchemaUpdate ensures that no value is present for LastSchemaUpdate, not even an explicit nil
### GetPlayground

`func (o *ApidefGraphQLConfig) GetPlayground() ApidefGraphQLPlayground`

GetPlayground returns the Playground field if non-nil, zero value otherwise.

### GetPlaygroundOk

`func (o *ApidefGraphQLConfig) GetPlaygroundOk() (*ApidefGraphQLPlayground, bool)`

GetPlaygroundOk returns a tuple with the Playground field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayground

`func (o *ApidefGraphQLConfig) SetPlayground(v ApidefGraphQLPlayground)`

SetPlayground sets Playground field to given value.

### HasPlayground

`func (o *ApidefGraphQLConfig) HasPlayground() bool`

HasPlayground returns a boolean if a field has been set.

### GetProxy

`func (o *ApidefGraphQLConfig) GetProxy() ApidefGraphQLProxyConfig`

GetProxy returns the Proxy field if non-nil, zero value otherwise.

### GetProxyOk

`func (o *ApidefGraphQLConfig) GetProxyOk() (*ApidefGraphQLProxyConfig, bool)`

GetProxyOk returns a tuple with the Proxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxy

`func (o *ApidefGraphQLConfig) SetProxy(v ApidefGraphQLProxyConfig)`

SetProxy sets Proxy field to given value.

### HasProxy

`func (o *ApidefGraphQLConfig) HasProxy() bool`

HasProxy returns a boolean if a field has been set.

### GetSchema

`func (o *ApidefGraphQLConfig) GetSchema() string`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *ApidefGraphQLConfig) GetSchemaOk() (*string, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *ApidefGraphQLConfig) SetSchema(v string)`

SetSchema sets Schema field to given value.

### HasSchema

`func (o *ApidefGraphQLConfig) HasSchema() bool`

HasSchema returns a boolean if a field has been set.

### GetSubgraph

`func (o *ApidefGraphQLConfig) GetSubgraph() ApidefGraphQLSubgraphConfig`

GetSubgraph returns the Subgraph field if non-nil, zero value otherwise.

### GetSubgraphOk

`func (o *ApidefGraphQLConfig) GetSubgraphOk() (*ApidefGraphQLSubgraphConfig, bool)`

GetSubgraphOk returns a tuple with the Subgraph field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubgraph

`func (o *ApidefGraphQLConfig) SetSubgraph(v ApidefGraphQLSubgraphConfig)`

SetSubgraph sets Subgraph field to given value.

### HasSubgraph

`func (o *ApidefGraphQLConfig) HasSubgraph() bool`

HasSubgraph returns a boolean if a field has been set.

### GetSupergraph

`func (o *ApidefGraphQLConfig) GetSupergraph() ApidefGraphQLSupergraphConfig`

GetSupergraph returns the Supergraph field if non-nil, zero value otherwise.

### GetSupergraphOk

`func (o *ApidefGraphQLConfig) GetSupergraphOk() (*ApidefGraphQLSupergraphConfig, bool)`

GetSupergraphOk returns a tuple with the Supergraph field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupergraph

`func (o *ApidefGraphQLConfig) SetSupergraph(v ApidefGraphQLSupergraphConfig)`

SetSupergraph sets Supergraph field to given value.

### HasSupergraph

`func (o *ApidefGraphQLConfig) HasSupergraph() bool`

HasSupergraph returns a boolean if a field has been set.

### GetTypeFieldConfigurations

`func (o *ApidefGraphQLConfig) GetTypeFieldConfigurations() []DatasourceTypeFieldConfiguration`

GetTypeFieldConfigurations returns the TypeFieldConfigurations field if non-nil, zero value otherwise.

### GetTypeFieldConfigurationsOk

`func (o *ApidefGraphQLConfig) GetTypeFieldConfigurationsOk() (*[]DatasourceTypeFieldConfiguration, bool)`

GetTypeFieldConfigurationsOk returns a tuple with the TypeFieldConfigurations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeFieldConfigurations

`func (o *ApidefGraphQLConfig) SetTypeFieldConfigurations(v []DatasourceTypeFieldConfiguration)`

SetTypeFieldConfigurations sets TypeFieldConfigurations field to given value.

### HasTypeFieldConfigurations

`func (o *ApidefGraphQLConfig) HasTypeFieldConfigurations() bool`

HasTypeFieldConfigurations returns a boolean if a field has been set.

### SetTypeFieldConfigurationsNil

`func (o *ApidefGraphQLConfig) SetTypeFieldConfigurationsNil(b bool)`

 SetTypeFieldConfigurationsNil sets the value for TypeFieldConfigurations to be an explicit nil

### UnsetTypeFieldConfigurations
`func (o *ApidefGraphQLConfig) UnsetTypeFieldConfigurations()`

UnsetTypeFieldConfigurations ensures that no value is present for TypeFieldConfigurations, not even an explicit nil
### GetVersion

`func (o *ApidefGraphQLConfig) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ApidefGraphQLConfig) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ApidefGraphQLConfig) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ApidefGraphQLConfig) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


