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

// checks if the StringRegexMap type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StringRegexMap{}

// StringRegexMap struct for StringRegexMap
type StringRegexMap struct {
	MatchRx *string `json:"match_rx,omitempty"`
	Reverse *bool `json:"reverse,omitempty"`
}

// NewStringRegexMap instantiates a new StringRegexMap object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStringRegexMap() *StringRegexMap {
	this := StringRegexMap{}
	return &this
}

// NewStringRegexMapWithDefaults instantiates a new StringRegexMap object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStringRegexMapWithDefaults() *StringRegexMap {
	this := StringRegexMap{}
	return &this
}

// GetMatchRx returns the MatchRx field value if set, zero value otherwise.
func (o *StringRegexMap) GetMatchRx() string {
	if o == nil || IsNil(o.MatchRx) {
		var ret string
		return ret
	}
	return *o.MatchRx
}

// GetMatchRxOk returns a tuple with the MatchRx field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StringRegexMap) GetMatchRxOk() (*string, bool) {
	if o == nil || IsNil(o.MatchRx) {
		return nil, false
	}
	return o.MatchRx, true
}

// HasMatchRx returns a boolean if a field has been set.
func (o *StringRegexMap) HasMatchRx() bool {
	if o != nil && !IsNil(o.MatchRx) {
		return true
	}

	return false
}

// SetMatchRx gets a reference to the given string and assigns it to the MatchRx field.
func (o *StringRegexMap) SetMatchRx(v string) {
	o.MatchRx = &v
}

// GetReverse returns the Reverse field value if set, zero value otherwise.
func (o *StringRegexMap) GetReverse() bool {
	if o == nil || IsNil(o.Reverse) {
		var ret bool
		return ret
	}
	return *o.Reverse
}

// GetReverseOk returns a tuple with the Reverse field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StringRegexMap) GetReverseOk() (*bool, bool) {
	if o == nil || IsNil(o.Reverse) {
		return nil, false
	}
	return o.Reverse, true
}

// HasReverse returns a boolean if a field has been set.
func (o *StringRegexMap) HasReverse() bool {
	if o != nil && !IsNil(o.Reverse) {
		return true
	}

	return false
}

// SetReverse gets a reference to the given bool and assigns it to the Reverse field.
func (o *StringRegexMap) SetReverse(v bool) {
	o.Reverse = &v
}

func (o StringRegexMap) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StringRegexMap) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MatchRx) {
		toSerialize["match_rx"] = o.MatchRx
	}
	if !IsNil(o.Reverse) {
		toSerialize["reverse"] = o.Reverse
	}
	return toSerialize, nil
}

type NullableStringRegexMap struct {
	value *StringRegexMap
	isSet bool
}

func (v NullableStringRegexMap) Get() *StringRegexMap {
	return v.value
}

func (v *NullableStringRegexMap) Set(val *StringRegexMap) {
	v.value = val
	v.isSet = true
}

func (v NullableStringRegexMap) IsSet() bool {
	return v.isSet
}

func (v *NullableStringRegexMap) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStringRegexMap(val *StringRegexMap) *NullableStringRegexMap {
	return &NullableStringRegexMap{value: val, isSet: true}
}

func (v NullableStringRegexMap) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStringRegexMap) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


