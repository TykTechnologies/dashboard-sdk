# ApiDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiDefinition** | Pointer to [**ApidefAPIDefinition**](ApidefAPIDefinition.md) |  | [optional] 
**ApiModel** | Pointer to **map[string]interface{}** |  | [optional] 
**Categories** | Pointer to **[]string** |  | [optional] 
**CreatedAt** | Pointer to **NullableTime** |  | [optional] 
**HookReferences** | Pointer to [**[]HookReference**](HookReference.md) |  | [optional] 
**IsSite** | Pointer to **bool** |  | [optional] 
**Oas** | Pointer to [**OasOAS**](OasOAS.md) |  | [optional] 
**SortBy** | Pointer to **int32** |  | [optional] 
**UpdatedAt** | Pointer to **NullableTime** |  | [optional] 
**UserGroupOwners** | Pointer to **[]string** |  | [optional] 
**UserOwners** | Pointer to **[]string** |  | [optional] 

## Methods

### NewApiDefinition

`func NewApiDefinition() *ApiDefinition`

NewApiDefinition instantiates a new ApiDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiDefinitionWithDefaults

`func NewApiDefinitionWithDefaults() *ApiDefinition`

NewApiDefinitionWithDefaults instantiates a new ApiDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiDefinition

`func (o *ApiDefinition) GetApiDefinition() ApidefAPIDefinition`

GetApiDefinition returns the ApiDefinition field if non-nil, zero value otherwise.

### GetApiDefinitionOk

`func (o *ApiDefinition) GetApiDefinitionOk() (*ApidefAPIDefinition, bool)`

GetApiDefinitionOk returns a tuple with the ApiDefinition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiDefinition

`func (o *ApiDefinition) SetApiDefinition(v ApidefAPIDefinition)`

SetApiDefinition sets ApiDefinition field to given value.

### HasApiDefinition

`func (o *ApiDefinition) HasApiDefinition() bool`

HasApiDefinition returns a boolean if a field has been set.

### GetApiModel

`func (o *ApiDefinition) GetApiModel() map[string]interface{}`

GetApiModel returns the ApiModel field if non-nil, zero value otherwise.

### GetApiModelOk

`func (o *ApiDefinition) GetApiModelOk() (*map[string]interface{}, bool)`

GetApiModelOk returns a tuple with the ApiModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiModel

`func (o *ApiDefinition) SetApiModel(v map[string]interface{})`

SetApiModel sets ApiModel field to given value.

### HasApiModel

`func (o *ApiDefinition) HasApiModel() bool`

HasApiModel returns a boolean if a field has been set.

### GetCategories

