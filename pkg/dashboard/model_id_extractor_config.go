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

// checks if the IDExtractorConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IDExtractorConfig{}

// IDExtractorConfig struct for IDExtractorConfig
type IDExtractorConfig struct {
	FormParamName    *string `json:"formParamName,omitempty"`
	HeaderName       *string `json:"headerName,omitempty"`
	Regexp           *string `json:"regexp,omitempty"`
	RegexpMatchIndex *int32  `json:"regexpMatchIndex,omitempty"`
	XPathExp         *string `json:"xPathExp,omitempty"`
}

// NewIDExtractorConfig instantiates a new IDExtractorConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIDExtractorConfig() *IDExtractorConfig {
	this := IDExtractorConfig{}
	return &this
}

// NewIDExtractorConfigWithDefaults instantiates a new IDExtractorConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIDExtractorConfigWithDefaults() *IDExtractorConfig {
	this := IDExtractorConfig{}
	return &this
}

// GetFormParamName returns the FormParamName field value if set, zero value otherwise.
func (o *IDExtractorConfig) GetFormParamName() string {
	if o == nil || IsNil(o.FormParamName) {
		var ret string
		return ret
	}
	return *o.FormParamName
}

// GetFormParamNameOk returns a tuple with the FormParamName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IDExtractorConfig) GetFormParamNameOk() (*string, bool) {
	if o == nil || IsNil(o.FormParamName) {
		return nil, false
	}
	return o.FormParamName, true
}

// HasFormParamName returns a boolean if a field has been set.
func (o *IDExtractorConfig) HasFormParamName() bool {
	if o != nil && !IsNil(o.FormParamName) {
		return true
	}

	return false
}

// SetFormParamName gets a reference to the given string and assigns it to the FormParamName field.
func (o *IDExtractorConfig) SetFormParamName(v string) {
	o.FormParamName = &v
}

// GetHeaderName returns the HeaderName field value if set, zero value otherwise.
func (o *IDExtractorConfig) GetHeaderName() string {
	if o == nil || IsNil(o.HeaderName) {
		var ret string
		return ret
	}
	return *o.HeaderName
}

// GetHeaderNameOk returns a tuple with the HeaderName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IDExtractorConfig) GetHeaderNameOk() (*string, bool) {
	if o == nil || IsNil(o.HeaderName) {
		return nil, false
	}
	return o.HeaderName, true
}

// HasHeaderName returns a boolean if a field has been set.
func (o *IDExtractorConfig) HasHeaderName() bool {
	if o != nil && !IsNil(o.HeaderName) {
		return true
	}

	return false
}

// SetHeaderName gets a reference to the given string and assigns it to the HeaderName field.
func (o *IDExtractorConfig) SetHeaderName(v string) {
	o.HeaderName = &v
}

// GetRegexp returns the Regexp field value if set, zero value otherwise.
func (o *IDExtractorConfig) GetRegexp() string {
	if o == nil || IsNil(o.Regexp) {
		var ret string
		return ret
	}
	return *o.Regexp
}

// GetRegexpOk returns a tuple with the Regexp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IDExtractorConfig) GetRegexpOk() (*string, bool) {
	if o == nil || IsNil(o.Regexp) {
		return nil, false
	}
	return o.Regexp, true
}

// HasRegexp returns a boolean if a field has been set.
func (o *IDExtractorConfig) HasRegexp() bool {
	if o != nil && !IsNil(o.Regexp) {
		return true
	}

	return false
}

// SetRegexp gets a reference to the given string and assigns it to the Regexp field.
func (o *IDExtractorConfig) SetRegexp(v string) {
	o.Regexp = &v
}

// GetRegexpMatchIndex returns the RegexpMatchIndex field value if set, zero value otherwise.
func (o *IDExtractorConfig) GetRegexpMatchIndex() int32 {
	if o == nil || IsNil(o.RegexpMatchIndex) {
		var ret int32
		return ret
	}
	return *o.RegexpMatchIndex
}

// GetRegexpMatchIndexOk returns a tuple with the RegexpMatchIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IDExtractorConfig) GetRegexpMatchIndexOk() (*int32, bool) {
	if o == nil || IsNil(o.RegexpMatchIndex) {
		return nil, false
	}
	return o.RegexpMatchIndex, true
}

// HasRegexpMatchIndex returns a boolean if a field has been set.
func (o *IDExtractorConfig) HasRegexpMatchIndex() bool {
	if o != nil && !IsNil(o.RegexpMatchIndex) {
		return true
	}

	return false
}

// SetRegexpMatchIndex gets a reference to the given int32 and assigns it to the RegexpMatchIndex field.
func (o *IDExtractorConfig) SetRegexpMatchIndex(v int32) {
	o.RegexpMatchIndex = &v
}

// GetXPathExp returns the XPathExp field value if set, zero value otherwise.
func (o *IDExtractorConfig) GetXPathExp() string {
	if o == nil || IsNil(o.XPathExp) {
		var ret string
		return ret
	}
	return *o.XPathExp
}

// GetXPathExpOk returns a tuple with the XPathExp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IDExtractorConfig) GetXPathExpOk() (*string, bool) {
	if o == nil || IsNil(o.XPathExp) {
		return nil, false
	}
	return o.XPathExp, true
}

// HasXPathExp returns a boolean if a field has been set.
func (o *IDExtractorConfig) HasXPathExp() bool {
	if o != nil && !IsNil(o.XPathExp) {
		return true
	}

	return false
}

// SetXPathExp gets a reference to the given string and assigns it to the XPathExp field.
func (o *IDExtractorConfig) SetXPathExp(v string) {
	o.XPathExp = &v
}

func (o IDExtractorConfig) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IDExtractorConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FormParamName) {
		toSerialize["formParamName"] = o.FormParamName
	}
	if !IsNil(o.HeaderName) {
		toSerialize["headerName"] = o.HeaderName
	}
	if !IsNil(o.Regexp) {
		toSerialize["regexp"] = o.Regexp
	}
	if !IsNil(o.RegexpMatchIndex) {
		toSerialize["regexpMatchIndex"] = o.RegexpMatchIndex
	}
	if !IsNil(o.XPathExp) {
		toSerialize["xPathExp"] = o.XPathExp
	}
	return toSerialize, nil
}

type NullableIDExtractorConfig struct {
	value *IDExtractorConfig
	isSet bool
}

func (v NullableIDExtractorConfig) Get() *IDExtractorConfig {
	return v.value
}

func (v *NullableIDExtractorConfig) Set(val *IDExtractorConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableIDExtractorConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableIDExtractorConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIDExtractorConfig(val *IDExtractorConfig) *NullableIDExtractorConfig {
	return &NullableIDExtractorConfig{value: val, isSet: true}
}

func (v NullableIDExtractorConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIDExtractorConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
