# SessionEndpointMethod

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Limit** | Pointer to [**SessionEndpointRateLimit**](SessionEndpointRateLimit.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 

## Methods

### NewSessionEndpointMethod

`func NewSessionEndpointMethod() *SessionEndpointMethod`

NewSessionEndpointMethod instantiates a new SessionEndpointMethod object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSessionEndpointMethodWithDefaults

`func NewSessionEndpointMethodWithDefaults() *SessionEndpointMethod`

NewSessionEndpointMethodWithDefaults instantiates a new SessionEndpointMethod object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLimit

`func (o *SessionEndpointMethod) GetLimit() SessionEndpointRateLimit`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *SessionEndpointMethod) GetLimitOk() (*SessionEndpointRateLimit, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *SessionEndpointMethod) SetLimit(v SessionEndpointRateLimit)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *SessionEndpointMethod) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetName

`func (o *SessionEndpointMethod) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SessionEndpointMethod) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SessionEndpointMethod) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SessionEndpointMethod) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


