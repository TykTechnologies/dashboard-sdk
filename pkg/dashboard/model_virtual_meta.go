/*
Tyk Dashboard API

## <a name=\"introduction\"></a> Introduction  The Tyk Dashboard API offers granular, programmatic access to a centralised database of resources that your Tyk nodes can pull from. This API has a dynamic user administrative structure which means the secret key that is used to communicate with your Tyk nodes can be kept secret and access to the wider management functions can be handled on a user-by-user and organisation-by-organisation basis.  A common question around using a database-backed configuration is how to programatically add API definitions to your Tyk nodes, the Dashboard API allows much more fine-grained, secure and multi-user access to your Tyk cluster, and should be used to manage a database-backed Tyk node.  The Tyk Dashboard API works seamlessly with the Tyk Dashboard (and the two come bundled together).  ## <a name=\"security-hierarchy\"></a> Security Hierarchy  The Dashboard API provides a more structured security layer to managing Tyk nodes.  ### Organisations, APIs and Users  With the Dashboard API and a database-backed Tyk setup, (and to an extent with file-based API setups - if diligence is used in naming an creating definitions), the following security model is applied to the management of Upstream APIs:  * **Organisations**: All APIs are *owned* by an organisation, this is designated by the `OrgID` parameter in the API Definition. * **Users**: All users created in the Dashboard belong to an organisation (unless an exception is made for super-administrative access). * **APIs**: All APIs belong to an Organisation and only Users that belong to that organisation can see the analytics for those APIs and manage their configurations. * **API Keys**: API Keys are designated by organisation, this means an API key that has full access rights will not be allowed to access the APIs of another organisation on the same system, but can have full access to all APIs within the organisation. * **Access Rights**: Access rights are stored with the key, this enables a key to give access to multiple APIs, this is defined by the session object in the core Tyk API.  In order to use the Dashboard API, you'll need to get the `Tyk Dashboard API Access Credentials` secret from your user profile on the Dashboard UI.  The secret you set should then be sent along as a header with each Dashboard API Request in order for it to be successful:  ``` authorization: <your-secret> ```

API version: 5.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dashboard

import (
	"encoding/json"
)

// checks if the VirtualMeta type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VirtualMeta{}

// VirtualMeta struct for VirtualMeta
type VirtualMeta struct {
	FunctionSourceType *string `json:"function_source_type,omitempty"`
	FunctionSourceUri *string `json:"function_source_uri,omitempty"`
	Method *string `json:"method,omitempty"`
	Path *string `json:"path,omitempty"`
	ProxyOnError *bool `json:"proxy_on_error,omitempty"`
	ResponseFunctionName *string `json:"response_function_name,omitempty"`
	UseSession *bool `json:"use_session,omitempty"`
}

// NewVirtualMeta instantiates a new VirtualMeta object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVirtualMeta() *VirtualMeta {
	this := VirtualMeta{}
	return &this
}

// NewVirtualMetaWithDefaults instantiates a new VirtualMeta object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVirtualMetaWithDefaults() *VirtualMeta {
	this := VirtualMeta{}
	return &this
}

// GetFunctionSourceType returns the FunctionSourceType field value if set, zero value otherwise.
func (o *VirtualMeta) GetFunctionSourceType() string {
	if o == nil || IsNil(o.FunctionSourceType) {
		var ret string
		return ret
	}
	return *o.FunctionSourceType
}

// GetFunctionSourceTypeOk returns a tuple with the FunctionSourceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualMeta) GetFunctionSourceTypeOk() (*string, bool) {
	if o == nil || IsNil(o.FunctionSourceType) {
		return nil, false
	}
	return o.FunctionSourceType, true
}

// HasFunctionSourceType returns a boolean if a field has been set.
func (o *VirtualMeta) HasFunctionSourceType() bool {
	if o != nil && !IsNil(o.FunctionSourceType) {
		return true
	}

	return false
}

// SetFunctionSourceType gets a reference to the given string and assigns it to the FunctionSourceType field.
func (o *VirtualMeta) SetFunctionSourceType(v string) {
	o.FunctionSourceType = &v
}

// GetFunctionSourceUri returns the FunctionSourceUri field value if set, zero value otherwise.
func (o *VirtualMeta) GetFunctionSourceUri() string {
	if o == nil || IsNil(o.FunctionSourceUri) {
		var ret string
		return ret
	}
	return *o.FunctionSourceUri
}

// GetFunctionSourceUriOk returns a tuple with the FunctionSourceUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualMeta) GetFunctionSourceUriOk() (*string, bool) {
	if o == nil || IsNil(o.FunctionSourceUri) {
		return nil, false
	}
	return o.FunctionSourceUri, true
}

// HasFunctionSourceUri returns a boolean if a field has been set.
func (o *VirtualMeta) HasFunctionSourceUri() bool {
	if o != nil && !IsNil(o.FunctionSourceUri) {
		return true
	}

	return false
}

// SetFunctionSourceUri gets a reference to the given string and assigns it to the FunctionSourceUri field.
func (o *VirtualMeta) SetFunctionSourceUri(v string) {
	o.FunctionSourceUri = &v
}

// GetMethod returns the Method field value if set, zero value otherwise.
func (o *VirtualMeta) GetMethod() string {
	if o == nil || IsNil(o.Method) {
		var ret string
		return ret
	}
	return *o.Method
}

// GetMethodOk returns a tuple with the Method field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualMeta) GetMethodOk() (*string, bool) {
	if o == nil || IsNil(o.Method) {
		return nil, false
	}
	return o.Method, true
}

// HasMethod returns a boolean if a field has been set.
func (o *VirtualMeta) HasMethod() bool {
	if o != nil && !IsNil(o.Method) {
		return true
	}

	return false
}

// SetMethod gets a reference to the given string and assigns it to the Method field.
func (o *VirtualMeta) SetMethod(v string) {
	o.Method = &v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *VirtualMeta) GetPath() string {
	if o == nil || IsNil(o.Path) {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualMeta) GetPathOk() (*string, bool) {
	if o == nil || IsNil(o.Path) {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *VirtualMeta) HasPath() bool {
	if o != nil && !IsNil(o.Path) {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *VirtualMeta) SetPath(v string) {
	o.Path = &v
}

// GetProxyOnError returns the ProxyOnError field value if set, zero value otherwise.
func (o *VirtualMeta) GetProxyOnError() bool {
	if o == nil || IsNil(o.ProxyOnError) {
		var ret bool
		return ret
	}
	return *o.ProxyOnError
}

// GetProxyOnErrorOk returns a tuple with the ProxyOnError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualMeta) GetProxyOnErrorOk() (*bool, bool) {
	if o == nil || IsNil(o.ProxyOnError) {
		return nil, false
	}
	return o.ProxyOnError, true
}

// HasProxyOnError returns a boolean if a field has been set.
func (o *VirtualMeta) HasProxyOnError() bool {
	if o != nil && !IsNil(o.ProxyOnError) {
		return true
	}

	return false
}

// SetProxyOnError gets a reference to the given bool and assigns it to the ProxyOnError field.
func (o *VirtualMeta) SetProxyOnError(v bool) {
	o.ProxyOnError = &v
}

// GetResponseFunctionName returns the ResponseFunctionName field value if set, zero value otherwise.
func (o *VirtualMeta) GetResponseFunctionName() string {
	if o == nil || IsNil(o.ResponseFunctionName) {
		var ret string
		return ret
	}
	return *o.ResponseFunctionName
}

// GetResponseFunctionNameOk returns a tuple with the ResponseFunctionName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualMeta) GetResponseFunctionNameOk() (*string, bool) {
	if o == nil || IsNil(o.ResponseFunctionName) {
		return nil, false
	}
	return o.ResponseFunctionName, true
}

// HasResponseFunctionName returns a boolean if a field has been set.
func (o *VirtualMeta) HasResponseFunctionName() bool {
	if o != nil && !IsNil(o.ResponseFunctionName) {
		return true
	}

	return false
}

// SetResponseFunctionName gets a reference to the given string and assigns it to the ResponseFunctionName field.
func (o *VirtualMeta) SetResponseFunctionName(v string) {
	o.ResponseFunctionName = &v
}

// GetUseSession returns the UseSession field value if set, zero value otherwise.
func (o *VirtualMeta) GetUseSession() bool {
	if o == nil || IsNil(o.UseSession) {
		var ret bool
		return ret
	}
	return *o.UseSession
}

// GetUseSessionOk returns a tuple with the UseSession field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualMeta) GetUseSessionOk() (*bool, bool) {
	if o == nil || IsNil(o.UseSession) {
		return nil, false
	}
	return o.UseSession, true
}

// HasUseSession returns a boolean if a field has been set.
func (o *VirtualMeta) HasUseSession() bool {
	if o != nil && !IsNil(o.UseSession) {
		return true
	}

	return false
}

// SetUseSession gets a reference to the given bool and assigns it to the UseSession field.
func (o *VirtualMeta) SetUseSession(v bool) {
	o.UseSession = &v
}

func (o VirtualMeta) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VirtualMeta) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FunctionSourceType) {
		toSerialize["function_source_type"] = o.FunctionSourceType
	}
	if !IsNil(o.FunctionSourceUri) {
		toSerialize["function_source_uri"] = o.FunctionSourceUri
	}
	if !IsNil(o.Method) {
		toSerialize["method"] = o.Method
	}
	if !IsNil(o.Path) {
		toSerialize["path"] = o.Path
	}
	if !IsNil(o.ProxyOnError) {
		toSerialize["proxy_on_error"] = o.ProxyOnError
	}
	if !IsNil(o.ResponseFunctionName) {
		toSerialize["response_function_name"] = o.ResponseFunctionName
	}
	if !IsNil(o.UseSession) {
		toSerialize["use_session"] = o.UseSession
	}
	return toSerialize, nil
}

type NullableVirtualMeta struct {
	value *VirtualMeta
	isSet bool
}

func (v NullableVirtualMeta) Get() *VirtualMeta {
	return v.value
}

func (v *NullableVirtualMeta) Set(val *VirtualMeta) {
	v.value = val
	v.isSet = true
}

func (v NullableVirtualMeta) IsSet() bool {
	return v.isSet
}

func (v *NullableVirtualMeta) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVirtualMeta(val *VirtualMeta) *NullableVirtualMeta {
	return &NullableVirtualMeta{value: val, isSet: true}
}

func (v NullableVirtualMeta) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVirtualMeta) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


