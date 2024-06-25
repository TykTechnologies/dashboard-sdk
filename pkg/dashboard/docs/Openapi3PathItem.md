# Openapi3PathItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** |  | [optional] 
**Connect** | Pointer to [**Openapi3Operation**](Openapi3Operation.md) |  | [optional] 
**Delete** | Pointer to [**Openapi3Operation**](Openapi3Operation.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Get** | Pointer to [**Openapi3Operation**](Openapi3Operation.md) |  | [optional] 
**Head** | Pointer to [**Openapi3Operation**](Openapi3Operation.md) |  | [optional] 
**Options** | Pointer to [**Openapi3Operation**](Openapi3Operation.md) |  | [optional] 
**Parameters** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Patch** | Pointer to [**Openapi3Operation**](Openapi3Operation.md) |  | [optional] 
**Post** | Pointer to [**Openapi3Operation**](Openapi3Operation.md) |  | [optional] 
**Put** | Pointer to [**Openapi3Operation**](Openapi3Operation.md) |  | [optional] 
**Servers** | Pointer to [**[]Openapi3Server**](Openapi3Server.md) |  | [optional] 
**Summary** | Pointer to **string** |  | [optional] 
**Trace** | Pointer to [**Openapi3Operation**](Openapi3Operation.md) |  | [optional] 

## Methods

### NewOpenapi3PathItem

`func NewOpenapi3PathItem() *Openapi3PathItem`

NewOpenapi3PathItem instantiates a new Openapi3PathItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpenapi3PathItemWithDefaults

`func NewOpenapi3PathItemWithDefaults() *Openapi3PathItem`

NewOpenapi3PathItemWithDefaults instantiates a new Openapi3PathItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *Openapi3PathItem) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *Openapi3PathItem) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *Openapi3PathItem) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *Openapi3PathItem) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetConnect

`func (o *Openapi3PathItem) GetConnect() Openapi3Operation`

GetConnect returns the Connect field if non-nil, zero value otherwise.

### GetConnectOk

`func (o *Openapi3PathItem) GetConnectOk() (*Openapi3Operation, bool)`

GetConnectOk returns a tuple with the Connect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnect

`func (o *Openapi3PathItem) SetConnect(v Openapi3Operation)`

SetConnect sets Connect field to given value.

### HasConnect

`func (o *Openapi3PathItem) HasConnect() bool`

HasConnect returns a boolean if a field has been set.

### GetDelete

`func (o *Openapi3PathItem) GetDelete() Openapi3Operation`

GetDelete returns the Delete field if non-nil, zero value otherwise.

### GetDeleteOk

`func (o *Openapi3PathItem) GetDeleteOk() (*Openapi3Operation, bool)`

GetDeleteOk returns a tuple with the Delete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelete

`func (o *Openapi3PathItem) SetDelete(v Openapi3Operation)`

SetDelete sets Delete field to given value.

### HasDelete

`func (o *Openapi3PathItem) HasDelete() bool`

HasDelete returns a boolean if a field has been set.

### GetDescription

