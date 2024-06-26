# ApidefEndPointMeta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Disabled** | Pointer to **bool** |  | [optional] 
**IgnoreCase** | Pointer to **bool** |  | [optional] 
**Method** | Pointer to **string** |  | [optional] 
**MethodActions** | Pointer to [**map[string]ApidefEndpointMethodMeta**](ApidefEndpointMethodMeta.md) |  | [optional] 
**Path** | Pointer to **string** |  | [optional] 

## Methods

### NewApidefEndPointMeta

`func NewApidefEndPointMeta() *ApidefEndPointMeta`

NewApidefEndPointMeta instantiates a new ApidefEndPointMeta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApidefEndPointMetaWithDefaults

`func NewApidefEndPointMetaWithDefaults() *ApidefEndPointMeta`

NewApidefEndPointMetaWithDefaults instantiates a new ApidefEndPointMeta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisabled

`func (o *ApidefEndPointMeta) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *ApidefEndPointMeta) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *ApidefEndPointMeta) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *ApidefEndPointMeta) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetIgnoreCase

`func (o *ApidefEndPointMeta) GetIgnoreCase() bool`

GetIgnoreCase returns the IgnoreCase field if non-nil, zero value otherwise.

### GetIgnoreCaseOk

`func (o *ApidefEndPointMeta) GetIgnoreCaseOk() (*bool, bool)`

GetIgnoreCaseOk returns a tuple with the IgnoreCase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreCase

`func (o *ApidefEndPointMeta) SetIgnoreCase(v bool)`

SetIgnoreCase sets IgnoreCase field to given value.

### HasIgnoreCase

`func (o *ApidefEndPointMeta) HasIgnoreCase() bool`

HasIgnoreCase returns a boolean if a field has been set.

### GetMethod

`func (o *ApidefEndPointMeta) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *ApidefEndPointMeta) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *ApidefEndPointMeta) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *ApidefEndPointMeta) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetMethodActions

`func (o *ApidefEndPointMeta) GetMethodActions() map[string]ApidefEndpointMethodMeta`

GetMethodActions returns the MethodActions field if non-nil, zero value otherwise.

### GetMethodActionsOk

`func (o *ApidefEndPointMeta) GetMethodActionsOk() (*map[string]ApidefEndpointMethodMeta, bool)`

GetMethodActionsOk returns a tuple with the MethodActions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethodActions

`func (o *ApidefEndPointMeta) SetMethodActions(v map[string]ApidefEndpointMethodMeta)`

SetMethodActions sets MethodActions field to given value.

### HasMethodActions

`func (o *ApidefEndPointMeta) HasMethodActions() bool`

HasMethodActions returns a boolean if a field has been set.

### GetPath

`func (o *ApidefEndPointMeta) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *ApidefEndPointMeta) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *ApidefEndPointMeta) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *ApidefEndPointMeta) HasPath() bool`

HasPath returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


