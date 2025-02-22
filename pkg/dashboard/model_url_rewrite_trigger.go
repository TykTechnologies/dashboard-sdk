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

// checks if the URLRewriteTrigger type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &URLRewriteTrigger{}

// URLRewriteTrigger struct for URLRewriteTrigger
type URLRewriteTrigger struct {
	Condition *string          `json:"condition,omitempty"`
	RewriteTo *string          `json:"rewriteTo,omitempty"`
	Rules     []URLRewriteRule `json:"rules,omitempty"`
}

// NewURLRewriteTrigger instantiates a new URLRewriteTrigger object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewURLRewriteTrigger() *URLRewriteTrigger {
	this := URLRewriteTrigger{}
	return &this
}

// NewURLRewriteTriggerWithDefaults instantiates a new URLRewriteTrigger object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewURLRewriteTriggerWithDefaults() *URLRewriteTrigger {
	this := URLRewriteTrigger{}
	return &this
}

// GetCondition returns the Condition field value if set, zero value otherwise.
func (o *URLRewriteTrigger) GetCondition() string {
	if o == nil || IsNil(o.Condition) {
		var ret string
		return ret
	}
	return *o.Condition
}

// GetConditionOk returns a tuple with the Condition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *URLRewriteTrigger) GetConditionOk() (*string, bool) {
	if o == nil || IsNil(o.Condition) {
		return nil, false
	}
	return o.Condition, true
}

// HasCondition returns a boolean if a field has been set.
func (o *URLRewriteTrigger) HasCondition() bool {
	if o != nil && !IsNil(o.Condition) {
		return true
	}

	return false
}

// SetCondition gets a reference to the given string and assigns it to the Condition field.
func (o *URLRewriteTrigger) SetCondition(v string) {
	o.Condition = &v
}

// GetRewriteTo returns the RewriteTo field value if set, zero value otherwise.
func (o *URLRewriteTrigger) GetRewriteTo() string {
	if o == nil || IsNil(o.RewriteTo) {
		var ret string
		return ret
	}
	return *o.RewriteTo
}

// GetRewriteToOk returns a tuple with the RewriteTo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *URLRewriteTrigger) GetRewriteToOk() (*string, bool) {
	if o == nil || IsNil(o.RewriteTo) {
		return nil, false
	}
	return o.RewriteTo, true
}

// HasRewriteTo returns a boolean if a field has been set.
func (o *URLRewriteTrigger) HasRewriteTo() bool {
	if o != nil && !IsNil(o.RewriteTo) {
		return true
	}

	return false
}

// SetRewriteTo gets a reference to the given string and assigns it to the RewriteTo field.
func (o *URLRewriteTrigger) SetRewriteTo(v string) {
	o.RewriteTo = &v
}

// GetRules returns the Rules field value if set, zero value otherwise.
func (o *URLRewriteTrigger) GetRules() []URLRewriteRule {
	if o == nil || IsNil(o.Rules) {
		var ret []URLRewriteRule
		return ret
	}
	return o.Rules
}

// GetRulesOk returns a tuple with the Rules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *URLRewriteTrigger) GetRulesOk() ([]URLRewriteRule, bool) {
	if o == nil || IsNil(o.Rules) {
		return nil, false
	}
	return o.Rules, true
}

// HasRules returns a boolean if a field has been set.
func (o *URLRewriteTrigger) HasRules() bool {
	if o != nil && !IsNil(o.Rules) {
		return true
	}

	return false
}

// SetRules gets a reference to the given []URLRewriteRule and assigns it to the Rules field.
func (o *URLRewriteTrigger) SetRules(v []URLRewriteRule) {
	o.Rules = v
}

func (o URLRewriteTrigger) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o URLRewriteTrigger) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Condition) {
		toSerialize["condition"] = o.Condition
	}
	if !IsNil(o.RewriteTo) {
		toSerialize["rewriteTo"] = o.RewriteTo
	}
	if !IsNil(o.Rules) {
		toSerialize["rules"] = o.Rules
	}
	return toSerialize, nil
}

type NullableURLRewriteTrigger struct {
	value *URLRewriteTrigger
	isSet bool
}

func (v NullableURLRewriteTrigger) Get() *URLRewriteTrigger {
	return v.value
}

func (v *NullableURLRewriteTrigger) Set(val *URLRewriteTrigger) {
	v.value = val
	v.isSet = true
}

func (v NullableURLRewriteTrigger) IsSet() bool {
	return v.isSet
}

func (v *NullableURLRewriteTrigger) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableURLRewriteTrigger(val *URLRewriteTrigger) *NullableURLRewriteTrigger {
	return &NullableURLRewriteTrigger{value: val, isSet: true}
}

func (v NullableURLRewriteTrigger) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableURLRewriteTrigger) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
