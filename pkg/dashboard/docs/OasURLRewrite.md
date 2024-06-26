# OasURLRewrite

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** |  | [optional] 
**Pattern** | Pointer to **string** |  | [optional] 
**RewriteTo** | Pointer to **string** |  | [optional] 
**Triggers** | Pointer to [**[]OasURLRewriteTrigger**](OasURLRewriteTrigger.md) |  | [optional] 

## Methods

### NewOasURLRewrite

`func NewOasURLRewrite() *OasURLRewrite`

NewOasURLRewrite instantiates a new OasURLRewrite object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOasURLRewriteWithDefaults

`func NewOasURLRewriteWithDefaults() *OasURLRewrite`

NewOasURLRewriteWithDefaults instantiates a new OasURLRewrite object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *OasURLRewrite) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *OasURLRewrite) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *OasURLRewrite) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *OasURLRewrite) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetPattern

`func (o *OasURLRewrite) GetPattern() string`

GetPattern returns the Pattern field if non-nil, zero value otherwise.

### GetPatternOk

`func (o *OasURLRewrite) GetPatternOk() (*string, bool)`

GetPatternOk returns a tuple with the Pattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPattern

`func (o *OasURLRewrite) SetPattern(v string)`

SetPattern sets Pattern field to given value.

### HasPattern

`func (o *OasURLRewrite) HasPattern() bool`

HasPattern returns a boolean if a field has been set.

### GetRewriteTo

`func (o *OasURLRewrite) GetRewriteTo() string`

GetRewriteTo returns the RewriteTo field if non-nil, zero value otherwise.

### GetRewriteToOk

`func (o *OasURLRewrite) GetRewriteToOk() (*string, bool)`

GetRewriteToOk returns a tuple with the RewriteTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRewriteTo

`func (o *OasURLRewrite) SetRewriteTo(v string)`

SetRewriteTo sets RewriteTo field to given value.

### HasRewriteTo

`func (o *OasURLRewrite) HasRewriteTo() bool`

HasRewriteTo returns a boolean if a field has been set.

### GetTriggers

`func (o *OasURLRewrite) GetTriggers() []OasURLRewriteTrigger`

GetTriggers returns the Triggers field if non-nil, zero value otherwise.

### GetTriggersOk

`func (o *OasURLRewrite) GetTriggersOk() (*[]OasURLRewriteTrigger, bool)`

GetTriggersOk returns a tuple with the Triggers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggers

`func (o *OasURLRewrite) SetTriggers(v []OasURLRewriteTrigger)`

SetTriggers sets Triggers field to given value.

### HasTriggers

`func (o *OasURLRewrite) HasTriggers() bool`

HasTriggers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


