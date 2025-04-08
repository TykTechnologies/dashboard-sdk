# MigrateAPIResponseSuccessDefinitionsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiID** | Pointer to **string** |  | [optional] 
**BaseAPI** | Pointer to **map[string]interface{}** | contains the migrated base API in Tyk OAS format | [optional] 
**Versions** | Pointer to **[]map[string]interface{}** | contains the versioned API definitions in Tyk OAS format | [optional] 

## Methods

### NewMigrateAPIResponseSuccessDefinitionsInner

`func NewMigrateAPIResponseSuccessDefinitionsInner() *MigrateAPIResponseSuccessDefinitionsInner`

NewMigrateAPIResponseSuccessDefinitionsInner instantiates a new MigrateAPIResponseSuccessDefinitionsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMigrateAPIResponseSuccessDefinitionsInnerWithDefaults

`func NewMigrateAPIResponseSuccessDefinitionsInnerWithDefaults() *MigrateAPIResponseSuccessDefinitionsInner`

NewMigrateAPIResponseSuccessDefinitionsInnerWithDefaults instantiates a new MigrateAPIResponseSuccessDefinitionsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiID

`func (o *MigrateAPIResponseSuccessDefinitionsInner) GetApiID() string`

GetApiID returns the ApiID field if non-nil, zero value otherwise.

### GetApiIDOk

`func (o *MigrateAPIResponseSuccessDefinitionsInner) GetApiIDOk() (*string, bool)`

GetApiIDOk returns a tuple with the ApiID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiID

`func (o *MigrateAPIResponseSuccessDefinitionsInner) SetApiID(v string)`

SetApiID sets ApiID field to given value.

### HasApiID

`func (o *MigrateAPIResponseSuccessDefinitionsInner) HasApiID() bool`

HasApiID returns a boolean if a field has been set.

### GetBaseAPI

`func (o *MigrateAPIResponseSuccessDefinitionsInner) GetBaseAPI() map[string]interface{}`

GetBaseAPI returns the BaseAPI field if non-nil, zero value otherwise.

### GetBaseAPIOk

`func (o *MigrateAPIResponseSuccessDefinitionsInner) GetBaseAPIOk() (*map[string]interface{}, bool)`

GetBaseAPIOk returns a tuple with the BaseAPI field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseAPI

`func (o *MigrateAPIResponseSuccessDefinitionsInner) SetBaseAPI(v map[string]interface{})`

SetBaseAPI sets BaseAPI field to given value.

### HasBaseAPI

`func (o *MigrateAPIResponseSuccessDefinitionsInner) HasBaseAPI() bool`

HasBaseAPI returns a boolean if a field has been set.

### GetVersions

`func (o *MigrateAPIResponseSuccessDefinitionsInner) GetVersions() []map[string]interface{}`

GetVersions returns the Versions field if non-nil, zero value otherwise.

### GetVersionsOk

`func (o *MigrateAPIResponseSuccessDefinitionsInner) GetVersionsOk() (*[]map[string]interface{}, bool)`

GetVersionsOk returns a tuple with the Versions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersions

`func (o *MigrateAPIResponseSuccessDefinitionsInner) SetVersions(v []map[string]interface{})`

SetVersions sets Versions field to given value.

### HasVersions

`func (o *MigrateAPIResponseSuccessDefinitionsInner) HasVersions() bool`

HasVersions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


