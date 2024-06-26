# ApidefUptimeTests

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CheckList** | Pointer to [**[]ApidefHostCheckObject**](ApidefHostCheckObject.md) |  | [optional] 
**Config** | Pointer to [**ApidefUptimeTestsConfig**](ApidefUptimeTestsConfig.md) |  | [optional] 

## Methods

### NewApidefUptimeTests

`func NewApidefUptimeTests() *ApidefUptimeTests`

NewApidefUptimeTests instantiates a new ApidefUptimeTests object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApidefUptimeTestsWithDefaults

`func NewApidefUptimeTestsWithDefaults() *ApidefUptimeTests`

NewApidefUptimeTestsWithDefaults instantiates a new ApidefUptimeTests object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCheckList

`func (o *ApidefUptimeTests) GetCheckList() []ApidefHostCheckObject`

GetCheckList returns the CheckList field if non-nil, zero value otherwise.

### GetCheckListOk

`func (o *ApidefUptimeTests) GetCheckListOk() (*[]ApidefHostCheckObject, bool)`

GetCheckListOk returns a tuple with the CheckList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckList

`func (o *ApidefUptimeTests) SetCheckList(v []ApidefHostCheckObject)`

SetCheckList sets CheckList field to given value.

### HasCheckList

`func (o *ApidefUptimeTests) HasCheckList() bool`

HasCheckList returns a boolean if a field has been set.

### SetCheckListNil

`func (o *ApidefUptimeTests) SetCheckListNil(b bool)`

 SetCheckListNil sets the value for CheckList to be an explicit nil

### UnsetCheckList
`func (o *ApidefUptimeTests) UnsetCheckList()`

UnsetCheckList ensures that no value is present for CheckList, not even an explicit nil
### GetConfig

`func (o *ApidefUptimeTests) GetConfig() ApidefUptimeTestsConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *ApidefUptimeTests) GetConfigOk() (*ApidefUptimeTestsConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *ApidefUptimeTests) SetConfig(v ApidefUptimeTestsConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *ApidefUptimeTests) HasConfig() bool`

HasConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


