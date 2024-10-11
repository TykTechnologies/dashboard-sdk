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

// checks if the EventHandler type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EventHandler{}

// EventHandler struct for EventHandler
type EventHandler struct {
	Enabled *bool   `json:"enabled,omitempty"`
	Id      *string `json:"id,omitempty"`
	Name    *string `json:"name,omitempty"`
	Trigger *string `json:"trigger,omitempty"`
	Type    *string `json:"type,omitempty"`
}

// NewEventHandler instantiates a new EventHandler object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEventHandler() *EventHandler {
	this := EventHandler{}
	return &this
}

// NewEventHandlerWithDefaults instantiates a new EventHandler object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventHandlerWithDefaults() *EventHandler {
	this := EventHandler{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *EventHandler) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventHandler) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *EventHandler) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *EventHandler) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *EventHandler) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventHandler) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *EventHandler) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *EventHandler) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *EventHandler) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventHandler) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *EventHandler) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *EventHandler) SetName(v string) {
	o.Name = &v
}

// GetTrigger returns the Trigger field value if set, zero value otherwise.
func (o *EventHandler) GetTrigger() string {
	if o == nil || IsNil(o.Trigger) {
		var ret string
		return ret
	}
	return *o.Trigger
}

// GetTriggerOk returns a tuple with the Trigger field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventHandler) GetTriggerOk() (*string, bool) {
	if o == nil || IsNil(o.Trigger) {
		return nil, false
	}
	return o.Trigger, true
}

// HasTrigger returns a boolean if a field has been set.
func (o *EventHandler) HasTrigger() bool {
	if o != nil && !IsNil(o.Trigger) {
		return true
	}

	return false
}

// SetTrigger gets a reference to the given string and assigns it to the Trigger field.
func (o *EventHandler) SetTrigger(v string) {
	o.Trigger = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *EventHandler) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventHandler) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *EventHandler) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *EventHandler) SetType(v string) {
	o.Type = &v
}

func (o EventHandler) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EventHandler) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Trigger) {
		toSerialize["trigger"] = o.Trigger
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableEventHandler struct {
	value *EventHandler
	isSet bool
}

func (v NullableEventHandler) Get() *EventHandler {
	return v.value
}

func (v *NullableEventHandler) Set(val *EventHandler) {
	v.value = val
	v.isSet = true
}

func (v NullableEventHandler) IsSet() bool {
	return v.isSet
}

func (v *NullableEventHandler) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventHandler(val *EventHandler) *NullableEventHandler {
	return &NullableEventHandler{value: val, isSet: true}
}

func (v NullableEventHandler) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventHandler) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
