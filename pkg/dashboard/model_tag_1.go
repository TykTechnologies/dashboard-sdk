/*
Tyk Dashboard API

 ## <a name=\"introduction\"></a> Introduction  The Tyk Dashboard API offers granular, programmatic access to a centralised database of resources that your Tyk nodes can pull from. This API has a dynamic user administrative structure which means the secret key that is used to communicate with your Tyk nodes can be kept secret and access to the wider management functions can be handled on a user-by-user and organisation-by-organisation basis.  A common question around using a database-backed configuration is how to programmatically add API definitions to your Tyk nodes, the Dashboard API allows much more fine-grained, secure and multi-user access to your Tyk cluster, and should be used to manage a database-backed Tyk node.  The Tyk Dashboard API works seamlessly with the Tyk Dashboard (and the two come bundled together).  ## <a name=\"security-hierarchy\"></a> Security Hierarchy  The Dashboard API provides a more structured security layer to managing Tyk nodes.  ### Organisations, APIs and Users  With the Dashboard API and a database-backed Tyk setup, (and to an extent with file-based API setups - if diligence is used in naming and creating definitions), the following security model is applied to the management of Upstream APIs:  * **Organisations**: All APIs are *owned* by an organisation, this is designated by the 'OrgID' parameter in the API Definition. * **Users**: All users created in the Dashboard belong to an organisation (unless an exception is made for super-administrative access). * **APIs**: All APIs belong to an Organisation and only Users that belong to that organisation can see the analytics for those APIs and manage their configurations. * **API Keys**: API Keys are designated by organisation, this means an API key that has full access rights will not be allowed to access the APIs of another organisation on the same system, but can have full access to all APIs within the organisation. * **Access Rights**: Access rights are stored with the key, this enables a key to give access to multiple APIs, this is defined by the session object in the core Tyk API.  In order to use the Dashboard API, you'll need to get the 'Tyk Dashboard API Access Credentials' secret from your user profile on the Dashboard UI.  The secret you set should then be sent along as a header with each Dashboard API Request in order for it to be successful:   authorization: <your-secret>

API version: 5.7.0
Contact: support@tyk.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dashboard

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the Tag1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Tag1{}

// Tag1 struct for Tag1
type Tag1 struct {
	Name         string                 `json:"name"`
	Description  *string                `json:"description,omitempty"`
	ExternalDocs *ExternalDocumentation `json:"externalDocs,omitempty"`
}

type _Tag1 Tag1

// NewTag1 instantiates a new Tag1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTag1(name string) *Tag1 {
	this := Tag1{}
	this.Name = name
	return &this
}

// NewTag1WithDefaults instantiates a new Tag1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTag1WithDefaults() *Tag1 {
	this := Tag1{}
	return &this
}

// GetName returns the Name field value
func (o *Tag1) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Tag1) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Tag1) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Tag1) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Tag1) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Tag1) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Tag1) SetDescription(v string) {
	o.Description = &v
}

// GetExternalDocs returns the ExternalDocs field value if set, zero value otherwise.
func (o *Tag1) GetExternalDocs() ExternalDocumentation {
	if o == nil || IsNil(o.ExternalDocs) {
		var ret ExternalDocumentation
		return ret
	}
	return *o.ExternalDocs
}

// GetExternalDocsOk returns a tuple with the ExternalDocs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Tag1) GetExternalDocsOk() (*ExternalDocumentation, bool) {
	if o == nil || IsNil(o.ExternalDocs) {
		return nil, false
	}
	return o.ExternalDocs, true
}

// HasExternalDocs returns a boolean if a field has been set.
func (o *Tag1) HasExternalDocs() bool {
	if o != nil && !IsNil(o.ExternalDocs) {
		return true
	}

	return false
}

// SetExternalDocs gets a reference to the given ExternalDocumentation and assigns it to the ExternalDocs field.
func (o *Tag1) SetExternalDocs(v ExternalDocumentation) {
	o.ExternalDocs = &v
}

func (o Tag1) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Tag1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.ExternalDocs) {
		toSerialize["externalDocs"] = o.ExternalDocs
	}
	return toSerialize, nil
}

func (o *Tag1) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
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

	varTag1 := _Tag1{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTag1)

	if err != nil {
		return err
	}

	*o = Tag1(varTag1)

	return err
}

type NullableTag1 struct {
	value *Tag1
	isSet bool
}

func (v NullableTag1) Get() *Tag1 {
	return v.value
}

func (v *NullableTag1) Set(val *Tag1) {
	v.value = val
	v.isSet = true
}

func (v NullableTag1) IsSet() bool {
	return v.isSet
}

func (v *NullableTag1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTag1(val *Tag1) *NullableTag1 {
	return &NullableTag1{value: val, isSet: true}
}

func (v NullableTag1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTag1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
