/*
Tyk Dashboard API

 ## <a name=\"introduction\"></a> Introduction  The Tyk Dashboard API offers granular, programmatic access to a centralised database of resources that your Tyk nodes can pull from. This API has a dynamic user administrative structure which means the secret key that is used to communicate with your Tyk nodes can be kept secret and access to the wider management functions can be handled on a user-by-user and organisation-by-organisation basis.  A common question around using a database-backed configuration is how to programmatically add API definitions to your Tyk nodes, the Dashboard API allows much more fine-grained, secure and multi-user access to your Tyk cluster, and should be used to manage a database-backed Tyk node.  The Tyk Dashboard API works seamlessly with the Tyk Dashboard (and the two come bundled together).  ## <a name=\"security-hierarchy\"></a> Security Hierarchy  The Dashboard API provides a more structured security layer to managing Tyk nodes.  ### Organisations, APIs and Users  With the Dashboard API and a database-backed Tyk setup, (and to an extent with file-based API setups - if diligence is used in naming and creating definitions), the following security model is applied to the management of Upstream APIs:  * **Organisations**: All APIs are *owned* by an organisation, this is designated by the 'OrgID' parameter in the API Definition. * **Users**: All users created in the Dashboard belong to an organisation (unless an exception is made for super-administrative access). * **APIs**: All APIs belong to an Organisation and only Users that belong to that organisation can see the analytics for those APIs and manage their configurations. * **API Keys**: API Keys are designated by organisation, this means an API key that has full access rights will not be allowed to access the APIs of another organisation on the same system, but can have full access to all APIs within the organisation. * **Access Rights**: Access rights are stored with the key, this enables a key to give access to multiple APIs, this is defined by the session object in the core Tyk API.  In order to use the Dashboard API, you'll need to get the 'Tyk Dashboard API Access Credentials' secret from your user profile on the Dashboard UI.  The secret you set should then be sent along as a header with each Dashboard API Request in order for it to be successful:   authorization: <your-secret>

API version: 5.7.1
Contact: support@tyk.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dashboard

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the ProxyRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProxyRequest{}

// ProxyRequest struct for ProxyRequest
type ProxyRequest struct {
	// HTTP method for the proxy request (GET, POST, PUT, DELETE, etc.)
	Method string `json:"method"`
	// Full URL for the proxy request (valid Gateway url), including scheme, host, and path
	Url string `json:"url"`
	// Headers to be sent with the proxy request
	Headers *map[string]string `json:"headers,omitempty"`
	// Body of the proxy request, typically used for POST or PUT requests
	Body map[string]interface{} `json:"body,omitempty"`
}

type _ProxyRequest ProxyRequest

// NewProxyRequest instantiates a new ProxyRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProxyRequest(method string, url string) *ProxyRequest {
	this := ProxyRequest{}
	this.Method = method
	this.Url = url
	return &this
}

// NewProxyRequestWithDefaults instantiates a new ProxyRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProxyRequestWithDefaults() *ProxyRequest {
	this := ProxyRequest{}
	return &this
}

// GetMethod returns the Method field value
func (o *ProxyRequest) GetMethod() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Method
}

// GetMethodOk returns a tuple with the Method field value
// and a boolean to check if the value has been set.
func (o *ProxyRequest) GetMethodOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Method, true
}

// SetMethod sets field value
func (o *ProxyRequest) SetMethod(v string) {
	o.Method = v
}

// GetUrl returns the Url field value
func (o *ProxyRequest) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *ProxyRequest) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *ProxyRequest) SetUrl(v string) {
	o.Url = v
}

// GetHeaders returns the Headers field value if set, zero value otherwise.
func (o *ProxyRequest) GetHeaders() map[string]string {
	if o == nil || IsNil(o.Headers) {
		var ret map[string]string
		return ret
	}
	return *o.Headers
}

// GetHeadersOk returns a tuple with the Headers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProxyRequest) GetHeadersOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Headers) {
		return nil, false
	}
	return o.Headers, true
}

// HasHeaders returns a boolean if a field has been set.
func (o *ProxyRequest) HasHeaders() bool {
	if o != nil && !IsNil(o.Headers) {
		return true
	}

	return false
}

// SetHeaders gets a reference to the given map[string]string and assigns it to the Headers field.
func (o *ProxyRequest) SetHeaders(v map[string]string) {
	o.Headers = &v
}

// GetBody returns the Body field value if set, zero value otherwise.
func (o *ProxyRequest) GetBody() map[string]interface{} {
	if o == nil || IsNil(o.Body) {
		var ret map[string]interface{}
		return ret
	}
	return o.Body
}

// GetBodyOk returns a tuple with the Body field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProxyRequest) GetBodyOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Body) {
		return map[string]interface{}{}, false
	}
	return o.Body, true
}

// HasBody returns a boolean if a field has been set.
func (o *ProxyRequest) HasBody() bool {
	if o != nil && !IsNil(o.Body) {
		return true
	}

	return false
}

// SetBody gets a reference to the given map[string]interface{} and assigns it to the Body field.
func (o *ProxyRequest) SetBody(v map[string]interface{}) {
	o.Body = v
}

func (o ProxyRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProxyRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["method"] = o.Method
	toSerialize["url"] = o.Url
	if !IsNil(o.Headers) {
		toSerialize["headers"] = o.Headers
	}
	if !IsNil(o.Body) {
		toSerialize["body"] = o.Body
	}
	return toSerialize, nil
}

func (o *ProxyRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"method",
		"url",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varProxyRequest := _ProxyRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varProxyRequest)

	if err != nil {
		return err
	}

	*o = ProxyRequest(varProxyRequest)

	return err
}

type NullableProxyRequest struct {
	value *ProxyRequest
	isSet bool
}

func (v NullableProxyRequest) Get() *ProxyRequest {
	return v.value
}

func (v *NullableProxyRequest) Set(val *ProxyRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableProxyRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableProxyRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProxyRequest(val *ProxyRequest) *NullableProxyRequest {
	return &NullableProxyRequest{value: val, isSet: true}
}

func (v NullableProxyRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProxyRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
