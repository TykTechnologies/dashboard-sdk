# OasGlobal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cache** | Pointer to [**OasCache**](OasCache.md) |  | [optional] 
**ContextVariables** | Pointer to [**OasContextVariables**](OasContextVariables.md) |  | [optional] 
**Cors** | Pointer to [**OasCORS**](OasCORS.md) |  | [optional] 
**PluginConfig** | Pointer to [**OasPluginConfig**](OasPluginConfig.md) |  | [optional] 
**PostAuthenticationPlugin** | Pointer to [**OasPostAuthenticationPlugin**](OasPostAuthenticationPlugin.md) |  | [optional] 
**PostAuthenticationPlugins** | Pointer to [**[]OasCustomPlugin**](OasCustomPlugin.md) |  | [optional] 
**PostPlugin** | Pointer to [**OasPostPlugin**](OasPostPlugin.md) |  | [optional] 
**PostPlugins** | Pointer to [**[]OasCustomPlugin**](OasCustomPlugin.md) |  | [optional] 
**PrePlugin** | Pointer to [**OasPrePlugin**](OasPrePlugin.md) |  | [optional] 
**PrePlugins** | Pointer to [**[]OasCustomPlugin**](OasCustomPlugin.md) |  | [optional] 
**ResponsePlugin** | Pointer to [**OasResponsePlugin**](OasResponsePlugin.md) |  | [optional] 
**ResponsePlugins** | Pointer to [**[]OasCustomPlugin**](OasCustomPlugin.md) |  | [optional] 
**TrafficLogs** | Pointer to [**OasTrafficLogs**](OasTrafficLogs.md) |  | [optional] 
**TransformRequestHeaders** | Pointer to [**OasTransformHeaders**](OasTransformHeaders.md) |  | [optional] 
**TransformResponseHeaders** | Pointer to [**OasTransformHeaders**](OasTransformHeaders.md) |  | [optional] 

## Methods

### NewOasGlobal

`func NewOasGlobal() *OasGlobal`

NewOasGlobal instantiates a new OasGlobal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOasGlobalWithDefaults

`func NewOasGlobalWithDefaults() *OasGlobal`

NewOasGlobalWithDefaults instantiates a new OasGlobal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCache

`func (o *OasGlobal) GetCache() OasCache`

GetCache returns the Cache field if non-nil, zero value otherwise.

### GetCacheOk

`func (o *OasGlobal) GetCacheOk() (*OasCache, bool)`

GetCacheOk returns a tuple with the Cache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCache

`func (o *OasGlobal) SetCache(v OasCache)`

SetCache sets Cache field to given value.

### HasCache

`func (o *OasGlobal) HasCache() bool`

HasCache returns a boolean if a field has been set.

### GetContextVariables

`func (o *OasGlobal) GetContextVariables() OasContextVariables`

GetContextVariables returns the ContextVariables field if non-nil, zero value otherwise.

### GetContextVariablesOk

`func (o *OasGlobal) GetContextVariablesOk() (*OasContextVariables, bool)`

GetContextVariablesOk returns a tuple with the ContextVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContextVariables

`func (o *OasGlobal) SetContextVariables(v OasContextVariables)`

SetContextVariables sets ContextVariables field to given value.

### HasContextVariables

`func (o *OasGlobal) HasContextVariables() bool`

HasContextVariables returns a boolean if a field has been set.

### GetCors

`func (o *OasGlobal) GetCors() OasCORS`

GetCors returns the Cors field if non-nil, zero value otherwise.

### GetCorsOk

`func (o *OasGlobal) GetCorsOk() (*OasCORS, bool)`

GetCorsOk returns a tuple with the Cors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCors

`func (o *OasGlobal) SetCors(v OasCORS)`

SetCors sets Cors field to given value.

### HasCors

`func (o *OasGlobal) HasCors() bool`

HasCors returns a boolean if a field has been set.

### GetPluginConfig

`func (o *OasGlobal) GetPluginConfig() OasPluginConfig`

GetPluginConfig returns the PluginConfig field if non-nil, zero value otherwise.

### GetPluginConfigOk

`func (o *OasGlobal) GetPluginConfigOk() (*OasPluginConfig, bool)`

GetPluginConfigOk returns a tuple with the PluginConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPluginConfig

`func (o *OasGlobal) SetPluginConfig(v OasPluginConfig)`

