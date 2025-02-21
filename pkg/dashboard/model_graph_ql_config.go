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
	"time"
)

// checks if the GraphQLConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GraphQLConfig{}

// GraphQLConfig struct for GraphQLConfig
type GraphQLConfig struct {
	Enabled                 *bool                              `json:"enabled,omitempty"`
	Engine                  *GraphQLEngineConfig               `json:"engine,omitempty"`
	ExecutionMode           *string                            `json:"execution_mode,omitempty"`
	Introspection           *GraphQLIntrospectionConfig        `json:"introspection,omitempty"`
	LastSchemaUpdate        NullableTime                       `json:"last_schema_update,omitempty"`
	Playground              *GraphQLPlayground                 `json:"playground,omitempty"`
	Proxy                   *GraphQLProxyConfig                `json:"proxy,omitempty"`
	Schema                  *string                            `json:"schema,omitempty"`
	Subgraph                *GraphQLSubgraphConfig             `json:"subgraph,omitempty"`
	Supergraph              *GraphQLSupergraphConfig           `json:"supergraph,omitempty"`
	TypeFieldConfigurations []DatasourceTypeFieldConfiguration `json:"type_field_configurations,omitempty"`
	Version                 *string                            `json:"version,omitempty"`
}

// NewGraphQLConfig instantiates a new GraphQLConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGraphQLConfig() *GraphQLConfig {
	this := GraphQLConfig{}
	return &this
}

// NewGraphQLConfigWithDefaults instantiates a new GraphQLConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGraphQLConfigWithDefaults() *GraphQLConfig {
	this := GraphQLConfig{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *GraphQLConfig) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GraphQLConfig) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *GraphQLConfig) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *GraphQLConfig) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetEngine returns the Engine field value if set, zero value otherwise.
func (o *GraphQLConfig) GetEngine() GraphQLEngineConfig {
	if o == nil || IsNil(o.Engine) {
		var ret GraphQLEngineConfig
		return ret
	}
	return *o.Engine
}

// GetEngineOk returns a tuple with the Engine field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GraphQLConfig) GetEngineOk() (*GraphQLEngineConfig, bool) {
	if o == nil || IsNil(o.Engine) {
		return nil, false
	}
	return o.Engine, true
}

// HasEngine returns a boolean if a field has been set.
func (o *GraphQLConfig) HasEngine() bool {
	if o != nil && !IsNil(o.Engine) {
		return true
	}

	return false
}

// SetEngine gets a reference to the given GraphQLEngineConfig and assigns it to the Engine field.
func (o *GraphQLConfig) SetEngine(v GraphQLEngineConfig) {
	o.Engine = &v
}

// GetExecutionMode returns the ExecutionMode field value if set, zero value otherwise.
func (o *GraphQLConfig) GetExecutionMode() string {
	if o == nil || IsNil(o.ExecutionMode) {
		var ret string
		return ret
	}
	return *o.ExecutionMode
}

// GetExecutionModeOk returns a tuple with the ExecutionMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GraphQLConfig) GetExecutionModeOk() (*string, bool) {
	if o == nil || IsNil(o.ExecutionMode) {
		return nil, false
	}
	return o.ExecutionMode, true
}

// HasExecutionMode returns a boolean if a field has been set.
func (o *GraphQLConfig) HasExecutionMode() bool {
	if o != nil && !IsNil(o.ExecutionMode) {
		return true
	}

	return false
}

// SetExecutionMode gets a reference to the given string and assigns it to the ExecutionMode field.
func (o *GraphQLConfig) SetExecutionMode(v string) {
	o.ExecutionMode = &v
}

// GetIntrospection returns the Introspection field value if set, zero value otherwise.
func (o *GraphQLConfig) GetIntrospection() GraphQLIntrospectionConfig {
	if o == nil || IsNil(o.Introspection) {
		var ret GraphQLIntrospectionConfig
		return ret
	}
	return *o.Introspection
}

// GetIntrospectionOk returns a tuple with the Introspection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GraphQLConfig) GetIntrospectionOk() (*GraphQLIntrospectionConfig, bool) {
	if o == nil || IsNil(o.Introspection) {
		return nil, false
	}
	return o.Introspection, true
}

// HasIntrospection returns a boolean if a field has been set.
func (o *GraphQLConfig) HasIntrospection() bool {
	if o != nil && !IsNil(o.Introspection) {
		return true
	}

	return false
}

// SetIntrospection gets a reference to the given GraphQLIntrospectionConfig and assigns it to the Introspection field.
func (o *GraphQLConfig) SetIntrospection(v GraphQLIntrospectionConfig) {
	o.Introspection = &v
}

