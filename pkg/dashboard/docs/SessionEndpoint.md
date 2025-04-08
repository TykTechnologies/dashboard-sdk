# SessionEndpoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Methods** | Pointer to [**[]SessionEndpointMethod**](SessionEndpointMethod.md) |  | [optional] 
**Path** | Pointer to **string** |  | [optional] 

## Methods

### NewSessionEndpoint

`func NewSessionEndpoint() *SessionEndpoint`

NewSessionEndpoint instantiates a new SessionEndpoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSessionEndpointWithDefaults

`func NewSessionEndpointWithDefaults() *SessionEndpoint`

NewSessionEndpointWithDefaults instantiates a new SessionEndpoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMethods

`func (o *SessionEndpoint) GetMethods() []SessionEndpointMethod`

GetMethods returns the Methods field if non-nil, zero value otherwise.

### GetMethodsOk

`func (o *SessionEndpoint) GetMethodsOk() (*[]SessionEndpointMethod, bool)`

GetMethodsOk returns a tuple with the Methods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethods

`func (o *SessionEndpoint) SetMethods(v []SessionEndpointMethod)`

SetMethods sets Methods field to given value.

### HasMethods

`func (o *SessionEndpoint) HasMethods() bool`

HasMethods returns a boolean if a field has been set.

### GetPath

`func (o *SessionEndpoint) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *SessionEndpoint) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *SessionEndpoint) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *SessionEndpoint) HasPath() bool`

HasPath returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


