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

// checks if the SSOAccessData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SSOAccessData{}

// SSOAccessData struct for SSOAccessData
type SSOAccessData struct {
	DisplayName               *string `json:"DisplayName,omitempty"`
	EmailAddress              *string `json:"EmailAddress,omitempty"`
	ForSection                *string `json:"ForSection,omitempty"`
	GroupID                   *string `json:"GroupID,omitempty"`
	OrgID                     *string `json:"OrgID,omitempty"`
	SSOOnlyForRegisteredUsers *bool   `json:"SSOOnlyForRegisteredUsers,omitempty"`
	UserNotAllowed            *bool   `json:"UserNotAllowed,omitempty"`
}

// NewSSOAccessData instantiates a new SSOAccessData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSSOAccessData() *SSOAccessData {
	this := SSOAccessData{}
	return &this
}

// NewSSOAccessDataWithDefaults instantiates a new SSOAccessData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSSOAccessDataWithDefaults() *SSOAccessData {
	this := SSOAccessData{}
	return &this
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *SSOAccessData) GetDisplayName() string {
	if o == nil || IsNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SSOAccessData) GetDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *SSOAccessData) HasDisplayName() bool {
	if o != nil && !IsNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *SSOAccessData) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetEmailAddress returns the EmailAddress field value if set, zero value otherwise.
func (o *SSOAccessData) GetEmailAddress() string {
	if o == nil || IsNil(o.EmailAddress) {
		var ret string
		return ret
	}
	return *o.EmailAddress
}

// GetEmailAddressOk returns a tuple with the EmailAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SSOAccessData) GetEmailAddressOk() (*string, bool) {
	if o == nil || IsNil(o.EmailAddress) {
		return nil, false
	}
	return o.EmailAddress, true
}

// HasEmailAddress returns a boolean if a field has been set.
func (o *SSOAccessData) HasEmailAddress() bool {
	if o != nil && !IsNil(o.EmailAddress) {
		return true
	}

	return false
}

// SetEmailAddress gets a reference to the given string and assigns it to the EmailAddress field.
func (o *SSOAccessData) SetEmailAddress(v string) {
	o.EmailAddress = &v
}

// GetForSection returns the ForSection field value if set, zero value otherwise.
func (o *SSOAccessData) GetForSection() string {
	if o == nil || IsNil(o.ForSection) {
		var ret string
		return ret
	}
	return *o.ForSection
}

// GetForSectionOk returns a tuple with the ForSection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SSOAccessData) GetForSectionOk() (*string, bool) {
	if o == nil || IsNil(o.ForSection) {
		return nil, false
	}
	return o.ForSection, true
}

// HasForSection returns a boolean if a field has been set.
func (o *SSOAccessData) HasForSection() bool {
	if o != nil && !IsNil(o.ForSection) {
		return true
	}

	return false
}

// SetForSection gets a reference to the given string and assigns it to the ForSection field.
func (o *SSOAccessData) SetForSection(v string) {
	o.ForSection = &v
}

// GetGroupID returns the GroupID field value if set, zero value otherwise.
func (o *SSOAccessData) GetGroupID() string {
	if o == nil || IsNil(o.GroupID) {
		var ret string
		return ret
	}
	return *o.GroupID
}

// GetGroupIDOk returns a tuple with the GroupID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SSOAccessData) GetGroupIDOk() (*string, bool) {
	if o == nil || IsNil(o.GroupID) {
		return nil, false
	}
	return o.GroupID, true
}

// HasGroupID returns a boolean if a field has been set.
func (o *SSOAccessData) HasGroupID() bool {
	if o != nil && !IsNil(o.GroupID) {
		return true
	}

	return false
}

// SetGroupID gets a reference to the given string and assigns it to the GroupID field.
func (o *SSOAccessData) SetGroupID(v string) {
	o.GroupID = &v
}

// GetOrgID returns the OrgID field value if set, zero value otherwise.
func (o *SSOAccessData) GetOrgID() string {
	if o == nil || IsNil(o.OrgID) {
		var ret string
		return ret
	}
	return *o.OrgID
}

// GetOrgIDOk returns a tuple with the OrgID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SSOAccessData) GetOrgIDOk() (*string, bool) {
	if o == nil || IsNil(o.OrgID) {
		return nil, false
	}
	return o.OrgID, true
}

