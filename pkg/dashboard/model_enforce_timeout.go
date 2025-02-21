/*
Tyk Dashboard API

 ## <a name=\"introduction\"></a> Introduction  The Tyk Dashboard API offers granular, programmatic access to a centralised database of resources that your Tyk nodes can pull from. This API has a dynamic user administrative structure which means the secret key that is used to communicate with your Tyk nodes can be kept secret and access to the wider management functions can be handled on a user-by-user and organisation-by-organisation basis.  A common question around using a database-backed configuration is how to programmatically add API definitions to your Tyk nodes, the Dashboard API allows much more fine-grained, secure and multi-user access to your Tyk cluster, and should be used to manage a database-backed Tyk node.  The Tyk Dashboard API works seamlessly with the Tyk Dashboard (and the two come bundled together).  ## <a name=\"security-hierarchy\"></a> Security Hierarchy  The Dashboard API provides a more structured security layer to managing Tyk nodes.  ### Organisations, APIs and Users  With the Dashboard API and a database-backed Tyk setup, (and to an extent with file-based API setups - if diligence is used in naming and creating definitions), the following security model is applied to the management of Upstream APIs:  * **Organisations**: All APIs are *owned* by an organisation, this is designated by the 'OrgID' parameter in the API Definition. * **Users**: All users created in the Dashboard belong to an organisation (unless an exception is made for super-administrative access). * **APIs**: All APIs belong to an Organisation and only Users that belong to that organisation can see the analytics for those APIs and manage their configurations. * **API Keys**: API Keys are designated by organisation, this means an API key that has full access rights will not be allowed to access the APIs of another organisation on the same system, but can have full access to all APIs within the organisation. * **Access Rights**: Access rights are stored with the key, this enables a key to give access to multiple APIs, this is defined by the session object in the core Tyk API.  In order to use the Dashboard API, you'll need to get the 'Tyk Dashboard API Access Credentials' secret from your user profile on the Dashboard UI.  The secret you set should then be sent along as a header with each Dashboard API Request in order for it to be successful:   authorization: <your-secret>

API version: 5.7.1
Contact: support@tyk.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dashboard

import (
	"encoding/json"
)

// checks if the EnforceTimeout type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EnforceTimeout{}

// EnforceTimeout struct for EnforceTimeout
type EnforceTimeout struct {
	Enabled *bool  `json:"enabled,omitempty"`
	Value   *int32 `json:"value,omitempty"`
}

// NewEnforceTimeout instantiates a new EnforceTimeout object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnforceTimeout() *EnforceTimeout {
	this := EnforceTimeout{}
	return &this
}

// NewEnforceTimeoutWithDefaults instantiates a new EnforceTimeout object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnforceTimeoutWithDefaults() *EnforceTimeout {
	this := EnforceTimeout{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *EnforceTimeout) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnforceTimeout) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *EnforceTimeout) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *EnforceTimeout) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *EnforceTimeout) GetValue() int32 {
	if o == nil || IsNil(o.Value) {
		var ret int32
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnforceTimeout) GetValueOk() (*int32, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *EnforceTimeout) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given int32 and assigns it to the Value field.
func (o *EnforceTimeout) SetValue(v int32) {
	o.Value = &v
}

func (o EnforceTimeout) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EnforceTimeout) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableEnforceTimeout struct {
	value *EnforceTimeout
	isSet bool
}

func (v NullableEnforceTimeout) Get() *EnforceTimeout {
	return v.value
}

func (v *NullableEnforceTimeout) Set(val *EnforceTimeout) {
	v.value = val
	v.isSet = true
}

func (v NullableEnforceTimeout) IsSet() bool {
	return v.isSet
}

func (v *NullableEnforceTimeout) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnforceTimeout(val *EnforceTimeout) *NullableEnforceTimeout {
	return &NullableEnforceTimeout{value: val, isSet: true}
}

func (v NullableEnforceTimeout) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnforceTimeout) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
