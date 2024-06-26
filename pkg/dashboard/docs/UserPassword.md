# UserPassword

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiModel** | Pointer to **map[string]interface{}** |  | [optional] 
**CurrentPassword** | Pointer to **string** |  | [optional] 
**NewPassword** | Pointer to **string** |  | [optional] 

## Methods

### NewUserPassword

`func NewUserPassword() *UserPassword`

NewUserPassword instantiates a new UserPassword object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserPasswordWithDefaults

`func NewUserPasswordWithDefaults() *UserPassword`

NewUserPasswordWithDefaults instantiates a new UserPassword object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiModel

`func (o *UserPassword) GetApiModel() map[string]interface{}`

GetApiModel returns the ApiModel field if non-nil, zero value otherwise.

### GetApiModelOk

`func (o *UserPassword) GetApiModelOk() (*map[string]interface{}, bool)`

GetApiModelOk returns a tuple with the ApiModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiModel

`func (o *UserPassword) SetApiModel(v map[string]interface{})`

SetApiModel sets ApiModel field to given value.

### HasApiModel

`func (o *UserPassword) HasApiModel() bool`

HasApiModel returns a boolean if a field has been set.

### GetCurrentPassword

`func (o *UserPassword) GetCurrentPassword() string`

GetCurrentPassword returns the CurrentPassword field if non-nil, zero value otherwise.

### GetCurrentPasswordOk

`func (o *UserPassword) GetCurrentPasswordOk() (*string, bool)`

GetCurrentPasswordOk returns a tuple with the CurrentPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPassword

`func (o *UserPassword) SetCurrentPassword(v string)`

SetCurrentPassword sets CurrentPassword field to given value.

### HasCurrentPassword

`func (o *UserPassword) HasCurrentPassword() bool`

HasCurrentPassword returns a boolean if a field has been set.

### GetNewPassword

`func (o *UserPassword) GetNewPassword() string`

GetNewPassword returns the NewPassword field if non-nil, zero value otherwise.

### GetNewPasswordOk

`func (o *UserPassword) GetNewPasswordOk() (*string, bool)`

GetNewPasswordOk returns a tuple with the NewPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewPassword

`func (o *UserPassword) SetNewPassword(v string)`

SetNewPassword sets NewPassword field to given value.

### HasNewPassword

`func (o *UserPassword) HasNewPassword() bool`

HasNewPassword returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


