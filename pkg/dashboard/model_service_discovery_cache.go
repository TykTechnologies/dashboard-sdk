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

// checks if the ServiceDiscoveryCache type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServiceDiscoveryCache{}

// ServiceDiscoveryCache struct for ServiceDiscoveryCache
type ServiceDiscoveryCache struct {
	Enabled *bool  `json:"enabled,omitempty"`
	Timeout *int64 `json:"timeout,omitempty"`
}

// NewServiceDiscoveryCache instantiates a new ServiceDiscoveryCache object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceDiscoveryCache() *ServiceDiscoveryCache {
	this := ServiceDiscoveryCache{}
	return &this
}

// NewServiceDiscoveryCacheWithDefaults instantiates a new ServiceDiscoveryCache object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceDiscoveryCacheWithDefaults() *ServiceDiscoveryCache {
	this := ServiceDiscoveryCache{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *ServiceDiscoveryCache) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceDiscoveryCache) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *ServiceDiscoveryCache) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *ServiceDiscoveryCache) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetTimeout returns the Timeout field value if set, zero value otherwise.
func (o *ServiceDiscoveryCache) GetTimeout() int64 {
	if o == nil || IsNil(o.Timeout) {
		var ret int64
		return ret
	}
	return *o.Timeout
}

// GetTimeoutOk returns a tuple with the Timeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceDiscoveryCache) GetTimeoutOk() (*int64, bool) {
	if o == nil || IsNil(o.Timeout) {
		return nil, false
	}
	return o.Timeout, true
}

// HasTimeout returns a boolean if a field has been set.
func (o *ServiceDiscoveryCache) HasTimeout() bool {
	if o != nil && !IsNil(o.Timeout) {
		return true
	}

	return false
}

// SetTimeout gets a reference to the given int64 and assigns it to the Timeout field.
func (o *ServiceDiscoveryCache) SetTimeout(v int64) {
	o.Timeout = &v
}

func (o ServiceDiscoveryCache) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServiceDiscoveryCache) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.Timeout) {
		toSerialize["timeout"] = o.Timeout
	}
	return toSerialize, nil
}

type NullableServiceDiscoveryCache struct {
	value *ServiceDiscoveryCache
	isSet bool
}

func (v NullableServiceDiscoveryCache) Get() *ServiceDiscoveryCache {
	return v.value
}

func (v *NullableServiceDiscoveryCache) Set(val *ServiceDiscoveryCache) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceDiscoveryCache) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceDiscoveryCache) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceDiscoveryCache(val *ServiceDiscoveryCache) *NullableServiceDiscoveryCache {
	return &NullableServiceDiscoveryCache{value: val, isSet: true}
}

func (v NullableServiceDiscoveryCache) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceDiscoveryCache) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
