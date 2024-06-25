# ApidefServiceDiscoveryConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CacheDisabled** | Pointer to **bool** |  | [optional] 
**CacheTimeout** | Pointer to **int32** |  | [optional] 
**DataPath** | Pointer to **string** |  | [optional] 
**EndpointReturnsList** | Pointer to **bool** |  | [optional] 
**ParentDataPath** | Pointer to **string** |  | [optional] 
**PortDataPath** | Pointer to **string** |  | [optional] 
**QueryEndpoint** | Pointer to **string** |  | [optional] 
**TargetPath** | Pointer to **string** |  | [optional] 
**UseDiscoveryService** | Pointer to **bool** |  | [optional] 
**UseNestedQuery** | Pointer to **bool** |  | [optional] 
**UseTargetList** | Pointer to **bool** |  | [optional] 

## Methods

### NewApidefServiceDiscoveryConfiguration

`func NewApidefServiceDiscoveryConfiguration() *ApidefServiceDiscoveryConfiguration`

NewApidefServiceDiscoveryConfiguration instantiates a new ApidefServiceDiscoveryConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApidefServiceDiscoveryConfigurationWithDefaults

`func NewApidefServiceDiscoveryConfigurationWithDefaults() *ApidefServiceDiscoveryConfiguration`

NewApidefServiceDiscoveryConfigurationWithDefaults instantiates a new ApidefServiceDiscoveryConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCacheDisabled

`func (o *ApidefServiceDiscoveryConfiguration) GetCacheDisabled() bool`

GetCacheDisabled returns the CacheDisabled field if non-nil, zero value otherwise.

### GetCacheDisabledOk

`func (o *ApidefServiceDiscoveryConfiguration) GetCacheDisabledOk() (*bool, bool)`

GetCacheDisabledOk returns a tuple with the CacheDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheDisabled

`func (o *ApidefServiceDiscoveryConfiguration) SetCacheDisabled(v bool)`

SetCacheDisabled sets CacheDisabled field to given value.

### HasCacheDisabled

`func (o *ApidefServiceDiscoveryConfiguration) HasCacheDisabled() bool`

HasCacheDisabled returns a boolean if a field has been set.

### GetCacheTimeout

`func (o *ApidefServiceDiscoveryConfiguration) GetCacheTimeout() int32`

GetCacheTimeout returns the CacheTimeout field if non-nil, zero value otherwise.

### GetCacheTimeoutOk

`func (o *ApidefServiceDiscoveryConfiguration) GetCacheTimeoutOk() (*int32, bool)`

GetCacheTimeoutOk returns a tuple with the CacheTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheTimeout

`func (o *ApidefServiceDiscoveryConfiguration) SetCacheTimeout(v int32)`

SetCacheTimeout sets CacheTimeout field to given value.

### HasCacheTimeout

`func (o *ApidefServiceDiscoveryConfiguration) HasCacheTimeout() bool`

HasCacheTimeout returns a boolean if a field has been set.

### GetDataPath

`func (o *ApidefServiceDiscoveryConfiguration) GetDataPath() string`

GetDataPath returns the DataPath field if non-nil, zero value otherwise.

### GetDataPathOk

`func (o *ApidefServiceDiscoveryConfiguration) GetDataPathOk() (*string, bool)`

GetDataPathOk returns a tuple with the DataPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataPath

`func (o *ApidefServiceDiscoveryConfiguration) SetDataPath(v string)`

SetDataPath sets DataPath field to given value.

### HasDataPath

`func (o *ApidefServiceDiscoveryConfiguration) HasDataPath() bool`

HasDataPath returns a boolean if a field has been set.

### GetEndpointReturnsList

`func (o *ApidefServiceDiscoveryConfiguration) GetEndpointReturnsList() bool`

GetEndpointReturnsList returns the EndpointReturnsList field if non-nil, zero value otherwise.

### GetEndpointReturnsListOk

`func (o *ApidefServiceDiscoveryConfiguration) GetEndpointReturnsListOk() (*bool, bool)`

GetEndpointReturnsListOk returns a tuple with the EndpointReturnsList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointReturnsList

`func (o *ApidefServiceDiscoveryConfiguration) SetEndpointReturnsList(v bool)`

SetEndpointReturnsList sets EndpointReturnsList field to given value.

### HasEndpointReturnsList

`func (o *ApidefServiceDiscoveryConfiguration) HasEndpointReturnsList() bool`

HasEndpointReturnsList returns a boolean if a field has been set.

### GetParentDataPath

`func (o *ApidefServiceDiscoveryConfiguration) GetParentDataPath() string`

GetParentDataPath returns the ParentDataPath field if non-nil, zero value otherwise.

### GetParentDataPathOk

`func (o *ApidefServiceDiscoveryConfiguration) GetParentDataPathOk() (*string, bool)`

GetParentDataPathOk returns a tuple with the ParentDataPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentDataPath

