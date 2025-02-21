# APIDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CORS** | Pointer to [**CORSConfig**](CORSConfig.md) |  | [optional] 
**Active** | Pointer to **bool** |  | [optional] 
**AllowedIps** | Pointer to **[]string** |  | [optional] 
**AnalyticsPlugin** | Pointer to [**AnalyticsPluginConfig**](AnalyticsPluginConfig.md) |  | [optional] 
**ApiId** | Pointer to **string** |  | [optional] 
**Auth** | Pointer to [**AuthConfig**](AuthConfig.md) |  | [optional] 
**AuthConfigs** | Pointer to [**map[string]AuthConfig**](AuthConfig.md) |  | [optional] 
**AuthProvider** | Pointer to [**AuthProviderMeta**](AuthProviderMeta.md) |  | [optional] 
**BaseIdentityProvidedBy** | Pointer to **string** |  | [optional] 
**BasicAuth** | Pointer to [**APIDefinitionBasicAuth**](APIDefinitionBasicAuth.md) |  | [optional] 
**BlacklistedIps** | Pointer to **[]string** |  | [optional] 
**CacheOptions** | Pointer to [**CacheOptions**](CacheOptions.md) |  | [optional] 
**CertificatePinningDisabled** | Pointer to **bool** |  | [optional] 
**Certificates** | Pointer to **[]string** |  | [optional] 
**ClientCertificates** | Pointer to **[]string** |  | [optional] 
**ConfigData** | Pointer to **map[string]interface{}** |  | [optional] 
**ConfigDataDisabled** | Pointer to **bool** |  | [optional] 
**CustomMiddleware** | Pointer to [**MiddlewareSection**](MiddlewareSection.md) |  | [optional] 
**CustomMiddlewareBundle** | Pointer to **string** |  | [optional] 
**CustomMiddlewareBundleDisabled** | Pointer to **bool** |  | [optional] 
**CustomPluginAuthEnabled** | Pointer to **bool** |  | [optional] 
**Definition** | Pointer to [**VersionDefinition**](VersionDefinition.md) |  | [optional] 
**DetailedTracing** | Pointer to **bool** |  | [optional] 
**DisableQuota** | Pointer to **bool** |  | [optional] 
**DisableRateLimit** | Pointer to **bool** |  | [optional] 
**DoNotTrack** | Pointer to **bool** |  | [optional] 
**Domain** | Pointer to **string** |  | [optional] 
**DomainDisabled** | Pointer to **bool** |  | [optional] 
**DontSetQuotaOnCreate** | Pointer to **bool** |  | [optional] 
**EnableBatchRequestSupport** | Pointer to **bool** |  | [optional] 
**EnableContextVars** | Pointer to **bool** |  | [optional] 
**EnableCoprocessAuth** | Pointer to **bool** |  | [optional] 
**EnableDetailedRecording** | Pointer to **bool** |  | [optional] 
**EnableIpBlacklisting** | Pointer to **bool** |  | [optional] 
**EnableIpWhitelisting** | Pointer to **bool** |  | [optional] 
**EnableJwt** | Pointer to **bool** |  | [optional] 
**EnableProxyProtocol** | Pointer to **bool** |  | [optional] 
**EnableSignatureChecking** | Pointer to **bool** |  | [optional] 
**EventHandlers** | Pointer to [**EventHandlerMetaConfig**](EventHandlerMetaConfig.md) |  | [optional] 
**Expiration** | Pointer to **string** |  | [optional] 
**ExpireAnalyticsAfter** | Pointer to **int64** |  | [optional] 
**ExternalOauth** | Pointer to [**ExternalOAuth**](ExternalOAuth.md) |  | [optional] 
**GlobalRateLimit** | Pointer to [**GlobalRateLimit**](GlobalRateLimit.md) |  | [optional] 
**Graphql** | Pointer to [**GraphQLConfig**](GraphQLConfig.md) |  | [optional] 
**HmacAllowedAlgorithms** | Pointer to **[]string** |  | [optional] 
**HmacAllowedClockSkew** | Pointer to **float64** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**IdpClientIdMappingDisabled** | Pointer to **bool** |  | [optional] 
**Internal** | Pointer to **bool** |  | [optional] 
**IsOas** | Pointer to **bool** |  | [optional] 
**JwtClientBaseField** | Pointer to **string** |  | [optional] 
**JwtDefaultPolicies** | Pointer to **[]string** |  | [optional] 
**JwtExpiresAtValidationSkew** | Pointer to **int32** |  | [optional] 
**JwtIdentityBaseField** | Pointer to **string** |  | [optional] 
**JwtIssuedAtValidationSkew** | Pointer to **int32** |  | [optional] 
**JwtNotBeforeValidationSkew** | Pointer to **int32** |  | [optional] 
**JwtPolicyFieldName** | Pointer to **string** |  | [optional] 
**JwtScopeClaimName** | Pointer to **string** |  | [optional] 
**JwtScopeToPolicyMapping** | Pointer to **map[string]string** |  | [optional] 
**JwtSigningMethod** | Pointer to **string** |  | [optional] 
**JwtSkipKid** | Pointer to **bool** |  | [optional] 
**JwtSource** | Pointer to **string** |  | [optional] 
**ListenPort** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Notifications** | Pointer to [**NotificationsManager**](NotificationsManager.md) |  | [optional] 
**OauthMeta** | Pointer to [**APIDefinitionOauthMeta**](APIDefinitionOauthMeta.md) |  | [optional] 
**OpenidOptions** | Pointer to [**OpenIDOptions**](OpenIDOptions.md) |  | [optional] 
**OrgId** | Pointer to **string** |  | [optional] 
**PinnedPublicKeys** | Pointer to **map[string]string** |  | [optional] 
**Protocol** | Pointer to **string** |  | [optional] 
**Proxy** | Pointer to [**ProxyConfig**](ProxyConfig.md) |  | [optional] 
**RequestSigning** | Pointer to [**RequestSigningMeta**](RequestSigningMeta.md) |  | [optional] 
**ResponseProcessors** | Pointer to [**[]ResponseProcessor**](ResponseProcessor.md) |  | [optional] 
**Scopes** | Pointer to [**ScopesType2**](ScopesType2.md) |  | [optional] 
**SessionLifetime** | Pointer to **int64** |  | [optional] 
**SessionLifetimeRespectsKeyExpiration** | Pointer to **bool** |  | [optional] 
**SessionProvider** | Pointer to [**SessionProviderMeta**](SessionProviderMeta.md) |  | [optional] 
**Slug** | Pointer to **string** |  | [optional] 
**StripAuthData** | Pointer to **bool** |  | [optional] 
**TagHeaders** | Pointer to **[]string** |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**TagsDisabled** | Pointer to **bool** |  | [optional] 
**UpstreamCertificates** | Pointer to **map[string]string** |  | [optional] 
**UpstreamCertificatesDisabled** | Pointer to **bool** |  | [optional] 
**UptimeTests** | Pointer to [**UptimeTests**](UptimeTests.md) |  | [optional] 
**UseBasicAuth** | Pointer to **bool** |  | [optional] 
**UseGoPluginAuth** | Pointer to **bool** |  | [optional] 
**UseKeyless** | Pointer to **bool** |  | [optional] 
**UseMutualTlsAuth** | Pointer to **bool** |  | [optional] 
**UseOauth2** | Pointer to **bool** |  | [optional] 
**UseOpenid** | Pointer to **bool** |  | [optional] 
**UseStandardAuth** | Pointer to **bool** |  | [optional] 
**VersionData** | Pointer to [**VersionData**](VersionData.md) |  | [optional] 

## Methods

### NewAPIDefinition

`func NewAPIDefinition() *APIDefinition`

NewAPIDefinition instantiates a new APIDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAPIDefinitionWithDefaults

`func NewAPIDefinitionWithDefaults() *APIDefinition`

NewAPIDefinitionWithDefaults instantiates a new APIDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCORS

`func (o *APIDefinition) GetCORS() CORSConfig`

GetCORS returns the CORS field if non-nil, zero value otherwise.

### GetCORSOk

`func (o *APIDefinition) GetCORSOk() (*CORSConfig, bool)`

GetCORSOk returns a tuple with the CORS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCORS

`func (o *APIDefinition) SetCORS(v CORSConfig)`

SetCORS sets CORS field to given value.

### HasCORS

`func (o *APIDefinition) HasCORS() bool`

HasCORS returns a boolean if a field has been set.

### GetActive

