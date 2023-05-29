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

// checks if the CertificateMeta type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CertificateMeta{}

// CertificateMeta CertificateBasics represents basic details of a certificate
type CertificateMeta struct {
	Id *string `json:"id,omitempty"`
	Fingerprint *string `json:"fingerprint,omitempty"`
	HasPrivate *string `json:"has_private,omitempty"`
	Issuer *PkixName `json:"issuer,omitempty"`
	Subject *PkixName `json:"subject,omitempty"`
	NotBefore *string `json:"not_before,omitempty"`
	NotAfter *string `json:"not_after,omitempty"`
	DnsNames []string `json:"dns_names,omitempty"`
}

// NewCertificateMeta instantiates a new CertificateMeta object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCertificateMeta() *CertificateMeta {
	this := CertificateMeta{}
	return &this
}

// NewCertificateMetaWithDefaults instantiates a new CertificateMeta object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCertificateMetaWithDefaults() *CertificateMeta {
	this := CertificateMeta{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CertificateMeta) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateMeta) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CertificateMeta) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CertificateMeta) SetId(v string) {
	o.Id = &v
}

// GetFingerprint returns the Fingerprint field value if set, zero value otherwise.
func (o *CertificateMeta) GetFingerprint() string {
	if o == nil || IsNil(o.Fingerprint) {
		var ret string
		return ret
	}
	return *o.Fingerprint
}

// GetFingerprintOk returns a tuple with the Fingerprint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateMeta) GetFingerprintOk() (*string, bool) {
	if o == nil || IsNil(o.Fingerprint) {
		return nil, false
	}
	return o.Fingerprint, true
}

// HasFingerprint returns a boolean if a field has been set.
func (o *CertificateMeta) HasFingerprint() bool {
	if o != nil && !IsNil(o.Fingerprint) {
		return true
	}

	return false
}

// SetFingerprint gets a reference to the given string and assigns it to the Fingerprint field.
func (o *CertificateMeta) SetFingerprint(v string) {
	o.Fingerprint = &v
}

// GetHasPrivate returns the HasPrivate field value if set, zero value otherwise.
func (o *CertificateMeta) GetHasPrivate() string {
	if o == nil || IsNil(o.HasPrivate) {
		var ret string
		return ret
	}
	return *o.HasPrivate
}

// GetHasPrivateOk returns a tuple with the HasPrivate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateMeta) GetHasPrivateOk() (*string, bool) {
	if o == nil || IsNil(o.HasPrivate) {
		return nil, false
	}
	return o.HasPrivate, true
}

// HasHasPrivate returns a boolean if a field has been set.
func (o *CertificateMeta) HasHasPrivate() bool {
	if o != nil && !IsNil(o.HasPrivate) {
		return true
	}

	return false
}

// SetHasPrivate gets a reference to the given string and assigns it to the HasPrivate field.
func (o *CertificateMeta) SetHasPrivate(v string) {
	o.HasPrivate = &v
}

// GetIssuer returns the Issuer field value if set, zero value otherwise.
func (o *CertificateMeta) GetIssuer() PkixName {
	if o == nil || IsNil(o.Issuer) {
		var ret PkixName
		return ret
	}
	return *o.Issuer
}

// GetIssuerOk returns a tuple with the Issuer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateMeta) GetIssuerOk() (*PkixName, bool) {
	if o == nil || IsNil(o.Issuer) {
		return nil, false
	}
	return o.Issuer, true
}

// HasIssuer returns a boolean if a field has been set.
func (o *CertificateMeta) HasIssuer() bool {
	if o != nil && !IsNil(o.Issuer) {
		return true
	}

	return false
}

// SetIssuer gets a reference to the given PkixName and assigns it to the Issuer field.
func (o *CertificateMeta) SetIssuer(v PkixName) {
	o.Issuer = &v
}

// GetSubject returns the Subject field value if set, zero value otherwise.
func (o *CertificateMeta) GetSubject() PkixName {
	if o == nil || IsNil(o.Subject) {
		var ret PkixName
		return ret
	}
	return *o.Subject
}

// GetSubjectOk returns a tuple with the Subject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateMeta) GetSubjectOk() (*PkixName, bool) {
	if o == nil || IsNil(o.Subject) {
		return nil, false
	}
	return o.Subject, true
}

