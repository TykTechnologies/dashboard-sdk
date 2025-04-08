# AuditLogs

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pages** | Pointer to **int32** |  | [optional] 
**AuditLogs** | Pointer to [**[]AuditLog**](AuditLog.md) |  | [optional] 

## Methods

### NewAuditLogs

`func NewAuditLogs() *AuditLogs`

NewAuditLogs instantiates a new AuditLogs object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuditLogsWithDefaults

`func NewAuditLogsWithDefaults() *AuditLogs`

NewAuditLogsWithDefaults instantiates a new AuditLogs object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPages

`func (o *AuditLogs) GetPages() int32`

GetPages returns the Pages field if non-nil, zero value otherwise.

### GetPagesOk

`func (o *AuditLogs) GetPagesOk() (*int32, bool)`

GetPagesOk returns a tuple with the Pages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPages

`func (o *AuditLogs) SetPages(v int32)`

SetPages sets Pages field to given value.

### HasPages

`func (o *AuditLogs) HasPages() bool`

HasPages returns a boolean if a field has been set.

### GetAuditLogs

`func (o *AuditLogs) GetAuditLogs() []AuditLog`

GetAuditLogs returns the AuditLogs field if non-nil, zero value otherwise.

### GetAuditLogsOk

`func (o *AuditLogs) GetAuditLogsOk() (*[]AuditLog, bool)`

GetAuditLogsOk returns a tuple with the AuditLogs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditLogs

`func (o *AuditLogs) SetAuditLogs(v []AuditLog)`

SetAuditLogs sets AuditLogs field to given value.

### HasAuditLogs

`func (o *AuditLogs) HasAuditLogs() bool`

HasAuditLogs returns a boolean if a field has been set.

### SetAuditLogsNil

`func (o *AuditLogs) SetAuditLogsNil(b bool)`

 SetAuditLogsNil sets the value for AuditLogs to be an explicit nil

### UnsetAuditLogs
`func (o *AuditLogs) UnsetAuditLogs()`

UnsetAuditLogs ensures that no value is present for AuditLogs, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


