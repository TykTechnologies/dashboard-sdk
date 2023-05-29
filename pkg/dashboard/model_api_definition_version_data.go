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

// checks if the APIDefinitionVersionData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &APIDefinitionVersionData{}

// APIDefinitionVersionData struct for APIDefinitionVersionData
type APIDefinitionVersionData struct {
	DefaultVersion *string `json:"default_version,omitempty"`
	NotVersioned *bool `json:"not_versioned,omitempty"`
	Versions *map[string]VersionInfo `json:"versions,omitempty"`
}

// NewAPIDefinitionVersionData instantiates a new APIDefinitionVersionData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAPIDefinitionVersionData() *APIDefinitionVersionData {
	this := APIDefinitionVersionData{}
	return &this
}

// NewAPIDefinitionVersionDataWithDefaults instantiates a new APIDefinitionVersionData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAPIDefinitionVersionDataWithDefaults() *APIDefinitionVersionData {
	this := APIDefinitionVersionData{}
	return &this
}

// GetDefaultVersion returns the DefaultVersion field value if set, zero value otherwise.
func (o *APIDefinitionVersionData) GetDefaultVersion() string {
	if o == nil || IsNil(o.DefaultVersion) {
		var ret string
		return ret
	}
	return *o.DefaultVersion
}

// GetDefaultVersionOk returns a tuple with the DefaultVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APIDefinitionVersionData) GetDefaultVersionOk() (*string, bool) {
	if o == nil || IsNil(o.DefaultVersion) {
		return nil, false
	}
	return o.DefaultVersion, true
}

// HasDefaultVersion returns a boolean if a field has been set.
func (o *APIDefinitionVersionData) HasDefaultVersion() bool {
	if o != nil && !IsNil(o.DefaultVersion) {
		return true
	}

	return false
}

// SetDefaultVersion gets a reference to the given string and assigns it to the DefaultVersion field.
func (o *APIDefinitionVersionData) SetDefaultVersion(v string) {
	o.DefaultVersion = &v
}

// GetNotVersioned returns the NotVersioned field value if set, zero value otherwise.
func (o *APIDefinitionVersionData) GetNotVersioned() bool {
	if o == nil || IsNil(o.NotVersioned) {
		var ret bool
		return ret
	}
	return *o.NotVersioned
}

// GetNotVersionedOk returns a tuple with the NotVersioned field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APIDefinitionVersionData) GetNotVersionedOk() (*bool, bool) {
	if o == nil || IsNil(o.NotVersioned) {
		return nil, false
	}
	return o.NotVersioned, true
}

// HasNotVersioned returns a boolean if a field has been set.
func (o *APIDefinitionVersionData) HasNotVersioned() bool {
	if o != nil && !IsNil(o.NotVersioned) {
		return true
	}

	return false
}

// SetNotVersioned gets a reference to the given bool and assigns it to the NotVersioned field.
func (o *APIDefinitionVersionData) SetNotVersioned(v bool) {
	o.NotVersioned = &v
}

// GetVersions returns the Versions field value if set, zero value otherwise.
func (o *APIDefinitionVersionData) GetVersions() map[string]VersionInfo {
	if o == nil || IsNil(o.Versions) {
		var ret map[string]VersionInfo
		return ret
	}
	return *o.Versions
}

// GetVersionsOk returns a tuple with the Versions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APIDefinitionVersionData) GetVersionsOk() (*map[string]VersionInfo, bool) {
	if o == nil || IsNil(o.Versions) {
		return nil, false
	}
	return o.Versions, true
}

// HasVersions returns a boolean if a field has been set.
func (o *APIDefinitionVersionData) HasVersions() bool {
	if o != nil && !IsNil(o.Versions) {
		return true
	}

	return false
}

// SetVersions gets a reference to the given map[string]VersionInfo and assigns it to the Versions field.
func (o *APIDefinitionVersionData) SetVersions(v map[string]VersionInfo) {
	o.Versions = &v
}

func (o APIDefinitionVersionData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o APIDefinitionVersionData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DefaultVersion) {
		toSerialize["default_version"] = o.DefaultVersion
	}
	if !IsNil(o.NotVersioned) {
		toSerialize["not_versioned"] = o.NotVersioned
	}
	if !IsNil(o.Versions) {
		toSerialize["versions"] = o.Versions
	}
	return toSerialize, nil
}

type NullableAPIDefinitionVersionData struct {
	value *APIDefinitionVersionData
	isSet bool
}

func (v NullableAPIDefinitionVersionData) Get() *APIDefinitionVersionData {
	return v.value
}

func (v *NullableAPIDefinitionVersionData) Set(val *APIDefinitionVersionData) {
	v.value = val
	v.isSet = true
}

func (v NullableAPIDefinitionVersionData) IsSet() bool {
	return v.isSet
}

func (v *NullableAPIDefinitionVersionData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAPIDefinitionVersionData(val *APIDefinitionVersionData) *NullableAPIDefinitionVersionData {
	return &NullableAPIDefinitionVersionData{value: val, isSet: true}
}

func (v NullableAPIDefinitionVersionData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAPIDefinitionVersionData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


