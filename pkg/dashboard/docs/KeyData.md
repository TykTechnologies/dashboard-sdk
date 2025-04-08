# KeyData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiModel** | Pointer to **map[string]interface{}** |  | [optional] 
**Data** | Pointer to [**SessionState**](SessionState.md) |  | [optional] 
**KeyHash** | Pointer to **string** |  | [optional] 
**KeyId** | Pointer to **string** |  | [optional] 

## Methods

### NewKeyData

`func NewKeyData() *KeyData`

NewKeyData instantiates a new KeyData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeyDataWithDefaults

`func NewKeyDataWithDefaults() *KeyData`

NewKeyDataWithDefaults instantiates a new KeyData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiModel

`func (o *KeyData) GetApiModel() map[string]interface{}`

GetApiModel returns the ApiModel field if non-nil, zero value otherwise.

### GetApiModelOk

`func (o *KeyData) GetApiModelOk() (*map[string]interface{}, bool)`

GetApiModelOk returns a tuple with the ApiModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiModel

`func (o *KeyData) SetApiModel(v map[string]interface{})`

SetApiModel sets ApiModel field to given value.

### HasApiModel

`func (o *KeyData) HasApiModel() bool`

HasApiModel returns a boolean if a field has been set.

### GetData

`func (o *KeyData) GetData() SessionState`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *KeyData) GetDataOk() (*SessionState, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *KeyData) SetData(v SessionState)`

SetData sets Data field to given value.

### HasData

`func (o *KeyData) HasData() bool`

HasData returns a boolean if a field has been set.

### GetKeyHash

`func (o *KeyData) GetKeyHash() string`

GetKeyHash returns the KeyHash field if non-nil, zero value otherwise.

### GetKeyHashOk

`func (o *KeyData) GetKeyHashOk() (*string, bool)`

GetKeyHashOk returns a tuple with the KeyHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyHash

`func (o *KeyData) SetKeyHash(v string)`

SetKeyHash sets KeyHash field to given value.

### HasKeyHash

`func (o *KeyData) HasKeyHash() bool`

HasKeyHash returns a boolean if a field has been set.

### GetKeyId

`func (o *KeyData) GetKeyId() string`

GetKeyId returns the KeyId field if non-nil, zero value otherwise.

### GetKeyIdOk

`func (o *KeyData) GetKeyIdOk() (*string, bool)`

GetKeyIdOk returns a tuple with the KeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyId

`func (o *KeyData) SetKeyId(v string)`

SetKeyId sets KeyId field to given value.

### HasKeyId

`func (o *KeyData) HasKeyId() bool`

HasKeyId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


