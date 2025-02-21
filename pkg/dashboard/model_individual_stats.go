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

// checks if the IndividualStats type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IndividualStats{}

// IndividualStats struct for IndividualStats
type IndividualStats struct {
	AvgActiveUsage           *int32  `json:"avg_active_usage,omitempty"`
	AvgUsage                 *int32  `json:"avg_usage,omitempty"`
	Date                     *string `json:"date,omitempty"`
	LicenseEntitlement       *int32  `json:"license_entitlement,omitempty"`
	LicenseEntitlementActive *int32  `json:"license_entitlement_active,omitempty"`
	MaxActiveUsage           *int32  `json:"max_active_usage,omitempty"`
	MaxUsage                 *int32  `json:"max_usage,omitempty"`
	MinActiveUsage           *int32  `json:"min_active_usage,omitempty"`
	MinUsage                 *int32  `json:"min_usage,omitempty"`
}

// NewIndividualStats instantiates a new IndividualStats object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIndividualStats() *IndividualStats {
	this := IndividualStats{}
	return &this
}

// NewIndividualStatsWithDefaults instantiates a new IndividualStats object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIndividualStatsWithDefaults() *IndividualStats {
	this := IndividualStats{}
	return &this
}

// GetAvgActiveUsage returns the AvgActiveUsage field value if set, zero value otherwise.
func (o *IndividualStats) GetAvgActiveUsage() int32 {
	if o == nil || IsNil(o.AvgActiveUsage) {
		var ret int32
		return ret
	}
	return *o.AvgActiveUsage
}

// GetAvgActiveUsageOk returns a tuple with the AvgActiveUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndividualStats) GetAvgActiveUsageOk() (*int32, bool) {
	if o == nil || IsNil(o.AvgActiveUsage) {
		return nil, false
	}
	return o.AvgActiveUsage, true
}

// HasAvgActiveUsage returns a boolean if a field has been set.
func (o *IndividualStats) HasAvgActiveUsage() bool {
	if o != nil && !IsNil(o.AvgActiveUsage) {
		return true
	}

	return false
}

// SetAvgActiveUsage gets a reference to the given int32 and assigns it to the AvgActiveUsage field.
func (o *IndividualStats) SetAvgActiveUsage(v int32) {
	o.AvgActiveUsage = &v
}

// GetAvgUsage returns the AvgUsage field value if set, zero value otherwise.
func (o *IndividualStats) GetAvgUsage() int32 {
	if o == nil || IsNil(o.AvgUsage) {
		var ret int32
		return ret
	}
	return *o.AvgUsage
}

// GetAvgUsageOk returns a tuple with the AvgUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndividualStats) GetAvgUsageOk() (*int32, bool) {
	if o == nil || IsNil(o.AvgUsage) {
		return nil, false
	}
	return o.AvgUsage, true
}

// HasAvgUsage returns a boolean if a field has been set.
func (o *IndividualStats) HasAvgUsage() bool {
	if o != nil && !IsNil(o.AvgUsage) {
		return true
	}

	return false
}

// SetAvgUsage gets a reference to the given int32 and assigns it to the AvgUsage field.
func (o *IndividualStats) SetAvgUsage(v int32) {
	o.AvgUsage = &v
}

// GetDate returns the Date field value if set, zero value otherwise.
func (o *IndividualStats) GetDate() string {
	if o == nil || IsNil(o.Date) {
		var ret string
		return ret
	}
	return *o.Date
}

// GetDateOk returns a tuple with the Date field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndividualStats) GetDateOk() (*string, bool) {
	if o == nil || IsNil(o.Date) {
		return nil, false
	}
	return o.Date, true
}

// HasDate returns a boolean if a field has been set.
func (o *IndividualStats) HasDate() bool {
	if o != nil && !IsNil(o.Date) {
		return true
	}

	return false
}

// SetDate gets a reference to the given string and assigns it to the Date field.
func (o *IndividualStats) SetDate(v string) {
	o.Date = &v
}

