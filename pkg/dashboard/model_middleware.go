/*
Tyk Dashboard API

 ## <a name=\"introduction\"></a> Introduction  The Tyk Dashboard API offers granular, programmatic access to a centralised database of resources that your Tyk nodes can pull from. This API has a dynamic user administrative structure which means the secret key that is used to communicate with your Tyk nodes can be kept secret and access to the wider management functions can be handled on a user-by-user and organisation-by-organisation basis.  A common question around using a database-backed configuration is how to programmatically add API definitions to your Tyk nodes, the Dashboard API allows much more fine-grained, secure and multi-user access to your Tyk cluster, and should be used to manage a database-backed Tyk node.  The Tyk Dashboard API works seamlessly with the Tyk Dashboard (and the two come bundled together).  ## <a name=\"security-hierarchy\"></a> Security Hierarchy  The Dashboard API provides a more structured security layer to managing Tyk nodes.  ### Organisations, APIs and Users  With the Dashboard API and a database-backed Tyk setup, (and to an extent with file-based API setups - if diligence is used in naming and creating definitions), the following security model is applied to the management of Upstream APIs:  * **Organisations**: All APIs are *owned* by an organisation, this is designated by the 'OrgID' parameter in the API Definition. * **Users**: All users created in the Dashboard belong to an organisation (unless an exception is made for super-administrative access). * **APIs**: All APIs belong to an Organisation and only Users that belong to that organisation can see the analytics for those APIs and manage their configurations. * **API Keys**: API Keys are designated by organisation, this means an API key that has full access rights will not be allowed to access the APIs of another organisation on the same system, but can have full access to all APIs within the organisation. * **Access Rights**: Access rights are stored with the key, this enables a key to give access to multiple APIs, this is defined by the session object in the core Tyk API.  In order to use the Dashboard API, you'll need to get the 'Tyk Dashboard API Access Credentials' secret from your user profile on the Dashboard UI.  The secret you set should then be sent along as a header with each Dashboard API Request in order for it to be successful:   authorization: <your-secret>

API version: 5.6.0
Contact: support@tyk.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dashboard

import (
	"encoding/json"
)

// checks if the Middleware type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Middleware{}

// Middleware struct for Middleware
type Middleware struct {
	Global     NullableGlobal        `json:"global,omitempty"`
	Operations *map[string]Operation `json:"operations,omitempty"`
}

// NewMiddleware instantiates a new Middleware object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMiddleware() *Middleware {
	this := Middleware{}
	return &this
}

// NewMiddlewareWithDefaults instantiates a new Middleware object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMiddlewareWithDefaults() *Middleware {
	this := Middleware{}
	return &this
}

// GetGlobal returns the Global field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Middleware) GetGlobal() Global {
	if o == nil || IsNil(o.Global.Get()) {
		var ret Global
		return ret
	}
	return *o.Global.Get()
}

// GetGlobalOk returns a tuple with the Global field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Middleware) GetGlobalOk() (*Global, bool) {
	if o == nil {
		return nil, false
	}
	return o.Global.Get(), o.Global.IsSet()
}

// HasGlobal returns a boolean if a field has been set.
func (o *Middleware) HasGlobal() bool {
	if o != nil && o.Global.IsSet() {
		return true
	}

	return false
}

// SetGlobal gets a reference to the given NullableGlobal and assigns it to the Global field.
func (o *Middleware) SetGlobal(v Global) {
	o.Global.Set(&v)
}

// SetGlobalNil sets the value for Global to be an explicit nil
func (o *Middleware) SetGlobalNil() {
	o.Global.Set(nil)
}

// UnsetGlobal ensures that no value is present for Global, not even an explicit nil
func (o *Middleware) UnsetGlobal() {
	o.Global.Unset()
}

// GetOperations returns the Operations field value if set, zero value otherwise.
func (o *Middleware) GetOperations() map[string]Operation {
	if o == nil || IsNil(o.Operations) {
		var ret map[string]Operation
		return ret
	}
	return *o.Operations
}

// GetOperationsOk returns a tuple with the Operations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Middleware) GetOperationsOk() (*map[string]Operation, bool) {
	if o == nil || IsNil(o.Operations) {
		return nil, false
	}
	return o.Operations, true
}

// HasOperations returns a boolean if a field has been set.
func (o *Middleware) HasOperations() bool {
	if o != nil && !IsNil(o.Operations) {
		return true
	}

	return false
}

// SetOperations gets a reference to the given map[string]Operation and assigns it to the Operations field.
func (o *Middleware) SetOperations(v map[string]Operation) {
	o.Operations = &v
}

func (o Middleware) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Middleware) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Global.IsSet() {
		toSerialize["global"] = o.Global.Get()
	}
	if !IsNil(o.Operations) {
		toSerialize["operations"] = o.Operations
	}
	return toSerialize, nil
}

type NullableMiddleware struct {
	value *Middleware
	isSet bool
}

func (v NullableMiddleware) Get() *Middleware {
	return v.value
}

func (v *NullableMiddleware) Set(val *Middleware) {
	v.value = val
	v.isSet = true
}

func (v NullableMiddleware) IsSet() bool {
	return v.isSet
}

func (v *NullableMiddleware) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMiddleware(val *Middleware) *NullableMiddleware {
	return &NullableMiddleware{value: val, isSet: true}
}

func (v NullableMiddleware) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMiddleware) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
