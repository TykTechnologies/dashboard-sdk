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

// checks if the UserPassword type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserPassword{}

// UserPassword struct for UserPassword
type UserPassword struct {
	ApiModel        map[string]interface{} `json:"api_model,omitempty"`
	CurrentPassword *string                `json:"current_password,omitempty"`
	NewPassword     *string                `json:"new_password,omitempty"`
}

// NewUserPassword instantiates a new UserPassword object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserPassword() *UserPassword {
	this := UserPassword{}
	return &this
}

// NewUserPasswordWithDefaults instantiates a new UserPassword object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserPasswordWithDefaults() *UserPassword {
	this := UserPassword{}
	return &this
}

// GetApiModel returns the ApiModel field value if set, zero value otherwise.
func (o *UserPassword) GetApiModel() map[string]interface{} {
	if o == nil || IsNil(o.ApiModel) {
		var ret map[string]interface{}
		return ret
	}
	return o.ApiModel
}

// GetApiModelOk returns a tuple with the ApiModel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserPassword) GetApiModelOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.ApiModel) {
		return map[string]interface{}{}, false
	}
	return o.ApiModel, true
}

// HasApiModel returns a boolean if a field has been set.
func (o *UserPassword) HasApiModel() bool {
	if o != nil && !IsNil(o.ApiModel) {
		return true
	}

	return false
}

// SetApiModel gets a reference to the given map[string]interface{} and assigns it to the ApiModel field.
func (o *UserPassword) SetApiModel(v map[string]interface{}) {
	o.ApiModel = v
}

// GetCurrentPassword returns the CurrentPassword field value if set, zero value otherwise.
func (o *UserPassword) GetCurrentPassword() string {
	if o == nil || IsNil(o.CurrentPassword) {
		var ret string
		return ret
	}
	return *o.CurrentPassword
}

// GetCurrentPasswordOk returns a tuple with the CurrentPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserPassword) GetCurrentPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.CurrentPassword) {
		return nil, false
	}
	return o.CurrentPassword, true
}

// HasCurrentPassword returns a boolean if a field has been set.
func (o *UserPassword) HasCurrentPassword() bool {
	if o != nil && !IsNil(o.CurrentPassword) {
		return true
	}

	return false
}

// SetCurrentPassword gets a reference to the given string and assigns it to the CurrentPassword field.
func (o *UserPassword) SetCurrentPassword(v string) {
	o.CurrentPassword = &v
}

// GetNewPassword returns the NewPassword field value if set, zero value otherwise.
func (o *UserPassword) GetNewPassword() string {
	if o == nil || IsNil(o.NewPassword) {
		var ret string
		return ret
	}
	return *o.NewPassword
}

// GetNewPasswordOk returns a tuple with the NewPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserPassword) GetNewPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.NewPassword) {
		return nil, false
	}
	return o.NewPassword, true
}

// HasNewPassword returns a boolean if a field has been set.
func (o *UserPassword) HasNewPassword() bool {
	if o != nil && !IsNil(o.NewPassword) {
		return true
	}

	return false
}

// SetNewPassword gets a reference to the given string and assigns it to the NewPassword field.
func (o *UserPassword) SetNewPassword(v string) {
	o.NewPassword = &v
}

func (o UserPassword) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserPassword) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ApiModel) {
		toSerialize["api_model"] = o.ApiModel
	}
	if !IsNil(o.CurrentPassword) {
		toSerialize["current_password"] = o.CurrentPassword
	}
	if !IsNil(o.NewPassword) {
		toSerialize["new_password"] = o.NewPassword
	}
	return toSerialize, nil
}

type NullableUserPassword struct {
	value *UserPassword
	isSet bool
}

func (v NullableUserPassword) Get() *UserPassword {
	return v.value
}

func (v *NullableUserPassword) Set(val *UserPassword) {
	v.value = val
	v.isSet = true
}

func (v NullableUserPassword) IsSet() bool {
	return v.isSet
}

func (v *NullableUserPassword) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserPassword(val *UserPassword) *NullableUserPassword {
	return &NullableUserPassword{value: val, isSet: true}
}

func (v NullableUserPassword) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserPassword) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