// HasSubject returns a boolean if a field has been set.
func (o *CertificateMeta) HasSubject() bool {
	if o != nil && !IsNil(o.Subject) {
		return true
	}

	return false
}

// SetSubject gets a reference to the given PkixName and assigns it to the Subject field.
func (o *CertificateMeta) SetSubject(v PkixName) {
	o.Subject = &v
}

// GetNotBefore returns the NotBefore field value if set, zero value otherwise.
func (o *CertificateMeta) GetNotBefore() string {
	if o == nil || IsNil(o.NotBefore) {
		var ret string
		return ret
	}
	return *o.NotBefore
}

// GetNotBeforeOk returns a tuple with the NotBefore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateMeta) GetNotBeforeOk() (*string, bool) {
	if o == nil || IsNil(o.NotBefore) {
		return nil, false
	}
	return o.NotBefore, true
}

// HasNotBefore returns a boolean if a field has been set.
func (o *CertificateMeta) HasNotBefore() bool {
	if o != nil && !IsNil(o.NotBefore) {
		return true
	}

	return false
}

// SetNotBefore gets a reference to the given string and assigns it to the NotBefore field.
func (o *CertificateMeta) SetNotBefore(v string) {
	o.NotBefore = &v
}

// GetNotAfter returns the NotAfter field value if set, zero value otherwise.
func (o *CertificateMeta) GetNotAfter() string {
	if o == nil || IsNil(o.NotAfter) {
		var ret string
		return ret
	}
	return *o.NotAfter
}

// GetNotAfterOk returns a tuple with the NotAfter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateMeta) GetNotAfterOk() (*string, bool) {
	if o == nil || IsNil(o.NotAfter) {
		return nil, false
	}
	return o.NotAfter, true
}

// HasNotAfter returns a boolean if a field has been set.
func (o *CertificateMeta) HasNotAfter() bool {
	if o != nil && !IsNil(o.NotAfter) {
		return true
	}

	return false
}

// SetNotAfter gets a reference to the given string and assigns it to the NotAfter field.
func (o *CertificateMeta) SetNotAfter(v string) {
	o.NotAfter = &v
}

// GetDnsNames returns the DnsNames field value if set, zero value otherwise.
func (o *CertificateMeta) GetDnsNames() []string {
	if o == nil || IsNil(o.DnsNames) {
		var ret []string
		return ret
	}
	return o.DnsNames
}

// GetDnsNamesOk returns a tuple with the DnsNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateMeta) GetDnsNamesOk() ([]string, bool) {
	if o == nil || IsNil(o.DnsNames) {
		return nil, false
	}
	return o.DnsNames, true
}

// HasDnsNames returns a boolean if a field has been set.
func (o *CertificateMeta) HasDnsNames() bool {
	if o != nil && !IsNil(o.DnsNames) {
		return true
	}

	return false
}

// SetDnsNames gets a reference to the given []string and assigns it to the DnsNames field.
func (o *CertificateMeta) SetDnsNames(v []string) {
	o.DnsNames = v
}

func (o CertificateMeta) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CertificateMeta) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Fingerprint) {
		toSerialize["fingerprint"] = o.Fingerprint
	}
	if !IsNil(o.HasPrivate) {
		toSerialize["has_private"] = o.HasPrivate
	}
	if !IsNil(o.Issuer) {
		toSerialize["issuer"] = o.Issuer
	}
	if !IsNil(o.Subject) {
		toSerialize["subject"] = o.Subject
	}
	if !IsNil(o.NotBefore) {
		toSerialize["not_before"] = o.NotBefore
	}
	if !IsNil(o.NotAfter) {
		toSerialize["not_after"] = o.NotAfter
	}
	if !IsNil(o.DnsNames) {
		toSerialize["dns_names"] = o.DnsNames
	}
	return toSerialize, nil
}

type NullableCertificateMeta struct {
	value *CertificateMeta
	isSet bool
}

func (v NullableCertificateMeta) Get() *CertificateMeta {
	return v.value
}

func (v *NullableCertificateMeta) Set(val *CertificateMeta) {
	v.value = val
	v.isSet = true
}

func (v NullableCertificateMeta) IsSet() bool {
	return v.isSet
}

func (v *NullableCertificateMeta) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCertificateMeta(val *CertificateMeta) *NullableCertificateMeta {
	return &NullableCertificateMeta{value: val, isSet: true}
}

func (v NullableCertificateMeta) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCertificateMeta) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