// HasOrgID returns a boolean if a field has been set.
func (o *SSOAccessData) HasOrgID() bool {
	if o != nil && !IsNil(o.OrgID) {
		return true
	}

	return false
}

// SetOrgID gets a reference to the given string and assigns it to the OrgID field.
func (o *SSOAccessData) SetOrgID(v string) {
	o.OrgID = &v
}

// GetSSOOnlyForRegisteredUsers returns the SSOOnlyForRegisteredUsers field value if set, zero value otherwise.
func (o *SSOAccessData) GetSSOOnlyForRegisteredUsers() bool {
	if o == nil || IsNil(o.SSOOnlyForRegisteredUsers) {
		var ret bool
		return ret
	}
	return *o.SSOOnlyForRegisteredUsers
}

// GetSSOOnlyForRegisteredUsersOk returns a tuple with the SSOOnlyForRegisteredUsers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SSOAccessData) GetSSOOnlyForRegisteredUsersOk() (*bool, bool) {
	if o == nil || IsNil(o.SSOOnlyForRegisteredUsers) {
		return nil, false
	}
	return o.SSOOnlyForRegisteredUsers, true
}

// HasSSOOnlyForRegisteredUsers returns a boolean if a field has been set.
func (o *SSOAccessData) HasSSOOnlyForRegisteredUsers() bool {
	if o != nil && !IsNil(o.SSOOnlyForRegisteredUsers) {
		return true
	}

	return false
}

// SetSSOOnlyForRegisteredUsers gets a reference to the given bool and assigns it to the SSOOnlyForRegisteredUsers field.
func (o *SSOAccessData) SetSSOOnlyForRegisteredUsers(v bool) {
	o.SSOOnlyForRegisteredUsers = &v
}

// GetUserNotAllowed returns the UserNotAllowed field value if set, zero value otherwise.
func (o *SSOAccessData) GetUserNotAllowed() bool {
	if o == nil || IsNil(o.UserNotAllowed) {
		var ret bool
		return ret
	}
	return *o.UserNotAllowed
}

// GetUserNotAllowedOk returns a tuple with the UserNotAllowed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SSOAccessData) GetUserNotAllowedOk() (*bool, bool) {
	if o == nil || IsNil(o.UserNotAllowed) {
		return nil, false
	}
	return o.UserNotAllowed, true
}

// HasUserNotAllowed returns a boolean if a field has been set.
func (o *SSOAccessData) HasUserNotAllowed() bool {
	if o != nil && !IsNil(o.UserNotAllowed) {
		return true
	}

	return false
}

// SetUserNotAllowed gets a reference to the given bool and assigns it to the UserNotAllowed field.
func (o *SSOAccessData) SetUserNotAllowed(v bool) {
	o.UserNotAllowed = &v
}

func (o SSOAccessData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SSOAccessData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DisplayName) {
		toSerialize["DisplayName"] = o.DisplayName
	}
	if !IsNil(o.EmailAddress) {
		toSerialize["EmailAddress"] = o.EmailAddress
	}
	if !IsNil(o.ForSection) {
		toSerialize["ForSection"] = o.ForSection
	}
	if !IsNil(o.GroupID) {
		toSerialize["GroupID"] = o.GroupID
	}
	if !IsNil(o.OrgID) {
		toSerialize["OrgID"] = o.OrgID
	}
	if !IsNil(o.SSOOnlyForRegisteredUsers) {
		toSerialize["SSOOnlyForRegisteredUsers"] = o.SSOOnlyForRegisteredUsers
	}
	if !IsNil(o.UserNotAllowed) {
		toSerialize["UserNotAllowed"] = o.UserNotAllowed
	}
	return toSerialize, nil
}

type NullableSSOAccessData struct {
	value *SSOAccessData
	isSet bool
}

func (v NullableSSOAccessData) Get() *SSOAccessData {
	return v.value
}

func (v *NullableSSOAccessData) Set(val *SSOAccessData) {
	v.value = val
	v.isSet = true
}

func (v NullableSSOAccessData) IsSet() bool {
	return v.isSet
}

func (v *NullableSSOAccessData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSSOAccessData(val *SSOAccessData) *NullableSSOAccessData {
	return &NullableSSOAccessData{value: val, isSet: true}
}

func (v NullableSSOAccessData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSSOAccessData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
