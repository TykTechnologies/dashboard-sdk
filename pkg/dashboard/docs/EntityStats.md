# EntityStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Apis** | Pointer to [**[]IndividualStats**](IndividualStats.md) |  | [optional] 
**Dataplanes** | Pointer to [**[]IndividualStats**](IndividualStats.md) |  | [optional] 
**DataplanesGateways** | Pointer to [**[]IndividualStats**](IndividualStats.md) |  | [optional] 

## Methods

### NewEntityStats

`func NewEntityStats() *EntityStats`

NewEntityStats instantiates a new EntityStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntityStatsWithDefaults

`func NewEntityStatsWithDefaults() *EntityStats`

NewEntityStatsWithDefaults instantiates a new EntityStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApis

`func (o *EntityStats) GetApis() []IndividualStats`

GetApis returns the Apis field if non-nil, zero value otherwise.

### GetApisOk

`func (o *EntityStats) GetApisOk() (*[]IndividualStats, bool)`

GetApisOk returns a tuple with the Apis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApis

`func (o *EntityStats) SetApis(v []IndividualStats)`

SetApis sets Apis field to given value.

### HasApis

`func (o *EntityStats) HasApis() bool`

HasApis returns a boolean if a field has been set.

### SetApisNil

`func (o *EntityStats) SetApisNil(b bool)`

 SetApisNil sets the value for Apis to be an explicit nil

### UnsetApis
`func (o *EntityStats) UnsetApis()`

UnsetApis ensures that no value is present for Apis, not even an explicit nil
### GetDataplanes

`func (o *EntityStats) GetDataplanes() []IndividualStats`

GetDataplanes returns the Dataplanes field if non-nil, zero value otherwise.

### GetDataplanesOk

`func (o *EntityStats) GetDataplanesOk() (*[]IndividualStats, bool)`

GetDataplanesOk returns a tuple with the Dataplanes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataplanes

`func (o *EntityStats) SetDataplanes(v []IndividualStats)`

SetDataplanes sets Dataplanes field to given value.

### HasDataplanes

`func (o *EntityStats) HasDataplanes() bool`

HasDataplanes returns a boolean if a field has been set.

### GetDataplanesGateways

`func (o *EntityStats) GetDataplanesGateways() []IndividualStats`

GetDataplanesGateways returns the DataplanesGateways field if non-nil, zero value otherwise.

### GetDataplanesGatewaysOk

`func (o *EntityStats) GetDataplanesGatewaysOk() (*[]IndividualStats, bool)`

GetDataplanesGatewaysOk returns a tuple with the DataplanesGateways field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataplanesGateways

`func (o *EntityStats) SetDataplanesGateways(v []IndividualStats)`

SetDataplanesGateways sets DataplanesGateways field to given value.

### HasDataplanesGateways

`func (o *EntityStats) HasDataplanesGateways() bool`

HasDataplanesGateways returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