// GetLastSchemaUpdate returns the LastSchemaUpdate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GraphQLConfig) GetLastSchemaUpdate() time.Time {
	if o == nil || IsNil(o.LastSchemaUpdate.Get()) {
		var ret time.Time
		return ret
	}
	return *o.LastSchemaUpdate.Get()
}

// GetLastSchemaUpdateOk returns a tuple with the LastSchemaUpdate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GraphQLConfig) GetLastSchemaUpdateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastSchemaUpdate.Get(), o.LastSchemaUpdate.IsSet()
}

// HasLastSchemaUpdate returns a boolean if a field has been set.
func (o *GraphQLConfig) HasLastSchemaUpdate() bool {
	if o != nil && o.LastSchemaUpdate.IsSet() {
		return true
	}

	return false
}

// SetLastSchemaUpdate gets a reference to the given NullableTime and assigns it to the LastSchemaUpdate field.
func (o *GraphQLConfig) SetLastSchemaUpdate(v time.Time) {
	o.LastSchemaUpdate.Set(&v)
}

// SetLastSchemaUpdateNil sets the value for LastSchemaUpdate to be an explicit nil
func (o *GraphQLConfig) SetLastSchemaUpdateNil() {
	o.LastSchemaUpdate.Set(nil)
}

// UnsetLastSchemaUpdate ensures that no value is present for LastSchemaUpdate, not even an explicit nil
func (o *GraphQLConfig) UnsetLastSchemaUpdate() {
	o.LastSchemaUpdate.Unset()
}

// GetPlayground returns the Playground field value if set, zero value otherwise.
func (o *GraphQLConfig) GetPlayground() GraphQLPlayground {
	if o == nil || IsNil(o.Playground) {
		var ret GraphQLPlayground
		return ret
	}
	return *o.Playground
}

// GetPlaygroundOk returns a tuple with the Playground field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GraphQLConfig) GetPlaygroundOk() (*GraphQLPlayground, bool) {
	if o == nil || IsNil(o.Playground) {
		return nil, false
	}
	return o.Playground, true
}

// HasPlayground returns a boolean if a field has been set.
func (o *GraphQLConfig) HasPlayground() bool {
	if o != nil && !IsNil(o.Playground) {
		return true
	}

	return false
}

// SetPlayground gets a reference to the given GraphQLPlayground and assigns it to the Playground field.
func (o *GraphQLConfig) SetPlayground(v GraphQLPlayground) {
	o.Playground = &v
}

// GetProxy returns the Proxy field value if set, zero value otherwise.
func (o *GraphQLConfig) GetProxy() GraphQLProxyConfig {
	if o == nil || IsNil(o.Proxy) {
		var ret GraphQLProxyConfig
		return ret
	}
	return *o.Proxy
}

// GetProxyOk returns a tuple with the Proxy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GraphQLConfig) GetProxyOk() (*GraphQLProxyConfig, bool) {
	if o == nil || IsNil(o.Proxy) {
		return nil, false
	}
	return o.Proxy, true
}

// HasProxy returns a boolean if a field has been set.
func (o *GraphQLConfig) HasProxy() bool {
	if o != nil && !IsNil(o.Proxy) {
		return true
	}

	return false
}

// SetProxy gets a reference to the given GraphQLProxyConfig and assigns it to the Proxy field.
func (o *GraphQLConfig) SetProxy(v GraphQLProxyConfig) {
	o.Proxy = &v
}

// GetSchema returns the Schema field value if set, zero value otherwise.
func (o *GraphQLConfig) GetSchema() string {
	if o == nil || IsNil(o.Schema) {
		var ret string
		return ret
	}
	return *o.Schema
}

// GetSchemaOk returns a tuple with the Schema field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GraphQLConfig) GetSchemaOk() (*string, bool) {
	if o == nil || IsNil(o.Schema) {
		return nil, false
	}
	return o.Schema, true
}

// HasSchema returns a boolean if a field has been set.
func (o *GraphQLConfig) HasSchema() bool {
	if o != nil && !IsNil(o.Schema) {
		return true
	}

	return false
}

// SetSchema gets a reference to the given string and assigns it to the Schema field.
func (o *GraphQLConfig) SetSchema(v string) {
	o.Schema = &v
}

// GetSubgraph returns the Subgraph field value if set, zero value otherwise.
func (o *GraphQLConfig) GetSubgraph() GraphQLSubgraphConfig {
	if o == nil || IsNil(o.Subgraph) {
		var ret GraphQLSubgraphConfig
		return ret
	}
	return *o.Subgraph
}

// GetSubgraphOk returns a tuple with the Subgraph field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GraphQLConfig) GetSubgraphOk() (*GraphQLSubgraphConfig, bool) {
	if o == nil || IsNil(o.Subgraph) {
		return nil, false
	}
	return o.Subgraph, true
}

