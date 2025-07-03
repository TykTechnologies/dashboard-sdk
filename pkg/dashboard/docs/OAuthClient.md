# OAuthClient

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | Pointer to **string** |  | [optional] 
**ClientName** | Pointer to **string** |  | [optional] 
**DcrRegistration** | Pointer to [**DcrpRegistration**](DcrpRegistration.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**MetaData** | Pointer to **map[string]string** |  | [optional] 
**PolicyId** | Pointer to **string** |  | [optional] 
**RedirectUri** | Pointer to **string** |  | [optional] 
**Secret** | Pointer to **string** |  | [optional] 

## Methods

### NewOAuthClient

`func NewOAuthClient() *OAuthClient`

NewOAuthClient instantiates a new OAuthClient object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOAuthClientWithDefaults

`func NewOAuthClientWithDefaults() *OAuthClient`

NewOAuthClientWithDefaults instantiates a new OAuthClient object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *OAuthClient) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *OAuthClient) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *OAuthClient) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *OAuthClient) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetClientName

`func (o *OAuthClient) GetClientName() string`

GetClientName returns the ClientName field if non-nil, zero value otherwise.

### GetClientNameOk

`func (o *OAuthClient) GetClientNameOk() (*string, bool)`

GetClientNameOk returns a tuple with the ClientName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientName

`func (o *OAuthClient) SetClientName(v string)`

SetClientName sets ClientName field to given value.

### HasClientName

`func (o *OAuthClient) HasClientName() bool`

HasClientName returns a boolean if a field has been set.

### GetDcrRegistration

`func (o *OAuthClient) GetDcrRegistration() DcrpRegistration`

GetDcrRegistration returns the DcrRegistration field if non-nil, zero value otherwise.

### GetDcrRegistrationOk

`func (o *OAuthClient) GetDcrRegistrationOk() (*DcrpRegistration, bool)`

GetDcrRegistrationOk returns a tuple with the DcrRegistration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDcrRegistration

`func (o *OAuthClient) SetDcrRegistration(v DcrpRegistration)`

SetDcrRegistration sets DcrRegistration field to given value.

### HasDcrRegistration

`func (o *OAuthClient) HasDcrRegistration() bool`

HasDcrRegistration returns a boolean if a field has been set.

### GetDescription

`func (o *OAuthClient) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *OAuthClient) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *OAuthClient) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *OAuthClient) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMetaData

`func (o *OAuthClient) GetMetaData() map[string]string`

GetMetaData returns the MetaData field if non-nil, zero value otherwise.

### GetMetaDataOk

`func (o *OAuthClient) GetMetaDataOk() (*map[string]string, bool)`

GetMetaDataOk returns a tuple with the MetaData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaData

`func (o *OAuthClient) SetMetaData(v map[string]string)`

SetMetaData sets MetaData field to given value.

### HasMetaData

`func (o *OAuthClient) HasMetaData() bool`

HasMetaData returns a boolean if a field has been set.

### GetPolicyId

`func (o *OAuthClient) GetPolicyId() string`

GetPolicyId returns the PolicyId field if non-nil, zero value otherwise.

### GetPolicyIdOk

`func (o *OAuthClient) GetPolicyIdOk() (*string, bool)`

GetPolicyIdOk returns a tuple with the PolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyId

`func (o *OAuthClient) SetPolicyId(v string)`

SetPolicyId sets PolicyId field to given value.

### HasPolicyId

`func (o *OAuthClient) HasPolicyId() bool`

HasPolicyId returns a boolean if a field has been set.

### GetRedirectUri

`func (o *OAuthClient) GetRedirectUri() string`

GetRedirectUri returns the RedirectUri field if non-nil, zero value otherwise.

### GetRedirectUriOk

`func (o *OAuthClient) GetRedirectUriOk() (*string, bool)`

GetRedirectUriOk returns a tuple with the RedirectUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUri

`func (o *OAuthClient) SetRedirectUri(v string)`

SetRedirectUri sets RedirectUri field to given value.

### HasRedirectUri

`func (o *OAuthClient) HasRedirectUri() bool`

HasRedirectUri returns a boolean if a field has been set.

### GetSecret

`func (o *OAuthClient) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *OAuthClient) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *OAuthClient) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *OAuthClient) HasSecret() bool`

HasSecret returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


