# DryRunRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Oas** | Pointer to [**Model30**](Model30.md) |  | [optional] 
**TykOas** | Pointer to [**CreateApiOASRequest**](CreateApiOASRequest.md) |  | [optional] 

## Methods

### NewDryRunRequest

`func NewDryRunRequest() *DryRunRequest`

NewDryRunRequest instantiates a new DryRunRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDryRunRequestWithDefaults

`func NewDryRunRequestWithDefaults() *DryRunRequest`

NewDryRunRequestWithDefaults instantiates a new DryRunRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOas

`func (o *DryRunRequest) GetOas() Model30`

GetOas returns the Oas field if non-nil, zero value otherwise.

### GetOasOk

`func (o *DryRunRequest) GetOasOk() (*Model30, bool)`

GetOasOk returns a tuple with the Oas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOas

`func (o *DryRunRequest) SetOas(v Model30)`

SetOas sets Oas field to given value.

### HasOas

`func (o *DryRunRequest) HasOas() bool`

HasOas returns a boolean if a field has been set.

### GetTykOas

`func (o *DryRunRequest) GetTykOas() CreateApiOASRequest`

GetTykOas returns the TykOas field if non-nil, zero value otherwise.

### GetTykOasOk

`func (o *DryRunRequest) GetTykOasOk() (*CreateApiOASRequest, bool)`

GetTykOasOk returns a tuple with the TykOas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTykOas

`func (o *DryRunRequest) SetTykOas(v CreateApiOASRequest)`

SetTykOas sets TykOas field to given value.

### HasTykOas

`func (o *DryRunRequest) HasTykOas() bool`

HasTykOas returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