`func (o *APIDefinition) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *APIDefinition) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *APIDefinition) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *APIDefinition) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetAllowedIps

`func (o *APIDefinition) GetAllowedIps() []string`

GetAllowedIps returns the AllowedIps field if non-nil, zero value otherwise.

### GetAllowedIpsOk

`func (o *APIDefinition) GetAllowedIpsOk() (*[]string, bool)`

GetAllowedIpsOk returns a tuple with the AllowedIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedIps

`func (o *APIDefinition) SetAllowedIps(v []string)`

SetAllowedIps sets AllowedIps field to given value.

### HasAllowedIps

`func (o *APIDefinition) HasAllowedIps() bool`

HasAllowedIps returns a boolean if a field has been set.

### SetAllowedIpsNil

`func (o *APIDefinition) SetAllowedIpsNil(b bool)`

 SetAllowedIpsNil sets the value for AllowedIps to be an explicit nil

### UnsetAllowedIps
`func (o *APIDefinition) UnsetAllowedIps()`

UnsetAllowedIps ensures that no value is present for AllowedIps, not even an explicit nil
### GetAnalyticsPlugin

`func (o *APIDefinition) GetAnalyticsPlugin() AnalyticsPluginConfig`

GetAnalyticsPlugin returns the AnalyticsPlugin field if non-nil, zero value otherwise.

### GetAnalyticsPluginOk

`func (o *APIDefinition) GetAnalyticsPluginOk() (*AnalyticsPluginConfig, bool)`

GetAnalyticsPluginOk returns a tuple with the AnalyticsPlugin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalyticsPlugin

`func (o *APIDefinition) SetAnalyticsPlugin(v AnalyticsPluginConfig)`

SetAnalyticsPlugin sets AnalyticsPlugin field to given value.

### HasAnalyticsPlugin

`func (o *APIDefinition) HasAnalyticsPlugin() bool`

HasAnalyticsPlugin returns a boolean if a field has been set.

### GetApiId

`func (o *APIDefinition) GetApiId() string`

GetApiId returns the ApiId field if non-nil, zero value otherwise.

### GetApiIdOk

`func (o *APIDefinition) GetApiIdOk() (*string, bool)`

GetApiIdOk returns a tuple with the ApiId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiId

`func (o *APIDefinition) SetApiId(v string)`

SetApiId sets ApiId field to given value.

### HasApiId

`func (o *APIDefinition) HasApiId() bool`

HasApiId returns a boolean if a field has been set.

### GetAuth

`func (o *APIDefinition) GetAuth() AuthConfig`

GetAuth returns the Auth field if non-nil, zero value otherwise.

### GetAuthOk

`func (o *APIDefinition) GetAuthOk() (*AuthConfig, bool)`

GetAuthOk returns a tuple with the Auth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuth

`func (o *APIDefinition) SetAuth(v AuthConfig)`

SetAuth sets Auth field to given value.

### HasAuth

`func (o *APIDefinition) HasAuth() bool`

HasAuth returns a boolean if a field has been set.

### GetAuthConfigs

`func (o *APIDefinition) GetAuthConfigs() map[string]AuthConfig`

GetAuthConfigs returns the AuthConfigs field if non-nil, zero value otherwise.

### GetAuthConfigsOk

`func (o *APIDefinition) GetAuthConfigsOk() (*map[string]AuthConfig, bool)`

GetAuthConfigsOk returns a tuple with the AuthConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthConfigs

`func (o *APIDefinition) SetAuthConfigs(v map[string]AuthConfig)`

SetAuthConfigs sets AuthConfigs field to given value.

### HasAuthConfigs

`func (o *APIDefinition) HasAuthConfigs() bool`

HasAuthConfigs returns a boolean if a field has been set.

### SetAuthConfigsNil

`func (o *APIDefinition) SetAuthConfigsNil(b bool)`

 SetAuthConfigsNil sets the value for AuthConfigs to be an explicit nil

### UnsetAuthConfigs
`func (o *APIDefinition) UnsetAuthConfigs()`

UnsetAuthConfigs ensures that no value is present for AuthConfigs, not even an explicit nil
### GetAuthProvider

`func (o *APIDefinition) GetAuthProvider() AuthProviderMeta`

GetAuthProvider returns the AuthProvider field if non-nil, zero value otherwise.

### GetAuthProviderOk

`func (o *APIDefinition) GetAuthProviderOk() (*AuthProviderMeta, bool)`

GetAuthProviderOk returns a tuple with the AuthProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthProvider

`func (o *APIDefinition) SetAuthProvider(v AuthProviderMeta)`

SetAuthProvider sets AuthProvider field to given value.

### HasAuthProvider

`func (o *APIDefinition) HasAuthProvider() bool`

HasAuthProvider returns a boolean if a field has been set.

### GetBaseIdentityProvidedBy

`func (o *APIDefinition) GetBaseIdentityProvidedBy() string`

GetBaseIdentityProvidedBy returns the BaseIdentityProvidedBy field if non-nil, zero value otherwise.

### GetBaseIdentityProvidedByOk

`func (o *APIDefinition) GetBaseIdentityProvidedByOk() (*string, bool)`

GetBaseIdentityProvidedByOk returns a tuple with the BaseIdentityProvidedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseIdentityProvidedBy

`func (o *APIDefinition) SetBaseIdentityProvidedBy(v string)`

SetBaseIdentityProvidedBy sets BaseIdentityProvidedBy field to given value.

### HasBaseIdentityProvidedBy

`func (o *APIDefinition) HasBaseIdentityProvidedBy() bool`

HasBaseIdentityProvidedBy returns a boolean if a field has been set.

### GetBasicAuth

`func (o *APIDefinition) GetBasicAuth() APIDefinitionBasicAuth`

GetBasicAuth returns the BasicAuth field if non-nil, zero value otherwise.

### GetBasicAuthOk

`func (o *APIDefinition) GetBasicAuthOk() (*APIDefinitionBasicAuth, bool)`

GetBasicAuthOk returns a tuple with the BasicAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasicAuth

`func (o *APIDefinition) SetBasicAuth(v APIDefinitionBasicAuth)`

SetBasicAuth sets BasicAuth field to given value.

### HasBasicAuth

`func (o *APIDefinition) HasBasicAuth() bool`

HasBasicAuth returns a boolean if a field has been set.

### GetBlacklistedIps

`func (o *APIDefinition) GetBlacklistedIps() []string`

GetBlacklistedIps returns the BlacklistedIps field if non-nil, zero value otherwise.

### GetBlacklistedIpsOk

`func (o *APIDefinition) GetBlacklistedIpsOk() (*[]string, bool)`

GetBlacklistedIpsOk returns a tuple with the BlacklistedIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlacklistedIps

`func (o *APIDefinition) SetBlacklistedIps(v []string)`

SetBlacklistedIps sets BlacklistedIps field to given value.

### HasBlacklistedIps

`func (o *APIDefinition) HasBlacklistedIps() bool`

HasBlacklistedIps returns a boolean if a field has been set.

### SetBlacklistedIpsNil

`func (o *APIDefinition) SetBlacklistedIpsNil(b bool)`

 SetBlacklistedIpsNil sets the value for BlacklistedIps to be an explicit nil

### UnsetBlacklistedIps
`func (o *APIDefinition) UnsetBlacklistedIps()`

UnsetBlacklistedIps ensures that no value is present for BlacklistedIps, not even an explicit nil
### GetCacheOptions

`func (o *APIDefinition) GetCacheOptions() CacheOptions`

GetCacheOptions returns the CacheOptions field if non-nil, zero value otherwise.

### GetCacheOptionsOk

`func (o *APIDefinition) GetCacheOptionsOk() (*CacheOptions, bool)`

GetCacheOptionsOk returns a tuple with the CacheOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheOptions

`func (o *APIDefinition) SetCacheOptions(v CacheOptions)`

SetCacheOptions sets CacheOptions field to given value.

### HasCacheOptions

`func (o *APIDefinition) HasCacheOptions() bool`

HasCacheOptions returns a boolean if a field has been set.

### GetCertificatePinningDisabled

`func (o *APIDefinition) GetCertificatePinningDisabled() bool`

GetCertificatePinningDisabled returns the CertificatePinningDisabled field if non-nil, zero value otherwise.

### GetCertificatePinningDisabledOk

`func (o *APIDefinition) GetCertificatePinningDisabledOk() (*bool, bool)`

GetCertificatePinningDisabledOk returns a tuple with the CertificatePinningDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificatePinningDisabled

`func (o *APIDefinition) SetCertificatePinningDisabled(v bool)`

SetCertificatePinningDisabled sets CertificatePinningDisabled field to given value.

### HasCertificatePinningDisabled

`func (o *APIDefinition) HasCertificatePinningDisabled() bool`

HasCertificatePinningDisabled returns a boolean if a field has been set.

### GetCertificates

`func (o *APIDefinition) GetCertificates() []string`

GetCertificates returns the Certificates field if non-nil, zero value otherwise.

### GetCertificatesOk

`func (o *APIDefinition) GetCertificatesOk() (*[]string, bool)`

GetCertificatesOk returns a tuple with the Certificates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificates

`func (o *APIDefinition) SetCertificates(v []string)`

SetCertificates sets Certificates field to given value.

### HasCertificates

`func (o *APIDefinition) HasCertificates() bool`

HasCertificates returns a boolean if a field has been set.

### SetCertificatesNil

`func (o *APIDefinition) SetCertificatesNil(b bool)`

 SetCertificatesNil sets the value for Certificates to be an explicit nil

### UnsetCertificates
`func (o *APIDefinition) UnsetCertificates()`

UnsetCertificates ensures that no value is present for Certificates, not even an explicit nil
### GetClientCertificates

`func (o *APIDefinition) GetClientCertificates() []string`

GetClientCertificates returns the ClientCertificates field if non-nil, zero value otherwise.

### GetClientCertificatesOk

`func (o *APIDefinition) GetClientCertificatesOk() (*[]string, bool)`

GetClientCertificatesOk returns a tuple with the ClientCertificates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientCertificates

`func (o *APIDefinition) SetClientCertificates(v []string)`

SetClientCertificates sets ClientCertificates field to given value.

### HasClientCertificates

`func (o *APIDefinition) HasClientCertificates() bool`

HasClientCertificates returns a boolean if a field has been set.

### SetClientCertificatesNil

`func (o *APIDefinition) SetClientCertificatesNil(b bool)`

 SetClientCertificatesNil sets the value for ClientCertificates to be an explicit nil

### UnsetClientCertificates
`func (o *APIDefinition) UnsetClientCertificates()`

UnsetClientCertificates ensures that no value is present for ClientCertificates, not even an explicit nil
### GetConfigData

`func (o *APIDefinition) GetConfigData() map[string]interface{}`

GetConfigData returns the ConfigData field if non-nil, zero value otherwise.

### GetConfigDataOk

`func (o *APIDefinition) GetConfigDataOk() (*map[string]interface{}, bool)`

GetConfigDataOk returns a tuple with the ConfigData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigData

`func (o *APIDefinition) SetConfigData(v map[string]interface{})`

SetConfigData sets ConfigData field to given value.

### HasConfigData

`func (o *APIDefinition) HasConfigData() bool`

HasConfigData returns a boolean if a field has been set.

### SetConfigDataNil

`func (o *APIDefinition) SetConfigDataNil(b bool)`

 SetConfigDataNil sets the value for ConfigData to be an explicit nil

### UnsetConfigData
`func (o *APIDefinition) UnsetConfigData()`

UnsetConfigData ensures that no value is present for ConfigData, not even an explicit nil
### GetConfigDataDisabled

`func (o *APIDefinition) GetConfigDataDisabled() bool`

GetConfigDataDisabled returns the ConfigDataDisabled field if non-nil, zero value otherwise.

### GetConfigDataDisabledOk

`func (o *APIDefinition) GetConfigDataDisabledOk() (*bool, bool)`

GetConfigDataDisabledOk returns a tuple with the ConfigDataDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigDataDisabled

`func (o *APIDefinition) SetConfigDataDisabled(v bool)`

SetConfigDataDisabled sets ConfigDataDisabled field to given value.

### HasConfigDataDisabled

`func (o *APIDefinition) HasConfigDataDisabled() bool`

HasConfigDataDisabled returns a boolean if a field has been set.

### GetCustomMiddleware

`func (o *APIDefinition) GetCustomMiddleware() MiddlewareSection`

GetCustomMiddleware returns the CustomMiddleware field if non-nil, zero value otherwise.

### GetCustomMiddlewareOk

`func (o *APIDefinition) GetCustomMiddlewareOk() (*MiddlewareSection, bool)`

GetCustomMiddlewareOk returns a tuple with the CustomMiddleware field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomMiddleware

`func (o *APIDefinition) SetCustomMiddleware(v MiddlewareSection)`

SetCustomMiddleware sets CustomMiddleware field to given value.

### HasCustomMiddleware

`func (o *APIDefinition) HasCustomMiddleware() bool`

HasCustomMiddleware returns a boolean if a field has been set.

### GetCustomMiddlewareBundle

`func (o *APIDefinition) GetCustomMiddlewareBundle() string`

GetCustomMiddlewareBundle returns the CustomMiddlewareBundle field if non-nil, zero value otherwise.

### GetCustomMiddlewareBundleOk

`func (o *APIDefinition) GetCustomMiddlewareBundleOk() (*string, bool)`

GetCustomMiddlewareBundleOk returns a tuple with the CustomMiddlewareBundle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomMiddlewareBundle

`func (o *APIDefinition) SetCustomMiddlewareBundle(v string)`

SetCustomMiddlewareBundle sets CustomMiddlewareBundle field to given value.

### HasCustomMiddlewareBundle

`func (o *APIDefinition) HasCustomMiddlewareBundle() bool`

HasCustomMiddlewareBundle returns a boolean if a field has been set.

### GetCustomMiddlewareBundleDisabled

`func (o *APIDefinition) GetCustomMiddlewareBundleDisabled() bool`

GetCustomMiddlewareBundleDisabled returns the CustomMiddlewareBundleDisabled field if non-nil, zero value otherwise.

### GetCustomMiddlewareBundleDisabledOk

`func (o *APIDefinition) GetCustomMiddlewareBundleDisabledOk() (*bool, bool)`

GetCustomMiddlewareBundleDisabledOk returns a tuple with the CustomMiddlewareBundleDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomMiddlewareBundleDisabled

`func (o *APIDefinition) SetCustomMiddlewareBundleDisabled(v bool)`

SetCustomMiddlewareBundleDisabled sets CustomMiddlewareBundleDisabled field to given value.

### HasCustomMiddlewareBundleDisabled

`func (o *APIDefinition) HasCustomMiddlewareBundleDisabled() bool`

HasCustomMiddlewareBundleDisabled returns a boolean if a field has been set.

### GetCustomPluginAuthEnabled

`func (o *APIDefinition) GetCustomPluginAuthEnabled() bool`

GetCustomPluginAuthEnabled returns the CustomPluginAuthEnabled field if non-nil, zero value otherwise.

### GetCustomPluginAuthEnabledOk

`func (o *APIDefinition) GetCustomPluginAuthEnabledOk() (*bool, bool)`

GetCustomPluginAuthEnabledOk returns a tuple with the CustomPluginAuthEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomPluginAuthEnabled

`func (o *APIDefinition) SetCustomPluginAuthEnabled(v bool)`

SetCustomPluginAuthEnabled sets CustomPluginAuthEnabled field to given value.

### HasCustomPluginAuthEnabled

`func (o *APIDefinition) HasCustomPluginAuthEnabled() bool`

HasCustomPluginAuthEnabled returns a boolean if a field has been set.

### GetDefinition

`func (o *APIDefinition) GetDefinition() VersionDefinition`

GetDefinition returns the Definition field if non-nil, zero value otherwise.

### GetDefinitionOk

`func (o *APIDefinition) GetDefinitionOk() (*VersionDefinition, bool)`

GetDefinitionOk returns a tuple with the Definition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefinition

`func (o *APIDefinition) SetDefinition(v VersionDefinition)`

SetDefinition sets Definition field to given value.

### HasDefinition

`func (o *APIDefinition) HasDefinition() bool`

HasDefinition returns a boolean if a field has been set.

### GetDetailedTracing

`func (o *APIDefinition) GetDetailedTracing() bool`

GetDetailedTracing returns the DetailedTracing field if non-nil, zero value otherwise.

### GetDetailedTracingOk

`func (o *APIDefinition) GetDetailedTracingOk() (*bool, bool)`

GetDetailedTracingOk returns a tuple with the DetailedTracing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetailedTracing

`func (o *APIDefinition) SetDetailedTracing(v bool)`

SetDetailedTracing sets DetailedTracing field to given value.

### HasDetailedTracing

`func (o *APIDefinition) HasDetailedTracing() bool`

HasDetailedTracing returns a boolean if a field has been set.

### GetDisableQuota

`func (o *APIDefinition) GetDisableQuota() bool`

GetDisableQuota returns the DisableQuota field if non-nil, zero value otherwise.

### GetDisableQuotaOk

`func (o *APIDefinition) GetDisableQuotaOk() (*bool, bool)`

GetDisableQuotaOk returns a tuple with the DisableQuota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableQuota

`func (o *APIDefinition) SetDisableQuota(v bool)`

SetDisableQuota sets DisableQuota field to given value.

### HasDisableQuota

`func (o *APIDefinition) HasDisableQuota() bool`

HasDisableQuota returns a boolean if a field has been set.

### GetDisableRateLimit

`func (o *APIDefinition) GetDisableRateLimit() bool`

GetDisableRateLimit returns the DisableRateLimit field if non-nil, zero value otherwise.

### GetDisableRateLimitOk

`func (o *APIDefinition) GetDisableRateLimitOk() (*bool, bool)`

GetDisableRateLimitOk returns a tuple with the DisableRateLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableRateLimit

`func (o *APIDefinition) SetDisableRateLimit(v bool)`

SetDisableRateLimit sets DisableRateLimit field to given value.

### HasDisableRateLimit

`func (o *APIDefinition) HasDisableRateLimit() bool`

HasDisableRateLimit returns a boolean if a field has been set.

### GetDoNotTrack

`func (o *APIDefinition) GetDoNotTrack() bool`

GetDoNotTrack returns the DoNotTrack field if non-nil, zero value otherwise.

### GetDoNotTrackOk

`func (o *APIDefinition) GetDoNotTrackOk() (*bool, bool)`

GetDoNotTrackOk returns a tuple with the DoNotTrack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoNotTrack

`func (o *APIDefinition) SetDoNotTrack(v bool)`

SetDoNotTrack sets DoNotTrack field to given value.

### HasDoNotTrack

`func (o *APIDefinition) HasDoNotTrack() bool`

HasDoNotTrack returns a boolean if a field has been set.

### GetDomain

`func (o *APIDefinition) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *APIDefinition) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *APIDefinition) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *APIDefinition) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetDomainDisabled

