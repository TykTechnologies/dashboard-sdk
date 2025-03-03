# MigrateAPIResponseSuccess

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int32** | count of APIs migrated successfully | [optional] 
**ApiIDs** | Pointer to **[]string** | list of API IDs migrated successfully | [optional] 
**StagedAPIs** | Pointer to [**[]MigrateAPIResponseSuccessStagedAPIsInner**](MigrateAPIResponseSuccessStagedAPIsInner.md) | reports the details of APIs staged in stage mode | [optional] 
**Definitions** | Pointer to [**[]MigrateAPIResponseSuccessDefinitionsInner**](MigrateAPIResponseSuccessDefinitionsInner.md) |  | [optional] 

## Methods

### NewMigrateAPIResponseSuccess

`func NewMigrateAPIResponseSuccess() *MigrateAPIResponseSuccess`

NewMigrateAPIResponseSuccess instantiates a new MigrateAPIResponseSuccess object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMigrateAPIResponseSuccessWithDefaults

`func NewMigrateAPIResponseSuccessWithDefaults() *MigrateAPIResponseSuccess`

NewMigrateAPIResponseSuccessWithDefaults instantiates a new MigrateAPIResponseSuccess object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *MigrateAPIResponseSuccess) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *MigrateAPIResponseSuccess) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *MigrateAPIResponseSuccess) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *MigrateAPIResponseSuccess) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetApiIDs

`func (o *MigrateAPIResponseSuccess) GetApiIDs() []string`

GetApiIDs returns the ApiIDs field if non-nil, zero value otherwise.

### GetApiIDsOk

`func (o *MigrateAPIResponseSuccess) GetApiIDsOk() (*[]string, bool)`

GetApiIDsOk returns a tuple with the ApiIDs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiIDs

`func (o *MigrateAPIResponseSuccess) SetApiIDs(v []string)`

SetApiIDs sets ApiIDs field to given value.

### HasApiIDs

`func (o *MigrateAPIResponseSuccess) HasApiIDs() bool`

HasApiIDs returns a boolean if a field has been set.

### GetStagedAPIs

`func (o *MigrateAPIResponseSuccess) GetStagedAPIs() []MigrateAPIResponseSuccessStagedAPIsInner`

GetStagedAPIs returns the StagedAPIs field if non-nil, zero value otherwise.

### GetStagedAPIsOk

`func (o *MigrateAPIResponseSuccess) GetStagedAPIsOk() (*[]MigrateAPIResponseSuccessStagedAPIsInner, bool)`

GetStagedAPIsOk returns a tuple with the StagedAPIs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStagedAPIs

`func (o *MigrateAPIResponseSuccess) SetStagedAPIs(v []MigrateAPIResponseSuccessStagedAPIsInner)`

SetStagedAPIs sets StagedAPIs field to given value.

### HasStagedAPIs

`func (o *MigrateAPIResponseSuccess) HasStagedAPIs() bool`

HasStagedAPIs returns a boolean if a field has been set.

### GetDefinitions

`func (o *MigrateAPIResponseSuccess) GetDefinitions() []MigrateAPIResponseSuccessDefinitionsInner`

GetDefinitions returns the Definitions field if non-nil, zero value otherwise.

### GetDefinitionsOk

`func (o *MigrateAPIResponseSuccess) GetDefinitionsOk() (*[]MigrateAPIResponseSuccessDefinitionsInner, bool)`

GetDefinitionsOk returns a tuple with the Definitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefinitions

`func (o *MigrateAPIResponseSuccess) SetDefinitions(v []MigrateAPIResponseSuccessDefinitionsInner)`

SetDefinitions sets Definitions field to given value.

### HasDefinitions

`func (o *MigrateAPIResponseSuccess) HasDefinitions() bool`

HasDefinitions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


