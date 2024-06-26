# ApidefEventHandlerMetaConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Events** | Pointer to [**map[string][]ApidefEventHandlerTriggerConfig**](array.md) |  | [optional] 

## Methods

### NewApidefEventHandlerMetaConfig

`func NewApidefEventHandlerMetaConfig() *ApidefEventHandlerMetaConfig`

NewApidefEventHandlerMetaConfig instantiates a new ApidefEventHandlerMetaConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApidefEventHandlerMetaConfigWithDefaults

`func NewApidefEventHandlerMetaConfigWithDefaults() *ApidefEventHandlerMetaConfig`

NewApidefEventHandlerMetaConfigWithDefaults instantiates a new ApidefEventHandlerMetaConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvents

`func (o *ApidefEventHandlerMetaConfig) GetEvents() map[string][]ApidefEventHandlerTriggerConfig`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *ApidefEventHandlerMetaConfig) GetEventsOk() (*map[string][]ApidefEventHandlerTriggerConfig, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *ApidefEventHandlerMetaConfig) SetEvents(v map[string][]ApidefEventHandlerTriggerConfig)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *ApidefEventHandlerMetaConfig) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### SetEventsNil

`func (o *ApidefEventHandlerMetaConfig) SetEventsNil(b bool)`

 SetEventsNil sets the value for Events to be an explicit nil

### UnsetEvents
`func (o *ApidefEventHandlerMetaConfig) UnsetEvents()`

UnsetEvents ensures that no value is present for Events, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


