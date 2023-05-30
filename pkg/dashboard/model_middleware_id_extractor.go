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

// checks if the MiddlewareIdExtractor type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MiddlewareIdExtractor{}

// MiddlewareIdExtractor struct for MiddlewareIdExtractor
type MiddlewareIdExtractor struct {
	ExtractFrom     *string                           `json:"extract_from,omitempty"`
	ExtractWith     *string                           `json:"extract_with,omitempty"`
	ExtractorConfig map[string]map[string]interface{} `json:"extractor_config,omitempty"`
}

// NewMiddlewareIdExtractor instantiates a new MiddlewareIdExtractor object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMiddlewareIdExtractor() *MiddlewareIdExtractor {
	this := MiddlewareIdExtractor{}
	return &this
}

// NewMiddlewareIdExtractorWithDefaults instantiates a new MiddlewareIdExtractor object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMiddlewareIdExtractorWithDefaults() *MiddlewareIdExtractor {
	this := MiddlewareIdExtractor{}
	return &this
}

// GetExtractFrom returns the ExtractFrom field value if set, zero value otherwise.
func (o *MiddlewareIdExtractor) GetExtractFrom() string {
	if o == nil || IsNil(o.ExtractFrom) {
		var ret string
		return ret
	}
	return *o.ExtractFrom
}

// GetExtractFromOk returns a tuple with the ExtractFrom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MiddlewareIdExtractor) GetExtractFromOk() (*string, bool) {
	if o == nil || IsNil(o.ExtractFrom) {
		return nil, false
	}
	return o.ExtractFrom, true
}

// HasExtractFrom returns a boolean if a field has been set.
func (o *MiddlewareIdExtractor) HasExtractFrom() bool {
	if o != nil && !IsNil(o.ExtractFrom) {
		return true
	}

	return false
}

// SetExtractFrom gets a reference to the given string and assigns it to the ExtractFrom field.
func (o *MiddlewareIdExtractor) SetExtractFrom(v string) {
	o.ExtractFrom = &v
}

// GetExtractWith returns the ExtractWith field value if set, zero value otherwise.
func (o *MiddlewareIdExtractor) GetExtractWith() string {
	if o == nil || IsNil(o.ExtractWith) {
		var ret string
		return ret
	}
	return *o.ExtractWith
}

// GetExtractWithOk returns a tuple with the ExtractWith field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MiddlewareIdExtractor) GetExtractWithOk() (*string, bool) {
	if o == nil || IsNil(o.ExtractWith) {
		return nil, false
	}
	return o.ExtractWith, true
}

// HasExtractWith returns a boolean if a field has been set.
func (o *MiddlewareIdExtractor) HasExtractWith() bool {
	if o != nil && !IsNil(o.ExtractWith) {
		return true
	}

	return false
}

// SetExtractWith gets a reference to the given string and assigns it to the ExtractWith field.
func (o *MiddlewareIdExtractor) SetExtractWith(v string) {
	o.ExtractWith = &v
}

// GetExtractorConfig returns the ExtractorConfig field value if set, zero value otherwise.
func (o *MiddlewareIdExtractor) GetExtractorConfig() map[string]map[string]interface{} {
	if o == nil || IsNil(o.ExtractorConfig) {
		var ret map[string]map[string]interface{}
		return ret
	}
	return o.ExtractorConfig
}

// GetExtractorConfigOk returns a tuple with the ExtractorConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MiddlewareIdExtractor) GetExtractorConfigOk() (map[string]map[string]interface{}, bool) {
	if o == nil || IsNil(o.ExtractorConfig) {
		return map[string]map[string]interface{}{}, false
	}
	return o.ExtractorConfig, true
}

// HasExtractorConfig returns a boolean if a field has been set.
func (o *MiddlewareIdExtractor) HasExtractorConfig() bool {
	if o != nil && !IsNil(o.ExtractorConfig) {
		return true
	}

	return false
}

// SetExtractorConfig gets a reference to the given map[string]map[string]interface{} and assigns it to the ExtractorConfig field.
func (o *MiddlewareIdExtractor) SetExtractorConfig(v map[string]map[string]interface{}) {
	o.ExtractorConfig = v
}

func (o MiddlewareIdExtractor) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MiddlewareIdExtractor) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ExtractFrom) {
		toSerialize["extract_from"] = o.ExtractFrom
	}
	if !IsNil(o.ExtractWith) {
		toSerialize["extract_with"] = o.ExtractWith
	}
	if !IsNil(o.ExtractorConfig) {
		toSerialize["extractor_config"] = o.ExtractorConfig
	}
	return toSerialize, nil
}

type NullableMiddlewareIdExtractor struct {
	value *MiddlewareIdExtractor
	isSet bool
}

func (v NullableMiddlewareIdExtractor) Get() *MiddlewareIdExtractor {
	return v.value
}

func (v *NullableMiddlewareIdExtractor) Set(val *MiddlewareIdExtractor) {
	v.value = val
	v.isSet = true
}

func (v NullableMiddlewareIdExtractor) IsSet() bool {
	return v.isSet
}

func (v *NullableMiddlewareIdExtractor) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMiddlewareIdExtractor(val *MiddlewareIdExtractor) *NullableMiddlewareIdExtractor {
	return &NullableMiddlewareIdExtractor{value: val, isSet: true}
}

func (v NullableMiddlewareIdExtractor) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMiddlewareIdExtractor) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
