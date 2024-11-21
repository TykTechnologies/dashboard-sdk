# ApiDefinitionsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Apis** | Pointer to [**[]ApiDefinitionWrapper**](ApiDefinitionWrapper.md) |  | [optional] 
**Pages** | Pointer to **int32** |  | [optional] 

## Methods

### NewApiDefinitionsResponse

`func NewApiDefinitionsResponse() *ApiDefinitionsResponse`

NewApiDefinitionsResponse instantiates a new ApiDefinitionsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiDefinitionsResponseWithDefaults

`func NewApiDefinitionsResponseWithDefaults() *ApiDefinitionsResponse`

NewApiDefinitionsResponseWithDefaults instantiates a new ApiDefinitionsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApis

`func (o *ApiDefinitionsResponse) GetApis() []ApiDefinitionWrapper`

GetApis returns the Apis field if non-nil, zero value otherwise.

### GetApisOk

`func (o *ApiDefinitionsResponse) GetApisOk() (*[]ApiDefinitionWrapper, bool)`

GetApisOk returns a tuple with the Apis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApis

`func (o *ApiDefinitionsResponse) SetApis(v []ApiDefinitionWrapper)`

SetApis sets Apis field to given value.

### HasApis

`func (o *ApiDefinitionsResponse) HasApis() bool`

HasApis returns a boolean if a field has been set.

### SetApisNil

`func (o *ApiDefinitionsResponse) SetApisNil(b bool)`

 SetApisNil sets the value for Apis to be an explicit nil

### UnsetApis
`func (o *ApiDefinitionsResponse) UnsetApis()`

UnsetApis ensures that no value is present for Apis, not even an explicit nil
### GetPages

`func (o *ApiDefinitionsResponse) GetPages() int32`

GetPages returns the Pages field if non-nil, zero value otherwise.

### GetPagesOk

`func (o *ApiDefinitionsResponse) GetPagesOk() (*int32, bool)`

GetPagesOk returns a tuple with the Pages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPages

`func (o *ApiDefinitionsResponse) SetPages(v int32)`

SetPages sets Pages field to given value.

### HasPages

`func (o *ApiDefinitionsResponse) HasPages() bool`

HasPages returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


