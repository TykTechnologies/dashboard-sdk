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

// checks if the NewAdditionalPermissions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NewAdditionalPermissions{}

// NewAdditionalPermissions struct for NewAdditionalPermissions
type NewAdditionalPermissions struct {
	AdditionalPermissions map[string]string `json:"additional_permissions,omitempty"`
}

// NewNewAdditionalPermissions instantiates a new NewAdditionalPermissions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNewAdditionalPermissions() *NewAdditionalPermissions {
	this := NewAdditionalPermissions{}
	return &this
}

// NewNewAdditionalPermissionsWithDefaults instantiates a new NewAdditionalPermissions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNewAdditionalPermissionsWithDefaults() *NewAdditionalPermissions {
	this := NewAdditionalPermissions{}
	return &this
}

// GetAdditionalPermissions returns the AdditionalPermissions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NewAdditionalPermissions) GetAdditionalPermissions() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}
	return o.AdditionalPermissions
}

// GetAdditionalPermissionsOk returns a tuple with the AdditionalPermissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NewAdditionalPermissions) GetAdditionalPermissionsOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.AdditionalPermissions) {
		return nil, false
	}
	return &o.AdditionalPermissions, true
}

// HasAdditionalPermissions returns a boolean if a field has been set.
func (o *NewAdditionalPermissions) HasAdditionalPermissions() bool {
	if o != nil && !IsNil(o.AdditionalPermissions) {
		return true
	}

	return false
}

// SetAdditionalPermissions gets a reference to the given map[string]string and assigns it to the AdditionalPermissions field.
func (o *NewAdditionalPermissions) SetAdditionalPermissions(v map[string]string) {
	o.AdditionalPermissions = v
}

func (o NewAdditionalPermissions) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NewAdditionalPermissions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.AdditionalPermissions != nil {
		toSerialize["additional_permissions"] = o.AdditionalPermissions
	}
	return toSerialize, nil
}

type NullableNewAdditionalPermissions struct {
	value *NewAdditionalPermissions
	isSet bool
}

func (v NullableNewAdditionalPermissions) Get() *NewAdditionalPermissions {
	return v.value
}

func (v *NullableNewAdditionalPermissions) Set(val *NewAdditionalPermissions) {
	v.value = val
	v.isSet = true
}

func (v NullableNewAdditionalPermissions) IsSet() bool {
	return v.isSet
}

func (v *NullableNewAdditionalPermissions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNewAdditionalPermissions(val *NewAdditionalPermissions) *NullableNewAdditionalPermissions {
	return &NullableNewAdditionalPermissions{value: val, isSet: true}
}

func (v NullableNewAdditionalPermissions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNewAdditionalPermissions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
