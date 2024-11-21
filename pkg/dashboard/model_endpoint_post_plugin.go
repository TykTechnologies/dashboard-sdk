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

// checks if the EndpointPostPlugin type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EndpointPostPlugin{}

// EndpointPostPlugin struct for EndpointPostPlugin
type EndpointPostPlugin struct {
	Enabled      *bool   `json:"enabled,omitempty"`
	FunctionName *string `json:"functionName,omitempty"`
	Name         *string `json:"name,omitempty"`
	Path         *string `json:"path,omitempty"`
}

// NewEndpointPostPlugin instantiates a new EndpointPostPlugin object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEndpointPostPlugin() *EndpointPostPlugin {
	this := EndpointPostPlugin{}
	return &this
}

// NewEndpointPostPluginWithDefaults instantiates a new EndpointPostPlugin object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEndpointPostPluginWithDefaults() *EndpointPostPlugin {
	this := EndpointPostPlugin{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *EndpointPostPlugin) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndpointPostPlugin) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *EndpointPostPlugin) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *EndpointPostPlugin) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetFunctionName returns the FunctionName field value if set, zero value otherwise.
func (o *EndpointPostPlugin) GetFunctionName() string {
	if o == nil || IsNil(o.FunctionName) {
		var ret string
		return ret
	}
	return *o.FunctionName
}

// GetFunctionNameOk returns a tuple with the FunctionName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndpointPostPlugin) GetFunctionNameOk() (*string, bool) {
	if o == nil || IsNil(o.FunctionName) {
		return nil, false
	}
	return o.FunctionName, true
}

// HasFunctionName returns a boolean if a field has been set.
func (o *EndpointPostPlugin) HasFunctionName() bool {
	if o != nil && !IsNil(o.FunctionName) {
		return true
	}

	return false
}

// SetFunctionName gets a reference to the given string and assigns it to the FunctionName field.
func (o *EndpointPostPlugin) SetFunctionName(v string) {
	o.FunctionName = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *EndpointPostPlugin) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndpointPostPlugin) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *EndpointPostPlugin) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *EndpointPostPlugin) SetName(v string) {
	o.Name = &v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *EndpointPostPlugin) GetPath() string {
	if o == nil || IsNil(o.Path) {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndpointPostPlugin) GetPathOk() (*string, bool) {
	if o == nil || IsNil(o.Path) {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *EndpointPostPlugin) HasPath() bool {
	if o != nil && !IsNil(o.Path) {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *EndpointPostPlugin) SetPath(v string) {
	o.Path = &v
}

func (o EndpointPostPlugin) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EndpointPostPlugin) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.FunctionName) {
		toSerialize["functionName"] = o.FunctionName
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Path) {
		toSerialize["path"] = o.Path
	}
	return toSerialize, nil
}

type NullableEndpointPostPlugin struct {
	value *EndpointPostPlugin
	isSet bool
}

func (v NullableEndpointPostPlugin) Get() *EndpointPostPlugin {
	return v.value
}

func (v *NullableEndpointPostPlugin) Set(val *EndpointPostPlugin) {
	v.value = val
	v.isSet = true
}

func (v NullableEndpointPostPlugin) IsSet() bool {
	return v.isSet
}

func (v *NullableEndpointPostPlugin) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEndpointPostPlugin(val *EndpointPostPlugin) *NullableEndpointPostPlugin {
	return &NullableEndpointPostPlugin{value: val, isSet: true}
}

func (v NullableEndpointPostPlugin) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEndpointPostPlugin) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
