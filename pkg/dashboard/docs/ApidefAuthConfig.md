# ApidefAuthConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthHeaderName** | Pointer to **string** |  | [optional] 
**CookieName** | Pointer to **string** |  | [optional] 
**DisableHeader** | Pointer to **bool** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**ParamName** | Pointer to **string** |  | [optional] 
**Signature** | Pointer to [**ApidefSignatureConfig**](ApidefSignatureConfig.md) |  | [optional] 
**UseCertificate** | Pointer to **bool** |  | [optional] 
**UseCookie** | Pointer to **bool** |  | [optional] 
**UseParam** | Pointer to **bool** |  | [optional] 
**ValidateSignature** | Pointer to **bool** |  | [optional] 

## Methods

### NewApidefAuthConfig

`func NewApidefAuthConfig() *ApidefAuthConfig`

NewApidefAuthConfig instantiates a new ApidefAuthConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApidefAuthConfigWithDefaults

`func NewApidefAuthConfigWithDefaults() *ApidefAuthConfig`

NewApidefAuthConfigWithDefaults instantiates a new ApidefAuthConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthHeaderName

`func (o *ApidefAuthConfig) GetAuthHeaderName() string`

GetAuthHeaderName returns the AuthHeaderName field if non-nil, zero value otherwise.

### GetAuthHeaderNameOk

`func (o *ApidefAuthConfig) GetAuthHeaderNameOk() (*string, bool)`

GetAuthHeaderNameOk returns a tuple with the AuthHeaderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthHeaderName

`func (o *ApidefAuthConfig) SetAuthHeaderName(v string)`

SetAuthHeaderName sets AuthHeaderName field to given value.

### HasAuthHeaderName

`func (o *ApidefAuthConfig) HasAuthHeaderName() bool`

HasAuthHeaderName returns a boolean if a field has been set.

### GetCookieName

`func (o *ApidefAuthConfig) GetCookieName() string`

GetCookieName returns the CookieName field if non-nil, zero value otherwise.

### GetCookieNameOk

`func (o *ApidefAuthConfig) GetCookieNameOk() (*string, bool)`

GetCookieNameOk returns a tuple with the CookieName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCookieName

`func (o *ApidefAuthConfig) SetCookieName(v string)`

SetCookieName sets CookieName field to given value.

### HasCookieName

`func (o *ApidefAuthConfig) HasCookieName() bool`

HasCookieName returns a boolean if a field has been set.

### GetDisableHeader

`func (o *ApidefAuthConfig) GetDisableHeader() bool`

GetDisableHeader returns the DisableHeader field if non-nil, zero value otherwise.

### GetDisableHeaderOk

`func (o *ApidefAuthConfig) GetDisableHeaderOk() (*bool, bool)`

GetDisableHeaderOk returns a tuple with the DisableHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableHeader

`func (o *ApidefAuthConfig) SetDisableHeader(v bool)`

SetDisableHeader sets DisableHeader field to given value.

### HasDisableHeader

`func (o *ApidefAuthConfig) HasDisableHeader() bool`

HasDisableHeader returns a boolean if a field has been set.

### GetName

`func (o *ApidefAuthConfig) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApidefAuthConfig) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApidefAuthConfig) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ApidefAuthConfig) HasName() bool`

HasName returns a boolean if a field has been set.

### GetParamName

`func (o *ApidefAuthConfig) GetParamName() string`

GetParamName returns the ParamName field if non-nil, zero value otherwise.

### GetParamNameOk

`func (o *ApidefAuthConfig) GetParamNameOk() (*string, bool)`

GetParamNameOk returns a tuple with the ParamName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParamName

`func (o *ApidefAuthConfig) SetParamName(v string)`

SetParamName sets ParamName field to given value.

### HasParamName

`func (o *ApidefAuthConfig) HasParamName() bool`

HasParamName returns a boolean if a field has been set.

### GetSignature

`func (o *ApidefAuthConfig) GetSignature() ApidefSignatureConfig`

GetSignature returns the Signature field if non-nil, zero value otherwise.

### GetSignatureOk

`func (o *ApidefAuthConfig) GetSignatureOk() (*ApidefSignatureConfig, bool)`

GetSignatureOk returns a tuple with the Signature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignature

`func (o *ApidefAuthConfig) SetSignature(v ApidefSignatureConfig)`

SetSignature sets Signature field to given value.

### HasSignature

`func (o *ApidefAuthConfig) HasSignature() bool`

HasSignature returns a boolean if a field has been set.

### GetUseCertificate

`func (o *ApidefAuthConfig) GetUseCertificate() bool`

GetUseCertificate returns the UseCertificate field if non-nil, zero value otherwise.

### GetUseCertificateOk

`func (o *ApidefAuthConfig) GetUseCertificateOk() (*bool, bool)`

GetUseCertificateOk returns a tuple with the UseCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseCertificate

`func (o *ApidefAuthConfig) SetUseCertificate(v bool)`

SetUseCertificate sets UseCertificate field to given value.

### HasUseCertificate

`func (o *ApidefAuthConfig) HasUseCertificate() bool`

HasUseCertificate returns a boolean if a field has been set.

### GetUseCookie

`func (o *ApidefAuthConfig) GetUseCookie() bool`

GetUseCookie returns the UseCookie field if non-nil, zero value otherwise.

### GetUseCookieOk

`func (o *ApidefAuthConfig) GetUseCookieOk() (*bool, bool)`

GetUseCookieOk returns a tuple with the UseCookie field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseCookie

`func (o *ApidefAuthConfig) SetUseCookie(v bool)`

SetUseCookie sets UseCookie field to given value.

### HasUseCookie

`func (o *ApidefAuthConfig) HasUseCookie() bool`

HasUseCookie returns a boolean if a field has been set.

### GetUseParam

`func (o *ApidefAuthConfig) GetUseParam() bool`

GetUseParam returns the UseParam field if non-nil, zero value otherwise.

### GetUseParamOk

`func (o *ApidefAuthConfig) GetUseParamOk() (*bool, bool)`

GetUseParamOk returns a tuple with the UseParam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseParam

`func (o *ApidefAuthConfig) SetUseParam(v bool)`

SetUseParam sets UseParam field to given value.

### HasUseParam

`func (o *ApidefAuthConfig) HasUseParam() bool`

HasUseParam returns a boolean if a field has been set.

### GetValidateSignature

`func (o *ApidefAuthConfig) GetValidateSignature() bool`

GetValidateSignature returns the ValidateSignature field if non-nil, zero value otherwise.

### GetValidateSignatureOk

`func (o *ApidefAuthConfig) GetValidateSignatureOk() (*bool, bool)`

GetValidateSignatureOk returns a tuple with the ValidateSignature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidateSignature

`func (o *ApidefAuthConfig) SetValidateSignature(v bool)`

SetValidateSignature sets ValidateSignature field to given value.

### HasValidateSignature

`func (o *ApidefAuthConfig) HasValidateSignature() bool`

HasValidateSignature returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


