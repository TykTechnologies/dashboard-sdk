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

// checks if the Openapi3PathItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Openapi3PathItem{}

// Openapi3PathItem struct for Openapi3PathItem
type Openapi3PathItem struct {
	Ref         *string                  `json:"$ref,omitempty"`
	Connect     *Openapi3Operation       `json:"connect,omitempty"`
	Delete      *Openapi3Operation       `json:"delete,omitempty"`
	Description *string                  `json:"description,omitempty"`
	Get         *Openapi3Operation       `json:"get,omitempty"`
	Head        *Openapi3Operation       `json:"head,omitempty"`
	Options     *Openapi3Operation       `json:"options,omitempty"`
	Parameters  []map[string]interface{} `json:"parameters,omitempty"`
	Patch       *Openapi3Operation       `json:"patch,omitempty"`
	Post        *Openapi3Operation       `json:"post,omitempty"`
	Put         *Openapi3Operation       `json:"put,omitempty"`
	Servers     []Openapi3Server         `json:"servers,omitempty"`
	Summary     *string                  `json:"summary,omitempty"`
	Trace       *Openapi3Operation       `json:"trace,omitempty"`
}

// NewOpenapi3PathItem instantiates a new Openapi3PathItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOpenapi3PathItem() *Openapi3PathItem {
	this := Openapi3PathItem{}
	return &this
}

// NewOpenapi3PathItemWithDefaults instantiates a new Openapi3PathItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOpenapi3PathItemWithDefaults() *Openapi3PathItem {
	this := Openapi3PathItem{}
	return &this
}

// GetRef returns the Ref field value if set, zero value otherwise.
func (o *Openapi3PathItem) GetRef() string {
	if o == nil || IsNil(o.Ref) {
		var ret string
		return ret
	}
	return *o.Ref
}

// GetRefOk returns a tuple with the Ref field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Openapi3PathItem) GetRefOk() (*string, bool) {
	if o == nil || IsNil(o.Ref) {
		return nil, false
	}
	return o.Ref, true
}

// HasRef returns a boolean if a field has been set.
func (o *Openapi3PathItem) HasRef() bool {
	if o != nil && !IsNil(o.Ref) {
		return true
	}

	return false
}

// SetRef gets a reference to the given string and assigns it to the Ref field.
func (o *Openapi3PathItem) SetRef(v string) {
	o.Ref = &v
}

// GetConnect returns the Connect field value if set, zero value otherwise.
func (o *Openapi3PathItem) GetConnect() Openapi3Operation {
	if o == nil || IsNil(o.Connect) {
		var ret Openapi3Operation
		return ret
	}
	return *o.Connect
}

// GetConnectOk returns a tuple with the Connect field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Openapi3PathItem) GetConnectOk() (*Openapi3Operation, bool) {
	if o == nil || IsNil(o.Connect) {
		return nil, false
	}
	return o.Connect, true
}

// HasConnect returns a boolean if a field has been set.
func (o *Openapi3PathItem) HasConnect() bool {
	if o != nil && !IsNil(o.Connect) {
		return true
	}

	return false
}

// SetConnect gets a reference to the given Openapi3Operation and assigns it to the Connect field.
func (o *Openapi3PathItem) SetConnect(v Openapi3Operation) {
	o.Connect = &v
}

// GetDelete returns the Delete field value if set, zero value otherwise.
func (o *Openapi3PathItem) GetDelete() Openapi3Operation {
	if o == nil || IsNil(o.Delete) {
		var ret Openapi3Operation
		return ret
	}
	return *o.Delete
}

// GetDeleteOk returns a tuple with the Delete field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Openapi3PathItem) GetDeleteOk() (*Openapi3Operation, bool) {
	if o == nil || IsNil(o.Delete) {
		return nil, false
	}
	return o.Delete, true
}

// HasDelete returns a boolean if a field has been set.
func (o *Openapi3PathItem) HasDelete() bool {
	if o != nil && !IsNil(o.Delete) {
		return true
	}

	return false
}

// SetDelete gets a reference to the given Openapi3Operation and assigns it to the Delete field.
func (o *Openapi3PathItem) SetDelete(v Openapi3Operation) {
	o.Delete = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Openapi3PathItem) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Openapi3PathItem) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Openapi3PathItem) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Openapi3PathItem) SetDescription(v string) {
	o.Description = &v
}

// GetGet returns the Get field value if set, zero value otherwise.
func (o *Openapi3PathItem) GetGet() Openapi3Operation {
	if o == nil || IsNil(o.Get) {
		var ret Openapi3Operation
		return ret
	}
	return *o.Get
}

// GetGetOk returns a tuple with the Get field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Openapi3PathItem) GetGetOk() (*Openapi3Operation, bool) {
	if o == nil || IsNil(o.Get) {
		return nil, false
	}
	return o.Get, true
}

