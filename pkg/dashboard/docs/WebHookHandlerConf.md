# WebHookHandlerConf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiModel** | Pointer to **map[string]interface{}** |  | [optional] 
**EventTimeout** | Pointer to **int64** |  | [optional] 
**HeaderMap** | Pointer to **map[string]string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Method** | **string** |  | 
**Name** | Pointer to **string** |  | [optional] 
**OrgId** | Pointer to **string** |  | [optional] 
**TargetPath** | **string** |  | 
**TemplatePath** | Pointer to **string** |  | [optional] 
**WebhookId** | Pointer to **string** |  | [optional] 

## Methods

### NewWebHookHandlerConf

`func NewWebHookHandlerConf(method string, targetPath string, ) *WebHookHandlerConf`

NewWebHookHandlerConf instantiates a new WebHookHandlerConf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebHookHandlerConfWithDefaults

`func NewWebHookHandlerConfWithDefaults() *WebHookHandlerConf`

NewWebHookHandlerConfWithDefaults instantiates a new WebHookHandlerConf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiModel

`func (o *WebHookHandlerConf) GetApiModel() map[string]interface{}`

GetApiModel returns the ApiModel field if non-nil, zero value otherwise.

### GetApiModelOk

`func (o *WebHookHandlerConf) GetApiModelOk() (*map[string]interface{}, bool)`

GetApiModelOk returns a tuple with the ApiModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiModel

`func (o *WebHookHandlerConf) SetApiModel(v map[string]interface{})`

SetApiModel sets ApiModel field to given value.

### HasApiModel

`func (o *WebHookHandlerConf) HasApiModel() bool`

HasApiModel returns a boolean if a field has been set.

### GetEventTimeout

`func (o *WebHookHandlerConf) GetEventTimeout() int64`

GetEventTimeout returns the EventTimeout field if non-nil, zero value otherwise.

### GetEventTimeoutOk

`func (o *WebHookHandlerConf) GetEventTimeoutOk() (*int64, bool)`

GetEventTimeoutOk returns a tuple with the EventTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTimeout

`func (o *WebHookHandlerConf) SetEventTimeout(v int64)`

SetEventTimeout sets EventTimeout field to given value.

### HasEventTimeout

`func (o *WebHookHandlerConf) HasEventTimeout() bool`

HasEventTimeout returns a boolean if a field has been set.

### GetHeaderMap

`func (o *WebHookHandlerConf) GetHeaderMap() map[string]string`

GetHeaderMap returns the HeaderMap field if non-nil, zero value otherwise.

### GetHeaderMapOk

`func (o *WebHookHandlerConf) GetHeaderMapOk() (*map[string]string, bool)`

GetHeaderMapOk returns a tuple with the HeaderMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaderMap

`func (o *WebHookHandlerConf) SetHeaderMap(v map[string]string)`

SetHeaderMap sets HeaderMap field to given value.

### HasHeaderMap

`func (o *WebHookHandlerConf) HasHeaderMap() bool`

HasHeaderMap returns a boolean if a field has been set.

### SetHeaderMapNil

`func (o *WebHookHandlerConf) SetHeaderMapNil(b bool)`

 SetHeaderMapNil sets the value for HeaderMap to be an explicit nil

### UnsetHeaderMap
`func (o *WebHookHandlerConf) UnsetHeaderMap()`

UnsetHeaderMap ensures that no value is present for HeaderMap, not even an explicit nil
### GetId

`func (o *WebHookHandlerConf) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WebHookHandlerConf) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WebHookHandlerConf) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *WebHookHandlerConf) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMethod

`func (o *WebHookHandlerConf) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *WebHookHandlerConf) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *WebHookHandlerConf) SetMethod(v string)`

SetMethod sets Method field to given value.


### GetName

`func (o *WebHookHandlerConf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WebHookHandlerConf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WebHookHandlerConf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WebHookHandlerConf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrgId

`func (o *WebHookHandlerConf) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *WebHookHandlerConf) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *WebHookHandlerConf) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *WebHookHandlerConf) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetTargetPath

`func (o *WebHookHandlerConf) GetTargetPath() string`

GetTargetPath returns the TargetPath field if non-nil, zero value otherwise.

### GetTargetPathOk

`func (o *WebHookHandlerConf) GetTargetPathOk() (*string, bool)`

GetTargetPathOk returns a tuple with the TargetPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetPath

`func (o *WebHookHandlerConf) SetTargetPath(v string)`

SetTargetPath sets TargetPath field to given value.


### GetTemplatePath

`func (o *WebHookHandlerConf) GetTemplatePath() string`

GetTemplatePath returns the TemplatePath field if non-nil, zero value otherwise.

### GetTemplatePathOk

`func (o *WebHookHandlerConf) GetTemplatePathOk() (*string, bool)`

GetTemplatePathOk returns a tuple with the TemplatePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplatePath

`func (o *WebHookHandlerConf) SetTemplatePath(v string)`

SetTemplatePath sets TemplatePath field to given value.

### HasTemplatePath

`func (o *WebHookHandlerConf) HasTemplatePath() bool`

HasTemplatePath returns a boolean if a field has been set.

### GetWebhookId

`func (o *WebHookHandlerConf) GetWebhookId() string`

GetWebhookId returns the WebhookId field if non-nil, zero value otherwise.

### GetWebhookIdOk

`func (o *WebHookHandlerConf) GetWebhookIdOk() (*string, bool)`

GetWebhookIdOk returns a tuple with the WebhookId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookId

`func (o *WebHookHandlerConf) SetWebhookId(v string)`

SetWebhookId sets WebhookId field to given value.

### HasWebhookId

`func (o *WebHookHandlerConf) HasWebhookId() bool`

HasWebhookId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


