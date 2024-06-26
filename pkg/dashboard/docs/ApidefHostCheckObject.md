# ApidefHostCheckObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Body** | Pointer to **string** |  | [optional] 
**Commands** | Pointer to [**[]ApidefCheckCommand**](ApidefCheckCommand.md) |  | [optional] 
**EnableProxyProtocol** | Pointer to **bool** |  | [optional] 
**Headers** | Pointer to **map[string]string** |  | [optional] 
**Method** | Pointer to **string** |  | [optional] 
**Protocol** | Pointer to **string** |  | [optional] 
**Timeout** | Pointer to **int32** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 

## Methods

### NewApidefHostCheckObject

`func NewApidefHostCheckObject() *ApidefHostCheckObject`

NewApidefHostCheckObject instantiates a new ApidefHostCheckObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApidefHostCheckObjectWithDefaults

`func NewApidefHostCheckObjectWithDefaults() *ApidefHostCheckObject`

NewApidefHostCheckObjectWithDefaults instantiates a new ApidefHostCheckObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBody

`func (o *ApidefHostCheckObject) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *ApidefHostCheckObject) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *ApidefHostCheckObject) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *ApidefHostCheckObject) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetCommands

`func (o *ApidefHostCheckObject) GetCommands() []ApidefCheckCommand`

GetCommands returns the Commands field if non-nil, zero value otherwise.

### GetCommandsOk

`func (o *ApidefHostCheckObject) GetCommandsOk() (*[]ApidefCheckCommand, bool)`

GetCommandsOk returns a tuple with the Commands field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommands

`func (o *ApidefHostCheckObject) SetCommands(v []ApidefCheckCommand)`

SetCommands sets Commands field to given value.

### HasCommands

`func (o *ApidefHostCheckObject) HasCommands() bool`

HasCommands returns a boolean if a field has been set.

### SetCommandsNil

`func (o *ApidefHostCheckObject) SetCommandsNil(b bool)`

 SetCommandsNil sets the value for Commands to be an explicit nil

### UnsetCommands
`func (o *ApidefHostCheckObject) UnsetCommands()`

UnsetCommands ensures that no value is present for Commands, not even an explicit nil
### GetEnableProxyProtocol

`func (o *ApidefHostCheckObject) GetEnableProxyProtocol() bool`

GetEnableProxyProtocol returns the EnableProxyProtocol field if non-nil, zero value otherwise.

### GetEnableProxyProtocolOk

`func (o *ApidefHostCheckObject) GetEnableProxyProtocolOk() (*bool, bool)`

GetEnableProxyProtocolOk returns a tuple with the EnableProxyProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableProxyProtocol

`func (o *ApidefHostCheckObject) SetEnableProxyProtocol(v bool)`

SetEnableProxyProtocol sets EnableProxyProtocol field to given value.

### HasEnableProxyProtocol

`func (o *ApidefHostCheckObject) HasEnableProxyProtocol() bool`

HasEnableProxyProtocol returns a boolean if a field has been set.

### GetHeaders

`func (o *ApidefHostCheckObject) GetHeaders() map[string]string`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *ApidefHostCheckObject) GetHeadersOk() (*map[string]string, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *ApidefHostCheckObject) SetHeaders(v map[string]string)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *ApidefHostCheckObject) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### SetHeadersNil

`func (o *ApidefHostCheckObject) SetHeadersNil(b bool)`

 SetHeadersNil sets the value for Headers to be an explicit nil

### UnsetHeaders
`func (o *ApidefHostCheckObject) UnsetHeaders()`

UnsetHeaders ensures that no value is present for Headers, not even an explicit nil
### GetMethod

`func (o *ApidefHostCheckObject) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *ApidefHostCheckObject) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *ApidefHostCheckObject) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *ApidefHostCheckObject) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetProtocol

`func (o *ApidefHostCheckObject) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *ApidefHostCheckObject) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *ApidefHostCheckObject) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *ApidefHostCheckObject) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetTimeout

`func (o *ApidefHostCheckObject) GetTimeout() int32`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *ApidefHostCheckObject) GetTimeoutOk() (*int32, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *ApidefHostCheckObject) SetTimeout(v int32)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *ApidefHostCheckObject) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### GetUrl

`func (o *ApidefHostCheckObject) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ApidefHostCheckObject) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ApidefHostCheckObject) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *ApidefHostCheckObject) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