`func (o *Openapi3PathItem) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Openapi3PathItem) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Openapi3PathItem) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Openapi3PathItem) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetGet

`func (o *Openapi3PathItem) GetGet() Openapi3Operation`

GetGet returns the Get field if non-nil, zero value otherwise.

### GetGetOk

`func (o *Openapi3PathItem) GetGetOk() (*Openapi3Operation, bool)`

GetGetOk returns a tuple with the Get field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGet

`func (o *Openapi3PathItem) SetGet(v Openapi3Operation)`

SetGet sets Get field to given value.

### HasGet

`func (o *Openapi3PathItem) HasGet() bool`

HasGet returns a boolean if a field has been set.

### GetHead

`func (o *Openapi3PathItem) GetHead() Openapi3Operation`

GetHead returns the Head field if non-nil, zero value otherwise.

### GetHeadOk

`func (o *Openapi3PathItem) GetHeadOk() (*Openapi3Operation, bool)`

GetHeadOk returns a tuple with the Head field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHead

`func (o *Openapi3PathItem) SetHead(v Openapi3Operation)`

SetHead sets Head field to given value.

### HasHead

`func (o *Openapi3PathItem) HasHead() bool`

HasHead returns a boolean if a field has been set.

### GetOptions

`func (o *Openapi3PathItem) GetOptions() Openapi3Operation`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *Openapi3PathItem) GetOptionsOk() (*Openapi3Operation, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *Openapi3PathItem) SetOptions(v Openapi3Operation)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *Openapi3PathItem) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetParameters

`func (o *Openapi3PathItem) GetParameters() []map[string]interface{}`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *Openapi3PathItem) GetParametersOk() (*[]map[string]interface{}, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *Openapi3PathItem) SetParameters(v []map[string]interface{})`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *Openapi3PathItem) HasParameters() bool`

HasParameters returns a boolean if a field has been set.

### GetPatch

`func (o *Openapi3PathItem) GetPatch() Openapi3Operation`

GetPatch returns the Patch field if non-nil, zero value otherwise.

### GetPatchOk

`func (o *Openapi3PathItem) GetPatchOk() (*Openapi3Operation, bool)`

GetPatchOk returns a tuple with the Patch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatch

`func (o *Openapi3PathItem) SetPatch(v Openapi3Operation)`

SetPatch sets Patch field to given value.

### HasPatch

`func (o *Openapi3PathItem) HasPatch() bool`

HasPatch returns a boolean if a field has been set.

### GetPost

`func (o *Openapi3PathItem) GetPost() Openapi3Operation`

GetPost returns the Post field if non-nil, zero value otherwise.

### GetPostOk

`func (o *Openapi3PathItem) GetPostOk() (*Openapi3Operation, bool)`

GetPostOk returns a tuple with the Post field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPost

`func (o *Openapi3PathItem) SetPost(v Openapi3Operation)`

SetPost sets Post field to given value.

### HasPost

`func (o *Openapi3PathItem) HasPost() bool`

HasPost returns a boolean if a field has been set.

### GetPut

`func (o *Openapi3PathItem) GetPut() Openapi3Operation`

GetPut returns the Put field if non-nil, zero value otherwise.

### GetPutOk

`func (o *Openapi3PathItem) GetPutOk() (*Openapi3Operation, bool)`

GetPutOk returns a tuple with the Put field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPut

`func (o *Openapi3PathItem) SetPut(v Openapi3Operation)`

SetPut sets Put field to given value.

### HasPut

`func (o *Openapi3PathItem) HasPut() bool`

HasPut returns a boolean if a field has been set.

### GetServers

`func (o *Openapi3PathItem) GetServers() []Openapi3Server`

GetServers returns the Servers field if non-nil, zero value otherwise.

### GetServersOk

`func (o *Openapi3PathItem) GetServersOk() (*[]Openapi3Server, bool)`

GetServersOk returns a tuple with the Servers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServers

`func (o *Openapi3PathItem) SetServers(v []Openapi3Server)`

SetServers sets Servers field to given value.

### HasServers

`func (o *Openapi3PathItem) HasServers() bool`

HasServers returns a boolean if a field has been set.

### SetServersNil

`func (o *Openapi3PathItem) SetServersNil(b bool)`

 SetServersNil sets the value for Servers to be an explicit nil

### UnsetServers
`func (o *Openapi3PathItem) UnsetServers()`

UnsetServers ensures that no value is present for Servers, not even an explicit nil
### GetSummary

`func (o *Openapi3PathItem) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *Openapi3PathItem) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *Openapi3PathItem) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *Openapi3PathItem) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetTrace

`func (o *Openapi3PathItem) GetTrace() Openapi3Operation`

GetTrace returns the Trace field if non-nil, zero value otherwise.

### GetTraceOk

`func (o *Openapi3PathItem) GetTraceOk() (*Openapi3Operation, bool)`

GetTraceOk returns a tuple with the Trace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrace

`func (o *Openapi3PathItem) SetTrace(v Openapi3Operation)`

SetTrace sets Trace field to given value.

### HasTrace

`func (o *Openapi3PathItem) HasTrace() bool`

HasTrace returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


