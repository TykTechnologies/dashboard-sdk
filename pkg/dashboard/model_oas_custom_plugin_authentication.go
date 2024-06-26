/*
NEW Tyk DASH API

## <a name=\"introduction\"></a> Introduction      The Tyk Dashboard API offers granular, programmatic access to a centralised database of resources that your Tyk nodes can pull from. This API has a dynamic user administrative structure which means the secret key that is used to communicate with your Tyk nodes can be kept secret and access to the wider management functions can be handled on a user-by-user and organisation-by-organisation basis.      A common question around using a database-backed configuration is how to programatically add API definitions to your Tyk nodes, the Dashboard API allows much more fine-grained, secure and multi-user access to your Tyk cluster, and should be used to manage a database-backed Tyk node.      The Tyk Dashboard API works seamlessly with the Tyk Dashboard (and the two come bundled together).      ## <a name=\"security-hierarchy\"></a> Security Hierarchy      The Dashboard API provides a more structured security layer to managing Tyk nodes.      ### Organisations, APIs and Users      With the Dashboard API and a database-backed Tyk setup, (and to an extent with file-based API setups - if diligence is used in naming an creating definitions), the following security model is applied to the management of Upstream APIs:      * **Organisations**: All APIs are *owned* by an organisation, this is designated by the OrgID parameter in the API Definition.     * **Users**: All users created in the Dashboard belong to an organisation (unless an exception is made for super-administrative access).     * **APIs**: All APIs belong to an Organisation and only Users that belong to that organisation can see the analytics for those APIs and manage their configurations.     * **API Keys**: API Keys are designated by organisation, this means an API key that has full access rights will not be allowed to access the APIs of another organisation on the same system, but can have full access to all APIs within the organisation.     * **Access Rights**: Access rights are stored with the key, this enables a key to give access to multiple APIs, this is defined by the session object in the core Tyk API.      In order to use the Dashboard API, you'll need to get the Tyk Dashboard API Access Credentials secret from your user profile on the Dashboard UI.      The secret you set should then be sent along as a header with each Dashboard API Request in order for it to be successful:  authorization: <your-secret>

API version: 5.4.0
Contact: support@tyk.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dashboard

import (
	"encoding/json"
)

// checks if the OasCustomPluginAuthentication type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OasCustomPluginAuthentication{}

// OasCustomPluginAuthentication struct for OasCustomPluginAuthentication
type OasCustomPluginAuthentication struct {
	AuthSources *OasAuthSources          `json:"AuthSources,omitempty"`
	Config      *OasAuthenticationPlugin `json:"config,omitempty"`
	Enabled     *bool                    `json:"enabled,omitempty"`
}

// NewOasCustomPluginAuthentication instantiates a new OasCustomPluginAuthentication object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOasCustomPluginAuthentication() *OasCustomPluginAuthentication {
	this := OasCustomPluginAuthentication{}
	return &this
}

// NewOasCustomPluginAuthenticationWithDefaults instantiates a new OasCustomPluginAuthentication object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOasCustomPluginAuthenticationWithDefaults() *OasCustomPluginAuthentication {
	this := OasCustomPluginAuthentication{}
	return &this
}

// GetAuthSources returns the AuthSources field value if set, zero value otherwise.
func (o *OasCustomPluginAuthentication) GetAuthSources() OasAuthSources {
	if o == nil || IsNil(o.AuthSources) {
		var ret OasAuthSources
		return ret
	}
	return *o.AuthSources
}

// GetAuthSourcesOk returns a tuple with the AuthSources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OasCustomPluginAuthentication) GetAuthSourcesOk() (*OasAuthSources, bool) {
	if o == nil || IsNil(o.AuthSources) {
		return nil, false
	}
	return o.AuthSources, true
}

// HasAuthSources returns a boolean if a field has been set.
func (o *OasCustomPluginAuthentication) HasAuthSources() bool {
	if o != nil && !IsNil(o.AuthSources) {
		return true
	}

	return false
}

// SetAuthSources gets a reference to the given OasAuthSources and assigns it to the AuthSources field.
func (o *OasCustomPluginAuthentication) SetAuthSources(v OasAuthSources) {
	o.AuthSources = &v
}

// GetConfig returns the Config field value if set, zero value otherwise.
func (o *OasCustomPluginAuthentication) GetConfig() OasAuthenticationPlugin {
	if o == nil || IsNil(o.Config) {
		var ret OasAuthenticationPlugin
		return ret
	}
	return *o.Config
}

// GetConfigOk returns a tuple with the Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OasCustomPluginAuthentication) GetConfigOk() (*OasAuthenticationPlugin, bool) {
	if o == nil || IsNil(o.Config) {
		return nil, false
	}
	return o.Config, true
}

// HasConfig returns a boolean if a field has been set.
func (o *OasCustomPluginAuthentication) HasConfig() bool {
	if o != nil && !IsNil(o.Config) {
		return true
	}

	return false
}

// SetConfig gets a reference to the given OasAuthenticationPlugin and assigns it to the Config field.
func (o *OasCustomPluginAuthentication) SetConfig(v OasAuthenticationPlugin) {
	o.Config = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *OasCustomPluginAuthentication) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OasCustomPluginAuthentication) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *OasCustomPluginAuthentication) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *OasCustomPluginAuthentication) SetEnabled(v bool) {
	o.Enabled = &v
}

func (o OasCustomPluginAuthentication) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OasCustomPluginAuthentication) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AuthSources) {
		toSerialize["AuthSources"] = o.AuthSources
	}
	if !IsNil(o.Config) {
		toSerialize["config"] = o.Config
	}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	return toSerialize, nil
}

type NullableOasCustomPluginAuthentication struct {
	value *OasCustomPluginAuthentication
	isSet bool
}

func (v NullableOasCustomPluginAuthentication) Get() *OasCustomPluginAuthentication {
	return v.value
}

func (v *NullableOasCustomPluginAuthentication) Set(val *OasCustomPluginAuthentication) {
	v.value = val
	v.isSet = true
}

func (v NullableOasCustomPluginAuthentication) IsSet() bool {
	return v.isSet
}

func (v *NullableOasCustomPluginAuthentication) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOasCustomPluginAuthentication(val *OasCustomPluginAuthentication) *NullableOasCustomPluginAuthentication {
	return &NullableOasCustomPluginAuthentication{value: val, isSet: true}
}

func (v NullableOasCustomPluginAuthentication) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOasCustomPluginAuthentication) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
