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

// checks if the UptimeTests type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UptimeTests{}

// UptimeTests struct for UptimeTests
type UptimeTests struct {
	CheckList []HostCheckObject  `json:"check_list,omitempty"`
	Config    *UptimeTestsConfig `json:"config,omitempty"`
}

// NewUptimeTests instantiates a new UptimeTests object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUptimeTests() *UptimeTests {
	this := UptimeTests{}
	return &this
}

// NewUptimeTestsWithDefaults instantiates a new UptimeTests object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUptimeTestsWithDefaults() *UptimeTests {
	this := UptimeTests{}
	return &this
}

// GetCheckList returns the CheckList field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UptimeTests) GetCheckList() []HostCheckObject {
	if o == nil {
		var ret []HostCheckObject
		return ret
	}
	return o.CheckList
}

// GetCheckListOk returns a tuple with the CheckList field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UptimeTests) GetCheckListOk() ([]HostCheckObject, bool) {
	if o == nil || IsNil(o.CheckList) {
		return nil, false
	}
	return o.CheckList, true
}

// HasCheckList returns a boolean if a field has been set.
func (o *UptimeTests) HasCheckList() bool {
	if o != nil && !IsNil(o.CheckList) {
		return true
	}

	return false
}

// SetCheckList gets a reference to the given []HostCheckObject and assigns it to the CheckList field.
func (o *UptimeTests) SetCheckList(v []HostCheckObject) {
	o.CheckList = v
}

// GetConfig returns the Config field value if set, zero value otherwise.
func (o *UptimeTests) GetConfig() UptimeTestsConfig {
	if o == nil || IsNil(o.Config) {
		var ret UptimeTestsConfig
		return ret
	}
	return *o.Config
}

// GetConfigOk returns a tuple with the Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UptimeTests) GetConfigOk() (*UptimeTestsConfig, bool) {
	if o == nil || IsNil(o.Config) {
		return nil, false
	}
	return o.Config, true
}

// HasConfig returns a boolean if a field has been set.
func (o *UptimeTests) HasConfig() bool {
	if o != nil && !IsNil(o.Config) {
		return true
	}

	return false
}

// SetConfig gets a reference to the given UptimeTestsConfig and assigns it to the Config field.
func (o *UptimeTests) SetConfig(v UptimeTestsConfig) {
	o.Config = &v
}

func (o UptimeTests) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UptimeTests) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.CheckList != nil {
		toSerialize["check_list"] = o.CheckList
	}
	if !IsNil(o.Config) {
		toSerialize["config"] = o.Config
	}
	return toSerialize, nil
}

type NullableUptimeTests struct {
	value *UptimeTests
	isSet bool
}

func (v NullableUptimeTests) Get() *UptimeTests {
	return v.value
}

func (v *NullableUptimeTests) Set(val *UptimeTests) {
	v.value = val
	v.isSet = true
}

func (v NullableUptimeTests) IsSet() bool {
	return v.isSet
}

func (v *NullableUptimeTests) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUptimeTests(val *UptimeTests) *NullableUptimeTests {
	return &NullableUptimeTests{value: val, isSet: true}
}

func (v NullableUptimeTests) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUptimeTests) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
