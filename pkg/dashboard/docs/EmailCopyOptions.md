# EmailCopyOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**KeyEmail** | Pointer to [**EmailConfigMeta**](EmailConfigMeta.md) |  | [optional] 
**ResetPasswordEmail** | Pointer to [**EmailConfigMeta**](EmailConfigMeta.md) |  | [optional] 
**WelcomeEmail** | Pointer to [**EmailConfigMeta**](EmailConfigMeta.md) |  | [optional] 

## Methods

### NewEmailCopyOptions

`func NewEmailCopyOptions() *EmailCopyOptions`

NewEmailCopyOptions instantiates a new EmailCopyOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailCopyOptionsWithDefaults

`func NewEmailCopyOptionsWithDefaults() *EmailCopyOptions`

NewEmailCopyOptionsWithDefaults instantiates a new EmailCopyOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKeyEmail

`func (o *EmailCopyOptions) GetKeyEmail() EmailConfigMeta`

GetKeyEmail returns the KeyEmail field if non-nil, zero value otherwise.

### GetKeyEmailOk

`func (o *EmailCopyOptions) GetKeyEmailOk() (*EmailConfigMeta, bool)`

GetKeyEmailOk returns a tuple with the KeyEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyEmail

`func (o *EmailCopyOptions) SetKeyEmail(v EmailConfigMeta)`

SetKeyEmail sets KeyEmail field to given value.

### HasKeyEmail

`func (o *EmailCopyOptions) HasKeyEmail() bool`

HasKeyEmail returns a boolean if a field has been set.

### GetResetPasswordEmail

`func (o *EmailCopyOptions) GetResetPasswordEmail() EmailConfigMeta`

GetResetPasswordEmail returns the ResetPasswordEmail field if non-nil, zero value otherwise.

### GetResetPasswordEmailOk

`func (o *EmailCopyOptions) GetResetPasswordEmailOk() (*EmailConfigMeta, bool)`

GetResetPasswordEmailOk returns a tuple with the ResetPasswordEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResetPasswordEmail

`func (o *EmailCopyOptions) SetResetPasswordEmail(v EmailConfigMeta)`

SetResetPasswordEmail sets ResetPasswordEmail field to given value.

### HasResetPasswordEmail

`func (o *EmailCopyOptions) HasResetPasswordEmail() bool`

HasResetPasswordEmail returns a boolean if a field has been set.

### GetWelcomeEmail

`func (o *EmailCopyOptions) GetWelcomeEmail() EmailConfigMeta`

GetWelcomeEmail returns the WelcomeEmail field if non-nil, zero value otherwise.

### GetWelcomeEmailOk

`func (o *EmailCopyOptions) GetWelcomeEmailOk() (*EmailConfigMeta, bool)`

GetWelcomeEmailOk returns a tuple with the WelcomeEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWelcomeEmail

`func (o *EmailCopyOptions) SetWelcomeEmail(v EmailConfigMeta)`

SetWelcomeEmail sets WelcomeEmail field to given value.

### HasWelcomeEmail

`func (o *EmailCopyOptions) HasWelcomeEmail() bool`

HasWelcomeEmail returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


