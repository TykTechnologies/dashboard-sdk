# ProxyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Method** | **string** | HTTP method for the proxy request (GET, POST, PUT, DELETE, etc.) | 
**Url** | **string** | Full URL for the proxy request (valid Gateway url), including scheme, host, and path | 
**Headers** | Pointer to **map[string]string** | Headers to be sent with the proxy request | [optional] 
**Body** | Pointer to **map[string]interface{}** | Body of the proxy request, typically used for POST or PUT requests | [optional] 

## Methods

### NewProxyRequest

`func NewProxyRequest(method string, url string, ) *ProxyRequest`

NewProxyRequest instantiates a new ProxyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProxyRequestWithDefaults

`func NewProxyRequestWithDefaults() *ProxyRequest`

NewProxyRequestWithDefaults instantiates a new ProxyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMethod

`func (o *ProxyRequest) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *ProxyRequest) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *ProxyRequest) SetMethod(v string)`

SetMethod sets Method field to given value.


### GetUrl

`func (o *ProxyRequest) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ProxyRequest) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ProxyRequest) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetHeaders

`func (o *ProxyRequest) GetHeaders() map[string]string`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *ProxyRequest) GetHeadersOk() (*map[string]string, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *ProxyRequest) SetHeaders(v map[string]string)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *ProxyRequest) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### GetBody

`func (o *ProxyRequest) GetBody() map[string]interface{}`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *ProxyRequest) GetBodyOk() (*map[string]interface{}, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *ProxyRequest) SetBody(v map[string]interface{})`

SetBody sets Body field to given value.

### HasBody

`func (o *ProxyRequest) HasBody() bool`

HasBody returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


