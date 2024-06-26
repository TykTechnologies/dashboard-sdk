# ApidefRoutingTriggerOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HeaderMatches** | Pointer to [**map[string]ApidefStringRegexMap**](ApidefStringRegexMap.md) |  | [optional] 
**PathPartMatches** | Pointer to [**map[string]ApidefStringRegexMap**](ApidefStringRegexMap.md) |  | [optional] 
**PayloadMatches** | Pointer to [**ApidefStringRegexMap**](ApidefStringRegexMap.md) |  | [optional] 
**QueryValMatches** | Pointer to [**map[string]ApidefStringRegexMap**](ApidefStringRegexMap.md) |  | [optional] 
**RequestContextMatches** | Pointer to [**map[string]ApidefStringRegexMap**](ApidefStringRegexMap.md) |  | [optional] 
**SessionMetaMatches** | Pointer to [**map[string]ApidefStringRegexMap**](ApidefStringRegexMap.md) |  | [optional] 

## Methods

### NewApidefRoutingTriggerOptions

`func NewApidefRoutingTriggerOptions() *ApidefRoutingTriggerOptions`

NewApidefRoutingTriggerOptions instantiates a new ApidefRoutingTriggerOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApidefRoutingTriggerOptionsWithDefaults

`func NewApidefRoutingTriggerOptionsWithDefaults() *ApidefRoutingTriggerOptions`

NewApidefRoutingTriggerOptionsWithDefaults instantiates a new ApidefRoutingTriggerOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHeaderMatches

`func (o *ApidefRoutingTriggerOptions) GetHeaderMatches() map[string]ApidefStringRegexMap`

GetHeaderMatches returns the HeaderMatches field if non-nil, zero value otherwise.

### GetHeaderMatchesOk

`func (o *ApidefRoutingTriggerOptions) GetHeaderMatchesOk() (*map[string]ApidefStringRegexMap, bool)`

GetHeaderMatchesOk returns a tuple with the HeaderMatches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaderMatches

`func (o *ApidefRoutingTriggerOptions) SetHeaderMatches(v map[string]ApidefStringRegexMap)`

SetHeaderMatches sets HeaderMatches field to given value.

### HasHeaderMatches

`func (o *ApidefRoutingTriggerOptions) HasHeaderMatches() bool`

HasHeaderMatches returns a boolean if a field has been set.

### SetHeaderMatchesNil

`func (o *ApidefRoutingTriggerOptions) SetHeaderMatchesNil(b bool)`

 SetHeaderMatchesNil sets the value for HeaderMatches to be an explicit nil

### UnsetHeaderMatches
`func (o *ApidefRoutingTriggerOptions) UnsetHeaderMatches()`

UnsetHeaderMatches ensures that no value is present for HeaderMatches, not even an explicit nil
### GetPathPartMatches

`func (o *ApidefRoutingTriggerOptions) GetPathPartMatches() map[string]ApidefStringRegexMap`

GetPathPartMatches returns the PathPartMatches field if non-nil, zero value otherwise.

### GetPathPartMatchesOk

`func (o *ApidefRoutingTriggerOptions) GetPathPartMatchesOk() (*map[string]ApidefStringRegexMap, bool)`

GetPathPartMatchesOk returns a tuple with the PathPartMatches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPathPartMatches

`func (o *ApidefRoutingTriggerOptions) SetPathPartMatches(v map[string]ApidefStringRegexMap)`

SetPathPartMatches sets PathPartMatches field to given value.

### HasPathPartMatches

`func (o *ApidefRoutingTriggerOptions) HasPathPartMatches() bool`

HasPathPartMatches returns a boolean if a field has been set.

### SetPathPartMatchesNil

`func (o *ApidefRoutingTriggerOptions) SetPathPartMatchesNil(b bool)`

 SetPathPartMatchesNil sets the value for PathPartMatches to be an explicit nil

### UnsetPathPartMatches
`func (o *ApidefRoutingTriggerOptions) UnsetPathPartMatches()`

UnsetPathPartMatches ensures that no value is present for PathPartMatches, not even an explicit nil
### GetPayloadMatches

`func (o *ApidefRoutingTriggerOptions) GetPayloadMatches() ApidefStringRegexMap`

GetPayloadMatches returns the PayloadMatches field if non-nil, zero value otherwise.

### GetPayloadMatchesOk

`func (o *ApidefRoutingTriggerOptions) GetPayloadMatchesOk() (*ApidefStringRegexMap, bool)`

