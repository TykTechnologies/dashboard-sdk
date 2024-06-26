# OasServiceDiscovery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cache** | Pointer to [**OasServiceDiscoveryCache**](OasServiceDiscoveryCache.md) |  | [optional] 
**CacheTimeout** | Pointer to **int32** |  | [optional] 
**DataPath** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**EndpointReturnsList** | Pointer to **bool** |  | [optional] 
**ParentDataPath** | Pointer to **string** |  | [optional] 
**PortDataPath** | Pointer to **string** |  | [optional] 
**QueryEndpoint** | Pointer to **string** |  | [optional] 
**TargetPath** | Pointer to **string** |  | [optional] 
**UseNestedQuery** | Pointer to **bool** |  | [optional] 
**UseTargetList** | Pointer to **bool** |  | [optional] 

## Methods

### NewOasServiceDiscovery

`func NewOasServiceDiscovery() *OasServiceDiscovery`

NewOasServiceDiscovery instantiates a new OasServiceDiscovery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOasServiceDiscoveryWithDefaults

`func NewOasServiceDiscoveryWithDefaults() *OasServiceDiscovery`

NewOasServiceDiscoveryWithDefaults instantiates a new OasServiceDiscovery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCache

`func (o *OasServiceDiscovery) GetCache() OasServiceDiscoveryCache`

GetCache returns the Cache field if non-nil, zero value otherwise.

### GetCacheOk

`func (o *OasServiceDiscovery) GetCacheOk() (*OasServiceDiscoveryCache, bool)`

GetCacheOk returns a tuple with the Cache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCache

`func (o *OasServiceDiscovery) SetCache(v OasServiceDiscoveryCache)`

SetCache sets Cache field to given value.

### HasCache

`func (o *OasServiceDiscovery) HasCache() bool`

HasCache returns a boolean if a field has been set.

### GetCacheTimeout

`func (o *OasServiceDiscovery) GetCacheTimeout() int32`

GetCacheTimeout returns the CacheTimeout field if non-nil, zero value otherwise.

### GetCacheTimeoutOk

`func (o *OasServiceDiscovery) GetCacheTimeoutOk() (*int32, bool)`

GetCacheTimeoutOk returns a tuple with the CacheTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheTimeout

`func (o *OasServiceDiscovery) SetCacheTimeout(v int32)`

SetCacheTimeout sets CacheTimeout field to given value.

### HasCacheTimeout

`func (o *OasServiceDiscovery) HasCacheTimeout() bool`

HasCacheTimeout returns a boolean if a field has been set.

### GetDataPath

`func (o *OasServiceDiscovery) GetDataPath() string`

GetDataPath returns the DataPath field if non-nil, zero value otherwise.

### GetDataPathOk

`func (o *OasServiceDiscovery) GetDataPathOk() (*string, bool)`

GetDataPathOk returns a tuple with the DataPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataPath

`func (o *OasServiceDiscovery) SetDataPath(v string)`

SetDataPath sets DataPath field to given value.

### HasDataPath

`func (o *OasServiceDiscovery) HasDataPath() bool`

HasDataPath returns a boolean if a field has been set.

### GetEnabled