`func (o *APIDefinition) GetDomainDisabled() bool`

GetDomainDisabled returns the DomainDisabled field if non-nil, zero value otherwise.

### GetDomainDisabledOk

`func (o *APIDefinition) GetDomainDisabledOk() (*bool, bool)`

GetDomainDisabledOk returns a tuple with the DomainDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainDisabled

`func (o *APIDefinition) SetDomainDisabled(v bool)`

SetDomainDisabled sets DomainDisabled field to given value.

### HasDomainDisabled

`func (o *APIDefinition) HasDomainDisabled() bool`

HasDomainDisabled returns a boolean if a field has been set.

### GetDontSetQuotaOnCreate

`func (o *APIDefinition) GetDontSetQuotaOnCreate() bool`

GetDontSetQuotaOnCreate returns the DontSetQuotaOnCreate field if non-nil, zero value otherwise.

### GetDontSetQuotaOnCreateOk

`func (o *APIDefinition) GetDontSetQuotaOnCreateOk() (*bool, bool)`

GetDontSetQuotaOnCreateOk returns a tuple with the DontSetQuotaOnCreate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDontSetQuotaOnCreate

`func (o *APIDefinition) SetDontSetQuotaOnCreate(v bool)`

SetDontSetQuotaOnCreate sets DontSetQuotaOnCreate field to given value.

