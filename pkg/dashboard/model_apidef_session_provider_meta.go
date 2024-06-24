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

// checks if the ApidefSessionProviderMeta type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApidefSessionProviderMeta{}

// ApidefSessionProviderMeta struct for ApidefSessionProviderMeta
type ApidefSessionProviderMeta struct {
	Meta          map[string]interface{} `json:"meta,omitempty"`
	Name          *string                `json:"name,omitempty"`
	StorageEngine *string                `json:"storage_engine,omitempty"`
}

// NewApidefSessionProviderMeta instantiates a new ApidefSessionProviderMeta object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApidefSessionProviderMeta() *ApidefSessionProviderMeta {
	this := ApidefSessionProviderMeta{}
	return &this
}

// NewApidefSessionProviderMetaWithDefaults instantiates a new ApidefSessionProviderMeta object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApidefSessionProviderMetaWithDefaults() *ApidefSessionProviderMeta {
	this := ApidefSessionProviderMeta{}
	return &this
}

// GetMeta returns the Meta field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApidefSessionProviderMeta) GetMeta() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApidefSessionProviderMeta) GetMetaOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Meta) {
		return map[string]interface{}{}, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *ApidefSessionProviderMeta) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given map[string]interface{} and assigns it to the Meta field.
func (o *ApidefSessionProviderMeta) SetMeta(v map[string]interface{}) {
	o.Meta = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ApidefSessionProviderMeta) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApidefSessionProviderMeta) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ApidefSessionProviderMeta) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ApidefSessionProviderMeta) SetName(v string) {
	o.Name = &v
}

// GetStorageEngine returns the StorageEngine field value if set, zero value otherwise.
func (o *ApidefSessionProviderMeta) GetStorageEngine() string {
	if o == nil || IsNil(o.StorageEngine) {
		var ret string
		return ret
	}
	return *o.StorageEngine
}

// GetStorageEngineOk returns a tuple with the StorageEngine field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApidefSessionProviderMeta) GetStorageEngineOk() (*string, bool) {
	if o == nil || IsNil(o.StorageEngine) {
		return nil, false
	}
	return o.StorageEngine, true
}

// HasStorageEngine returns a boolean if a field has been set.
func (o *ApidefSessionProviderMeta) HasStorageEngine() bool {
	if o != nil && !IsNil(o.StorageEngine) {
		return true
	}

	return false
}

// SetStorageEngine gets a reference to the given string and assigns it to the StorageEngine field.
func (o *ApidefSessionProviderMeta) SetStorageEngine(v string) {
	o.StorageEngine = &v
}

func (o ApidefSessionProviderMeta) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApidefSessionProviderMeta) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Meta != nil {
		toSerialize["meta"] = o.Meta
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.StorageEngine) {
		toSerialize["storage_engine"] = o.StorageEngine
	}
	return toSerialize, nil
}

type NullableApidefSessionProviderMeta struct {
	value *ApidefSessionProviderMeta
	isSet bool
}

func (v NullableApidefSessionProviderMeta) Get() *ApidefSessionProviderMeta {
	return v.value
}

func (v *NullableApidefSessionProviderMeta) Set(val *ApidefSessionProviderMeta) {
	v.value = val
	v.isSet = true
}

func (v NullableApidefSessionProviderMeta) IsSet() bool {
	return v.isSet
}

func (v *NullableApidefSessionProviderMeta) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApidefSessionProviderMeta(val *ApidefSessionProviderMeta) *NullableApidefSessionProviderMeta {
	return &NullableApidefSessionProviderMeta{value: val, isSet: true}
}

func (v NullableApidefSessionProviderMeta) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApidefSessionProviderMeta) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
