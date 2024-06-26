# Keys

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**AllKeys**](AllKeys.md) |  | [optional] 
**Pages** | Pointer to **int32** |  | [optional] 

## Methods

### NewKeys

`func NewKeys() *Keys`

NewKeys instantiates a new Keys object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeysWithDefaults

`func NewKeysWithDefaults() *Keys`

NewKeysWithDefaults instantiates a new Keys object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *Keys) GetData() AllKeys`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *Keys) GetDataOk() (*AllKeys, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *Keys) SetData(v AllKeys)`

SetData sets Data field to given value.

### HasData

`func (o *Keys) HasData() bool`

HasData returns a boolean if a field has been set.

### GetPages

`func (o *Keys) GetPages() int32`

GetPages returns the Pages field if non-nil, zero value otherwise.

### GetPagesOk

`func (o *Keys) GetPagesOk() (*int32, bool)`

GetPagesOk returns a tuple with the Pages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPages

`func (o *Keys) SetPages(v int32)`

SetPages sets Pages field to given value.

### HasPages

`func (o *Keys) HasPages() bool`

HasPages returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


