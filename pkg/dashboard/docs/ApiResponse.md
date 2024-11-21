# ApiResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ID** | Pointer to **string** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Meta** | Pointer to **interface{}** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 

## Methods

### NewApiResponse

`func NewApiResponse() *ApiResponse`

NewApiResponse instantiates a new ApiResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiResponseWithDefaults

`func NewApiResponseWithDefaults() *ApiResponse`

NewApiResponseWithDefaults instantiates a new ApiResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetID

`func (o *ApiResponse) GetID() string`

GetID returns the ID field if non-nil, zero value otherwise.

### GetIDOk

`func (o *ApiResponse) GetIDOk() (*string, bool)`

GetIDOk returns a tuple with the ID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetID

`func (o *ApiResponse) SetID(v string)`

SetID sets ID field to given value.

### HasID

`func (o *ApiResponse) HasID() bool`

HasID returns a boolean if a field has been set.

### GetMessage

`func (o *ApiResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ApiResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ApiResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ApiResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetMeta

`func (o *ApiResponse) GetMeta() interface{}`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ApiResponse) GetMetaOk() (*interface{}, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ApiResponse) SetMeta(v interface{})`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ApiResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### SetMetaNil

`func (o *ApiResponse) SetMetaNil(b bool)`

 SetMetaNil sets the value for Meta to be an explicit nil

### UnsetMeta
`func (o *ApiResponse) UnsetMeta()`

UnsetMeta ensures that no value is present for Meta, not even an explicit nil
### GetStatus

`func (o *ApiResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ApiResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ApiResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ApiResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


