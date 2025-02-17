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

// checks if the CertificateList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CertificateList{}

// CertificateList struct for CertificateList
type CertificateList struct {
	Certs []string `json:"certs,omitempty"`
	Pages *int32   `json:"pages,omitempty"`
}

// NewCertificateList instantiates a new CertificateList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCertificateList() *CertificateList {
	this := CertificateList{}
	return &this
}

// NewCertificateListWithDefaults instantiates a new CertificateList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCertificateListWithDefaults() *CertificateList {
	this := CertificateList{}
	return &this
}

// GetCerts returns the Certs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CertificateList) GetCerts() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Certs
}

// GetCertsOk returns a tuple with the Certs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CertificateList) GetCertsOk() ([]string, bool) {
	if o == nil || IsNil(o.Certs) {
		return nil, false
	}
	return o.Certs, true
}

// HasCerts returns a boolean if a field has been set.
func (o *CertificateList) HasCerts() bool {
	if o != nil && !IsNil(o.Certs) {
		return true
	}

	return false
}

// SetCerts gets a reference to the given []string and assigns it to the Certs field.
func (o *CertificateList) SetCerts(v []string) {
	o.Certs = v
}

// GetPages returns the Pages field value if set, zero value otherwise.
func (o *CertificateList) GetPages() int32 {
	if o == nil || IsNil(o.Pages) {
		var ret int32
		return ret
	}
	return *o.Pages
}

// GetPagesOk returns a tuple with the Pages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateList) GetPagesOk() (*int32, bool) {
	if o == nil || IsNil(o.Pages) {
		return nil, false
	}
	return o.Pages, true
}

// HasPages returns a boolean if a field has been set.
func (o *CertificateList) HasPages() bool {
	if o != nil && !IsNil(o.Pages) {
		return true
	}

	return false
}

// SetPages gets a reference to the given int32 and assigns it to the Pages field.
func (o *CertificateList) SetPages(v int32) {
	o.Pages = &v
}

func (o CertificateList) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CertificateList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Certs != nil {
		toSerialize["certs"] = o.Certs
	}
	if !IsNil(o.Pages) {
		toSerialize["pages"] = o.Pages
	}
	return toSerialize, nil
}

type NullableCertificateList struct {
	value *CertificateList
	isSet bool
}

func (v NullableCertificateList) Get() *CertificateList {
	return v.value
}

func (v *NullableCertificateList) Set(val *CertificateList) {
	v.value = val
	v.isSet = true
}

func (v NullableCertificateList) IsSet() bool {
	return v.isSet
}

func (v *NullableCertificateList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCertificateList(val *CertificateList) *NullableCertificateList {
	return &NullableCertificateList{value: val, isSet: true}
}

func (v NullableCertificateList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCertificateList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