`func (o *ApidefServiceDiscoveryConfiguration) SetParentDataPath(v string)`

SetParentDataPath sets ParentDataPath field to given value.

### HasParentDataPath

`func (o *ApidefServiceDiscoveryConfiguration) HasParentDataPath() bool`

HasParentDataPath returns a boolean if a field has been set.

### GetPortDataPath

`func (o *ApidefServiceDiscoveryConfiguration) GetPortDataPath() string`

GetPortDataPath returns the PortDataPath field if non-nil, zero value otherwise.

### GetPortDataPathOk

`func (o *ApidefServiceDiscoveryConfiguration) GetPortDataPathOk() (*string, bool)`

GetPortDataPathOk returns a tuple with the PortDataPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortDataPath

`func (o *ApidefServiceDiscoveryConfiguration) SetPortDataPath(v string)`

SetPortDataPath sets PortDataPath field to given value.

### HasPortDataPath

`func (o *ApidefServiceDiscoveryConfiguration) HasPortDataPath() bool`

HasPortDataPath returns a boolean if a field has been set.

### GetQueryEndpoint

`func (o *ApidefServiceDiscoveryConfiguration) GetQueryEndpoint() string`

GetQueryEndpoint returns the QueryEndpoint field if non-nil, zero value otherwise.

### GetQueryEndpointOk

`func (o *ApidefServiceDiscoveryConfiguration) GetQueryEndpointOk() (*string, bool)`

GetQueryEndpointOk returns a tuple with the QueryEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryEndpoint

`func (o *ApidefServiceDiscoveryConfiguration) SetQueryEndpoint(v string)`

SetQueryEndpoint sets QueryEndpoint field to given value.

### HasQueryEndpoint

`func (o *ApidefServiceDiscoveryConfiguration) HasQueryEndpoint() bool`

HasQueryEndpoint returns a boolean if a field has been set.

### GetTargetPath

`func (o *ApidefServiceDiscoveryConfiguration) GetTargetPath() string`

GetTargetPath returns the TargetPath field if non-nil, zero value otherwise.

### GetTargetPathOk

`func (o *ApidefServiceDiscoveryConfiguration) GetTargetPathOk() (*string, bool)`

GetTargetPathOk returns a tuple with the TargetPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetPath

`func (o *ApidefServiceDiscoveryConfiguration) SetTargetPath(v string)`

SetTargetPath sets TargetPath field to given value.

### HasTargetPath

`func (o *ApidefServiceDiscoveryConfiguration) HasTargetPath() bool`

HasTargetPath returns a boolean if a field has been set.

### GetUseDiscoveryService

`func (o *ApidefServiceDiscoveryConfiguration) GetUseDiscoveryService() bool`

GetUseDiscoveryService returns the UseDiscoveryService field if non-nil, zero value otherwise.

### GetUseDiscoveryServiceOk

`func (o *ApidefServiceDiscoveryConfiguration) GetUseDiscoveryServiceOk() (*bool, bool)`

GetUseDiscoveryServiceOk returns a tuple with the UseDiscoveryService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseDiscoveryService

`func (o *ApidefServiceDiscoveryConfiguration) SetUseDiscoveryService(v bool)`

SetUseDiscoveryService sets UseDiscoveryService field to given value.

### HasUseDiscoveryService

`func (o *ApidefServiceDiscoveryConfiguration) HasUseDiscoveryService() bool`

HasUseDiscoveryService returns a boolean if a field has been set.

### GetUseNestedQuery

`func (o *ApidefServiceDiscoveryConfiguration) GetUseNestedQuery() bool`

GetUseNestedQuery returns the UseNestedQuery field if non-nil, zero value otherwise.

### GetUseNestedQueryOk

`func (o *ApidefServiceDiscoveryConfiguration) GetUseNestedQueryOk() (*bool, bool)`

GetUseNestedQueryOk returns a tuple with the UseNestedQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseNestedQuery

`func (o *ApidefServiceDiscoveryConfiguration) SetUseNestedQuery(v bool)`

SetUseNestedQuery sets UseNestedQuery field to given value.

### HasUseNestedQuery

`func (o *ApidefServiceDiscoveryConfiguration) HasUseNestedQuery() bool`

HasUseNestedQuery returns a boolean if a field has been set.

### GetUseTargetList

`func (o *ApidefServiceDiscoveryConfiguration) GetUseTargetList() bool`

GetUseTargetList returns the UseTargetList field if non-nil, zero value otherwise.

### GetUseTargetListOk

`func (o *ApidefServiceDiscoveryConfiguration) GetUseTargetListOk() (*bool, bool)`

GetUseTargetListOk returns a tuple with the UseTargetList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseTargetList

`func (o *ApidefServiceDiscoveryConfiguration) SetUseTargetList(v bool)`

SetUseTargetList sets UseTargetList field to given value.

### HasUseTargetList

`func (o *ApidefServiceDiscoveryConfiguration) HasUseTargetList() bool`

HasUseTargetList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


