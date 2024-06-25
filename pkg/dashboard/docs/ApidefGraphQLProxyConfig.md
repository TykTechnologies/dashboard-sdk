# ApidefGraphQLProxyConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthHeaders** | Pointer to **map[string]string** |  | [optional] 
**Features** | Pointer to [**ApidefGraphQLProxyFeaturesConfig**](ApidefGraphQLProxyFeaturesConfig.md) |  | [optional] 
**RequestHeaders** | Pointer to **map[string]string** |  | [optional] 
**RequestHeadersRewrite** | Pointer to [**map[string]ApidefRequestHeadersRewriteConfig**](ApidefRequestHeadersRewriteConfig.md) |  | [optional] 
**SubscriptionType** | Pointer to **string** |  | [optional] 
**UseResponseExtensions** | Pointer to [**ApidefGraphQLResponseExtensions**](ApidefGraphQLResponseExtensions.md) |  | [optional] 

## Methods

### NewApidefGraphQLProxyConfig

`func NewApidefGraphQLProxyConfig() *ApidefGraphQLProxyConfig`

NewApidefGraphQLProxyConfig instantiates a new ApidefGraphQLProxyConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApidefGraphQLProxyConfigWithDefaults

`func NewApidefGraphQLProxyConfigWithDefaults() *ApidefGraphQLProxyConfig`

NewApidefGraphQLProxyConfigWithDefaults instantiates a new ApidefGraphQLProxyConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthHeaders

`func (o *ApidefGraphQLProxyConfig) GetAuthHeaders() map[string]string`

GetAuthHeaders returns the AuthHeaders field if non-nil, zero value otherwise.

### GetAuthHeadersOk

`func (o *ApidefGraphQLProxyConfig) GetAuthHeadersOk() (*map[string]string, bool)`

GetAuthHeadersOk returns a tuple with the AuthHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthHeaders

`func (o *ApidefGraphQLProxyConfig) SetAuthHeaders(v map[string]string)`

SetAuthHeaders sets AuthHeaders field to given value.

### HasAuthHeaders

`func (o *ApidefGraphQLProxyConfig) HasAuthHeaders() bool`

HasAuthHeaders returns a boolean if a field has been set.

### SetAuthHeadersNil

`func (o *ApidefGraphQLProxyConfig) SetAuthHeadersNil(b bool)`

 SetAuthHeadersNil sets the value for AuthHeaders to be an explicit nil

### UnsetAuthHeaders
`func (o *ApidefGraphQLProxyConfig) UnsetAuthHeaders()`

UnsetAuthHeaders ensures that no value is present for AuthHeaders, not even an explicit nil
### GetFeatures

`func (o *ApidefGraphQLProxyConfig) GetFeatures() ApidefGraphQLProxyFeaturesConfig`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *ApidefGraphQLProxyConfig) GetFeaturesOk() (*ApidefGraphQLProxyFeaturesConfig, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *ApidefGraphQLProxyConfig) SetFeatures(v ApidefGraphQLProxyFeaturesConfig)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *ApidefGraphQLProxyConfig) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.

### GetRequestHeaders

`func (o *ApidefGraphQLProxyConfig) GetRequestHeaders() map[string]string`

GetRequestHeaders returns the RequestHeaders field if non-nil, zero value otherwise.

### GetRequestHeadersOk

`func (o *ApidefGraphQLProxyConfig) GetRequestHeadersOk() (*map[string]string, bool)`

GetRequestHeadersOk returns a tuple with the RequestHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestHeaders

`func (o *ApidefGraphQLProxyConfig) SetRequestHeaders(v map[string]string)`

SetRequestHeaders sets RequestHeaders field to given value.

### HasRequestHeaders

`func (o *ApidefGraphQLProxyConfig) HasRequestHeaders() bool`

HasRequestHeaders returns a boolean if a field has been set.

### SetRequestHeadersNil

`func (o *ApidefGraphQLProxyConfig) SetRequestHeadersNil(b bool)`

 SetRequestHeadersNil sets the value for RequestHeaders to be an explicit nil

### UnsetRequestHeaders
`func (o *ApidefGraphQLProxyConfig) UnsetRequestHeaders()`

UnsetRequestHeaders ensures that no value is present for RequestHeaders, not even an explicit nil
### GetRequestHeadersRewrite

`func (o *ApidefGraphQLProxyConfig) GetRequestHeadersRewrite() map[string]ApidefRequestHeadersRewriteConfig`

GetRequestHeadersRewrite returns the RequestHeadersRewrite field if non-nil, zero value otherwise.

### GetRequestHeadersRewriteOk

`func (o *ApidefGraphQLProxyConfig) GetRequestHeadersRewriteOk() (*map[string]ApidefRequestHeadersRewriteConfig, bool)`

GetRequestHeadersRewriteOk returns a tuple with the RequestHeadersRewrite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestHeadersRewrite

`func (o *ApidefGraphQLProxyConfig) SetRequestHeadersRewrite(v map[string]ApidefRequestHeadersRewriteConfig)`

SetRequestHeadersRewrite sets RequestHeadersRewrite field to given value.

### HasRequestHeadersRewrite

`func (o *ApidefGraphQLProxyConfig) HasRequestHeadersRewrite() bool`

HasRequestHeadersRewrite returns a boolean if a field has been set.

### SetRequestHeadersRewriteNil

`func (o *ApidefGraphQLProxyConfig) SetRequestHeadersRewriteNil(b bool)`

 SetRequestHeadersRewriteNil sets the value for RequestHeadersRewrite to be an explicit nil

### UnsetRequestHeadersRewrite
`func (o *ApidefGraphQLProxyConfig) UnsetRequestHeadersRewrite()`

UnsetRequestHeadersRewrite ensures that no value is present for RequestHeadersRewrite, not even an explicit nil
### GetSubscriptionType

`func (o *ApidefGraphQLProxyConfig) GetSubscriptionType() string`

GetSubscriptionType returns the SubscriptionType field if non-nil, zero value otherwise.

### GetSubscriptionTypeOk

`func (o *ApidefGraphQLProxyConfig) GetSubscriptionTypeOk() (*string, bool)`

GetSubscriptionTypeOk returns a tuple with the SubscriptionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionType

`func (o *ApidefGraphQLProxyConfig) SetSubscriptionType(v string)`

SetSubscriptionType sets SubscriptionType field to given value.

### HasSubscriptionType

`func (o *ApidefGraphQLProxyConfig) HasSubscriptionType() bool`

HasSubscriptionType returns a boolean if a field has been set.

### GetUseResponseExtensions

`func (o *ApidefGraphQLProxyConfig) GetUseResponseExtensions() ApidefGraphQLResponseExtensions`

GetUseResponseExtensions returns the UseResponseExtensions field if non-nil, zero value otherwise.

### GetUseResponseExtensionsOk

`func (o *ApidefGraphQLProxyConfig) GetUseResponseExtensionsOk() (*ApidefGraphQLResponseExtensions, bool)`

GetUseResponseExtensionsOk returns a tuple with the UseResponseExtensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseResponseExtensions

`func (o *ApidefGraphQLProxyConfig) SetUseResponseExtensions(v ApidefGraphQLResponseExtensions)`

SetUseResponseExtensions sets UseResponseExtensions field to given value.

### HasUseResponseExtensions

`func (o *ApidefGraphQLProxyConfig) HasUseResponseExtensions() bool`

HasUseResponseExtensions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


