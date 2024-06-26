# ReturnDataStruct

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]Policy**](Policy.md) |  | [optional] 
**Pages** | Pointer to **int32** |  | [optional] 

## Methods

### NewReturnDataStruct

`func NewReturnDataStruct() *ReturnDataStruct`

NewReturnDataStruct instantiates a new ReturnDataStruct object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReturnDataStructWithDefaults

`func NewReturnDataStructWithDefaults() *ReturnDataStruct`

NewReturnDataStructWithDefaults instantiates a new ReturnDataStruct object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ReturnDataStruct) GetData() []Policy`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ReturnDataStruct) GetDataOk() (*[]Policy, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ReturnDataStruct) SetData(v []Policy)`

SetData sets Data field to given value.

### HasData

`func (o *ReturnDataStruct) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *ReturnDataStruct) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *ReturnDataStruct) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetPages

`func (o *ReturnDataStruct) GetPages() int32`

GetPages returns the Pages field if non-nil, zero value otherwise.

### GetPagesOk

`func (o *ReturnDataStruct) GetPagesOk() (*int32, bool)`

GetPagesOk returns a tuple with the Pages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPages

`func (o *ReturnDataStruct) SetPages(v int32)`

SetPages sets Pages field to given value.

### HasPages

`func (o *ReturnDataStruct) HasPages() bool`

HasPages returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


