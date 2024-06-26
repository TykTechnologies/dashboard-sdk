# ApidefURLRewriteMeta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Disabled** | Pointer to **bool** |  | [optional] 
**MatchPattern** | Pointer to **string** |  | [optional] 
**Method** | Pointer to **string** |  | [optional] 
**Path** | Pointer to **string** |  | [optional] 
**RewriteTo** | Pointer to **string** |  | [optional] 
**Triggers** | Pointer to [**[]ApidefRoutingTrigger**](ApidefRoutingTrigger.md) |  | [optional] 

## Methods

### NewApidefURLRewriteMeta

`func NewApidefURLRewriteMeta() *ApidefURLRewriteMeta`

NewApidefURLRewriteMeta instantiates a new ApidefURLRewriteMeta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApidefURLRewriteMetaWithDefaults

`func NewApidefURLRewriteMetaWithDefaults() *ApidefURLRewriteMeta`

NewApidefURLRewriteMetaWithDefaults instantiates a new ApidefURLRewriteMeta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisabled

`func (o *ApidefURLRewriteMeta) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *ApidefURLRewriteMeta) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *ApidefURLRewriteMeta) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *ApidefURLRewriteMeta) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetMatchPattern

`func (o *ApidefURLRewriteMeta) GetMatchPattern() string`

GetMatchPattern returns the MatchPattern field if non-nil, zero value otherwise.

### GetMatchPatternOk

`func (o *ApidefURLRewriteMeta) GetMatchPatternOk() (*string, bool)`

GetMatchPatternOk returns a tuple with the MatchPattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchPattern

`func (o *ApidefURLRewriteMeta) SetMatchPattern(v string)`

SetMatchPattern sets MatchPattern field to given value.

### HasMatchPattern

`func (o *ApidefURLRewriteMeta) HasMatchPattern() bool`

HasMatchPattern returns a boolean if a field has been set.

### GetMethod

`func (o *ApidefURLRewriteMeta) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *ApidefURLRewriteMeta) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *ApidefURLRewriteMeta) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *ApidefURLRewriteMeta) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetPath

`func (o *ApidefURLRewriteMeta) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *ApidefURLRewriteMeta) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *ApidefURLRewriteMeta) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *ApidefURLRewriteMeta) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetRewriteTo

`func (o *ApidefURLRewriteMeta) GetRewriteTo() string`

GetRewriteTo returns the RewriteTo field if non-nil, zero value otherwise.

### GetRewriteToOk

`func (o *ApidefURLRewriteMeta) GetRewriteToOk() (*string, bool)`

GetRewriteToOk returns a tuple with the RewriteTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRewriteTo

`func (o *ApidefURLRewriteMeta) SetRewriteTo(v string)`

SetRewriteTo sets RewriteTo field to given value.

### HasRewriteTo

`func (o *ApidefURLRewriteMeta) HasRewriteTo() bool`

HasRewriteTo returns a boolean if a field has been set.

### GetTriggers

`func (o *ApidefURLRewriteMeta) GetTriggers() []ApidefRoutingTrigger`

GetTriggers returns the Triggers field if non-nil, zero value otherwise.

### GetTriggersOk

`func (o *ApidefURLRewriteMeta) GetTriggersOk() (*[]ApidefRoutingTrigger, bool)`

GetTriggersOk returns a tuple with the Triggers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggers

`func (o *ApidefURLRewriteMeta) SetTriggers(v []ApidefRoutingTrigger)`

SetTriggers sets Triggers field to given value.

### HasTriggers

`func (o *ApidefURLRewriteMeta) HasTriggers() bool`

HasTriggers returns a boolean if a field has been set.

### SetTriggersNil

`func (o *ApidefURLRewriteMeta) SetTriggersNil(b bool)`

 SetTriggersNil sets the value for Triggers to be an explicit nil

### UnsetTriggers
`func (o *ApidefURLRewriteMeta) UnsetTriggers()`

UnsetTriggers ensures that no value is present for Triggers, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


