# ApiDefinitionWrapper

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiDefinition** | Pointer to [**APIDefinition**](APIDefinition.md) |  | [optional] 
**ApiModel** | Pointer to **map[string]interface{}** |  | [optional] 
**Categories** | Pointer to **[]string** |  | [optional] 
**CreatedAt** | Pointer to **NullableTime** |  | [optional] 
**HookReferences** | Pointer to [**[]HookReference**](HookReference.md) |  | [optional] 
**IsSite** | Pointer to **bool** |  | [optional] 
**Oas** | Pointer to [**NullableOAS**](OAS.md) |  | [optional] 
**SortBy** | Pointer to **int32** |  | [optional] 
**UpdatedAt** | Pointer to **NullableTime** |  | [optional] 
**UserGroupOwners** | Pointer to **[]string** |  | [optional] 
**UserOwners** | Pointer to **[]string** |  | [optional] 

## Methods

### NewApiDefinitionWrapper

`func NewApiDefinitionWrapper() *ApiDefinitionWrapper`

NewApiDefinitionWrapper instantiates a new ApiDefinitionWrapper object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiDefinitionWrapperWithDefaults

`func NewApiDefinitionWrapperWithDefaults() *ApiDefinitionWrapper`

NewApiDefinitionWrapperWithDefaults instantiates a new ApiDefinitionWrapper object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiDefinition

`func (o *ApiDefinitionWrapper) GetApiDefinition() APIDefinition`

GetApiDefinition returns the ApiDefinition field if non-nil, zero value otherwise.

### GetApiDefinitionOk

`func (o *ApiDefinitionWrapper) GetApiDefinitionOk() (*APIDefinition, bool)`

GetApiDefinitionOk returns a tuple with the ApiDefinition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiDefinition

`func (o *ApiDefinitionWrapper) SetApiDefinition(v APIDefinition)`

SetApiDefinition sets ApiDefinition field to given value.

### HasApiDefinition

`func (o *ApiDefinitionWrapper) HasApiDefinition() bool`

HasApiDefinition returns a boolean if a field has been set.

### GetApiModel

`func (o *ApiDefinitionWrapper) GetApiModel() map[string]interface{}`

GetApiModel returns the ApiModel field if non-nil, zero value otherwise.

### GetApiModelOk

`func (o *ApiDefinitionWrapper) GetApiModelOk() (*map[string]interface{}, bool)`

GetApiModelOk returns a tuple with the ApiModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiModel

`func (o *ApiDefinitionWrapper) SetApiModel(v map[string]interface{})`

SetApiModel sets ApiModel field to given value.

### HasApiModel

`func (o *ApiDefinitionWrapper) HasApiModel() bool`

HasApiModel returns a boolean if a field has been set.

### GetCategories

