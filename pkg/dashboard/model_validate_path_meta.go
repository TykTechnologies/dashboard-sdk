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

// checks if the ValidatePathMeta type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ValidatePathMeta{}

// ValidatePathMeta struct for ValidatePathMeta
type ValidatePathMeta struct {
	Disabled          *bool                  `json:"disabled,omitempty"`
	ErrorResponseCode *int32                 `json:"error_response_code,omitempty"`
	Method            *string                `json:"method,omitempty"`
	Path              *string                `json:"path,omitempty"`
	Schema            map[string]interface{} `json:"schema,omitempty"`
	SchemaB64         *string                `json:"schema_b64,omitempty"`
}

// NewValidatePathMeta instantiates a new ValidatePathMeta object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewValidatePathMeta() *ValidatePathMeta {
	this := ValidatePathMeta{}
	return &this
}

// NewValidatePathMetaWithDefaults instantiates a new ValidatePathMeta object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewValidatePathMetaWithDefaults() *ValidatePathMeta {
	this := ValidatePathMeta{}
	return &this
}

// GetDisabled returns the Disabled field value if set, zero value otherwise.
func (o *ValidatePathMeta) GetDisabled() bool {
	if o == nil || IsNil(o.Disabled) {
		var ret bool
		return ret
	}
	return *o.Disabled
}

// GetDisabledOk returns a tuple with the Disabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidatePathMeta) GetDisabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Disabled) {
		return nil, false
	}
	return o.Disabled, true
}

// HasDisabled returns a boolean if a field has been set.
func (o *ValidatePathMeta) HasDisabled() bool {
	if o != nil && !IsNil(o.Disabled) {
		return true
	}

	return false
}

// SetDisabled gets a reference to the given bool and assigns it to the Disabled field.
func (o *ValidatePathMeta) SetDisabled(v bool) {
	o.Disabled = &v
}

// GetErrorResponseCode returns the ErrorResponseCode field value if set, zero value otherwise.
func (o *ValidatePathMeta) GetErrorResponseCode() int32 {
	if o == nil || IsNil(o.ErrorResponseCode) {
		var ret int32
		return ret
	}
	return *o.ErrorResponseCode
}

// GetErrorResponseCodeOk returns a tuple with the ErrorResponseCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidatePathMeta) GetErrorResponseCodeOk() (*int32, bool) {
	if o == nil || IsNil(o.ErrorResponseCode) {
		return nil, false
	}
	return o.ErrorResponseCode, true
}

// HasErrorResponseCode returns a boolean if a field has been set.
func (o *ValidatePathMeta) HasErrorResponseCode() bool {
	if o != nil && !IsNil(o.ErrorResponseCode) {
		return true
	}

	return false
}

// SetErrorResponseCode gets a reference to the given int32 and assigns it to the ErrorResponseCode field.
func (o *ValidatePathMeta) SetErrorResponseCode(v int32) {
	o.ErrorResponseCode = &v
}

// GetMethod returns the Method field value if set, zero value otherwise.
func (o *ValidatePathMeta) GetMethod() string {
	if o == nil || IsNil(o.Method) {
		var ret string
		return ret
	}
	return *o.Method
}

// GetMethodOk returns a tuple with the Method field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidatePathMeta) GetMethodOk() (*string, bool) {
	if o == nil || IsNil(o.Method) {
		return nil, false
	}
	return o.Method, true
}

// HasMethod returns a boolean if a field has been set.
func (o *ValidatePathMeta) HasMethod() bool {
	if o != nil && !IsNil(o.Method) {
		return true
	}

	return false
}

// SetMethod gets a reference to the given string and assigns it to the Method field.
func (o *ValidatePathMeta) SetMethod(v string) {
	o.Method = &v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *ValidatePathMeta) GetPath() string {
	if o == nil || IsNil(o.Path) {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidatePathMeta) GetPathOk() (*string, bool) {
	if o == nil || IsNil(o.Path) {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *ValidatePathMeta) HasPath() bool {
	if o != nil && !IsNil(o.Path) {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *ValidatePathMeta) SetPath(v string) {
	o.Path = &v
}

// GetSchema returns the Schema field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ValidatePathMeta) GetSchema() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Schema
}

// GetSchemaOk returns a tuple with the Schema field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ValidatePathMeta) GetSchemaOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Schema) {
		return map[string]interface{}{}, false
	}
	return o.Schema, true
}

// HasSchema returns a boolean if a field has been set.
func (o *ValidatePathMeta) HasSchema() bool {
	if o != nil && !IsNil(o.Schema) {
		return true
	}

	return false
}

// SetSchema gets a reference to the given map[string]interface{} and assigns it to the Schema field.
func (o *ValidatePathMeta) SetSchema(v map[string]interface{}) {
	o.Schema = v
}

// GetSchemaB64 returns the SchemaB64 field value if set, zero value otherwise.
func (o *ValidatePathMeta) GetSchemaB64() string {
	if o == nil || IsNil(o.SchemaB64) {
		var ret string
		return ret
	}
	return *o.SchemaB64
}

// GetSchemaB64Ok returns a tuple with the SchemaB64 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidatePathMeta) GetSchemaB64Ok() (*string, bool) {
	if o == nil || IsNil(o.SchemaB64) {
		return nil, false
	}
	return o.SchemaB64, true
}

// HasSchemaB64 returns a boolean if a field has been set.
func (o *ValidatePathMeta) HasSchemaB64() bool {
	if o != nil && !IsNil(o.SchemaB64) {
		return true
	}

	return false
}

// SetSchemaB64 gets a reference to the given string and assigns it to the SchemaB64 field.
func (o *ValidatePathMeta) SetSchemaB64(v string) {
	o.SchemaB64 = &v
}

func (o ValidatePathMeta) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ValidatePathMeta) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Disabled) {
		toSerialize["disabled"] = o.Disabled
	}
	if !IsNil(o.ErrorResponseCode) {
		toSerialize["error_response_code"] = o.ErrorResponseCode
	}
	if !IsNil(o.Method) {
		toSerialize["method"] = o.Method
	}
	if !IsNil(o.Path) {
		toSerialize["path"] = o.Path
	}
	if o.Schema != nil {
		toSerialize["schema"] = o.Schema
	}
	if !IsNil(o.SchemaB64) {
		toSerialize["schema_b64"] = o.SchemaB64
	}
	return toSerialize, nil
}

type NullableValidatePathMeta struct {
	value *ValidatePathMeta
	isSet bool
}

func (v NullableValidatePathMeta) Get() *ValidatePathMeta {
	return v.value
}

func (v *NullableValidatePathMeta) Set(val *ValidatePathMeta) {
	v.value = val
	v.isSet = true
}

func (v NullableValidatePathMeta) IsSet() bool {
	return v.isSet
}

func (v *NullableValidatePathMeta) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableValidatePathMeta(val *ValidatePathMeta) *NullableValidatePathMeta {
	return &NullableValidatePathMeta{value: val, isSet: true}
}

func (v NullableValidatePathMeta) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableValidatePathMeta) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
