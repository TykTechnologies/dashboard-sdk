# PortalConfigMailOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EmailCopy** | Pointer to [**EmailCopyOptions**](EmailCopyOptions.md) |  | [optional] 
**MailFromEmail** | Pointer to **string** |  | [optional] 
**MailFromName** | Pointer to **string** |  | [optional] 

## Methods

### NewPortalConfigMailOptions

`func NewPortalConfigMailOptions() *PortalConfigMailOptions`

NewPortalConfigMailOptions instantiates a new PortalConfigMailOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPortalConfigMailOptionsWithDefaults

`func NewPortalConfigMailOptionsWithDefaults() *PortalConfigMailOptions`

NewPortalConfigMailOptionsWithDefaults instantiates a new PortalConfigMailOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmailCopy

`func (o *PortalConfigMailOptions) GetEmailCopy() EmailCopyOptions`

GetEmailCopy returns the EmailCopy field if non-nil, zero value otherwise.

### GetEmailCopyOk

`func (o *PortalConfigMailOptions) GetEmailCopyOk() (*EmailCopyOptions, bool)`

GetEmailCopyOk returns a tuple with the EmailCopy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailCopy

`func (o *PortalConfigMailOptions) SetEmailCopy(v EmailCopyOptions)`

SetEmailCopy sets EmailCopy field to given value.

### HasEmailCopy

`func (o *PortalConfigMailOptions) HasEmailCopy() bool`

HasEmailCopy returns a boolean if a field has been set.

### GetMailFromEmail

`func (o *PortalConfigMailOptions) GetMailFromEmail() string`

GetMailFromEmail returns the MailFromEmail field if non-nil, zero value otherwise.

### GetMailFromEmailOk

`func (o *PortalConfigMailOptions) GetMailFromEmailOk() (*string, bool)`

GetMailFromEmailOk returns a tuple with the MailFromEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMailFromEmail

`func (o *PortalConfigMailOptions) SetMailFromEmail(v string)`

SetMailFromEmail sets MailFromEmail field to given value.

### HasMailFromEmail

`func (o *PortalConfigMailOptions) HasMailFromEmail() bool`

HasMailFromEmail returns a boolean if a field has been set.

### GetMailFromName

`func (o *PortalConfigMailOptions) GetMailFromName() string`

GetMailFromName returns the MailFromName field if non-nil, zero value otherwise.

### GetMailFromNameOk

`func (o *PortalConfigMailOptions) GetMailFromNameOk() (*string, bool)`

GetMailFromNameOk returns a tuple with the MailFromName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMailFromName

`func (o *PortalConfigMailOptions) SetMailFromName(v string)`

SetMailFromName sets MailFromName field to given value.

### HasMailFromName

`func (o *PortalConfigMailOptions) HasMailFromName() bool`

HasMailFromName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


