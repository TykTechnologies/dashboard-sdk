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

// checks if the HeaderInjectionMeta type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HeaderInjectionMeta{}

// HeaderInjectionMeta struct for HeaderInjectionMeta
type HeaderInjectionMeta struct {
	ActOn         *bool             `json:"act_on,omitempty"`
	AddHeaders    map[string]string `json:"add_headers,omitempty"`
	DeleteHeaders []string          `json:"delete_headers,omitempty"`
	Disabled      *bool             `json:"disabled,omitempty"`
	Method        *string           `json:"method,omitempty"`
	Path          *string           `json:"path,omitempty"`
}

// NewHeaderInjectionMeta instantiates a new HeaderInjectionMeta object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHeaderInjectionMeta() *HeaderInjectionMeta {
	this := HeaderInjectionMeta{}
	return &this
}

// NewHeaderInjectionMetaWithDefaults instantiates a new HeaderInjectionMeta object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHeaderInjectionMetaWithDefaults() *HeaderInjectionMeta {
	this := HeaderInjectionMeta{}
	return &this
}

// GetActOn returns the ActOn field value if set, zero value otherwise.
func (o *HeaderInjectionMeta) GetActOn() bool {
	if o == nil || IsNil(o.ActOn) {
		var ret bool
		return ret
	}
	return *o.ActOn
}

// GetActOnOk returns a tuple with the ActOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HeaderInjectionMeta) GetActOnOk() (*bool, bool) {
	if o == nil || IsNil(o.ActOn) {
		return nil, false
	}
	return o.ActOn, true
}

// HasActOn returns a boolean if a field has been set.
func (o *HeaderInjectionMeta) HasActOn() bool {
	if o != nil && !IsNil(o.ActOn) {
		return true
	}

	return false
}

// SetActOn gets a reference to the given bool and assigns it to the ActOn field.
func (o *HeaderInjectionMeta) SetActOn(v bool) {
	o.ActOn = &v
}

// GetAddHeaders returns the AddHeaders field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HeaderInjectionMeta) GetAddHeaders() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}
	return o.AddHeaders
}

// GetAddHeadersOk returns a tuple with the AddHeaders field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HeaderInjectionMeta) GetAddHeadersOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.AddHeaders) {
		return nil, false
	}
	return &o.AddHeaders, true
}

// HasAddHeaders returns a boolean if a field has been set.
func (o *HeaderInjectionMeta) HasAddHeaders() bool {
	if o != nil && !IsNil(o.AddHeaders) {
		return true
	}

	return false
}

// SetAddHeaders gets a reference to the given map[string]string and assigns it to the AddHeaders field.
func (o *HeaderInjectionMeta) SetAddHeaders(v map[string]string) {
	o.AddHeaders = v
}

// GetDeleteHeaders returns the DeleteHeaders field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HeaderInjectionMeta) GetDeleteHeaders() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.DeleteHeaders
}

// GetDeleteHeadersOk returns a tuple with the DeleteHeaders field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HeaderInjectionMeta) GetDeleteHeadersOk() ([]string, bool) {
	if o == nil || IsNil(o.DeleteHeaders) {
		return nil, false
	}
	return o.DeleteHeaders, true
}

// HasDeleteHeaders returns a boolean if a field has been set.
func (o *HeaderInjectionMeta) HasDeleteHeaders() bool {
	if o != nil && !IsNil(o.DeleteHeaders) {
		return true
	}

	return false
}

// SetDeleteHeaders gets a reference to the given []string and assigns it to the DeleteHeaders field.
func (o *HeaderInjectionMeta) SetDeleteHeaders(v []string) {
	o.DeleteHeaders = v
}

// GetDisabled returns the Disabled field value if set, zero value otherwise.
func (o *HeaderInjectionMeta) GetDisabled() bool {
	if o == nil || IsNil(o.Disabled) {
		var ret bool
		return ret
	}
	return *o.Disabled
}

// GetDisabledOk returns a tuple with the Disabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HeaderInjectionMeta) GetDisabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Disabled) {
		return nil, false
	}
	return o.Disabled, true
}

// HasDisabled returns a boolean if a field has been set.
func (o *HeaderInjectionMeta) HasDisabled() bool {
	if o != nil && !IsNil(o.Disabled) {
		return true
	}

	return false
}

// SetDisabled gets a reference to the given bool and assigns it to the Disabled field.
func (o *HeaderInjectionMeta) SetDisabled(v bool) {
	o.Disabled = &v
}

// GetMethod returns the Method field value if set, zero value otherwise.
func (o *HeaderInjectionMeta) GetMethod() string {
	if o == nil || IsNil(o.Method) {
		var ret string
		return ret
	}
	return *o.Method
}

// GetMethodOk returns a tuple with the Method field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HeaderInjectionMeta) GetMethodOk() (*string, bool) {
	if o == nil || IsNil(o.Method) {
		return nil, false
	}
	return o.Method, true
}

// HasMethod returns a boolean if a field has been set.
func (o *HeaderInjectionMeta) HasMethod() bool {
	if o != nil && !IsNil(o.Method) {
		return true
	}

	return false
}

// SetMethod gets a reference to the given string and assigns it to the Method field.
func (o *HeaderInjectionMeta) SetMethod(v string) {
	o.Method = &v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *HeaderInjectionMeta) GetPath() string {
	if o == nil || IsNil(o.Path) {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HeaderInjectionMeta) GetPathOk() (*string, bool) {
	if o == nil || IsNil(o.Path) {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *HeaderInjectionMeta) HasPath() bool {
	if o != nil && !IsNil(o.Path) {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *HeaderInjectionMeta) SetPath(v string) {
	o.Path = &v
}

func (o HeaderInjectionMeta) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HeaderInjectionMeta) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ActOn) {
		toSerialize["act_on"] = o.ActOn
	}
	if o.AddHeaders != nil {
		toSerialize["add_headers"] = o.AddHeaders
	}
	if o.DeleteHeaders != nil {
		toSerialize["delete_headers"] = o.DeleteHeaders
	}
	if !IsNil(o.Disabled) {
		toSerialize["disabled"] = o.Disabled
	}
	if !IsNil(o.Method) {
		toSerialize["method"] = o.Method
	}
	if !IsNil(o.Path) {
		toSerialize["path"] = o.Path
	}
	return toSerialize, nil
}

type NullableHeaderInjectionMeta struct {
	value *HeaderInjectionMeta
	isSet bool
}

func (v NullableHeaderInjectionMeta) Get() *HeaderInjectionMeta {
	return v.value
}

func (v *NullableHeaderInjectionMeta) Set(val *HeaderInjectionMeta) {
	v.value = val
	v.isSet = true
}

func (v NullableHeaderInjectionMeta) IsSet() bool {
	return v.isSet
}

func (v *NullableHeaderInjectionMeta) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHeaderInjectionMeta(val *HeaderInjectionMeta) *NullableHeaderInjectionMeta {
	return &NullableHeaderInjectionMeta{value: val, isSet: true}
}

func (v NullableHeaderInjectionMeta) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHeaderInjectionMeta) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
