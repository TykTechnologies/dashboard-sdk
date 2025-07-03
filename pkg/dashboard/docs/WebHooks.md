# WebHooks

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hooks** | Pointer to [**[]WebHookHandlerConf**](WebHookHandlerConf.md) |  | [optional] 
**Pages** | Pointer to **int32** |  | [optional] 

## Methods

### NewWebHooks

`func NewWebHooks() *WebHooks`

NewWebHooks instantiates a new WebHooks object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebHooksWithDefaults

`func NewWebHooksWithDefaults() *WebHooks`

NewWebHooksWithDefaults instantiates a new WebHooks object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHooks

`func (o *WebHooks) GetHooks() []WebHookHandlerConf`

GetHooks returns the Hooks field if non-nil, zero value otherwise.

### GetHooksOk

`func (o *WebHooks) GetHooksOk() (*[]WebHookHandlerConf, bool)`

GetHooksOk returns a tuple with the Hooks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHooks

`func (o *WebHooks) SetHooks(v []WebHookHandlerConf)`

SetHooks sets Hooks field to given value.

### HasHooks

`func (o *WebHooks) HasHooks() bool`

HasHooks returns a boolean if a field has been set.

### SetHooksNil

`func (o *WebHooks) SetHooksNil(b bool)`

 SetHooksNil sets the value for Hooks to be an explicit nil

### UnsetHooks
`func (o *WebHooks) UnsetHooks()`

UnsetHooks ensures that no value is present for Hooks, not even an explicit nil
### GetPages

`func (o *WebHooks) GetPages() int32`

GetPages returns the Pages field if non-nil, zero value otherwise.

### GetPagesOk

`func (o *WebHooks) GetPagesOk() (*int32, bool)`

GetPagesOk returns a tuple with the Pages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPages

`func (o *WebHooks) SetPages(v int32)`

SetPages sets Pages field to given value.

### HasPages

`func (o *WebHooks) HasPages() bool`

HasPages returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


