# CreateWebhookRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HeaderMap** | Pointer to **map[string]string** |  | [optional] 
**Method** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**TargetPath** | Pointer to **string** |  | [optional] 

## Methods

### NewCreateWebhookRequest

`func NewCreateWebhookRequest() *CreateWebhookRequest`

NewCreateWebhookRequest instantiates a new CreateWebhookRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateWebhookRequestWithDefaults

`func NewCreateWebhookRequestWithDefaults() *CreateWebhookRequest`

NewCreateWebhookRequestWithDefaults instantiates a new CreateWebhookRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHeaderMap

`func (o *CreateWebhookRequest) GetHeaderMap() map[string]string`

GetHeaderMap returns the HeaderMap field if non-nil, zero value otherwise.

### GetHeaderMapOk

`func (o *CreateWebhookRequest) GetHeaderMapOk() (*map[string]string, bool)`

GetHeaderMapOk returns a tuple with the HeaderMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaderMap

`func (o *CreateWebhookRequest) SetHeaderMap(v map[string]string)`

SetHeaderMap sets HeaderMap field to given value.

### HasHeaderMap

`func (o *CreateWebhookRequest) HasHeaderMap() bool`

HasHeaderMap returns a boolean if a field has been set.

### SetHeaderMapNil

`func (o *CreateWebhookRequest) SetHeaderMapNil(b bool)`

 SetHeaderMapNil sets the value for HeaderMap to be an explicit nil

### UnsetHeaderMap
`func (o *CreateWebhookRequest) UnsetHeaderMap()`

UnsetHeaderMap ensures that no value is present for HeaderMap, not even an explicit nil
### GetMethod

`func (o *CreateWebhookRequest) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *CreateWebhookRequest) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *CreateWebhookRequest) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *CreateWebhookRequest) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetName

`func (o *CreateWebhookRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateWebhookRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateWebhookRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateWebhookRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTargetPath

`func (o *CreateWebhookRequest) GetTargetPath() string`

GetTargetPath returns the TargetPath field if non-nil, zero value otherwise.

### GetTargetPathOk

`func (o *CreateWebhookRequest) GetTargetPathOk() (*string, bool)`

GetTargetPathOk returns a tuple with the TargetPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetPath

`func (o *CreateWebhookRequest) SetTargetPath(v string)`

SetTargetPath sets TargetPath field to given value.

### HasTargetPath

`func (o *CreateWebhookRequest) HasTargetPath() bool`

HasTargetPath returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


