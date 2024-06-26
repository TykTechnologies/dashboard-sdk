# OasXTykAPIGateway

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Info** | Pointer to [**OasInfo**](OasInfo.md) |  | [optional] 
**Middleware** | Pointer to [**OasMiddleware**](OasMiddleware.md) |  | [optional] 
**Server** | Pointer to [**OasServer**](OasServer.md) |  | [optional] 
**Upstream** | Pointer to [**OasUpstream**](OasUpstream.md) |  | [optional] 

## Methods

### NewOasXTykAPIGateway

`func NewOasXTykAPIGateway() *OasXTykAPIGateway`

NewOasXTykAPIGateway instantiates a new OasXTykAPIGateway object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOasXTykAPIGatewayWithDefaults

`func NewOasXTykAPIGatewayWithDefaults() *OasXTykAPIGateway`

NewOasXTykAPIGatewayWithDefaults instantiates a new OasXTykAPIGateway object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInfo

`func (o *OasXTykAPIGateway) GetInfo() OasInfo`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *OasXTykAPIGateway) GetInfoOk() (*OasInfo, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *OasXTykAPIGateway) SetInfo(v OasInfo)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *OasXTykAPIGateway) HasInfo() bool`

HasInfo returns a boolean if a field has been set.

### GetMiddleware

`func (o *OasXTykAPIGateway) GetMiddleware() OasMiddleware`

GetMiddleware returns the Middleware field if non-nil, zero value otherwise.

### GetMiddlewareOk

`func (o *OasXTykAPIGateway) GetMiddlewareOk() (*OasMiddleware, bool)`

GetMiddlewareOk returns a tuple with the Middleware field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMiddleware

`func (o *OasXTykAPIGateway) SetMiddleware(v OasMiddleware)`

SetMiddleware sets Middleware field to given value.

### HasMiddleware

`func (o *OasXTykAPIGateway) HasMiddleware() bool`

HasMiddleware returns a boolean if a field has been set.

### GetServer

`func (o *OasXTykAPIGateway) GetServer() OasServer`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *OasXTykAPIGateway) GetServerOk() (*OasServer, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *OasXTykAPIGateway) SetServer(v OasServer)`

SetServer sets Server field to given value.

### HasServer

`func (o *OasXTykAPIGateway) HasServer() bool`

HasServer returns a boolean if a field has been set.

### GetUpstream

`func (o *OasXTykAPIGateway) GetUpstream() OasUpstream`

GetUpstream returns the Upstream field if non-nil, zero value otherwise.

### GetUpstreamOk

`func (o *OasXTykAPIGateway) GetUpstreamOk() (*OasUpstream, bool)`

GetUpstreamOk returns a tuple with the Upstream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpstream

`func (o *OasXTykAPIGateway) SetUpstream(v OasUpstream)`

SetUpstream sets Upstream field to given value.

### HasUpstream

`func (o *OasXTykAPIGateway) HasUpstream() bool`

HasUpstream returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


