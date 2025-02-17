/*
Tyk Dashboard API

 ## <a name=\"introduction\"></a> Introduction  The Tyk Dashboard API offers granular, programmatic access to a centralised database of resources that your Tyk nodes can pull from. This API has a dynamic user administrative structure which means the secret key that is used to communicate with your Tyk nodes can be kept secret and access to the wider management functions can be handled on a user-by-user and organisation-by-organisation basis.  A common question around using a database-backed configuration is how to programmatically add API definitions to your Tyk nodes, the Dashboard API allows much more fine-grained, secure and multi-user access to your Tyk cluster, and should be used to manage a database-backed Tyk node.  The Tyk Dashboard API works seamlessly with the Tyk Dashboard (and the two come bundled together).  ## <a name=\"security-hierarchy\"></a> Security Hierarchy  The Dashboard API provides a more structured security layer to managing Tyk nodes.  ### Organisations, APIs and Users  With the Dashboard API and a database-backed Tyk setup, (and to an extent with file-based API setups - if diligence is used in naming and creating definitions), the following security model is applied to the management of Upstream APIs:  * **Organisations**: All APIs are *owned* by an organisation, this is designated by the 'OrgID' parameter in the API Definition. * **Users**: All users created in the Dashboard belong to an organisation (unless an exception is made for super-administrative access). * **APIs**: All APIs belong to an Organisation and only Users that belong to that organisation can see the analytics for those APIs and manage their configurations. * **API Keys**: API Keys are designated by organisation, this means an API key that has full access rights will not be allowed to access the APIs of another organisation on the same system, but can have full access to all APIs within the organisation. * **Access Rights**: Access rights are stored with the key, this enables a key to give access to multiple APIs, this is defined by the session object in the core Tyk API.  In order to use the Dashboard API, you'll need to get the 'Tyk Dashboard API Access Credentials' secret from your user profile on the Dashboard UI.  The secret you set should then be sent along as a header with each Dashboard API Request in order for it to be successful:   authorization: <your-secret>

API version: 5.7.1
Contact: support@tyk.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dashboard

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"
)

// checks if the User type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &User{}

// User struct for User
type User struct {
	AccessKey       *string                `json:"access_key,omitempty"`
	Active          *bool                  `json:"active,omitempty"`
	ApiModel        map[string]interface{} `json:"api_model,omitempty"`
	CreatedAt       *time.Time             `json:"created_at,omitempty"`
	EmailAddress    string                 `json:"email_address"`
	FirstName       string                 `json:"first_name"`
	GroupId         *string                `json:"group_id,omitempty"`
	Id              *string                `json:"id,omitempty"`
	LastLoginDate   *time.Time             `json:"last_login_date,omitempty"`
	LastName        string                 `json:"last_name"`
	OrgId           *string                `json:"org_id,omitempty"`
	PasswordMaxDays *int32                 `json:"password_max_days,omitempty"`
	PasswordUpdated *time.Time             `json:"password_updated,omitempty"`
	UserPermissions map[string]string      `json:"user_permissions"`
}

type _User User

// NewUser instantiates a new User object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUser(emailAddress string, firstName string, lastName string, userPermissions map[string]string) *User {
	this := User{}
	this.EmailAddress = emailAddress
	this.FirstName = firstName
	this.LastName = lastName
	this.UserPermissions = userPermissions
	return &this
}

// NewUserWithDefaults instantiates a new User object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserWithDefaults() *User {
	this := User{}
	return &this
}

// GetAccessKey returns the AccessKey field value if set, zero value otherwise.
func (o *User) GetAccessKey() string {
	if o == nil || IsNil(o.AccessKey) {
		var ret string
		return ret
	}
	return *o.AccessKey
}

// GetAccessKeyOk returns a tuple with the AccessKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetAccessKeyOk() (*string, bool) {
	if o == nil || IsNil(o.AccessKey) {
		return nil, false
	}
	return o.AccessKey, true
}

// HasAccessKey returns a boolean if a field has been set.
func (o *User) HasAccessKey() bool {
	if o != nil && !IsNil(o.AccessKey) {
		return true
	}

	return false
}

// SetAccessKey gets a reference to the given string and assigns it to the AccessKey field.
func (o *User) SetAccessKey(v string) {
	o.AccessKey = &v
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *User) GetActive() bool {
	if o == nil || IsNil(o.Active) {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.Active) {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *User) HasActive() bool {
	if o != nil && !IsNil(o.Active) {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *User) SetActive(v bool) {
	o.Active = &v
}

// GetApiModel returns the ApiModel field value if set, zero value otherwise.
func (o *User) GetApiModel() map[string]interface{} {
	if o == nil || IsNil(o.ApiModel) {
		var ret map[string]interface{}
		return ret
	}
	return o.ApiModel
}

// GetApiModelOk returns a tuple with the ApiModel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetApiModelOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.ApiModel) {
		return map[string]interface{}{}, false
	}
	return o.ApiModel, true
}

// HasApiModel returns a boolean if a field has been set.
func (o *User) HasApiModel() bool {
	if o != nil && !IsNil(o.ApiModel) {
		return true
	}

	return false
}

// SetApiModel gets a reference to the given map[string]interface{} and assigns it to the ApiModel field.
func (o *User) SetApiModel(v map[string]interface{}) {
	o.ApiModel = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *User) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *User) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *User) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetEmailAddress returns the EmailAddress field value
func (o *User) GetEmailAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EmailAddress
}

// GetEmailAddressOk returns a tuple with the EmailAddress field value
// and a boolean to check if the value has been set.
func (o *User) GetEmailAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EmailAddress, true
}

// SetEmailAddress sets field value
func (o *User) SetEmailAddress(v string) {
	o.EmailAddress = v
}

// GetFirstName returns the FirstName field value
func (o *User) GetFirstName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value
// and a boolean to check if the value has been set.
func (o *User) GetFirstNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FirstName, true
}

// SetFirstName sets field value
func (o *User) SetFirstName(v string) {
	o.FirstName = v
}

// GetGroupId returns the GroupId field value if set, zero value otherwise.
func (o *User) GetGroupId() string {
	if o == nil || IsNil(o.GroupId) {
		var ret string
		return ret
	}
	return *o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetGroupIdOk() (*string, bool) {
	if o == nil || IsNil(o.GroupId) {
		return nil, false
	}
	return o.GroupId, true
}

// HasGroupId returns a boolean if a field has been set.
func (o *User) HasGroupId() bool {
	if o != nil && !IsNil(o.GroupId) {
		return true
	}

	return false
}

// SetGroupId gets a reference to the given string and assigns it to the GroupId field.
func (o *User) SetGroupId(v string) {
	o.GroupId = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *User) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *User) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *User) SetId(v string) {
	o.Id = &v
}

// GetLastLoginDate returns the LastLoginDate field value if set, zero value otherwise.
func (o *User) GetLastLoginDate() time.Time {
	if o == nil || IsNil(o.LastLoginDate) {
		var ret time.Time
		return ret
	}
	return *o.LastLoginDate
}

// GetLastLoginDateOk returns a tuple with the LastLoginDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetLastLoginDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastLoginDate) {
		return nil, false
	}
	return o.LastLoginDate, true
}

// HasLastLoginDate returns a boolean if a field has been set.
func (o *User) HasLastLoginDate() bool {
	if o != nil && !IsNil(o.LastLoginDate) {
		return true
	}

	return false
}

// SetLastLoginDate gets a reference to the given time.Time and assigns it to the LastLoginDate field.
func (o *User) SetLastLoginDate(v time.Time) {
	o.LastLoginDate = &v
}

// GetLastName returns the LastName field value
func (o *User) GetLastName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value
// and a boolean to check if the value has been set.
func (o *User) GetLastNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastName, true
}

// SetLastName sets field value
func (o *User) SetLastName(v string) {
	o.LastName = v
}

// GetOrgId returns the OrgId field value if set, zero value otherwise.
func (o *User) GetOrgId() string {
	if o == nil || IsNil(o.OrgId) {
		var ret string
		return ret
	}
	return *o.OrgId
}

// GetOrgIdOk returns a tuple with the OrgId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetOrgIdOk() (*string, bool) {
	if o == nil || IsNil(o.OrgId) {
		return nil, false
	}
	return o.OrgId, true
}

// HasOrgId returns a boolean if a field has been set.
func (o *User) HasOrgId() bool {
	if o != nil && !IsNil(o.OrgId) {
		return true
	}

	return false
}

// SetOrgId gets a reference to the given string and assigns it to the OrgId field.
func (o *User) SetOrgId(v string) {
	o.OrgId = &v
}

// GetPasswordMaxDays returns the PasswordMaxDays field value if set, zero value otherwise.
func (o *User) GetPasswordMaxDays() int32 {
	if o == nil || IsNil(o.PasswordMaxDays) {
		var ret int32
		return ret
	}
	return *o.PasswordMaxDays
}

// GetPasswordMaxDaysOk returns a tuple with the PasswordMaxDays field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetPasswordMaxDaysOk() (*int32, bool) {
	if o == nil || IsNil(o.PasswordMaxDays) {
		return nil, false
	}
	return o.PasswordMaxDays, true
}

// HasPasswordMaxDays returns a boolean if a field has been set.
func (o *User) HasPasswordMaxDays() bool {
	if o != nil && !IsNil(o.PasswordMaxDays) {
		return true
	}

	return false
}

// SetPasswordMaxDays gets a reference to the given int32 and assigns it to the PasswordMaxDays field.
func (o *User) SetPasswordMaxDays(v int32) {
	o.PasswordMaxDays = &v
}

// GetPasswordUpdated returns the PasswordUpdated field value if set, zero value otherwise.
func (o *User) GetPasswordUpdated() time.Time {
	if o == nil || IsNil(o.PasswordUpdated) {
		var ret time.Time
		return ret
	}
	return *o.PasswordUpdated
}

// GetPasswordUpdatedOk returns a tuple with the PasswordUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetPasswordUpdatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.PasswordUpdated) {
		return nil, false
	}
	return o.PasswordUpdated, true
}

// HasPasswordUpdated returns a boolean if a field has been set.
func (o *User) HasPasswordUpdated() bool {
	if o != nil && !IsNil(o.PasswordUpdated) {
		return true
	}

	return false
}

// SetPasswordUpdated gets a reference to the given time.Time and assigns it to the PasswordUpdated field.
func (o *User) SetPasswordUpdated(v time.Time) {
	o.PasswordUpdated = &v
}

// GetUserPermissions returns the UserPermissions field value
func (o *User) GetUserPermissions() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}

	return o.UserPermissions
}

// GetUserPermissionsOk returns a tuple with the UserPermissions field value
// and a boolean to check if the value has been set.
func (o *User) GetUserPermissionsOk() (*map[string]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserPermissions, true
}

// SetUserPermissions sets field value
func (o *User) SetUserPermissions(v map[string]string) {
	o.UserPermissions = v
}

func (o User) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o User) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccessKey) {
		toSerialize["access_key"] = o.AccessKey
	}
	if !IsNil(o.Active) {
		toSerialize["active"] = o.Active
	}
	if !IsNil(o.ApiModel) {
		toSerialize["api_model"] = o.ApiModel
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	toSerialize["email_address"] = o.EmailAddress
	toSerialize["first_name"] = o.FirstName
	if !IsNil(o.GroupId) {
		toSerialize["group_id"] = o.GroupId
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.LastLoginDate) {
		toSerialize["last_login_date"] = o.LastLoginDate
	}
	toSerialize["last_name"] = o.LastName
	if !IsNil(o.OrgId) {
		toSerialize["org_id"] = o.OrgId
	}
	if !IsNil(o.PasswordMaxDays) {
		toSerialize["password_max_days"] = o.PasswordMaxDays
	}
	if !IsNil(o.PasswordUpdated) {
		toSerialize["password_updated"] = o.PasswordUpdated
	}
	toSerialize["user_permissions"] = o.UserPermissions
	return toSerialize, nil
}

func (o *User) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"email_address",
		"first_name",
		"last_name",
		"user_permissions",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varUser := _User{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUser)

	if err != nil {
		return err
	}

	*o = User(varUser)

	return err
}

type NullableUser struct {
	value *User
	isSet bool
}

func (v NullableUser) Get() *User {
	return v.value
}

func (v *NullableUser) Set(val *User) {
	v.value = val
	v.isSet = true
}

func (v NullableUser) IsSet() bool {
	return v.isSet
}

func (v *NullableUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUser(val *User) *NullableUser {
	return &NullableUser{value: val, isSet: true}
}

func (v NullableUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
