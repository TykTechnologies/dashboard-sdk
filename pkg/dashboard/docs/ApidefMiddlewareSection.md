# ApidefMiddlewareSection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthCheck** | Pointer to [**ApidefMiddlewareDefinition**](ApidefMiddlewareDefinition.md) |  | [optional] 
**Driver** | Pointer to **string** |  | [optional] 
**IdExtractor** | Pointer to [**ApidefMiddlewareIdExtractor**](ApidefMiddlewareIdExtractor.md) |  | [optional] 
**Post** | Pointer to [**[]ApidefMiddlewareDefinition**](ApidefMiddlewareDefinition.md) |  | [optional] 
**PostKeyAuth** | Pointer to [**[]ApidefMiddlewareDefinition**](ApidefMiddlewareDefinition.md) |  | [optional] 
**Pre** | Pointer to [**[]ApidefMiddlewareDefinition**](ApidefMiddlewareDefinition.md) |  | [optional] 
**Response** | Pointer to [**[]ApidefMiddlewareDefinition**](ApidefMiddlewareDefinition.md) |  | [optional] 

## Methods

### NewApidefMiddlewareSection

`func NewApidefMiddlewareSection() *ApidefMiddlewareSection`

NewApidefMiddlewareSection instantiates a new ApidefMiddlewareSection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApidefMiddlewareSectionWithDefaults

`func NewApidefMiddlewareSectionWithDefaults() *ApidefMiddlewareSection`

NewApidefMiddlewareSectionWithDefaults instantiates a new ApidefMiddlewareSection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthCheck

`func (o *ApidefMiddlewareSection) GetAuthCheck() ApidefMiddlewareDefinition`

GetAuthCheck returns the AuthCheck field if non-nil, zero value otherwise.

### GetAuthCheckOk

`func (o *ApidefMiddlewareSection) GetAuthCheckOk() (*ApidefMiddlewareDefinition, bool)`

GetAuthCheckOk returns a tuple with the AuthCheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthCheck

`func (o *ApidefMiddlewareSection) SetAuthCheck(v ApidefMiddlewareDefinition)`

SetAuthCheck sets AuthCheck field to given value.

### HasAuthCheck

`func (o *ApidefMiddlewareSection) HasAuthCheck() bool`

HasAuthCheck returns a boolean if a field has been set.

### GetDriver

`func (o *ApidefMiddlewareSection) GetDriver() string`

GetDriver returns the Driver field if non-nil, zero value otherwise.

### GetDriverOk

`func (o *ApidefMiddlewareSection) GetDriverOk() (*string, bool)`

GetDriverOk returns a tuple with the Driver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriver

`func (o *ApidefMiddlewareSection) SetDriver(v string)`

SetDriver sets Driver field to given value.

### HasDriver

`func (o *ApidefMiddlewareSection) HasDriver() bool`

HasDriver returns a boolean if a field has been set.

### GetIdExtractor

`func (o *ApidefMiddlewareSection) GetIdExtractor() ApidefMiddlewareIdExtractor`

GetIdExtractor returns the IdExtractor field if non-nil, zero value otherwise.

### GetIdExtractorOk

`func (o *ApidefMiddlewareSection) GetIdExtractorOk() (*ApidefMiddlewareIdExtractor, bool)`

GetIdExtractorOk returns a tuple with the IdExtractor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdExtractor

`func (o *ApidefMiddlewareSection) SetIdExtractor(v ApidefMiddlewareIdExtractor)`

SetIdExtractor sets IdExtractor field to given value.

### HasIdExtractor

`func (o *ApidefMiddlewareSection) HasIdExtractor() bool`

HasIdExtractor returns a boolean if a field has been set.

### GetPost

`func (o *ApidefMiddlewareSection) GetPost() []ApidefMiddlewareDefinition`

GetPost returns the Post field if non-nil, zero value otherwise.

### GetPostOk

`func (o *ApidefMiddlewareSection) GetPostOk() (*[]ApidefMiddlewareDefinition, bool)`

GetPostOk returns a tuple with the Post field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPost

`func (o *ApidefMiddlewareSection) SetPost(v []ApidefMiddlewareDefinition)`

SetPost sets Post field to given value.

### HasPost

`func (o *ApidefMiddlewareSection) HasPost() bool`

HasPost returns a boolean if a field has been set.

### SetPostNil

`func (o *ApidefMiddlewareSection) SetPostNil(b bool)`

 SetPostNil sets the value for Post to be an explicit nil

### UnsetPost
`func (o *ApidefMiddlewareSection) UnsetPost()`

UnsetPost ensures that no value is present for Post, not even an explicit nil
### GetPostKeyAuth

`func (o *ApidefMiddlewareSection) GetPostKeyAuth() []ApidefMiddlewareDefinition`

GetPostKeyAuth returns the PostKeyAuth field if non-nil, zero value otherwise.

### GetPostKeyAuthOk

`func (o *ApidefMiddlewareSection) GetPostKeyAuthOk() (*[]ApidefMiddlewareDefinition, bool)`

GetPostKeyAuthOk returns a tuple with the PostKeyAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostKeyAuth

`func (o *ApidefMiddlewareSection) SetPostKeyAuth(v []ApidefMiddlewareDefinition)`

SetPostKeyAuth sets PostKeyAuth field to given value.

### HasPostKeyAuth

`func (o *ApidefMiddlewareSection) HasPostKeyAuth() bool`

HasPostKeyAuth returns a boolean if a field has been set.

### SetPostKeyAuthNil

`func (o *ApidefMiddlewareSection) SetPostKeyAuthNil(b bool)`

 SetPostKeyAuthNil sets the value for PostKeyAuth to be an explicit nil

### UnsetPostKeyAuth
`func (o *ApidefMiddlewareSection) UnsetPostKeyAuth()`

UnsetPostKeyAuth ensures that no value is present for PostKeyAuth, not even an explicit nil
### GetPre

`func (o *ApidefMiddlewareSection) GetPre() []ApidefMiddlewareDefinition`

GetPre returns the Pre field if non-nil, zero value otherwise.

### GetPreOk

`func (o *ApidefMiddlewareSection) GetPreOk() (*[]ApidefMiddlewareDefinition, bool)`

GetPreOk returns a tuple with the Pre field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPre

`func (o *ApidefMiddlewareSection) SetPre(v []ApidefMiddlewareDefinition)`

SetPre sets Pre field to given value.

### HasPre

`func (o *ApidefMiddlewareSection) HasPre() bool`

HasPre returns a boolean if a field has been set.

### SetPreNil

`func (o *ApidefMiddlewareSection) SetPreNil(b bool)`

 SetPreNil sets the value for Pre to be an explicit nil

### UnsetPre
`func (o *ApidefMiddlewareSection) UnsetPre()`

UnsetPre ensures that no value is present for Pre, not even an explicit nil
### GetResponse

`func (o *ApidefMiddlewareSection) GetResponse() []ApidefMiddlewareDefinition`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *ApidefMiddlewareSection) GetResponseOk() (*[]ApidefMiddlewareDefinition, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *ApidefMiddlewareSection) SetResponse(v []ApidefMiddlewareDefinition)`

SetResponse sets Response field to given value.

### HasResponse

`func (o *ApidefMiddlewareSection) HasResponse() bool`

HasResponse returns a boolean if a field has been set.

### SetResponseNil

`func (o *ApidefMiddlewareSection) SetResponseNil(b bool)`

 SetResponseNil sets the value for Response to be an explicit nil

### UnsetResponse
`func (o *ApidefMiddlewareSection) UnsetResponse()`

UnsetResponse ensures that no value is present for Response, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