// GetLicenseEntitlement returns the LicenseEntitlement field value if set, zero value otherwise.
func (o *IndividualStats) GetLicenseEntitlement() int32 {
	if o == nil || IsNil(o.LicenseEntitlement) {
		var ret int32
		return ret
	}
	return *o.LicenseEntitlement
}

// GetLicenseEntitlementOk returns a tuple with the LicenseEntitlement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndividualStats) GetLicenseEntitlementOk() (*int32, bool) {
	if o == nil || IsNil(o.LicenseEntitlement) {
		return nil, false
	}
	return o.LicenseEntitlement, true
}

// HasLicenseEntitlement returns a boolean if a field has been set.
func (o *IndividualStats) HasLicenseEntitlement() bool {
	if o != nil && !IsNil(o.LicenseEntitlement) {
		return true
	}

	return false
}

// SetLicenseEntitlement gets a reference to the given int32 and assigns it to the LicenseEntitlement field.
func (o *IndividualStats) SetLicenseEntitlement(v int32) {
	o.LicenseEntitlement = &v
}

// GetLicenseEntitlementActive returns the LicenseEntitlementActive field value if set, zero value otherwise.
func (o *IndividualStats) GetLicenseEntitlementActive() int32 {
	if o == nil || IsNil(o.LicenseEntitlementActive) {
		var ret int32
		return ret
	}
	return *o.LicenseEntitlementActive
}

// GetLicenseEntitlementActiveOk returns a tuple with the LicenseEntitlementActive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndividualStats) GetLicenseEntitlementActiveOk() (*int32, bool) {
	if o == nil || IsNil(o.LicenseEntitlementActive) {
		return nil, false
	}
	return o.LicenseEntitlementActive, true
}

// HasLicenseEntitlementActive returns a boolean if a field has been set.
func (o *IndividualStats) HasLicenseEntitlementActive() bool {
	if o != nil && !IsNil(o.LicenseEntitlementActive) {
		return true
	}

	return false
}

// SetLicenseEntitlementActive gets a reference to the given int32 and assigns it to the LicenseEntitlementActive field.
func (o *IndividualStats) SetLicenseEntitlementActive(v int32) {
	o.LicenseEntitlementActive = &v
}

// GetMaxActiveUsage returns the MaxActiveUsage field value if set, zero value otherwise.
func (o *IndividualStats) GetMaxActiveUsage() int32 {
	if o == nil || IsNil(o.MaxActiveUsage) {
		var ret int32
		return ret
	}
	return *o.MaxActiveUsage
}

// GetMaxActiveUsageOk returns a tuple with the MaxActiveUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndividualStats) GetMaxActiveUsageOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxActiveUsage) {
		return nil, false
	}
	return o.MaxActiveUsage, true
}

// HasMaxActiveUsage returns a boolean if a field has been set.
func (o *IndividualStats) HasMaxActiveUsage() bool {
	if o != nil && !IsNil(o.MaxActiveUsage) {
		return true
	}

	return false
}

// SetMaxActiveUsage gets a reference to the given int32 and assigns it to the MaxActiveUsage field.
func (o *IndividualStats) SetMaxActiveUsage(v int32) {
	o.MaxActiveUsage = &v
}

// GetMaxUsage returns the MaxUsage field value if set, zero value otherwise.
func (o *IndividualStats) GetMaxUsage() int32 {
	if o == nil || IsNil(o.MaxUsage) {
		var ret int32
		return ret
	}
	return *o.MaxUsage
}

// GetMaxUsageOk returns a tuple with the MaxUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndividualStats) GetMaxUsageOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxUsage) {
		return nil, false
	}
	return o.MaxUsage, true
}

// HasMaxUsage returns a boolean if a field has been set.
func (o *IndividualStats) HasMaxUsage() bool {
	if o != nil && !IsNil(o.MaxUsage) {
		return true
	}

	return false
}