### HasDontSetQuotaOnCreate

`func (o *APIDefinition) HasDontSetQuotaOnCreate() bool`

HasDontSetQuotaOnCreate returns a boolean if a field has been set.

### GetEnableBatchRequestSupport

`func (o *APIDefinition) GetEnableBatchRequestSupport() bool`

GetEnableBatchRequestSupport returns the EnableBatchRequestSupport field if non-nil, zero value otherwise.

### GetEnableBatchRequestSupportOk

`func (o *APIDefinition) GetEnableBatchRequestSupportOk() (*bool, bool)`

GetEnableBatchRequestSupportOk returns a tuple with the EnableBatchRequestSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableBatchRequestSupport

`func (o *APIDefinition) SetEnableBatchRequestSupport(v bool)`

SetEnableBatchRequestSupport sets EnableBatchRequestSupport field to given value.

### HasEnableBatchRequestSupport

`func (o *APIDefinition) HasEnableBatchRequestSupport() bool`

HasEnableBatchRequestSupport returns a boolean if a field has been set.

### GetEnableContextVars

`func (o *APIDefinition) GetEnableContextVars() bool`

GetEnableContextVars returns the EnableContextVars field if non-nil, zero value otherwise.

### GetEnableContextVarsOk

`func (o *APIDefinition) GetEnableContextVarsOk() (*bool, bool)`

GetEnableContextVarsOk returns a tuple with the EnableContextVars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableContextVars

`func (o *APIDefinition) SetEnableContextVars(v bool)`

SetEnableContextVars sets EnableContextVars field to given value.

### HasEnableContextVars

`func (o *APIDefinition) HasEnableContextVars() bool`

HasEnableContextVars returns a boolean if a field has been set.

### GetEnableCoprocessAuth

`func (o *APIDefinition) GetEnableCoprocessAuth() bool`

GetEnableCoprocessAuth returns the EnableCoprocessAuth field if non-nil, zero value otherwise.

### GetEnableCoprocessAuthOk

`func (o *APIDefinition) GetEnableCoprocessAuthOk() (*bool, bool)`

GetEnableCoprocessAuthOk returns a tuple with the EnableCoprocessAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableCoprocessAuth

`func (o *APIDefinition) SetEnableCoprocessAuth(v bool)`

SetEnableCoprocessAuth sets EnableCoprocessAuth field to given value.

### HasEnableCoprocessAuth

`func (o *APIDefinition) HasEnableCoprocessAuth() bool`

HasEnableCoprocessAuth returns a boolean if a field has been set.

### GetEnableDetailedRecording

`func (o *APIDefinition) GetEnableDetailedRecording() bool`

GetEnableDetailedRecording returns the EnableDetailedRecording field if non-nil, zero value otherwise.

### GetEnableDetailedRecordingOk

`func (o *APIDefinition) GetEnableDetailedRecordingOk() (*bool, bool)`

GetEnableDetailedRecordingOk returns a tuple with the EnableDetailedRecording field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableDetailedRecording

`func (o *APIDefinition) SetEnableDetailedRecording(v bool)`

SetEnableDetailedRecording sets EnableDetailedRecording field to given value.

### HasEnableDetailedRecording

`func (o *APIDefinition) HasEnableDetailedRecording() bool`

HasEnableDetailedRecording returns a boolean if a field has been set.

### GetEnableIpBlacklisting

`func (o *APIDefinition) GetEnableIpBlacklisting() bool`

GetEnableIpBlacklisting returns the EnableIpBlacklisting field if non-nil, zero value otherwise.

### GetEnableIpBlacklistingOk

`func (o *APIDefinition) GetEnableIpBlacklistingOk() (*bool, bool)`

GetEnableIpBlacklistingOk returns a tuple with the EnableIpBlacklisting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableIpBlacklisting

`func (o *APIDefinition) SetEnableIpBlacklisting(v bool)`

SetEnableIpBlacklisting sets EnableIpBlacklisting field to given value.

### HasEnableIpBlacklisting

`func (o *APIDefinition) HasEnableIpBlacklisting() bool`

HasEnableIpBlacklisting returns a boolean if a field has been set.

### GetEnableIpWhitelisting

`func (o *APIDefinition) GetEnableIpWhitelisting() bool`

GetEnableIpWhitelisting returns the EnableIpWhitelisting field if non-nil, zero value otherwise.

### GetEnableIpWhitelistingOk

`func (o *APIDefinition) GetEnableIpWhitelistingOk() (*bool, bool)`

GetEnableIpWhitelistingOk returns a tuple with the EnableIpWhitelisting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableIpWhitelisting

`func (o *APIDefinition) SetEnableIpWhitelisting(v bool)`

SetEnableIpWhitelisting sets EnableIpWhitelisting field to given value.

### HasEnableIpWhitelisting

`func (o *APIDefinition) HasEnableIpWhitelisting() bool`

HasEnableIpWhitelisting returns a boolean if a field has been set.

### GetEnableJwt

`func (o *APIDefinition) GetEnableJwt() bool`

GetEnableJwt returns the EnableJwt field if non-nil, zero value otherwise.

### GetEnableJwtOk

`func (o *APIDefinition) GetEnableJwtOk() (*bool, bool)`

GetEnableJwtOk returns a tuple with the EnableJwt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableJwt

`func (o *APIDefinition) SetEnableJwt(v bool)`

SetEnableJwt sets EnableJwt field to given value.

### HasEnableJwt

`func (o *APIDefinition) HasEnableJwt() bool`

HasEnableJwt returns a boolean if a field has been set.

### GetEnableProxyProtocol

`func (o *APIDefinition) GetEnableProxyProtocol() bool`

GetEnableProxyProtocol returns the EnableProxyProtocol field if non-nil, zero value otherwise.

### GetEnableProxyProtocolOk

`func (o *APIDefinition) GetEnableProxyProtocolOk() (*bool, bool)`

GetEnableProxyProtocolOk returns a tuple with the EnableProxyProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableProxyProtocol

`func (o *APIDefinition) SetEnableProxyProtocol(v bool)`

SetEnableProxyProtocol sets EnableProxyProtocol field to given value.

### HasEnableProxyProtocol

`func (o *APIDefinition) HasEnableProxyProtocol() bool`

HasEnableProxyProtocol returns a boolean if a field has been set.

### GetEnableSignatureChecking

`func (o *APIDefinition) GetEnableSignatureChecking() bool`

GetEnableSignatureChecking returns the EnableSignatureChecking field if non-nil, zero value otherwise.

### GetEnableSignatureCheckingOk

`func (o *APIDefinition) GetEnableSignatureCheckingOk() (*bool, bool)`

GetEnableSignatureCheckingOk returns a tuple with the EnableSignatureChecking field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableSignatureChecking

`func (o *APIDefinition) SetEnableSignatureChecking(v bool)`

SetEnableSignatureChecking sets EnableSignatureChecking field to given value.

### HasEnableSignatureChecking

`func (o *APIDefinition) HasEnableSignatureChecking() bool`

HasEnableSignatureChecking returns a boolean if a field has been set.

### GetEventHandlers

`func (o *APIDefinition) GetEventHandlers() EventHandlerMetaConfig`

GetEventHandlers returns the EventHandlers field if non-nil, zero value otherwise.

### GetEventHandlersOk

`func (o *APIDefinition) GetEventHandlersOk() (*EventHandlerMetaConfig, bool)`

GetEventHandlersOk returns a tuple with the EventHandlers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventHandlers

`func (o *APIDefinition) SetEventHandlers(v EventHandlerMetaConfig)`

SetEventHandlers sets EventHandlers field to given value.

### HasEventHandlers

`func (o *APIDefinition) HasEventHandlers() bool`

HasEventHandlers returns a boolean if a field has been set.

### GetExpiration

`func (o *APIDefinition) GetExpiration() string`

GetExpiration returns the Expiration field if non-nil, zero value otherwise.

### GetExpirationOk

`func (o *APIDefinition) GetExpirationOk() (*string, bool)`

GetExpirationOk returns a tuple with the Expiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiration

`func (o *APIDefinition) SetExpiration(v string)`

SetExpiration sets Expiration field to given value.

### HasExpiration

`func (o *APIDefinition) HasExpiration() bool`

HasExpiration returns a boolean if a field has been set.

### GetExpireAnalyticsAfter

`func (o *APIDefinition) GetExpireAnalyticsAfter() int64`

GetExpireAnalyticsAfter returns the ExpireAnalyticsAfter field if non-nil, zero value otherwise.