`func (o *ApiDefinitionWrapper) GetCategories() []string`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *ApiDefinitionWrapper) GetCategoriesOk() (*[]string, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *ApiDefinitionWrapper) SetCategories(v []string)`

SetCategories sets Categories field to given value.

### HasCategories

`func (o *ApiDefinitionWrapper) HasCategories() bool`

HasCategories returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ApiDefinitionWrapper) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ApiDefinitionWrapper) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ApiDefinitionWrapper) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ApiDefinitionWrapper) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *ApiDefinitionWrapper) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *ApiDefinitionWrapper) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetHookReferences

`func (o *ApiDefinitionWrapper) GetHookReferences() []HookReference`

GetHookReferences returns the HookReferences field if non-nil, zero value otherwise.

### GetHookReferencesOk

`func (o *ApiDefinitionWrapper) GetHookReferencesOk() (*[]HookReference, bool)`

GetHookReferencesOk returns a tuple with the HookReferences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHookReferences

`func (o *ApiDefinitionWrapper) SetHookReferences(v []HookReference)`

SetHookReferences sets HookReferences field to given value.

### HasHookReferences

`func (o *ApiDefinitionWrapper) HasHookReferences() bool`

HasHookReferences returns a boolean if a field has been set.

### SetHookReferencesNil

`func (o *ApiDefinitionWrapper) SetHookReferencesNil(b bool)`

 SetHookReferencesNil sets the value for HookReferences to be an explicit nil

### UnsetHookReferences
`func (o *ApiDefinitionWrapper) UnsetHookReferences()`

UnsetHookReferences ensures that no value is present for HookReferences, not even an explicit nil
### GetIsSite

`func (o *ApiDefinitionWrapper) GetIsSite() bool`

GetIsSite returns the IsSite field if non-nil, zero value otherwise.

### GetIsSiteOk

`func (o *ApiDefinitionWrapper) GetIsSiteOk() (*bool, bool)`

GetIsSiteOk returns a tuple with the IsSite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSite

`func (o *ApiDefinitionWrapper) SetIsSite(v bool)`

SetIsSite sets IsSite field to given value.

### HasIsSite

`func (o *ApiDefinitionWrapper) HasIsSite() bool`

HasIsSite returns a boolean if a field has been set.

### GetOas

`func (o *ApiDefinitionWrapper) GetOas() OAS`

GetOas returns the Oas field if non-nil, zero value otherwise.

### GetOasOk

`func (o *ApiDefinitionWrapper) GetOasOk() (*OAS, bool)`

GetOasOk returns a tuple with the Oas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOas

`func (o *ApiDefinitionWrapper) SetOas(v OAS)`

SetOas sets Oas field to given value.

### HasOas

`func (o *ApiDefinitionWrapper) HasOas() bool`

HasOas returns a boolean if a field has been set.

### SetOasNil

`func (o *ApiDefinitionWrapper) SetOasNil(b bool)`

 SetOasNil sets the value for Oas to be an explicit nil

### UnsetOas
`func (o *ApiDefinitionWrapper) UnsetOas()`

UnsetOas ensures that no value is present for Oas, not even an explicit nil
### GetSortBy

`func (o *ApiDefinitionWrapper) GetSortBy() int32`

GetSortBy returns the SortBy field if non-nil, zero value otherwise.

### GetSortByOk

`func (o *ApiDefinitionWrapper) GetSortByOk() (*int32, bool)`

GetSortByOk returns a tuple with the SortBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortBy

`func (o *ApiDefinitionWrapper) SetSortBy(v int32)`

SetSortBy sets SortBy field to given value.

### HasSortBy

`func (o *ApiDefinitionWrapper) HasSortBy() bool`

HasSortBy returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ApiDefinitionWrapper) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ApiDefinitionWrapper) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ApiDefinitionWrapper) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ApiDefinitionWrapper) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *ApiDefinitionWrapper) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *ApiDefinitionWrapper) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetUserGroupOwners

`func (o *ApiDefinitionWrapper) GetUserGroupOwners() []string`

GetUserGroupOwners returns the UserGroupOwners field if non-nil, zero value otherwise.

### GetUserGroupOwnersOk

`func (o *ApiDefinitionWrapper) GetUserGroupOwnersOk() (*[]string, bool)`

GetUserGroupOwnersOk returns a tuple with the UserGroupOwners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserGroupOwners

`func (o *ApiDefinitionWrapper) SetUserGroupOwners(v []string)`

SetUserGroupOwners sets UserGroupOwners field to given value.

### HasUserGroupOwners

`func (o *ApiDefinitionWrapper) HasUserGroupOwners() bool`

HasUserGroupOwners returns a boolean if a field has been set.

### SetUserGroupOwnersNil

`func (o *ApiDefinitionWrapper) SetUserGroupOwnersNil(b bool)`

 SetUserGroupOwnersNil sets the value for UserGroupOwners to be an explicit nil

### UnsetUserGroupOwners
`func (o *ApiDefinitionWrapper) UnsetUserGroupOwners()`

UnsetUserGroupOwners ensures that no value is present for UserGroupOwners, not even an explicit nil
### GetUserOwners

`func (o *ApiDefinitionWrapper) GetUserOwners() []string`

GetUserOwners returns the UserOwners field if non-nil, zero value otherwise.

### GetUserOwnersOk

`func (o *ApiDefinitionWrapper) GetUserOwnersOk() (*[]string, bool)`

GetUserOwnersOk returns a tuple with the UserOwners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserOwners

`func (o *ApiDefinitionWrapper) SetUserOwners(v []string)`

SetUserOwners sets UserOwners field to given value.

### HasUserOwners

`func (o *ApiDefinitionWrapper) HasUserOwners() bool`

HasUserOwners returns a boolean if a field has been set.

### SetUserOwnersNil

`func (o *ApiDefinitionWrapper) SetUserOwnersNil(b bool)`

 SetUserOwnersNil sets the value for UserOwners to be an explicit nil

### UnsetUserOwners
`func (o *ApiDefinitionWrapper) UnsetUserOwners()`

UnsetUserOwners ensures that no value is present for UserOwners, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


