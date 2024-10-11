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

// checks if the MethodTransformMeta type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MethodTransformMeta{}

// MethodTransformMeta struct for MethodTransformMeta
type MethodTransformMeta struct {
	Disabled *bool   `json:"disabled,omitempty"`
	Method   *string `json:"method,omitempty"`
	Path     *string `json:"path,omitempty"`
	ToMethod *string `json:"to_method,omitempty"`
}

// NewMethodTransformMeta instantiates a new MethodTransformMeta object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMethodTransformMeta() *MethodTransformMeta {
	this := MethodTransformMeta{}
	return &this
}

// NewMethodTransformMetaWithDefaults instantiates a new MethodTransformMeta object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMethodTransformMetaWithDefaults() *MethodTransformMeta {
	this := MethodTransformMeta{}
	return &this
}

// GetDisabled returns the Disabled field value if set, zero value otherwise.
func (o *MethodTransformMeta) GetDisabled() bool {
	if o == nil || IsNil(o.Disabled) {
		var ret bool
		return ret
	}
	return *o.Disabled
}

// GetDisabledOk returns a tuple with the Disabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MethodTransformMeta) GetDisabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Disabled) {
		return nil, false
	}
	return o.Disabled, true
}

// HasDisabled returns a boolean if a field has been set.
func (o *MethodTransformMeta) HasDisabled() bool {
	if o != nil && !IsNil(o.Disabled) {
		return true
	}

	return false
}

// SetDisabled gets a reference to the given bool and assigns it to the Disabled field.
func (o *MethodTransformMeta) SetDisabled(v bool) {
	o.Disabled = &v
}

// GetMethod returns the Method field value if set, zero value otherwise.
func (o *MethodTransformMeta) GetMethod() string {
	if o == nil || IsNil(o.Method) {
		var ret string
		return ret
	}
	return *o.Method
}

// GetMethodOk returns a tuple with the Method field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MethodTransformMeta) GetMethodOk() (*string, bool) {
	if o == nil || IsNil(o.Method) {
		return nil, false
	}
	return o.Method, true
}

// HasMethod returns a boolean if a field has been set.
func (o *MethodTransformMeta) HasMethod() bool {
	if o != nil && !IsNil(o.Method) {
		return true
	}

	return false
}

// SetMethod gets a reference to the given string and assigns it to the Method field.
func (o *MethodTransformMeta) SetMethod(v string) {
	o.Method = &v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *MethodTransformMeta) GetPath() string {
	if o == nil || IsNil(o.Path) {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MethodTransformMeta) GetPathOk() (*string, bool) {
	if o == nil || IsNil(o.Path) {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *MethodTransformMeta) HasPath() bool {
	if o != nil && !IsNil(o.Path) {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *MethodTransformMeta) SetPath(v string) {
	o.Path = &v
}

// GetToMethod returns the ToMethod field value if set, zero value otherwise.
func (o *MethodTransformMeta) GetToMethod() string {
	if o == nil || IsNil(o.ToMethod) {
		var ret string
		return ret
	}
	return *o.ToMethod
}

// GetToMethodOk returns a tuple with the ToMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MethodTransformMeta) GetToMethodOk() (*string, bool) {
	if o == nil || IsNil(o.ToMethod) {
		return nil, false
	}
	return o.ToMethod, true
}

// HasToMethod returns a boolean if a field has been set.
func (o *MethodTransformMeta) HasToMethod() bool {
	if o != nil && !IsNil(o.ToMethod) {
		return true
	}

	return false
}

// SetToMethod gets a reference to the given string and assigns it to the ToMethod field.
func (o *MethodTransformMeta) SetToMethod(v string) {
	o.ToMethod = &v
}

func (o MethodTransformMeta) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MethodTransformMeta) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Disabled) {
		toSerialize["disabled"] = o.Disabled
	}
	if !IsNil(o.Method) {
		toSerialize["method"] = o.Method
	}
	if !IsNil(o.Path) {
		toSerialize["path"] = o.Path
	}
	if !IsNil(o.ToMethod) {
		toSerialize["to_method"] = o.ToMethod
	}
	return toSerialize, nil
}

type NullableMethodTransformMeta struct {
	value *MethodTransformMeta
	isSet bool
}

func (v NullableMethodTransformMeta) Get() *MethodTransformMeta {
	return v.value
}

func (v *NullableMethodTransformMeta) Set(val *MethodTransformMeta) {
	v.value = val
	v.isSet = true
}

func (v NullableMethodTransformMeta) IsSet() bool {
	return v.isSet
}

func (v *NullableMethodTransformMeta) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMethodTransformMeta(val *MethodTransformMeta) *NullableMethodTransformMeta {
	return &NullableMethodTransformMeta{value: val, isSet: true}
}

func (v NullableMethodTransformMeta) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMethodTransformMeta) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
