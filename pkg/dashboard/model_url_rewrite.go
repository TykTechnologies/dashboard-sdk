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

// checks if the URLRewrite type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &URLRewrite{}

// URLRewrite struct for URLRewrite
type URLRewrite struct {
	Enabled   *bool               `json:"enabled,omitempty"`
	Pattern   *string             `json:"pattern,omitempty"`
	RewriteTo *string             `json:"rewriteTo,omitempty"`
	Triggers  []URLRewriteTrigger `json:"triggers,omitempty"`
}

// NewURLRewrite instantiates a new URLRewrite object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewURLRewrite() *URLRewrite {
	this := URLRewrite{}
	return &this
}

// NewURLRewriteWithDefaults instantiates a new URLRewrite object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewURLRewriteWithDefaults() *URLRewrite {
	this := URLRewrite{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *URLRewrite) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *URLRewrite) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *URLRewrite) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *URLRewrite) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetPattern returns the Pattern field value if set, zero value otherwise.
func (o *URLRewrite) GetPattern() string {
	if o == nil || IsNil(o.Pattern) {
		var ret string
		return ret
	}
	return *o.Pattern
}

// GetPatternOk returns a tuple with the Pattern field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *URLRewrite) GetPatternOk() (*string, bool) {
	if o == nil || IsNil(o.Pattern) {
		return nil, false
	}
	return o.Pattern, true
}

// HasPattern returns a boolean if a field has been set.
func (o *URLRewrite) HasPattern() bool {
	if o != nil && !IsNil(o.Pattern) {
		return true
	}

	return false
}

// SetPattern gets a reference to the given string and assigns it to the Pattern field.
func (o *URLRewrite) SetPattern(v string) {
	o.Pattern = &v
}

// GetRewriteTo returns the RewriteTo field value if set, zero value otherwise.
func (o *URLRewrite) GetRewriteTo() string {
	if o == nil || IsNil(o.RewriteTo) {
		var ret string
		return ret
	}
	return *o.RewriteTo
}

// GetRewriteToOk returns a tuple with the RewriteTo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *URLRewrite) GetRewriteToOk() (*string, bool) {
	if o == nil || IsNil(o.RewriteTo) {
		return nil, false
	}
	return o.RewriteTo, true
}

// HasRewriteTo returns a boolean if a field has been set.
func (o *URLRewrite) HasRewriteTo() bool {
	if o != nil && !IsNil(o.RewriteTo) {
		return true
	}

	return false
}

// SetRewriteTo gets a reference to the given string and assigns it to the RewriteTo field.
func (o *URLRewrite) SetRewriteTo(v string) {
	o.RewriteTo = &v
}

// GetTriggers returns the Triggers field value if set, zero value otherwise.
func (o *URLRewrite) GetTriggers() []URLRewriteTrigger {
	if o == nil || IsNil(o.Triggers) {
		var ret []URLRewriteTrigger
		return ret
	}
	return o.Triggers
}

// GetTriggersOk returns a tuple with the Triggers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *URLRewrite) GetTriggersOk() ([]URLRewriteTrigger, bool) {
	if o == nil || IsNil(o.Triggers) {
		return nil, false
	}
	return o.Triggers, true
}

// HasTriggers returns a boolean if a field has been set.
func (o *URLRewrite) HasTriggers() bool {
	if o != nil && !IsNil(o.Triggers) {
		return true
	}

	return false
}

// SetTriggers gets a reference to the given []URLRewriteTrigger and assigns it to the Triggers field.
func (o *URLRewrite) SetTriggers(v []URLRewriteTrigger) {
	o.Triggers = v
}

func (o URLRewrite) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o URLRewrite) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.Pattern) {
		toSerialize["pattern"] = o.Pattern
	}
	if !IsNil(o.RewriteTo) {
		toSerialize["rewriteTo"] = o.RewriteTo
	}
	if !IsNil(o.Triggers) {
		toSerialize["triggers"] = o.Triggers
	}
	return toSerialize, nil
}

type NullableURLRewrite struct {
	value *URLRewrite
	isSet bool
}

func (v NullableURLRewrite) Get() *URLRewrite {
	return v.value
}

func (v *NullableURLRewrite) Set(val *URLRewrite) {
	v.value = val
	v.isSet = true
}

func (v NullableURLRewrite) IsSet() bool {
	return v.isSet
}

func (v *NullableURLRewrite) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableURLRewrite(val *URLRewrite) *NullableURLRewrite {
	return &NullableURLRewrite{value: val, isSet: true}
}

func (v NullableURLRewrite) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableURLRewrite) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