SetPluginConfig sets PluginConfig field to given value.

### HasPluginConfig

`func (o *OasGlobal) HasPluginConfig() bool`

HasPluginConfig returns a boolean if a field has been set.

### GetPostAuthenticationPlugin

`func (o *OasGlobal) GetPostAuthenticationPlugin() OasPostAuthenticationPlugin`

GetPostAuthenticationPlugin returns the PostAuthenticationPlugin field if non-nil, zero value otherwise.

### GetPostAuthenticationPluginOk

`func (o *OasGlobal) GetPostAuthenticationPluginOk() (*OasPostAuthenticationPlugin, bool)`

GetPostAuthenticationPluginOk returns a tuple with the PostAuthenticationPlugin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostAuthenticationPlugin

`func (o *OasGlobal) SetPostAuthenticationPlugin(v OasPostAuthenticationPlugin)`

SetPostAuthenticationPlugin sets PostAuthenticationPlugin field to given value.

### HasPostAuthenticationPlugin

`func (o *OasGlobal) HasPostAuthenticationPlugin() bool`

HasPostAuthenticationPlugin returns a boolean if a field has been set.

### GetPostAuthenticationPlugins

`func (o *OasGlobal) GetPostAuthenticationPlugins() []OasCustomPlugin`

GetPostAuthenticationPlugins returns the PostAuthenticationPlugins field if non-nil, zero value otherwise.

### GetPostAuthenticationPluginsOk

`func (o *OasGlobal) GetPostAuthenticationPluginsOk() (*[]OasCustomPlugin, bool)`

GetPostAuthenticationPluginsOk returns a tuple with the PostAuthenticationPlugins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostAuthenticationPlugins

`func (o *OasGlobal) SetPostAuthenticationPlugins(v []OasCustomPlugin)`

SetPostAuthenticationPlugins sets PostAuthenticationPlugins field to given value.

### HasPostAuthenticationPlugins

`func (o *OasGlobal) HasPostAuthenticationPlugins() bool`

HasPostAuthenticationPlugins returns a boolean if a field has been set.

### GetPostPlugin

`func (o *OasGlobal) GetPostPlugin() OasPostPlugin`

GetPostPlugin returns the PostPlugin field if non-nil, zero value otherwise.

### GetPostPluginOk

`func (o *OasGlobal) GetPostPluginOk() (*OasPostPlugin, bool)`

GetPostPluginOk returns a tuple with the PostPlugin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostPlugin

`func (o *OasGlobal) SetPostPlugin(v OasPostPlugin)`

SetPostPlugin sets PostPlugin field to given value.

### HasPostPlugin

`func (o *OasGlobal) HasPostPlugin() bool`

HasPostPlugin returns a boolean if a field has been set.

### GetPostPlugins

`func (o *OasGlobal) GetPostPlugins() []OasCustomPlugin`

GetPostPlugins returns the PostPlugins field if non-nil, zero value otherwise.

### GetPostPluginsOk

`func (o *OasGlobal) GetPostPluginsOk() (*[]OasCustomPlugin, bool)`

GetPostPluginsOk returns a tuple with the PostPlugins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostPlugins

`func (o *OasGlobal) SetPostPlugins(v []OasCustomPlugin)`

SetPostPlugins sets PostPlugins field to given value.

### HasPostPlugins

`func (o *OasGlobal) HasPostPlugins() bool`

HasPostPlugins returns a boolean if a field has been set.

### GetPrePlugin

`func (o *OasGlobal) GetPrePlugin() OasPrePlugin`

GetPrePlugin returns the PrePlugin field if non-nil, zero value otherwise.

### GetPrePluginOk

`func (o *OasGlobal) GetPrePluginOk() (*OasPrePlugin, bool)`

GetPrePluginOk returns a tuple with the PrePlugin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrePlugin

`func (o *OasGlobal) SetPrePlugin(v OasPrePlugin)`

SetPrePlugin sets PrePlugin field to given value.

### HasPrePlugin

`func (o *OasGlobal) HasPrePlugin() bool`

HasPrePlugin returns a boolean if a field has been set.

### GetPrePlugins

`func (o *OasGlobal) GetPrePlugins() []OasCustomPlugin`

GetPrePlugins returns the PrePlugins field if non-nil, zero value otherwise.