// HasGet returns a boolean if a field has been set.
func (o *Openapi3PathItem) HasGet() bool {
	if o != nil && !IsNil(o.Get) {
		return true
	}

	return false
}

// SetGet gets a reference to the given Openapi3Operation and assigns it to the Get field.
func (o *Openapi3PathItem) SetGet(v Openapi3Operation) {
	o.Get = &v
}

// GetHead returns the Head field value if set, zero value otherwise.
func (o *Openapi3PathItem) GetHead() Openapi3Operation {
	if o == nil || IsNil(o.Head) {
		var ret Openapi3Operation
		return ret
	}
	return *o.Head
}

// GetHeadOk returns a tuple with the Head field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Openapi3PathItem) GetHeadOk() (*Openapi3Operation, bool) {
	if o == nil || IsNil(o.Head) {
		return nil, false
	}
	return o.Head, true
}

// HasHead returns a boolean if a field has been set.
func (o *Openapi3PathItem) HasHead() bool {
	if o != nil && !IsNil(o.Head) {
		return true
	}

	return false
}

// SetHead gets a reference to the given Openapi3Operation and assigns it to the Head field.
func (o *Openapi3PathItem) SetHead(v Openapi3Operation) {
	o.Head = &v
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *Openapi3PathItem) GetOptions() Openapi3Operation {
	if o == nil || IsNil(o.Options) {
		var ret Openapi3Operation
		return ret
	}
	return *o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Openapi3PathItem) GetOptionsOk() (*Openapi3Operation, bool) {
	if o == nil || IsNil(o.Options) {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *Openapi3PathItem) HasOptions() bool {
	if o != nil && !IsNil(o.Options) {
		return true
	}

	return false
}

// SetOptions gets a reference to the given Openapi3Operation and assigns it to the Options field.
func (o *Openapi3PathItem) SetOptions(v Openapi3Operation) {
	o.Options = &v
}

// GetParameters returns the Parameters field value if set, zero value otherwise.
func (o *Openapi3PathItem) GetParameters() []map[string]interface{} {
	if o == nil || IsNil(o.Parameters) {
		var ret []map[string]interface{}
		return ret
	}
	return o.Parameters
}

// GetParametersOk returns a tuple with the Parameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Openapi3PathItem) GetParametersOk() ([]map[string]interface{}, bool) {
	if o == nil || IsNil(o.Parameters) {
		return nil, false
	}
	return o.Parameters, true
}

// HasParameters returns a boolean if a field has been set.
func (o *Openapi3PathItem) HasParameters() bool {
	if o != nil && !IsNil(o.Parameters) {
		return true
	}

	return false
}

// SetParameters gets a reference to the given []map[string]interface{} and assigns it to the Parameters field.
func (o *Openapi3PathItem) SetParameters(v []map[string]interface{}) {
	o.Parameters = v
}

// GetPatch returns the Patch field value if set, zero value otherwise.
func (o *Openapi3PathItem) GetPatch() Openapi3Operation {
	if o == nil || IsNil(o.Patch) {
		var ret Openapi3Operation
		return ret
	}
	return *o.Patch
}

// GetPatchOk returns a tuple with the Patch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Openapi3PathItem) GetPatchOk() (*Openapi3Operation, bool) {
	if o == nil || IsNil(o.Patch) {
		return nil, false
	}
	return o.Patch, true
}

// HasPatch returns a boolean if a field has been set.
func (o *Openapi3PathItem) HasPatch() bool {
	if o != nil && !IsNil(o.Patch) {
		return true
	}

	return false
}

// SetPatch gets a reference to the given Openapi3Operation and assigns it to the Patch field.
func (o *Openapi3PathItem) SetPatch(v Openapi3Operation) {
	o.Patch = &v
}

// GetPost returns the Post field value if set, zero value otherwise.
func (o *Openapi3PathItem) GetPost() Openapi3Operation {
	if o == nil || IsNil(o.Post) {
		var ret Openapi3Operation
		return ret
	}
	return *o.Post
}

// GetPostOk returns a tuple with the Post field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Openapi3PathItem) GetPostOk() (*Openapi3Operation, bool) {
	if o == nil || IsNil(o.Post) {
		return nil, false
	}
	return o.Post, true
}

// HasPost returns a boolean if a field has been set.
func (o *Openapi3PathItem) HasPost() bool {
	if o != nil && !IsNil(o.Post) {
		return true
	}

	return false
}

// SetPost gets a reference to the given Openapi3Operation and assigns it to the Post field.
func (o *Openapi3PathItem) SetPost(v Openapi3Operation) {
	o.Post = &v
}

// GetPut returns the Put field value if set, zero value otherwise.
func (o *Openapi3PathItem) GetPut() Openapi3Operation {
	if o == nil || IsNil(o.Put) {
		var ret Openapi3Operation
		return ret
	}
	return *o.Put
}

// GetPutOk returns a tuple with the Put field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Openapi3PathItem) GetPutOk() (*Openapi3Operation, bool) {
	if o == nil || IsNil(o.Put) {
		return nil, false
	}
	return o.Put, true
}

