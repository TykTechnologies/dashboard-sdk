# OasTransformHeaders

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Add** | Pointer to [**[]OasHeader**](OasHeader.md) |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Remove** | Pointer to **[]string** |  | [optional] 

## Methods

### NewOasTransformHeaders

`func NewOasTransformHeaders() *OasTransformHeaders`

NewOasTransformHeaders instantiates a new OasTransformHeaders object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOasTransformHeadersWithDefaults

`func NewOasTransformHeadersWithDefaults() *OasTransformHeaders`

NewOasTransformHeadersWithDefaults instantiates a new OasTransformHeaders object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdd

`func (o *OasTransformHeaders) GetAdd() []OasHeader`

GetAdd returns the Add field if non-nil, zero value otherwise.

### GetAddOk

`func (o *OasTransformHeaders) GetAddOk() (*[]OasHeader, bool)`

GetAddOk returns a tuple with the Add field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdd

`func (o *OasTransformHeaders) SetAdd(v []OasHeader)`

SetAdd sets Add field to given value.

### HasAdd

`func (o *OasTransformHeaders) HasAdd() bool`

HasAdd returns a boolean if a field has been set.

### GetEnabled

`func (o *OasTransformHeaders) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *OasTransformHeaders) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *OasTransformHeaders) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *OasTransformHeaders) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetRemove

`func (o *OasTransformHeaders) GetRemove() []string`

GetRemove returns the Remove field if non-nil, zero value otherwise.

### GetRemoveOk

`func (o *OasTransformHeaders) GetRemoveOk() (*[]string, bool)`

GetRemoveOk returns a tuple with the Remove field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemove

`func (o *OasTransformHeaders) SetRemove(v []string)`

SetRemove sets Remove field to given value.

### HasRemove

`func (o *OasTransformHeaders) HasRemove() bool`

HasRemove returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