### GetPrePluginsOk

`func (o *OasGlobal) GetPrePluginsOk() (*[]OasCustomPlugin, bool)`

GetPrePluginsOk returns a tuple with the PrePlugins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrePlugins

`func (o *OasGlobal) SetPrePlugins(v []OasCustomPlugin)`

SetPrePlugins sets PrePlugins field to given value.

### HasPrePlugins

`func (o *OasGlobal) HasPrePlugins() bool`

HasPrePlugins returns a boolean if a field has been set.

### GetResponsePlugin

`func (o *OasGlobal) GetResponsePlugin() OasResponsePlugin`

GetResponsePlugin returns the ResponsePlugin field if non-nil, zero value otherwise.

### GetResponsePluginOk

`func (o *OasGlobal) GetResponsePluginOk() (*OasResponsePlugin, bool)`

GetResponsePluginOk returns a tuple with the ResponsePlugin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponsePlugin

`func (o *OasGlobal) SetResponsePlugin(v OasResponsePlugin)`

SetResponsePlugin sets ResponsePlugin field to given value.

### HasResponsePlugin

`func (o *OasGlobal) HasResponsePlugin() bool`

HasResponsePlugin returns a boolean if a field has been set.

### GetResponsePlugins

`func (o *OasGlobal) GetResponsePlugins() []OasCustomPlugin`

GetResponsePlugins returns the ResponsePlugins field if non-nil, zero value otherwise.

### GetResponsePluginsOk

`func (o *OasGlobal) GetResponsePluginsOk() (*[]OasCustomPlugin, bool)`

GetResponsePluginsOk returns a tuple with the ResponsePlugins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponsePlugins

`func (o *OasGlobal) SetResponsePlugins(v []OasCustomPlugin)`

SetResponsePlugins sets ResponsePlugins field to given value.

### HasResponsePlugins

`func (o *OasGlobal) HasResponsePlugins() bool`

HasResponsePlugins returns a boolean if a field has been set.

### GetTrafficLogs

`func (o *OasGlobal) GetTrafficLogs() OasTrafficLogs`

GetTrafficLogs returns the TrafficLogs field if non-nil, zero value otherwise.

### GetTrafficLogsOk

`func (o *OasGlobal) GetTrafficLogsOk() (*OasTrafficLogs, bool)`

GetTrafficLogsOk returns a tuple with the TrafficLogs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrafficLogs

`func (o *OasGlobal) SetTrafficLogs(v OasTrafficLogs)`

SetTrafficLogs sets TrafficLogs field to given value.

### HasTrafficLogs

`func (o *OasGlobal) HasTrafficLogs() bool`

HasTrafficLogs returns a boolean if a field has been set.

### GetTransformRequestHeaders

`func (o *OasGlobal) GetTransformRequestHeaders() OasTransformHeaders`

GetTransformRequestHeaders returns the TransformRequestHeaders field if non-nil, zero value otherwise.

### GetTransformRequestHeadersOk

`func (o *OasGlobal) GetTransformRequestHeadersOk() (*OasTransformHeaders, bool)`

GetTransformRequestHeadersOk returns a tuple with the TransformRequestHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransformRequestHeaders

`func (o *OasGlobal) SetTransformRequestHeaders(v OasTransformHeaders)`

SetTransformRequestHeaders sets TransformRequestHeaders field to given value.

### HasTransformRequestHeaders

`func (o *OasGlobal) HasTransformRequestHeaders() bool`

HasTransformRequestHeaders returns a boolean if a field has been set.

### GetTransformResponseHeaders

`func (o *OasGlobal) GetTransformResponseHeaders() OasTransformHeaders`

GetTransformResponseHeaders returns the TransformResponseHeaders field if non-nil, zero value otherwise.

### GetTransformResponseHeadersOk

`func (o *OasGlobal) GetTransformResponseHeadersOk() (*OasTransformHeaders, bool)`

GetTransformResponseHeadersOk returns a tuple with the TransformResponseHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransformResponseHeaders

`func (o *OasGlobal) SetTransformResponseHeaders(v OasTransformHeaders)`

SetTransformResponseHeaders sets TransformResponseHeaders field to given value.

### HasTransformResponseHeaders

`func (o *OasGlobal) HasTransformResponseHeaders() bool`

HasTransformResponseHeaders returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


