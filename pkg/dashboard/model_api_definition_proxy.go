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

// checks if the APIDefinitionProxy type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &APIDefinitionProxy{}

// APIDefinitionProxy struct for APIDefinitionProxy
type APIDefinitionProxy struct {
	CheckHostAgainstUptimeTests *bool `json:"check_host_against_uptime_tests,omitempty"`
	DisableStripSlash *bool `json:"disable_strip_slash,omitempty"`
	EnableLoadBalancing *bool `json:"enable_load_balancing,omitempty"`
	ListenPath *string `json:"listen_path,omitempty"`
	PreserveHostHeader *bool `json:"preserve_host_header,omitempty"`
	ServiceDiscovery *ServiceDiscoveryConfiguration `json:"service_discovery,omitempty"`
	StripListenPath *bool `json:"strip_listen_path,omitempty"`
	TargetList []string `json:"target_list,omitempty"`
	TargetUrl *string `json:"target_url,omitempty"`
	Transport *APIDefinitionProxyTransport `json:"transport,omitempty"`
}

// NewAPIDefinitionProxy instantiates a new APIDefinitionProxy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAPIDefinitionProxy() *APIDefinitionProxy {
	this := APIDefinitionProxy{}
	return &this
}

// NewAPIDefinitionProxyWithDefaults instantiates a new APIDefinitionProxy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAPIDefinitionProxyWithDefaults() *APIDefinitionProxy {
	this := APIDefinitionProxy{}
	return &this
}

// GetCheckHostAgainstUptimeTests returns the CheckHostAgainstUptimeTests field value if set, zero value otherwise.
func (o *APIDefinitionProxy) GetCheckHostAgainstUptimeTests() bool {
	if o == nil || IsNil(o.CheckHostAgainstUptimeTests) {
		var ret bool
		return ret
	}
	return *o.CheckHostAgainstUptimeTests
}

// GetCheckHostAgainstUptimeTestsOk returns a tuple with the CheckHostAgainstUptimeTests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APIDefinitionProxy) GetCheckHostAgainstUptimeTestsOk() (*bool, bool) {
	if o == nil || IsNil(o.CheckHostAgainstUptimeTests) {
		return nil, false
	}
	return o.CheckHostAgainstUptimeTests, true
}

// HasCheckHostAgainstUptimeTests returns a boolean if a field has been set.
func (o *APIDefinitionProxy) HasCheckHostAgainstUptimeTests() bool {
	if o != nil && !IsNil(o.CheckHostAgainstUptimeTests) {
		return true
	}

	return false
}

// SetCheckHostAgainstUptimeTests gets a reference to the given bool and assigns it to the CheckHostAgainstUptimeTests field.
func (o *APIDefinitionProxy) SetCheckHostAgainstUptimeTests(v bool) {
	o.CheckHostAgainstUptimeTests = &v
}

// GetDisableStripSlash returns the DisableStripSlash field value if set, zero value otherwise.
func (o *APIDefinitionProxy) GetDisableStripSlash() bool {
	if o == nil || IsNil(o.DisableStripSlash) {
		var ret bool
		return ret
	}
	return *o.DisableStripSlash
}

// GetDisableStripSlashOk returns a tuple with the DisableStripSlash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APIDefinitionProxy) GetDisableStripSlashOk() (*bool, bool) {
	if o == nil || IsNil(o.DisableStripSlash) {
		return nil, false
	}
	return o.DisableStripSlash, true
}

// HasDisableStripSlash returns a boolean if a field has been set.
func (o *APIDefinitionProxy) HasDisableStripSlash() bool {
	if o != nil && !IsNil(o.DisableStripSlash) {
		return true
	}

	return false
}

// SetDisableStripSlash gets a reference to the given bool and assigns it to the DisableStripSlash field.
func (o *APIDefinitionProxy) SetDisableStripSlash(v bool) {
	o.DisableStripSlash = &v
}

// GetEnableLoadBalancing returns the EnableLoadBalancing field value if set, zero value otherwise.
func (o *APIDefinitionProxy) GetEnableLoadBalancing() bool {
	if o == nil || IsNil(o.EnableLoadBalancing) {
		var ret bool
		return ret
	}
	return *o.EnableLoadBalancing
}

// GetEnableLoadBalancingOk returns a tuple with the EnableLoadBalancing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APIDefinitionProxy) GetEnableLoadBalancingOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableLoadBalancing) {
		return nil, false
	}
	return o.EnableLoadBalancing, true
}

// HasEnableLoadBalancing returns a boolean if a field has been set.
func (o *APIDefinitionProxy) HasEnableLoadBalancing() bool {
	if o != nil && !IsNil(o.EnableLoadBalancing) {
		return true
	}

	return false
}

