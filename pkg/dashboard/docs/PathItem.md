# PathItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** |  | [optional] 
**Connect** | Pointer to [**NullableOperationType2**](OperationType2.md) |  | [optional] 
**Delete** | Pointer to [**NullableOperationType2**](OperationType2.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Get** | Pointer to [**NullableOperationType2**](OperationType2.md) |  | [optional] 
**Head** | Pointer to [**NullableOperationType2**](OperationType2.md) |  | [optional] 
**Options** | Pointer to [**NullableOperationType2**](OperationType2.md) |  | [optional] 
**Parameters** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Patch** | Pointer to [**NullableOperationType2**](OperationType2.md) |  | [optional] 
**Post** | Pointer to [**NullableOperationType2**](OperationType2.md) |  | [optional] 
**Put** | Pointer to [**NullableOperationType2**](OperationType2.md) |  | [optional] 
**Servers** | Pointer to [**[]ServerType2**](ServerType2.md) |  | [optional] 
**Summary** | Pointer to **string** |  | [optional] 
**Trace** | Pointer to [**NullableOperationType2**](OperationType2.md) |  | [optional] 

## Methods

### NewPathItem

`func NewPathItem() *PathItem`

NewPathItem instantiates a new PathItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPathItemWithDefaults

`func NewPathItemWithDefaults() *PathItem`

NewPathItemWithDefaults instantiates a new PathItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *PathItem) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *PathItem) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *PathItem) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *PathItem) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetConnect

`func (o *PathItem) GetConnect() OperationType2`

GetConnect returns the Connect field if non-nil, zero value otherwise.

### GetConnectOk

`func (o *PathItem) GetConnectOk() (*OperationType2, bool)`

GetConnectOk returns a tuple with the Connect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnect

`func (o *PathItem) SetConnect(v OperationType2)`

SetConnect sets Connect field to given value.

### HasConnect

`func (o *PathItem) HasConnect() bool`

HasConnect returns a boolean if a field has been set.

### SetConnectNil

`func (o *PathItem) SetConnectNil(b bool)`

 SetConnectNil sets the value for Connect to be an explicit nil

### UnsetConnect
`func (o *PathItem) UnsetConnect()`

UnsetConnect ensures that no value is present for Connect, not even an explicit nil
### GetDelete

`func (o *PathItem) GetDelete() OperationType2`

GetDelete returns the Delete field if non-nil, zero value otherwise.

### GetDeleteOk

`func (o *PathItem) GetDeleteOk() (*OperationType2, bool)`

GetDeleteOk returns a tuple with the Delete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelete

`func (o *PathItem) SetDelete(v OperationType2)`

SetDelete sets Delete field to given value.

### HasDelete

`func (o *PathItem) HasDelete() bool`

HasDelete returns a boolean if a field has been set.

### SetDeleteNil

`func (o *PathItem) SetDeleteNil(b bool)`

 SetDeleteNil sets the value for Delete to be an explicit nil

### UnsetDelete
`func (o *PathItem) UnsetDelete()`

UnsetDelete ensures that no value is present for Delete, not even an explicit nil
### GetDescription

