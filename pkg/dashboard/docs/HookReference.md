# HookReference

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventName** | Pointer to **string** |  | [optional] 
**EventTimeout** | Pointer to **int64** |  | [optional] 
**Hook** | Pointer to [**WebHookHandlerConf**](WebHookHandlerConf.md) |  | [optional] 

## Methods

### NewHookReference

`func NewHookReference() *HookReference`

NewHookReference instantiates a new HookReference object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHookReferenceWithDefaults

`func NewHookReferenceWithDefaults() *HookReference`

NewHookReferenceWithDefaults instantiates a new HookReference object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventName

`func (o *HookReference) GetEventName() string`

GetEventName returns the EventName field if non-nil, zero value otherwise.

### GetEventNameOk

`func (o *HookReference) GetEventNameOk() (*string, bool)`

GetEventNameOk returns a tuple with the EventName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventName

`func (o *HookReference) SetEventName(v string)`

SetEventName sets EventName field to given value.

### HasEventName

`func (o *HookReference) HasEventName() bool`

HasEventName returns a boolean if a field has been set.

### GetEventTimeout

`func (o *HookReference) GetEventTimeout() int64`

GetEventTimeout returns the EventTimeout field if non-nil, zero value otherwise.

### GetEventTimeoutOk

`func (o *HookReference) GetEventTimeoutOk() (*int64, bool)`

GetEventTimeoutOk returns a tuple with the EventTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTimeout

`func (o *HookReference) SetEventTimeout(v int64)`

SetEventTimeout sets EventTimeout field to given value.

### HasEventTimeout

`func (o *HookReference) HasEventTimeout() bool`

HasEventTimeout returns a boolean if a field has been set.

### GetHook

`func (o *HookReference) GetHook() WebHookHandlerConf`

GetHook returns the Hook field if non-nil, zero value otherwise.

### GetHookOk

`func (o *HookReference) GetHookOk() (*WebHookHandlerConf, bool)`

GetHookOk returns a tuple with the Hook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHook

`func (o *HookReference) SetHook(v WebHookHandlerConf)`

SetHook sets Hook field to given value.

### HasHook

`func (o *HookReference) HasHook() bool`

HasHook returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


