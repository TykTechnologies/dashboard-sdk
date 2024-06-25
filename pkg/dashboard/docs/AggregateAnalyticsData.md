# AggregateAnalyticsData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]ResultUnit**](ResultUnit.md) |  | [optional] 
**Pages** | Pointer to **int32** |  | [optional] 

## Methods

### NewAggregateAnalyticsData

`func NewAggregateAnalyticsData() *AggregateAnalyticsData`

NewAggregateAnalyticsData instantiates a new AggregateAnalyticsData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAggregateAnalyticsDataWithDefaults

`func NewAggregateAnalyticsDataWithDefaults() *AggregateAnalyticsData`

NewAggregateAnalyticsDataWithDefaults instantiates a new AggregateAnalyticsData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *AggregateAnalyticsData) GetData() []ResultUnit`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AggregateAnalyticsData) GetDataOk() (*[]ResultUnit, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AggregateAnalyticsData) SetData(v []ResultUnit)`

SetData sets Data field to given value.

### HasData

`func (o *AggregateAnalyticsData) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *AggregateAnalyticsData) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *AggregateAnalyticsData) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetPages

`func (o *AggregateAnalyticsData) GetPages() int32`

GetPages returns the Pages field if non-nil, zero value otherwise.

### GetPagesOk

`func (o *AggregateAnalyticsData) GetPagesOk() (*int32, bool)`

GetPagesOk returns a tuple with the Pages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPages

`func (o *AggregateAnalyticsData) SetPages(v int32)`

SetPages sets Pages field to given value.

### HasPages

`func (o *AggregateAnalyticsData) HasPages() bool`

HasPages returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


