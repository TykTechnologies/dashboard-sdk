# MigrateAPIResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | Pointer to [**MigrateAPIResponseSuccess**](MigrateAPIResponseSuccess.md) |  | [optional] 
**Failure** | Pointer to [**MigrateAPIResponseFailure**](MigrateAPIResponseFailure.md) |  | [optional] 
**Skipped** | Pointer to [**MigrateAPIResponseSkipped**](MigrateAPIResponseSkipped.md) |  | [optional] 
**AbortedOnFailure** | Pointer to **bool** | reports whether migration process aborted on first failure | [optional] 

## Methods

### NewMigrateAPIResponse

`func NewMigrateAPIResponse() *MigrateAPIResponse`

NewMigrateAPIResponse instantiates a new MigrateAPIResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMigrateAPIResponseWithDefaults

`func NewMigrateAPIResponseWithDefaults() *MigrateAPIResponse`

NewMigrateAPIResponseWithDefaults instantiates a new MigrateAPIResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *MigrateAPIResponse) GetSuccess() MigrateAPIResponseSuccess`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *MigrateAPIResponse) GetSuccessOk() (*MigrateAPIResponseSuccess, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *MigrateAPIResponse) SetSuccess(v MigrateAPIResponseSuccess)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *MigrateAPIResponse) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetFailure

`func (o *MigrateAPIResponse) GetFailure() MigrateAPIResponseFailure`

GetFailure returns the Failure field if non-nil, zero value otherwise.

### GetFailureOk

`func (o *MigrateAPIResponse) GetFailureOk() (*MigrateAPIResponseFailure, bool)`

GetFailureOk returns a tuple with the Failure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailure

`func (o *MigrateAPIResponse) SetFailure(v MigrateAPIResponseFailure)`

SetFailure sets Failure field to given value.

### HasFailure

`func (o *MigrateAPIResponse) HasFailure() bool`

HasFailure returns a boolean if a field has been set.

### GetSkipped

`func (o *MigrateAPIResponse) GetSkipped() MigrateAPIResponseSkipped`

GetSkipped returns the Skipped field if non-nil, zero value otherwise.

### GetSkippedOk

`func (o *MigrateAPIResponse) GetSkippedOk() (*MigrateAPIResponseSkipped, bool)`

GetSkippedOk returns a tuple with the Skipped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipped

`func (o *MigrateAPIResponse) SetSkipped(v MigrateAPIResponseSkipped)`

SetSkipped sets Skipped field to given value.

### HasSkipped

`func (o *MigrateAPIResponse) HasSkipped() bool`

HasSkipped returns a boolean if a field has been set.

### GetAbortedOnFailure

`func (o *MigrateAPIResponse) GetAbortedOnFailure() bool`

GetAbortedOnFailure returns the AbortedOnFailure field if non-nil, zero value otherwise.

### GetAbortedOnFailureOk

`func (o *MigrateAPIResponse) GetAbortedOnFailureOk() (*bool, bool)`

GetAbortedOnFailureOk returns a tuple with the AbortedOnFailure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbortedOnFailure

`func (o *MigrateAPIResponse) SetAbortedOnFailure(v bool)`

SetAbortedOnFailure sets AbortedOnFailure field to given value.

### HasAbortedOnFailure

`func (o *MigrateAPIResponse) HasAbortedOnFailure() bool`

HasAbortedOnFailure returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


