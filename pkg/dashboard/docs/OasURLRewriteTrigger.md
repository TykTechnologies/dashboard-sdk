# OasURLRewriteTrigger

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Condition** | Pointer to **string** |  | [optional] 
**RewriteTo** | Pointer to **string** |  | [optional] 
**Rules** | Pointer to [**[]OasURLRewriteRule**](OasURLRewriteRule.md) |  | [optional] 

## Methods

### NewOasURLRewriteTrigger

`func NewOasURLRewriteTrigger() *OasURLRewriteTrigger`

NewOasURLRewriteTrigger instantiates a new OasURLRewriteTrigger object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOasURLRewriteTriggerWithDefaults

`func NewOasURLRewriteTriggerWithDefaults() *OasURLRewriteTrigger`

NewOasURLRewriteTriggerWithDefaults instantiates a new OasURLRewriteTrigger object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCondition

`func (o *OasURLRewriteTrigger) GetCondition() string`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *OasURLRewriteTrigger) GetConditionOk() (*string, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondition

`func (o *OasURLRewriteTrigger) SetCondition(v string)`

SetCondition sets Condition field to given value.

### HasCondition

`func (o *OasURLRewriteTrigger) HasCondition() bool`

HasCondition returns a boolean if a field has been set.

### GetRewriteTo

`func (o *OasURLRewriteTrigger) GetRewriteTo() string`

GetRewriteTo returns the RewriteTo field if non-nil, zero value otherwise.

### GetRewriteToOk

`func (o *OasURLRewriteTrigger) GetRewriteToOk() (*string, bool)`

GetRewriteToOk returns a tuple with the RewriteTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRewriteTo

`func (o *OasURLRewriteTrigger) SetRewriteTo(v string)`

SetRewriteTo sets RewriteTo field to given value.

### HasRewriteTo

`func (o *OasURLRewriteTrigger) HasRewriteTo() bool`

HasRewriteTo returns a boolean if a field has been set.

### GetRules

`func (o *OasURLRewriteTrigger) GetRules() []OasURLRewriteRule`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *OasURLRewriteTrigger) GetRulesOk() (*[]OasURLRewriteRule, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *OasURLRewriteTrigger) SetRules(v []OasURLRewriteRule)`

SetRules sets Rules field to given value.

### HasRules

`func (o *OasURLRewriteTrigger) HasRules() bool`

HasRules returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


