# ApiCatalogue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Apis** | [**[]APIDescription**](APIDescription.md) |  | 
**Email** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**OrgId** | Pointer to **string** |  | [optional] 

## Methods

### NewApiCatalogue

`func NewApiCatalogue(apis []APIDescription, ) *ApiCatalogue`

NewApiCatalogue instantiates a new ApiCatalogue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiCatalogueWithDefaults

`func NewApiCatalogueWithDefaults() *ApiCatalogue`

NewApiCatalogueWithDefaults instantiates a new ApiCatalogue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApis

`func (o *ApiCatalogue) GetApis() []APIDescription`

GetApis returns the Apis field if non-nil, zero value otherwise.

### GetApisOk

`func (o *ApiCatalogue) GetApisOk() (*[]APIDescription, bool)`

GetApisOk returns a tuple with the Apis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApis

`func (o *ApiCatalogue) SetApis(v []APIDescription)`

SetApis sets Apis field to given value.


### SetApisNil

`func (o *ApiCatalogue) SetApisNil(b bool)`

 SetApisNil sets the value for Apis to be an explicit nil

### UnsetApis
`func (o *ApiCatalogue) UnsetApis()`

UnsetApis ensures that no value is present for Apis, not even an explicit nil
### GetEmail

`func (o *ApiCatalogue) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *ApiCatalogue) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *ApiCatalogue) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *ApiCatalogue) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetId

`func (o *ApiCatalogue) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApiCatalogue) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApiCatalogue) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ApiCatalogue) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOrgId

`func (o *ApiCatalogue) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *ApiCatalogue) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *ApiCatalogue) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *ApiCatalogue) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


