# PortalConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HashKeys** | Pointer to **bool** |  | [optional] 
**CatalogueLoginOnly** | Pointer to **bool** |  | [optional] 
**DcrOptions** | Pointer to [**DCROptions**](DCROptions.md) |  | [optional] 
**DisableAutoLogin** | Pointer to **bool** |  | [optional] 
**DisableLogin** | Pointer to **bool** |  | [optional] 
**DisableSignup** | Pointer to **bool** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**EnableDcr** | Pointer to **bool** |  | [optional] 
**EnableMultiSelection** | Pointer to **bool** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**KeyRequestFields** | Pointer to **[]string** |  | [optional] 
**MailOptions** | Pointer to [**PortalConfigMailOptions**](PortalConfigMailOptions.md) |  | [optional] 
**OauthUsageLimit** | Pointer to **int32** |  | [optional] 
**OrgId** | Pointer to **string** |  | [optional] 
**Override** | Pointer to **bool** |  | [optional] 
**RedirectOnKeyRequest** | Pointer to **bool** |  | [optional] 
**RedirectTo** | Pointer to **string** |  | [optional] 
**RequireKeyApproval** | Pointer to **bool** |  | [optional] 
**SecureKeyApproval** | Pointer to **bool** |  | [optional] 
**SignupFields** | Pointer to **[]string** |  | [optional] 

## Methods

### NewPortalConfig

`func NewPortalConfig() *PortalConfig`

NewPortalConfig instantiates a new PortalConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPortalConfigWithDefaults

`func NewPortalConfigWithDefaults() *PortalConfig`

NewPortalConfigWithDefaults instantiates a new PortalConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHashKeys

`func (o *PortalConfig) GetHashKeys() bool`

GetHashKeys returns the HashKeys field if non-nil, zero value otherwise.

### GetHashKeysOk

`func (o *PortalConfig) GetHashKeysOk() (*bool, bool)`

GetHashKeysOk returns a tuple with the HashKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHashKeys

`func (o *PortalConfig) SetHashKeys(v bool)`

SetHashKeys sets HashKeys field to given value.

### HasHashKeys

`func (o *PortalConfig) HasHashKeys() bool`

HasHashKeys returns a boolean if a field has been set.

### GetCatalogueLoginOnly

`func (o *PortalConfig) GetCatalogueLoginOnly() bool`

GetCatalogueLoginOnly returns the CatalogueLoginOnly field if non-nil, zero value otherwise.

### GetCatalogueLoginOnlyOk

`func (o *PortalConfig) GetCatalogueLoginOnlyOk() (*bool, bool)`

GetCatalogueLoginOnlyOk returns a tuple with the CatalogueLoginOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalogueLoginOnly

`func (o *PortalConfig) SetCatalogueLoginOnly(v bool)`

SetCatalogueLoginOnly sets CatalogueLoginOnly field to given value.

### HasCatalogueLoginOnly

`func (o *PortalConfig) HasCatalogueLoginOnly() bool`

HasCatalogueLoginOnly returns a boolean if a field has been set.

### GetDcrOptions

`func (o *PortalConfig) GetDcrOptions() DCROptions`

GetDcrOptions returns the DcrOptions field if non-nil, zero value otherwise.

### GetDcrOptionsOk

`func (o *PortalConfig) GetDcrOptionsOk() (*DCROptions, bool)`

GetDcrOptionsOk returns a tuple with the DcrOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDcrOptions

`func (o *PortalConfig) SetDcrOptions(v DCROptions)`

SetDcrOptions sets DcrOptions field to given value.

### HasDcrOptions

`func (o *PortalConfig) HasDcrOptions() bool`

HasDcrOptions returns a boolean if a field has been set.

### GetDisableAutoLogin

`func (o *PortalConfig) GetDisableAutoLogin() bool`

GetDisableAutoLogin returns the DisableAutoLogin field if non-nil, zero value otherwise.

### GetDisableAutoLoginOk

`func (o *PortalConfig) GetDisableAutoLoginOk() (*bool, bool)`

GetDisableAutoLoginOk returns a tuple with the DisableAutoLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableAutoLogin

`func (o *PortalConfig) SetDisableAutoLogin(v bool)`

SetDisableAutoLogin sets DisableAutoLogin field to given value.

### HasDisableAutoLogin

`func (o *PortalConfig) HasDisableAutoLogin() bool`

HasDisableAutoLogin returns a boolean if a field has been set.

### GetDisableLogin

`func (o *PortalConfig) GetDisableLogin() bool`

GetDisableLogin returns the DisableLogin field if non-nil, zero value otherwise.

### GetDisableLoginOk

`func (o *PortalConfig) GetDisableLoginOk() (*bool, bool)`

GetDisableLoginOk returns a tuple with the DisableLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableLogin

`func (o *PortalConfig) SetDisableLogin(v bool)`

SetDisableLogin sets DisableLogin field to given value.

### HasDisableLogin

`func (o *PortalConfig) HasDisableLogin() bool`

HasDisableLogin returns a boolean if a field has been set.

### GetDisableSignup

`func (o *PortalConfig) GetDisableSignup() bool`

GetDisableSignup returns the DisableSignup field if non-nil, zero value otherwise.

### GetDisableSignupOk

`func (o *PortalConfig) GetDisableSignupOk() (*bool, bool)`

GetDisableSignupOk returns a tuple with the DisableSignup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableSignup

`func (o *PortalConfig) SetDisableSignup(v bool)`

SetDisableSignup sets DisableSignup field to given value.

### HasDisableSignup

`func (o *PortalConfig) HasDisableSignup() bool`

HasDisableSignup returns a boolean if a field has been set.

### GetEmail

`func (o *PortalConfig) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *PortalConfig) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *PortalConfig) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *PortalConfig) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetEnableDcr

`func (o *PortalConfig) GetEnableDcr() bool`

GetEnableDcr returns the EnableDcr field if non-nil, zero value otherwise.

### GetEnableDcrOk

`func (o *PortalConfig) GetEnableDcrOk() (*bool, bool)`

GetEnableDcrOk returns a tuple with the EnableDcr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableDcr

`func (o *PortalConfig) SetEnableDcr(v bool)`

SetEnableDcr sets EnableDcr field to given value.

### HasEnableDcr

`func (o *PortalConfig) HasEnableDcr() bool`

HasEnableDcr returns a boolean if a field has been set.

### GetEnableMultiSelection

`func (o *PortalConfig) GetEnableMultiSelection() bool`

GetEnableMultiSelection returns the EnableMultiSelection field if non-nil, zero value otherwise.

### GetEnableMultiSelectionOk

`func (o *PortalConfig) GetEnableMultiSelectionOk() (*bool, bool)`

GetEnableMultiSelectionOk returns a tuple with the EnableMultiSelection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableMultiSelection

`func (o *PortalConfig) SetEnableMultiSelection(v bool)`

SetEnableMultiSelection sets EnableMultiSelection field to given value.

### HasEnableMultiSelection

`func (o *PortalConfig) HasEnableMultiSelection() bool`

HasEnableMultiSelection returns a boolean if a field has been set.

### GetId

`func (o *PortalConfig) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PortalConfig) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PortalConfig) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PortalConfig) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKeyRequestFields

`func (o *PortalConfig) GetKeyRequestFields() []string`

GetKeyRequestFields returns the KeyRequestFields field if non-nil, zero value otherwise.

### GetKeyRequestFieldsOk

`func (o *PortalConfig) GetKeyRequestFieldsOk() (*[]string, bool)`

GetKeyRequestFieldsOk returns a tuple with the KeyRequestFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyRequestFields

`func (o *PortalConfig) SetKeyRequestFields(v []string)`

SetKeyRequestFields sets KeyRequestFields field to given value.

### HasKeyRequestFields

`func (o *PortalConfig) HasKeyRequestFields() bool`

HasKeyRequestFields returns a boolean if a field has been set.

### SetKeyRequestFieldsNil

`func (o *PortalConfig) SetKeyRequestFieldsNil(b bool)`

 SetKeyRequestFieldsNil sets the value for KeyRequestFields to be an explicit nil

### UnsetKeyRequestFields
`func (o *PortalConfig) UnsetKeyRequestFields()`

UnsetKeyRequestFields ensures that no value is present for KeyRequestFields, not even an explicit nil
### GetMailOptions

`func (o *PortalConfig) GetMailOptions() PortalConfigMailOptions`

GetMailOptions returns the MailOptions field if non-nil, zero value otherwise.

### GetMailOptionsOk

`func (o *PortalConfig) GetMailOptionsOk() (*PortalConfigMailOptions, bool)`

GetMailOptionsOk returns a tuple with the MailOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMailOptions

`func (o *PortalConfig) SetMailOptions(v PortalConfigMailOptions)`

SetMailOptions sets MailOptions field to given value.

### HasMailOptions

`func (o *PortalConfig) HasMailOptions() bool`

HasMailOptions returns a boolean if a field has been set.

### GetOauthUsageLimit

`func (o *PortalConfig) GetOauthUsageLimit() int32`

GetOauthUsageLimit returns the OauthUsageLimit field if non-nil, zero value otherwise.

### GetOauthUsageLimitOk

`func (o *PortalConfig) GetOauthUsageLimitOk() (*int32, bool)`

GetOauthUsageLimitOk returns a tuple with the OauthUsageLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauthUsageLimit

`func (o *PortalConfig) SetOauthUsageLimit(v int32)`

