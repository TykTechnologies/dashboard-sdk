# ApidefVersionData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultVersion** | Pointer to **string** |  | [optional] 
**NotVersioned** | Pointer to **bool** |  | [optional] 
**Versions** | Pointer to [**map[string]ApidefVersionInfo**](ApidefVersionInfo.md) |  | [optional] 

## Methods

### NewApidefVersionData

`func NewApidefVersionData() *ApidefVersionData`

NewApidefVersionData instantiates a new ApidefVersionData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApidefVersionDataWithDefaults

`func NewApidefVersionDataWithDefaults() *ApidefVersionData`

NewApidefVersionDataWithDefaults instantiates a new ApidefVersionData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultVersion

`func (o *ApidefVersionData) GetDefaultVersion() string`

GetDefaultVersion returns the DefaultVersion field if non-nil, zero value otherwise.

### GetDefaultVersionOk

`func (o *ApidefVersionData) GetDefaultVersionOk() (*string, bool)`

GetDefaultVersionOk returns a tuple with the DefaultVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultVersion

`func (o *ApidefVersionData) SetDefaultVersion(v string)`

SetDefaultVersion sets DefaultVersion field to given value.

### HasDefaultVersion

`func (o *ApidefVersionData) HasDefaultVersion() bool`

HasDefaultVersion returns a boolean if a field has been set.

### GetNotVersioned

`func (o *ApidefVersionData) GetNotVersioned() bool`

GetNotVersioned returns the NotVersioned field if non-nil, zero value otherwise.

### GetNotVersionedOk

`func (o *ApidefVersionData) GetNotVersionedOk() (*bool, bool)`

GetNotVersionedOk returns a tuple with the NotVersioned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotVersioned

`func (o *ApidefVersionData) SetNotVersioned(v bool)`

SetNotVersioned sets NotVersioned field to given value.

### HasNotVersioned

`func (o *ApidefVersionData) HasNotVersioned() bool`

HasNotVersioned returns a boolean if a field has been set.

### GetVersions

`func (o *ApidefVersionData) GetVersions() map[string]ApidefVersionInfo`

GetVersions returns the Versions field if non-nil, zero value otherwise.

### GetVersionsOk

`func (o *ApidefVersionData) GetVersionsOk() (*map[string]ApidefVersionInfo, bool)`

GetVersionsOk returns a tuple with the Versions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersions

`func (o *ApidefVersionData) SetVersions(v map[string]ApidefVersionInfo)`

SetVersions sets Versions field to given value.

### HasVersions

`func (o *ApidefVersionData) HasVersions() bool`

HasVersions returns a boolean if a field has been set.

### SetVersionsNil

`func (o *ApidefVersionData) SetVersionsNil(b bool)`

 SetVersionsNil sets the value for Versions to be an explicit nil

### UnsetVersions
`func (o *ApidefVersionData) UnsetVersions()`

UnsetVersions ensures that no value is present for Versions, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


