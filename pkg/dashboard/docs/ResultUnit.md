# ResultUnit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | Pointer to **int32** |  | [optional] 
**Hits** | Pointer to **int32** |  | [optional] 
**Id** | Pointer to [**ResultId**](ResultId.md) |  | [optional] 
**LastHit** | Pointer to **time.Time** |  | [optional] 
**Latency** | Pointer to **float32** |  | [optional] 
**MaxLatency** | Pointer to **int32** |  | [optional] 
**MaxUpstreamLatency** | Pointer to **int32** |  | [optional] 
**MinLatency** | Pointer to **int32** |  | [optional] 
**MinUpstreamLatency** | Pointer to **int32** |  | [optional] 
**RequestTime** | Pointer to **float32** |  | [optional] 
**Success** | Pointer to **int32** |  | [optional] 
**UpstreamLatency** | Pointer to **float32** |  | [optional] 

## Methods

### NewResultUnit

`func NewResultUnit() *ResultUnit`

NewResultUnit instantiates a new ResultUnit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResultUnitWithDefaults

`func NewResultUnitWithDefaults() *ResultUnit`

NewResultUnitWithDefaults instantiates a new ResultUnit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *ResultUnit) GetError() int32`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ResultUnit) GetErrorOk() (*int32, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ResultUnit) SetError(v int32)`

SetError sets Error field to given value.

### HasError

`func (o *ResultUnit) HasError() bool`

HasError returns a boolean if a field has been set.

### GetHits

`func (o *ResultUnit) GetHits() int32`

GetHits returns the Hits field if non-nil, zero value otherwise.

### GetHitsOk

`func (o *ResultUnit) GetHitsOk() (*int32, bool)`

GetHitsOk returns a tuple with the Hits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHits

`func (o *ResultUnit) SetHits(v int32)`

SetHits sets Hits field to given value.

### HasHits

`func (o *ResultUnit) HasHits() bool`

HasHits returns a boolean if a field has been set.

### GetId

`func (o *ResultUnit) GetId() ResultId`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ResultUnit) GetIdOk() (*ResultId, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ResultUnit) SetId(v ResultId)`

SetId sets Id field to given value.

### HasId

`func (o *ResultUnit) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastHit

`func (o *ResultUnit) GetLastHit() time.Time`

GetLastHit returns the LastHit field if non-nil, zero value otherwise.

### GetLastHitOk

`func (o *ResultUnit) GetLastHitOk() (*time.Time, bool)`

GetLastHitOk returns a tuple with the LastHit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastHit

`func (o *ResultUnit) SetLastHit(v time.Time)`

SetLastHit sets LastHit field to given value.

### HasLastHit

`func (o *ResultUnit) HasLastHit() bool`

HasLastHit returns a boolean if a field has been set.

### GetLatency

`func (o *ResultUnit) GetLatency() float32`

GetLatency returns the Latency field if non-nil, zero value otherwise.

### GetLatencyOk

`func (o *ResultUnit) GetLatencyOk() (*float32, bool)`

GetLatencyOk returns a tuple with the Latency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatency

`func (o *ResultUnit) SetLatency(v float32)`

SetLatency sets Latency field to given value.

### HasLatency

`func (o *ResultUnit) HasLatency() bool`

HasLatency returns a boolean if a field has been set.

### GetMaxLatency

`func (o *ResultUnit) GetMaxLatency() int32`

GetMaxLatency returns the MaxLatency field if non-nil, zero value otherwise.

### GetMaxLatencyOk

`func (o *ResultUnit) GetMaxLatencyOk() (*int32, bool)`

GetMaxLatencyOk returns a tuple with the MaxLatency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxLatency

`func (o *ResultUnit) SetMaxLatency(v int32)`

SetMaxLatency sets MaxLatency field to given value.

### HasMaxLatency

`func (o *ResultUnit) HasMaxLatency() bool`

HasMaxLatency returns a boolean if a field has been set.

### GetMaxUpstreamLatency

`func (o *ResultUnit) GetMaxUpstreamLatency() int32`

GetMaxUpstreamLatency returns the MaxUpstreamLatency field if non-nil, zero value otherwise.

### GetMaxUpstreamLatencyOk

`func (o *ResultUnit) GetMaxUpstreamLatencyOk() (*int32, bool)`

GetMaxUpstreamLatencyOk returns a tuple with the MaxUpstreamLatency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxUpstreamLatency

`func (o *ResultUnit) SetMaxUpstreamLatency(v int32)`

SetMaxUpstreamLatency sets MaxUpstreamLatency field to given value.

### HasMaxUpstreamLatency

`func (o *ResultUnit) HasMaxUpstreamLatency() bool`

HasMaxUpstreamLatency returns a boolean if a field has been set.

### GetMinLatency

`func (o *ResultUnit) GetMinLatency() int32`

GetMinLatency returns the MinLatency field if non-nil, zero value otherwise.

### GetMinLatencyOk

`func (o *ResultUnit) GetMinLatencyOk() (*int32, bool)`

GetMinLatencyOk returns a tuple with the MinLatency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinLatency

`func (o *ResultUnit) SetMinLatency(v int32)`

SetMinLatency sets MinLatency field to given value.

### HasMinLatency

`func (o *ResultUnit) HasMinLatency() bool`

HasMinLatency returns a boolean if a field has been set.

### GetMinUpstreamLatency

`func (o *ResultUnit) GetMinUpstreamLatency() int32`

GetMinUpstreamLatency returns the MinUpstreamLatency field if non-nil, zero value otherwise.

### GetMinUpstreamLatencyOk

`func (o *ResultUnit) GetMinUpstreamLatencyOk() (*int32, bool)`

GetMinUpstreamLatencyOk returns a tuple with the MinUpstreamLatency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinUpstreamLatency

`func (o *ResultUnit) SetMinUpstreamLatency(v int32)`

SetMinUpstreamLatency sets MinUpstreamLatency field to given value.

### HasMinUpstreamLatency

`func (o *ResultUnit) HasMinUpstreamLatency() bool`

HasMinUpstreamLatency returns a boolean if a field has been set.

### GetRequestTime

`func (o *ResultUnit) GetRequestTime() float32`

GetRequestTime returns the RequestTime field if non-nil, zero value otherwise.

### GetRequestTimeOk

`func (o *ResultUnit) GetRequestTimeOk() (*float32, bool)`

GetRequestTimeOk returns a tuple with the RequestTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestTime

`func (o *ResultUnit) SetRequestTime(v float32)`

SetRequestTime sets RequestTime field to given value.

### HasRequestTime

`func (o *ResultUnit) HasRequestTime() bool`

HasRequestTime returns a boolean if a field has been set.

### GetSuccess

`func (o *ResultUnit) GetSuccess() int32`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *ResultUnit) GetSuccessOk() (*int32, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *ResultUnit) SetSuccess(v int32)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *ResultUnit) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetUpstreamLatency

`func (o *ResultUnit) GetUpstreamLatency() float32`

GetUpstreamLatency returns the UpstreamLatency field if non-nil, zero value otherwise.

### GetUpstreamLatencyOk

`func (o *ResultUnit) GetUpstreamLatencyOk() (*float32, bool)`

GetUpstreamLatencyOk returns a tuple with the UpstreamLatency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpstreamLatency

`func (o *ResultUnit) SetUpstreamLatency(v float32)`

SetUpstreamLatency sets UpstreamLatency field to given value.

### HasUpstreamLatency

`func (o *ResultUnit) HasUpstreamLatency() bool`

HasUpstreamLatency returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


