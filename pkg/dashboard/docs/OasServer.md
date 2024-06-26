# OasServer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Authentication** | Pointer to [**OasAuthentication**](OasAuthentication.md) |  | [optional] 
**ClientCertificates** | Pointer to [**OasClientCertificates**](OasClientCertificates.md) |  | [optional] 
**CustomDomain** | Pointer to [**OasDomain**](OasDomain.md) |  | [optional] 
**DetailedActivityLogs** | Pointer to [**OasDetailedActivityLogs**](OasDetailedActivityLogs.md) |  | [optional] 
**DetailedTracing** | Pointer to [**OasDetailedTracing**](OasDetailedTracing.md) |  | [optional] 
**EventHandlers** | Pointer to [**[]OasEventHandler**](OasEventHandler.md) |  | [optional] 
**GatewayTags** | Pointer to [**OasGatewayTags**](OasGatewayTags.md) |  | [optional] 
**ListenPath** | Pointer to [**OasListenPath**](OasListenPath.md) |  | [optional] 

## Methods

### NewOasServer

`func NewOasServer() *OasServer`

NewOasServer instantiates a new OasServer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOasServerWithDefaults

`func NewOasServerWithDefaults() *OasServer`

NewOasServerWithDefaults instantiates a new OasServer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthentication

`func (o *OasServer) GetAuthentication() OasAuthentication`

GetAuthentication returns the Authentication field if non-nil, zero value otherwise.

### GetAuthenticationOk

`func (o *OasServer) GetAuthenticationOk() (*OasAuthentication, bool)`

GetAuthenticationOk returns a tuple with the Authentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthentication

`func (o *OasServer) SetAuthentication(v OasAuthentication)`

SetAuthentication sets Authentication field to given value.

### HasAuthentication

`func (o *OasServer) HasAuthentication() bool`

HasAuthentication returns a boolean if a field has been set.

### GetClientCertificates

`func (o *OasServer) GetClientCertificates() OasClientCertificates`

GetClientCertificates returns the ClientCertificates field if non-nil, zero value otherwise.

### GetClientCertificatesOk

`func (o *OasServer) GetClientCertificatesOk() (*OasClientCertificates, bool)`

GetClientCertificatesOk returns a tuple with the ClientCertificates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientCertificates

`func (o *OasServer) SetClientCertificates(v OasClientCertificates)`

SetClientCertificates sets ClientCertificates field to given value.

### HasClientCertificates

`func (o *OasServer) HasClientCertificates() bool`

HasClientCertificates returns a boolean if a field has been set.

### GetCustomDomain

`func (o *OasServer) GetCustomDomain() OasDomain`

GetCustomDomain returns the CustomDomain field if non-nil, zero value otherwise.

### GetCustomDomainOk

`func (o *OasServer) GetCustomDomainOk() (*OasDomain, bool)`

GetCustomDomainOk returns a tuple with the CustomDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomDomain

`func (o *OasServer) SetCustomDomain(v OasDomain)`

SetCustomDomain sets CustomDomain field to given value.

### HasCustomDomain

`func (o *OasServer) HasCustomDomain() bool`

HasCustomDomain returns a boolean if a field has been set.

### GetDetailedActivityLogs

`func (o *OasServer) GetDetailedActivityLogs() OasDetailedActivityLogs`

GetDetailedActivityLogs returns the DetailedActivityLogs field if non-nil, zero value otherwise.

### GetDetailedActivityLogsOk

`func (o *OasServer) GetDetailedActivityLogsOk() (*OasDetailedActivityLogs, bool)`

GetDetailedActivityLogsOk returns a tuple with the DetailedActivityLogs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetailedActivityLogs

`func (o *OasServer) SetDetailedActivityLogs(v OasDetailedActivityLogs)`

SetDetailedActivityLogs sets DetailedActivityLogs field to given value.

### HasDetailedActivityLogs

`func (o *OasServer) HasDetailedActivityLogs() bool`

HasDetailedActivityLogs returns a boolean if a field has been set.

### GetDetailedTracing

`func (o *OasServer) GetDetailedTracing() OasDetailedTracing`

GetDetailedTracing returns the DetailedTracing field if non-nil, zero value otherwise.

### GetDetailedTracingOk

`func (o *OasServer) GetDetailedTracingOk() (*OasDetailedTracing, bool)`

GetDetailedTracingOk returns a tuple with the DetailedTracing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetailedTracing

`func (o *OasServer) SetDetailedTracing(v OasDetailedTracing)`

SetDetailedTracing sets DetailedTracing field to given value.

### HasDetailedTracing

`func (o *OasServer) HasDetailedTracing() bool`

HasDetailedTracing returns a boolean if a field has been set.

### GetEventHandlers

`func (o *OasServer) GetEventHandlers() []OasEventHandler`

GetEventHandlers returns the EventHandlers field if non-nil, zero value otherwise.

### GetEventHandlersOk

`func (o *OasServer) GetEventHandlersOk() (*[]OasEventHandler, bool)`

GetEventHandlersOk returns a tuple with the EventHandlers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventHandlers

`func (o *OasServer) SetEventHandlers(v []OasEventHandler)`

SetEventHandlers sets EventHandlers field to given value.

### HasEventHandlers

`func (o *OasServer) HasEventHandlers() bool`

HasEventHandlers returns a boolean if a field has been set.

### GetGatewayTags

`func (o *OasServer) GetGatewayTags() OasGatewayTags`

GetGatewayTags returns the GatewayTags field if non-nil, zero value otherwise.

### GetGatewayTagsOk

`func (o *OasServer) GetGatewayTagsOk() (*OasGatewayTags, bool)`

GetGatewayTagsOk returns a tuple with the GatewayTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayTags

`func (o *OasServer) SetGatewayTags(v OasGatewayTags)`

SetGatewayTags sets GatewayTags field to given value.

### HasGatewayTags

`func (o *OasServer) HasGatewayTags() bool`

HasGatewayTags returns a boolean if a field has been set.

### GetListenPath

`func (o *OasServer) GetListenPath() OasListenPath`

GetListenPath returns the ListenPath field if non-nil, zero value otherwise.

### GetListenPathOk

`func (o *OasServer) GetListenPathOk() (*OasListenPath, bool)`

GetListenPathOk returns a tuple with the ListenPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListenPath

`func (o *OasServer) SetListenPath(v OasListenPath)`

SetListenPath sets ListenPath field to given value.

### HasListenPath

`func (o *OasServer) HasListenPath() bool`

HasListenPath returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


