# AuthSources

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cookie** | Pointer to [**NullableAuthSource**](AuthSource.md) |  | [optional] 
**Header** | Pointer to [**NullableAuthSource**](AuthSource.md) |  | [optional] 
**Query** | Pointer to [**NullableAuthSource**](AuthSource.md) |  | [optional] 

## Methods

### NewAuthSources

`func NewAuthSources() *AuthSources`

NewAuthSources instantiates a new AuthSources object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthSourcesWithDefaults

`func NewAuthSourcesWithDefaults() *AuthSources`

NewAuthSourcesWithDefaults instantiates a new AuthSources object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCookie

`func (o *AuthSources) GetCookie() AuthSource`

GetCookie returns the Cookie field if non-nil, zero value otherwise.

### GetCookieOk

`func (o *AuthSources) GetCookieOk() (*AuthSource, bool)`

GetCookieOk returns a tuple with the Cookie field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCookie

`func (o *AuthSources) SetCookie(v AuthSource)`

SetCookie sets Cookie field to given value.

### HasCookie

`func (o *AuthSources) HasCookie() bool`

HasCookie returns a boolean if a field has been set.

### SetCookieNil

`func (o *AuthSources) SetCookieNil(b bool)`

 SetCookieNil sets the value for Cookie to be an explicit nil

### UnsetCookie
`func (o *AuthSources) UnsetCookie()`

UnsetCookie ensures that no value is present for Cookie, not even an explicit nil
### GetHeader

`func (o *AuthSources) GetHeader() AuthSource`

GetHeader returns the Header field if non-nil, zero value otherwise.

### GetHeaderOk

`func (o *AuthSources) GetHeaderOk() (*AuthSource, bool)`

GetHeaderOk returns a tuple with the Header field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeader

`func (o *AuthSources) SetHeader(v AuthSource)`

SetHeader sets Header field to given value.

### HasHeader

`func (o *AuthSources) HasHeader() bool`

HasHeader returns a boolean if a field has been set.

### SetHeaderNil

`func (o *AuthSources) SetHeaderNil(b bool)`

 SetHeaderNil sets the value for Header to be an explicit nil

### UnsetHeader
`func (o *AuthSources) UnsetHeader()`

UnsetHeader ensures that no value is present for Header, not even an explicit nil
### GetQuery

`func (o *AuthSources) GetQuery() AuthSource`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *AuthSources) GetQueryOk() (*AuthSource, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *AuthSources) SetQuery(v AuthSource)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *AuthSources) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### SetQueryNil

`func (o *AuthSources) SetQueryNil(b bool)`

 SetQueryNil sets the value for Query to be an explicit nil

### UnsetQuery
`func (o *AuthSources) UnsetQuery()`

UnsetQuery ensures that no value is present for Query, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


