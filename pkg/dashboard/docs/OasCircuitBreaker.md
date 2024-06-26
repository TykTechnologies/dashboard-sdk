# OasCircuitBreaker

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CoolDownPeriod** | Pointer to **int32** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**HalfOpenStateEnabled** | Pointer to **bool** |  | [optional] 
**SampleSize** | Pointer to **int32** |  | [optional] 
**Threshold** | Pointer to **float32** |  | [optional] 

## Methods

### NewOasCircuitBreaker

`func NewOasCircuitBreaker() *OasCircuitBreaker`

NewOasCircuitBreaker instantiates a new OasCircuitBreaker object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOasCircuitBreakerWithDefaults

`func NewOasCircuitBreakerWithDefaults() *OasCircuitBreaker`

NewOasCircuitBreakerWithDefaults instantiates a new OasCircuitBreaker object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCoolDownPeriod

`func (o *OasCircuitBreaker) GetCoolDownPeriod() int32`

GetCoolDownPeriod returns the CoolDownPeriod field if non-nil, zero value otherwise.

### GetCoolDownPeriodOk

`func (o *OasCircuitBreaker) GetCoolDownPeriodOk() (*int32, bool)`

GetCoolDownPeriodOk returns a tuple with the CoolDownPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoolDownPeriod

`func (o *OasCircuitBreaker) SetCoolDownPeriod(v int32)`

SetCoolDownPeriod sets CoolDownPeriod field to given value.

### HasCoolDownPeriod

`func (o *OasCircuitBreaker) HasCoolDownPeriod() bool`

HasCoolDownPeriod returns a boolean if a field has been set.

### GetEnabled

`func (o *OasCircuitBreaker) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *OasCircuitBreaker) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *OasCircuitBreaker) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *OasCircuitBreaker) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetHalfOpenStateEnabled

`func (o *OasCircuitBreaker) GetHalfOpenStateEnabled() bool`

GetHalfOpenStateEnabled returns the HalfOpenStateEnabled field if non-nil, zero value otherwise.

### GetHalfOpenStateEnabledOk

`func (o *OasCircuitBreaker) GetHalfOpenStateEnabledOk() (*bool, bool)`

GetHalfOpenStateEnabledOk returns a tuple with the HalfOpenStateEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHalfOpenStateEnabled

`func (o *OasCircuitBreaker) SetHalfOpenStateEnabled(v bool)`

SetHalfOpenStateEnabled sets HalfOpenStateEnabled field to given value.

### HasHalfOpenStateEnabled

`func (o *OasCircuitBreaker) HasHalfOpenStateEnabled() bool`

HasHalfOpenStateEnabled returns a boolean if a field has been set.

### GetSampleSize

`func (o *OasCircuitBreaker) GetSampleSize() int32`

GetSampleSize returns the SampleSize field if non-nil, zero value otherwise.

### GetSampleSizeOk

`func (o *OasCircuitBreaker) GetSampleSizeOk() (*int32, bool)`

GetSampleSizeOk returns a tuple with the SampleSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSampleSize

`func (o *OasCircuitBreaker) SetSampleSize(v int32)`

SetSampleSize sets SampleSize field to given value.

### HasSampleSize

`func (o *OasCircuitBreaker) HasSampleSize() bool`

HasSampleSize returns a boolean if a field has been set.

### GetThreshold

`func (o *OasCircuitBreaker) GetThreshold() float32`

GetThreshold returns the Threshold field if non-nil, zero value otherwise.

### GetThresholdOk

`func (o *OasCircuitBreaker) GetThresholdOk() (*float32, bool)`

GetThresholdOk returns a tuple with the Threshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreshold

`func (o *OasCircuitBreaker) SetThreshold(v float32)`

SetThreshold sets Threshold field to given value.

### HasThreshold

`func (o *OasCircuitBreaker) HasThreshold() bool`

HasThreshold returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


