# KeysDetailed

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Keys** | Pointer to [**[]KeyData**](KeyData.md) |  | [optional] 
**Pages** | Pointer to **int32** |  | [optional] 

## Methods

### NewKeysDetailed

`func NewKeysDetailed() *KeysDetailed`

NewKeysDetailed instantiates a new KeysDetailed object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeysDetailedWithDefaults

`func NewKeysDetailedWithDefaults() *KeysDetailed`

NewKeysDetailedWithDefaults instantiates a new KeysDetailed object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKeys

`func (o *KeysDetailed) GetKeys() []KeyData`

GetKeys returns the Keys field if non-nil, zero value otherwise.

### GetKeysOk

`func (o *KeysDetailed) GetKeysOk() (*[]KeyData, bool)`

GetKeysOk returns a tuple with the Keys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeys

`func (o *KeysDetailed) SetKeys(v []KeyData)`

SetKeys sets Keys field to given value.

### HasKeys

`func (o *KeysDetailed) HasKeys() bool`

HasKeys returns a boolean if a field has been set.

### SetKeysNil

`func (o *KeysDetailed) SetKeysNil(b bool)`

 SetKeysNil sets the value for Keys to be an explicit nil

### UnsetKeys
`func (o *KeysDetailed) UnsetKeys()`

UnsetKeys ensures that no value is present for Keys, not even an explicit nil
### GetPages

`func (o *KeysDetailed) GetPages() int32`

GetPages returns the Pages field if non-nil, zero value otherwise.

### GetPagesOk

`func (o *KeysDetailed) GetPagesOk() (*int32, bool)`

GetPagesOk returns a tuple with the Pages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPages

`func (o *KeysDetailed) SetPages(v int32)`

SetPages sets Pages field to given value.

### HasPages

`func (o *KeysDetailed) HasPages() bool`

HasPages returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


