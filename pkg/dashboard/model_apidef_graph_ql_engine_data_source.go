/*
NEW Tyk DASH API

## <a name=\"introduction\"></a> Introduction      The Tyk Dashboard API offers granular, programmatic access to a centralised database of resources that your Tyk nodes can pull from. This API has a dynamic user administrative structure which means the secret key that is used to communicate with your Tyk nodes can be kept secret and access to the wider management functions can be handled on a user-by-user and organisation-by-organisation basis.      A common question around using a database-backed configuration is how to programatically add API definitions to your Tyk nodes, the Dashboard API allows much more fine-grained, secure and multi-user access to your Tyk cluster, and should be used to manage a database-backed Tyk node.      The Tyk Dashboard API works seamlessly with the Tyk Dashboard (and the two come bundled together).      ## <a name=\"security-hierarchy\"></a> Security Hierarchy      The Dashboard API provides a more structured security layer to managing Tyk nodes.      ### Organisations, APIs and Users      With the Dashboard API and a database-backed Tyk setup, (and to an extent with file-based API setups - if diligence is used in naming an creating definitions), the following security model is applied to the management of Upstream APIs:      * **Organisations**: All APIs are *owned* by an organisation, this is designated by the OrgID parameter in the API Definition.     * **Users**: All users created in the Dashboard belong to an organisation (unless an exception is made for super-administrative access).     * **APIs**: All APIs belong to an Organisation and only Users that belong to that organisation can see the analytics for those APIs and manage their configurations.     * **API Keys**: API Keys are designated by organisation, this means an API key that has full access rights will not be allowed to access the APIs of another organisation on the same system, but can have full access to all APIs within the organisation.     * **Access Rights**: Access rights are stored with the key, this enables a key to give access to multiple APIs, this is defined by the session object in the core Tyk API.      In order to use the Dashboard API, you'll need to get the Tyk Dashboard API Access Credentials secret from your user profile on the Dashboard UI.      The secret you set should then be sent along as a header with each Dashboard API Request in order for it to be successful:  authorization: <your-secret>

API version: 5.4.0
Contact: support@tyk.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dashboard

import (
	"encoding/json"
)

// checks if the ApidefGraphQLEngineDataSource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApidefGraphQLEngineDataSource{}

// ApidefGraphQLEngineDataSource struct for ApidefGraphQLEngineDataSource
type ApidefGraphQLEngineDataSource struct {
	Config     interface{}               `json:"config,omitempty"`
	Internal   *bool                     `json:"internal,omitempty"`
	Kind       *string                   `json:"kind,omitempty"`
	Name       *string                   `json:"name,omitempty"`
	RootFields []ApidefGraphQLTypeFields `json:"root_fields,omitempty"`
}

// NewApidefGraphQLEngineDataSource instantiates a new ApidefGraphQLEngineDataSource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApidefGraphQLEngineDataSource() *ApidefGraphQLEngineDataSource {
	this := ApidefGraphQLEngineDataSource{}
	return &this
}

// NewApidefGraphQLEngineDataSourceWithDefaults instantiates a new ApidefGraphQLEngineDataSource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApidefGraphQLEngineDataSourceWithDefaults() *ApidefGraphQLEngineDataSource {
	this := ApidefGraphQLEngineDataSource{}
	return &this
}

// GetConfig returns the Config field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApidefGraphQLEngineDataSource) GetConfig() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Config
}

// GetConfigOk returns a tuple with the Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApidefGraphQLEngineDataSource) GetConfigOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Config) {
		return nil, false
	}
	return &o.Config, true
}

// HasConfig returns a boolean if a field has been set.
func (o *ApidefGraphQLEngineDataSource) HasConfig() bool {
	if o != nil && !IsNil(o.Config) {
		return true
	}

	return false
}

// SetConfig gets a reference to the given interface{} and assigns it to the Config field.
func (o *ApidefGraphQLEngineDataSource) SetConfig(v interface{}) {
	o.Config = v
}

// GetInternal returns the Internal field value if set, zero value otherwise.
func (o *ApidefGraphQLEngineDataSource) GetInternal() bool {
	if o == nil || IsNil(o.Internal) {
		var ret bool
		return ret
	}
	return *o.Internal
}

// GetInternalOk returns a tuple with the Internal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApidefGraphQLEngineDataSource) GetInternalOk() (*bool, bool) {
	if o == nil || IsNil(o.Internal) {
		return nil, false
	}
	return o.Internal, true
}

// HasInternal returns a boolean if a field has been set.
func (o *ApidefGraphQLEngineDataSource) HasInternal() bool {
	if o != nil && !IsNil(o.Internal) {
		return true
	}

	return false
}

// SetInternal gets a reference to the given bool and assigns it to the Internal field.
func (o *ApidefGraphQLEngineDataSource) SetInternal(v bool) {
	o.Internal = &v
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *ApidefGraphQLEngineDataSource) GetKind() string {
	if o == nil || IsNil(o.Kind) {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApidefGraphQLEngineDataSource) GetKindOk() (*string, bool) {
	if o == nil || IsNil(o.Kind) {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *ApidefGraphQLEngineDataSource) HasKind() bool {
	if o != nil && !IsNil(o.Kind) {
		return true
	}

	return false
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *ApidefGraphQLEngineDataSource) SetKind(v string) {
	o.Kind = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ApidefGraphQLEngineDataSource) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApidefGraphQLEngineDataSource) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ApidefGraphQLEngineDataSource) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ApidefGraphQLEngineDataSource) SetName(v string) {
	o.Name = &v
}

// GetRootFields returns the RootFields field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApidefGraphQLEngineDataSource) GetRootFields() []ApidefGraphQLTypeFields {
	if o == nil {
		var ret []ApidefGraphQLTypeFields
		return ret
	}
	return o.RootFields
}

// GetRootFieldsOk returns a tuple with the RootFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApidefGraphQLEngineDataSource) GetRootFieldsOk() ([]ApidefGraphQLTypeFields, bool) {
	if o == nil || IsNil(o.RootFields) {
		return nil, false
	}
	return o.RootFields, true
}

// HasRootFields returns a boolean if a field has been set.
func (o *ApidefGraphQLEngineDataSource) HasRootFields() bool {
	if o != nil && !IsNil(o.RootFields) {
		return true
	}

	return false
}

// SetRootFields gets a reference to the given []ApidefGraphQLTypeFields and assigns it to the RootFields field.
func (o *ApidefGraphQLEngineDataSource) SetRootFields(v []ApidefGraphQLTypeFields) {
	o.RootFields = v
}

func (o ApidefGraphQLEngineDataSource) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApidefGraphQLEngineDataSource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Config != nil {
		toSerialize["config"] = o.Config
	}
	if !IsNil(o.Internal) {
		toSerialize["internal"] = o.Internal
	}
	if !IsNil(o.Kind) {
		toSerialize["kind"] = o.Kind
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if o.RootFields != nil {
		toSerialize["root_fields"] = o.RootFields
	}
	return toSerialize, nil
}

type NullableApidefGraphQLEngineDataSource struct {
	value *ApidefGraphQLEngineDataSource
	isSet bool
}

func (v NullableApidefGraphQLEngineDataSource) Get() *ApidefGraphQLEngineDataSource {
	return v.value
}

func (v *NullableApidefGraphQLEngineDataSource) Set(val *ApidefGraphQLEngineDataSource) {
	v.value = val
	v.isSet = true
}

func (v NullableApidefGraphQLEngineDataSource) IsSet() bool {
	return v.isSet
}

func (v *NullableApidefGraphQLEngineDataSource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApidefGraphQLEngineDataSource(val *ApidefGraphQLEngineDataSource) *NullableApidefGraphQLEngineDataSource {
	return &NullableApidefGraphQLEngineDataSource{value: val, isSet: true}
}

func (v NullableApidefGraphQLEngineDataSource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApidefGraphQLEngineDataSource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
