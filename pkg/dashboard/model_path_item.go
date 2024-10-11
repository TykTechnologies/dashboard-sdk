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

// checks if the PathItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PathItem{}

// PathItem struct for PathItem
type PathItem struct {
	Ref         *string                  `json:"$ref,omitempty"`
	Connect     NullableOperationType2   `json:"connect,omitempty"`
	Delete      NullableOperationType2   `json:"delete,omitempty"`
	Description *string                  `json:"description,omitempty"`
	Get         NullableOperationType2   `json:"get,omitempty"`
	Head        NullableOperationType2   `json:"head,omitempty"`
	Options     NullableOperationType2   `json:"options,omitempty"`
	Parameters  []map[string]interface{} `json:"parameters,omitempty"`
	Patch       NullableOperationType2   `json:"patch,omitempty"`
	Post        NullableOperationType2   `json:"post,omitempty"`
	Put         NullableOperationType2   `json:"put,omitempty"`
	Servers     []ServerType2            `json:"servers,omitempty"`
	Summary     *string                  `json:"summary,omitempty"`
	Trace       NullableOperationType2   `json:"trace,omitempty"`
}

// NewPathItem instantiates a new PathItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPathItem() *PathItem {
	this := PathItem{}
	return &this
}

// NewPathItemWithDefaults instantiates a new PathItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPathItemWithDefaults() *PathItem {
	this := PathItem{}
	return &this
}

// GetRef returns the Ref field value if set, zero value otherwise.
func (o *PathItem) GetRef() string {
	if o == nil || IsNil(o.Ref) {
		var ret string
		return ret
	}
	return *o.Ref
}

// GetRefOk returns a tuple with the Ref field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PathItem) GetRefOk() (*string, bool) {
	if o == nil || IsNil(o.Ref) {
		return nil, false
	}
	return o.Ref, true
}

// HasRef returns a boolean if a field has been set.
func (o *PathItem) HasRef() bool {
	if o != nil && !IsNil(o.Ref) {
		return true
	}

	return false
}

// SetRef gets a reference to the given string and assigns it to the Ref field.
func (o *PathItem) SetRef(v string) {
	o.Ref = &v
}

// GetConnect returns the Connect field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PathItem) GetConnect() OperationType2 {
	if o == nil || IsNil(o.Connect.Get()) {
		var ret OperationType2
		return ret
	}
	return *o.Connect.Get()
}

// GetConnectOk returns a tuple with the Connect field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PathItem) GetConnectOk() (*OperationType2, bool) {
	if o == nil {
		return nil, false
	}
	return o.Connect.Get(), o.Connect.IsSet()
}

// HasConnect returns a boolean if a field has been set.
func (o *PathItem) HasConnect() bool {
	if o != nil && o.Connect.IsSet() {
		return true
	}

	return false
}

// SetConnect gets a reference to the given NullableOperationType2 and assigns it to the Connect field.
func (o *PathItem) SetConnect(v OperationType2) {
	o.Connect.Set(&v)
}

// SetConnectNil sets the value for Connect to be an explicit nil
func (o *PathItem) SetConnectNil() {
	o.Connect.Set(nil)
}

// UnsetConnect ensures that no value is present for Connect, not even an explicit nil
func (o *PathItem) UnsetConnect() {
	o.Connect.Unset()
}

// GetDelete returns the Delete field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PathItem) GetDelete() OperationType2 {
	if o == nil || IsNil(o.Delete.Get()) {
		var ret OperationType2
		return ret
	}
	return *o.Delete.Get()
}

// GetDeleteOk returns a tuple with the Delete field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PathItem) GetDeleteOk() (*OperationType2, bool) {
	if o == nil {
		return nil, false
	}
	return o.Delete.Get(), o.Delete.IsSet()
}

// HasDelete returns a boolean if a field has been set.
func (o *PathItem) HasDelete() bool {
	if o != nil && o.Delete.IsSet() {
		return true
	}

	return false
}

// SetDelete gets a reference to the given NullableOperationType2 and assigns it to the Delete field.
func (o *PathItem) SetDelete(v OperationType2) {
	o.Delete.Set(&v)
}

