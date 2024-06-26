# ApidefEndpointMethodMeta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **string** |  | [optional] 
**Code** | Pointer to **int32** |  | [optional] 
**Data** | Pointer to **string** |  | [optional] 
**Headers** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewApidefEndpointMethodMeta

`func NewApidefEndpointMethodMeta() *ApidefEndpointMethodMeta`

NewApidefEndpointMethodMeta instantiates a new ApidefEndpointMethodMeta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApidefEndpointMethodMetaWithDefaults

`func NewApidefEndpointMethodMetaWithDefaults() *ApidefEndpointMethodMeta`

NewApidefEndpointMethodMetaWithDefaults instantiates a new ApidefEndpointMethodMeta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *ApidefEndpointMethodMeta) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *ApidefEndpointMethodMeta) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *ApidefEndpointMethodMeta) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *ApidefEndpointMethodMeta) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetCode

`func (o *ApidefEndpointMethodMeta) GetCode() int32`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ApidefEndpointMethodMeta) GetCodeOk() (*int32, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ApidefEndpointMethodMeta) SetCode(v int32)`

SetCode sets Code field to given value.

### HasCode

`func (o *ApidefEndpointMethodMeta) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetData

`func (o *ApidefEndpointMethodMeta) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ApidefEndpointMethodMeta) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ApidefEndpointMethodMeta) SetData(v string)`

SetData sets Data field to given value.

### HasData

`func (o *ApidefEndpointMethodMeta) HasData() bool`

HasData returns a boolean if a field has been set.

### GetHeaders

`func (o *ApidefEndpointMethodMeta) GetHeaders() map[string]string`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *ApidefEndpointMethodMeta) GetHeadersOk() (*map[string]string, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *ApidefEndpointMethodMeta) SetHeaders(v map[string]string)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *ApidefEndpointMethodMeta) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### SetHeadersNil

`func (o *ApidefEndpointMethodMeta) SetHeadersNil(b bool)`

 SetHeadersNil sets the value for Headers to be an explicit nil

### UnsetHeaders
`func (o *ApidefEndpointMethodMeta) UnsetHeaders()`

UnsetHeaders ensures that no value is present for Headers, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