### GetExpireAnalyticsAfterOk

`func (o *APIDefinition) GetExpireAnalyticsAfterOk() (*int64, bool)`

GetExpireAnalyticsAfterOk returns a tuple with the ExpireAnalyticsAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpireAnalyticsAfter

`func (o *APIDefinition) SetExpireAnalyticsAfter(v int64)`

SetExpireAnalyticsAfter sets ExpireAnalyticsAfter field to given value.

### HasExpireAnalyticsAfter

`func (o *APIDefinition) HasExpireAnalyticsAfter() bool`

HasExpireAnalyticsAfter returns a boolean if a field has been set.

### GetExternalOauth

`func (o *APIDefinition) GetExternalOauth() ExternalOAuth`

GetExternalOauth returns the ExternalOauth field if non-nil, zero value otherwise.

### GetExternalOauthOk

`func (o *APIDefinition) GetExternalOauthOk() (*ExternalOAuth, bool)`

GetExternalOauthOk returns a tuple with the ExternalOauth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalOauth

`func (o *APIDefinition) SetExternalOauth(v ExternalOAuth)`

SetExternalOauth sets ExternalOauth field to given value.

### HasExternalOauth

`func (o *APIDefinition) HasExternalOauth() bool`

HasExternalOauth returns a boolean if a field has been set.

### GetGlobalRateLimit

`func (o *APIDefinition) GetGlobalRateLimit() GlobalRateLimit`

GetGlobalRateLimit returns the GlobalRateLimit field if non-nil, zero value otherwise.

### GetGlobalRateLimitOk

`func (o *APIDefinition) GetGlobalRateLimitOk() (*GlobalRateLimit, bool)`

GetGlobalRateLimitOk returns a tuple with the GlobalRateLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalRateLimit

`func (o *APIDefinition) SetGlobalRateLimit(v GlobalRateLimit)`

SetGlobalRateLimit sets GlobalRateLimit field to given value.

### HasGlobalRateLimit

`func (o *APIDefinition) HasGlobalRateLimit() bool`

HasGlobalRateLimit returns a boolean if a field has been set.

### GetGraphql

`func (o *APIDefinition) GetGraphql() GraphQLConfig`

GetGraphql returns the Graphql field if non-nil, zero value otherwise.

### GetGraphqlOk

`func (o *APIDefinition) GetGraphqlOk() (*GraphQLConfig, bool)`

GetGraphqlOk returns a tuple with the Graphql field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGraphql

`func (o *APIDefinition) SetGraphql(v GraphQLConfig)`

SetGraphql sets Graphql field to given value.

### HasGraphql

`func (o *APIDefinition) HasGraphql() bool`

HasGraphql returns a boolean if a field has been set.

### GetHmacAllowedAlgorithms

`func (o *APIDefinition) GetHmacAllowedAlgorithms() []string`

GetHmacAllowedAlgorithms returns the HmacAllowedAlgorithms field if non-nil, zero value otherwise.

### GetHmacAllowedAlgorithmsOk

`func (o *APIDefinition) GetHmacAllowedAlgorithmsOk() (*[]string, bool)`

GetHmacAllowedAlgorithmsOk returns a tuple with the HmacAllowedAlgorithms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHmacAllowedAlgorithms

`func (o *APIDefinition) SetHmacAllowedAlgorithms(v []string)`

SetHmacAllowedAlgorithms sets HmacAllowedAlgorithms field to given value.

### HasHmacAllowedAlgorithms

`func (o *APIDefinition) HasHmacAllowedAlgorithms() bool`

HasHmacAllowedAlgorithms returns a boolean if a field has been set.

### SetHmacAllowedAlgorithmsNil

`func (o *APIDefinition) SetHmacAllowedAlgorithmsNil(b bool)`

 SetHmacAllowedAlgorithmsNil sets the value for HmacAllowedAlgorithms to be an explicit nil

### UnsetHmacAllowedAlgorithms
`func (o *APIDefinition) UnsetHmacAllowedAlgorithms()`

UnsetHmacAllowedAlgorithms ensures that no value is present for HmacAllowedAlgorithms, not even an explicit nil
### GetHmacAllowedClockSkew

`func (o *APIDefinition) GetHmacAllowedClockSkew() float64`

GetHmacAllowedClockSkew returns the HmacAllowedClockSkew field if non-nil, zero value otherwise.

### GetHmacAllowedClockSkewOk

`func (o *APIDefinition) GetHmacAllowedClockSkewOk() (*float64, bool)`

GetHmacAllowedClockSkewOk returns a tuple with the HmacAllowedClockSkew field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHmacAllowedClockSkew

`func (o *APIDefinition) SetHmacAllowedClockSkew(v float64)`

SetHmacAllowedClockSkew sets HmacAllowedClockSkew field to given value.

### HasHmacAllowedClockSkew

`func (o *APIDefinition) HasHmacAllowedClockSkew() bool`

HasHmacAllowedClockSkew returns a boolean if a field has been set.

### GetId

