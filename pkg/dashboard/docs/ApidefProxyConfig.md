# ApidefProxyConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CheckHostAgainstUptimeTests** | Pointer to **bool** |  | [optional] 
**DisableStripSlash** | Pointer to **bool** |  | [optional] 
**EnableLoadBalancing** | Pointer to **bool** |  | [optional] 
**ListenPath** | Pointer to **string** |  | [optional] 
**PreserveHostHeader** | Pointer to **bool** |  | [optional] 
**ServiceDiscovery** | Pointer to [**ApidefServiceDiscoveryConfiguration**](ApidefServiceDiscoveryConfiguration.md) |  | [optional] 
**StripListenPath** | Pointer to **bool** |  | [optional] 
**TargetList** | Pointer to **[]string** |  | [optional] 
**TargetUrl** | Pointer to **string** |  | [optional] 
**Transport** | Pointer to [**ApidefProxyConfigTransport**](ApidefProxyConfigTransport.md) |  | [optional] 

## Methods

### NewApidefProxyConfig

`func NewApidefProxyConfig() *ApidefProxyConfig`

NewApidefProxyConfig instantiates a new ApidefProxyConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApidefProxyConfigWithDefaults

`func NewApidefProxyConfigWithDefaults() *ApidefProxyConfig`

NewApidefProxyConfigWithDefaults instantiates a new ApidefProxyConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCheckHostAgainstUptimeTests

`func (o *ApidefProxyConfig) GetCheckHostAgainstUptimeTests() bool`

GetCheckHostAgainstUptimeTests returns the CheckHostAgainstUptimeTests field if non-nil, zero value otherwise.

### GetCheckHostAgainstUptimeTestsOk

`func (o *ApidefProxyConfig) GetCheckHostAgainstUptimeTestsOk() (*bool, bool)`

GetCheckHostAgainstUptimeTestsOk returns a tuple with the CheckHostAgainstUptimeTests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckHostAgainstUptimeTests

`func (o *ApidefProxyConfig) SetCheckHostAgainstUptimeTests(v bool)`

SetCheckHostAgainstUptimeTests sets CheckHostAgainstUptimeTests field to given value.

### HasCheckHostAgainstUptimeTests

`func (o *ApidefProxyConfig) HasCheckHostAgainstUptimeTests() bool`

HasCheckHostAgainstUptimeTests returns a boolean if a field has been set.

### GetDisableStripSlash

`func (o *ApidefProxyConfig) GetDisableStripSlash() bool`

GetDisableStripSlash returns the DisableStripSlash field if non-nil, zero value otherwise.

### GetDisableStripSlashOk

`func (o *ApidefProxyConfig) GetDisableStripSlashOk() (*bool, bool)`

GetDisableStripSlashOk returns a tuple with the DisableStripSlash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableStripSlash

`func (o *ApidefProxyConfig) SetDisableStripSlash(v bool)`

SetDisableStripSlash sets DisableStripSlash field to given value.

### HasDisableStripSlash

`func (o *ApidefProxyConfig) HasDisableStripSlash() bool`

HasDisableStripSlash returns a boolean if a field has been set.

### GetEnableLoadBalancing

`func (o *ApidefProxyConfig) GetEnableLoadBalancing() bool`

GetEnableLoadBalancing returns the EnableLoadBalancing field if non-nil, zero value otherwise.

### GetEnableLoadBalancingOk

`func (o *ApidefProxyConfig) GetEnableLoadBalancingOk() (*bool, bool)`

GetEnableLoadBalancingOk returns a tuple with the EnableLoadBalancing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableLoadBalancing

`func (o *ApidefProxyConfig) SetEnableLoadBalancing(v bool)`

SetEnableLoadBalancing sets EnableLoadBalancing field to given value.

### HasEnableLoadBalancing

`func (o *ApidefProxyConfig) HasEnableLoadBalancing() bool`

HasEnableLoadBalancing returns a boolean if a field has been set.

### GetListenPath

`func (o *ApidefProxyConfig) GetListenPath() string`

GetListenPath returns the ListenPath field if non-nil, zero value otherwise.

### GetListenPathOk

`func (o *ApidefProxyConfig) GetListenPathOk() (*string, bool)`

GetListenPathOk returns a tuple with the ListenPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListenPath

`func (o *ApidefProxyConfig) SetListenPath(v string)`

SetListenPath sets ListenPath field to given value.

### HasListenPath

`func (o *ApidefProxyConfig) HasListenPath() bool`

HasListenPath returns a boolean if a field has been set.

