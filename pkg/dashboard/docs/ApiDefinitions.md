# ApiDefinitions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Apis** | Pointer to [**[]ApiDefinition**](ApiDefinition.md) |  | [optional] 
**Pages** | Pointer to **int32** |  | [optional] 

## Methods

### NewApiDefinitions

`func NewApiDefinitions() *ApiDefinitions`

NewApiDefinitions instantiates a new ApiDefinitions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiDefinitionsWithDefaults

`func NewApiDefinitionsWithDefaults() *ApiDefinitions`

NewApiDefinitionsWithDefaults instantiates a new ApiDefinitions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApis

`func (o *ApiDefinitions) GetApis() []ApiDefinition`

GetApis returns the Apis field if non-nil, zero value otherwise.

### GetApisOk

`func (o *ApiDefinitions) GetApisOk() (*[]ApiDefinition, bool)`

GetApisOk returns a tuple with the Apis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApis

`func (o *ApiDefinitions) SetApis(v []ApiDefinition)`

SetApis sets Apis field to given value.

### HasApis

`func (o *ApiDefinitions) HasApis() bool`

HasApis returns a boolean if a field has been set.

### SetApisNil

`func (o *ApiDefinitions) SetApisNil(b bool)`

 SetApisNil sets the value for Apis to be an explicit nil

### UnsetApis
`func (o *ApiDefinitions) UnsetApis()`

UnsetApis ensures that no value is present for Apis, not even an explicit nil
### GetPages

`func (o *ApiDefinitions) GetPages() int32`

GetPages returns the Pages field if non-nil, zero value otherwise.

### GetPagesOk

`func (o *ApiDefinitions) GetPagesOk() (*int32, bool)`

GetPagesOk returns a tuple with the Pages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPages

`func (o *ApiDefinitions) SetPages(v int32)`

SetPages sets Pages field to given value.

### HasPages

`func (o *ApiDefinitions) HasPages() bool`

HasPages returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


