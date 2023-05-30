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

// checks if the SessionProviderMeta type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SessionProviderMeta{}

// SessionProviderMeta struct for SessionProviderMeta
type SessionProviderMeta struct {
	Meta          map[string]map[string]interface{} `json:"meta,omitempty"`
	Name          *string                           `json:"name,omitempty"`
	StorageEngine *string                           `json:"storage_engine,omitempty"`
}

// NewSessionProviderMeta instantiates a new SessionProviderMeta object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSessionProviderMeta() *SessionProviderMeta {
	this := SessionProviderMeta{}
	return &this
}

// NewSessionProviderMetaWithDefaults instantiates a new SessionProviderMeta object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSessionProviderMetaWithDefaults() *SessionProviderMeta {
	this := SessionProviderMeta{}
	return &this
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *SessionProviderMeta) GetMeta() map[string]map[string]interface{} {
	if o == nil || IsNil(o.Meta) {
		var ret map[string]map[string]interface{}
		return ret
	}
	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionProviderMeta) GetMetaOk() (map[string]map[string]interface{}, bool) {
	if o == nil || IsNil(o.Meta) {
		return map[string]map[string]interface{}{}, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *SessionProviderMeta) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given map[string]map[string]interface{} and assigns it to the Meta field.
func (o *SessionProviderMeta) SetMeta(v map[string]map[string]interface{}) {
	o.Meta = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SessionProviderMeta) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionProviderMeta) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SessionProviderMeta) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SessionProviderMeta) SetName(v string) {
	o.Name = &v
}

// GetStorageEngine returns the StorageEngine field value if set, zero value otherwise.
func (o *SessionProviderMeta) GetStorageEngine() string {
	if o == nil || IsNil(o.StorageEngine) {
		var ret string
		return ret
	}
	return *o.StorageEngine
}

// GetStorageEngineOk returns a tuple with the StorageEngine field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionProviderMeta) GetStorageEngineOk() (*string, bool) {
	if o == nil || IsNil(o.StorageEngine) {
		return nil, false
	}
	return o.StorageEngine, true
}

// HasStorageEngine returns a boolean if a field has been set.
func (o *SessionProviderMeta) HasStorageEngine() bool {
	if o != nil && !IsNil(o.StorageEngine) {
		return true
	}

	return false
}

// SetStorageEngine gets a reference to the given string and assigns it to the StorageEngine field.
func (o *SessionProviderMeta) SetStorageEngine(v string) {
	o.StorageEngine = &v
}

func (o SessionProviderMeta) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SessionProviderMeta) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Meta) {
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

type NullableSessionProviderMeta struct {
	value *SessionProviderMeta
	isSet bool
}

func (v NullableSessionProviderMeta) Get() *SessionProviderMeta {
	return v.value
}

func (v *NullableSessionProviderMeta) Set(val *SessionProviderMeta) {
	v.value = val
	v.isSet = true
}

func (v NullableSessionProviderMeta) IsSet() bool {
	return v.isSet
}

func (v *NullableSessionProviderMeta) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSessionProviderMeta(val *SessionProviderMeta) *NullableSessionProviderMeta {
	return &NullableSessionProviderMeta{value: val, isSet: true}
}

func (v NullableSessionProviderMeta) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSessionProviderMeta) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
