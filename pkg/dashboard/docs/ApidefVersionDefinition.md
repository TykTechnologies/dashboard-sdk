# ApidefVersionDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Default** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**FallbackToDefault** | Pointer to **bool** |  | [optional] 
**Key** | Pointer to **string** |  | [optional] 
**Location** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**StripPath** | Pointer to **bool** |  | [optional] 
**StripVersioningData** | Pointer to **bool** |  | [optional] 
**Versions** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewApidefVersionDefinition

`func NewApidefVersionDefinition() *ApidefVersionDefinition`

NewApidefVersionDefinition instantiates a new ApidefVersionDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApidefVersionDefinitionWithDefaults

`func NewApidefVersionDefinitionWithDefaults() *ApidefVersionDefinition`

NewApidefVersionDefinitionWithDefaults instantiates a new ApidefVersionDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefault

`func (o *ApidefVersionDefinition) GetDefault() string`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *ApidefVersionDefinition) GetDefaultOk() (*string, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *ApidefVersionDefinition) SetDefault(v string)`

SetDefault sets Default field to given value.

### HasDefault

`func (o *ApidefVersionDefinition) HasDefault() bool`

HasDefault returns a boolean if a field has been set.

### GetEnabled

`func (o *ApidefVersionDefinition) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ApidefVersionDefinition) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ApidefVersionDefinition) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ApidefVersionDefinition) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetFallbackToDefault

`func (o *ApidefVersionDefinition) GetFallbackToDefault() bool`

GetFallbackToDefault returns the FallbackToDefault field if non-nil, zero value otherwise.

### GetFallbackToDefaultOk

`func (o *ApidefVersionDefinition) GetFallbackToDefaultOk() (*bool, bool)`

GetFallbackToDefaultOk returns a tuple with the FallbackToDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFallbackToDefault

`func (o *ApidefVersionDefinition) SetFallbackToDefault(v bool)`

SetFallbackToDefault sets FallbackToDefault field to given value.

### HasFallbackToDefault

`func (o *ApidefVersionDefinition) HasFallbackToDefault() bool`

HasFallbackToDefault returns a boolean if a field has been set.

### GetKey

`func (o *ApidefVersionDefinition) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *ApidefVersionDefinition) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *ApidefVersionDefinition) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *ApidefVersionDefinition) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetLocation

`func (o *ApidefVersionDefinition) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *ApidefVersionDefinition) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *ApidefVersionDefinition) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *ApidefVersionDefinition) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetName

`func (o *ApidefVersionDefinition) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApidefVersionDefinition) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApidefVersionDefinition) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ApidefVersionDefinition) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStripPath

`func (o *ApidefVersionDefinition) GetStripPath() bool`

GetStripPath returns the StripPath field if non-nil, zero value otherwise.

### GetStripPathOk

`func (o *ApidefVersionDefinition) GetStripPathOk() (*bool, bool)`

GetStripPathOk returns a tuple with the StripPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStripPath

`func (o *ApidefVersionDefinition) SetStripPath(v bool)`

SetStripPath sets StripPath field to given value.

### HasStripPath

`func (o *ApidefVersionDefinition) HasStripPath() bool`

HasStripPath returns a boolean if a field has been set.

### GetStripVersioningData

`func (o *ApidefVersionDefinition) GetStripVersioningData() bool`

GetStripVersioningData returns the StripVersioningData field if non-nil, zero value otherwise.

### GetStripVersioningDataOk

`func (o *ApidefVersionDefinition) GetStripVersioningDataOk() (*bool, bool)`

GetStripVersioningDataOk returns a tuple with the StripVersioningData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStripVersioningData

`func (o *ApidefVersionDefinition) SetStripVersioningData(v bool)`

SetStripVersioningData sets StripVersioningData field to given value.

### HasStripVersioningData

`func (o *ApidefVersionDefinition) HasStripVersioningData() bool`

HasStripVersioningData returns a boolean if a field has been set.

### GetVersions

`func (o *ApidefVersionDefinition) GetVersions() map[string]string`

GetVersions returns the Versions field if non-nil, zero value otherwise.

### GetVersionsOk

`func (o *ApidefVersionDefinition) GetVersionsOk() (*map[string]string, bool)`

GetVersionsOk returns a tuple with the Versions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersions

`func (o *ApidefVersionDefinition) SetVersions(v map[string]string)`

SetVersions sets Versions field to given value.

### HasVersions

`func (o *ApidefVersionDefinition) HasVersions() bool`

HasVersions returns a boolean if a field has been set.

### SetVersionsNil

`func (o *ApidefVersionDefinition) SetVersionsNil(b bool)`

 SetVersionsNil sets the value for Versions to be an explicit nil

### UnsetVersions
`func (o *ApidefVersionDefinition) UnsetVersions()`

UnsetVersions ensures that no value is present for Versions, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