// SetDeleteNil sets the value for Delete to be an explicit nil
func (o *PathItem) SetDeleteNil() {
	o.Delete.Set(nil)
}

// UnsetDelete ensures that no value is present for Delete, not even an explicit nil
func (o *PathItem) UnsetDelete() {
	o.Delete.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PathItem) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PathItem) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PathItem) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PathItem) SetDescription(v string) {
	o.Description = &v
}

// GetGet returns the Get field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PathItem) GetGet() OperationType2 {
	if o == nil || IsNil(o.Get.Get()) {
		var ret OperationType2
		return ret
	}
	return *o.Get.Get()
}

// GetGetOk returns a tuple with the Get field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PathItem) GetGetOk() (*OperationType2, bool) {
	if o == nil {
		return nil, false
	}
	return o.Get.Get(), o.Get.IsSet()
}

// HasGet returns a boolean if a field has been set.
func (o *PathItem) HasGet() bool {
	if o != nil && o.Get.IsSet() {
		return true
	}

	return false
}

// SetGet gets a reference to the given NullableOperationType2 and assigns it to the Get field.
func (o *PathItem) SetGet(v OperationType2) {
	o.Get.Set(&v)
}

// SetGetNil sets the value for Get to be an explicit nil
func (o *PathItem) SetGetNil() {
	o.Get.Set(nil)
}

// UnsetGet ensures that no value is present for Get, not even an explicit nil
func (o *PathItem) UnsetGet() {
	o.Get.Unset()
}

// GetHead returns the Head field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PathItem) GetHead() OperationType2 {
	if o == nil || IsNil(o.Head.Get()) {
		var ret OperationType2
		return ret
	}
	return *o.Head.Get()
}

// GetHeadOk returns a tuple with the Head field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PathItem) GetHeadOk() (*OperationType2, bool) {
	if o == nil {
		return nil, false
	}
	return o.Head.Get(), o.Head.IsSet()
}

// HasHead returns a boolean if a field has been set.
func (o *PathItem) HasHead() bool {
	if o != nil && o.Head.IsSet() {
		return true
	}

	return false
}

// SetHead gets a reference to the given NullableOperationType2 and assigns it to the Head field.
func (o *PathItem) SetHead(v OperationType2) {
	o.Head.Set(&v)
}

// SetHeadNil sets the value for Head to be an explicit nil
func (o *PathItem) SetHeadNil() {
	o.Head.Set(nil)
}

// UnsetHead ensures that no value is present for Head, not even an explicit nil
func (o *PathItem) UnsetHead() {
	o.Head.Unset()
}

// GetOptions returns the Options field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PathItem) GetOptions() OperationType2 {
	if o == nil || IsNil(o.Options.Get()) {
		var ret OperationType2
		return ret
	}
	return *o.Options.Get()
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PathItem) GetOptionsOk() (*OperationType2, bool) {
	if o == nil {
		return nil, false
	}
	return o.Options.Get(), o.Options.IsSet()
}

// HasOptions returns a boolean if a field has been set.
func (o *PathItem) HasOptions() bool {
	if o != nil && o.Options.IsSet() {
		return true
	}

	return false
}

// SetOptions gets a reference to the given NullableOperationType2 and assigns it to the Options field.
func (o *PathItem) SetOptions(v OperationType2) {
	o.Options.Set(&v)
}

// SetOptionsNil sets the value for Options to be an explicit nil
func (o *PathItem) SetOptionsNil() {
	o.Options.Set(nil)
}

// UnsetOptions ensures that no value is present for Options, not even an explicit nil
func (o *PathItem) UnsetOptions() {
	o.Options.Unset()
}

// GetParameters returns the Parameters field value if set, zero value otherwise.
func (o *PathItem) GetParameters() []map[string]interface{} {
	if o == nil || IsNil(o.Parameters) {
		var ret []map[string]interface{}
		return ret
	}
	return o.Parameters
}

// GetParametersOk returns a tuple with the Parameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PathItem) GetParametersOk() ([]map[string]interface{}, bool) {
	if o == nil || IsNil(o.Parameters) {
		return nil, false
	}
	return o.Parameters, true
}