`func (o *APIDefinition) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *APIDefinition) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *APIDefinition) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *APIDefinition) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIdpClientIdMappingDisabled

`func (o *APIDefinition) GetIdpClientIdMappingDisabled() bool`

GetIdpClientIdMappingDisabled returns the IdpClientIdMappingDisabled field if non-nil, zero value otherwise.

### GetIdpClientIdMappingDisabledOk

`func (o *APIDefinition) GetIdpClientIdMappingDisabledOk() (*bool, bool)`

GetIdpClientIdMappingDisabledOk returns a tuple with the IdpClientIdMappingDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpClientIdMappingDisabled

`func (o *APIDefinition) SetIdpClientIdMappingDisabled(v bool)`

SetIdpClientIdMappingDisabled sets IdpClientIdMappingDisabled field to given value.

### HasIdpClientIdMappingDisabled

`func (o *APIDefinition) HasIdpClientIdMappingDisabled() bool`

HasIdpClientIdMappingDisabled returns a boolean if a field has been set.

### GetInternal

`func (o *APIDefinition) GetInternal() bool`

GetInternal returns the Internal field if non-nil, zero value otherwise.

### GetInternalOk

`func (o *APIDefinition) GetInternalOk() (*bool, bool)`

GetInternalOk returns a tuple with the Internal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternal

`func (o *APIDefinition) SetInternal(v bool)`

SetInternal sets Internal field to given value.

### HasInternal

`func (o *APIDefinition) HasInternal() bool`

HasInternal returns a boolean if a field has been set.

### GetIsOas

`func (o *APIDefinition) GetIsOas() bool`

GetIsOas returns the IsOas field if non-nil, zero value otherwise.

### GetIsOasOk

`func (o *APIDefinition) GetIsOasOk() (*bool, bool)`

GetIsOasOk returns a tuple with the IsOas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOas

`func (o *APIDefinition) SetIsOas(v bool)`

SetIsOas sets IsOas field to given value.

### HasIsOas

`func (o *APIDefinition) HasIsOas() bool`

HasIsOas returns a boolean if a field has been set.

### GetJwtClientBaseField

`func (o *APIDefinition) GetJwtClientBaseField() string`

GetJwtClientBaseField returns the JwtClientBaseField field if non-nil, zero value otherwise.

### GetJwtClientBaseFieldOk

`func (o *APIDefinition) GetJwtClientBaseFieldOk() (*string, bool)`

GetJwtClientBaseFieldOk returns a tuple with the JwtClientBaseField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwtClientBaseField

`func (o *APIDefinition) SetJwtClientBaseField(v string)`

SetJwtClientBaseField sets JwtClientBaseField field to given value.

### HasJwtClientBaseField

`func (o *APIDefinition) HasJwtClientBaseField() bool`

HasJwtClientBaseField returns a boolean if a field has been set.

### GetJwtDefaultPolicies

`func (o *APIDefinition) GetJwtDefaultPolicies() []string`

GetJwtDefaultPolicies returns the JwtDefaultPolicies field if non-nil, zero value otherwise.

### GetJwtDefaultPoliciesOk

`func (o *APIDefinition) GetJwtDefaultPoliciesOk() (*[]string, bool)`

GetJwtDefaultPoliciesOk returns a tuple with the JwtDefaultPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwtDefaultPolicies

`func (o *APIDefinition) SetJwtDefaultPolicies(v []string)`

SetJwtDefaultPolicies sets JwtDefaultPolicies field to given value.

### HasJwtDefaultPolicies

`func (o *APIDefinition) HasJwtDefaultPolicies() bool`

HasJwtDefaultPolicies returns a boolean if a field has been set.

### SetJwtDefaultPoliciesNil

`func (o *APIDefinition) SetJwtDefaultPoliciesNil(b bool)`

 SetJwtDefaultPoliciesNil sets the value for JwtDefaultPolicies to be an explicit nil

### UnsetJwtDefaultPolicies
`func (o *APIDefinition) UnsetJwtDefaultPolicies()`

UnsetJwtDefaultPolicies ensures that no value is present for JwtDefaultPolicies, not even an explicit nil
### GetJwtExpiresAtValidationSkew

`func (o *APIDefinition) GetJwtExpiresAtValidationSkew() int32`

GetJwtExpiresAtValidationSkew returns the JwtExpiresAtValidationSkew field if non-nil, zero value otherwise.

### GetJwtExpiresAtValidationSkewOk

`func (o *APIDefinition) GetJwtExpiresAtValidationSkewOk() (*int32, bool)`

GetJwtExpiresAtValidationSkewOk returns a tuple with the JwtExpiresAtValidationSkew field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwtExpiresAtValidationSkew

`func (o *APIDefinition) SetJwtExpiresAtValidationSkew(v int32)`

SetJwtExpiresAtValidationSkew sets JwtExpiresAtValidationSkew field to given value.

### HasJwtExpiresAtValidationSkew

`func (o *APIDefinition) HasJwtExpiresAtValidationSkew() bool`

HasJwtExpiresAtValidationSkew returns a boolean if a field has been set.

### GetJwtIdentityBaseField

`func (o *APIDefinition) GetJwtIdentityBaseField() string`

GetJwtIdentityBaseField returns the JwtIdentityBaseField field if non-nil, zero value otherwise.

### GetJwtIdentityBaseFieldOk

`func (o *APIDefinition) GetJwtIdentityBaseFieldOk() (*string, bool)`

GetJwtIdentityBaseFieldOk returns a tuple with the JwtIdentityBaseField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwtIdentityBaseField

`func (o *APIDefinition) SetJwtIdentityBaseField(v string)`

SetJwtIdentityBaseField sets JwtIdentityBaseField field to given value.

### HasJwtIdentityBaseField

`func (o *APIDefinition) HasJwtIdentityBaseField() bool`

HasJwtIdentityBaseField returns a boolean if a field has been set.

### GetJwtIssuedAtValidationSkew

`func (o *APIDefinition) GetJwtIssuedAtValidationSkew() int32`

GetJwtIssuedAtValidationSkew returns the JwtIssuedAtValidationSkew field if non-nil, zero value otherwise.

### GetJwtIssuedAtValidationSkewOk

`func (o *APIDefinition) GetJwtIssuedAtValidationSkewOk() (*int32, bool)`

GetJwtIssuedAtValidationSkewOk returns a tuple with the JwtIssuedAtValidationSkew field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwtIssuedAtValidationSkew

`func (o *APIDefinition) SetJwtIssuedAtValidationSkew(v int32)`

SetJwtIssuedAtValidationSkew sets JwtIssuedAtValidationSkew field to given value.

### HasJwtIssuedAtValidationSkew

`func (o *APIDefinition) HasJwtIssuedAtValidationSkew() bool`

HasJwtIssuedAtValidationSkew returns a boolean if a field has been set.

### GetJwtNotBeforeValidationSkew

`func (o *APIDefinition) GetJwtNotBeforeValidationSkew() int32`

GetJwtNotBeforeValidationSkew returns the JwtNotBeforeValidationSkew field if non-nil, zero value otherwise.

### GetJwtNotBeforeValidationSkewOk

`func (o *APIDefinition) GetJwtNotBeforeValidationSkewOk() (*int32, bool)`

GetJwtNotBeforeValidationSkewOk returns a tuple with the JwtNotBeforeValidationSkew field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwtNotBeforeValidationSkew

`func (o *APIDefinition) SetJwtNotBeforeValidationSkew(v int32)`

SetJwtNotBeforeValidationSkew sets JwtNotBeforeValidationSkew field to given value.

### HasJwtNotBeforeValidationSkew

`func (o *APIDefinition) HasJwtNotBeforeValidationSkew() bool`

HasJwtNotBeforeValidationSkew returns a boolean if a field has been set.

### GetJwtPolicyFieldName

`func (o *APIDefinition) GetJwtPolicyFieldName() string`

GetJwtPolicyFieldName returns the JwtPolicyFieldName field if non-nil, zero value otherwise.

### GetJwtPolicyFieldNameOk

`func (o *APIDefinition) GetJwtPolicyFieldNameOk() (*string, bool)`

GetJwtPolicyFieldNameOk returns a tuple with the JwtPolicyFieldName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwtPolicyFieldName

`func (o *APIDefinition) SetJwtPolicyFieldName(v string)`

SetJwtPolicyFieldName sets JwtPolicyFieldName field to given value.

### HasJwtPolicyFieldName

`func (o *APIDefinition) HasJwtPolicyFieldName() bool`

HasJwtPolicyFieldName returns a boolean if a field has been set.

### GetJwtScopeClaimName

`func (o *APIDefinition) GetJwtScopeClaimName() string`

GetJwtScopeClaimName returns the JwtScopeClaimName field if non-nil, zero value otherwise.

### GetJwtScopeClaimNameOk

`func (o *APIDefinition) GetJwtScopeClaimNameOk() (*string, bool)`

GetJwtScopeClaimNameOk returns a tuple with the JwtScopeClaimName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwtScopeClaimName

`func (o *APIDefinition) SetJwtScopeClaimName(v string)`

SetJwtScopeClaimName sets JwtScopeClaimName field to given value.

### HasJwtScopeClaimName

`func (o *APIDefinition) HasJwtScopeClaimName() bool`

HasJwtScopeClaimName returns a boolean if a field has been set.

### GetJwtScopeToPolicyMapping

`func (o *APIDefinition) GetJwtScopeToPolicyMapping() map[string]string`

GetJwtScopeToPolicyMapping returns the JwtScopeToPolicyMapping field if non-nil, zero value otherwise.

### GetJwtScopeToPolicyMappingOk

`func (o *APIDefinition) GetJwtScopeToPolicyMappingOk() (*map[string]string, bool)`

GetJwtScopeToPolicyMappingOk returns a tuple with the JwtScopeToPolicyMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwtScopeToPolicyMapping

`func (o *APIDefinition) SetJwtScopeToPolicyMapping(v map[string]string)`

SetJwtScopeToPolicyMapping sets JwtScopeToPolicyMapping field to given value.

### HasJwtScopeToPolicyMapping

`func (o *APIDefinition) HasJwtScopeToPolicyMapping() bool`

HasJwtScopeToPolicyMapping returns a boolean if a field has been set.

### SetJwtScopeToPolicyMappingNil

`func (o *APIDefinition) SetJwtScopeToPolicyMappingNil(b bool)`

 SetJwtScopeToPolicyMappingNil sets the value for JwtScopeToPolicyMapping to be an explicit nil

### UnsetJwtScopeToPolicyMapping
`func (o *APIDefinition) UnsetJwtScopeToPolicyMapping()`

UnsetJwtScopeToPolicyMapping ensures that no value is present for JwtScopeToPolicyMapping, not even an explicit nil
### GetJwtSigningMethod

`func (o *APIDefinition) GetJwtSigningMethod() string`

GetJwtSigningMethod returns the JwtSigningMethod field if non-nil, zero value otherwise.

### GetJwtSigningMethodOk

`func (o *APIDefinition) GetJwtSigningMethodOk() (*string, bool)`

GetJwtSigningMethodOk returns a tuple with the JwtSigningMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwtSigningMethod

`func (o *APIDefinition) SetJwtSigningMethod(v string)`

SetJwtSigningMethod sets JwtSigningMethod field to given value.

### HasJwtSigningMethod

`func (o *APIDefinition) HasJwtSigningMethod() bool`

HasJwtSigningMethod returns a boolean if a field has been set.

### GetJwtSkipKid

`func (o *APIDefinition) GetJwtSkipKid() bool`

GetJwtSkipKid returns the JwtSkipKid field if non-nil, zero value otherwise.

### GetJwtSkipKidOk

`func (o *APIDefinition) GetJwtSkipKidOk() (*bool, bool)`

GetJwtSkipKidOk returns a tuple with the JwtSkipKid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwtSkipKid

`func (o *APIDefinition) SetJwtSkipKid(v bool)`

SetJwtSkipKid sets JwtSkipKid field to given value.

### HasJwtSkipKid

`func (o *APIDefinition) HasJwtSkipKid() bool`

HasJwtSkipKid returns a boolean if a field has been set.

### GetJwtSource

`func (o *APIDefinition) GetJwtSource() string`

GetJwtSource returns the JwtSource field if non-nil, zero value otherwise.

### GetJwtSourceOk

`func (o *APIDefinition) GetJwtSourceOk() (*string, bool)`

GetJwtSourceOk returns a tuple with the JwtSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwtSource

`func (o *APIDefinition) SetJwtSource(v string)`

SetJwtSource sets JwtSource field to given value.

### HasJwtSource

`func (o *APIDefinition) HasJwtSource() bool`

HasJwtSource returns a boolean if a field has been set.

### GetListenPort

`func (o *APIDefinition) GetListenPort() int32`

GetListenPort returns the ListenPort field if non-nil, zero value otherwise.

### GetListenPortOk

`func (o *APIDefinition) GetListenPortOk() (*int32, bool)`

GetListenPortOk returns a tuple with the ListenPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListenPort

`func (o *APIDefinition) SetListenPort(v int32)`

SetListenPort sets ListenPort field to given value.

### HasListenPort

`func (o *APIDefinition) HasListenPort() bool`

HasListenPort returns a boolean if a field has been set.

### GetName

`func (o *APIDefinition) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *APIDefinition) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *APIDefinition) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *APIDefinition) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNotifications

`func (o *APIDefinition) GetNotifications() NotificationsManager`

GetNotifications returns the Notifications field if non-nil, zero value otherwise.

### GetNotificationsOk

`func (o *APIDefinition) GetNotificationsOk() (*NotificationsManager, bool)`

GetNotificationsOk returns a tuple with the Notifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifications

