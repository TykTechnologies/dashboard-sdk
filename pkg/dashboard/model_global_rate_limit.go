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

// checks if the GlobalRateLimit type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GlobalRateLimit{}

// GlobalRateLimit struct for GlobalRateLimit
type GlobalRateLimit struct {
	Disabled *bool    `json:"disabled,omitempty"`
	Per      *float64 `json:"per,omitempty"`
	Rate     *float64 `json:"rate,omitempty"`
}

// NewGlobalRateLimit instantiates a new GlobalRateLimit object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGlobalRateLimit() *GlobalRateLimit {
	this := GlobalRateLimit{}
	return &this
}

// NewGlobalRateLimitWithDefaults instantiates a new GlobalRateLimit object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGlobalRateLimitWithDefaults() *GlobalRateLimit {
	this := GlobalRateLimit{}
	return &this
}

// GetDisabled returns the Disabled field value if set, zero value otherwise.
func (o *GlobalRateLimit) GetDisabled() bool {
	if o == nil || IsNil(o.Disabled) {
		var ret bool
		return ret
	}
	return *o.Disabled
}

// GetDisabledOk returns a tuple with the Disabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalRateLimit) GetDisabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Disabled) {
		return nil, false
	}
	return o.Disabled, true
}

// HasDisabled returns a boolean if a field has been set.
func (o *GlobalRateLimit) HasDisabled() bool {
	if o != nil && !IsNil(o.Disabled) {
		return true
	}

	return false
}

// SetDisabled gets a reference to the given bool and assigns it to the Disabled field.
func (o *GlobalRateLimit) SetDisabled(v bool) {
	o.Disabled = &v
}

// GetPer returns the Per field value if set, zero value otherwise.
func (o *GlobalRateLimit) GetPer() float64 {
	if o == nil || IsNil(o.Per) {
		var ret float64
		return ret
	}
	return *o.Per
}

// GetPerOk returns a tuple with the Per field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalRateLimit) GetPerOk() (*float64, bool) {
	if o == nil || IsNil(o.Per) {
		return nil, false
	}
	return o.Per, true
}

// HasPer returns a boolean if a field has been set.
func (o *GlobalRateLimit) HasPer() bool {
	if o != nil && !IsNil(o.Per) {
		return true
	}

	return false
}

// SetPer gets a reference to the given float64 and assigns it to the Per field.
func (o *GlobalRateLimit) SetPer(v float64) {
	o.Per = &v
}

// GetRate returns the Rate field value if set, zero value otherwise.
func (o *GlobalRateLimit) GetRate() float64 {
	if o == nil || IsNil(o.Rate) {
		var ret float64
		return ret
	}
	return *o.Rate
}

// GetRateOk returns a tuple with the Rate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalRateLimit) GetRateOk() (*float64, bool) {
	if o == nil || IsNil(o.Rate) {
		return nil, false
	}
	return o.Rate, true
}

// HasRate returns a boolean if a field has been set.
func (o *GlobalRateLimit) HasRate() bool {
	if o != nil && !IsNil(o.Rate) {
		return true
	}

	return false
}

// SetRate gets a reference to the given float64 and assigns it to the Rate field.
func (o *GlobalRateLimit) SetRate(v float64) {
	o.Rate = &v
}

func (o GlobalRateLimit) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GlobalRateLimit) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Disabled) {
		toSerialize["disabled"] = o.Disabled
	}
	if !IsNil(o.Per) {
		toSerialize["per"] = o.Per
	}
	if !IsNil(o.Rate) {
		toSerialize["rate"] = o.Rate
	}
	return toSerialize, nil
}

type NullableGlobalRateLimit struct {
	value *GlobalRateLimit
	isSet bool
}

func (v NullableGlobalRateLimit) Get() *GlobalRateLimit {
	return v.value
}

func (v *NullableGlobalRateLimit) Set(val *GlobalRateLimit) {
	v.value = val
	v.isSet = true
}

func (v NullableGlobalRateLimit) IsSet() bool {
	return v.isSet
}

func (v *NullableGlobalRateLimit) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGlobalRateLimit(val *GlobalRateLimit) *NullableGlobalRateLimit {
	return &NullableGlobalRateLimit{value: val, isSet: true}
}

func (v NullableGlobalRateLimit) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGlobalRateLimit) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
