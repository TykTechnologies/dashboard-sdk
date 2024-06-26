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

// checks if the AccessDefinition type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccessDefinition{}

// AccessDefinition struct for AccessDefinition
type AccessDefinition struct {
	AllowanceScope       *string                 `json:"allowance_scope,omitempty"`
	AllowedTypes         []GraphqlType           `json:"allowed_types,omitempty"`
	AllowedUrls          []AccessSpec            `json:"allowed_urls,omitempty"`
	ApiId                *string                 `json:"api_id,omitempty"`
	ApiName              *string                 `json:"api_name,omitempty"`
	DisableIntrospection *bool                   `json:"disable_introspection,omitempty"`
	FieldAccessRights    []FieldAccessDefinition `json:"field_access_rights,omitempty"`
	Limit                *APILimit               `json:"limit,omitempty"`
	RestrictedTypes      []GraphqlType           `json:"restricted_types,omitempty"`
	Versions             []string                `json:"versions,omitempty"`
}

// NewAccessDefinition instantiates a new AccessDefinition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccessDefinition() *AccessDefinition {
	this := AccessDefinition{}
	return &this
}

// NewAccessDefinitionWithDefaults instantiates a new AccessDefinition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccessDefinitionWithDefaults() *AccessDefinition {
	this := AccessDefinition{}
	return &this
}

// GetAllowanceScope returns the AllowanceScope field value if set, zero value otherwise.
func (o *AccessDefinition) GetAllowanceScope() string {
	if o == nil || IsNil(o.AllowanceScope) {
		var ret string
		return ret
	}
	return *o.AllowanceScope
}

// GetAllowanceScopeOk returns a tuple with the AllowanceScope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessDefinition) GetAllowanceScopeOk() (*string, bool) {
	if o == nil || IsNil(o.AllowanceScope) {
		return nil, false
	}
	return o.AllowanceScope, true
}

// HasAllowanceScope returns a boolean if a field has been set.
func (o *AccessDefinition) HasAllowanceScope() bool {
	if o != nil && !IsNil(o.AllowanceScope) {
		return true
	}

	return false
}

// SetAllowanceScope gets a reference to the given string and assigns it to the AllowanceScope field.
func (o *AccessDefinition) SetAllowanceScope(v string) {
	o.AllowanceScope = &v
}

// GetAllowedTypes returns the AllowedTypes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AccessDefinition) GetAllowedTypes() []GraphqlType {
	if o == nil {
		var ret []GraphqlType
		return ret
	}
	return o.AllowedTypes
}

// GetAllowedTypesOk returns a tuple with the AllowedTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccessDefinition) GetAllowedTypesOk() ([]GraphqlType, bool) {
	if o == nil || IsNil(o.AllowedTypes) {
		return nil, false
	}
	return o.AllowedTypes, true
}

// HasAllowedTypes returns a boolean if a field has been set.
func (o *AccessDefinition) HasAllowedTypes() bool {
	if o != nil && !IsNil(o.AllowedTypes) {
		return true
	}

	return false
}

// SetAllowedTypes gets a reference to the given []GraphqlType and assigns it to the AllowedTypes field.
func (o *AccessDefinition) SetAllowedTypes(v []GraphqlType) {
	o.AllowedTypes = v
}

// GetAllowedUrls returns the AllowedUrls field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AccessDefinition) GetAllowedUrls() []AccessSpec {
	if o == nil {
		var ret []AccessSpec
		return ret
	}
	return o.AllowedUrls
}

// GetAllowedUrlsOk returns a tuple with the AllowedUrls field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccessDefinition) GetAllowedUrlsOk() ([]AccessSpec, bool) {
	if o == nil || IsNil(o.AllowedUrls) {
		return nil, false
	}
	return o.AllowedUrls, true
}

// HasAllowedUrls returns a boolean if a field has been set.
func (o *AccessDefinition) HasAllowedUrls() bool {
	if o != nil && !IsNil(o.AllowedUrls) {
		return true
	}

	return false
}

