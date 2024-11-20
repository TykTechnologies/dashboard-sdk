/*
Tyk Dashboard API

 ## <a name=\"introduction\"></a> Introduction  The Tyk Dashboard API offers granular, programmatic access to a centralised database of resources that your Tyk nodes can pull from. This API has a dynamic user administrative structure which means the secret key that is used to communicate with your Tyk nodes can be kept secret and access to the wider management functions can be handled on a user-by-user and organisation-by-organisation basis.  A common question around using a database-backed configuration is how to programmatically add API definitions to your Tyk nodes, the Dashboard API allows much more fine-grained, secure and multi-user access to your Tyk cluster, and should be used to manage a database-backed Tyk node.  The Tyk Dashboard API works seamlessly with the Tyk Dashboard (and the two come bundled together).  ## <a name=\"security-hierarchy\"></a> Security Hierarchy  The Dashboard API provides a more structured security layer to managing Tyk nodes.  ### Organisations, APIs and Users  With the Dashboard API and a database-backed Tyk setup, (and to an extent with file-based API setups - if diligence is used in naming and creating definitions), the following security model is applied to the management of Upstream APIs:  * **Organisations**: All APIs are *owned* by an organisation, this is designated by the 'OrgID' parameter in the API Definition. * **Users**: All users created in the Dashboard belong to an organisation (unless an exception is made for super-administrative access). * **APIs**: All APIs belong to an Organisation and only Users that belong to that organisation can see the analytics for those APIs and manage their configurations. * **API Keys**: API Keys are designated by organisation, this means an API key that has full access rights will not be allowed to access the APIs of another organisation on the same system, but can have full access to all APIs within the organisation. * **Access Rights**: Access rights are stored with the key, this enables a key to give access to multiple APIs, this is defined by the session object in the core Tyk API.  In order to use the Dashboard API, you'll need to get the 'Tyk Dashboard API Access Credentials' secret from your user profile on the Dashboard UI.  The secret you set should then be sent along as a header with each Dashboard API Request in order for it to be successful:   authorization: <your-secret>

API version: 5.7.0
Contact: support@tyk.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dashboard

import (
	"encoding/json"
)

// checks if the OASSchemaResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OASSchemaResponse{}

// OASSchemaResponse struct for OASSchemaResponse
type OASSchemaResponse struct {
	Message *string     `json:"message,omitempty"`
	Schema  interface{} `json:"schema,omitempty"`
	Status  *string     `json:"status,omitempty"`
}

// NewOASSchemaResponse instantiates a new OASSchemaResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOASSchemaResponse() *OASSchemaResponse {
	this := OASSchemaResponse{}
	return &this
}

// NewOASSchemaResponseWithDefaults instantiates a new OASSchemaResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOASSchemaResponseWithDefaults() *OASSchemaResponse {
	this := OASSchemaResponse{}
	return &this
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *OASSchemaResponse) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OASSchemaResponse) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *OASSchemaResponse) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *OASSchemaResponse) SetMessage(v string) {
	o.Message = &v
}

// GetSchema returns the Schema field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OASSchemaResponse) GetSchema() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Schema
}

// GetSchemaOk returns a tuple with the Schema field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OASSchemaResponse) GetSchemaOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Schema) {
		return nil, false
	}
	return &o.Schema, true
}

// HasSchema returns a boolean if a field has been set.
func (o *OASSchemaResponse) HasSchema() bool {
	if o != nil && !IsNil(o.Schema) {
		return true
	}

	return false
}

// SetSchema gets a reference to the given interface{} and assigns it to the Schema field.
func (o *OASSchemaResponse) SetSchema(v interface{}) {
	o.Schema = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *OASSchemaResponse) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OASSchemaResponse) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *OASSchemaResponse) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *OASSchemaResponse) SetStatus(v string) {
	o.Status = &v
}

func (o OASSchemaResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OASSchemaResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if o.Schema != nil {
		toSerialize["schema"] = o.Schema
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}

type NullableOASSchemaResponse struct {
	value *OASSchemaResponse
	isSet bool
}

func (v NullableOASSchemaResponse) Get() *OASSchemaResponse {
	return v.value
}

func (v *NullableOASSchemaResponse) Set(val *OASSchemaResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableOASSchemaResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableOASSchemaResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOASSchemaResponse(val *OASSchemaResponse) *NullableOASSchemaResponse {
	return &NullableOASSchemaResponse{value: val, isSet: true}
}

func (v NullableOASSchemaResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOASSchemaResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
