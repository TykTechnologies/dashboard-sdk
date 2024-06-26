/*
NEW Tyk DASH API

## <a name=\"introduction\"></a> Introduction      The Tyk Dashboard API offers granular, programmatic access to a centralised database of resources that your Tyk nodes can pull from. This API has a dynamic user administrative structure which means the secret key that is used to communicate with your Tyk nodes can be kept secret and access to the wider management functions can be handled on a user-by-user and organisation-by-organisation basis.      A common question around using a database-backed configuration is how to programatically add API definitions to your Tyk nodes, the Dashboard API allows much more fine-grained, secure and multi-user access to your Tyk cluster, and should be used to manage a database-backed Tyk node.      The Tyk Dashboard API works seamlessly with the Tyk Dashboard (and the two come bundled together).      ## <a name=\"security-hierarchy\"></a> Security Hierarchy      The Dashboard API provides a more structured security layer to managing Tyk nodes.      ### Organisations, APIs and Users      With the Dashboard API and a database-backed Tyk setup, (and to an extent with file-based API setups - if diligence is used in naming an creating definitions), the following security model is applied to the management of Upstream APIs:      * **Organisations**: All APIs are *owned* by an organisation, this is designated by the OrgID parameter in the API Definition.     * **Users**: All users created in the Dashboard belong to an organisation (unless an exception is made for super-administrative access).     * **APIs**: All APIs belong to an Organisation and only Users that belong to that organisation can see the analytics for those APIs and manage their configurations.     * **API Keys**: API Keys are designated by organisation, this means an API key that has full access rights will not be allowed to access the APIs of another organisation on the same system, but can have full access to all APIs within the organisation.     * **Access Rights**: Access rights are stored with the key, this enables a key to give access to multiple APIs, this is defined by the session object in the core Tyk API.      In order to use the Dashboard API, you'll need to get the Tyk Dashboard API Access Credentials secret from your user profile on the Dashboard UI.      The secret you set should then be sent along as a header with each Dashboard API Request in order for it to be successful:  authorization: <your-secret>

API version: 5.4.0
Contact: support@tyk.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dashboard

import (
	"encoding/json"
	"time"
)

// checks if the Policy type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Policy{}

// Policy struct for Policy
type Policy struct {
	Id                 *string                     `json:"_id,omitempty"`
	AccessRights       map[string]AccessDefinition `json:"access_rights,omitempty"`
	Active             *bool                       `json:"active,omitempty"`
	DateCreated        *time.Time                  `json:"date_created,omitempty"`
	HmacEnabled        *bool                       `json:"hmac_enabled,omitempty"`
	Id                 *string                     `json:"id,omitempty"`
	IsInactive         *bool                       `json:"is_inactive,omitempty"`
	KeyExpiresIn       *int32                      `json:"key_expires_in,omitempty"`
	LastUpdated        *string                     `json:"last_updated,omitempty"`
	MaxQueryDepth      *int32                      `json:"max_query_depth,omitempty"`
	MetaData           map[string]interface{}      `json:"meta_data,omitempty"`
	Name               *string                     `json:"name,omitempty"`
	OrgId              *string                     `json:"org_id,omitempty"`
	Partitions         *PolicyPartitions           `json:"partitions,omitempty"`
	Per                *float32                    `json:"per,omitempty"`
	QuotaMax           *int32                      `json:"quota_max,omitempty"`
	QuotaRenewalRate   *int32                      `json:"quota_renewal_rate,omitempty"`
	Rate               *float32                    `json:"rate,omitempty"`
	Smoothing          *RateLimitSmoothing         `json:"smoothing,omitempty"`
	Tags               []string                    `json:"tags,omitempty"`
	ThrottleInterval   *float32                    `json:"throttle_interval,omitempty"`
	ThrottleRetryLimit *int32                      `json:"throttle_retry_limit,omitempty"`
}

// NewPolicy instantiates a new Policy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPolicy() *Policy {
	this := Policy{}
	return &this
}

// NewPolicyWithDefaults instantiates a new Policy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPolicyWithDefaults() *Policy {
	this := Policy{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Policy) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Policy) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Policy) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Policy) SetId(v string) {
	o.Id = &v
}

// GetAccessRights returns the AccessRights field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Policy) GetAccessRights() map[string]AccessDefinition {
	if o == nil {
		var ret map[string]AccessDefinition
		return ret
	}
	return o.AccessRights
}

// GetAccessRightsOk returns a tuple with the AccessRights field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Policy) GetAccessRightsOk() (*map[string]AccessDefinition, bool) {
	if o == nil || IsNil(o.AccessRights) {
		return nil, false
	}
	return &o.AccessRights, true
}

// HasAccessRights returns a boolean if a field has been set.
func (o *Policy) HasAccessRights() bool {
	if o != nil && !IsNil(o.AccessRights) {
		return true
	}

	return false
}

// SetAccessRights gets a reference to the given map[string]AccessDefinition and assigns it to the AccessRights field.
func (o *Policy) SetAccessRights(v map[string]AccessDefinition) {
	o.AccessRights = v
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *Policy) GetActive() bool {
	if o == nil || IsNil(o.Active) {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Policy) GetActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.Active) {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *Policy) HasActive() bool {
	if o != nil && !IsNil(o.Active) {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *Policy) SetActive(v bool) {
	o.Active = &v
}

// GetDateCreated returns the DateCreated field value if set, zero value otherwise.
func (o *Policy) GetDateCreated() time.Time {
	if o == nil || IsNil(o.DateCreated) {
		var ret time.Time
		return ret
	}
	return *o.DateCreated
}

// GetDateCreatedOk returns a tuple with the DateCreated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Policy) GetDateCreatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.DateCreated) {
		return nil, false
	}
	return o.DateCreated, true
}

// HasDateCreated returns a boolean if a field has been set.
func (o *Policy) HasDateCreated() bool {
	if o != nil && !IsNil(o.DateCreated) {
		return true
	}

	return false
}

// SetDateCreated gets a reference to the given time.Time and assigns it to the DateCreated field.
func (o *Policy) SetDateCreated(v time.Time) {
	o.DateCreated = &v
}

// GetHmacEnabled returns the HmacEnabled field value if set, zero value otherwise.
func (o *Policy) GetHmacEnabled() bool {
	if o == nil || IsNil(o.HmacEnabled) {
		var ret bool
		return ret
	}
	return *o.HmacEnabled
}

// GetHmacEnabledOk returns a tuple with the HmacEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Policy) GetHmacEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.HmacEnabled) {
		return nil, false
	}
	return o.HmacEnabled, true
}

// HasHmacEnabled returns a boolean if a field has been set.
func (o *Policy) HasHmacEnabled() bool {
	if o != nil && !IsNil(o.HmacEnabled) {
		return true
	}

	return false
}

// SetHmacEnabled gets a reference to the given bool and assigns it to the HmacEnabled field.
func (o *Policy) SetHmacEnabled(v bool) {
	o.HmacEnabled = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Policy) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Policy) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Policy) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Policy) SetId(v string) {
	o.Id = &v
}

// GetIsInactive returns the IsInactive field value if set, zero value otherwise.
func (o *Policy) GetIsInactive() bool {
	if o == nil || IsNil(o.IsInactive) {
		var ret bool
		return ret
	}
	return *o.IsInactive
}

// GetIsInactiveOk returns a tuple with the IsInactive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Policy) GetIsInactiveOk() (*bool, bool) {
	if o == nil || IsNil(o.IsInactive) {
		return nil, false
	}
	return o.IsInactive, true
}

// HasIsInactive returns a boolean if a field has been set.
func (o *Policy) HasIsInactive() bool {
	if o != nil && !IsNil(o.IsInactive) {
		return true
	}

	return false
}

// SetIsInactive gets a reference to the given bool and assigns it to the IsInactive field.
func (o *Policy) SetIsInactive(v bool) {
	o.IsInactive = &v
}

// GetKeyExpiresIn returns the KeyExpiresIn field value if set, zero value otherwise.
func (o *Policy) GetKeyExpiresIn() int32 {
	if o == nil || IsNil(o.KeyExpiresIn) {
		var ret int32
		return ret
	}
	return *o.KeyExpiresIn
}

// GetKeyExpiresInOk returns a tuple with the KeyExpiresIn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Policy) GetKeyExpiresInOk() (*int32, bool) {
	if o == nil || IsNil(o.KeyExpiresIn) {
		return nil, false
	}
	return o.KeyExpiresIn, true
}

// HasKeyExpiresIn returns a boolean if a field has been set.
func (o *Policy) HasKeyExpiresIn() bool {
	if o != nil && !IsNil(o.KeyExpiresIn) {
		return true
	}

	return false
}

// SetKeyExpiresIn gets a reference to the given int32 and assigns it to the KeyExpiresIn field.
func (o *Policy) SetKeyExpiresIn(v int32) {
	o.KeyExpiresIn = &v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *Policy) GetLastUpdated() string {
	if o == nil || IsNil(o.LastUpdated) {
		var ret string
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Policy) GetLastUpdatedOk() (*string, bool) {
	if o == nil || IsNil(o.LastUpdated) {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *Policy) HasLastUpdated() bool {
	if o != nil && !IsNil(o.LastUpdated) {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given string and assigns it to the LastUpdated field.
func (o *Policy) SetLastUpdated(v string) {
	o.LastUpdated = &v
}

// GetMaxQueryDepth returns the MaxQueryDepth field value if set, zero value otherwise.
func (o *Policy) GetMaxQueryDepth() int32 {
	if o == nil || IsNil(o.MaxQueryDepth) {
		var ret int32
		return ret
	}
	return *o.MaxQueryDepth
}

// GetMaxQueryDepthOk returns a tuple with the MaxQueryDepth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Policy) GetMaxQueryDepthOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxQueryDepth) {
		return nil, false
	}
	return o.MaxQueryDepth, true
}

// HasMaxQueryDepth returns a boolean if a field has been set.
func (o *Policy) HasMaxQueryDepth() bool {
	if o != nil && !IsNil(o.MaxQueryDepth) {
		return true
	}

	return false
}

// SetMaxQueryDepth gets a reference to the given int32 and assigns it to the MaxQueryDepth field.
func (o *Policy) SetMaxQueryDepth(v int32) {
	o.MaxQueryDepth = &v
}

// GetMetaData returns the MetaData field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Policy) GetMetaData() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.MetaData
}

// GetMetaDataOk returns a tuple with the MetaData field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Policy) GetMetaDataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.MetaData) {
		return map[string]interface{}{}, false
	}
	return o.MetaData, true
}

// HasMetaData returns a boolean if a field has been set.
func (o *Policy) HasMetaData() bool {
	if o != nil && !IsNil(o.MetaData) {
		return true
	}

	return false
}

// SetMetaData gets a reference to the given map[string]interface{} and assigns it to the MetaData field.
func (o *Policy) SetMetaData(v map[string]interface{}) {
	o.MetaData = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Policy) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Policy) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Policy) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Policy) SetName(v string) {
	o.Name = &v
}

// GetOrgId returns the OrgId field value if set, zero value otherwise.
func (o *Policy) GetOrgId() string {
	if o == nil || IsNil(o.OrgId) {
		var ret string
		return ret
	}
	return *o.OrgId
}

// GetOrgIdOk returns a tuple with the OrgId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Policy) GetOrgIdOk() (*string, bool) {
	if o == nil || IsNil(o.OrgId) {
		return nil, false
	}
	return o.OrgId, true
}

// HasOrgId returns a boolean if a field has been set.
func (o *Policy) HasOrgId() bool {
	if o != nil && !IsNil(o.OrgId) {
		return true
	}

	return false
}

// SetOrgId gets a reference to the given string and assigns it to the OrgId field.
func (o *Policy) SetOrgId(v string) {
	o.OrgId = &v
}

// GetPartitions returns the Partitions field value if set, zero value otherwise.
func (o *Policy) GetPartitions() PolicyPartitions {
	if o == nil || IsNil(o.Partitions) {
		var ret PolicyPartitions
		return ret
	}
	return *o.Partitions
}

// GetPartitionsOk returns a tuple with the Partitions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Policy) GetPartitionsOk() (*PolicyPartitions, bool) {
	if o == nil || IsNil(o.Partitions) {
		return nil, false
	}
	return o.Partitions, true
}

// HasPartitions returns a boolean if a field has been set.
func (o *Policy) HasPartitions() bool {
	if o != nil && !IsNil(o.Partitions) {
		return true
	}

	return false
}

// SetPartitions gets a reference to the given PolicyPartitions and assigns it to the Partitions field.
func (o *Policy) SetPartitions(v PolicyPartitions) {
	o.Partitions = &v
}

// GetPer returns the Per field value if set, zero value otherwise.
func (o *Policy) GetPer() float32 {
	if o == nil || IsNil(o.Per) {
		var ret float32
		return ret
	}
	return *o.Per
}

// GetPerOk returns a tuple with the Per field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Policy) GetPerOk() (*float32, bool) {
	if o == nil || IsNil(o.Per) {
		return nil, false
	}
	return o.Per, true
}

// HasPer returns a boolean if a field has been set.
func (o *Policy) HasPer() bool {
	if o != nil && !IsNil(o.Per) {
		return true
	}

	return false
}

// SetPer gets a reference to the given float32 and assigns it to the Per field.
func (o *Policy) SetPer(v float32) {
	o.Per = &v
}

// GetQuotaMax returns the QuotaMax field value if set, zero value otherwise.
func (o *Policy) GetQuotaMax() int32 {
	if o == nil || IsNil(o.QuotaMax) {
		var ret int32
		return ret
	}
	return *o.QuotaMax
}

// GetQuotaMaxOk returns a tuple with the QuotaMax field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Policy) GetQuotaMaxOk() (*int32, bool) {
	if o == nil || IsNil(o.QuotaMax) {
		return nil, false
	}
	return o.QuotaMax, true
}

// HasQuotaMax returns a boolean if a field has been set.
func (o *Policy) HasQuotaMax() bool {
	if o != nil && !IsNil(o.QuotaMax) {
		return true
	}

	return false
}

// SetQuotaMax gets a reference to the given int32 and assigns it to the QuotaMax field.
func (o *Policy) SetQuotaMax(v int32) {
	o.QuotaMax = &v
}

// GetQuotaRenewalRate returns the QuotaRenewalRate field value if set, zero value otherwise.
func (o *Policy) GetQuotaRenewalRate() int32 {
	if o == nil || IsNil(o.QuotaRenewalRate) {
		var ret int32
		return ret
	}
	return *o.QuotaRenewalRate
}

// GetQuotaRenewalRateOk returns a tuple with the QuotaRenewalRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Policy) GetQuotaRenewalRateOk() (*int32, bool) {
	if o == nil || IsNil(o.QuotaRenewalRate) {
		return nil, false
	}
	return o.QuotaRenewalRate, true
}

// HasQuotaRenewalRate returns a boolean if a field has been set.
func (o *Policy) HasQuotaRenewalRate() bool {
	if o != nil && !IsNil(o.QuotaRenewalRate) {
		return true
	}

	return false
}

// SetQuotaRenewalRate gets a reference to the given int32 and assigns it to the QuotaRenewalRate field.
func (o *Policy) SetQuotaRenewalRate(v int32) {
	o.QuotaRenewalRate = &v
}

// GetRate returns the Rate field value if set, zero value otherwise.
func (o *Policy) GetRate() float32 {
	if o == nil || IsNil(o.Rate) {
		var ret float32
		return ret
	}
	return *o.Rate
}

// GetRateOk returns a tuple with the Rate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Policy) GetRateOk() (*float32, bool) {
	if o == nil || IsNil(o.Rate) {
		return nil, false
	}
	return o.Rate, true
}

// HasRate returns a boolean if a field has been set.
func (o *Policy) HasRate() bool {
	if o != nil && !IsNil(o.Rate) {
		return true
	}

	return false
}

// SetRate gets a reference to the given float32 and assigns it to the Rate field.
func (o *Policy) SetRate(v float32) {
	o.Rate = &v
}

// GetSmoothing returns the Smoothing field value if set, zero value otherwise.
func (o *Policy) GetSmoothing() RateLimitSmoothing {
	if o == nil || IsNil(o.Smoothing) {
		var ret RateLimitSmoothing
		return ret
	}
	return *o.Smoothing
}

// GetSmoothingOk returns a tuple with the Smoothing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Policy) GetSmoothingOk() (*RateLimitSmoothing, bool) {
	if o == nil || IsNil(o.Smoothing) {
		return nil, false
	}
	return o.Smoothing, true
}

// HasSmoothing returns a boolean if a field has been set.
func (o *Policy) HasSmoothing() bool {
	if o != nil && !IsNil(o.Smoothing) {
		return true
	}

	return false
}

// SetSmoothing gets a reference to the given RateLimitSmoothing and assigns it to the Smoothing field.
func (o *Policy) SetSmoothing(v RateLimitSmoothing) {
	o.Smoothing = &v
}

// GetTags returns the Tags field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Policy) GetTags() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Policy) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *Policy) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *Policy) SetTags(v []string) {
	o.Tags = v
}

// GetThrottleInterval returns the ThrottleInterval field value if set, zero value otherwise.
func (o *Policy) GetThrottleInterval() float32 {
	if o == nil || IsNil(o.ThrottleInterval) {
		var ret float32
		return ret
	}
	return *o.ThrottleInterval
}

// GetThrottleIntervalOk returns a tuple with the ThrottleInterval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Policy) GetThrottleIntervalOk() (*float32, bool) {
	if o == nil || IsNil(o.ThrottleInterval) {
		return nil, false
	}
	return o.ThrottleInterval, true
}

// HasThrottleInterval returns a boolean if a field has been set.
func (o *Policy) HasThrottleInterval() bool {
	if o != nil && !IsNil(o.ThrottleInterval) {
		return true
	}

	return false
}

// SetThrottleInterval gets a reference to the given float32 and assigns it to the ThrottleInterval field.
func (o *Policy) SetThrottleInterval(v float32) {
	o.ThrottleInterval = &v
}

// GetThrottleRetryLimit returns the ThrottleRetryLimit field value if set, zero value otherwise.
func (o *Policy) GetThrottleRetryLimit() int32 {
	if o == nil || IsNil(o.ThrottleRetryLimit) {
		var ret int32
		return ret
	}
	return *o.ThrottleRetryLimit
}

// GetThrottleRetryLimitOk returns a tuple with the ThrottleRetryLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Policy) GetThrottleRetryLimitOk() (*int32, bool) {
	if o == nil || IsNil(o.ThrottleRetryLimit) {
		return nil, false
	}
	return o.ThrottleRetryLimit, true
}

// HasThrottleRetryLimit returns a boolean if a field has been set.
func (o *Policy) HasThrottleRetryLimit() bool {
	if o != nil && !IsNil(o.ThrottleRetryLimit) {
		return true
	}

	return false
}

// SetThrottleRetryLimit gets a reference to the given int32 and assigns it to the ThrottleRetryLimit field.
func (o *Policy) SetThrottleRetryLimit(v int32) {
	o.ThrottleRetryLimit = &v
}

func (o Policy) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Policy) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["_id"] = o.Id
	}
	if o.AccessRights != nil {
		toSerialize["access_rights"] = o.AccessRights
	}
	if !IsNil(o.Active) {
		toSerialize["active"] = o.Active
	}
	if !IsNil(o.DateCreated) {
		toSerialize["date_created"] = o.DateCreated
	}
	if !IsNil(o.HmacEnabled) {
		toSerialize["hmac_enabled"] = o.HmacEnabled
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.IsInactive) {
		toSerialize["is_inactive"] = o.IsInactive
	}
	if !IsNil(o.KeyExpiresIn) {
		toSerialize["key_expires_in"] = o.KeyExpiresIn
	}
	if !IsNil(o.LastUpdated) {
		toSerialize["last_updated"] = o.LastUpdated
	}
	if !IsNil(o.MaxQueryDepth) {
		toSerialize["max_query_depth"] = o.MaxQueryDepth
	}
	if o.MetaData != nil {
		toSerialize["meta_data"] = o.MetaData
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.OrgId) {
		toSerialize["org_id"] = o.OrgId
	}
	if !IsNil(o.Partitions) {
		toSerialize["partitions"] = o.Partitions
	}
	if !IsNil(o.Per) {
		toSerialize["per"] = o.Per
	}
	if !IsNil(o.QuotaMax) {
		toSerialize["quota_max"] = o.QuotaMax
	}
	if !IsNil(o.QuotaRenewalRate) {
		toSerialize["quota_renewal_rate"] = o.QuotaRenewalRate
	}
	if !IsNil(o.Rate) {
		toSerialize["rate"] = o.Rate
	}
	if !IsNil(o.Smoothing) {
		toSerialize["smoothing"] = o.Smoothing
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.ThrottleInterval) {
		toSerialize["throttle_interval"] = o.ThrottleInterval
	}
	if !IsNil(o.ThrottleRetryLimit) {
		toSerialize["throttle_retry_limit"] = o.ThrottleRetryLimit
	}
	return toSerialize, nil
}

type NullablePolicy struct {
	value *Policy
	isSet bool
}

func (v NullablePolicy) Get() *Policy {
	return v.value
}

func (v *NullablePolicy) Set(val *Policy) {
	v.value = val
	v.isSet = true
}

func (v NullablePolicy) IsSet() bool {
	return v.isSet
}

func (v *NullablePolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePolicy(val *Policy) *NullablePolicy {
	return &NullablePolicy{value: val, isSet: true}
}

func (v NullablePolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