`func (o *APIDefinition) SetNotifications(v NotificationsManager)`

SetNotifications sets Notifications field to given value.

### HasNotifications

`func (o *APIDefinition) HasNotifications() bool`

HasNotifications returns a boolean if a field has been set.

### GetOauthMeta

`func (o *APIDefinition) GetOauthMeta() APIDefinitionOauthMeta`

GetOauthMeta returns the OauthMeta field if non-nil, zero value otherwise.

### GetOauthMetaOk

`func (o *APIDefinition) GetOauthMetaOk() (*APIDefinitionOauthMeta, bool)`

GetOauthMetaOk returns a tuple with the OauthMeta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauthMeta

`func (o *APIDefinition) SetOauthMeta(v APIDefinitionOauthMeta)`

SetOauthMeta sets OauthMeta field to given value.

### HasOauthMeta

`func (o *APIDefinition) HasOauthMeta() bool`

HasOauthMeta returns a boolean if a field has been set.

### GetOpenidOptions

`func (o *APIDefinition) GetOpenidOptions() OpenIDOptions`

GetOpenidOptions returns the OpenidOptions field if non-nil, zero value otherwise.

### GetOpenidOptionsOk

`func (o *APIDefinition) GetOpenidOptionsOk() (*OpenIDOptions, bool)`

GetOpenidOptionsOk returns a tuple with the OpenidOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenidOptions

`func (o *APIDefinition) SetOpenidOptions(v OpenIDOptions)`

SetOpenidOptions sets OpenidOptions field to given value.

### HasOpenidOptions

`func (o *APIDefinition) HasOpenidOptions() bool`

HasOpenidOptions returns a boolean if a field has been set.

### GetOrgId

`func (o *APIDefinition) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *APIDefinition) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *APIDefinition) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *APIDefinition) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetPinnedPublicKeys

`func (o *APIDefinition) GetPinnedPublicKeys() map[string]string`

GetPinnedPublicKeys returns the PinnedPublicKeys field if non-nil, zero value otherwise.

### GetPinnedPublicKeysOk

`func (o *APIDefinition) GetPinnedPublicKeysOk() (*map[string]string, bool)`

GetPinnedPublicKeysOk returns a tuple with the PinnedPublicKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPinnedPublicKeys

`func (o *APIDefinition) SetPinnedPublicKeys(v map[string]string)`

SetPinnedPublicKeys sets PinnedPublicKeys field to given value.

### HasPinnedPublicKeys

`func (o *APIDefinition) HasPinnedPublicKeys() bool`

HasPinnedPublicKeys returns a boolean if a field has been set.

### SetPinnedPublicKeysNil

`func (o *APIDefinition) SetPinnedPublicKeysNil(b bool)`

 SetPinnedPublicKeysNil sets the value for PinnedPublicKeys to be an explicit nil

### UnsetPinnedPublicKeys
`func (o *APIDefinition) UnsetPinnedPublicKeys()`

UnsetPinnedPublicKeys ensures that no value is present for PinnedPublicKeys, not even an explicit nil
### GetProtocol

`func (o *APIDefinition) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *APIDefinition) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *APIDefinition) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *APIDefinition) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetProxy

`func (o *APIDefinition) GetProxy() ProxyConfig`

GetProxy returns the Proxy field if non-nil, zero value otherwise.

### GetProxyOk

`func (o *APIDefinition) GetProxyOk() (*ProxyConfig, bool)`

GetProxyOk returns a tuple with the Proxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxy

`func (o *APIDefinition) SetProxy(v ProxyConfig)`

SetProxy sets Proxy field to given value.

### HasProxy

`func (o *APIDefinition) HasProxy() bool`

HasProxy returns a boolean if a field has been set.

### GetRequestSigning

`func (o *APIDefinition) GetRequestSigning() RequestSigningMeta`

GetRequestSigning returns the RequestSigning field if non-nil, zero value otherwise.

### GetRequestSigningOk

`func (o *APIDefinition) GetRequestSigningOk() (*RequestSigningMeta, bool)`

GetRequestSigningOk returns a tuple with the RequestSigning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestSigning

`func (o *APIDefinition) SetRequestSigning(v RequestSigningMeta)`

SetRequestSigning sets RequestSigning field to given value.

### HasRequestSigning

`func (o *APIDefinition) HasRequestSigning() bool`

HasRequestSigning returns a boolean if a field has been set.

### GetResponseProcessors

`func (o *APIDefinition) GetResponseProcessors() []ResponseProcessor`

GetResponseProcessors returns the ResponseProcessors field if non-nil, zero value otherwise.

### GetResponseProcessorsOk

`func (o *APIDefinition) GetResponseProcessorsOk() (*[]ResponseProcessor, bool)`

GetResponseProcessorsOk returns a tuple with the ResponseProcessors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseProcessors

`func (o *APIDefinition) SetResponseProcessors(v []ResponseProcessor)`

SetResponseProcessors sets ResponseProcessors field to given value.

### HasResponseProcessors

`func (o *APIDefinition) HasResponseProcessors() bool`

HasResponseProcessors returns a boolean if a field has been set.

### SetResponseProcessorsNil

`func (o *APIDefinition) SetResponseProcessorsNil(b bool)`

 SetResponseProcessorsNil sets the value for ResponseProcessors to be an explicit nil

### UnsetResponseProcessors
`func (o *APIDefinition) UnsetResponseProcessors()`

UnsetResponseProcessors ensures that no value is present for ResponseProcessors, not even an explicit nil
### GetScopes

`func (o *APIDefinition) GetScopes() ScopesType2`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *APIDefinition) GetScopesOk() (*ScopesType2, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *APIDefinition) SetScopes(v ScopesType2)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *APIDefinition) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### GetSessionLifetime

`func (o *APIDefinition) GetSessionLifetime() int64`

GetSessionLifetime returns the SessionLifetime field if non-nil, zero value otherwise.

### GetSessionLifetimeOk

`func (o *APIDefinition) GetSessionLifetimeOk() (*int64, bool)`

GetSessionLifetimeOk returns a tuple with the SessionLifetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionLifetime

`func (o *APIDefinition) SetSessionLifetime(v int64)`

SetSessionLifetime sets SessionLifetime field to given value.

### HasSessionLifetime

`func (o *APIDefinition) HasSessionLifetime() bool`

HasSessionLifetime returns a boolean if a field has been set.

### GetSessionLifetimeRespectsKeyExpiration

`func (o *APIDefinition) GetSessionLifetimeRespectsKeyExpiration() bool`

GetSessionLifetimeRespectsKeyExpiration returns the SessionLifetimeRespectsKeyExpiration field if non-nil, zero value otherwise.

### GetSessionLifetimeRespectsKeyExpirationOk

`func (o *APIDefinition) GetSessionLifetimeRespectsKeyExpirationOk() (*bool, bool)`

GetSessionLifetimeRespectsKeyExpirationOk returns a tuple with the SessionLifetimeRespectsKeyExpiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionLifetimeRespectsKeyExpiration

`func (o *APIDefinition) SetSessionLifetimeRespectsKeyExpiration(v bool)`

SetSessionLifetimeRespectsKeyExpiration sets SessionLifetimeRespectsKeyExpiration field to given value.

### HasSessionLifetimeRespectsKeyExpiration

`func (o *APIDefinition) HasSessionLifetimeRespectsKeyExpiration() bool`

HasSessionLifetimeRespectsKeyExpiration returns a boolean if a field has been set.

### GetSessionProvider

`func (o *APIDefinition) GetSessionProvider() SessionProviderMeta`

GetSessionProvider returns the SessionProvider field if non-nil, zero value otherwise.

### GetSessionProviderOk

`func (o *APIDefinition) GetSessionProviderOk() (*SessionProviderMeta, bool)`

GetSessionProviderOk returns a tuple with the SessionProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionProvider

`func (o *APIDefinition) SetSessionProvider(v SessionProviderMeta)`

SetSessionProvider sets SessionProvider field to given value.

### HasSessionProvider

`func (o *APIDefinition) HasSessionProvider() bool`

HasSessionProvider returns a boolean if a field has been set.

### GetSlug

`func (o *APIDefinition) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *APIDefinition) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *APIDefinition) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *APIDefinition) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetStripAuthData

`func (o *APIDefinition) GetStripAuthData() bool`

GetStripAuthData returns the StripAuthData field if non-nil, zero value otherwise.

### GetStripAuthDataOk

`func (o *APIDefinition) GetStripAuthDataOk() (*bool, bool)`

GetStripAuthDataOk returns a tuple with the StripAuthData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStripAuthData

`func (o *APIDefinition) SetStripAuthData(v bool)`

SetStripAuthData sets StripAuthData field to given value.

### HasStripAuthData

`func (o *APIDefinition) HasStripAuthData() bool`

HasStripAuthData returns a boolean if a field has been set.

### GetTagHeaders

`func (o *APIDefinition) GetTagHeaders() []string`

GetTagHeaders returns the TagHeaders field if non-nil, zero value otherwise.

### GetTagHeadersOk

`func (o *APIDefinition) GetTagHeadersOk() (*[]string, bool)`

GetTagHeadersOk returns a tuple with the TagHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagHeaders

`func (o *APIDefinition) SetTagHeaders(v []string)`

SetTagHeaders sets TagHeaders field to given value.

### HasTagHeaders

`func (o *APIDefinition) HasTagHeaders() bool`

HasTagHeaders returns a boolean if a field has been set.

### SetTagHeadersNil

`func (o *APIDefinition) SetTagHeadersNil(b bool)`

 SetTagHeadersNil sets the value for TagHeaders to be an explicit nil

### UnsetTagHeaders
`func (o *APIDefinition) UnsetTagHeaders()`

UnsetTagHeaders ensures that no value is present for TagHeaders, not even an explicit nil
### GetTags

`func (o *APIDefinition) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *APIDefinition) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *APIDefinition) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *APIDefinition) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *APIDefinition) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *APIDefinition) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetTagsDisabled

