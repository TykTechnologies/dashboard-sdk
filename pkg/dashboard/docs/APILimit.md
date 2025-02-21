# APILimit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxQueryDepth** | Pointer to **int32** |  | [optional] 
**Per** | Pointer to **float64** |  | [optional] 
**QuotaMax** | Pointer to **int64** |  | [optional] 
**QuotaRemaining** | Pointer to **int64** |  | [optional] 
**QuotaRenewalRate** | Pointer to **int64** |  | [optional] 
**QuotaRenews** | Pointer to **int64** |  | [optional] 
**Rate** | Pointer to **float64** |  | [optional] 
**SetByPolicy** | Pointer to **bool** |  | [optional] 
**Smoothing** | Pointer to [**NullableRateLimitSmoothing**](RateLimitSmoothing.md) |  | [optional] 
**ThrottleInterval** | Pointer to **float64** |  | [optional] 
**ThrottleRetryLimit** | Pointer to **int32** |  | [optional] 

## Methods

### NewAPILimit

`func NewAPILimit() *APILimit`

NewAPILimit instantiates a new APILimit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAPILimitWithDefaults

`func NewAPILimitWithDefaults() *APILimit`

NewAPILimitWithDefaults instantiates a new APILimit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaxQueryDepth

`func (o *APILimit) GetMaxQueryDepth() int32`

GetMaxQueryDepth returns the MaxQueryDepth field if non-nil, zero value otherwise.

### GetMaxQueryDepthOk

`func (o *APILimit) GetMaxQueryDepthOk() (*int32, bool)`

GetMaxQueryDepthOk returns a tuple with the MaxQueryDepth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxQueryDepth

`func (o *APILimit) SetMaxQueryDepth(v int32)`

SetMaxQueryDepth sets MaxQueryDepth field to given value.

### HasMaxQueryDepth

`func (o *APILimit) HasMaxQueryDepth() bool`

HasMaxQueryDepth returns a boolean if a field has been set.

### GetPer

`func (o *APILimit) GetPer() float64`

GetPer returns the Per field if non-nil, zero value otherwise.

### GetPerOk

`func (o *APILimit) GetPerOk() (*float64, bool)`

GetPerOk returns a tuple with the Per field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPer

`func (o *APILimit) SetPer(v float64)`

SetPer sets Per field to given value.

### HasPer

`func (o *APILimit) HasPer() bool`

HasPer returns a boolean if a field has been set.

### GetQuotaMax

`func (o *APILimit) GetQuotaMax() int64`

GetQuotaMax returns the QuotaMax field if non-nil, zero value otherwise.

### GetQuotaMaxOk

`func (o *APILimit) GetQuotaMaxOk() (*int64, bool)`

GetQuotaMaxOk returns a tuple with the QuotaMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotaMax

`func (o *APILimit) SetQuotaMax(v int64)`

SetQuotaMax sets QuotaMax field to given value.

### HasQuotaMax

`func (o *APILimit) HasQuotaMax() bool`

HasQuotaMax returns a boolean if a field has been set.

### GetQuotaRemaining

`func (o *APILimit) GetQuotaRemaining() int64`

GetQuotaRemaining returns the QuotaRemaining field if non-nil, zero value otherwise.

### GetQuotaRemainingOk

`func (o *APILimit) GetQuotaRemainingOk() (*int64, bool)`

GetQuotaRemainingOk returns a tuple with the QuotaRemaining field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotaRemaining

`func (o *APILimit) SetQuotaRemaining(v int64)`

SetQuotaRemaining sets QuotaRemaining field to given value.

### HasQuotaRemaining

`func (o *APILimit) HasQuotaRemaining() bool`

HasQuotaRemaining returns a boolean if a field has been set.

### GetQuotaRenewalRate

`func (o *APILimit) GetQuotaRenewalRate() int64`

GetQuotaRenewalRate returns the QuotaRenewalRate field if non-nil, zero value otherwise.

### GetQuotaRenewalRateOk

`func (o *APILimit) GetQuotaRenewalRateOk() (*int64, bool)`

GetQuotaRenewalRateOk returns a tuple with the QuotaRenewalRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotaRenewalRate

`func (o *APILimit) SetQuotaRenewalRate(v int64)`

SetQuotaRenewalRate sets QuotaRenewalRate field to given value.

