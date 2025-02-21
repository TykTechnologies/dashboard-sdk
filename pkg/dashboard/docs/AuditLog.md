# AuditLog

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MID** | Pointer to **string** |  | [optional] 
**ReqId** | Pointer to **string** |  | [optional] 
**OrgId** | Pointer to **string** |  | [optional] 
**Date** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **int32** |  | [optional] 
**Ip** | Pointer to **string** |  | [optional] 
**User** | Pointer to **string** |  | [optional] 
**Action** | Pointer to **string** |  | [optional] 
**Method** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **int32** |  | [optional] 

## Methods

### NewAuditLog

`func NewAuditLog() *AuditLog`

NewAuditLog instantiates a new AuditLog object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuditLogWithDefaults

`func NewAuditLogWithDefaults() *AuditLog`

NewAuditLogWithDefaults instantiates a new AuditLog object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMID

`func (o *AuditLog) GetMID() string`

GetMID returns the MID field if non-nil, zero value otherwise.

### GetMIDOk

`func (o *AuditLog) GetMIDOk() (*string, bool)`

GetMIDOk returns a tuple with the MID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMID

`func (o *AuditLog) SetMID(v string)`

SetMID sets MID field to given value.

### HasMID

`func (o *AuditLog) HasMID() bool`

HasMID returns a boolean if a field has been set.

### GetReqId

`func (o *AuditLog) GetReqId() string`

GetReqId returns the ReqId field if non-nil, zero value otherwise.

### GetReqIdOk

`func (o *AuditLog) GetReqIdOk() (*string, bool)`

GetReqIdOk returns a tuple with the ReqId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReqId

`func (o *AuditLog) SetReqId(v string)`

SetReqId sets ReqId field to given value.

### HasReqId

`func (o *AuditLog) HasReqId() bool`

HasReqId returns a boolean if a field has been set.

### GetOrgId

`func (o *AuditLog) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *AuditLog) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *AuditLog) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *AuditLog) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetDate

`func (o *AuditLog) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *AuditLog) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *AuditLog) SetDate(v string)`

SetDate sets Date field to given value.

### HasDate

`func (o *AuditLog) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetTimestamp

`func (o *AuditLog) GetTimestamp() int32`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *AuditLog) GetTimestampOk() (*int32, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *AuditLog) SetTimestamp(v int32)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *AuditLog) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetIp

`func (o *AuditLog) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *AuditLog) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *AuditLog) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *AuditLog) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetUser

`func (o *AuditLog) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *AuditLog) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *AuditLog) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *AuditLog) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetAction

`func (o *AuditLog) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *AuditLog) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *AuditLog) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *AuditLog) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetMethod

`func (o *AuditLog) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *AuditLog) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *AuditLog) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *AuditLog) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetUrl

`func (o *AuditLog) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *AuditLog) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *AuditLog) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *AuditLog) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetStatus

`func (o *AuditLog) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AuditLog) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AuditLog) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AuditLog) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


