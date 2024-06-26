# OasMiddleware

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Global** | Pointer to [**OasGlobal**](OasGlobal.md) |  | [optional] 
**Operations** | Pointer to [**map[string]OasOperation**](OasOperation.md) |  | [optional] 

## Methods

### NewOasMiddleware

`func NewOasMiddleware() *OasMiddleware`

NewOasMiddleware instantiates a new OasMiddleware object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOasMiddlewareWithDefaults

`func NewOasMiddlewareWithDefaults() *OasMiddleware`

NewOasMiddlewareWithDefaults instantiates a new OasMiddleware object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGlobal

`func (o *OasMiddleware) GetGlobal() OasGlobal`

GetGlobal returns the Global field if non-nil, zero value otherwise.

### GetGlobalOk

`func (o *OasMiddleware) GetGlobalOk() (*OasGlobal, bool)`

GetGlobalOk returns a tuple with the Global field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobal

`func (o *OasMiddleware) SetGlobal(v OasGlobal)`

SetGlobal sets Global field to given value.

### HasGlobal

`func (o *OasMiddleware) HasGlobal() bool`

HasGlobal returns a boolean if a field has been set.

### GetOperations

`func (o *OasMiddleware) GetOperations() map[string]OasOperation`

GetOperations returns the Operations field if non-nil, zero value otherwise.

### GetOperationsOk

`func (o *OasMiddleware) GetOperationsOk() (*map[string]OasOperation, bool)`

GetOperationsOk returns a tuple with the Operations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperations

`func (o *OasMiddleware) SetOperations(v map[string]OasOperation)`

SetOperations sets Operations field to given value.

### HasOperations

`func (o *OasMiddleware) HasOperations() bool`

HasOperations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