### GetPreserveHostHeader

`func (o *ApidefProxyConfig) GetPreserveHostHeader() bool`

GetPreserveHostHeader returns the PreserveHostHeader field if non-nil, zero value otherwise.

### GetPreserveHostHeaderOk

`func (o *ApidefProxyConfig) GetPreserveHostHeaderOk() (*bool, bool)`

GetPreserveHostHeaderOk returns a tuple with the PreserveHostHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreserveHostHeader

`func (o *ApidefProxyConfig) SetPreserveHostHeader(v bool)`

SetPreserveHostHeader sets PreserveHostHeader field to given value.

### HasPreserveHostHeader

`func (o *ApidefProxyConfig) HasPreserveHostHeader() bool`

HasPreserveHostHeader returns a boolean if a field has been set.

### GetServiceDiscovery

`func (o *ApidefProxyConfig) GetServiceDiscovery() ApidefServiceDiscoveryConfiguration`

GetServiceDiscovery returns the ServiceDiscovery field if non-nil, zero value otherwise.

### GetServiceDiscoveryOk

`func (o *ApidefProxyConfig) GetServiceDiscoveryOk() (*ApidefServiceDiscoveryConfiguration, bool)`

GetServiceDiscoveryOk returns a tuple with the ServiceDiscovery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceDiscovery

`func (o *ApidefProxyConfig) SetServiceDiscovery(v ApidefServiceDiscoveryConfiguration)`

SetServiceDiscovery sets ServiceDiscovery field to given value.

### HasServiceDiscovery

`func (o *ApidefProxyConfig) HasServiceDiscovery() bool`

HasServiceDiscovery returns a boolean if a field has been set.

### GetStripListenPath

`func (o *ApidefProxyConfig) GetStripListenPath() bool`

GetStripListenPath returns the StripListenPath field if non-nil, zero value otherwise.

### GetStripListenPathOk

`func (o *ApidefProxyConfig) GetStripListenPathOk() (*bool, bool)`

GetStripListenPathOk returns a tuple with the StripListenPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStripListenPath

`func (o *ApidefProxyConfig) SetStripListenPath(v bool)`

SetStripListenPath sets StripListenPath field to given value.

### HasStripListenPath

`func (o *ApidefProxyConfig) HasStripListenPath() bool`

HasStripListenPath returns a boolean if a field has been set.

### GetTargetList

`func (o *ApidefProxyConfig) GetTargetList() []string`

GetTargetList returns the TargetList field if non-nil, zero value otherwise.

### GetTargetListOk

`func (o *ApidefProxyConfig) GetTargetListOk() (*[]string, bool)`

GetTargetListOk returns a tuple with the TargetList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetList

`func (o *ApidefProxyConfig) SetTargetList(v []string)`

SetTargetList sets TargetList field to given value.

### HasTargetList

`func (o *ApidefProxyConfig) HasTargetList() bool`

HasTargetList returns a boolean if a field has been set.

### SetTargetListNil

`func (o *ApidefProxyConfig) SetTargetListNil(b bool)`

 SetTargetListNil sets the value for TargetList to be an explicit nil

### UnsetTargetList
`func (o *ApidefProxyConfig) UnsetTargetList()`

UnsetTargetList ensures that no value is present for TargetList, not even an explicit nil
### GetTargetUrl

`func (o *ApidefProxyConfig) GetTargetUrl() string`

GetTargetUrl returns the TargetUrl field if non-nil, zero value otherwise.

### GetTargetUrlOk

`func (o *ApidefProxyConfig) GetTargetUrlOk() (*string, bool)`

GetTargetUrlOk returns a tuple with the TargetUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetUrl

`func (o *ApidefProxyConfig) SetTargetUrl(v string)`

SetTargetUrl sets TargetUrl field to given value.

### HasTargetUrl

`func (o *ApidefProxyConfig) HasTargetUrl() bool`

HasTargetUrl returns a boolean if a field has been set.

### GetTransport

`func (o *ApidefProxyConfig) GetTransport() ApidefProxyConfigTransport`

GetTransport returns the Transport field if non-nil, zero value otherwise.

### GetTransportOk

`func (o *ApidefProxyConfig) GetTransportOk() (*ApidefProxyConfigTransport, bool)`

GetTransportOk returns a tuple with the Transport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransport

`func (o *ApidefProxyConfig) SetTransport(v ApidefProxyConfigTransport)`

SetTransport sets Transport field to given value.

### HasTransport

`func (o *ApidefProxyConfig) HasTransport() bool`

HasTransport returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


