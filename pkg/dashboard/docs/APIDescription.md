# APIDescription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiId** | Pointer to **string** |  | [optional] 
**AuthType** | Pointer to **string** |  | [optional] 
**Config** | Pointer to [**PortalConfig**](PortalConfig.md) |  | [optional] 
**Documentation** | Pointer to **string** |  | [optional] 
**Fields** | Pointer to **map[string]string** |  | [optional] 
**IsKeyless** | Pointer to **bool** |  | [optional] 
**LongDescription** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**PolicyId** | Pointer to **string** |  | [optional] 
**ShortDescription** | Pointer to **string** |  | [optional] 
**Show** | Pointer to **bool** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 

## Methods

### NewAPIDescription

`func NewAPIDescription() *APIDescription`

NewAPIDescription instantiates a new APIDescription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAPIDescriptionWithDefaults

`func NewAPIDescriptionWithDefaults() *APIDescription`

NewAPIDescriptionWithDefaults instantiates a new APIDescription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiId

`func (o *APIDescription) GetApiId() string`

GetApiId returns the ApiId field if non-nil, zero value otherwise.

### GetApiIdOk

`func (o *APIDescription) GetApiIdOk() (*string, bool)`

GetApiIdOk returns a tuple with the ApiId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiId

`func (o *APIDescription) SetApiId(v string)`

SetApiId sets ApiId field to given value.

### HasApiId

`func (o *APIDescription) HasApiId() bool`

HasApiId returns a boolean if a field has been set.

### GetAuthType

`func (o *APIDescription) GetAuthType() string`

GetAuthType returns the AuthType field if non-nil, zero value otherwise.

### GetAuthTypeOk

`func (o *APIDescription) GetAuthTypeOk() (*string, bool)`

GetAuthTypeOk returns a tuple with the AuthType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthType

`func (o *APIDescription) SetAuthType(v string)`

SetAuthType sets AuthType field to given value.

### HasAuthType

`func (o *APIDescription) HasAuthType() bool`

HasAuthType returns a boolean if a field has been set.

### GetConfig

`func (o *APIDescription) GetConfig() PortalConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *APIDescription) GetConfigOk() (*PortalConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *APIDescription) SetConfig(v PortalConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *APIDescription) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetDocumentation

`func (o *APIDescription) GetDocumentation() string`

GetDocumentation returns the Documentation field if non-nil, zero value otherwise.

### GetDocumentationOk

`func (o *APIDescription) GetDocumentationOk() (*string, bool)`

GetDocumentationOk returns a tuple with the Documentation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentation

`func (o *APIDescription) SetDocumentation(v string)`

SetDocumentation sets Documentation field to given value.

### HasDocumentation

`func (o *APIDescription) HasDocumentation() bool`

HasDocumentation returns a boolean if a field has been set.

### GetFields

`func (o *APIDescription) GetFields() map[string]string`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *APIDescription) GetFieldsOk() (*map[string]string, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *APIDescription) SetFields(v map[string]string)`

SetFields sets Fields field to given value.

### HasFields

`func (o *APIDescription) HasFields() bool`

HasFields returns a boolean if a field has been set.

### SetFieldsNil

`func (o *APIDescription) SetFieldsNil(b bool)`

 SetFieldsNil sets the value for Fields to be an explicit nil

### UnsetFields
`func (o *APIDescription) UnsetFields()`

UnsetFields ensures that no value is present for Fields, not even an explicit nil
### GetIsKeyless

`func (o *APIDescription) GetIsKeyless() bool`

GetIsKeyless returns the IsKeyless field if non-nil, zero value otherwise.

### GetIsKeylessOk

`func (o *APIDescription) GetIsKeylessOk() (*bool, bool)`

GetIsKeylessOk returns a tuple with the IsKeyless field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsKeyless

`func (o *APIDescription) SetIsKeyless(v bool)`

SetIsKeyless sets IsKeyless field to given value.

### HasIsKeyless

`func (o *APIDescription) HasIsKeyless() bool`

HasIsKeyless returns a boolean if a field has been set.

### GetLongDescription

`func (o *APIDescription) GetLongDescription() string`

GetLongDescription returns the LongDescription field if non-nil, zero value otherwise.

### GetLongDescriptionOk

`func (o *APIDescription) GetLongDescriptionOk() (*string, bool)`

GetLongDescriptionOk returns a tuple with the LongDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongDescription

`func (o *APIDescription) SetLongDescription(v string)`

SetLongDescription sets LongDescription field to given value.

### HasLongDescription

`func (o *APIDescription) HasLongDescription() bool`

HasLongDescription returns a boolean if a field has been set.

### GetName

`func (o *APIDescription) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *APIDescription) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *APIDescription) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *APIDescription) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPolicyId

`func (o *APIDescription) GetPolicyId() string`

GetPolicyId returns the PolicyId field if non-nil, zero value otherwise.

### GetPolicyIdOk

`func (o *APIDescription) GetPolicyIdOk() (*string, bool)`

GetPolicyIdOk returns a tuple with the PolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyId

`func (o *APIDescription) SetPolicyId(v string)`

SetPolicyId sets PolicyId field to given value.

### HasPolicyId

`func (o *APIDescription) HasPolicyId() bool`

HasPolicyId returns a boolean if a field has been set.

### GetShortDescription

`func (o *APIDescription) GetShortDescription() string`

GetShortDescription returns the ShortDescription field if non-nil, zero value otherwise.

### GetShortDescriptionOk

`func (o *APIDescription) GetShortDescriptionOk() (*string, bool)`

GetShortDescriptionOk returns a tuple with the ShortDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortDescription

`func (o *APIDescription) SetShortDescription(v string)`

SetShortDescription sets ShortDescription field to given value.

### HasShortDescription

`func (o *APIDescription) HasShortDescription() bool`

HasShortDescription returns a boolean if a field has been set.

### GetShow

`func (o *APIDescription) GetShow() bool`

GetShow returns the Show field if non-nil, zero value otherwise.

### GetShowOk

`func (o *APIDescription) GetShowOk() (*bool, bool)`

GetShowOk returns a tuple with the Show field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShow

`func (o *APIDescription) SetShow(v bool)`

SetShow sets Show field to given value.

### HasShow

`func (o *APIDescription) HasShow() bool`

HasShow returns a boolean if a field has been set.

### GetVersion

`func (o *APIDescription) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *APIDescription) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *APIDescription) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *APIDescription) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