// HasParameters returns a boolean if a field has been set.
func (o *PathItem) HasParameters() bool {
	if o != nil && !IsNil(o.Parameters) {
		return true
	}

	return false
}

// SetParameters gets a reference to the given []map[string]interface{} and assigns it to the Parameters field.
func (o *PathItem) SetParameters(v []map[string]interface{}) {
	o.Parameters = v
}

// GetPatch returns the Patch field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PathItem) GetPatch() OperationType2 {
	if o == nil || IsNil(o.Patch.Get()) {
		var ret OperationType2
		return ret
	}
	return *o.Patch.Get()
}

// GetPatchOk returns a tuple with the Patch field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PathItem) GetPatchOk() (*OperationType2, bool) {
	if o == nil {
		return nil, false
	}
	return o.Patch.Get(), o.Patch.IsSet()
}

// HasPatch returns a boolean if a field has been set.
func (o *PathItem) HasPatch() bool {
	if o != nil && o.Patch.IsSet() {
		return true
	}

	return false
}

// SetPatch gets a reference to the given NullableOperationType2 and assigns it to the Patch field.
func (o *PathItem) SetPatch(v OperationType2) {
	o.Patch.Set(&v)
}

// SetPatchNil sets the value for Patch to be an explicit nil
func (o *PathItem) SetPatchNil() {
	o.Patch.Set(nil)
}

// UnsetPatch ensures that no value is present for Patch, not even an explicit nil
func (o *PathItem) UnsetPatch() {
	o.Patch.Unset()
}

// GetPost returns the Post field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PathItem) GetPost() OperationType2 {
	if o == nil || IsNil(o.Post.Get()) {
		var ret OperationType2
		return ret
	}
	return *o.Post.Get()
}

// GetPostOk returns a tuple with the Post field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PathItem) GetPostOk() (*OperationType2, bool) {
	if o == nil {
		return nil, false
	}
	return o.Post.Get(), o.Post.IsSet()
}

// HasPost returns a boolean if a field has been set.
func (o *PathItem) HasPost() bool {
	if o != nil && o.Post.IsSet() {
		return true
	}

	return false
}

// SetPost gets a reference to the given NullableOperationType2 and assigns it to the Post field.
func (o *PathItem) SetPost(v OperationType2) {
	o.Post.Set(&v)
}

// SetPostNil sets the value for Post to be an explicit nil
func (o *PathItem) SetPostNil() {
	o.Post.Set(nil)
}

// UnsetPost ensures that no value is present for Post, not even an explicit nil
func (o *PathItem) UnsetPost() {
	o.Post.Unset()
}

// GetPut returns the Put field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PathItem) GetPut() OperationType2 {
	if o == nil || IsNil(o.Put.Get()) {
		var ret OperationType2
		return ret
	}
	return *o.Put.Get()
}

// GetPutOk returns a tuple with the Put field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PathItem) GetPutOk() (*OperationType2, bool) {
	if o == nil {
		return nil, false
	}
	return o.Put.Get(), o.Put.IsSet()
}

// HasPut returns a boolean if a field has been set.
func (o *PathItem) HasPut() bool {
	if o != nil && o.Put.IsSet() {
		return true
	}

	return false
}

// SetPut gets a reference to the given NullableOperationType2 and assigns it to the Put field.
func (o *PathItem) SetPut(v OperationType2) {
	o.Put.Set(&v)
}

// SetPutNil sets the value for Put to be an explicit nil
func (o *PathItem) SetPutNil() {
	o.Put.Set(nil)
}

// UnsetPut ensures that no value is present for Put, not even an explicit nil
func (o *PathItem) UnsetPut() {
	o.Put.Unset()
}

// GetServers returns the Servers field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PathItem) GetServers() []ServerType2 {
	if o == nil {
		var ret []ServerType2
		return ret
	}
	return o.Servers
}

// GetServersOk returns a tuple with the Servers field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PathItem) GetServersOk() ([]ServerType2, bool) {
	if o == nil || IsNil(o.Servers) {
		return nil, false
	}
	return o.Servers, true
}

