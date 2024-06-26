# OasMockResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Body** | Pointer to **string** |  | [optional] 
**Code** | Pointer to **int32** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**FromOASExamples** | Pointer to [**OasFromOASExamples**](OasFromOASExamples.md) |  | [optional] 
**Headers** | Pointer to [**[]OasHeader**](OasHeader.md) |  | [optional] 

## Methods

### NewOasMockResponse

`func NewOasMockResponse() *OasMockResponse`

NewOasMockResponse instantiates a new OasMockResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOasMockResponseWithDefaults

`func NewOasMockResponseWithDefaults() *OasMockResponse`

NewOasMockResponseWithDefaults instantiates a new OasMockResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBody

`func (o *OasMockResponse) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *OasMockResponse) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *OasMockResponse) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *OasMockResponse) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetCode

`func (o *OasMockResponse) GetCode() int32`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *OasMockResponse) GetCodeOk() (*int32, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *OasMockResponse) SetCode(v int32)`

SetCode sets Code field to given value.

### HasCode

`func (o *OasMockResponse) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetEnabled

`func (o *OasMockResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *OasMockResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *OasMockResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *OasMockResponse) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetFromOASExamples

`func (o *OasMockResponse) GetFromOASExamples() OasFromOASExamples`

GetFromOASExamples returns the FromOASExamples field if non-nil, zero value otherwise.

### GetFromOASExamplesOk

`func (o *OasMockResponse) GetFromOASExamplesOk() (*OasFromOASExamples, bool)`

GetFromOASExamplesOk returns a tuple with the FromOASExamples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromOASExamples

`func (o *OasMockResponse) SetFromOASExamples(v OasFromOASExamples)`

SetFromOASExamples sets FromOASExamples field to given value.

### HasFromOASExamples

`func (o *OasMockResponse) HasFromOASExamples() bool`

HasFromOASExamples returns a boolean if a field has been set.

### GetHeaders

`func (o *OasMockResponse) GetHeaders() []OasHeader`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *OasMockResponse) GetHeadersOk() (*[]OasHeader, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *OasMockResponse) SetHeaders(v []OasHeader)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *OasMockResponse) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


