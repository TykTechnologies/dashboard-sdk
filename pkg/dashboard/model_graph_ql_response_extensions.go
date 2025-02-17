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

// checks if the GraphQLResponseExtensions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GraphQLResponseExtensions{}

// GraphQLResponseExtensions struct for GraphQLResponseExtensions
type GraphQLResponseExtensions struct {
	OnErrorForwarding *bool `json:"on_error_forwarding,omitempty"`
}

// NewGraphQLResponseExtensions instantiates a new GraphQLResponseExtensions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGraphQLResponseExtensions() *GraphQLResponseExtensions {
	this := GraphQLResponseExtensions{}
	return &this
}

// NewGraphQLResponseExtensionsWithDefaults instantiates a new GraphQLResponseExtensions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGraphQLResponseExtensionsWithDefaults() *GraphQLResponseExtensions {
	this := GraphQLResponseExtensions{}
	return &this
}

// GetOnErrorForwarding returns the OnErrorForwarding field value if set, zero value otherwise.
func (o *GraphQLResponseExtensions) GetOnErrorForwarding() bool {
	if o == nil || IsNil(o.OnErrorForwarding) {
		var ret bool
		return ret
	}
	return *o.OnErrorForwarding
}

// GetOnErrorForwardingOk returns a tuple with the OnErrorForwarding field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GraphQLResponseExtensions) GetOnErrorForwardingOk() (*bool, bool) {
	if o == nil || IsNil(o.OnErrorForwarding) {
		return nil, false
	}
	return o.OnErrorForwarding, true
}

// HasOnErrorForwarding returns a boolean if a field has been set.
func (o *GraphQLResponseExtensions) HasOnErrorForwarding() bool {
	if o != nil && !IsNil(o.OnErrorForwarding) {
		return true
	}

	return false
}

// SetOnErrorForwarding gets a reference to the given bool and assigns it to the OnErrorForwarding field.
func (o *GraphQLResponseExtensions) SetOnErrorForwarding(v bool) {
	o.OnErrorForwarding = &v
}

func (o GraphQLResponseExtensions) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GraphQLResponseExtensions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.OnErrorForwarding) {
		toSerialize["on_error_forwarding"] = o.OnErrorForwarding
	}
	return toSerialize, nil
}

type NullableGraphQLResponseExtensions struct {
	value *GraphQLResponseExtensions
	isSet bool
}

func (v NullableGraphQLResponseExtensions) Get() *GraphQLResponseExtensions {
	return v.value
}

func (v *NullableGraphQLResponseExtensions) Set(val *GraphQLResponseExtensions) {
	v.value = val
	v.isSet = true
}

func (v NullableGraphQLResponseExtensions) IsSet() bool {
	return v.isSet
}

func (v *NullableGraphQLResponseExtensions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGraphQLResponseExtensions(val *GraphQLResponseExtensions) *NullableGraphQLResponseExtensions {
	return &NullableGraphQLResponseExtensions{value: val, isSet: true}
}

func (v NullableGraphQLResponseExtensions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGraphQLResponseExtensions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