// SetMaxUsage gets a reference to the given int32 and assigns it to the MaxUsage field.
func (o *IndividualStats) SetMaxUsage(v int32) {
	o.MaxUsage = &v
}

// GetMinActiveUsage returns the MinActiveUsage field value if set, zero value otherwise.
func (o *IndividualStats) GetMinActiveUsage() int32 {
	if o == nil || IsNil(o.MinActiveUsage) {
		var ret int32
		return ret
	}
	return *o.MinActiveUsage
}

// GetMinActiveUsageOk returns a tuple with the MinActiveUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndividualStats) GetMinActiveUsageOk() (*int32, bool) {
	if o == nil || IsNil(o.MinActiveUsage) {
		return nil, false
	}
	return o.MinActiveUsage, true
}

// HasMinActiveUsage returns a boolean if a field has been set.
func (o *IndividualStats) HasMinActiveUsage() bool {
	if o != nil && !IsNil(o.MinActiveUsage) {
		return true
	}

	return false
}

// SetMinActiveUsage gets a reference to the given int32 and assigns it to the MinActiveUsage field.
func (o *IndividualStats) SetMinActiveUsage(v int32) {
	o.MinActiveUsage = &v
}

// GetMinUsage returns the MinUsage field value if set, zero value otherwise.
func (o *IndividualStats) GetMinUsage() int32 {
	if o == nil || IsNil(o.MinUsage) {
		var ret int32
		return ret
	}
	return *o.MinUsage
}

// GetMinUsageOk returns a tuple with the MinUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndividualStats) GetMinUsageOk() (*int32, bool) {
	if o == nil || IsNil(o.MinUsage) {
		return nil, false
	}
	return o.MinUsage, true
}

// HasMinUsage returns a boolean if a field has been set.
func (o *IndividualStats) HasMinUsage() bool {
	if o != nil && !IsNil(o.MinUsage) {
		return true
	}

	return false
}

// SetMinUsage gets a reference to the given int32 and assigns it to the MinUsage field.
func (o *IndividualStats) SetMinUsage(v int32) {
	o.MinUsage = &v
}

func (o IndividualStats) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IndividualStats) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AvgActiveUsage) {
		toSerialize["avg_active_usage"] = o.AvgActiveUsage
	}
	if !IsNil(o.AvgUsage) {
		toSerialize["avg_usage"] = o.AvgUsage
	}
	if !IsNil(o.Date) {
		toSerialize["date"] = o.Date
	}
	if !IsNil(o.LicenseEntitlement) {
		toSerialize["license_entitlement"] = o.LicenseEntitlement
	}
	if !IsNil(o.LicenseEntitlementActive) {
		toSerialize["license_entitlement_active"] = o.LicenseEntitlementActive
	}
	if !IsNil(o.MaxActiveUsage) {
		toSerialize["max_active_usage"] = o.MaxActiveUsage
	}
	if !IsNil(o.MaxUsage) {
		toSerialize["max_usage"] = o.MaxUsage
	}
	if !IsNil(o.MinActiveUsage) {
		toSerialize["min_active_usage"] = o.MinActiveUsage
	}
	if !IsNil(o.MinUsage) {
		toSerialize["min_usage"] = o.MinUsage
	}
	return toSerialize, nil
}

type NullableIndividualStats struct {
	value *IndividualStats
	isSet bool
}

func (v NullableIndividualStats) Get() *IndividualStats {
	return v.value
}

func (v *NullableIndividualStats) Set(val *IndividualStats) {
	v.value = val
	v.isSet = true
}

func (v NullableIndividualStats) IsSet() bool {
	return v.isSet
}

func (v *NullableIndividualStats) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIndividualStats(val *IndividualStats) *NullableIndividualStats {
	return &NullableIndividualStats{value: val, isSet: true}
}

func (v NullableIndividualStats) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIndividualStats) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
