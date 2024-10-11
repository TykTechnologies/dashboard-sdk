/*
Tyk Dashboard API

 ## <a name=\"introduction\"></a> Introduction  The Tyk Dashboard API offers granular, programmatic access to a centralised database of resources that your Tyk nodes can pull from. This API has a dynamic user administrative structure which means the secret key that is used to communicate with your Tyk nodes can be kept secret and access to the wider management functions can be handled on a user-by-user and organisation-by-organisation basis.  A common question around using a database-backed configuration is how to programmatically add API definitions to your Tyk nodes, the Dashboard API allows much more fine-grained, secure and multi-user access to your Tyk cluster, and should be used to manage a database-backed Tyk node.  The Tyk Dashboard API works seamlessly with the Tyk Dashboard (and the two come bundled together).  ## <a name=\"security-hierarchy\"></a> Security Hierarchy  The Dashboard API provides a more structured security layer to managing Tyk nodes.  ### Organisations, APIs and Users  With the Dashboard API and a database-backed Tyk setup, (and to an extent with file-based API setups - if diligence is used in naming and creating definitions), the following security model is applied to the management of Upstream APIs:  * **Organisations**: All APIs are *owned* by an organisation, this is designated by the 'OrgID' parameter in the API Definition. * **Users**: All users created in the Dashboard belong to an organisation (unless an exception is made for super-administrative access). * **APIs**: All APIs belong to an Organisation and only Users that belong to that organisation can see the analytics for those APIs and manage their configurations. * **API Keys**: API Keys are designated by organisation, this means an API key that has full access rights will not be allowed to access the APIs of another organisation on the same system, but can have full access to all APIs within the organisation. * **Access Rights**: Access rights are stored with the key, this enables a key to give access to multiple APIs, this is defined by the session object in the core Tyk API.  In order to use the Dashboard API, you'll need to get the 'Tyk Dashboard API Access Credentials' secret from your user profile on the Dashboard UI.  The secret you set should then be sent along as a header with each Dashboard API Request in order for it to be successful:   authorization: <your-secret>

API version: 5.6.0
Contact: support@tyk.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dashboard

import (
	"encoding/json"
)

// checks if the ResultId type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResultId{}

// ResultId struct for ResultId
type ResultId struct {
	Alias      *string           `json:"alias,omitempty"`
	ApiId      *string           `json:"api_id,omitempty"`
	ApiName    *string           `json:"api_name,omitempty"`
	Code       *int32            `json:"code,omitempty"`
	Day        *int32            `json:"day,omitempty"`
	Hour       *int32            `json:"hour,omitempty"`
	IsoCountry *string           `json:"iso_country,omitempty"`
	Key        *string           `json:"key,omitempty"`
	Month      *int32            `json:"month,omitempty"`
	Path       *string           `json:"path,omitempty"`
	Queries    *map[string]int32 `json:"queries,omitempty"`
	Url        *string           `json:"url,omitempty"`
	Year       *int32            `json:"year,omitempty"`
}

// NewResultId instantiates a new ResultId object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResultId() *ResultId {
	this := ResultId{}
	return &this
}

// NewResultIdWithDefaults instantiates a new ResultId object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResultIdWithDefaults() *ResultId {
	this := ResultId{}
	return &this
}

// GetAlias returns the Alias field value if set, zero value otherwise.
func (o *ResultId) GetAlias() string {
	if o == nil || IsNil(o.Alias) {
		var ret string
		return ret
	}
	return *o.Alias
}

// GetAliasOk returns a tuple with the Alias field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResultId) GetAliasOk() (*string, bool) {
	if o == nil || IsNil(o.Alias) {
		return nil, false
	}
	return o.Alias, true
}

// HasAlias returns a boolean if a field has been set.
func (o *ResultId) HasAlias() bool {
	if o != nil && !IsNil(o.Alias) {
		return true
	}

	return false
}

// SetAlias gets a reference to the given string and assigns it to the Alias field.
func (o *ResultId) SetAlias(v string) {
	o.Alias = &v
}

// GetApiId returns the ApiId field value if set, zero value otherwise.
func (o *ResultId) GetApiId() string {
	if o == nil || IsNil(o.ApiId) {
		var ret string
		return ret
	}
	return *o.ApiId
}

// GetApiIdOk returns a tuple with the ApiId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResultId) GetApiIdOk() (*string, bool) {
	if o == nil || IsNil(o.ApiId) {
		return nil, false
	}
	return o.ApiId, true
}

// HasApiId returns a boolean if a field has been set.
func (o *ResultId) HasApiId() bool {
	if o != nil && !IsNil(o.ApiId) {
		return true
	}

	return false
}

// SetApiId gets a reference to the given string and assigns it to the ApiId field.
func (o *ResultId) SetApiId(v string) {
	o.ApiId = &v
}

// GetApiName returns the ApiName field value if set, zero value otherwise.
func (o *ResultId) GetApiName() string {
	if o == nil || IsNil(o.ApiName) {
		var ret string
		return ret
	}
	return *o.ApiName
}

// GetApiNameOk returns a tuple with the ApiName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResultId) GetApiNameOk() (*string, bool) {
	if o == nil || IsNil(o.ApiName) {
		return nil, false
	}
	return o.ApiName, true
}

// HasApiName returns a boolean if a field has been set.
func (o *ResultId) HasApiName() bool {
	if o != nil && !IsNil(o.ApiName) {
		return true
	}

	return false
}

// SetApiName gets a reference to the given string and assigns it to the ApiName field.
func (o *ResultId) SetApiName(v string) {
	o.ApiName = &v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *ResultId) GetCode() int32 {
	if o == nil || IsNil(o.Code) {
		var ret int32
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResultId) GetCodeOk() (*int32, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *ResultId) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given int32 and assigns it to the Code field.
func (o *ResultId) SetCode(v int32) {
	o.Code = &v
}

// GetDay returns the Day field value if set, zero value otherwise.
func (o *ResultId) GetDay() int32 {
	if o == nil || IsNil(o.Day) {
		var ret int32
		return ret
	}
	return *o.Day
}

// GetDayOk returns a tuple with the Day field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResultId) GetDayOk() (*int32, bool) {
	if o == nil || IsNil(o.Day) {
		return nil, false
	}
	return o.Day, true
}

// HasDay returns a boolean if a field has been set.
func (o *ResultId) HasDay() bool {
	if o != nil && !IsNil(o.Day) {
		return true
	}

	return false
}

// SetDay gets a reference to the given int32 and assigns it to the Day field.
func (o *ResultId) SetDay(v int32) {
	o.Day = &v
}

// GetHour returns the Hour field value if set, zero value otherwise.
func (o *ResultId) GetHour() int32 {
	if o == nil || IsNil(o.Hour) {
		var ret int32
		return ret
	}
	return *o.Hour
}

// GetHourOk returns a tuple with the Hour field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResultId) GetHourOk() (*int32, bool) {
	if o == nil || IsNil(o.Hour) {
		return nil, false
	}
	return o.Hour, true
}

// HasHour returns a boolean if a field has been set.
func (o *ResultId) HasHour() bool {
	if o != nil && !IsNil(o.Hour) {
		return true
	}

	return false
}

// SetHour gets a reference to the given int32 and assigns it to the Hour field.
func (o *ResultId) SetHour(v int32) {
	o.Hour = &v
}

// GetIsoCountry returns the IsoCountry field value if set, zero value otherwise.
func (o *ResultId) GetIsoCountry() string {
	if o == nil || IsNil(o.IsoCountry) {
		var ret string
		return ret
	}
	return *o.IsoCountry
}

// GetIsoCountryOk returns a tuple with the IsoCountry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResultId) GetIsoCountryOk() (*string, bool) {
	if o == nil || IsNil(o.IsoCountry) {
		return nil, false
	}
	return o.IsoCountry, true
}

// HasIsoCountry returns a boolean if a field has been set.
func (o *ResultId) HasIsoCountry() bool {
	if o != nil && !IsNil(o.IsoCountry) {
		return true
	}

	return false
}

// SetIsoCountry gets a reference to the given string and assigns it to the IsoCountry field.
func (o *ResultId) SetIsoCountry(v string) {
	o.IsoCountry = &v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *ResultId) GetKey() string {
	if o == nil || IsNil(o.Key) {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResultId) GetKeyOk() (*string, bool) {
	if o == nil || IsNil(o.Key) {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *ResultId) HasKey() bool {
	if o != nil && !IsNil(o.Key) {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *ResultId) SetKey(v string) {
	o.Key = &v
}

// GetMonth returns the Month field value if set, zero value otherwise.
func (o *ResultId) GetMonth() int32 {
	if o == nil || IsNil(o.Month) {
		var ret int32
		return ret
	}
	return *o.Month
}

// GetMonthOk returns a tuple with the Month field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResultId) GetMonthOk() (*int32, bool) {
	if o == nil || IsNil(o.Month) {
		return nil, false
	}
	return o.Month, true
}

// HasMonth returns a boolean if a field has been set.
func (o *ResultId) HasMonth() bool {
	if o != nil && !IsNil(o.Month) {
		return true
	}

	return false
}

// SetMonth gets a reference to the given int32 and assigns it to the Month field.
func (o *ResultId) SetMonth(v int32) {
	o.Month = &v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *ResultId) GetPath() string {
	if o == nil || IsNil(o.Path) {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResultId) GetPathOk() (*string, bool) {
	if o == nil || IsNil(o.Path) {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *ResultId) HasPath() bool {
	if o != nil && !IsNil(o.Path) {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *ResultId) SetPath(v string) {
	o.Path = &v
}

// GetQueries returns the Queries field value if set, zero value otherwise.
func (o *ResultId) GetQueries() map[string]int32 {
	if o == nil || IsNil(o.Queries) {
		var ret map[string]int32
		return ret
	}
	return *o.Queries
}

// GetQueriesOk returns a tuple with the Queries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResultId) GetQueriesOk() (*map[string]int32, bool) {
	if o == nil || IsNil(o.Queries) {
		return nil, false
	}
	return o.Queries, true
}

// HasQueries returns a boolean if a field has been set.
func (o *ResultId) HasQueries() bool {
	if o != nil && !IsNil(o.Queries) {
		return true
	}

	return false
}

// SetQueries gets a reference to the given map[string]int32 and assigns it to the Queries field.
func (o *ResultId) SetQueries(v map[string]int32) {
	o.Queries = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *ResultId) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResultId) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *ResultId) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *ResultId) SetUrl(v string) {
	o.Url = &v
}

// GetYear returns the Year field value if set, zero value otherwise.
func (o *ResultId) GetYear() int32 {
	if o == nil || IsNil(o.Year) {
		var ret int32
		return ret
	}
	return *o.Year
}

// GetYearOk returns a tuple with the Year field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResultId) GetYearOk() (*int32, bool) {
	if o == nil || IsNil(o.Year) {
		return nil, false
	}
	return o.Year, true
}

// HasYear returns a boolean if a field has been set.
func (o *ResultId) HasYear() bool {
	if o != nil && !IsNil(o.Year) {
		return true
	}

	return false
}

// SetYear gets a reference to the given int32 and assigns it to the Year field.
func (o *ResultId) SetYear(v int32) {
	o.Year = &v
}

func (o ResultId) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResultId) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Alias) {
		toSerialize["alias"] = o.Alias
	}
	if !IsNil(o.ApiId) {
		toSerialize["api_id"] = o.ApiId
	}
	if !IsNil(o.ApiName) {
		toSerialize["api_name"] = o.ApiName
	}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.Day) {
		toSerialize["day"] = o.Day
	}
	if !IsNil(o.Hour) {
		toSerialize["hour"] = o.Hour
	}
	if !IsNil(o.IsoCountry) {
		toSerialize["iso_country"] = o.IsoCountry
	}
	if !IsNil(o.Key) {
		toSerialize["key"] = o.Key
	}
	if !IsNil(o.Month) {
		toSerialize["month"] = o.Month
	}
	if !IsNil(o.Path) {
		toSerialize["path"] = o.Path
	}
	if !IsNil(o.Queries) {
		toSerialize["queries"] = o.Queries
	}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	if !IsNil(o.Year) {
		toSerialize["year"] = o.Year
	}
	return toSerialize, nil
}

type NullableResultId struct {
	value *ResultId
	isSet bool
}

func (v NullableResultId) Get() *ResultId {
	return v.value
}

func (v *NullableResultId) Set(val *ResultId) {
	v.value = val
	v.isSet = true
}

func (v NullableResultId) IsSet() bool {
	return v.isSet
}

func (v *NullableResultId) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResultId(val *ResultId) *NullableResultId {
	return &NullableResultId{value: val, isSet: true}
}

func (v NullableResultId) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResultId) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
