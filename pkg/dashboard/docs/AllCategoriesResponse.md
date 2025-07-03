# AllCategoriesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Categories** | Pointer to [**[]CategoryCount**](CategoryCount.md) |  | [optional] 

## Methods

### NewAllCategoriesResponse

`func NewAllCategoriesResponse() *AllCategoriesResponse`

NewAllCategoriesResponse instantiates a new AllCategoriesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAllCategoriesResponseWithDefaults

`func NewAllCategoriesResponseWithDefaults() *AllCategoriesResponse`

NewAllCategoriesResponseWithDefaults instantiates a new AllCategoriesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategories

`func (o *AllCategoriesResponse) GetCategories() []CategoryCount`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *AllCategoriesResponse) GetCategoriesOk() (*[]CategoryCount, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *AllCategoriesResponse) SetCategories(v []CategoryCount)`

SetCategories sets Categories field to given value.

### HasCategories

`func (o *AllCategoriesResponse) HasCategories() bool`

HasCategories returns a boolean if a field has been set.

### SetCategoriesNil

`func (o *AllCategoriesResponse) SetCategoriesNil(b bool)`

 SetCategoriesNil sets the value for Categories to be an explicit nil

### UnsetCategories
`func (o *AllCategoriesResponse) UnsetCategories()`

UnsetCategories ensures that no value is present for Categories, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


