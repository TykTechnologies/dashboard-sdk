# OasAuthenticationPlugin

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** |  | [optional] 
**FunctionName** | Pointer to **string** |  | [optional] 
**IdExtractor** | Pointer to [**OasIDExtractor**](OasIDExtractor.md) |  | [optional] 
**Path** | Pointer to **string** |  | [optional] 
**RawBodyOnly** | Pointer to **bool** |  | [optional] 

## Methods

### NewOasAuthenticationPlugin

`func NewOasAuthenticationPlugin() *OasAuthenticationPlugin`

NewOasAuthenticationPlugin instantiates a new OasAuthenticationPlugin object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOasAuthenticationPluginWithDefaults

`func NewOasAuthenticationPluginWithDefaults() *OasAuthenticationPlugin`

NewOasAuthenticationPluginWithDefaults instantiates a new OasAuthenticationPlugin object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *OasAuthenticationPlugin) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *OasAuthenticationPlugin) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *OasAuthenticationPlugin) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *OasAuthenticationPlugin) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetFunctionName

`func (o *OasAuthenticationPlugin) GetFunctionName() string`

GetFunctionName returns the FunctionName field if non-nil, zero value otherwise.

### GetFunctionNameOk

`func (o *OasAuthenticationPlugin) GetFunctionNameOk() (*string, bool)`

GetFunctionNameOk returns a tuple with the FunctionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionName

`func (o *OasAuthenticationPlugin) SetFunctionName(v string)`

SetFunctionName sets FunctionName field to given value.

### HasFunctionName

`func (o *OasAuthenticationPlugin) HasFunctionName() bool`

HasFunctionName returns a boolean if a field has been set.

### GetIdExtractor

`func (o *OasAuthenticationPlugin) GetIdExtractor() OasIDExtractor`

GetIdExtractor returns the IdExtractor field if non-nil, zero value otherwise.

### GetIdExtractorOk

`func (o *OasAuthenticationPlugin) GetIdExtractorOk() (*OasIDExtractor, bool)`

GetIdExtractorOk returns a tuple with the IdExtractor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdExtractor

`func (o *OasAuthenticationPlugin) SetIdExtractor(v OasIDExtractor)`

SetIdExtractor sets IdExtractor field to given value.

### HasIdExtractor

`func (o *OasAuthenticationPlugin) HasIdExtractor() bool`

HasIdExtractor returns a boolean if a field has been set.

### GetPath

`func (o *OasAuthenticationPlugin) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *OasAuthenticationPlugin) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *OasAuthenticationPlugin) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *OasAuthenticationPlugin) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetRawBodyOnly

`func (o *OasAuthenticationPlugin) GetRawBodyOnly() bool`

GetRawBodyOnly returns the RawBodyOnly field if non-nil, zero value otherwise.

### GetRawBodyOnlyOk

`func (o *OasAuthenticationPlugin) GetRawBodyOnlyOk() (*bool, bool)`

GetRawBodyOnlyOk returns a tuple with the RawBodyOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawBodyOnly

`func (o *OasAuthenticationPlugin) SetRawBodyOnly(v bool)`

SetRawBodyOnly sets RawBodyOnly field to given value.

### HasRawBodyOnly

`func (o *OasAuthenticationPlugin) HasRawBodyOnly() bool`

HasRawBodyOnly returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


