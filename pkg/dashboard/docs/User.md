# User

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessKey** | Pointer to **string** |  | [optional] 
**Active** | Pointer to **bool** |  | [optional] 
**ApiModel** | Pointer to **map[string]interface{}** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**EmailAddress** | **string** |  | 
**FirstName** | **string** |  | 
**GroupId** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**LastLoginDate** | Pointer to **time.Time** |  | [optional] 
**LastName** | **string** |  | 
**OrgId** | Pointer to **string** |  | [optional] 
**PasswordMaxDays** | Pointer to **int32** |  | [optional] 
**PasswordUpdated** | Pointer to **time.Time** |  | [optional] 
**UserPermissions** | **map[string]string** |  | 

## Methods

### NewUser

`func NewUser(emailAddress string, firstName string, lastName string, userPermissions map[string]string, ) *User`

NewUser instantiates a new User object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserWithDefaults

`func NewUserWithDefaults() *User`

NewUserWithDefaults instantiates a new User object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessKey

`func (o *User) GetAccessKey() string`

GetAccessKey returns the AccessKey field if non-nil, zero value otherwise.

### GetAccessKeyOk

`func (o *User) GetAccessKeyOk() (*string, bool)`

GetAccessKeyOk returns a tuple with the AccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKey

`func (o *User) SetAccessKey(v string)`

SetAccessKey sets AccessKey field to given value.

### HasAccessKey

`func (o *User) HasAccessKey() bool`

HasAccessKey returns a boolean if a field has been set.

### GetActive

`func (o *User) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *User) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *User) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *User) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetApiModel

`func (o *User) GetApiModel() map[string]interface{}`

GetApiModel returns the ApiModel field if non-nil, zero value otherwise.

### GetApiModelOk

`func (o *User) GetApiModelOk() (*map[string]interface{}, bool)`

GetApiModelOk returns a tuple with the ApiModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiModel

`func (o *User) SetApiModel(v map[string]interface{})`

SetApiModel sets ApiModel field to given value.

### HasApiModel

`func (o *User) HasApiModel() bool`

HasApiModel returns a boolean if a field has been set.

### GetCreatedAt

`func (o *User) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *User) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *User) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *User) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetEmailAddress

`func (o *User) GetEmailAddress() string`

GetEmailAddress returns the EmailAddress field if non-nil, zero value otherwise.

### GetEmailAddressOk

`func (o *User) GetEmailAddressOk() (*string, bool)`

GetEmailAddressOk returns a tuple with the EmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAddress

`func (o *User) SetEmailAddress(v string)`

SetEmailAddress sets EmailAddress field to given value.


### GetFirstName

`func (o *User) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *User) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *User) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.


### GetGroupId

`func (o *User) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *User) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *User) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *User) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetId

`func (o *User) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *User) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *User) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *User) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastLoginDate

`func (o *User) GetLastLoginDate() time.Time`

GetLastLoginDate returns the LastLoginDate field if non-nil, zero value otherwise.

### GetLastLoginDateOk

`func (o *User) GetLastLoginDateOk() (*time.Time, bool)`

GetLastLoginDateOk returns a tuple with the LastLoginDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLoginDate

`func (o *User) SetLastLoginDate(v time.Time)`

SetLastLoginDate sets LastLoginDate field to given value.

### HasLastLoginDate

`func (o *User) HasLastLoginDate() bool`

HasLastLoginDate returns a boolean if a field has been set.

### GetLastName

`func (o *User) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *User) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *User) SetLastName(v string)`

SetLastName sets LastName field to given value.


### GetOrgId

`func (o *User) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *User) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *User) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *User) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetPasswordMaxDays

`func (o *User) GetPasswordMaxDays() int32`

GetPasswordMaxDays returns the PasswordMaxDays field if non-nil, zero value otherwise.

### GetPasswordMaxDaysOk

`func (o *User) GetPasswordMaxDaysOk() (*int32, bool)`

GetPasswordMaxDaysOk returns a tuple with the PasswordMaxDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordMaxDays

`func (o *User) SetPasswordMaxDays(v int32)`

SetPasswordMaxDays sets PasswordMaxDays field to given value.

### HasPasswordMaxDays

`func (o *User) HasPasswordMaxDays() bool`

HasPasswordMaxDays returns a boolean if a field has been set.

### GetPasswordUpdated

`func (o *User) GetPasswordUpdated() time.Time`

GetPasswordUpdated returns the PasswordUpdated field if non-nil, zero value otherwise.

### GetPasswordUpdatedOk

`func (o *User) GetPasswordUpdatedOk() (*time.Time, bool)`

GetPasswordUpdatedOk returns a tuple with the PasswordUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordUpdated

`func (o *User) SetPasswordUpdated(v time.Time)`

SetPasswordUpdated sets PasswordUpdated field to given value.

### HasPasswordUpdated

`func (o *User) HasPasswordUpdated() bool`

HasPasswordUpdated returns a boolean if a field has been set.

### GetUserPermissions

`func (o *User) GetUserPermissions() map[string]string`

GetUserPermissions returns the UserPermissions field if non-nil, zero value otherwise.

### GetUserPermissionsOk

`func (o *User) GetUserPermissionsOk() (*map[string]string, bool)`

GetUserPermissionsOk returns a tuple with the UserPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserPermissions

`func (o *User) SetUserPermissions(v map[string]string)`

SetUserPermissions sets UserPermissions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


