# MigrateAPIRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiIDs** | Pointer to **[]string** | List of API IDs to migrate. Cannot be used together with &#39;all&#39; | [optional] 
**All** | Pointer to **bool** | Migrate all APIs. Cannot be used together with &#39;apiIDs&#39; | [optional] 
**AbortOnFailure** | Pointer to **bool** | Stop migration process on first failure | [optional] 
**Mode** | **string** | Migration mode to use | 
**OverrideStaged** | Pointer to **bool** | When mode is staged and overrideStaged is set to true, migration process will overwrite already existing staged API with the same staged ID | [optional] 

## Methods

### NewMigrateAPIRequest

`func NewMigrateAPIRequest(mode string, ) *MigrateAPIRequest`

NewMigrateAPIRequest instantiates a new MigrateAPIRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMigrateAPIRequestWithDefaults

`func NewMigrateAPIRequestWithDefaults() *MigrateAPIRequest`

NewMigrateAPIRequestWithDefaults instantiates a new MigrateAPIRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiIDs

`func (o *MigrateAPIRequest) GetApiIDs() []string`

GetApiIDs returns the ApiIDs field if non-nil, zero value otherwise.

### GetApiIDsOk

`func (o *MigrateAPIRequest) GetApiIDsOk() (*[]string, bool)`

GetApiIDsOk returns a tuple with the ApiIDs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiIDs

`func (o *MigrateAPIRequest) SetApiIDs(v []string)`

SetApiIDs sets ApiIDs field to given value.

### HasApiIDs

`func (o *MigrateAPIRequest) HasApiIDs() bool`

HasApiIDs returns a boolean if a field has been set.

### GetAll

`func (o *MigrateAPIRequest) GetAll() bool`

GetAll returns the All field if non-nil, zero value otherwise.

### GetAllOk

`func (o *MigrateAPIRequest) GetAllOk() (*bool, bool)`

GetAllOk returns a tuple with the All field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAll

`func (o *MigrateAPIRequest) SetAll(v bool)`

SetAll sets All field to given value.

### HasAll

`func (o *MigrateAPIRequest) HasAll() bool`

HasAll returns a boolean if a field has been set.

### GetAbortOnFailure

`func (o *MigrateAPIRequest) GetAbortOnFailure() bool`

GetAbortOnFailure returns the AbortOnFailure field if non-nil, zero value otherwise.

### GetAbortOnFailureOk

`func (o *MigrateAPIRequest) GetAbortOnFailureOk() (*bool, bool)`

GetAbortOnFailureOk returns a tuple with the AbortOnFailure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbortOnFailure

`func (o *MigrateAPIRequest) SetAbortOnFailure(v bool)`

SetAbortOnFailure sets AbortOnFailure field to given value.

### HasAbortOnFailure

`func (o *MigrateAPIRequest) HasAbortOnFailure() bool`

HasAbortOnFailure returns a boolean if a field has been set.

### GetMode

`func (o *MigrateAPIRequest) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *MigrateAPIRequest) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *MigrateAPIRequest) SetMode(v string)`

SetMode sets Mode field to given value.


### GetOverrideStaged

`func (o *MigrateAPIRequest) GetOverrideStaged() bool`

GetOverrideStaged returns the OverrideStaged field if non-nil, zero value otherwise.

### GetOverrideStagedOk

`func (o *MigrateAPIRequest) GetOverrideStagedOk() (*bool, bool)`

GetOverrideStagedOk returns a tuple with the OverrideStaged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideStaged

`func (o *MigrateAPIRequest) SetOverrideStaged(v bool)`

SetOverrideStaged sets OverrideStaged field to given value.

### HasOverrideStaged

`func (o *MigrateAPIRequest) HasOverrideStaged() bool`

HasOverrideStaged returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


