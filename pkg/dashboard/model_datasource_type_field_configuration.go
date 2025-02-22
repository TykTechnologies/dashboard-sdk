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

// checks if the DatasourceTypeFieldConfiguration type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DatasourceTypeFieldConfiguration{}

// DatasourceTypeFieldConfiguration struct for DatasourceTypeFieldConfiguration
type DatasourceTypeFieldConfiguration struct {
	DataSource *DatasourceSourceConfig                `json:"data_source,omitempty"`
	FieldName  *string                                `json:"field_name,omitempty"`
	Mapping    NullableDatasourceMappingConfiguration `json:"mapping,omitempty"`
	TypeName   *string                                `json:"type_name,omitempty"`
}

// NewDatasourceTypeFieldConfiguration instantiates a new DatasourceTypeFieldConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDatasourceTypeFieldConfiguration() *DatasourceTypeFieldConfiguration {
	this := DatasourceTypeFieldConfiguration{}
	return &this
}

// NewDatasourceTypeFieldConfigurationWithDefaults instantiates a new DatasourceTypeFieldConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDatasourceTypeFieldConfigurationWithDefaults() *DatasourceTypeFieldConfiguration {
	this := DatasourceTypeFieldConfiguration{}
	return &this
}

// GetDataSource returns the DataSource field value if set, zero value otherwise.
func (o *DatasourceTypeFieldConfiguration) GetDataSource() DatasourceSourceConfig {
	if o == nil || IsNil(o.DataSource) {
		var ret DatasourceSourceConfig
		return ret
	}
	return *o.DataSource
}

// GetDataSourceOk returns a tuple with the DataSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatasourceTypeFieldConfiguration) GetDataSourceOk() (*DatasourceSourceConfig, bool) {
	if o == nil || IsNil(o.DataSource) {
		return nil, false
	}
	return o.DataSource, true
}

// HasDataSource returns a boolean if a field has been set.
func (o *DatasourceTypeFieldConfiguration) HasDataSource() bool {
	if o != nil && !IsNil(o.DataSource) {
		return true
	}

	return false
}

// SetDataSource gets a reference to the given DatasourceSourceConfig and assigns it to the DataSource field.
func (o *DatasourceTypeFieldConfiguration) SetDataSource(v DatasourceSourceConfig) {
	o.DataSource = &v
}

// GetFieldName returns the FieldName field value if set, zero value otherwise.
func (o *DatasourceTypeFieldConfiguration) GetFieldName() string {
	if o == nil || IsNil(o.FieldName) {
		var ret string
		return ret
	}
	return *o.FieldName
}

// GetFieldNameOk returns a tuple with the FieldName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatasourceTypeFieldConfiguration) GetFieldNameOk() (*string, bool) {
	if o == nil || IsNil(o.FieldName) {
		return nil, false
	}
	return o.FieldName, true
}

// HasFieldName returns a boolean if a field has been set.
func (o *DatasourceTypeFieldConfiguration) HasFieldName() bool {
	if o != nil && !IsNil(o.FieldName) {
		return true
	}

	return false
}

// SetFieldName gets a reference to the given string and assigns it to the FieldName field.
func (o *DatasourceTypeFieldConfiguration) SetFieldName(v string) {
	o.FieldName = &v
}

// GetMapping returns the Mapping field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DatasourceTypeFieldConfiguration) GetMapping() DatasourceMappingConfiguration {
	if o == nil || IsNil(o.Mapping.Get()) {
		var ret DatasourceMappingConfiguration
		return ret
	}
	return *o.Mapping.Get()
}

// GetMappingOk returns a tuple with the Mapping field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DatasourceTypeFieldConfiguration) GetMappingOk() (*DatasourceMappingConfiguration, bool) {
	if o == nil {
		return nil, false
	}
	return o.Mapping.Get(), o.Mapping.IsSet()
}

// HasMapping returns a boolean if a field has been set.
func (o *DatasourceTypeFieldConfiguration) HasMapping() bool {
	if o != nil && o.Mapping.IsSet() {
		return true
	}

	return false
}

// SetMapping gets a reference to the given NullableDatasourceMappingConfiguration and assigns it to the Mapping field.
func (o *DatasourceTypeFieldConfiguration) SetMapping(v DatasourceMappingConfiguration) {
	o.Mapping.Set(&v)
}

// SetMappingNil sets the value for Mapping to be an explicit nil
func (o *DatasourceTypeFieldConfiguration) SetMappingNil() {
	o.Mapping.Set(nil)
}

// UnsetMapping ensures that no value is present for Mapping, not even an explicit nil
func (o *DatasourceTypeFieldConfiguration) UnsetMapping() {
	o.Mapping.Unset()
}

// GetTypeName returns the TypeName field value if set, zero value otherwise.
func (o *DatasourceTypeFieldConfiguration) GetTypeName() string {
	if o == nil || IsNil(o.TypeName) {
		var ret string
		return ret
	}
	return *o.TypeName
}

// GetTypeNameOk returns a tuple with the TypeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatasourceTypeFieldConfiguration) GetTypeNameOk() (*string, bool) {
	if o == nil || IsNil(o.TypeName) {
		return nil, false
	}
	return o.TypeName, true
}

// HasTypeName returns a boolean if a field has been set.
func (o *DatasourceTypeFieldConfiguration) HasTypeName() bool {
	if o != nil && !IsNil(o.TypeName) {
		return true
	}

	return false
}

// SetTypeName gets a reference to the given string and assigns it to the TypeName field.
func (o *DatasourceTypeFieldConfiguration) SetTypeName(v string) {
	o.TypeName = &v
}

func (o DatasourceTypeFieldConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DatasourceTypeFieldConfiguration) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DataSource) {
		toSerialize["data_source"] = o.DataSource
	}
	if !IsNil(o.FieldName) {
		toSerialize["field_name"] = o.FieldName
	}
	if o.Mapping.IsSet() {
		toSerialize["mapping"] = o.Mapping.Get()
	}
	if !IsNil(o.TypeName) {
		toSerialize["type_name"] = o.TypeName
	}
	return toSerialize, nil
}

type NullableDatasourceTypeFieldConfiguration struct {
	value *DatasourceTypeFieldConfiguration
	isSet bool
}

func (v NullableDatasourceTypeFieldConfiguration) Get() *DatasourceTypeFieldConfiguration {
	return v.value
}

func (v *NullableDatasourceTypeFieldConfiguration) Set(val *DatasourceTypeFieldConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullableDatasourceTypeFieldConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullableDatasourceTypeFieldConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDatasourceTypeFieldConfiguration(val *DatasourceTypeFieldConfiguration) *NullableDatasourceTypeFieldConfiguration {
	return &NullableDatasourceTypeFieldConfiguration{value: val, isSet: true}
}

func (v NullableDatasourceTypeFieldConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDatasourceTypeFieldConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
