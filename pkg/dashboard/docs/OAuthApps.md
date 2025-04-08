# OAuthApps

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Apps** | Pointer to [**[]OAuthClient**](OAuthClient.md) |  | [optional] 
**Pages** | Pointer to **int32** |  | [optional] 

## Methods

### NewOAuthApps

`func NewOAuthApps() *OAuthApps`

NewOAuthApps instantiates a new OAuthApps object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOAuthAppsWithDefaults

`func NewOAuthAppsWithDefaults() *OAuthApps`

NewOAuthAppsWithDefaults instantiates a new OAuthApps object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApps

`func (o *OAuthApps) GetApps() []OAuthClient`

GetApps returns the Apps field if non-nil, zero value otherwise.

### GetAppsOk

`func (o *OAuthApps) GetAppsOk() (*[]OAuthClient, bool)`

GetAppsOk returns a tuple with the Apps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApps

`func (o *OAuthApps) SetApps(v []OAuthClient)`

SetApps sets Apps field to given value.

### HasApps

`func (o *OAuthApps) HasApps() bool`

HasApps returns a boolean if a field has been set.

### SetAppsNil

`func (o *OAuthApps) SetAppsNil(b bool)`

 SetAppsNil sets the value for Apps to be an explicit nil

### UnsetApps
`func (o *OAuthApps) UnsetApps()`

UnsetApps ensures that no value is present for Apps, not even an explicit nil
### GetPages

`func (o *OAuthApps) GetPages() int32`

GetPages returns the Pages field if non-nil, zero value otherwise.

### GetPagesOk

`func (o *OAuthApps) GetPagesOk() (*int32, bool)`

GetPagesOk returns a tuple with the Pages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPages

`func (o *OAuthApps) SetPages(v int32)`

SetPages sets Pages field to given value.

### HasPages

`func (o *OAuthApps) HasPages() bool`

HasPages returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


