/*
Tyk Dashboard API

 ## <a name=\"introduction\"></a> Introduction  The Tyk Dashboard API offers granular, programmatic access to a centralised database of resources that your Tyk nodes can pull from. This API has a dynamic user administrative structure which means the secret key that is used to communicate with your Tyk nodes can be kept secret and access to the wider management functions can be handled on a user-by-user and organisation-by-organisation basis.  A common question around using a database-backed configuration is how to programmatically add API definitions to your Tyk nodes, the Dashboard API allows much more fine-grained, secure and multi-user access to your Tyk cluster, and should be used to manage a database-backed Tyk node.  The Tyk Dashboard API works seamlessly with the Tyk Dashboard (and the two come bundled together).  ## <a name=\"security-hierarchy\"></a> Security Hierarchy  The Dashboard API provides a more structured security layer to managing Tyk nodes.  ### Organisations, APIs and Users  With the Dashboard API and a database-backed Tyk setup, (and to an extent with file-based API setups - if diligence is used in naming and creating definitions), the following security model is applied to the management of Upstream APIs:  * **Organisations**: All APIs are *owned* by an organisation, this is designated by the 'OrgID' parameter in the API Definition. * **Users**: All users created in the Dashboard belong to an organisation (unless an exception is made for super-administrative access). * **APIs**: All APIs belong to an Organisation and only Users that belong to that organisation can see the analytics for those APIs and manage their configurations. * **API Keys**: API Keys are designated by organisation, this means an API key that has full access rights will not be allowed to access the APIs of another organisation on the same system, but can have full access to all APIs within the organisation. * **Access Rights**: Access rights are stored with the key, this enables a key to give access to multiple APIs, this is defined by the session object in the core Tyk API.  In order to use the Dashboard API, you'll need to get the 'Tyk Dashboard API Access Credentials' secret from your user profile on the Dashboard UI.  The secret you set should then be sent along as a header with each Dashboard API Request in order for it to be successful:   authorization: <your-secret>

API version: 5.7.1
Contact: support@tyk.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dashboard

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the DryRunApiOAS200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DryRunApiOAS200Response{}

// DryRunApiOAS200Response struct for DryRunApiOAS200Response
type DryRunApiOAS200Response struct {
	XTykApiGateway *XTykAPIGateway        `json:"x-tyk-api-gateway,omitempty"`
	Openapi        string                 `json:"openapi" validate:"regexp=^3\\\\.0\\\\.\\\\d(-.+)?$"`
	Info           Info1                  `json:"info"`
	ExternalDocs   *ExternalDocumentation `json:"externalDocs,omitempty"`
	Servers        []Server1              `json:"servers,omitempty"`
	Security       []map[string][]string  `json:"security,omitempty"`
	Tags           []Tag1                 `json:"tags,omitempty"`
	Paths          map[string]interface{} `json:"paths"`
	Components     *Components1           `json:"components,omitempty"`
}

type _DryRunApiOAS200Response DryRunApiOAS200Response

// NewDryRunApiOAS200Response instantiates a new DryRunApiOAS200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDryRunApiOAS200Response(openapi string, info Info1, paths map[string]interface{}) *DryRunApiOAS200Response {
	this := DryRunApiOAS200Response{}
	this.Openapi = openapi
	this.Info = info
	this.Paths = paths
	return &this
}

// NewDryRunApiOAS200ResponseWithDefaults instantiates a new DryRunApiOAS200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDryRunApiOAS200ResponseWithDefaults() *DryRunApiOAS200Response {
	this := DryRunApiOAS200Response{}
	return &this
}

// GetXTykApiGateway returns the XTykApiGateway field value if set, zero value otherwise.
func (o *DryRunApiOAS200Response) GetXTykApiGateway() XTykAPIGateway {
	if o == nil || IsNil(o.XTykApiGateway) {
		var ret XTykAPIGateway
		return ret
	}
	return *o.XTykApiGateway
}

// GetXTykApiGatewayOk returns a tuple with the XTykApiGateway field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DryRunApiOAS200Response) GetXTykApiGatewayOk() (*XTykAPIGateway, bool) {
	if o == nil || IsNil(o.XTykApiGateway) {
		return nil, false
	}
	return o.XTykApiGateway, true
}

// HasXTykApiGateway returns a boolean if a field has been set.
func (o *DryRunApiOAS200Response) HasXTykApiGateway() bool {
	if o != nil && !IsNil(o.XTykApiGateway) {
		return true
	}

	return false
}

// SetXTykApiGateway gets a reference to the given XTykAPIGateway and assigns it to the XTykApiGateway field.
func (o *DryRunApiOAS200Response) SetXTykApiGateway(v XTykAPIGateway) {
	o.XTykApiGateway = &v
}

// GetOpenapi returns the Openapi field value
func (o *DryRunApiOAS200Response) GetOpenapi() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Openapi
}

// GetOpenapiOk returns a tuple with the Openapi field value
// and a boolean to check if the value has been set.
func (o *DryRunApiOAS200Response) GetOpenapiOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Openapi, true
}

// SetOpenapi sets field value
func (o *DryRunApiOAS200Response) SetOpenapi(v string) {
	o.Openapi = v
}

// GetInfo returns the Info field value
func (o *DryRunApiOAS200Response) GetInfo() Info1 {
	if o == nil {
		var ret Info1
		return ret
	}

	return o.Info
}

// GetInfoOk returns a tuple with the Info field value
// and a boolean to check if the value has been set.
func (o *DryRunApiOAS200Response) GetInfoOk() (*Info1, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Info, true
}

// SetInfo sets field value
func (o *DryRunApiOAS200Response) SetInfo(v Info1) {
	o.Info = v
}

// GetExternalDocs returns the ExternalDocs field value if set, zero value otherwise.
func (o *DryRunApiOAS200Response) GetExternalDocs() ExternalDocumentation {
	if o == nil || IsNil(o.ExternalDocs) {
		var ret ExternalDocumentation
		return ret
	}
	return *o.ExternalDocs
}

// GetExternalDocsOk returns a tuple with the ExternalDocs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DryRunApiOAS200Response) GetExternalDocsOk() (*ExternalDocumentation, bool) {
	if o == nil || IsNil(o.ExternalDocs) {
		return nil, false
	}
	return o.ExternalDocs, true
}

// HasExternalDocs returns a boolean if a field has been set.
func (o *DryRunApiOAS200Response) HasExternalDocs() bool {
	if o != nil && !IsNil(o.ExternalDocs) {
		return true
	}

	return false
}

// SetExternalDocs gets a reference to the given ExternalDocumentation and assigns it to the ExternalDocs field.
func (o *DryRunApiOAS200Response) SetExternalDocs(v ExternalDocumentation) {
	o.ExternalDocs = &v
}

// GetServers returns the Servers field value if set, zero value otherwise.
func (o *DryRunApiOAS200Response) GetServers() []Server1 {
	if o == nil || IsNil(o.Servers) {
		var ret []Server1
		return ret
	}
	return o.Servers
}

// GetServersOk returns a tuple with the Servers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DryRunApiOAS200Response) GetServersOk() ([]Server1, bool) {
	if o == nil || IsNil(o.Servers) {
		return nil, false
	}
	return o.Servers, true
}

// HasServers returns a boolean if a field has been set.
func (o *DryRunApiOAS200Response) HasServers() bool {
	if o != nil && !IsNil(o.Servers) {
		return true
	}

	return false
}

// SetServers gets a reference to the given []Server1 and assigns it to the Servers field.
func (o *DryRunApiOAS200Response) SetServers(v []Server1) {
	o.Servers = v
}

// GetSecurity returns the Security field value if set, zero value otherwise.
func (o *DryRunApiOAS200Response) GetSecurity() []map[string][]string {
	if o == nil || IsNil(o.Security) {
		var ret []map[string][]string
		return ret
	}
	return o.Security
}

// GetSecurityOk returns a tuple with the Security field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DryRunApiOAS200Response) GetSecurityOk() ([]map[string][]string, bool) {
	if o == nil || IsNil(o.Security) {
		return nil, false
	}
	return o.Security, true
}

// HasSecurity returns a boolean if a field has been set.
func (o *DryRunApiOAS200Response) HasSecurity() bool {
	if o != nil && !IsNil(o.Security) {
		return true
	}

	return false
}

// SetSecurity gets a reference to the given []map[string][]string and assigns it to the Security field.
func (o *DryRunApiOAS200Response) SetSecurity(v []map[string][]string) {
	o.Security = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *DryRunApiOAS200Response) GetTags() []Tag1 {
	if o == nil || IsNil(o.Tags) {
		var ret []Tag1
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DryRunApiOAS200Response) GetTagsOk() ([]Tag1, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *DryRunApiOAS200Response) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []Tag1 and assigns it to the Tags field.
func (o *DryRunApiOAS200Response) SetTags(v []Tag1) {
	o.Tags = v
}

// GetPaths returns the Paths field value
func (o *DryRunApiOAS200Response) GetPaths() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Paths
}

// GetPathsOk returns a tuple with the Paths field value
// and a boolean to check if the value has been set.
func (o *DryRunApiOAS200Response) GetPathsOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Paths, true
}

// SetPaths sets field value
func (o *DryRunApiOAS200Response) SetPaths(v map[string]interface{}) {
	o.Paths = v
}

// GetComponents returns the Components field value if set, zero value otherwise.
func (o *DryRunApiOAS200Response) GetComponents() Components1 {
	if o == nil || IsNil(o.Components) {
		var ret Components1
		return ret
	}
	return *o.Components
}

// GetComponentsOk returns a tuple with the Components field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DryRunApiOAS200Response) GetComponentsOk() (*Components1, bool) {
	if o == nil || IsNil(o.Components) {
		return nil, false
	}
	return o.Components, true
}

// HasComponents returns a boolean if a field has been set.
func (o *DryRunApiOAS200Response) HasComponents() bool {
	if o != nil && !IsNil(o.Components) {
		return true
	}

	return false
}

// SetComponents gets a reference to the given Components1 and assigns it to the Components field.
func (o *DryRunApiOAS200Response) SetComponents(v Components1) {
	o.Components = &v
}

func (o DryRunApiOAS200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DryRunApiOAS200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.XTykApiGateway) {
		toSerialize["x-tyk-api-gateway"] = o.XTykApiGateway
	}
	toSerialize["openapi"] = o.Openapi
	toSerialize["info"] = o.Info
	if !IsNil(o.ExternalDocs) {
		toSerialize["externalDocs"] = o.ExternalDocs
	}
	if !IsNil(o.Servers) {
		toSerialize["servers"] = o.Servers
	}
	if !IsNil(o.Security) {
		toSerialize["security"] = o.Security
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	toSerialize["paths"] = o.Paths
	if !IsNil(o.Components) {
		toSerialize["components"] = o.Components
	}
	return toSerialize, nil
}

func (o *DryRunApiOAS200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"openapi",
		"info",
		"paths",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varDryRunApiOAS200Response := _DryRunApiOAS200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDryRunApiOAS200Response)

	if err != nil {
		return err
	}

	*o = DryRunApiOAS200Response(varDryRunApiOAS200Response)

	return err
}

type NullableDryRunApiOAS200Response struct {
	value *DryRunApiOAS200Response
	isSet bool
}

func (v NullableDryRunApiOAS200Response) Get() *DryRunApiOAS200Response {
	return v.value
}

func (v *NullableDryRunApiOAS200Response) Set(val *DryRunApiOAS200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableDryRunApiOAS200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableDryRunApiOAS200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDryRunApiOAS200Response(val *DryRunApiOAS200Response) *NullableDryRunApiOAS200Response {
	return &NullableDryRunApiOAS200Response{value: val, isSet: true}
}

func (v NullableDryRunApiOAS200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDryRunApiOAS200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