// SetAllowedUrls gets a reference to the given []AccessSpec and assigns it to the AllowedUrls field.
func (o *AccessDefinition) SetAllowedUrls(v []AccessSpec) {
	o.AllowedUrls = v
}

// GetApiId returns the ApiId field value if set, zero value otherwise.
func (o *AccessDefinition) GetApiId() string {
	if o == nil || IsNil(o.ApiId) {
		var ret string
		return ret
	}
	return *o.ApiId
}

// GetApiIdOk returns a tuple with the ApiId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessDefinition) GetApiIdOk() (*string, bool) {
	if o == nil || IsNil(o.ApiId) {
		return nil, false
	}
	return o.ApiId, true
}

// HasApiId returns a boolean if a field has been set.
func (o *AccessDefinition) HasApiId() bool {
	if o != nil && !IsNil(o.ApiId) {
		return true
	}

	return false
}

// SetApiId gets a reference to the given string and assigns it to the ApiId field.
func (o *AccessDefinition) SetApiId(v string) {
	o.ApiId = &v
}

// GetApiName returns the ApiName field value if set, zero value otherwise.
func (o *AccessDefinition) GetApiName() string {
	if o == nil || IsNil(o.ApiName) {
		var ret string
		return ret
	}
	return *o.ApiName
}

// GetApiNameOk returns a tuple with the ApiName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessDefinition) GetApiNameOk() (*string, bool) {
	if o == nil || IsNil(o.ApiName) {
		return nil, false
	}
	return o.ApiName, true
}

// HasApiName returns a boolean if a field has been set.
func (o *AccessDefinition) HasApiName() bool {
	if o != nil && !IsNil(o.ApiName) {
		return true
	}

	return false
}

// SetApiName gets a reference to the given string and assigns it to the ApiName field.
func (o *AccessDefinition) SetApiName(v string) {
	o.ApiName = &v
}

// GetDisableIntrospection returns the DisableIntrospection field value if set, zero value otherwise.
func (o *AccessDefinition) GetDisableIntrospection() bool {
	if o == nil || IsNil(o.DisableIntrospection) {
		var ret bool
		return ret
	}
	return *o.DisableIntrospection
}

// GetDisableIntrospectionOk returns a tuple with the DisableIntrospection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessDefinition) GetDisableIntrospectionOk() (*bool, bool) {
	if o == nil || IsNil(o.DisableIntrospection) {
		return nil, false
	}
	return o.DisableIntrospection, true
}

// HasDisableIntrospection returns a boolean if a field has been set.
func (o *AccessDefinition) HasDisableIntrospection() bool {
	if o != nil && !IsNil(o.DisableIntrospection) {
		return true
	}

	return false
}

// SetDisableIntrospection gets a reference to the given bool and assigns it to the DisableIntrospection field.
func (o *AccessDefinition) SetDisableIntrospection(v bool) {
	o.DisableIntrospection = &v
}

// GetFieldAccessRights returns the FieldAccessRights field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AccessDefinition) GetFieldAccessRights() []FieldAccessDefinition {
	if o == nil {
		var ret []FieldAccessDefinition
		return ret
	}
	return o.FieldAccessRights
}

// GetFieldAccessRightsOk returns a tuple with the FieldAccessRights field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccessDefinition) GetFieldAccessRightsOk() ([]FieldAccessDefinition, bool) {
	if o == nil || IsNil(o.FieldAccessRights) {
		return nil, false
	}
	return o.FieldAccessRights, true
}

// HasFieldAccessRights returns a boolean if a field has been set.
func (o *AccessDefinition) HasFieldAccessRights() bool {
	if o != nil && !IsNil(o.FieldAccessRights) {
		return true
	}

	return false
}

