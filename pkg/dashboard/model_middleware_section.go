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

// checks if the MiddlewareSection type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MiddlewareSection{}

// MiddlewareSection struct for MiddlewareSection
type MiddlewareSection struct {
	AuthCheck *MiddlewareDefinition `json:"auth_check,omitempty"`
	Driver *string `json:"driver,omitempty"`
	IdExtractor *MiddlewareIdExtractor `json:"id_extractor,omitempty"`
	Post []MiddlewareDefinition `json:"post,omitempty"`
	PostKeyAuth []MiddlewareDefinition `json:"post_key_auth,omitempty"`
	Pre []MiddlewareDefinition `json:"pre,omitempty"`
	Response []MiddlewareDefinition `json:"response,omitempty"`
}

// NewMiddlewareSection instantiates a new MiddlewareSection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMiddlewareSection() *MiddlewareSection {
	this := MiddlewareSection{}
	return &this
}

// NewMiddlewareSectionWithDefaults instantiates a new MiddlewareSection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMiddlewareSectionWithDefaults() *MiddlewareSection {
	this := MiddlewareSection{}
	return &this
}

// GetAuthCheck returns the AuthCheck field value if set, zero value otherwise.
func (o *MiddlewareSection) GetAuthCheck() MiddlewareDefinition {
	if o == nil || IsNil(o.AuthCheck) {
		var ret MiddlewareDefinition
		return ret
	}
	return *o.AuthCheck
}

// GetAuthCheckOk returns a tuple with the AuthCheck field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MiddlewareSection) GetAuthCheckOk() (*MiddlewareDefinition, bool) {
	if o == nil || IsNil(o.AuthCheck) {
		return nil, false
	}
	return o.AuthCheck, true
}

// HasAuthCheck returns a boolean if a field has been set.
func (o *MiddlewareSection) HasAuthCheck() bool {
	if o != nil && !IsNil(o.AuthCheck) {
		return true
	}

	return false
}

// SetAuthCheck gets a reference to the given MiddlewareDefinition and assigns it to the AuthCheck field.
func (o *MiddlewareSection) SetAuthCheck(v MiddlewareDefinition) {
	o.AuthCheck = &v
}

// GetDriver returns the Driver field value if set, zero value otherwise.
func (o *MiddlewareSection) GetDriver() string {
	if o == nil || IsNil(o.Driver) {
		var ret string
		return ret
	}
	return *o.Driver
}

// GetDriverOk returns a tuple with the Driver field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MiddlewareSection) GetDriverOk() (*string, bool) {
	if o == nil || IsNil(o.Driver) {
		return nil, false
	}
	return o.Driver, true
}

// HasDriver returns a boolean if a field has been set.
func (o *MiddlewareSection) HasDriver() bool {
	if o != nil && !IsNil(o.Driver) {
		return true
	}

	return false
}

// SetDriver gets a reference to the given string and assigns it to the Driver field.
func (o *MiddlewareSection) SetDriver(v string) {
	o.Driver = &v
}

// GetIdExtractor returns the IdExtractor field value if set, zero value otherwise.
func (o *MiddlewareSection) GetIdExtractor() MiddlewareIdExtractor {
	if o == nil || IsNil(o.IdExtractor) {
		var ret MiddlewareIdExtractor
		return ret
	}
	return *o.IdExtractor
}

// GetIdExtractorOk returns a tuple with the IdExtractor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MiddlewareSection) GetIdExtractorOk() (*MiddlewareIdExtractor, bool) {
	if o == nil || IsNil(o.IdExtractor) {
		return nil, false
	}
	return o.IdExtractor, true
}

// HasIdExtractor returns a boolean if a field has been set.
func (o *MiddlewareSection) HasIdExtractor() bool {
	if o != nil && !IsNil(o.IdExtractor) {
		return true
	}

	return false
}

// SetIdExtractor gets a reference to the given MiddlewareIdExtractor and assigns it to the IdExtractor field.
func (o *MiddlewareSection) SetIdExtractor(v MiddlewareIdExtractor) {
	o.IdExtractor = &v
}

// GetPost returns the Post field value if set, zero value otherwise.
func (o *MiddlewareSection) GetPost() []MiddlewareDefinition {
	if o == nil || IsNil(o.Post) {
		var ret []MiddlewareDefinition
		return ret
	}
	return o.Post
}

// GetPostOk returns a tuple with the Post field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MiddlewareSection) GetPostOk() ([]MiddlewareDefinition, bool) {
	if o == nil || IsNil(o.Post) {
		return nil, false
	}
	return o.Post, true
}

// HasPost returns a boolean if a field has been set.
func (o *MiddlewareSection) HasPost() bool {
	if o != nil && !IsNil(o.Post) {
		return true
	}

	return false
}

