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

// checks if the SessionStateBasicAuthData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SessionStateBasicAuthData{}

// SessionStateBasicAuthData struct for SessionStateBasicAuthData
type SessionStateBasicAuthData struct {
	HashType *string `json:"hash_type,omitempty"`
	Password *string `json:"password,omitempty"`
	User     *string `json:"user,omitempty"`
}

// NewSessionStateBasicAuthData instantiates a new SessionStateBasicAuthData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSessionStateBasicAuthData() *SessionStateBasicAuthData {
	this := SessionStateBasicAuthData{}
	return &this
}

// NewSessionStateBasicAuthDataWithDefaults instantiates a new SessionStateBasicAuthData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSessionStateBasicAuthDataWithDefaults() *SessionStateBasicAuthData {
	this := SessionStateBasicAuthData{}
	return &this
}

// GetHashType returns the HashType field value if set, zero value otherwise.
func (o *SessionStateBasicAuthData) GetHashType() string {
	if o == nil || IsNil(o.HashType) {
		var ret string
		return ret
	}
	return *o.HashType
}

// GetHashTypeOk returns a tuple with the HashType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionStateBasicAuthData) GetHashTypeOk() (*string, bool) {
	if o == nil || IsNil(o.HashType) {
		return nil, false
	}
	return o.HashType, true
}

// HasHashType returns a boolean if a field has been set.
func (o *SessionStateBasicAuthData) HasHashType() bool {
	if o != nil && !IsNil(o.HashType) {
		return true
	}

	return false
}

// SetHashType gets a reference to the given string and assigns it to the HashType field.
func (o *SessionStateBasicAuthData) SetHashType(v string) {
	o.HashType = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *SessionStateBasicAuthData) GetPassword() string {
	if o == nil || IsNil(o.Password) {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionStateBasicAuthData) GetPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.Password) {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *SessionStateBasicAuthData) HasPassword() bool {
	if o != nil && !IsNil(o.Password) {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *SessionStateBasicAuthData) SetPassword(v string) {
	o.Password = &v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *SessionStateBasicAuthData) GetUser() string {
	if o == nil || IsNil(o.User) {
		var ret string
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionStateBasicAuthData) GetUserOk() (*string, bool) {
	if o == nil || IsNil(o.User) {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *SessionStateBasicAuthData) HasUser() bool {
	if o != nil && !IsNil(o.User) {
		return true
	}

	return false
}

// SetUser gets a reference to the given string and assigns it to the User field.
func (o *SessionStateBasicAuthData) SetUser(v string) {
	o.User = &v
}

func (o SessionStateBasicAuthData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SessionStateBasicAuthData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.HashType) {
		toSerialize["hash_type"] = o.HashType
	}
	if !IsNil(o.Password) {
		toSerialize["password"] = o.Password
	}
	if !IsNil(o.User) {
		toSerialize["user"] = o.User
	}
	return toSerialize, nil
}

type NullableSessionStateBasicAuthData struct {
	value *SessionStateBasicAuthData
	isSet bool
}

func (v NullableSessionStateBasicAuthData) Get() *SessionStateBasicAuthData {
	return v.value
}

func (v *NullableSessionStateBasicAuthData) Set(val *SessionStateBasicAuthData) {
	v.value = val
	v.isSet = true
}

func (v NullableSessionStateBasicAuthData) IsSet() bool {
	return v.isSet
}

func (v *NullableSessionStateBasicAuthData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSessionStateBasicAuthData(val *SessionStateBasicAuthData) *NullableSessionStateBasicAuthData {
	return &NullableSessionStateBasicAuthData{value: val, isSet: true}
}

func (v NullableSessionStateBasicAuthData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSessionStateBasicAuthData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