// SetEnableLoadBalancing gets a reference to the given bool and assigns it to the EnableLoadBalancing field.
func (o *APIDefinitionProxy) SetEnableLoadBalancing(v bool) {
	o.EnableLoadBalancing = &v
}

// GetListenPath returns the ListenPath field value if set, zero value otherwise.
func (o *APIDefinitionProxy) GetListenPath() string {
	if o == nil || IsNil(o.ListenPath) {
		var ret string
		return ret
	}
	return *o.ListenPath
}

// GetListenPathOk returns a tuple with the ListenPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APIDefinitionProxy) GetListenPathOk() (*string, bool) {
	if o == nil || IsNil(o.ListenPath) {
		return nil, false
	}
	return o.ListenPath, true
}

// HasListenPath returns a boolean if a field has been set.
func (o *APIDefinitionProxy) HasListenPath() bool {
	if o != nil && !IsNil(o.ListenPath) {
		return true
	}

	return false
}

// SetListenPath gets a reference to the given string and assigns it to the ListenPath field.
func (o *APIDefinitionProxy) SetListenPath(v string) {
	o.ListenPath = &v
}

// GetPreserveHostHeader returns the PreserveHostHeader field value if set, zero value otherwise.
func (o *APIDefinitionProxy) GetPreserveHostHeader() bool {
	if o == nil || IsNil(o.PreserveHostHeader) {
		var ret bool
		return ret
	}
	return *o.PreserveHostHeader
}

// GetPreserveHostHeaderOk returns a tuple with the PreserveHostHeader field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APIDefinitionProxy) GetPreserveHostHeaderOk() (*bool, bool) {
	if o == nil || IsNil(o.PreserveHostHeader) {
		return nil, false
	}
	return o.PreserveHostHeader, true
}

// HasPreserveHostHeader returns a boolean if a field has been set.
func (o *APIDefinitionProxy) HasPreserveHostHeader() bool {
	if o != nil && !IsNil(o.PreserveHostHeader) {
		return true
	}

	return false
}

// SetPreserveHostHeader gets a reference to the given bool and assigns it to the PreserveHostHeader field.
func (o *APIDefinitionProxy) SetPreserveHostHeader(v bool) {
	o.PreserveHostHeader = &v
}

// GetServiceDiscovery returns the ServiceDiscovery field value if set, zero value otherwise.
func (o *APIDefinitionProxy) GetServiceDiscovery() ServiceDiscoveryConfiguration {
	if o == nil || IsNil(o.ServiceDiscovery) {
		var ret ServiceDiscoveryConfiguration
		return ret
	}
	return *o.ServiceDiscovery
}

// GetServiceDiscoveryOk returns a tuple with the ServiceDiscovery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APIDefinitionProxy) GetServiceDiscoveryOk() (*ServiceDiscoveryConfiguration, bool) {
	if o == nil || IsNil(o.ServiceDiscovery) {
		return nil, false
	}
	return o.ServiceDiscovery, true
}

// HasServiceDiscovery returns a boolean if a field has been set.
func (o *APIDefinitionProxy) HasServiceDiscovery() bool {
	if o != nil && !IsNil(o.ServiceDiscovery) {
		return true
	}

	return false
}

// SetServiceDiscovery gets a reference to the given ServiceDiscoveryConfiguration and assigns it to the ServiceDiscovery field.
func (o *APIDefinitionProxy) SetServiceDiscovery(v ServiceDiscoveryConfiguration) {
	o.ServiceDiscovery = &v
}

// GetStripListenPath returns the StripListenPath field value if set, zero value otherwise.
func (o *APIDefinitionProxy) GetStripListenPath() bool {
	if o == nil || IsNil(o.StripListenPath) {
		var ret bool
		return ret
	}
	return *o.StripListenPath
}

// GetStripListenPathOk returns a tuple with the StripListenPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APIDefinitionProxy) GetStripListenPathOk() (*bool, bool) {
	if o == nil || IsNil(o.StripListenPath) {
		return nil, false
	}
	return o.StripListenPath, true
}

// HasStripListenPath returns a boolean if a field has been set.
func (o *APIDefinitionProxy) HasStripListenPath() bool {
	if o != nil && !IsNil(o.StripListenPath) {
		return true
	}

	return false
}

// SetStripListenPath gets a reference to the given bool and assigns it to the StripListenPath field.
func (o *APIDefinitionProxy) SetStripListenPath(v bool) {
	o.StripListenPath = &v
}

// GetTargetList returns the TargetList field value if set, zero value otherwise.
func (o *APIDefinitionProxy) GetTargetList() []string {
	if o == nil || IsNil(o.TargetList) {
		var ret []string
		return ret
	}
	return o.TargetList
}

