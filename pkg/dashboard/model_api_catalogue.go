/*
NEW Tyk DASH API

## <a name=\"introduction\"></a> Introduction      The Tyk Dashboard API offers granular, programmatic access to a centralised database of resources that your Tyk nodes can pull from. This API has a dynamic user administrative structure which means the secret key that is used to communicate with your Tyk nodes can be kept secret and access to the wider management functions can be handled on a user-by-user and organisation-by-organisation basis.      A common question around using a database-backed configuration is how to programatically add API definitions to your Tyk nodes, the Dashboard API allows much more fine-grained, secure and multi-user access to your Tyk cluster, and should be used to manage a database-backed Tyk node.      The Tyk Dashboard API works seamlessly with the Tyk Dashboard (and the two come bundled together).      ## <a name=\"security-hierarchy\"></a> Security Hierarchy      The Dashboard API provides a more structured security layer to managing Tyk nodes.      ### Organisations, APIs and Users      With the Dashboard API and a database-backed Tyk setup, (and to an extent with file-based API setups - if diligence is used in naming an creating definitions), the following security model is applied to the management of Upstream APIs:      * **Organisations**: All APIs are *owned* by an organisation, this is designated by the OrgID parameter in the API Definition.     * **Users**: All users created in the Dashboard belong to an organisation (unless an exception is made for super-administrative access).     * **APIs**: All APIs belong to an Organisation and only Users that belong to that organisation can see the analytics for those APIs and manage their configurations.     * **API Keys**: API Keys are designated by organisation, this means an API key that has full access rights will not be allowed to access the APIs of another organisation on the same system, but can have full access to all APIs within the organisation.     * **Access Rights**: Access rights are stored with the key, this enables a key to give access to multiple APIs, this is defined by the session object in the core Tyk API.      In order to use the Dashboard API, you'll need to get the Tyk Dashboard API Access Credentials secret from your user profile on the Dashboard UI.      The secret you set should then be sent along as a header with each Dashboard API Request in order for it to be successful:  authorization: <your-secret>

API version: 5.4.0
Contact: support@tyk.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dashboard

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the ApiCatalogue type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiCatalogue{}

// ApiCatalogue struct for ApiCatalogue
type ApiCatalogue struct {
	Apis  []APIDescription `json:"apis"`
	Email *string          `json:"email,omitempty"`
	Id    *string          `json:"id,omitempty"`
	OrgId *string          `json:"org_id,omitempty"`
}

type _ApiCatalogue ApiCatalogue

// NewApiCatalogue instantiates a new ApiCatalogue object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiCatalogue(apis []APIDescription) *ApiCatalogue {
	this := ApiCatalogue{}
	this.Apis = apis
	return &this
}

// NewApiCatalogueWithDefaults instantiates a new ApiCatalogue object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiCatalogueWithDefaults() *ApiCatalogue {
	this := ApiCatalogue{}
	return &this
}

// GetApis returns the Apis field value
// If the value is explicit nil, the zero value for []APIDescription will be returned
func (o *ApiCatalogue) GetApis() []APIDescription {
	if o == nil {
		var ret []APIDescription
		return ret
	}

	return o.Apis
}

// GetApisOk returns a tuple with the Apis field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiCatalogue) GetApisOk() ([]APIDescription, bool) {
	if o == nil || IsNil(o.Apis) {
		return nil, false
	}
	return o.Apis, true
}

// SetApis sets field value
func (o *ApiCatalogue) SetApis(v []APIDescription) {
	o.Apis = v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *ApiCatalogue) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiCatalogue) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *ApiCatalogue) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *ApiCatalogue) SetEmail(v string) {
	o.Email = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ApiCatalogue) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiCatalogue) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ApiCatalogue) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ApiCatalogue) SetId(v string) {
	o.Id = &v
}

// GetOrgId returns the OrgId field value if set, zero value otherwise.
func (o *ApiCatalogue) GetOrgId() string {
	if o == nil || IsNil(o.OrgId) {
		var ret string
		return ret
	}
	return *o.OrgId
}

// GetOrgIdOk returns a tuple with the OrgId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiCatalogue) GetOrgIdOk() (*string, bool) {
	if o == nil || IsNil(o.OrgId) {
		return nil, false
	}
	return o.OrgId, true
}

// HasOrgId returns a boolean if a field has been set.
func (o *ApiCatalogue) HasOrgId() bool {
	if o != nil && !IsNil(o.OrgId) {
		return true
	}

	return false
}

// SetOrgId gets a reference to the given string and assigns it to the OrgId field.
func (o *ApiCatalogue) SetOrgId(v string) {
	o.OrgId = &v
}

func (o ApiCatalogue) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiCatalogue) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Apis != nil {
		toSerialize["apis"] = o.Apis
	}
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.OrgId) {
		toSerialize["org_id"] = o.OrgId
	}
	return toSerialize, nil
}

func (o *ApiCatalogue) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"apis",
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

	varApiCatalogue := _ApiCatalogue{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varApiCatalogue)

	if err != nil {
		return err
	}

	*o = ApiCatalogue(varApiCatalogue)

	return err
}

type NullableApiCatalogue struct {
	value *ApiCatalogue
	isSet bool
}

func (v NullableApiCatalogue) Get() *ApiCatalogue {
	return v.value
}

func (v *NullableApiCatalogue) Set(val *ApiCatalogue) {
	v.value = val
	v.isSet = true
}

func (v NullableApiCatalogue) IsSet() bool {
	return v.isSet
}

func (v *NullableApiCatalogue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiCatalogue(val *ApiCatalogue) *NullableApiCatalogue {
	return &NullableApiCatalogue{value: val, isSet: true}
}

func (v NullableApiCatalogue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiCatalogue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
