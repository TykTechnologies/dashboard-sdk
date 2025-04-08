# SystemStatsResp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** |  | [optional] 
**Stats** | Pointer to [**NullableEntityStats**](EntityStats.md) |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 

## Methods

### NewSystemStatsResp

`func NewSystemStatsResp() *SystemStatsResp`

NewSystemStatsResp instantiates a new SystemStatsResp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSystemStatsRespWithDefaults

`func NewSystemStatsRespWithDefaults() *SystemStatsResp`

NewSystemStatsRespWithDefaults instantiates a new SystemStatsResp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *SystemStatsResp) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *SystemStatsResp) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *SystemStatsResp) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *SystemStatsResp) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetStats

`func (o *SystemStatsResp) GetStats() EntityStats`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *SystemStatsResp) GetStatsOk() (*EntityStats, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *SystemStatsResp) SetStats(v EntityStats)`

SetStats sets Stats field to given value.

### HasStats

`func (o *SystemStatsResp) HasStats() bool`

HasStats returns a boolean if a field has been set.

### SetStatsNil

`func (o *SystemStatsResp) SetStatsNil(b bool)`

 SetStatsNil sets the value for Stats to be an explicit nil

### UnsetStats
`func (o *SystemStatsResp) UnsetStats()`

UnsetStats ensures that no value is present for Stats, not even an explicit nil
### GetStatus

`func (o *SystemStatsResp) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SystemStatsResp) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SystemStatsResp) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SystemStatsResp) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


