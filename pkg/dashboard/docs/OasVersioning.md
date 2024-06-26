# OasVersioning

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Default** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**FallbackToDefault** | Pointer to **bool** |  | [optional] 
**Key** | Pointer to **string** |  | [optional] 
**Location** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**StripVersioningData** | Pointer to **bool** |  | [optional] 
**Versions** | Pointer to [**[]OasVersionToID**](OasVersionToID.md) |  | [optional] 

## Methods

### NewOasVersioning

`func NewOasVersioning() *OasVersioning`

NewOasVersioning instantiates a new OasVersioning object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOasVersioningWithDefaults

`func NewOasVersioningWithDefaults() *OasVersioning`

NewOasVersioningWithDefaults instantiates a new OasVersioning object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefault

`func (o *OasVersioning) GetDefault() string`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *OasVersioning) GetDefaultOk() (*string, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *OasVersioning) SetDefault(v string)`

SetDefault sets Default field to given value.

### HasDefault

`func (o *OasVersioning) HasDefault() bool`

HasDefault returns a boolean if a field has been set.

### GetEnabled

`func (o *OasVersioning) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *OasVersioning) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *OasVersioning) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *OasVersioning) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetFallbackToDefault

`func (o *OasVersioning) GetFallbackToDefault() bool`

GetFallbackToDefault returns the FallbackToDefault field if non-nil, zero value otherwise.

### GetFallbackToDefaultOk

`func (o *OasVersioning) GetFallbackToDefaultOk() (*bool, bool)`

GetFallbackToDefaultOk returns a tuple with the FallbackToDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFallbackToDefault

`func (o *OasVersioning) SetFallbackToDefault(v bool)`

SetFallbackToDefault sets FallbackToDefault field to given value.

### HasFallbackToDefault

`func (o *OasVersioning) HasFallbackToDefault() bool`

HasFallbackToDefault returns a boolean if a field has been set.

### GetKey

`func (o *OasVersioning) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *OasVersioning) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *OasVersioning) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *OasVersioning) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetLocation

`func (o *OasVersioning) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *OasVersioning) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *OasVersioning) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *OasVersioning) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetName

`func (o *OasVersioning) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OasVersioning) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OasVersioning) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OasVersioning) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStripVersioningData

`func (o *OasVersioning) GetStripVersioningData() bool`

GetStripVersioningData returns the StripVersioningData field if non-nil, zero value otherwise.

### GetStripVersioningDataOk

`func (o *OasVersioning) GetStripVersioningDataOk() (*bool, bool)`

GetStripVersioningDataOk returns a tuple with the StripVersioningData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStripVersioningData

`func (o *OasVersioning) SetStripVersioningData(v bool)`

SetStripVersioningData sets StripVersioningData field to given value.

### HasStripVersioningData

`func (o *OasVersioning) HasStripVersioningData() bool`

HasStripVersioningData returns a boolean if a field has been set.

### GetVersions

`func (o *OasVersioning) GetVersions() []OasVersionToID`

GetVersions returns the Versions field if non-nil, zero value otherwise.

### GetVersionsOk

`func (o *OasVersioning) GetVersionsOk() (*[]OasVersionToID, bool)`

GetVersionsOk returns a tuple with the Versions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersions

`func (o *OasVersioning) SetVersions(v []OasVersionToID)`

SetVersions sets Versions field to given value.

### HasVersions

`func (o *OasVersioning) HasVersions() bool`

HasVersions returns a boolean if a field has been set.

### SetVersionsNil

`func (o *OasVersioning) SetVersionsNil(b bool)`

 SetVersionsNil sets the value for Versions to be an explicit nil

### UnsetVersions
`func (o *OasVersioning) UnsetVersions()`

UnsetVersions ensures that no value is present for Versions, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


