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

// checks if the ProviderType2 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProviderType2{}

// ProviderType2 struct for ProviderType2
type ProviderType2 struct {
	Introspection *Introspection `json:"introspection,omitempty"`
	Jwt           *JWTValidation `json:"jwt,omitempty"`
}

// NewProviderType2 instantiates a new ProviderType2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProviderType2() *ProviderType2 {
	this := ProviderType2{}
	return &this
}

// NewProviderType2WithDefaults instantiates a new ProviderType2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProviderType2WithDefaults() *ProviderType2 {
	this := ProviderType2{}
	return &this
}

// GetIntrospection returns the Introspection field value if set, zero value otherwise.
func (o *ProviderType2) GetIntrospection() Introspection {
	if o == nil || IsNil(o.Introspection) {
		var ret Introspection
		return ret
	}
	return *o.Introspection
}

// GetIntrospectionOk returns a tuple with the Introspection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProviderType2) GetIntrospectionOk() (*Introspection, bool) {
	if o == nil || IsNil(o.Introspection) {
		return nil, false
	}
	return o.Introspection, true
}

// HasIntrospection returns a boolean if a field has been set.
func (o *ProviderType2) HasIntrospection() bool {
	if o != nil && !IsNil(o.Introspection) {
		return true
	}

	return false
}

// SetIntrospection gets a reference to the given Introspection and assigns it to the Introspection field.
func (o *ProviderType2) SetIntrospection(v Introspection) {
	o.Introspection = &v
}

// GetJwt returns the Jwt field value if set, zero value otherwise.
func (o *ProviderType2) GetJwt() JWTValidation {
	if o == nil || IsNil(o.Jwt) {
		var ret JWTValidation
		return ret
	}
	return *o.Jwt
}

// GetJwtOk returns a tuple with the Jwt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProviderType2) GetJwtOk() (*JWTValidation, bool) {
	if o == nil || IsNil(o.Jwt) {
		return nil, false
	}
	return o.Jwt, true
}

// HasJwt returns a boolean if a field has been set.
func (o *ProviderType2) HasJwt() bool {
	if o != nil && !IsNil(o.Jwt) {
		return true
	}

	return false
}

// SetJwt gets a reference to the given JWTValidation and assigns it to the Jwt field.
func (o *ProviderType2) SetJwt(v JWTValidation) {
	o.Jwt = &v
}

func (o ProviderType2) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProviderType2) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Introspection) {
		toSerialize["introspection"] = o.Introspection
	}
	if !IsNil(o.Jwt) {
		toSerialize["jwt"] = o.Jwt
	}
	return toSerialize, nil
}

type NullableProviderType2 struct {
	value *ProviderType2
	isSet bool
}

func (v NullableProviderType2) Get() *ProviderType2 {
	return v.value
}

func (v *NullableProviderType2) Set(val *ProviderType2) {
	v.value = val
	v.isSet = true
}

func (v NullableProviderType2) IsSet() bool {
	return v.isSet
}

func (v *NullableProviderType2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProviderType2(val *ProviderType2) *NullableProviderType2 {
	return &NullableProviderType2{value: val, isSet: true}
}

func (v NullableProviderType2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProviderType2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