// SetFieldAccessRights gets a reference to the given []FieldAccessDefinition and assigns it to the FieldAccessRights field.
func (o *AccessDefinition) SetFieldAccessRights(v []FieldAccessDefinition) {
	o.FieldAccessRights = v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *AccessDefinition) GetLimit() APILimit {
	if o == nil || IsNil(o.Limit) {
		var ret APILimit
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessDefinition) GetLimitOk() (*APILimit, bool) {
	if o == nil || IsNil(o.Limit) {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *AccessDefinition) HasLimit() bool {
	if o != nil && !IsNil(o.Limit) {
		return true
	}

	return false
}

// SetLimit gets a reference to the given APILimit and assigns it to the Limit field.
func (o *AccessDefinition) SetLimit(v APILimit) {
	o.Limit = &v
}

// GetRestrictedTypes returns the RestrictedTypes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AccessDefinition) GetRestrictedTypes() []GraphqlType {
	if o == nil {
		var ret []GraphqlType
		return ret
	}
	return o.RestrictedTypes
}

// GetRestrictedTypesOk returns a tuple with the RestrictedTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccessDefinition) GetRestrictedTypesOk() ([]GraphqlType, bool) {
	if o == nil || IsNil(o.RestrictedTypes) {
		return nil, false
	}
	return o.RestrictedTypes, true
}

// HasRestrictedTypes returns a boolean if a field has been set.
func (o *AccessDefinition) HasRestrictedTypes() bool {
	if o != nil && !IsNil(o.RestrictedTypes) {
		return true
	}

	return false
}

// SetRestrictedTypes gets a reference to the given []GraphqlType and assigns it to the RestrictedTypes field.
func (o *AccessDefinition) SetRestrictedTypes(v []GraphqlType) {
	o.RestrictedTypes = v
}

// GetVersions returns the Versions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AccessDefinition) GetVersions() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Versions
}

// GetVersionsOk returns a tuple with the Versions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccessDefinition) GetVersionsOk() ([]string, bool) {
	if o == nil || IsNil(o.Versions) {
		return nil, false
	}
	return o.Versions, true
}

// HasVersions returns a boolean if a field has been set.
func (o *AccessDefinition) HasVersions() bool {
	if o != nil && !IsNil(o.Versions) {
		return true
	}

	return false
}

// SetVersions gets a reference to the given []string and assigns it to the Versions field.
func (o *AccessDefinition) SetVersions(v []string) {
	o.Versions = v
}

func (o AccessDefinition) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccessDefinition) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AllowanceScope) {
		toSerialize["allowance_scope"] = o.AllowanceScope
	}
	if o.AllowedTypes != nil {
		toSerialize["allowed_types"] = o.AllowedTypes
	}
	if o.AllowedUrls != nil {
		toSerialize["allowed_urls"] = o.AllowedUrls
	}
	if !IsNil(o.ApiId) {
		toSerialize["api_id"] = o.ApiId
	}
	if !IsNil(o.ApiName) {
		toSerialize["api_name"] = o.ApiName
	}
	if !IsNil(o.DisableIntrospection) {
		toSerialize["disable_introspection"] = o.DisableIntrospection
	}
	if o.FieldAccessRights != nil {
		toSerialize["field_access_rights"] = o.FieldAccessRights
	}
	if !IsNil(o.Limit) {
		toSerialize["limit"] = o.Limit
	}
	if o.RestrictedTypes != nil {
		toSerialize["restricted_types"] = o.RestrictedTypes
	}
	if o.Versions != nil {
		toSerialize["versions"] = o.Versions
	}
	return toSerialize, nil
}

type NullableAccessDefinition struct {
	value *AccessDefinition
	isSet bool
}

func (v NullableAccessDefinition) Get() *AccessDefinition {
	return v.value
}

func (v *NullableAccessDefinition) Set(val *AccessDefinition) {
	v.value = val
	v.isSet = true
}

func (v NullableAccessDefinition) IsSet() bool {
	return v.isSet
}

func (v *NullableAccessDefinition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccessDefinition(val *AccessDefinition) *NullableAccessDefinition {
	return &NullableAccessDefinition{value: val, isSet: true}
}

func (v NullableAccessDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccessDefinition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