// HasSubgraph returns a boolean if a field has been set.
func (o *GraphQLConfig) HasSubgraph() bool {
	if o != nil && !IsNil(o.Subgraph) {
		return true
	}

	return false
}

// SetSubgraph gets a reference to the given GraphQLSubgraphConfig and assigns it to the Subgraph field.
func (o *GraphQLConfig) SetSubgraph(v GraphQLSubgraphConfig) {
	o.Subgraph = &v
}

// GetSupergraph returns the Supergraph field value if set, zero value otherwise.
func (o *GraphQLConfig) GetSupergraph() GraphQLSupergraphConfig {
	if o == nil || IsNil(o.Supergraph) {
		var ret GraphQLSupergraphConfig
		return ret
	}
	return *o.Supergraph
}

// GetSupergraphOk returns a tuple with the Supergraph field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GraphQLConfig) GetSupergraphOk() (*GraphQLSupergraphConfig, bool) {
	if o == nil || IsNil(o.Supergraph) {
		return nil, false
	}
	return o.Supergraph, true
}

// HasSupergraph returns a boolean if a field has been set.
func (o *GraphQLConfig) HasSupergraph() bool {
	if o != nil && !IsNil(o.Supergraph) {
		return true
	}

	return false
}

// SetSupergraph gets a reference to the given GraphQLSupergraphConfig and assigns it to the Supergraph field.
func (o *GraphQLConfig) SetSupergraph(v GraphQLSupergraphConfig) {
	o.Supergraph = &v
}

// GetTypeFieldConfigurations returns the TypeFieldConfigurations field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GraphQLConfig) GetTypeFieldConfigurations() []DatasourceTypeFieldConfiguration {
	if o == nil {
		var ret []DatasourceTypeFieldConfiguration
		return ret
	}
	return o.TypeFieldConfigurations
}

// GetTypeFieldConfigurationsOk returns a tuple with the TypeFieldConfigurations field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GraphQLConfig) GetTypeFieldConfigurationsOk() ([]DatasourceTypeFieldConfiguration, bool) {
	if o == nil || IsNil(o.TypeFieldConfigurations) {
		return nil, false
	}
	return o.TypeFieldConfigurations, true
}

// HasTypeFieldConfigurations returns a boolean if a field has been set.
func (o *GraphQLConfig) HasTypeFieldConfigurations() bool {
	if o != nil && !IsNil(o.TypeFieldConfigurations) {
		return true
	}

	return false
}

// SetTypeFieldConfigurations gets a reference to the given []DatasourceTypeFieldConfiguration and assigns it to the TypeFieldConfigurations field.
func (o *GraphQLConfig) SetTypeFieldConfigurations(v []DatasourceTypeFieldConfiguration) {
	o.TypeFieldConfigurations = v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *GraphQLConfig) GetVersion() string {
	if o == nil || IsNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GraphQLConfig) GetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *GraphQLConfig) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *GraphQLConfig) SetVersion(v string) {
	o.Version = &v
}

func (o GraphQLConfig) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GraphQLConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.Engine) {
		toSerialize["engine"] = o.Engine
	}
	if !IsNil(o.ExecutionMode) {
		toSerialize["execution_mode"] = o.ExecutionMode
	}
	if !IsNil(o.Introspection) {
		toSerialize["introspection"] = o.Introspection
	}
	if o.LastSchemaUpdate.IsSet() {
		toSerialize["last_schema_update"] = o.LastSchemaUpdate.Get()
	}
	if !IsNil(o.Playground) {
		toSerialize["playground"] = o.Playground
	}
	if !IsNil(o.Proxy) {
		toSerialize["proxy"] = o.Proxy
	}
	if !IsNil(o.Schema) {
		toSerialize["schema"] = o.Schema
	}
	if !IsNil(o.Subgraph) {
		toSerialize["subgraph"] = o.Subgraph
	}
	if !IsNil(o.Supergraph) {
		toSerialize["supergraph"] = o.Supergraph
	}
	if o.TypeFieldConfigurations != nil {
		toSerialize["type_field_configurations"] = o.TypeFieldConfigurations
	}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	return toSerialize, nil
}

type NullableGraphQLConfig struct {
	value *GraphQLConfig
	isSet bool
}

func (v NullableGraphQLConfig) Get() *GraphQLConfig {
	return v.value
}

func (v *NullableGraphQLConfig) Set(val *GraphQLConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableGraphQLConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableGraphQLConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGraphQLConfig(val *GraphQLConfig) *NullableGraphQLConfig {
	return &NullableGraphQLConfig{value: val, isSet: true}
}

func (v NullableGraphQLConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGraphQLConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