`func (o *OasServiceDiscovery) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *OasServiceDiscovery) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *OasServiceDiscovery) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *OasServiceDiscovery) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetEndpointReturnsList

`func (o *OasServiceDiscovery) GetEndpointReturnsList() bool`

GetEndpointReturnsList returns the EndpointReturnsList field if non-nil, zero value otherwise.

### GetEndpointReturnsListOk

`func (o *OasServiceDiscovery) GetEndpointReturnsListOk() (*bool, bool)`

GetEndpointReturnsListOk returns a tuple with the EndpointReturnsList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointReturnsList

`func (o *OasServiceDiscovery) SetEndpointReturnsList(v bool)`

SetEndpointReturnsList sets EndpointReturnsList field to given value.

### HasEndpointReturnsList

`func (o *OasServiceDiscovery) HasEndpointReturnsList() bool`

HasEndpointReturnsList returns a boolean if a field has been set.

### GetParentDataPath

`func (o *OasServiceDiscovery) GetParentDataPath() string`

GetParentDataPath returns the ParentDataPath field if non-nil, zero value otherwise.

### GetParentDataPathOk

`func (o *OasServiceDiscovery) GetParentDataPathOk() (*string, bool)`

GetParentDataPathOk returns a tuple with the ParentDataPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentDataPath

`func (o *OasServiceDiscovery) SetParentDataPath(v string)`

SetParentDataPath sets ParentDataPath field to given value.

### HasParentDataPath

`func (o *OasServiceDiscovery) HasParentDataPath() bool`

HasParentDataPath returns a boolean if a field has been set.

### GetPortDataPath

`func (o *OasServiceDiscovery) GetPortDataPath() string`

GetPortDataPath returns the PortDataPath field if non-nil, zero value otherwise.

### GetPortDataPathOk

`func (o *OasServiceDiscovery) GetPortDataPathOk() (*string, bool)`

GetPortDataPathOk returns a tuple with the PortDataPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortDataPath

`func (o *OasServiceDiscovery) SetPortDataPath(v string)`

SetPortDataPath sets PortDataPath field to given value.

### HasPortDataPath

`func (o *OasServiceDiscovery) HasPortDataPath() bool`

HasPortDataPath returns a boolean if a field has been set.

### GetQueryEndpoint

`func (o *OasServiceDiscovery) GetQueryEndpoint() string`

GetQueryEndpoint returns the QueryEndpoint field if non-nil, zero value otherwise.

### GetQueryEndpointOk

`func (o *OasServiceDiscovery) GetQueryEndpointOk() (*string, bool)`

GetQueryEndpointOk returns a tuple with the QueryEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryEndpoint

`func (o *OasServiceDiscovery) SetQueryEndpoint(v string)`

SetQueryEndpoint sets QueryEndpoint field to given value.

### HasQueryEndpoint

`func (o *OasServiceDiscovery) HasQueryEndpoint() bool`

HasQueryEndpoint returns a boolean if a field has been set.

### GetTargetPath

`func (o *OasServiceDiscovery) GetTargetPath() string`

GetTargetPath returns the TargetPath field if non-nil, zero value otherwise.

### GetTargetPathOk

`func (o *OasServiceDiscovery) GetTargetPathOk() (*string, bool)`

GetTargetPathOk returns a tuple with the TargetPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetPath

`func (o *OasServiceDiscovery) SetTargetPath(v string)`

SetTargetPath sets TargetPath field to given value.

### HasTargetPath

`func (o *OasServiceDiscovery) HasTargetPath() bool`

HasTargetPath returns a boolean if a field has been set.

### GetUseNestedQuery

`func (o *OasServiceDiscovery) GetUseNestedQuery() bool`

GetUseNestedQuery returns the UseNestedQuery field if non-nil, zero value otherwise.

### GetUseNestedQueryOk

`func (o *OasServiceDiscovery) GetUseNestedQueryOk() (*bool, bool)`

GetUseNestedQueryOk returns a tuple with the UseNestedQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseNestedQuery

`func (o *OasServiceDiscovery) SetUseNestedQuery(v bool)`

SetUseNestedQuery sets UseNestedQuery field to given value.

### HasUseNestedQuery

`func (o *OasServiceDiscovery) HasUseNestedQuery() bool`

HasUseNestedQuery returns a boolean if a field has been set.

### GetUseTargetList

`func (o *OasServiceDiscovery) GetUseTargetList() bool`

GetUseTargetList returns the UseTargetList field if non-nil, zero value otherwise.

### GetUseTargetListOk

`func (o *OasServiceDiscovery) GetUseTargetListOk() (*bool, bool)`

GetUseTargetListOk returns a tuple with the UseTargetList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseTargetList

`func (o *OasServiceDiscovery) SetUseTargetList(v bool)`

SetUseTargetList sets UseTargetList field to given value.

### HasUseTargetList

`func (o *OasServiceDiscovery) HasUseTargetList() bool`

HasUseTargetList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