// SetPost gets a reference to the given []MiddlewareDefinition and assigns it to the Post field.
func (o *MiddlewareSection) SetPost(v []MiddlewareDefinition) {
	o.Post = v
}

// GetPostKeyAuth returns the PostKeyAuth field value if set, zero value otherwise.
func (o *MiddlewareSection) GetPostKeyAuth() []MiddlewareDefinition {
	if o == nil || IsNil(o.PostKeyAuth) {
		var ret []MiddlewareDefinition
		return ret
	}
	return o.PostKeyAuth
}

// GetPostKeyAuthOk returns a tuple with the PostKeyAuth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MiddlewareSection) GetPostKeyAuthOk() ([]MiddlewareDefinition, bool) {
	if o == nil || IsNil(o.PostKeyAuth) {
		return nil, false
	}
	return o.PostKeyAuth, true
}

// HasPostKeyAuth returns a boolean if a field has been set.
func (o *MiddlewareSection) HasPostKeyAuth() bool {
	if o != nil && !IsNil(o.PostKeyAuth) {
		return true
	}

	return false
}

// SetPostKeyAuth gets a reference to the given []MiddlewareDefinition and assigns it to the PostKeyAuth field.
func (o *MiddlewareSection) SetPostKeyAuth(v []MiddlewareDefinition) {
	o.PostKeyAuth = v
}

// GetPre returns the Pre field value if set, zero value otherwise.
func (o *MiddlewareSection) GetPre() []MiddlewareDefinition {
	if o == nil || IsNil(o.Pre) {
		var ret []MiddlewareDefinition
		return ret
	}
	return o.Pre
}

// GetPreOk returns a tuple with the Pre field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MiddlewareSection) GetPreOk() ([]MiddlewareDefinition, bool) {
	if o == nil || IsNil(o.Pre) {
		return nil, false
	}
	return o.Pre, true
}

// HasPre returns a boolean if a field has been set.
func (o *MiddlewareSection) HasPre() bool {
	if o != nil && !IsNil(o.Pre) {
		return true
	}

	return false
}

// SetPre gets a reference to the given []MiddlewareDefinition and assigns it to the Pre field.
func (o *MiddlewareSection) SetPre(v []MiddlewareDefinition) {
	o.Pre = v
}

// GetResponse returns the Response field value if set, zero value otherwise.
func (o *MiddlewareSection) GetResponse() []MiddlewareDefinition {
	if o == nil || IsNil(o.Response) {
		var ret []MiddlewareDefinition
		return ret
	}
	return o.Response
}

// GetResponseOk returns a tuple with the Response field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MiddlewareSection) GetResponseOk() ([]MiddlewareDefinition, bool) {
	if o == nil || IsNil(o.Response) {
		return nil, false
	}
	return o.Response, true
}

// HasResponse returns a boolean if a field has been set.
func (o *MiddlewareSection) HasResponse() bool {
	if o != nil && !IsNil(o.Response) {
		return true
	}

	return false
}

// SetResponse gets a reference to the given []MiddlewareDefinition and assigns it to the Response field.
func (o *MiddlewareSection) SetResponse(v []MiddlewareDefinition) {
	o.Response = v
}

func (o MiddlewareSection) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MiddlewareSection) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AuthCheck) {
		toSerialize["auth_check"] = o.AuthCheck
	}
	if !IsNil(o.Driver) {
		toSerialize["driver"] = o.Driver
	}
	if !IsNil(o.IdExtractor) {
		toSerialize["id_extractor"] = o.IdExtractor
	}
	if !IsNil(o.Post) {
		toSerialize["post"] = o.Post
	}
	if !IsNil(o.PostKeyAuth) {
		toSerialize["post_key_auth"] = o.PostKeyAuth
	}
	if !IsNil(o.Pre) {
		toSerialize["pre"] = o.Pre
	}
	if !IsNil(o.Response) {
		toSerialize["response"] = o.Response
	}
	return toSerialize, nil
}

type NullableMiddlewareSection struct {
	value *MiddlewareSection
	isSet bool
}

func (v NullableMiddlewareSection) Get() *MiddlewareSection {
	return v.value
}

func (v *NullableMiddlewareSection) Set(val *MiddlewareSection) {
	v.value = val
	v.isSet = true
}

func (v NullableMiddlewareSection) IsSet() bool {
	return v.isSet
}

func (v *NullableMiddlewareSection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMiddlewareSection(val *MiddlewareSection) *NullableMiddlewareSection {
	return &NullableMiddlewareSection{value: val, isSet: true}
}

func (v NullableMiddlewareSection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMiddlewareSection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