`func (o *ApiDefinition) GetCategories() []string`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *ApiDefinition) GetCategoriesOk() (*[]string, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *ApiDefinition) SetCategories(v []string)`

SetCategories sets Categories field to given value.

### HasCategories

`func (o *ApiDefinition) HasCategories() bool`

HasCategories returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ApiDefinition) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ApiDefinition) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ApiDefinition) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ApiDefinition) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *ApiDefinition) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *ApiDefinition) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetHookReferences

`func (o *ApiDefinition) GetHookReferences() []HookReference`

GetHookReferences returns the HookReferences field if non-nil, zero value otherwise.

### GetHookReferencesOk

`func (o *ApiDefinition) GetHookReferencesOk() (*[]HookReference, bool)`

GetHookReferencesOk returns a tuple with the HookReferences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHookReferences

`func (o *ApiDefinition) SetHookReferences(v []HookReference)`

SetHookReferences sets HookReferences field to given value.

### HasHookReferences

`func (o *ApiDefinition) HasHookReferences() bool`

HasHookReferences returns a boolean if a field has been set.

### SetHookReferencesNil

`func (o *ApiDefinition) SetHookReferencesNil(b bool)`

 SetHookReferencesNil sets the value for HookReferences to be an explicit nil

### UnsetHookReferences
`func (o *ApiDefinition) UnsetHookReferences()`

UnsetHookReferences ensures that no value is present for HookReferences, not even an explicit nil
### GetIsSite

`func (o *ApiDefinition) GetIsSite() bool`

GetIsSite returns the IsSite field if non-nil, zero value otherwise.

### GetIsSiteOk

`func (o *ApiDefinition) GetIsSiteOk() (*bool, bool)`

GetIsSiteOk returns a tuple with the IsSite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSite

`func (o *ApiDefinition) SetIsSite(v bool)`

SetIsSite sets IsSite field to given value.

### HasIsSite

`func (o *ApiDefinition) HasIsSite() bool`

HasIsSite returns a boolean if a field has been set.

### GetOas

`func (o *ApiDefinition) GetOas() OasOAS`

GetOas returns the Oas field if non-nil, zero value otherwise.

### GetOasOk

`func (o *ApiDefinition) GetOasOk() (*OasOAS, bool)`

GetOasOk returns a tuple with the Oas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOas

`func (o *ApiDefinition) SetOas(v OasOAS)`

SetOas sets Oas field to given value.

### HasOas

`func (o *ApiDefinition) HasOas() bool`

HasOas returns a boolean if a field has been set.

### GetSortBy

`func (o *ApiDefinition) GetSortBy() int32`

GetSortBy returns the SortBy field if non-nil, zero value otherwise.

### GetSortByOk

`func (o *ApiDefinition) GetSortByOk() (*int32, bool)`

GetSortByOk returns a tuple with the SortBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortBy

`func (o *ApiDefinition) SetSortBy(v int32)`

SetSortBy sets SortBy field to given value.

### HasSortBy

`func (o *ApiDefinition) HasSortBy() bool`

HasSortBy returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ApiDefinition) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ApiDefinition) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ApiDefinition) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ApiDefinition) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *ApiDefinition) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *ApiDefinition) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetUserGroupOwners

`func (o *ApiDefinition) GetUserGroupOwners() []string`

GetUserGroupOwners returns the UserGroupOwners field if non-nil, zero value otherwise.

### GetUserGroupOwnersOk

`func (o *ApiDefinition) GetUserGroupOwnersOk() (*[]string, bool)`

GetUserGroupOwnersOk returns a tuple with the UserGroupOwners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserGroupOwners

`func (o *ApiDefinition) SetUserGroupOwners(v []string)`

SetUserGroupOwners sets UserGroupOwners field to given value.

### HasUserGroupOwners

`func (o *ApiDefinition) HasUserGroupOwners() bool`

HasUserGroupOwners returns a boolean if a field has been set.

### SetUserGroupOwnersNil

`func (o *ApiDefinition) SetUserGroupOwnersNil(b bool)`

 SetUserGroupOwnersNil sets the value for UserGroupOwners to be an explicit nil

### UnsetUserGroupOwners
`func (o *ApiDefinition) UnsetUserGroupOwners()`

UnsetUserGroupOwners ensures that no value is present for UserGroupOwners, not even an explicit nil
### GetUserOwners

`func (o *ApiDefinition) GetUserOwners() []string`

GetUserOwners returns the UserOwners field if non-nil, zero value otherwise.

### GetUserOwnersOk

`func (o *ApiDefinition) GetUserOwnersOk() (*[]string, bool)`

GetUserOwnersOk returns a tuple with the UserOwners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserOwners

`func (o *ApiDefinition) SetUserOwners(v []string)`

SetUserOwners sets UserOwners field to given value.

### HasUserOwners

`func (o *ApiDefinition) HasUserOwners() bool`

HasUserOwners returns a boolean if a field has been set.

### SetUserOwnersNil

`func (o *ApiDefinition) SetUserOwnersNil(b bool)`

 SetUserOwnersNil sets the value for UserOwners to be an explicit nil

### UnsetUserOwners
`func (o *ApiDefinition) UnsetUserOwners()`

UnsetUserOwners ensures that no value is present for UserOwners, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


