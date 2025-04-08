# ProxyResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StatusCode** | Pointer to **int32** | HTTP status code of the proxied response | [optional] 
**Headers** | Pointer to **map[string]string** | Headers received from the proxied response | [optional] 
**Body** | Pointer to **map[string]interface{}** | Body of the proxied response, parsed as JSON if possible | [optional] 

## Methods

### NewProxyResponse

`func NewProxyResponse() *ProxyResponse`

NewProxyResponse instantiates a new ProxyResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProxyResponseWithDefaults

`func NewProxyResponseWithDefaults() *ProxyResponse`

NewProxyResponseWithDefaults instantiates a new ProxyResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatusCode

`func (o *ProxyResponse) GetStatusCode() int32`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *ProxyResponse) GetStatusCodeOk() (*int32, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *ProxyResponse) SetStatusCode(v int32)`

SetStatusCode sets StatusCode field to given value.

### HasStatusCode

`func (o *ProxyResponse) HasStatusCode() bool`

HasStatusCode returns a boolean if a field has been set.

### GetHeaders

`func (o *ProxyResponse) GetHeaders() map[string]string`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *ProxyResponse) GetHeadersOk() (*map[string]string, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *ProxyResponse) SetHeaders(v map[string]string)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *ProxyResponse) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### GetBody

`func (o *ProxyResponse) GetBody() map[string]interface{}`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *ProxyResponse) GetBodyOk() (*map[string]interface{}, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *ProxyResponse) SetBody(v map[string]interface{})`

SetBody sets Body field to given value.

### HasBody

`func (o *ProxyResponse) HasBody() bool`

HasBody returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