// GetTargetListOk returns a tuple with the TargetList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APIDefinitionProxy) GetTargetListOk() ([]string, bool) {
	if o == nil || IsNil(o.TargetList) {
		return nil, false
	}
	return o.TargetList, true
}

// HasTargetList returns a boolean if a field has been set.
func (o *APIDefinitionProxy) HasTargetList() bool {
	if o != nil && !IsNil(o.TargetList) {
		return true
	}

	return false
}

// SetTargetList gets a reference to the given []string and assigns it to the TargetList field.
func (o *APIDefinitionProxy) SetTargetList(v []string) {
	o.TargetList = v
}

// GetTargetUrl returns the TargetUrl field value if set, zero value otherwise.
func (o *APIDefinitionProxy) GetTargetUrl() string {
	if o == nil || IsNil(o.TargetUrl) {
		var ret string
		return ret
	}
	return *o.TargetUrl
}

// GetTargetUrlOk returns a tuple with the TargetUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APIDefinitionProxy) GetTargetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.TargetUrl) {
		return nil, false
	}
	return o.TargetUrl, true
}

// HasTargetUrl returns a boolean if a field has been set.
func (o *APIDefinitionProxy) HasTargetUrl() bool {
	if o != nil && !IsNil(o.TargetUrl) {
		return true
	}

	return false
}

// SetTargetUrl gets a reference to the given string and assigns it to the TargetUrl field.
func (o *APIDefinitionProxy) SetTargetUrl(v string) {
	o.TargetUrl = &v
}

// GetTransport returns the Transport field value if set, zero value otherwise.
func (o *APIDefinitionProxy) GetTransport() APIDefinitionProxyTransport {
	if o == nil || IsNil(o.Transport) {
		var ret APIDefinitionProxyTransport
		return ret
	}
	return *o.Transport
}

// GetTransportOk returns a tuple with the Transport field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APIDefinitionProxy) GetTransportOk() (*APIDefinitionProxyTransport, bool) {
	if o == nil || IsNil(o.Transport) {
		return nil, false
	}
	return o.Transport, true
}

// HasTransport returns a boolean if a field has been set.
func (o *APIDefinitionProxy) HasTransport() bool {
	if o != nil && !IsNil(o.Transport) {
		return true
	}

	return false
}

// SetTransport gets a reference to the given APIDefinitionProxyTransport and assigns it to the Transport field.
func (o *APIDefinitionProxy) SetTransport(v APIDefinitionProxyTransport) {
	o.Transport = &v
}

func (o APIDefinitionProxy) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o APIDefinitionProxy) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CheckHostAgainstUptimeTests) {
		toSerialize["check_host_against_uptime_tests"] = o.CheckHostAgainstUptimeTests
	}
	if !IsNil(o.DisableStripSlash) {
		toSerialize["disable_strip_slash"] = o.DisableStripSlash
	}
	if !IsNil(o.EnableLoadBalancing) {
		toSerialize["enable_load_balancing"] = o.EnableLoadBalancing
	}
	if !IsNil(o.ListenPath) {
		toSerialize["listen_path"] = o.ListenPath
	}
	if !IsNil(o.PreserveHostHeader) {
		toSerialize["preserve_host_header"] = o.PreserveHostHeader
	}
	if !IsNil(o.ServiceDiscovery) {
		toSerialize["service_discovery"] = o.ServiceDiscovery
	}
	if !IsNil(o.StripListenPath) {
		toSerialize["strip_listen_path"] = o.StripListenPath
	}
	if !IsNil(o.TargetList) {
		toSerialize["target_list"] = o.TargetList
	}
	if !IsNil(o.TargetUrl) {
		toSerialize["target_url"] = o.TargetUrl
	}
	if !IsNil(o.Transport) {
		toSerialize["transport"] = o.Transport
	}
	return toSerialize, nil
}

type NullableAPIDefinitionProxy struct {
	value *APIDefinitionProxy
	isSet bool
}

func (v NullableAPIDefinitionProxy) Get() *APIDefinitionProxy {
	return v.value
}

func (v *NullableAPIDefinitionProxy) Set(val *APIDefinitionProxy) {
	v.value = val
	v.isSet = true
}

func (v NullableAPIDefinitionProxy) IsSet() bool {
	return v.isSet
}

func (v *NullableAPIDefinitionProxy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAPIDefinitionProxy(val *APIDefinitionProxy) *NullableAPIDefinitionProxy {
	return &NullableAPIDefinitionProxy{value: val, isSet: true}
}

func (v NullableAPIDefinitionProxy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAPIDefinitionProxy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


