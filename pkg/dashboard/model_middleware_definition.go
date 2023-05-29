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

// checks if the MiddlewareDefinition type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MiddlewareDefinition{}

// MiddlewareDefinition struct for MiddlewareDefinition
type MiddlewareDefinition struct {
	Name *string `json:"name,omitempty"`
	Path *string `json:"path,omitempty"`
	RequireSession *bool `json:"require_session,omitempty"`
}

// NewMiddlewareDefinition instantiates a new MiddlewareDefinition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMiddlewareDefinition() *MiddlewareDefinition {
	this := MiddlewareDefinition{}
	return &this
}

// NewMiddlewareDefinitionWithDefaults instantiates a new MiddlewareDefinition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMiddlewareDefinitionWithDefaults() *MiddlewareDefinition {
	this := MiddlewareDefinition{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *MiddlewareDefinition) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MiddlewareDefinition) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *MiddlewareDefinition) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *MiddlewareDefinition) SetName(v string) {
	o.Name = &v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *MiddlewareDefinition) GetPath() string {
	if o == nil || IsNil(o.Path) {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MiddlewareDefinition) GetPathOk() (*string, bool) {
	if o == nil || IsNil(o.Path) {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *MiddlewareDefinition) HasPath() bool {
	if o != nil && !IsNil(o.Path) {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *MiddlewareDefinition) SetPath(v string) {
	o.Path = &v
}

// GetRequireSession returns the RequireSession field value if set, zero value otherwise.
func (o *MiddlewareDefinition) GetRequireSession() bool {
	if o == nil || IsNil(o.RequireSession) {
		var ret bool
		return ret
	}
	return *o.RequireSession
}

// GetRequireSessionOk returns a tuple with the RequireSession field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MiddlewareDefinition) GetRequireSessionOk() (*bool, bool) {
	if o == nil || IsNil(o.RequireSession) {
		return nil, false
	}
	return o.RequireSession, true
}

// HasRequireSession returns a boolean if a field has been set.
func (o *MiddlewareDefinition) HasRequireSession() bool {
	if o != nil && !IsNil(o.RequireSession) {
		return true
	}

	return false
}

// SetRequireSession gets a reference to the given bool and assigns it to the RequireSession field.
func (o *MiddlewareDefinition) SetRequireSession(v bool) {
	o.RequireSession = &v
}

func (o MiddlewareDefinition) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MiddlewareDefinition) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Path) {
		toSerialize["path"] = o.Path
	}
	if !IsNil(o.RequireSession) {
		toSerialize["require_session"] = o.RequireSession
	}
	return toSerialize, nil
}

type NullableMiddlewareDefinition struct {
	value *MiddlewareDefinition
	isSet bool
}

func (v NullableMiddlewareDefinition) Get() *MiddlewareDefinition {
	return v.value
}

func (v *NullableMiddlewareDefinition) Set(val *MiddlewareDefinition) {
	v.value = val
	v.isSet = true
}

func (v NullableMiddlewareDefinition) IsSet() bool {
	return v.isSet
}

func (v *NullableMiddlewareDefinition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMiddlewareDefinition(val *MiddlewareDefinition) *NullableMiddlewareDefinition {
	return &NullableMiddlewareDefinition{value: val, isSet: true}
}

func (v NullableMiddlewareDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMiddlewareDefinition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


