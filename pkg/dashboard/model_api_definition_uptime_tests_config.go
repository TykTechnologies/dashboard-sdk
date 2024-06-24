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

// checks if the APIDefinitionUptimeTestsConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &APIDefinitionUptimeTestsConfig{}

// APIDefinitionUptimeTestsConfig struct for APIDefinitionUptimeTestsConfig
type APIDefinitionUptimeTestsConfig struct {
	ExpireUtimeAfter *int64                         `json:"expire_utime_after,omitempty"`
	RecheckWait      *int64                         `json:"recheck_wait,omitempty"`
	ServiceDiscovery *ServiceDiscoveryConfiguration `json:"service_discovery,omitempty"`
}

// NewAPIDefinitionUptimeTestsConfig instantiates a new APIDefinitionUptimeTestsConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAPIDefinitionUptimeTestsConfig() *APIDefinitionUptimeTestsConfig {
	this := APIDefinitionUptimeTestsConfig{}
	return &this
}

// NewAPIDefinitionUptimeTestsConfigWithDefaults instantiates a new APIDefinitionUptimeTestsConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAPIDefinitionUptimeTestsConfigWithDefaults() *APIDefinitionUptimeTestsConfig {
	this := APIDefinitionUptimeTestsConfig{}
	return &this
}

// GetExpireUtimeAfter returns the ExpireUtimeAfter field value if set, zero value otherwise.
func (o *APIDefinitionUptimeTestsConfig) GetExpireUtimeAfter() int64 {
	if o == nil || IsNil(o.ExpireUtimeAfter) {
		var ret int64
		return ret
	}
	return *o.ExpireUtimeAfter
}

// GetExpireUtimeAfterOk returns a tuple with the ExpireUtimeAfter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APIDefinitionUptimeTestsConfig) GetExpireUtimeAfterOk() (*int64, bool) {
	if o == nil || IsNil(o.ExpireUtimeAfter) {
		return nil, false
	}
	return o.ExpireUtimeAfter, true
}

// HasExpireUtimeAfter returns a boolean if a field has been set.
func (o *APIDefinitionUptimeTestsConfig) HasExpireUtimeAfter() bool {
	if o != nil && !IsNil(o.ExpireUtimeAfter) {
		return true
	}

	return false
}

// SetExpireUtimeAfter gets a reference to the given int64 and assigns it to the ExpireUtimeAfter field.
func (o *APIDefinitionUptimeTestsConfig) SetExpireUtimeAfter(v int64) {
	o.ExpireUtimeAfter = &v
}

// GetRecheckWait returns the RecheckWait field value if set, zero value otherwise.
func (o *APIDefinitionUptimeTestsConfig) GetRecheckWait() int64 {
	if o == nil || IsNil(o.RecheckWait) {
		var ret int64
		return ret
	}
	return *o.RecheckWait
}

// GetRecheckWaitOk returns a tuple with the RecheckWait field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APIDefinitionUptimeTestsConfig) GetRecheckWaitOk() (*int64, bool) {
	if o == nil || IsNil(o.RecheckWait) {
		return nil, false
	}
	return o.RecheckWait, true
}

// HasRecheckWait returns a boolean if a field has been set.
func (o *APIDefinitionUptimeTestsConfig) HasRecheckWait() bool {
	if o != nil && !IsNil(o.RecheckWait) {
		return true
	}

	return false
}

// SetRecheckWait gets a reference to the given int64 and assigns it to the RecheckWait field.
func (o *APIDefinitionUptimeTestsConfig) SetRecheckWait(v int64) {
	o.RecheckWait = &v
}

// GetServiceDiscovery returns the ServiceDiscovery field value if set, zero value otherwise.
func (o *APIDefinitionUptimeTestsConfig) GetServiceDiscovery() ServiceDiscoveryConfiguration {
	if o == nil || IsNil(o.ServiceDiscovery) {
		var ret ServiceDiscoveryConfiguration
		return ret
	}
	return *o.ServiceDiscovery
}

// GetServiceDiscoveryOk returns a tuple with the ServiceDiscovery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APIDefinitionUptimeTestsConfig) GetServiceDiscoveryOk() (*ServiceDiscoveryConfiguration, bool) {
	if o == nil || IsNil(o.ServiceDiscovery) {
		return nil, false
	}
	return o.ServiceDiscovery, true
}

// HasServiceDiscovery returns a boolean if a field has been set.
func (o *APIDefinitionUptimeTestsConfig) HasServiceDiscovery() bool {
	if o != nil && !IsNil(o.ServiceDiscovery) {
		return true
	}

	return false
}

// SetServiceDiscovery gets a reference to the given ServiceDiscoveryConfiguration and assigns it to the ServiceDiscovery field.
func (o *APIDefinitionUptimeTestsConfig) SetServiceDiscovery(v ServiceDiscoveryConfiguration) {
	o.ServiceDiscovery = &v
}

func (o APIDefinitionUptimeTestsConfig) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o APIDefinitionUptimeTestsConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ExpireUtimeAfter) {
		toSerialize["expire_utime_after"] = o.ExpireUtimeAfter
	}
	if !IsNil(o.RecheckWait) {
		toSerialize["recheck_wait"] = o.RecheckWait
	}
	if !IsNil(o.ServiceDiscovery) {
		toSerialize["service_discovery"] = o.ServiceDiscovery
	}
	return toSerialize, nil
}

type NullableAPIDefinitionUptimeTestsConfig struct {
	value *APIDefinitionUptimeTestsConfig
	isSet bool
}

func (v NullableAPIDefinitionUptimeTestsConfig) Get() *APIDefinitionUptimeTestsConfig {
	return v.value
}

func (v *NullableAPIDefinitionUptimeTestsConfig) Set(val *APIDefinitionUptimeTestsConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableAPIDefinitionUptimeTestsConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableAPIDefinitionUptimeTestsConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAPIDefinitionUptimeTestsConfig(val *APIDefinitionUptimeTestsConfig) *NullableAPIDefinitionUptimeTestsConfig {
	return &NullableAPIDefinitionUptimeTestsConfig{value: val, isSet: true}
}

func (v NullableAPIDefinitionUptimeTestsConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAPIDefinitionUptimeTestsConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
