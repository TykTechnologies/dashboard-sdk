# Asset

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mid** | Pointer to **string** | database id of asset | [optional] [readonly] 
**Data** | Pointer to **map[string]interface{}** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** | custom id of asset that can be provided while creation/updating. | [optional] 
**Kind** | Pointer to **string** | the asset type, which is set to oas-template | [optional] [default to "oas-template"]
**LastUpdated** | Pointer to **time.Time** | last updated time stamp | [optional] [readonly] 
**Name** | Pointer to **string** | human-readable name for the template | [optional] 
**OrgId** | Pointer to **string** |  | [optional] 

## Methods

### NewAsset

`func NewAsset() *Asset`

NewAsset instantiates a new Asset object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetWithDefaults

`func NewAssetWithDefaults() *Asset`

NewAssetWithDefaults instantiates a new Asset object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMid

`func (o *Asset) GetMid() string`

GetMid returns the Mid field if non-nil, zero value otherwise.

### GetMidOk

`func (o *Asset) GetMidOk() (*string, bool)`

GetMidOk returns a tuple with the Mid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMid

`func (o *Asset) SetMid(v string)`

SetMid sets Mid field to given value.

### HasMid

`func (o *Asset) HasMid() bool`

HasMid returns a boolean if a field has been set.

### GetData

`func (o *Asset) GetData() map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *Asset) GetDataOk() (*map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *Asset) SetData(v map[string]interface{})`

SetData sets Data field to given value.

### HasData

`func (o *Asset) HasData() bool`

HasData returns a boolean if a field has been set.

### GetDescription

`func (o *Asset) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Asset) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Asset) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Asset) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *Asset) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Asset) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Asset) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Asset) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKind

`func (o *Asset) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *Asset) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *Asset) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *Asset) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetLastUpdated

`func (o *Asset) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *Asset) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *Asset) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *Asset) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetName

`func (o *Asset) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Asset) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Asset) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Asset) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrgId

`func (o *Asset) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *Asset) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *Asset) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *Asset) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


