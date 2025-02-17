# ErrorResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** | Error status, typically \&quot;Error\&quot; | [optional] 
**Message** | Pointer to **string** | Detailed error message explaining what went wrong | [optional] 
**Meta** | Pointer to **map[string]interface{}** | Additional metadata about the error, if available | [optional] 

## Methods

### NewErrorResponse

`func NewErrorResponse() *ErrorResponse`

NewErrorResponse instantiates a new ErrorResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorResponseWithDefaults

`func NewErrorResponseWithDefaults() *ErrorResponse`

NewErrorResponseWithDefaults instantiates a new ErrorResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *ErrorResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ErrorResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ErrorResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ErrorResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetMessage

`func (o *ErrorResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ErrorResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ErrorResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ErrorResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetMeta

`func (o *ErrorResponse) GetMeta() map[string]interface{}`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ErrorResponse) GetMetaOk() (*map[string]interface{}, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ErrorResponse) SetMeta(v map[string]interface{})`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ErrorResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### SetMetaNil

`func (o *ErrorResponse) SetMetaNil(b bool)`

 SetMetaNil sets the value for Meta to be an explicit nil

### UnsetMeta
`func (o *ErrorResponse) UnsetMeta()`

UnsetMeta ensures that no value is present for Meta, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


