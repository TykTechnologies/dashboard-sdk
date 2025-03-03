# MigrateAPIResponseFailure

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int32** | count of APIs failed migration | [optional] 
**Errors** | Pointer to [**[]MigrateAPIResponseFailureErrorsInner**](MigrateAPIResponseFailureErrorsInner.md) | details of errors occurred during migration | [optional] 

## Methods

### NewMigrateAPIResponseFailure

`func NewMigrateAPIResponseFailure() *MigrateAPIResponseFailure`

NewMigrateAPIResponseFailure instantiates a new MigrateAPIResponseFailure object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMigrateAPIResponseFailureWithDefaults

`func NewMigrateAPIResponseFailureWithDefaults() *MigrateAPIResponseFailure`

NewMigrateAPIResponseFailureWithDefaults instantiates a new MigrateAPIResponseFailure object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *MigrateAPIResponseFailure) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *MigrateAPIResponseFailure) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *MigrateAPIResponseFailure) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *MigrateAPIResponseFailure) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetErrors

`func (o *MigrateAPIResponseFailure) GetErrors() []MigrateAPIResponseFailureErrorsInner`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *MigrateAPIResponseFailure) GetErrorsOk() (*[]MigrateAPIResponseFailureErrorsInner, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *MigrateAPIResponseFailure) SetErrors(v []MigrateAPIResponseFailureErrorsInner)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *MigrateAPIResponseFailure) HasErrors() bool`

HasErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