`func (o *PathItem) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PathItem) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PathItem) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PathItem) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetGet

`func (o *PathItem) GetGet() OperationType2`

GetGet returns the Get field if non-nil, zero value otherwise.

### GetGetOk

`func (o *PathItem) GetGetOk() (*OperationType2, bool)`

GetGetOk returns a tuple with the Get field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGet

`func (o *PathItem) SetGet(v OperationType2)`

SetGet sets Get field to given value.

### HasGet

`func (o *PathItem) HasGet() bool`

HasGet returns a boolean if a field has been set.

### SetGetNil

`func (o *PathItem) SetGetNil(b bool)`

 SetGetNil sets the value for Get to be an explicit nil

### UnsetGet
`func (o *PathItem) UnsetGet()`

UnsetGet ensures that no value is present for Get, not even an explicit nil
### GetHead

`func (o *PathItem) GetHead() OperationType2`

GetHead returns the Head field if non-nil, zero value otherwise.

### GetHeadOk

`func (o *PathItem) GetHeadOk() (*OperationType2, bool)`

GetHeadOk returns a tuple with the Head field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHead

`func (o *PathItem) SetHead(v OperationType2)`

SetHead sets Head field to given value.

### HasHead

`func (o *PathItem) HasHead() bool`

HasHead returns a boolean if a field has been set.

### SetHeadNil

`func (o *PathItem) SetHeadNil(b bool)`

 SetHeadNil sets the value for Head to be an explicit nil

### UnsetHead
`func (o *PathItem) UnsetHead()`

UnsetHead ensures that no value is present for Head, not even an explicit nil
### GetOptions

`func (o *PathItem) GetOptions() OperationType2`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *PathItem) GetOptionsOk() (*OperationType2, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *PathItem) SetOptions(v OperationType2)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *PathItem) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### SetOptionsNil

`func (o *PathItem) SetOptionsNil(b bool)`

 SetOptionsNil sets the value for Options to be an explicit nil

### UnsetOptions
`func (o *PathItem) UnsetOptions()`

UnsetOptions ensures that no value is present for Options, not even an explicit nil
### GetParameters

`func (o *PathItem) GetParameters() []map[string]interface{}`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *PathItem) GetParametersOk() (*[]map[string]interface{}, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *PathItem) SetParameters(v []map[string]interface{})`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *PathItem) HasParameters() bool`

HasParameters returns a boolean if a field has been set.

### GetPatch

`func (o *PathItem) GetPatch() OperationType2`

GetPatch returns the Patch field if non-nil, zero value otherwise.

### GetPatchOk

`func (o *PathItem) GetPatchOk() (*OperationType2, bool)`

GetPatchOk returns a tuple with the Patch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatch

`func (o *PathItem) SetPatch(v OperationType2)`

SetPatch sets Patch field to given value.

### HasPatch

`func (o *PathItem) HasPatch() bool`

HasPatch returns a boolean if a field has been set.

### SetPatchNil

`func (o *PathItem) SetPatchNil(b bool)`

 SetPatchNil sets the value for Patch to be an explicit nil

### UnsetPatch
`func (o *PathItem) UnsetPatch()`

UnsetPatch ensures that no value is present for Patch, not even an explicit nil
### GetPost

`func (o *PathItem) GetPost() OperationType2`

GetPost returns the Post field if non-nil, zero value otherwise.

### GetPostOk

`func (o *PathItem) GetPostOk() (*OperationType2, bool)`

GetPostOk returns a tuple with the Post field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPost

`func (o *PathItem) SetPost(v OperationType2)`

SetPost sets Post field to given value.

### HasPost

`func (o *PathItem) HasPost() bool`

HasPost returns a boolean if a field has been set.

### SetPostNil

`func (o *PathItem) SetPostNil(b bool)`

 SetPostNil sets the value for Post to be an explicit nil

### UnsetPost
`func (o *PathItem) UnsetPost()`

UnsetPost ensures that no value is present for Post, not even an explicit nil
### GetPut

`func (o *PathItem) GetPut() OperationType2`

GetPut returns the Put field if non-nil, zero value otherwise.

### GetPutOk

`func (o *PathItem) GetPutOk() (*OperationType2, bool)`

GetPutOk returns a tuple with the Put field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPut

`func (o *PathItem) SetPut(v OperationType2)`

SetPut sets Put field to given value.

### HasPut

`func (o *PathItem) HasPut() bool`

HasPut returns a boolean if a field has been set.

### SetPutNil

`func (o *PathItem) SetPutNil(b bool)`

 SetPutNil sets the value for Put to be an explicit nil

### UnsetPut
`func (o *PathItem) UnsetPut()`

UnsetPut ensures that no value is present for Put, not even an explicit nil
### GetServers

`func (o *PathItem) GetServers() []ServerType2`

GetServers returns the Servers field if non-nil, zero value otherwise.

### GetServersOk

`func (o *PathItem) GetServersOk() (*[]ServerType2, bool)`

GetServersOk returns a tuple with the Servers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServers

`func (o *PathItem) SetServers(v []ServerType2)`

SetServers sets Servers field to given value.

### HasServers

`func (o *PathItem) HasServers() bool`

HasServers returns a boolean if a field has been set.

### SetServersNil

`func (o *PathItem) SetServersNil(b bool)`

 SetServersNil sets the value for Servers to be an explicit nil

### UnsetServers
`func (o *PathItem) UnsetServers()`

UnsetServers ensures that no value is present for Servers, not even an explicit nil
### GetSummary

`func (o *PathItem) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *PathItem) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *PathItem) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *PathItem) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetTrace

`func (o *PathItem) GetTrace() OperationType2`

GetTrace returns the Trace field if non-nil, zero value otherwise.

### GetTraceOk

`func (o *PathItem) GetTraceOk() (*OperationType2, bool)`

GetTraceOk returns a tuple with the Trace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrace

`func (o *PathItem) SetTrace(v OperationType2)`

SetTrace sets Trace field to given value.

### HasTrace

`func (o *PathItem) HasTrace() bool`

HasTrace returns a boolean if a field has been set.

### SetTraceNil

`func (o *PathItem) SetTraceNil(b bool)`

 SetTraceNil sets the value for Trace to be an explicit nil

### UnsetTrace
`func (o *PathItem) UnsetTrace()`

UnsetTrace ensures that no value is present for Trace, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


