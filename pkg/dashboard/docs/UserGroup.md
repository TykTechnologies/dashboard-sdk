# UserGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **bool** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**OrgId** | Pointer to **string** |  | [optional] 
**PasswordMaxDays** | Pointer to **int32** |  | [optional] 
**UserPermissions** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewUserGroup

`func NewUserGroup() *UserGroup`

NewUserGroup instantiates a new UserGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserGroupWithDefaults

`func NewUserGroupWithDefaults() *UserGroup`

NewUserGroupWithDefaults instantiates a new UserGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *UserGroup) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *UserGroup) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *UserGroup) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *UserGroup) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetDescription

`func (o *UserGroup) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UserGroup) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UserGroup) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UserGroup) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *UserGroup) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserGroup) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserGroup) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *UserGroup) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *UserGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UserGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UserGroup) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UserGroup) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrgId

`func (o *UserGroup) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *UserGroup) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *UserGroup) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *UserGroup) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetPasswordMaxDays

`func (o *UserGroup) GetPasswordMaxDays() int32`

GetPasswordMaxDays returns the PasswordMaxDays field if non-nil, zero value otherwise.

### GetPasswordMaxDaysOk

`func (o *UserGroup) GetPasswordMaxDaysOk() (*int32, bool)`

GetPasswordMaxDaysOk returns a tuple with the PasswordMaxDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordMaxDays

`func (o *UserGroup) SetPasswordMaxDays(v int32)`

SetPasswordMaxDays sets PasswordMaxDays field to given value.

### HasPasswordMaxDays

`func (o *UserGroup) HasPasswordMaxDays() bool`

HasPasswordMaxDays returns a boolean if a field has been set.

### GetUserPermissions

`func (o *UserGroup) GetUserPermissions() map[string]string`

GetUserPermissions returns the UserPermissions field if non-nil, zero value otherwise.

### GetUserPermissionsOk

`func (o *UserGroup) GetUserPermissionsOk() (*map[string]string, bool)`

GetUserPermissionsOk returns a tuple with the UserPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserPermissions

`func (o *UserGroup) SetUserPermissions(v map[string]string)`

SetUserPermissions sets UserPermissions field to given value.

### HasUserPermissions

`func (o *UserGroup) HasUserPermissions() bool`

HasUserPermissions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