// HasServers returns a boolean if a field has been set.
func (o *PathItem) HasServers() bool {
	if o != nil && !IsNil(o.Servers) {
		return true
	}

	return false
}

// SetServers gets a reference to the given []ServerType2 and assigns it to the Servers field.
func (o *PathItem) SetServers(v []ServerType2) {
	o.Servers = v
}

// GetSummary returns the Summary field value if set, zero value otherwise.
func (o *PathItem) GetSummary() string {
	if o == nil || IsNil(o.Summary) {
		var ret string
		return ret
	}
	return *o.Summary
}

// GetSummaryOk returns a tuple with the Summary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PathItem) GetSummaryOk() (*string, bool) {
	if o == nil || IsNil(o.Summary) {
		return nil, false
	}
	return o.Summary, true
}

// HasSummary returns a boolean if a field has been set.
func (o *PathItem) HasSummary() bool {
	if o != nil && !IsNil(o.Summary) {
		return true
	}

	return false
}

// SetSummary gets a reference to the given string and assigns it to the Summary field.
func (o *PathItem) SetSummary(v string) {
	o.Summary = &v
}

// GetTrace returns the Trace field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PathItem) GetTrace() OperationType2 {
	if o == nil || IsNil(o.Trace.Get()) {
		var ret OperationType2
		return ret
	}
	return *o.Trace.Get()
}

// GetTraceOk returns a tuple with the Trace field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PathItem) GetTraceOk() (*OperationType2, bool) {
	if o == nil {
		return nil, false
	}
	return o.Trace.Get(), o.Trace.IsSet()
}

// HasTrace returns a boolean if a field has been set.
func (o *PathItem) HasTrace() bool {
	if o != nil && o.Trace.IsSet() {
		return true
	}

	return false
}

// SetTrace gets a reference to the given NullableOperationType2 and assigns it to the Trace field.
func (o *PathItem) SetTrace(v OperationType2) {
	o.Trace.Set(&v)
}

// SetTraceNil sets the value for Trace to be an explicit nil
func (o *PathItem) SetTraceNil() {
	o.Trace.Set(nil)
}

// UnsetTrace ensures that no value is present for Trace, not even an explicit nil
func (o *PathItem) UnsetTrace() {
	o.Trace.Unset()
}

func (o PathItem) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PathItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Ref) {
		toSerialize["$ref"] = o.Ref
	}
	if o.Connect.IsSet() {
		toSerialize["connect"] = o.Connect.Get()
	}
	if o.Delete.IsSet() {
		toSerialize["delete"] = o.Delete.Get()
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if o.Get.IsSet() {
		toSerialize["get"] = o.Get.Get()
	}
	if o.Head.IsSet() {
		toSerialize["head"] = o.Head.Get()
	}
	if o.Options.IsSet() {
		toSerialize["options"] = o.Options.Get()
	}
	if !IsNil(o.Parameters) {
		toSerialize["parameters"] = o.Parameters
	}
	if o.Patch.IsSet() {
		toSerialize["patch"] = o.Patch.Get()
	}
	if o.Post.IsSet() {
		toSerialize["post"] = o.Post.Get()
	}
	if o.Put.IsSet() {
		toSerialize["put"] = o.Put.Get()
	}
	if o.Servers != nil {
		toSerialize["servers"] = o.Servers
	}
	if !IsNil(o.Summary) {
		toSerialize["summary"] = o.Summary
	}
	if o.Trace.IsSet() {
		toSerialize["trace"] = o.Trace.Get()
	}
	return toSerialize, nil
}

type NullablePathItem struct {
	value *PathItem
	isSet bool
}

func (v NullablePathItem) Get() *PathItem {
	return v.value
}

func (v *NullablePathItem) Set(val *PathItem) {
	v.value = val
	v.isSet = true
}

func (v NullablePathItem) IsSet() bool {
	return v.isSet
}

func (v *NullablePathItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePathItem(val *PathItem) *NullablePathItem {
	return &NullablePathItem{value: val, isSet: true}
}

func (v NullablePathItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePathItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
