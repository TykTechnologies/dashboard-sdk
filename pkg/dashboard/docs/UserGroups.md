# UserGroups

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Groups** | Pointer to [**[]UserGroup**](UserGroup.md) |  | [optional] 
**Pages** | Pointer to **int32** |  | [optional] 

## Methods

### NewUserGroups

`func NewUserGroups() *UserGroups`

NewUserGroups instantiates a new UserGroups object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserGroupsWithDefaults

`func NewUserGroupsWithDefaults() *UserGroups`

NewUserGroupsWithDefaults instantiates a new UserGroups object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroups

`func (o *UserGroups) GetGroups() []UserGroup`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *UserGroups) GetGroupsOk() (*[]UserGroup, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *UserGroups) SetGroups(v []UserGroup)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *UserGroups) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### SetGroupsNil

`func (o *UserGroups) SetGroupsNil(b bool)`

 SetGroupsNil sets the value for Groups to be an explicit nil

### UnsetGroups
`func (o *UserGroups) UnsetGroups()`

UnsetGroups ensures that no value is present for Groups, not even an explicit nil
### GetPages

`func (o *UserGroups) GetPages() int32`

GetPages returns the Pages field if non-nil, zero value otherwise.

### GetPagesOk

`func (o *UserGroups) GetPagesOk() (*int32, bool)`

GetPagesOk returns a tuple with the Pages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPages

`func (o *UserGroups) SetPages(v int32)`

SetPages sets Pages field to given value.

### HasPages

`func (o *UserGroups) HasPages() bool`

HasPages returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