// HasPut returns a boolean if a field has been set.
func (o *Openapi3PathItem) HasPut() bool {
	if o != nil && !IsNil(o.Put) {
		return true
	}

	return false
}

// SetPut gets a reference to the given Openapi3Operation and assigns it to the Put field.
func (o *Openapi3PathItem) SetPut(v Openapi3Operation) {
	o.Put = &v
}

// GetServers returns the Servers field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Openapi3PathItem) GetServers() []Openapi3Server {
	if o == nil {
		var ret []Openapi3Server
		return ret
	}
	return o.Servers
}

// GetServersOk returns a tuple with the Servers field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Openapi3PathItem) GetServersOk() ([]Openapi3Server, bool) {
	if o == nil || IsNil(o.Servers) {
		return nil, false
	}
	return o.Servers, true
}

// HasServers returns a boolean if a field has been set.
func (o *Openapi3PathItem) HasServers() bool {
	if o != nil && !IsNil(o.Servers) {
		return true
	}

	return false
}

// SetServers gets a reference to the given []Openapi3Server and assigns it to the Servers field.
func (o *Openapi3PathItem) SetServers(v []Openapi3Server) {
	o.Servers = v
}

// GetSummary returns the Summary field value if set, zero value otherwise.
func (o *Openapi3PathItem) GetSummary() string {
	if o == nil || IsNil(o.Summary) {
		var ret string
		return ret
	}
	return *o.Summary
}

// GetSummaryOk returns a tuple with the Summary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Openapi3PathItem) GetSummaryOk() (*string, bool) {
	if o == nil || IsNil(o.Summary) {
		return nil, false
	}
	return o.Summary, true
}

// HasSummary returns a boolean if a field has been set.
func (o *Openapi3PathItem) HasSummary() bool {
	if o != nil && !IsNil(o.Summary) {
		return true
	}

	return false
}

// SetSummary gets a reference to the given string and assigns it to the Summary field.
func (o *Openapi3PathItem) SetSummary(v string) {
	o.Summary = &v
}

// GetTrace returns the Trace field value if set, zero value otherwise.
func (o *Openapi3PathItem) GetTrace() Openapi3Operation {
	if o == nil || IsNil(o.Trace) {
		var ret Openapi3Operation
		return ret
	}
	return *o.Trace
}

// GetTraceOk returns a tuple with the Trace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Openapi3PathItem) GetTraceOk() (*Openapi3Operation, bool) {
	if o == nil || IsNil(o.Trace) {
		return nil, false
	}
	return o.Trace, true
}

// HasTrace returns a boolean if a field has been set.
func (o *Openapi3PathItem) HasTrace() bool {
	if o != nil && !IsNil(o.Trace) {
		return true
	}

	return false
}

// SetTrace gets a reference to the given Openapi3Operation and assigns it to the Trace field.
func (o *Openapi3PathItem) SetTrace(v Openapi3Operation) {
	o.Trace = &v
}

func (o Openapi3PathItem) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Openapi3PathItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Ref) {
		toSerialize["$ref"] = o.Ref
	}
	if !IsNil(o.Connect) {
		toSerialize["connect"] = o.Connect
	}
	if !IsNil(o.Delete) {
		toSerialize["delete"] = o.Delete
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Get) {
		toSerialize["get"] = o.Get
	}
	if !IsNil(o.Head) {
		toSerialize["head"] = o.Head
	}
	if !IsNil(o.Options) {
		toSerialize["options"] = o.Options
	}
	if !IsNil(o.Parameters) {
		toSerialize["parameters"] = o.Parameters
	}
	if !IsNil(o.Patch) {
		toSerialize["patch"] = o.Patch
	}
	if !IsNil(o.Post) {
		toSerialize["post"] = o.Post
	}
	if !IsNil(o.Put) {
		toSerialize["put"] = o.Put
	}
	if o.Servers != nil {
		toSerialize["servers"] = o.Servers
	}
	if !IsNil(o.Summary) {
		toSerialize["summary"] = o.Summary
	}
	if !IsNil(o.Trace) {
		toSerialize["trace"] = o.Trace
	}
	return toSerialize, nil
}

type NullableOpenapi3PathItem struct {
	value *Openapi3PathItem
	isSet bool
}

func (v NullableOpenapi3PathItem) Get() *Openapi3PathItem {
	return v.value
}

func (v *NullableOpenapi3PathItem) Set(val *Openapi3PathItem) {
	v.value = val
	v.isSet = true
}

func (v NullableOpenapi3PathItem) IsSet() bool {
	return v.isSet
}

func (v *NullableOpenapi3PathItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOpenapi3PathItem(val *Openapi3PathItem) *NullableOpenapi3PathItem {
	return &NullableOpenapi3PathItem{value: val, isSet: true}
}

func (v NullableOpenapi3PathItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOpenapi3PathItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
