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

// checks if the KeyData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KeyData{}

// KeyData struct for KeyData
type KeyData struct {
	ApiModel map[string]interface{} `json:"api_model,omitempty"`
	Data     *SessionState          `json:"data,omitempty"`
	KeyHash  *string                `json:"key_hash,omitempty"`
	KeyId    *string                `json:"key_id,omitempty"`
}

// NewKeyData instantiates a new KeyData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKeyData() *KeyData {
	this := KeyData{}
	return &this
}

// NewKeyDataWithDefaults instantiates a new KeyData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKeyDataWithDefaults() *KeyData {
	this := KeyData{}
	return &this
}

// GetApiModel returns the ApiModel field value if set, zero value otherwise.
func (o *KeyData) GetApiModel() map[string]interface{} {
	if o == nil || IsNil(o.ApiModel) {
		var ret map[string]interface{}
		return ret
	}
	return o.ApiModel
}

// GetApiModelOk returns a tuple with the ApiModel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyData) GetApiModelOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.ApiModel) {
		return map[string]interface{}{}, false
	}
	return o.ApiModel, true
}

// HasApiModel returns a boolean if a field has been set.
func (o *KeyData) HasApiModel() bool {
	if o != nil && !IsNil(o.ApiModel) {
		return true
	}

	return false
}

// SetApiModel gets a reference to the given map[string]interface{} and assigns it to the ApiModel field.
func (o *KeyData) SetApiModel(v map[string]interface{}) {
	o.ApiModel = v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *KeyData) GetData() SessionState {
	if o == nil || IsNil(o.Data) {
		var ret SessionState
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyData) GetDataOk() (*SessionState, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *KeyData) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given SessionState and assigns it to the Data field.
func (o *KeyData) SetData(v SessionState) {
	o.Data = &v
}

// GetKeyHash returns the KeyHash field value if set, zero value otherwise.
func (o *KeyData) GetKeyHash() string {
	if o == nil || IsNil(o.KeyHash) {
		var ret string
		return ret
	}
	return *o.KeyHash
}

// GetKeyHashOk returns a tuple with the KeyHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyData) GetKeyHashOk() (*string, bool) {
	if o == nil || IsNil(o.KeyHash) {
		return nil, false
	}
	return o.KeyHash, true
}

// HasKeyHash returns a boolean if a field has been set.
func (o *KeyData) HasKeyHash() bool {
	if o != nil && !IsNil(o.KeyHash) {
		return true
	}

	return false
}

// SetKeyHash gets a reference to the given string and assigns it to the KeyHash field.
func (o *KeyData) SetKeyHash(v string) {
	o.KeyHash = &v
}

// GetKeyId returns the KeyId field value if set, zero value otherwise.
func (o *KeyData) GetKeyId() string {
	if o == nil || IsNil(o.KeyId) {
		var ret string
		return ret
	}
	return *o.KeyId
}

// GetKeyIdOk returns a tuple with the KeyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyData) GetKeyIdOk() (*string, bool) {
	if o == nil || IsNil(o.KeyId) {
		return nil, false
	}
	return o.KeyId, true
}

// HasKeyId returns a boolean if a field has been set.
func (o *KeyData) HasKeyId() bool {
	if o != nil && !IsNil(o.KeyId) {
		return true
	}

	return false
}

// SetKeyId gets a reference to the given string and assigns it to the KeyId field.
func (o *KeyData) SetKeyId(v string) {
	o.KeyId = &v
}

func (o KeyData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KeyData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ApiModel) {
		toSerialize["api_model"] = o.ApiModel
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.KeyHash) {
		toSerialize["key_hash"] = o.KeyHash
	}
	if !IsNil(o.KeyId) {
		toSerialize["key_id"] = o.KeyId
	}
	return toSerialize, nil
}

type NullableKeyData struct {
	value *KeyData
	isSet bool
}

func (v NullableKeyData) Get() *KeyData {
	return v.value
}

func (v *NullableKeyData) Set(val *KeyData) {
	v.value = val
	v.isSet = true
}

func (v NullableKeyData) IsSet() bool {
	return v.isSet
}

func (v *NullableKeyData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKeyData(val *KeyData) *NullableKeyData {
	return &NullableKeyData{value: val, isSet: true}
}

func (v NullableKeyData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKeyData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