### HasQuotaRenewalRate

`func (o *APILimit) HasQuotaRenewalRate() bool`

HasQuotaRenewalRate returns a boolean if a field has been set.

### GetQuotaRenews

`func (o *APILimit) GetQuotaRenews() int64`

GetQuotaRenews returns the QuotaRenews field if non-nil, zero value otherwise.

### GetQuotaRenewsOk

`func (o *APILimit) GetQuotaRenewsOk() (*int64, bool)`

GetQuotaRenewsOk returns a tuple with the QuotaRenews field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotaRenews

`func (o *APILimit) SetQuotaRenews(v int64)`

SetQuotaRenews sets QuotaRenews field to given value.

### HasQuotaRenews

`func (o *APILimit) HasQuotaRenews() bool`

HasQuotaRenews returns a boolean if a field has been set.

### GetRate

`func (o *APILimit) GetRate() float64`

GetRate returns the Rate field if non-nil, zero value otherwise.

### GetRateOk

`func (o *APILimit) GetRateOk() (*float64, bool)`

GetRateOk returns a tuple with the Rate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRate

`func (o *APILimit) SetRate(v float64)`

SetRate sets Rate field to given value.

### HasRate

`func (o *APILimit) HasRate() bool`

HasRate returns a boolean if a field has been set.

### GetSetByPolicy

`func (o *APILimit) GetSetByPolicy() bool`

GetSetByPolicy returns the SetByPolicy field if non-nil, zero value otherwise.

### GetSetByPolicyOk

`func (o *APILimit) GetSetByPolicyOk() (*bool, bool)`

GetSetByPolicyOk returns a tuple with the SetByPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetByPolicy

`func (o *APILimit) SetSetByPolicy(v bool)`

SetSetByPolicy sets SetByPolicy field to given value.

### HasSetByPolicy

`func (o *APILimit) HasSetByPolicy() bool`

HasSetByPolicy returns a boolean if a field has been set.

### GetSmoothing

`func (o *APILimit) GetSmoothing() RateLimitSmoothing`

GetSmoothing returns the Smoothing field if non-nil, zero value otherwise.

### GetSmoothingOk

`func (o *APILimit) GetSmoothingOk() (*RateLimitSmoothing, bool)`

GetSmoothingOk returns a tuple with the Smoothing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmoothing

`func (o *APILimit) SetSmoothing(v RateLimitSmoothing)`

SetSmoothing sets Smoothing field to given value.

### HasSmoothing

`func (o *APILimit) HasSmoothing() bool`

HasSmoothing returns a boolean if a field has been set.

### SetSmoothingNil

`func (o *APILimit) SetSmoothingNil(b bool)`

 SetSmoothingNil sets the value for Smoothing to be an explicit nil

### UnsetSmoothing
`func (o *APILimit) UnsetSmoothing()`

UnsetSmoothing ensures that no value is present for Smoothing, not even an explicit nil
### GetThrottleInterval

`func (o *APILimit) GetThrottleInterval() float64`

GetThrottleInterval returns the ThrottleInterval field if non-nil, zero value otherwise.

### GetThrottleIntervalOk

`func (o *APILimit) GetThrottleIntervalOk() (*float64, bool)`

GetThrottleIntervalOk returns a tuple with the ThrottleInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThrottleInterval

`func (o *APILimit) SetThrottleInterval(v float64)`

SetThrottleInterval sets ThrottleInterval field to given value.

### HasThrottleInterval

`func (o *APILimit) HasThrottleInterval() bool`

HasThrottleInterval returns a boolean if a field has been set.

### GetThrottleRetryLimit

`func (o *APILimit) GetThrottleRetryLimit() int32`

GetThrottleRetryLimit returns the ThrottleRetryLimit field if non-nil, zero value otherwise.

### GetThrottleRetryLimitOk

`func (o *APILimit) GetThrottleRetryLimitOk() (*int32, bool)`

GetThrottleRetryLimitOk returns a tuple with the ThrottleRetryLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThrottleRetryLimit

`func (o *APILimit) SetThrottleRetryLimit(v int32)`

SetThrottleRetryLimit sets ThrottleRetryLimit field to given value.

### HasThrottleRetryLimit

`func (o *APILimit) HasThrottleRetryLimit() bool`

HasThrottleRetryLimit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