GetPayloadMatchesOk returns a tuple with the PayloadMatches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayloadMatches

`func (o *ApidefRoutingTriggerOptions) SetPayloadMatches(v ApidefStringRegexMap)`

SetPayloadMatches sets PayloadMatches field to given value.

### HasPayloadMatches

`func (o *ApidefRoutingTriggerOptions) HasPayloadMatches() bool`

HasPayloadMatches returns a boolean if a field has been set.

### GetQueryValMatches

`func (o *ApidefRoutingTriggerOptions) GetQueryValMatches() map[string]ApidefStringRegexMap`

GetQueryValMatches returns the QueryValMatches field if non-nil, zero value otherwise.

### GetQueryValMatchesOk

`func (o *ApidefRoutingTriggerOptions) GetQueryValMatchesOk() (*map[string]ApidefStringRegexMap, bool)`

GetQueryValMatchesOk returns a tuple with the QueryValMatches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryValMatches

`func (o *ApidefRoutingTriggerOptions) SetQueryValMatches(v map[string]ApidefStringRegexMap)`

SetQueryValMatches sets QueryValMatches field to given value.

### HasQueryValMatches

`func (o *ApidefRoutingTriggerOptions) HasQueryValMatches() bool`

HasQueryValMatches returns a boolean if a field has been set.

### SetQueryValMatchesNil

`func (o *ApidefRoutingTriggerOptions) SetQueryValMatchesNil(b bool)`

 SetQueryValMatchesNil sets the value for QueryValMatches to be an explicit nil

### UnsetQueryValMatches
`func (o *ApidefRoutingTriggerOptions) UnsetQueryValMatches()`

UnsetQueryValMatches ensures that no value is present for QueryValMatches, not even an explicit nil
### GetRequestContextMatches

`func (o *ApidefRoutingTriggerOptions) GetRequestContextMatches() map[string]ApidefStringRegexMap`

GetRequestContextMatches returns the RequestContextMatches field if non-nil, zero value otherwise.

### GetRequestContextMatchesOk

`func (o *ApidefRoutingTriggerOptions) GetRequestContextMatchesOk() (*map[string]ApidefStringRegexMap, bool)`

GetRequestContextMatchesOk returns a tuple with the RequestContextMatches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestContextMatches

`func (o *ApidefRoutingTriggerOptions) SetRequestContextMatches(v map[string]ApidefStringRegexMap)`

SetRequestContextMatches sets RequestContextMatches field to given value.

### HasRequestContextMatches

`func (o *ApidefRoutingTriggerOptions) HasRequestContextMatches() bool`

HasRequestContextMatches returns a boolean if a field has been set.

### SetRequestContextMatchesNil

`func (o *ApidefRoutingTriggerOptions) SetRequestContextMatchesNil(b bool)`

 SetRequestContextMatchesNil sets the value for RequestContextMatches to be an explicit nil

### UnsetRequestContextMatches
`func (o *ApidefRoutingTriggerOptions) UnsetRequestContextMatches()`

UnsetRequestContextMatches ensures that no value is present for RequestContextMatches, not even an explicit nil
### GetSessionMetaMatches

`func (o *ApidefRoutingTriggerOptions) GetSessionMetaMatches() map[string]ApidefStringRegexMap`

GetSessionMetaMatches returns the SessionMetaMatches field if non-nil, zero value otherwise.

### GetSessionMetaMatchesOk

`func (o *ApidefRoutingTriggerOptions) GetSessionMetaMatchesOk() (*map[string]ApidefStringRegexMap, bool)`

GetSessionMetaMatchesOk returns a tuple with the SessionMetaMatches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionMetaMatches

`func (o *ApidefRoutingTriggerOptions) SetSessionMetaMatches(v map[string]ApidefStringRegexMap)`

SetSessionMetaMatches sets SessionMetaMatches field to given value.

### HasSessionMetaMatches

`func (o *ApidefRoutingTriggerOptions) HasSessionMetaMatches() bool`

HasSessionMetaMatches returns a boolean if a field has been set.

### SetSessionMetaMatchesNil

`func (o *ApidefRoutingTriggerOptions) SetSessionMetaMatchesNil(b bool)`

 SetSessionMetaMatchesNil sets the value for SessionMetaMatches to be an explicit nil

### UnsetSessionMetaMatches
`func (o *ApidefRoutingTriggerOptions) UnsetSessionMetaMatches()`

UnsetSessionMetaMatches ensures that no value is present for SessionMetaMatches, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


