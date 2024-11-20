/*
Tyk Dashboard API

 ## <a name=\"introduction\"></a> Introduction  The Tyk Dashboard API offers granular, programmatic access to a centralised database of resources that your Tyk nodes can pull from. This API has a dynamic user administrative structure which means the secret key that is used to communicate with your Tyk nodes can be kept secret and access to the wider management functions can be handled on a user-by-user and organisation-by-organisation basis.  A common question around using a database-backed configuration is how to programmatically add API definitions to your Tyk nodes, the Dashboard API allows much more fine-grained, secure and multi-user access to your Tyk cluster, and should be used to manage a database-backed Tyk node.  The Tyk Dashboard API works seamlessly with the Tyk Dashboard (and the two come bundled together).  ## <a name=\"security-hierarchy\"></a> Security Hierarchy  The Dashboard API provides a more structured security layer to managing Tyk nodes.  ### Organisations, APIs and Users  With the Dashboard API and a database-backed Tyk setup, (and to an extent with file-based API setups - if diligence is used in naming and creating definitions), the following security model is applied to the management of Upstream APIs:  * **Organisations**: All APIs are *owned* by an organisation, this is designated by the 'OrgID' parameter in the API Definition. * **Users**: All users created in the Dashboard belong to an organisation (unless an exception is made for super-administrative access). * **APIs**: All APIs belong to an Organisation and only Users that belong to that organisation can see the analytics for those APIs and manage their configurations. * **API Keys**: API Keys are designated by organisation, this means an API key that has full access rights will not be allowed to access the APIs of another organisation on the same system, but can have full access to all APIs within the organisation. * **Access Rights**: Access rights are stored with the key, this enables a key to give access to multiple APIs, this is defined by the session object in the core Tyk API.  In order to use the Dashboard API, you'll need to get the 'Tyk Dashboard API Access Credentials' secret from your user profile on the Dashboard UI.  The secret you set should then be sent along as a header with each Dashboard API Request in order for it to be successful:   authorization: <your-secret>

API version: 5.7.0
Contact: support@tyk.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dashboard

import (
	"encoding/json"
)

// checks if the XTykAPIGateway type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &XTykAPIGateway{}

// XTykAPIGateway struct for XTykAPIGateway
type XTykAPIGateway struct {
	Info       *Info              `json:"info,omitempty"`
	Middleware NullableMiddleware `json:"middleware,omitempty"`
	Server     *Server            `json:"server,omitempty"`
	Upstream   *Upstream          `json:"upstream,omitempty"`
}

// NewXTykAPIGateway instantiates a new XTykAPIGateway object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewXTykAPIGateway() *XTykAPIGateway {
	this := XTykAPIGateway{}
	return &this
}

// NewXTykAPIGatewayWithDefaults instantiates a new XTykAPIGateway object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewXTykAPIGatewayWithDefaults() *XTykAPIGateway {
	this := XTykAPIGateway{}
	return &this
}

// GetInfo returns the Info field value if set, zero value otherwise.
func (o *XTykAPIGateway) GetInfo() Info {
	if o == nil || IsNil(o.Info) {
		var ret Info
		return ret
	}
	return *o.Info
}

// GetInfoOk returns a tuple with the Info field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XTykAPIGateway) GetInfoOk() (*Info, bool) {
	if o == nil || IsNil(o.Info) {
		return nil, false
	}
	return o.Info, true
}

// HasInfo returns a boolean if a field has been set.
func (o *XTykAPIGateway) HasInfo() bool {
	if o != nil && !IsNil(o.Info) {
		return true
	}

	return false
}

// SetInfo gets a reference to the given Info and assigns it to the Info field.
func (o *XTykAPIGateway) SetInfo(v Info) {
	o.Info = &v
}

// GetMiddleware returns the Middleware field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *XTykAPIGateway) GetMiddleware() Middleware {
	if o == nil || IsNil(o.Middleware.Get()) {
		var ret Middleware
		return ret
	}
	return *o.Middleware.Get()
}

// GetMiddlewareOk returns a tuple with the Middleware field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *XTykAPIGateway) GetMiddlewareOk() (*Middleware, bool) {
	if o == nil {
		return nil, false
	}
	return o.Middleware.Get(), o.Middleware.IsSet()
}

// HasMiddleware returns a boolean if a field has been set.
func (o *XTykAPIGateway) HasMiddleware() bool {
	if o != nil && o.Middleware.IsSet() {
		return true
	}

	return false
}

// SetMiddleware gets a reference to the given NullableMiddleware and assigns it to the Middleware field.
func (o *XTykAPIGateway) SetMiddleware(v Middleware) {
	o.Middleware.Set(&v)
}

// SetMiddlewareNil sets the value for Middleware to be an explicit nil
func (o *XTykAPIGateway) SetMiddlewareNil() {
	o.Middleware.Set(nil)
}

// UnsetMiddleware ensures that no value is present for Middleware, not even an explicit nil
func (o *XTykAPIGateway) UnsetMiddleware() {
	o.Middleware.Unset()
}

// GetServer returns the Server field value if set, zero value otherwise.
func (o *XTykAPIGateway) GetServer() Server {
	if o == nil || IsNil(o.Server) {
		var ret Server
		return ret
	}
	return *o.Server
}

// GetServerOk returns a tuple with the Server field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XTykAPIGateway) GetServerOk() (*Server, bool) {
	if o == nil || IsNil(o.Server) {
		return nil, false
	}
	return o.Server, true
}

// HasServer returns a boolean if a field has been set.
func (o *XTykAPIGateway) HasServer() bool {
	if o != nil && !IsNil(o.Server) {
		return true
	}

	return false
}

// SetServer gets a reference to the given Server and assigns it to the Server field.
func (o *XTykAPIGateway) SetServer(v Server) {
	o.Server = &v
}

// GetUpstream returns the Upstream field value if set, zero value otherwise.
func (o *XTykAPIGateway) GetUpstream() Upstream {
	if o == nil || IsNil(o.Upstream) {
		var ret Upstream
		return ret
	}
	return *o.Upstream
}

// GetUpstreamOk returns a tuple with the Upstream field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XTykAPIGateway) GetUpstreamOk() (*Upstream, bool) {
	if o == nil || IsNil(o.Upstream) {
		return nil, false
	}
	return o.Upstream, true
}

// HasUpstream returns a boolean if a field has been set.
func (o *XTykAPIGateway) HasUpstream() bool {
	if o != nil && !IsNil(o.Upstream) {
		return true
	}

	return false
}

// SetUpstream gets a reference to the given Upstream and assigns it to the Upstream field.
func (o *XTykAPIGateway) SetUpstream(v Upstream) {
	o.Upstream = &v
}

func (o XTykAPIGateway) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o XTykAPIGateway) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Info) {
		toSerialize["info"] = o.Info
	}
	if o.Middleware.IsSet() {
		toSerialize["middleware"] = o.Middleware.Get()
	}
	if !IsNil(o.Server) {
		toSerialize["server"] = o.Server
	}
	if !IsNil(o.Upstream) {
		toSerialize["upstream"] = o.Upstream
	}
	return toSerialize, nil
}

type NullableXTykAPIGateway struct {
	value *XTykAPIGateway
	isSet bool
}

func (v NullableXTykAPIGateway) Get() *XTykAPIGateway {
	return v.value
}

func (v *NullableXTykAPIGateway) Set(val *XTykAPIGateway) {
	v.value = val
	v.isSet = true
}

func (v NullableXTykAPIGateway) IsSet() bool {
	return v.isSet
}

func (v *NullableXTykAPIGateway) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableXTykAPIGateway(val *XTykAPIGateway) *NullableXTykAPIGateway {
	return &NullableXTykAPIGateway{value: val, isSet: true}
}

func (v NullableXTykAPIGateway) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableXTykAPIGateway) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