SetOauthUsageLimit sets OauthUsageLimit field to given value.

### HasOauthUsageLimit

`func (o *PortalConfig) HasOauthUsageLimit() bool`

HasOauthUsageLimit returns a boolean if a field has been set.

### GetOrgId

`func (o *PortalConfig) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *PortalConfig) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *PortalConfig) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *PortalConfig) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetOverride

`func (o *PortalConfig) GetOverride() bool`

GetOverride returns the Override field if non-nil, zero value otherwise.

### GetOverrideOk

`func (o *PortalConfig) GetOverrideOk() (*bool, bool)`

GetOverrideOk returns a tuple with the Override field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverride

`func (o *PortalConfig) SetOverride(v bool)`

SetOverride sets Override field to given value.

### HasOverride

`func (o *PortalConfig) HasOverride() bool`

HasOverride returns a boolean if a field has been set.

### GetRedirectOnKeyRequest

`func (o *PortalConfig) GetRedirectOnKeyRequest() bool`

GetRedirectOnKeyRequest returns the RedirectOnKeyRequest field if non-nil, zero value otherwise.

### GetRedirectOnKeyRequestOk

`func (o *PortalConfig) GetRedirectOnKeyRequestOk() (*bool, bool)`

GetRedirectOnKeyRequestOk returns a tuple with the RedirectOnKeyRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectOnKeyRequest

`func (o *PortalConfig) SetRedirectOnKeyRequest(v bool)`

SetRedirectOnKeyRequest sets RedirectOnKeyRequest field to given value.

### HasRedirectOnKeyRequest

`func (o *PortalConfig) HasRedirectOnKeyRequest() bool`

HasRedirectOnKeyRequest returns a boolean if a field has been set.

### GetRedirectTo

`func (o *PortalConfig) GetRedirectTo() string`

GetRedirectTo returns the RedirectTo field if non-nil, zero value otherwise.

### GetRedirectToOk

`func (o *PortalConfig) GetRedirectToOk() (*string, bool)`

GetRedirectToOk returns a tuple with the RedirectTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectTo

`func (o *PortalConfig) SetRedirectTo(v string)`

SetRedirectTo sets RedirectTo field to given value.

### HasRedirectTo

`func (o *PortalConfig) HasRedirectTo() bool`

HasRedirectTo returns a boolean if a field has been set.

### GetRequireKeyApproval

`func (o *PortalConfig) GetRequireKeyApproval() bool`

GetRequireKeyApproval returns the RequireKeyApproval field if non-nil, zero value otherwise.

### GetRequireKeyApprovalOk

`func (o *PortalConfig) GetRequireKeyApprovalOk() (*bool, bool)`

GetRequireKeyApprovalOk returns a tuple with the RequireKeyApproval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireKeyApproval

`func (o *PortalConfig) SetRequireKeyApproval(v bool)`

SetRequireKeyApproval sets RequireKeyApproval field to given value.

### HasRequireKeyApproval

`func (o *PortalConfig) HasRequireKeyApproval() bool`

HasRequireKeyApproval returns a boolean if a field has been set.

### GetSecureKeyApproval

`func (o *PortalConfig) GetSecureKeyApproval() bool`

GetSecureKeyApproval returns the SecureKeyApproval field if non-nil, zero value otherwise.

### GetSecureKeyApprovalOk

`func (o *PortalConfig) GetSecureKeyApprovalOk() (*bool, bool)`

GetSecureKeyApprovalOk returns a tuple with the SecureKeyApproval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureKeyApproval

`func (o *PortalConfig) SetSecureKeyApproval(v bool)`

SetSecureKeyApproval sets SecureKeyApproval field to given value.

### HasSecureKeyApproval

`func (o *PortalConfig) HasSecureKeyApproval() bool`

HasSecureKeyApproval returns a boolean if a field has been set.

### GetSignupFields

`func (o *PortalConfig) GetSignupFields() []string`

GetSignupFields returns the SignupFields field if non-nil, zero value otherwise.

### GetSignupFieldsOk

`func (o *PortalConfig) GetSignupFieldsOk() (*[]string, bool)`

GetSignupFieldsOk returns a tuple with the SignupFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignupFields

`func (o *PortalConfig) SetSignupFields(v []string)`

SetSignupFields sets SignupFields field to given value.

### HasSignupFields

`func (o *PortalConfig) HasSignupFields() bool`

HasSignupFields returns a boolean if a field has been set.

### SetSignupFieldsNil

`func (o *PortalConfig) SetSignupFieldsNil(b bool)`

 SetSignupFieldsNil sets the value for SignupFields to be an explicit nil

### UnsetSignupFields
`func (o *PortalConfig) UnsetSignupFields()`

UnsetSignupFields ensures that no value is present for SignupFields, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


