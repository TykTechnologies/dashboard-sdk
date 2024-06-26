# ApidefUptimeTestsConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExpireUtimeAfter** | Pointer to **int32** |  | [optional] 
**RecheckWait** | Pointer to **int32** |  | [optional] 
**ServiceDiscovery** | Pointer to [**ApidefServiceDiscoveryConfiguration**](ApidefServiceDiscoveryConfiguration.md) |  | [optional] 

## Methods

### NewApidefUptimeTestsConfig

`func NewApidefUptimeTestsConfig() *ApidefUptimeTestsConfig`

NewApidefUptimeTestsConfig instantiates a new ApidefUptimeTestsConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApidefUptimeTestsConfigWithDefaults

`func NewApidefUptimeTestsConfigWithDefaults() *ApidefUptimeTestsConfig`

NewApidefUptimeTestsConfigWithDefaults instantiates a new ApidefUptimeTestsConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpireUtimeAfter

`func (o *ApidefUptimeTestsConfig) GetExpireUtimeAfter() int32`

GetExpireUtimeAfter returns the ExpireUtimeAfter field if non-nil, zero value otherwise.

### GetExpireUtimeAfterOk

`func (o *ApidefUptimeTestsConfig) GetExpireUtimeAfterOk() (*int32, bool)`

GetExpireUtimeAfterOk returns a tuple with the ExpireUtimeAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpireUtimeAfter

`func (o *ApidefUptimeTestsConfig) SetExpireUtimeAfter(v int32)`

SetExpireUtimeAfter sets ExpireUtimeAfter field to given value.

### HasExpireUtimeAfter

`func (o *ApidefUptimeTestsConfig) HasExpireUtimeAfter() bool`

HasExpireUtimeAfter returns a boolean if a field has been set.

### GetRecheckWait

`func (o *ApidefUptimeTestsConfig) GetRecheckWait() int32`

GetRecheckWait returns the RecheckWait field if non-nil, zero value otherwise.

### GetRecheckWaitOk

`func (o *ApidefUptimeTestsConfig) GetRecheckWaitOk() (*int32, bool)`

GetRecheckWaitOk returns a tuple with the RecheckWait field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecheckWait

`func (o *ApidefUptimeTestsConfig) SetRecheckWait(v int32)`

SetRecheckWait sets RecheckWait field to given value.

### HasRecheckWait

`func (o *ApidefUptimeTestsConfig) HasRecheckWait() bool`

HasRecheckWait returns a boolean if a field has been set.

### GetServiceDiscovery

`func (o *ApidefUptimeTestsConfig) GetServiceDiscovery() ApidefServiceDiscoveryConfiguration`

GetServiceDiscovery returns the ServiceDiscovery field if non-nil, zero value otherwise.

### GetServiceDiscoveryOk

`func (o *ApidefUptimeTestsConfig) GetServiceDiscoveryOk() (*ApidefServiceDiscoveryConfiguration, bool)`

GetServiceDiscoveryOk returns a tuple with the ServiceDiscovery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceDiscovery

`func (o *ApidefUptimeTestsConfig) SetServiceDiscovery(v ApidefServiceDiscoveryConfiguration)`

SetServiceDiscovery sets ServiceDiscovery field to given value.

### HasServiceDiscovery

`func (o *ApidefUptimeTestsConfig) HasServiceDiscovery() bool`

HasServiceDiscovery returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


