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

// checks if the OAuthApps type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OAuthApps{}

// OAuthApps struct for OAuthApps
type OAuthApps struct {
	Apps  []OAuthClient `json:"apps,omitempty"`
	Pages *int32        `json:"pages,omitempty"`
}

// NewOAuthApps instantiates a new OAuthApps object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOAuthApps() *OAuthApps {
	this := OAuthApps{}
	return &this
}

// NewOAuthAppsWithDefaults instantiates a new OAuthApps object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOAuthAppsWithDefaults() *OAuthApps {
	this := OAuthApps{}
	return &this
}

// GetApps returns the Apps field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OAuthApps) GetApps() []OAuthClient {
	if o == nil {
		var ret []OAuthClient
		return ret
	}
	return o.Apps
}

// GetAppsOk returns a tuple with the Apps field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OAuthApps) GetAppsOk() ([]OAuthClient, bool) {
	if o == nil || IsNil(o.Apps) {
		return nil, false
	}
	return o.Apps, true
}

// HasApps returns a boolean if a field has been set.
func (o *OAuthApps) HasApps() bool {
	if o != nil && !IsNil(o.Apps) {
		return true
	}

	return false
}

// SetApps gets a reference to the given []OAuthClient and assigns it to the Apps field.
func (o *OAuthApps) SetApps(v []OAuthClient) {
	o.Apps = v
}

// GetPages returns the Pages field value if set, zero value otherwise.
func (o *OAuthApps) GetPages() int32 {
	if o == nil || IsNil(o.Pages) {
		var ret int32
		return ret
	}
	return *o.Pages
}

// GetPagesOk returns a tuple with the Pages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuthApps) GetPagesOk() (*int32, bool) {
	if o == nil || IsNil(o.Pages) {
		return nil, false
	}
	return o.Pages, true
}

// HasPages returns a boolean if a field has been set.
func (o *OAuthApps) HasPages() bool {
	if o != nil && !IsNil(o.Pages) {
		return true
	}

	return false
}

// SetPages gets a reference to the given int32 and assigns it to the Pages field.
func (o *OAuthApps) SetPages(v int32) {
	o.Pages = &v
}

func (o OAuthApps) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OAuthApps) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Apps != nil {
		toSerialize["apps"] = o.Apps
	}
	if !IsNil(o.Pages) {
		toSerialize["pages"] = o.Pages
	}
	return toSerialize, nil
}

type NullableOAuthApps struct {
	value *OAuthApps
	isSet bool
}

func (v NullableOAuthApps) Get() *OAuthApps {
	return v.value
}

func (v *NullableOAuthApps) Set(val *OAuthApps) {
	v.value = val
	v.isSet = true
}

func (v NullableOAuthApps) IsSet() bool {
	return v.isSet
}

func (v *NullableOAuthApps) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOAuthApps(val *OAuthApps) *NullableOAuthApps {
	return &NullableOAuthApps{value: val, isSet: true}
}

func (v NullableOAuthApps) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOAuthApps) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
