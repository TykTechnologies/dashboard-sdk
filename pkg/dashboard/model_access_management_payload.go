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

// checks if the AccessManagementPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccessManagementPayload{}

// AccessManagementPayload struct for AccessManagementPayload
type AccessManagementPayload struct {
	UserGroupIds []string `json:"userGroupIds,omitempty"`
	UserIds      []string `json:"userIds,omitempty"`
}

// NewAccessManagementPayload instantiates a new AccessManagementPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccessManagementPayload() *AccessManagementPayload {
	this := AccessManagementPayload{}
	return &this
}

// NewAccessManagementPayloadWithDefaults instantiates a new AccessManagementPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccessManagementPayloadWithDefaults() *AccessManagementPayload {
	this := AccessManagementPayload{}
	return &this
}

// GetUserGroupIds returns the UserGroupIds field value if set, zero value otherwise.
func (o *AccessManagementPayload) GetUserGroupIds() []string {
	if o == nil || IsNil(o.UserGroupIds) {
		var ret []string
		return ret
	}
	return o.UserGroupIds
}

// GetUserGroupIdsOk returns a tuple with the UserGroupIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessManagementPayload) GetUserGroupIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.UserGroupIds) {
		return nil, false
	}
	return o.UserGroupIds, true
}

// HasUserGroupIds returns a boolean if a field has been set.
func (o *AccessManagementPayload) HasUserGroupIds() bool {
	if o != nil && !IsNil(o.UserGroupIds) {
		return true
	}

	return false
}

// SetUserGroupIds gets a reference to the given []string and assigns it to the UserGroupIds field.
func (o *AccessManagementPayload) SetUserGroupIds(v []string) {
	o.UserGroupIds = v
}

// GetUserIds returns the UserIds field value if set, zero value otherwise.
func (o *AccessManagementPayload) GetUserIds() []string {
	if o == nil || IsNil(o.UserIds) {
		var ret []string
		return ret
	}
	return o.UserIds
}

// GetUserIdsOk returns a tuple with the UserIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessManagementPayload) GetUserIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.UserIds) {
		return nil, false
	}
	return o.UserIds, true
}

// HasUserIds returns a boolean if a field has been set.
func (o *AccessManagementPayload) HasUserIds() bool {
	if o != nil && !IsNil(o.UserIds) {
		return true
	}

	return false
}

// SetUserIds gets a reference to the given []string and assigns it to the UserIds field.
func (o *AccessManagementPayload) SetUserIds(v []string) {
	o.UserIds = v
}

func (o AccessManagementPayload) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccessManagementPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.UserGroupIds) {
		toSerialize["userGroupIds"] = o.UserGroupIds
	}
	if !IsNil(o.UserIds) {
		toSerialize["userIds"] = o.UserIds
	}
	return toSerialize, nil
}

type NullableAccessManagementPayload struct {
	value *AccessManagementPayload
	isSet bool
}

func (v NullableAccessManagementPayload) Get() *AccessManagementPayload {
	return v.value
}

func (v *NullableAccessManagementPayload) Set(val *AccessManagementPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableAccessManagementPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableAccessManagementPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccessManagementPayload(val *AccessManagementPayload) *NullableAccessManagementPayload {
	return &NullableAccessManagementPayload{value: val, isSet: true}
}

func (v NullableAccessManagementPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccessManagementPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