`func (o *APIDefinition) GetTagsDisabled() bool`

GetTagsDisabled returns the TagsDisabled field if non-nil, zero value otherwise.

### GetTagsDisabledOk

`func (o *APIDefinition) GetTagsDisabledOk() (*bool, bool)`

GetTagsDisabledOk returns a tuple with the TagsDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagsDisabled

`func (o *APIDefinition) SetTagsDisabled(v bool)`

SetTagsDisabled sets TagsDisabled field to given value.

### HasTagsDisabled

`func (o *APIDefinition) HasTagsDisabled() bool`

HasTagsDisabled returns a boolean if a field has been set.

### GetUpstreamCertificates

`func (o *APIDefinition) GetUpstreamCertificates() map[string]string`

GetUpstreamCertificates returns the UpstreamCertificates field if non-nil, zero value otherwise.

### GetUpstreamCertificatesOk

`func (o *APIDefinition) GetUpstreamCertificatesOk() (*map[string]string, bool)`

GetUpstreamCertificatesOk returns a tuple with the UpstreamCertificates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpstreamCertificates

`func (o *APIDefinition) SetUpstreamCertificates(v map[string]string)`

SetUpstreamCertificates sets UpstreamCertificates field to given value.

### HasUpstreamCertificates

`func (o *APIDefinition) HasUpstreamCertificates() bool`

HasUpstreamCertificates returns a boolean if a field has been set.

### SetUpstreamCertificatesNil

`func (o *APIDefinition) SetUpstreamCertificatesNil(b bool)`

 SetUpstreamCertificatesNil sets the value for UpstreamCertificates to be an explicit nil

### UnsetUpstreamCertificates
`func (o *APIDefinition) UnsetUpstreamCertificates()`

UnsetUpstreamCertificates ensures that no value is present for UpstreamCertificates, not even an explicit nil
### GetUpstreamCertificatesDisabled

`func (o *APIDefinition) GetUpstreamCertificatesDisabled() bool`

GetUpstreamCertificatesDisabled returns the UpstreamCertificatesDisabled field if non-nil, zero value otherwise.

### GetUpstreamCertificatesDisabledOk

`func (o *APIDefinition) GetUpstreamCertificatesDisabledOk() (*bool, bool)`

GetUpstreamCertificatesDisabledOk returns a tuple with the UpstreamCertificatesDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpstreamCertificatesDisabled

`func (o *APIDefinition) SetUpstreamCertificatesDisabled(v bool)`

SetUpstreamCertificatesDisabled sets UpstreamCertificatesDisabled field to given value.

### HasUpstreamCertificatesDisabled

`func (o *APIDefinition) HasUpstreamCertificatesDisabled() bool`

HasUpstreamCertificatesDisabled returns a boolean if a field has been set.

### GetUptimeTests

`func (o *APIDefinition) GetUptimeTests() UptimeTests`

GetUptimeTests returns the UptimeTests field if non-nil, zero value otherwise.

### GetUptimeTestsOk

`func (o *APIDefinition) GetUptimeTestsOk() (*UptimeTests, bool)`

GetUptimeTestsOk returns a tuple with the UptimeTests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUptimeTests

`func (o *APIDefinition) SetUptimeTests(v UptimeTests)`

SetUptimeTests sets UptimeTests field to given value.

### HasUptimeTests

`func (o *APIDefinition) HasUptimeTests() bool`

HasUptimeTests returns a boolean if a field has been set.

### GetUseBasicAuth

`func (o *APIDefinition) GetUseBasicAuth() bool`

GetUseBasicAuth returns the UseBasicAuth field if non-nil, zero value otherwise.

### GetUseBasicAuthOk

`func (o *APIDefinition) GetUseBasicAuthOk() (*bool, bool)`

GetUseBasicAuthOk returns a tuple with the UseBasicAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseBasicAuth

`func (o *APIDefinition) SetUseBasicAuth(v bool)`

SetUseBasicAuth sets UseBasicAuth field to given value.

### HasUseBasicAuth

`func (o *APIDefinition) HasUseBasicAuth() bool`

HasUseBasicAuth returns a boolean if a field has been set.

### GetUseGoPluginAuth

`func (o *APIDefinition) GetUseGoPluginAuth() bool`

GetUseGoPluginAuth returns the UseGoPluginAuth field if non-nil, zero value otherwise.

### GetUseGoPluginAuthOk

`func (o *APIDefinition) GetUseGoPluginAuthOk() (*bool, bool)`

GetUseGoPluginAuthOk returns a tuple with the UseGoPluginAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseGoPluginAuth

`func (o *APIDefinition) SetUseGoPluginAuth(v bool)`

SetUseGoPluginAuth sets UseGoPluginAuth field to given value.

### HasUseGoPluginAuth

`func (o *APIDefinition) HasUseGoPluginAuth() bool`

HasUseGoPluginAuth returns a boolean if a field has been set.

### GetUseKeyless

`func (o *APIDefinition) GetUseKeyless() bool`

GetUseKeyless returns the UseKeyless field if non-nil, zero value otherwise.

### GetUseKeylessOk

`func (o *APIDefinition) GetUseKeylessOk() (*bool, bool)`

GetUseKeylessOk returns a tuple with the UseKeyless field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseKeyless

`func (o *APIDefinition) SetUseKeyless(v bool)`

SetUseKeyless sets UseKeyless field to given value.

### HasUseKeyless

`func (o *APIDefinition) HasUseKeyless() bool`

HasUseKeyless returns a boolean if a field has been set.

### GetUseMutualTlsAuth

`func (o *APIDefinition) GetUseMutualTlsAuth() bool`

GetUseMutualTlsAuth returns the UseMutualTlsAuth field if non-nil, zero value otherwise.

### GetUseMutualTlsAuthOk

`func (o *APIDefinition) GetUseMutualTlsAuthOk() (*bool, bool)`

GetUseMutualTlsAuthOk returns a tuple with the UseMutualTlsAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseMutualTlsAuth

`func (o *APIDefinition) SetUseMutualTlsAuth(v bool)`

SetUseMutualTlsAuth sets UseMutualTlsAuth field to given value.

### HasUseMutualTlsAuth

`func (o *APIDefinition) HasUseMutualTlsAuth() bool`

HasUseMutualTlsAuth returns a boolean if a field has been set.

### GetUseOauth2

`func (o *APIDefinition) GetUseOauth2() bool`

GetUseOauth2 returns the UseOauth2 field if non-nil, zero value otherwise.

### GetUseOauth2Ok

`func (o *APIDefinition) GetUseOauth2Ok() (*bool, bool)`

GetUseOauth2Ok returns a tuple with the UseOauth2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseOauth2

`func (o *APIDefinition) SetUseOauth2(v bool)`

SetUseOauth2 sets UseOauth2 field to given value.

### HasUseOauth2

`func (o *APIDefinition) HasUseOauth2() bool`

HasUseOauth2 returns a boolean if a field has been set.

### GetUseOpenid

`func (o *APIDefinition) GetUseOpenid() bool`

GetUseOpenid returns the UseOpenid field if non-nil, zero value otherwise.

### GetUseOpenidOk

`func (o *APIDefinition) GetUseOpenidOk() (*bool, bool)`

GetUseOpenidOk returns a tuple with the UseOpenid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseOpenid

`func (o *APIDefinition) SetUseOpenid(v bool)`

SetUseOpenid sets UseOpenid field to given value.

### HasUseOpenid

`func (o *APIDefinition) HasUseOpenid() bool`

HasUseOpenid returns a boolean if a field has been set.

### GetUseStandardAuth

`func (o *APIDefinition) GetUseStandardAuth() bool`

GetUseStandardAuth returns the UseStandardAuth field if non-nil, zero value otherwise.

### GetUseStandardAuthOk

`func (o *APIDefinition) GetUseStandardAuthOk() (*bool, bool)`

GetUseStandardAuthOk returns a tuple with the UseStandardAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseStandardAuth

`func (o *APIDefinition) SetUseStandardAuth(v bool)`

SetUseStandardAuth sets UseStandardAuth field to given value.

### HasUseStandardAuth

`func (o *APIDefinition) HasUseStandardAuth() bool`

HasUseStandardAuth returns a boolean if a field has been set.

### GetVersionData

`func (o *APIDefinition) GetVersionData() VersionData`

GetVersionData returns the VersionData field if non-nil, zero value otherwise.

### GetVersionDataOk

`func (o *APIDefinition) GetVersionDataOk() (*VersionData, bool)`

GetVersionDataOk returns a tuple with the VersionData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionData

`func (o *APIDefinition) SetVersionData(v VersionData)`

SetVersionData sets VersionData field to given value.

### HasVersionData

`func (o *APIDefinition) HasVersionData() bool`

HasVersionData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


