/*
Tyk Dashboard API

## <a name=\"introduction\"></a> Introduction  The Tyk Dashboard API offers granular, programmatic access to a centralised database of resources that your Tyk nodes can pull from. This API has a dynamic user administrative structure which means the secret key that is used to communicate with your Tyk nodes can be kept secret and access to the wider management functions can be handled on a user-by-user and organisation-by-organisation basis.  A common question around using a database-backed configuration is how to programatically add API definitions to your Tyk nodes, the Dashboard API allows much more fine-grained, secure and multi-user access to your Tyk cluster, and should be used to manage a database-backed Tyk node.  The Tyk Dashboard API works seamlessly with the Tyk Dashboard (and the two come bundled together).  ## <a name=\"security-hierarchy\"></a> Security Hierarchy  The Dashboard API provides a more structured security layer to managing Tyk nodes.  ### Organisations, APIs and Users  With the Dashboard API and a database-backed Tyk setup, (and to an extent with file-based API setups - if diligence is used in naming an creating definitions), the following security model is applied to the management of Upstream APIs:  * **Organisations**: All APIs are *owned* by an organisation, this is designated by the `OrgID` parameter in the API Definition. * **Users**: All users created in the Dashboard belong to an organisation (unless an exception is made for super-administrative access). * **APIs**: All APIs belong to an Organisation and only Users that belong to that organisation can see the analytics for those APIs and manage their configurations. * **API Keys**: API Keys are designated by organisation, this means an API key that has full access rights will not be allowed to access the APIs of another organisation on the same system, but can have full access to all APIs within the organisation. * **Access Rights**: Access rights are stored with the key, this enables a key to give access to multiple APIs, this is defined by the session object in the core Tyk API.  In order to use the Dashboard API, you'll need to get the `Tyk Dashboard API Access Credentials` secret from your user profile on the Dashboard UI.  The secret you set should then be sent along as a header with each Dashboard API Request in order for it to be successful:  ``` authorization: <your-secret> ```

API version: 5.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dashboard

import (
	"encoding/json"
)

// checks if the APIDefinitionProxyTransport type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &APIDefinitionProxyTransport{}

// APIDefinitionProxyTransport struct for APIDefinitionProxyTransport
type APIDefinitionProxyTransport struct {
	ProxyUrl *string `json:"proxy_url,omitempty"`
	SslCiphers []string `json:"ssl_ciphers,omitempty"`
	SslInsecureSkipVerify *bool `json:"ssl_insecure_skip_verify,omitempty"`
	SslMinVersion *int32 `json:"ssl_min_version,omitempty"`
}

// NewAPIDefinitionProxyTransport instantiates a new APIDefinitionProxyTransport object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAPIDefinitionProxyTransport() *APIDefinitionProxyTransport {
	this := APIDefinitionProxyTransport{}
	return &this
}

// NewAPIDefinitionProxyTransportWithDefaults instantiates a new APIDefinitionProxyTransport object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAPIDefinitionProxyTransportWithDefaults() *APIDefinitionProxyTransport {
	this := APIDefinitionProxyTransport{}
	return &this
}

// GetProxyUrl returns the ProxyUrl field value if set, zero value otherwise.
func (o *APIDefinitionProxyTransport) GetProxyUrl() string {
	if o == nil || IsNil(o.ProxyUrl) {
		var ret string
		return ret
	}
	return *o.ProxyUrl
}

// GetProxyUrlOk returns a tuple with the ProxyUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APIDefinitionProxyTransport) GetProxyUrlOk() (*string, bool) {
	if o == nil || IsNil(o.ProxyUrl) {
		return nil, false
	}
	return o.ProxyUrl, true
}

// HasProxyUrl returns a boolean if a field has been set.
func (o *APIDefinitionProxyTransport) HasProxyUrl() bool {
	if o != nil && !IsNil(o.ProxyUrl) {
		return true
	}

	return false
}

// SetProxyUrl gets a reference to the given string and assigns it to the ProxyUrl field.
func (o *APIDefinitionProxyTransport) SetProxyUrl(v string) {
	o.ProxyUrl = &v
}

// GetSslCiphers returns the SslCiphers field value if set, zero value otherwise.
func (o *APIDefinitionProxyTransport) GetSslCiphers() []string {
	if o == nil || IsNil(o.SslCiphers) {
		var ret []string
		return ret
	}
	return o.SslCiphers
}

// GetSslCiphersOk returns a tuple with the SslCiphers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APIDefinitionProxyTransport) GetSslCiphersOk() ([]string, bool) {
	if o == nil || IsNil(o.SslCiphers) {
		return nil, false
	}
	return o.SslCiphers, true
}

// HasSslCiphers returns a boolean if a field has been set.
func (o *APIDefinitionProxyTransport) HasSslCiphers() bool {
	if o != nil && !IsNil(o.SslCiphers) {
		return true
	}

	return false
}

// SetSslCiphers gets a reference to the given []string and assigns it to the SslCiphers field.
func (o *APIDefinitionProxyTransport) SetSslCiphers(v []string) {
	o.SslCiphers = v
}

// GetSslInsecureSkipVerify returns the SslInsecureSkipVerify field value if set, zero value otherwise.
func (o *APIDefinitionProxyTransport) GetSslInsecureSkipVerify() bool {
	if o == nil || IsNil(o.SslInsecureSkipVerify) {
		var ret bool
		return ret
	}
	return *o.SslInsecureSkipVerify
}

// GetSslInsecureSkipVerifyOk returns a tuple with the SslInsecureSkipVerify field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APIDefinitionProxyTransport) GetSslInsecureSkipVerifyOk() (*bool, bool) {
	if o == nil || IsNil(o.SslInsecureSkipVerify) {
		return nil, false
	}
	return o.SslInsecureSkipVerify, true
}

// HasSslInsecureSkipVerify returns a boolean if a field has been set.
func (o *APIDefinitionProxyTransport) HasSslInsecureSkipVerify() bool {
	if o != nil && !IsNil(o.SslInsecureSkipVerify) {
		return true
	}

	return false
}

// SetSslInsecureSkipVerify gets a reference to the given bool and assigns it to the SslInsecureSkipVerify field.
func (o *APIDefinitionProxyTransport) SetSslInsecureSkipVerify(v bool) {
	o.SslInsecureSkipVerify = &v
}

// GetSslMinVersion returns the SslMinVersion field value if set, zero value otherwise.
func (o *APIDefinitionProxyTransport) GetSslMinVersion() int32 {
	if o == nil || IsNil(o.SslMinVersion) {
		var ret int32
		return ret
	}
	return *o.SslMinVersion
}

// GetSslMinVersionOk returns a tuple with the SslMinVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APIDefinitionProxyTransport) GetSslMinVersionOk() (*int32, bool) {
	if o == nil || IsNil(o.SslMinVersion) {
		return nil, false
	}
	return o.SslMinVersion, true
}

// HasSslMinVersion returns a boolean if a field has been set.
func (o *APIDefinitionProxyTransport) HasSslMinVersion() bool {
	if o != nil && !IsNil(o.SslMinVersion) {
		return true
	}

	return false
}

// SetSslMinVersion gets a reference to the given int32 and assigns it to the SslMinVersion field.
func (o *APIDefinitionProxyTransport) SetSslMinVersion(v int32) {
	o.SslMinVersion = &v
}

func (o APIDefinitionProxyTransport) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o APIDefinitionProxyTransport) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ProxyUrl) {
		toSerialize["proxy_url"] = o.ProxyUrl
	}
	if !IsNil(o.SslCiphers) {
		toSerialize["ssl_ciphers"] = o.SslCiphers
	}
	if !IsNil(o.SslInsecureSkipVerify) {
		toSerialize["ssl_insecure_skip_verify"] = o.SslInsecureSkipVerify
	}
	if !IsNil(o.SslMinVersion) {
		toSerialize["ssl_min_version"] = o.SslMinVersion
	}
	return toSerialize, nil
}

type NullableAPIDefinitionProxyTransport struct {
	value *APIDefinitionProxyTransport
	isSet bool
}

func (v NullableAPIDefinitionProxyTransport) Get() *APIDefinitionProxyTransport {
	return v.value
}

func (v *NullableAPIDefinitionProxyTransport) Set(val *APIDefinitionProxyTransport) {
	v.value = val
	v.isSet = true
}

func (v NullableAPIDefinitionProxyTransport) IsSet() bool {
	return v.isSet
}

func (v *NullableAPIDefinitionProxyTransport) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAPIDefinitionProxyTransport(val *APIDefinitionProxyTransport) *NullableAPIDefinitionProxyTransport {
	return &NullableAPIDefinitionProxyTransport{value: val, isSet: true}
}

func (v NullableAPIDefinitionProxyTransport) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAPIDefinitionProxyTransport) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


