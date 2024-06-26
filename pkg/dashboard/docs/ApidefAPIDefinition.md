# ApidefAPIDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CORS** | Pointer to [**ApidefCORSConfig**](ApidefCORSConfig.md) |  | [optional] 
**Active** | Pointer to **bool** |  | [optional] 
**AllowedIps** | Pointer to **[]string** |  | [optional] 
**AnalyticsPlugin** | Pointer to [**ApidefAnalyticsPluginConfig**](ApidefAnalyticsPluginConfig.md) |  | [optional] 
**ApiId** | Pointer to **string** |  | [optional] 
**Auth** | Pointer to [**ApidefAuthConfig**](ApidefAuthConfig.md) |  | [optional] 
**AuthConfigs** | Pointer to [**map[string]ApidefAuthConfig**](ApidefAuthConfig.md) |  | [optional] 
**AuthProvider** | Pointer to [**ApidefAuthProviderMeta**](ApidefAuthProviderMeta.md) |  | [optional] 
**BaseIdentityProvidedBy** | Pointer to **string** |  | [optional] 
**BasicAuth** | Pointer to [**ApidefAPIDefinitionBasicAuth**](ApidefAPIDefinitionBasicAuth.md) |  | [optional] 
**BlacklistedIps** | Pointer to **[]string** |  | [optional] 
**CacheOptions** | Pointer to [**ApidefCacheOptions**](ApidefCacheOptions.md) |  | [optional] 
**CertificatePinningDisabled** | Pointer to **bool** |  | [optional] 
**Certificates** | Pointer to **[]string** |  | [optional] 
**ClientCertificates** | Pointer to **[]string** |  | [optional] 
**ConfigData** | Pointer to **map[string]interface{}** |  | [optional] 
**ConfigDataDisabled** | Pointer to **bool** |  | [optional] 
**CustomMiddleware** | Pointer to [**ApidefMiddlewareSection**](ApidefMiddlewareSection.md) |  | [optional] 
**CustomMiddlewareBundle** | Pointer to **string** |  | [optional] 
**CustomMiddlewareBundleDisabled** | Pointer to **bool** |  | [optional] 
**CustomPluginAuthEnabled** | Pointer to **bool** |  | [optional] 
**Definition** | Pointer to [**ApidefVersionDefinition**](ApidefVersionDefinition.md) |  | [optional] 
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
**EventHandlers** | Pointer to [**ApidefEventHandlerMetaConfig**](ApidefEventHandlerMetaConfig.md) |  | [optional] 
**Expiration** | Pointer to **string** |  | [optional] 
**ExpireAnalyticsAfter** | Pointer to **int32** |  | [optional] 
**ExternalOauth** | Pointer to [**ApidefExternalOAuth**](ApidefExternalOAuth.md) |  | [optional] 
**GlobalRateLimit** | Pointer to [**ApidefGlobalRateLimit**](ApidefGlobalRateLimit.md) |  | [optional] 
**Graphql** | Pointer to [**ApidefGraphQLConfig**](ApidefGraphQLConfig.md) |  | [optional] 
**HmacAllowedAlgorithms** | Pointer to **[]string** |  | [optional] 
**HmacAllowedClockSkew** | Pointer to **float32** |  | [optional] 
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
**Notifications** | Pointer to [**ApidefNotificationsManager**](ApidefNotificationsManager.md) |  | [optional] 
**OauthMeta** | Pointer to [**ApidefAPIDefinitionOauthMeta**](ApidefAPIDefinitionOauthMeta.md) |  | [optional] 
**OpenidOptions** | Pointer to [**ApidefOpenIDOptions**](ApidefOpenIDOptions.md) |  | [optional] 
**OrgId** | Pointer to **string** |  | [optional] 
**PinnedPublicKeys** | Pointer to **map[string]string** |  | [optional] 
**Protocol** | Pointer to **string** |  | [optional] 
**Proxy** | Pointer to [**ApidefProxyConfig**](ApidefProxyConfig.md) |  | [optional] 
**RequestSigning** | Pointer to [**ApidefRequestSigningMeta**](ApidefRequestSigningMeta.md) |  | [optional] 
**ResponseProcessors** | Pointer to [**[]ApidefResponseProcessor**](ApidefResponseProcessor.md) |  | [optional] 
**Scopes** | Pointer to [**ApidefScopes**](ApidefScopes.md) |  | [optional] 
**SessionLifetime** | Pointer to **int32** |  | [optional] 
**SessionLifetimeRespectsKeyExpiration** | Pointer to **bool** |  | [optional] 
**SessionProvider** | Pointer to [**ApidefSessionProviderMeta**](ApidefSessionProviderMeta.md) |  | [optional] 
**Slug** | Pointer to **string** |  | [optional] 
**StripAuthData** | Pointer to **bool** |  | [optional] 
**TagHeaders** | Pointer to **[]string** |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**TagsDisabled** | Pointer to **bool** |  | [optional] 
**UpstreamCertificates** | Pointer to **map[string]string** |  | [optional] 
**UpstreamCertificatesDisabled** | Pointer to **bool** |  | [optional] 
**UptimeTests** | Pointer to [**ApidefUptimeTests**](ApidefUptimeTests.md) |  | [optional] 
**UseBasicAuth** | Pointer to **bool** |  | [optional] 
**UseGoPluginAuth** | Pointer to **bool** |  | [optional] 
**UseKeyless** | Pointer to **bool** |  | [optional] 
**UseMutualTlsAuth** | Pointer to **bool** |  | [optional] 
**UseOauth2** | Pointer to **bool** |  | [optional] 
**UseOpenid** | Pointer to **bool** |  | [optional] 
**UseStandardAuth** | Pointer to **bool** |  | [optional] 
**VersionData** | Pointer to [**ApidefVersionData**](ApidefVersionData.md) |  | [optional] 

## Methods

### NewApidefAPIDefinition

`func NewApidefAPIDefinition() *ApidefAPIDefinition`

NewApidefAPIDefinition instantiates a new ApidefAPIDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApidefAPIDefinitionWithDefaults

`func NewApidefAPIDefinitionWithDefaults() *ApidefAPIDefinition`

NewApidefAPIDefinitionWithDefaults instantiates a new ApidefAPIDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCORS

`func (o *ApidefAPIDefinition) GetCORS() ApidefCORSConfig`

GetCORS returns the CORS field if non-nil, zero value otherwise.

### GetCORSOk

`func (o *ApidefAPIDefinition) GetCORSOk() (*ApidefCORSConfig, bool)`

GetCORSOk returns a tuple with the CORS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCORS

`func (o *ApidefAPIDefinition) SetCORS(v ApidefCORSConfig)`

SetCORS sets CORS field to given value.

### HasCORS

`func (o *ApidefAPIDefinition) HasCORS() bool`

HasCORS returns a boolean if a field has been set.

### GetActive

`func (o *ApidefAPIDefinition) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *ApidefAPIDefinition) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *ApidefAPIDefinition) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *ApidefAPIDefinition) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetAllowedIps

`func (o *ApidefAPIDefinition) GetAllowedIps() []string`

GetAllowedIps returns the AllowedIps field if non-nil, zero value otherwise.

### GetAllowedIpsOk

`func (o *ApidefAPIDefinition) GetAllowedIpsOk() (*[]string, bool)`

GetAllowedIpsOk returns a tuple with the AllowedIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedIps

`func (o *ApidefAPIDefinition) SetAllowedIps(v []string)`

SetAllowedIps sets AllowedIps field to given value.

### HasAllowedIps

`func (o *ApidefAPIDefinition) HasAllowedIps() bool`

HasAllowedIps returns a boolean if a field has been set.

### SetAllowedIpsNil

`func (o *ApidefAPIDefinition) SetAllowedIpsNil(b bool)`

 SetAllowedIpsNil sets the value for AllowedIps to be an explicit nil

### UnsetAllowedIps
`func (o *ApidefAPIDefinition) UnsetAllowedIps()`

UnsetAllowedIps ensures that no value is present for AllowedIps, not even an explicit nil
### GetAnalyticsPlugin

`func (o *ApidefAPIDefinition) GetAnalyticsPlugin() ApidefAnalyticsPluginConfig`

GetAnalyticsPlugin returns the AnalyticsPlugin field if non-nil, zero value otherwise.

### GetAnalyticsPluginOk

`func (o *ApidefAPIDefinition) GetAnalyticsPluginOk() (*ApidefAnalyticsPluginConfig, bool)`

GetAnalyticsPluginOk returns a tuple with the AnalyticsPlugin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalyticsPlugin

`func (o *ApidefAPIDefinition) SetAnalyticsPlugin(v ApidefAnalyticsPluginConfig)`

SetAnalyticsPlugin sets AnalyticsPlugin field to given value.

### HasAnalyticsPlugin

`func (o *ApidefAPIDefinition) HasAnalyticsPlugin() bool`

HasAnalyticsPlugin returns a boolean if a field has been set.

### GetApiId

`func (o *ApidefAPIDefinition) GetApiId() string`

GetApiId returns the ApiId field if non-nil, zero value otherwise.

### GetApiIdOk

`func (o *ApidefAPIDefinition) GetApiIdOk() (*string, bool)`

GetApiIdOk returns a tuple with the ApiId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiId

`func (o *ApidefAPIDefinition) SetApiId(v string)`

SetApiId sets ApiId field to given value.

### HasApiId

`func (o *ApidefAPIDefinition) HasApiId() bool`

HasApiId returns a boolean if a field has been set.

### GetAuth

`func (o *ApidefAPIDefinition) GetAuth() ApidefAuthConfig`

GetAuth returns the Auth field if non-nil, zero value otherwise.

### GetAuthOk

`func (o *ApidefAPIDefinition) GetAuthOk() (*ApidefAuthConfig, bool)`

GetAuthOk returns a tuple with the Auth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuth

`func (o *ApidefAPIDefinition) SetAuth(v ApidefAuthConfig)`

SetAuth sets Auth field to given value.

### HasAuth

`func (o *ApidefAPIDefinition) HasAuth() bool`

HasAuth returns a boolean if a field has been set.

### GetAuthConfigs

`func (o *ApidefAPIDefinition) GetAuthConfigs() map[string]ApidefAuthConfig`

GetAuthConfigs returns the AuthConfigs field if non-nil, zero value otherwise.

### GetAuthConfigsOk

`func (o *ApidefAPIDefinition) GetAuthConfigsOk() (*map[string]ApidefAuthConfig, bool)`

GetAuthConfigsOk returns a tuple with the AuthConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthConfigs

`func (o *ApidefAPIDefinition) SetAuthConfigs(v map[string]ApidefAuthConfig)`

SetAuthConfigs sets AuthConfigs field to given value.

### HasAuthConfigs

`func (o *ApidefAPIDefinition) HasAuthConfigs() bool`

HasAuthConfigs returns a boolean if a field has been set.

### SetAuthConfigsNil

`func (o *ApidefAPIDefinition) SetAuthConfigsNil(b bool)`

 SetAuthConfigsNil sets the value for AuthConfigs to be an explicit nil

### UnsetAuthConfigs
`func (o *ApidefAPIDefinition) UnsetAuthConfigs()`

UnsetAuthConfigs ensures that no value is present for AuthConfigs, not even an explicit nil
### GetAuthProvider

`func (o *ApidefAPIDefinition) GetAuthProvider() ApidefAuthProviderMeta`

GetAuthProvider returns the AuthProvider field if non-nil, zero value otherwise.

### GetAuthProviderOk

`func (o *ApidefAPIDefinition) GetAuthProviderOk() (*ApidefAuthProviderMeta, bool)`

GetAuthProviderOk returns a tuple with the AuthProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthProvider

`func (o *ApidefAPIDefinition) SetAuthProvider(v ApidefAuthProviderMeta)`

SetAuthProvider sets AuthProvider field to given value.

### HasAuthProvider

`func (o *ApidefAPIDefinition) HasAuthProvider() bool`

HasAuthProvider returns a boolean if a field has been set.

### GetBaseIdentityProvidedBy

`func (o *ApidefAPIDefinition) GetBaseIdentityProvidedBy() string`

GetBaseIdentityProvidedBy returns the BaseIdentityProvidedBy field if non-nil, zero value otherwise.

### GetBaseIdentityProvidedByOk

`func (o *ApidefAPIDefinition) GetBaseIdentityProvidedByOk() (*string, bool)`

GetBaseIdentityProvidedByOk returns a tuple with the BaseIdentityProvidedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseIdentityProvidedBy

`func (o *ApidefAPIDefinition) SetBaseIdentityProvidedBy(v string)`

SetBaseIdentityProvidedBy sets BaseIdentityProvidedBy field to given value.

### HasBaseIdentityProvidedBy

`func (o *ApidefAPIDefinition) HasBaseIdentityProvidedBy() bool`

HasBaseIdentityProvidedBy returns a boolean if a field has been set.

### GetBasicAuth

`func (o *ApidefAPIDefinition) GetBasicAuth() ApidefAPIDefinitionBasicAuth`

GetBasicAuth returns the BasicAuth field if non-nil, zero value otherwise.

### GetBasicAuthOk

`func (o *ApidefAPIDefinition) GetBasicAuthOk() (*ApidefAPIDefinitionBasicAuth, bool)`

GetBasicAuthOk returns a tuple with the BasicAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasicAuth

`func (o *ApidefAPIDefinition) SetBasicAuth(v ApidefAPIDefinitionBasicAuth)`

SetBasicAuth sets BasicAuth field to given value.

### HasBasicAuth

`func (o *ApidefAPIDefinition) HasBasicAuth() bool`

HasBasicAuth returns a boolean if a field has been set.

### GetBlacklistedIps

`func (o *ApidefAPIDefinition) GetBlacklistedIps() []string`

GetBlacklistedIps returns the BlacklistedIps field if non-nil, zero value otherwise.

### GetBlacklistedIpsOk

`func (o *ApidefAPIDefinition) GetBlacklistedIpsOk() (*[]string, bool)`

GetBlacklistedIpsOk returns a tuple with the BlacklistedIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlacklistedIps

`func (o *ApidefAPIDefinition) SetBlacklistedIps(v []string)`

SetBlacklistedIps sets BlacklistedIps field to given value.

### HasBlacklistedIps

`func (o *ApidefAPIDefinition) HasBlacklistedIps() bool`

HasBlacklistedIps returns a boolean if a field has been set.

### SetBlacklistedIpsNil

`func (o *ApidefAPIDefinition) SetBlacklistedIpsNil(b bool)`

 SetBlacklistedIpsNil sets the value for BlacklistedIps to be an explicit nil

### UnsetBlacklistedIps
`func (o *ApidefAPIDefinition) UnsetBlacklistedIps()`

UnsetBlacklistedIps ensures that no value is present for BlacklistedIps, not even an explicit nil
### GetCacheOptions

`func (o *ApidefAPIDefinition) GetCacheOptions() ApidefCacheOptions`

GetCacheOptions returns the CacheOptions field if non-nil, zero value otherwise.

### GetCacheOptionsOk

`func (o *ApidefAPIDefinition) GetCacheOptionsOk() (*ApidefCacheOptions, bool)`

GetCacheOptionsOk returns a tuple with the CacheOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheOptions

`func (o *ApidefAPIDefinition) SetCacheOptions(v ApidefCacheOptions)`

SetCacheOptions sets CacheOptions field to given value.

### HasCacheOptions

`func (o *ApidefAPIDefinition) HasCacheOptions() bool`

HasCacheOptions returns a boolean if a field has been set.

### GetCertificatePinningDisabled

`func (o *ApidefAPIDefinition) GetCertificatePinningDisabled() bool`

GetCertificatePinningDisabled returns the CertificatePinningDisabled field if non-nil, zero value otherwise.

### GetCertificatePinningDisabledOk

`func (o *ApidefAPIDefinition) GetCertificatePinningDisabledOk() (*bool, bool)`

GetCertificatePinningDisabledOk returns a tuple with the CertificatePinningDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificatePinningDisabled

`func (o *ApidefAPIDefinition) SetCertificatePinningDisabled(v bool)`

SetCertificatePinningDisabled sets CertificatePinningDisabled field to given value.

### HasCertificatePinningDisabled

`func (o *ApidefAPIDefinition) HasCertificatePinningDisabled() bool`

HasCertificatePinningDisabled returns a boolean if a field has been set.

### GetCertificates

`func (o *ApidefAPIDefinition) GetCertificates() []string`

GetCertificates returns the Certificates field if non-nil, zero value otherwise.

### GetCertificatesOk

`func (o *ApidefAPIDefinition) GetCertificatesOk() (*[]string, bool)`

GetCertificatesOk returns a tuple with the Certificates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificates

`func (o *ApidefAPIDefinition) SetCertificates(v []string)`

SetCertificates sets Certificates field to given value.

### HasCertificates

`func (o *ApidefAPIDefinition) HasCertificates() bool`

HasCertificates returns a boolean if a field has been set.

### SetCertificatesNil

`func (o *ApidefAPIDefinition) SetCertificatesNil(b bool)`

 SetCertificatesNil sets the value for Certificates to be an explicit nil

### UnsetCertificates
`func (o *ApidefAPIDefinition) UnsetCertificates()`

UnsetCertificates ensures that no value is present for Certificates, not even an explicit nil
### GetClientCertificates

`func (o *ApidefAPIDefinition) GetClientCertificates() []string`

GetClientCertificates returns the ClientCertificates field if non-nil, zero value otherwise.

### GetClientCertificatesOk

`func (o *ApidefAPIDefinition) GetClientCertificatesOk() (*[]string, bool)`

GetClientCertificatesOk returns a tuple with the ClientCertificates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientCertificates

`func (o *ApidefAPIDefinition) SetClientCertificates(v []string)`

SetClientCertificates sets ClientCertificates field to given value.

### HasClientCertificates

`func (o *ApidefAPIDefinition) HasClientCertificates() bool`

HasClientCertificates returns a boolean if a field has been set.

### SetClientCertificatesNil

`func (o *ApidefAPIDefinition) SetClientCertificatesNil(b bool)`

 SetClientCertificatesNil sets the value for ClientCertificates to be an explicit nil

### UnsetClientCertificates
`func (o *ApidefAPIDefinition) UnsetClientCertificates()`

UnsetClientCertificates ensures that no value is present for ClientCertificates, not even an explicit nil
### GetConfigData

`func (o *ApidefAPIDefinition) GetConfigData() map[string]interface{}`

GetConfigData returns the ConfigData field if non-nil, zero value otherwise.

### GetConfigDataOk

`func (o *ApidefAPIDefinition) GetConfigDataOk() (*map[string]interface{}, bool)`

GetConfigDataOk returns a tuple with the ConfigData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigData

`func (o *ApidefAPIDefinition) SetConfigData(v map[string]interface{})`

SetConfigData sets ConfigData field to given value.

### HasConfigData

`func (o *ApidefAPIDefinition) HasConfigData() bool`

HasConfigData returns a boolean if a field has been set.

### SetConfigDataNil

`func (o *ApidefAPIDefinition) SetConfigDataNil(b bool)`

 SetConfigDataNil sets the value for ConfigData to be an explicit nil

### UnsetConfigData
`func (o *ApidefAPIDefinition) UnsetConfigData()`

UnsetConfigData ensures that no value is present for ConfigData, not even an explicit nil
### GetConfigDataDisabled

`func (o *ApidefAPIDefinition) GetConfigDataDisabled() bool`

GetConfigDataDisabled returns the ConfigDataDisabled field if non-nil, zero value otherwise.

### GetConfigDataDisabledOk

`func (o *ApidefAPIDefinition) GetConfigDataDisabledOk() (*bool, bool)`

GetConfigDataDisabledOk returns a tuple with the ConfigDataDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigDataDisabled

`func (o *ApidefAPIDefinition) SetConfigDataDisabled(v bool)`

SetConfigDataDisabled sets ConfigDataDisabled field to given value.

### HasConfigDataDisabled

`func (o *ApidefAPIDefinition) HasConfigDataDisabled() bool`

HasConfigDataDisabled returns a boolean if a field has been set.

### GetCustomMiddleware

`func (o *ApidefAPIDefinition) GetCustomMiddleware() ApidefMiddlewareSection`

GetCustomMiddleware returns the CustomMiddleware field if non-nil, zero value otherwise.

### GetCustomMiddlewareOk

`func (o *ApidefAPIDefinition) GetCustomMiddlewareOk() (*ApidefMiddlewareSection, bool)`

GetCustomMiddlewareOk returns a tuple with the CustomMiddleware field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomMiddleware

`func (o *ApidefAPIDefinition) SetCustomMiddleware(v ApidefMiddlewareSection)`

SetCustomMiddleware sets CustomMiddleware field to given value.

### HasCustomMiddleware

`func (o *ApidefAPIDefinition) HasCustomMiddleware() bool`

HasCustomMiddleware returns a boolean if a field has been set.

### GetCustomMiddlewareBundle

`func (o *ApidefAPIDefinition) GetCustomMiddlewareBundle() string`

GetCustomMiddlewareBundle returns the CustomMiddlewareBundle field if non-nil, zero value otherwise.

### GetCustomMiddlewareBundleOk

`func (o *ApidefAPIDefinition) GetCustomMiddlewareBundleOk() (*string, bool)`

GetCustomMiddlewareBundleOk returns a tuple with the CustomMiddlewareBundle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomMiddlewareBundle

`func (o *ApidefAPIDefinition) SetCustomMiddlewareBundle(v string)`

SetCustomMiddlewareBundle sets CustomMiddlewareBundle field to given value.

### HasCustomMiddlewareBundle

`func (o *ApidefAPIDefinition) HasCustomMiddlewareBundle() bool`

HasCustomMiddlewareBundle returns a boolean if a field has been set.

### GetCustomMiddlewareBundleDisabled

`func (o *ApidefAPIDefinition) GetCustomMiddlewareBundleDisabled() bool`

GetCustomMiddlewareBundleDisabled returns the CustomMiddlewareBundleDisabled field if non-nil, zero value otherwise.

### GetCustomMiddlewareBundleDisabledOk

`func (o *ApidefAPIDefinition) GetCustomMiddlewareBundleDisabledOk() (*bool, bool)`

GetCustomMiddlewareBundleDisabledOk returns a tuple with the CustomMiddlewareBundleDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomMiddlewareBundleDisabled

`func (o *ApidefAPIDefinition) SetCustomMiddlewareBundleDisabled(v bool)`

SetCustomMiddlewareBundleDisabled sets CustomMiddlewareBundleDisabled field to given value.

### HasCustomMiddlewareBundleDisabled

`func (o *ApidefAPIDefinition) HasCustomMiddlewareBundleDisabled() bool`

HasCustomMiddlewareBundleDisabled returns a boolean if a field has been set.

### GetCustomPluginAuthEnabled

`func (o *ApidefAPIDefinition) GetCustomPluginAuthEnabled() bool`

GetCustomPluginAuthEnabled returns the CustomPluginAuthEnabled field if non-nil, zero value otherwise.

### GetCustomPluginAuthEnabledOk

`func (o *ApidefAPIDefinition) GetCustomPluginAuthEnabledOk() (*bool, bool)`

GetCustomPluginAuthEnabledOk returns a tuple with the CustomPluginAuthEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomPluginAuthEnabled

`func (o *ApidefAPIDefinition) SetCustomPluginAuthEnabled(v bool)`

SetCustomPluginAuthEnabled sets CustomPluginAuthEnabled field to given value.

### HasCustomPluginAuthEnabled

`func (o *ApidefAPIDefinition) HasCustomPluginAuthEnabled() bool`

HasCustomPluginAuthEnabled returns a boolean if a field has been set.

### GetDefinition

`func (o *ApidefAPIDefinition) GetDefinition() ApidefVersionDefinition`

GetDefinition returns the Definition field if non-nil, zero value otherwise.

### GetDefinitionOk

`func (o *ApidefAPIDefinition) GetDefinitionOk() (*ApidefVersionDefinition, bool)`

GetDefinitionOk returns a tuple with the Definition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefinition

`func (o *ApidefAPIDefinition) SetDefinition(v ApidefVersionDefinition)`

SetDefinition sets Definition field to given value.

### HasDefinition

`func (o *ApidefAPIDefinition) HasDefinition() bool`

HasDefinition returns a boolean if a field has been set.

### GetDetailedTracing

`func (o *ApidefAPIDefinition) GetDetailedTracing() bool`

GetDetailedTracing returns the DetailedTracing field if non-nil, zero value otherwise.

### GetDetailedTracingOk

`func (o *ApidefAPIDefinition) GetDetailedTracingOk() (*bool, bool)`

GetDetailedTracingOk returns a tuple with the DetailedTracing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetailedTracing

`func (o *ApidefAPIDefinition) SetDetailedTracing(v bool)`

SetDetailedTracing sets DetailedTracing field to given value.

### HasDetailedTracing

`func (o *ApidefAPIDefinition) HasDetailedTracing() bool`

HasDetailedTracing returns a boolean if a field has been set.

### GetDisableQuota

`func (o *ApidefAPIDefinition) GetDisableQuota() bool`

GetDisableQuota returns the DisableQuota field if non-nil, zero value otherwise.

### GetDisableQuotaOk

`func (o *ApidefAPIDefinition) GetDisableQuotaOk() (*bool, bool)`

GetDisableQuotaOk returns a tuple with the DisableQuota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableQuota

`func (o *ApidefAPIDefinition) SetDisableQuota(v bool)`

SetDisableQuota sets DisableQuota field to given value.

### HasDisableQuota

`func (o *ApidefAPIDefinition) HasDisableQuota() bool`

HasDisableQuota returns a boolean if a field has been set.

### GetDisableRateLimit

`func (o *ApidefAPIDefinition) GetDisableRateLimit() bool`

GetDisableRateLimit returns the DisableRateLimit field if non-nil, zero value otherwise.

### GetDisableRateLimitOk

`func (o *ApidefAPIDefinition) GetDisableRateLimitOk() (*bool, bool)`

GetDisableRateLimitOk returns a tuple with the DisableRateLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableRateLimit

`func (o *ApidefAPIDefinition) SetDisableRateLimit(v bool)`

SetDisableRateLimit sets DisableRateLimit field to given value.

### HasDisableRateLimit

`func (o *ApidefAPIDefinition) HasDisableRateLimit() bool`

HasDisableRateLimit returns a boolean if a field has been set.

### GetDoNotTrack

`func (o *ApidefAPIDefinition) GetDoNotTrack() bool`

GetDoNotTrack returns the DoNotTrack field if non-nil, zero value otherwise.

### GetDoNotTrackOk

`func (o *ApidefAPIDefinition) GetDoNotTrackOk() (*bool, bool)`

GetDoNotTrackOk returns a tuple with the DoNotTrack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoNotTrack

`func (o *ApidefAPIDefinition) SetDoNotTrack(v bool)`

SetDoNotTrack sets DoNotTrack field to given value.

### HasDoNotTrack

`func (o *ApidefAPIDefinition) HasDoNotTrack() bool`

HasDoNotTrack returns a boolean if a field has been set.

### GetDomain

`func (o *ApidefAPIDefinition) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *ApidefAPIDefinition) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *ApidefAPIDefinition) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *ApidefAPIDefinition) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetDomainDisabled

`func (o *ApidefAPIDefinition) GetDomainDisabled() bool`

GetDomainDisabled returns the DomainDisabled field if non-nil, zero value otherwise.

### GetDomainDisabledOk

`func (o *ApidefAPIDefinition) GetDomainDisabledOk() (*bool, bool)`

GetDomainDisabledOk returns a tuple with the DomainDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainDisabled

`func (o *ApidefAPIDefinition) SetDomainDisabled(v bool)`

SetDomainDisabled sets DomainDisabled field to given value.

### HasDomainDisabled

`func (o *ApidefAPIDefinition) HasDomainDisabled() bool`

HasDomainDisabled returns a boolean if a field has been set.

### GetDontSetQuotaOnCreate

`func (o *ApidefAPIDefinition) GetDontSetQuotaOnCreate() bool`

GetDontSetQuotaOnCreate returns the DontSetQuotaOnCreate field if non-nil, zero value otherwise.

### GetDontSetQuotaOnCreateOk

`func (o *ApidefAPIDefinition) GetDontSetQuotaOnCreateOk() (*bool, bool)`

GetDontSetQuotaOnCreateOk returns a tuple with the DontSetQuotaOnCreate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDontSetQuotaOnCreate

`func (o *ApidefAPIDefinition) SetDontSetQuotaOnCreate(v bool)`

SetDontSetQuotaOnCreate sets DontSetQuotaOnCreate field to given value.

### HasDontSetQuotaOnCreate

`func (o *ApidefAPIDefinition) HasDontSetQuotaOnCreate() bool`

HasDontSetQuotaOnCreate returns a boolean if a field has been set.

### GetEnableBatchRequestSupport

`func (o *ApidefAPIDefinition) GetEnableBatchRequestSupport() bool`

GetEnableBatchRequestSupport returns the EnableBatchRequestSupport field if non-nil, zero value otherwise.

### GetEnableBatchRequestSupportOk

`func (o *ApidefAPIDefinition) GetEnableBatchRequestSupportOk() (*bool, bool)`

GetEnableBatchRequestSupportOk returns a tuple with the EnableBatchRequestSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableBatchRequestSupport

`func (o *ApidefAPIDefinition) SetEnableBatchRequestSupport(v bool)`

SetEnableBatchRequestSupport sets EnableBatchRequestSupport field to given value.

### HasEnableBatchRequestSupport

`func (o *ApidefAPIDefinition) HasEnableBatchRequestSupport() bool`

HasEnableBatchRequestSupport returns a boolean if a field has been set.

### GetEnableContextVars

`func (o *ApidefAPIDefinition) GetEnableContextVars() bool`

GetEnableContextVars returns the EnableContextVars field if non-nil, zero value otherwise.

### GetEnableContextVarsOk

`func (o *ApidefAPIDefinition) GetEnableContextVarsOk() (*bool, bool)`

GetEnableContextVarsOk returns a tuple with the EnableContextVars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableContextVars

`func (o *ApidefAPIDefinition) SetEnableContextVars(v bool)`

SetEnableContextVars sets EnableContextVars field to given value.

### HasEnableContextVars

`func (o *ApidefAPIDefinition) HasEnableContextVars() bool`

HasEnableContextVars returns a boolean if a field has been set.

### GetEnableCoprocessAuth

`func (o *ApidefAPIDefinition) GetEnableCoprocessAuth() bool`

GetEnableCoprocessAuth returns the EnableCoprocessAuth field if non-nil, zero value otherwise.

### GetEnableCoprocessAuthOk

`func (o *ApidefAPIDefinition) GetEnableCoprocessAuthOk() (*bool, bool)`

GetEnableCoprocessAuthOk returns a tuple with the EnableCoprocessAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableCoprocessAuth

`func (o *ApidefAPIDefinition) SetEnableCoprocessAuth(v bool)`

SetEnableCoprocessAuth sets EnableCoprocessAuth field to given value.

### HasEnableCoprocessAuth

`func (o *ApidefAPIDefinition) HasEnableCoprocessAuth() bool`

HasEnableCoprocessAuth returns a boolean if a field has been set.

### GetEnableDetailedRecording

`func (o *ApidefAPIDefinition) GetEnableDetailedRecording() bool`

GetEnableDetailedRecording returns the EnableDetailedRecording field if non-nil, zero value otherwise.

### GetEnableDetailedRecordingOk

`func (o *ApidefAPIDefinition) GetEnableDetailedRecordingOk() (*bool, bool)`

GetEnableDetailedRecordingOk returns a tuple with the EnableDetailedRecording field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableDetailedRecording

`func (o *ApidefAPIDefinition) SetEnableDetailedRecording(v bool)`

SetEnableDetailedRecording sets EnableDetailedRecording field to given value.

### HasEnableDetailedRecording

`func (o *ApidefAPIDefinition) HasEnableDetailedRecording() bool`

HasEnableDetailedRecording returns a boolean if a field has been set.

### GetEnableIpBlacklisting

`func (o *ApidefAPIDefinition) GetEnableIpBlacklisting() bool`

GetEnableIpBlacklisting returns the EnableIpBlacklisting field if non-nil, zero value otherwise.

### GetEnableIpBlacklistingOk

`func (o *ApidefAPIDefinition) GetEnableIpBlacklistingOk() (*bool, bool)`

GetEnableIpBlacklistingOk returns a tuple with the EnableIpBlacklisting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableIpBlacklisting

`func (o *ApidefAPIDefinition) SetEnableIpBlacklisting(v bool)`

SetEnableIpBlacklisting sets EnableIpBlacklisting field to given value.

### HasEnableIpBlacklisting

`func (o *ApidefAPIDefinition) HasEnableIpBlacklisting() bool`

HasEnableIpBlacklisting returns a boolean if a field has been set.

### GetEnableIpWhitelisting

`func (o *ApidefAPIDefinition) GetEnableIpWhitelisting() bool`

GetEnableIpWhitelisting returns the EnableIpWhitelisting field if non-nil, zero value otherwise.

### GetEnableIpWhitelistingOk

`func (o *ApidefAPIDefinition) GetEnableIpWhitelistingOk() (*bool, bool)`

GetEnableIpWhitelistingOk returns a tuple with the EnableIpWhitelisting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableIpWhitelisting

`func (o *ApidefAPIDefinition) SetEnableIpWhitelisting(v bool)`

SetEnableIpWhitelisting sets EnableIpWhitelisting field to given value.

### HasEnableIpWhitelisting

`func (o *ApidefAPIDefinition) HasEnableIpWhitelisting() bool`

HasEnableIpWhitelisting returns a boolean if a field has been set.

### GetEnableJwt

`func (o *ApidefAPIDefinition) GetEnableJwt() bool`

GetEnableJwt returns the EnableJwt field if non-nil, zero value otherwise.

### GetEnableJwtOk

`func (o *ApidefAPIDefinition) GetEnableJwtOk() (*bool, bool)`

GetEnableJwtOk returns a tuple with the EnableJwt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableJwt

`func (o *ApidefAPIDefinition) SetEnableJwt(v bool)`

SetEnableJwt sets EnableJwt field to given value.

### HasEnableJwt

`func (o *ApidefAPIDefinition) HasEnableJwt() bool`

HasEnableJwt returns a boolean if a field has been set.

### GetEnableProxyProtocol

`func (o *ApidefAPIDefinition) GetEnableProxyProtocol() bool`

GetEnableProxyProtocol returns the EnableProxyProtocol field if non-nil, zero value otherwise.

### GetEnableProxyProtocolOk

`func (o *ApidefAPIDefinition) GetEnableProxyProtocolOk() (*bool, bool)`

GetEnableProxyProtocolOk returns a tuple with the EnableProxyProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableProxyProtocol

`func (o *ApidefAPIDefinition) SetEnableProxyProtocol(v bool)`

SetEnableProxyProtocol sets EnableProxyProtocol field to given value.

### HasEnableProxyProtocol

`func (o *ApidefAPIDefinition) HasEnableProxyProtocol() bool`

HasEnableProxyProtocol returns a boolean if a field has been set.

### GetEnableSignatureChecking

`func (o *ApidefAPIDefinition) GetEnableSignatureChecking() bool`

GetEnableSignatureChecking returns the EnableSignatureChecking field if non-nil, zero value otherwise.

### GetEnableSignatureCheckingOk

`func (o *ApidefAPIDefinition) GetEnableSignatureCheckingOk() (*bool, bool)`

GetEnableSignatureCheckingOk returns a tuple with the EnableSignatureChecking field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableSignatureChecking

`func (o *ApidefAPIDefinition) SetEnableSignatureChecking(v bool)`

SetEnableSignatureChecking sets EnableSignatureChecking field to given value.

### HasEnableSignatureChecking

`func (o *ApidefAPIDefinition) HasEnableSignatureChecking() bool`

HasEnableSignatureChecking returns a boolean if a field has been set.

### GetEventHandlers

`func (o *ApidefAPIDefinition) GetEventHandlers() ApidefEventHandlerMetaConfig`

GetEventHandlers returns the EventHandlers field if non-nil, zero value otherwise.

### GetEventHandlersOk

`func (o *ApidefAPIDefinition) GetEventHandlersOk() (*ApidefEventHandlerMetaConfig, bool)`

GetEventHandlersOk returns a tuple with the EventHandlers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventHandlers

`func (o *ApidefAPIDefinition) SetEventHandlers(v ApidefEventHandlerMetaConfig)`

SetEventHandlers sets EventHandlers field to given value.

### HasEventHandlers

`func (o *ApidefAPIDefinition) HasEventHandlers() bool`

HasEventHandlers returns a boolean if a field has been set.

### GetExpiration

`func (o *ApidefAPIDefinition) GetExpiration() string`

GetExpiration returns the Expiration field if non-nil, zero value otherwise.

### GetExpirationOk

`func (o *ApidefAPIDefinition) GetExpirationOk() (*string, bool)`

GetExpirationOk returns a tuple with the Expiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiration

`func (o *ApidefAPIDefinition) SetExpiration(v string)`

SetExpiration sets Expiration field to given value.

### HasExpiration

`func (o *ApidefAPIDefinition) HasExpiration() bool`

HasExpiration returns a boolean if a field has been set.

### GetExpireAnalyticsAfter

`func (o *ApidefAPIDefinition) GetExpireAnalyticsAfter() int32`

GetExpireAnalyticsAfter returns the ExpireAnalyticsAfter field if non-nil, zero value otherwise.

### GetExpireAnalyticsAfterOk

`func (o *ApidefAPIDefinition) GetExpireAnalyticsAfterOk() (*int32, bool)`

GetExpireAnalyticsAfterOk returns a tuple with the ExpireAnalyticsAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpireAnalyticsAfter

`func (o *ApidefAPIDefinition) SetExpireAnalyticsAfter(v int32)`

SetExpireAnalyticsAfter sets ExpireAnalyticsAfter field to given value.

### HasExpireAnalyticsAfter

`func (o *ApidefAPIDefinition) HasExpireAnalyticsAfter() bool`

HasExpireAnalyticsAfter returns a boolean if a field has been set.

### GetExternalOauth

`func (o *ApidefAPIDefinition) GetExternalOauth() ApidefExternalOAuth`

GetExternalOauth returns the ExternalOauth field if non-nil, zero value otherwise.

### GetExternalOauthOk

`func (o *ApidefAPIDefinition) GetExternalOauthOk() (*ApidefExternalOAuth, bool)`

GetExternalOauthOk returns a tuple with the ExternalOauth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalOauth

`func (o *ApidefAPIDefinition) SetExternalOauth(v ApidefExternalOAuth)`

SetExternalOauth sets ExternalOauth field to given value.

### HasExternalOauth

`func (o *ApidefAPIDefinition) HasExternalOauth() bool`

HasExternalOauth returns a boolean if a field has been set.

### GetGlobalRateLimit

`func (o *ApidefAPIDefinition) GetGlobalRateLimit() ApidefGlobalRateLimit`

GetGlobalRateLimit returns the GlobalRateLimit field if non-nil, zero value otherwise.

### GetGlobalRateLimitOk

`func (o *ApidefAPIDefinition) GetGlobalRateLimitOk() (*ApidefGlobalRateLimit, bool)`

GetGlobalRateLimitOk returns a tuple with the GlobalRateLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalRateLimit

`func (o *ApidefAPIDefinition) SetGlobalRateLimit(v ApidefGlobalRateLimit)`

SetGlobalRateLimit sets GlobalRateLimit field to given value.

### HasGlobalRateLimit

`func (o *ApidefAPIDefinition) HasGlobalRateLimit() bool`

HasGlobalRateLimit returns a boolean if a field has been set.

### GetGraphql

`func (o *ApidefAPIDefinition) GetGraphql() ApidefGraphQLConfig`

GetGraphql returns the Graphql field if non-nil, zero value otherwise.

### GetGraphqlOk

`func (o *ApidefAPIDefinition) GetGraphqlOk() (*ApidefGraphQLConfig, bool)`

GetGraphqlOk returns a tuple with the Graphql field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGraphql

`func (o *ApidefAPIDefinition) SetGraphql(v ApidefGraphQLConfig)`

SetGraphql sets Graphql field to given value.

### HasGraphql

`func (o *ApidefAPIDefinition) HasGraphql() bool`

HasGraphql returns a boolean if a field has been set.

### GetHmacAllowedAlgorithms

`func (o *ApidefAPIDefinition) GetHmacAllowedAlgorithms() []string`

GetHmacAllowedAlgorithms returns the HmacAllowedAlgorithms field if non-nil, zero value otherwise.

### GetHmacAllowedAlgorithmsOk

`func (o *ApidefAPIDefinition) GetHmacAllowedAlgorithmsOk() (*[]string, bool)`

GetHmacAllowedAlgorithmsOk returns a tuple with the HmacAllowedAlgorithms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHmacAllowedAlgorithms

`func (o *ApidefAPIDefinition) SetHmacAllowedAlgorithms(v []string)`

SetHmacAllowedAlgorithms sets HmacAllowedAlgorithms field to given value.

### HasHmacAllowedAlgorithms

`func (o *ApidefAPIDefinition) HasHmacAllowedAlgorithms() bool`

HasHmacAllowedAlgorithms returns a boolean if a field has been set.

### SetHmacAllowedAlgorithmsNil

`func (o *ApidefAPIDefinition) SetHmacAllowedAlgorithmsNil(b bool)`

 SetHmacAllowedAlgorithmsNil sets the value for HmacAllowedAlgorithms to be an explicit nil

### UnsetHmacAllowedAlgorithms
`func (o *ApidefAPIDefinition) UnsetHmacAllowedAlgorithms()`

UnsetHmacAllowedAlgorithms ensures that no value is present for HmacAllowedAlgorithms, not even an explicit nil
### GetHmacAllowedClockSkew

`func (o *ApidefAPIDefinition) GetHmacAllowedClockSkew() float32`

GetHmacAllowedClockSkew returns the HmacAllowedClockSkew field if non-nil, zero value otherwise.

### GetHmacAllowedClockSkewOk

`func (o *ApidefAPIDefinition) GetHmacAllowedClockSkewOk() (*float32, bool)`

GetHmacAllowedClockSkewOk returns a tuple with the HmacAllowedClockSkew field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHmacAllowedClockSkew

`func (o *ApidefAPIDefinition) SetHmacAllowedClockSkew(v float32)`

SetHmacAllowedClockSkew sets HmacAllowedClockSkew field to given value.

### HasHmacAllowedClockSkew

`func (o *ApidefAPIDefinition) HasHmacAllowedClockSkew() bool`

HasHmacAllowedClockSkew returns a boolean if a field has been set.

### GetId

`func (o *ApidefAPIDefinition) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApidefAPIDefinition) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApidefAPIDefinition) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ApidefAPIDefinition) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIdpClientIdMappingDisabled

`func (o *ApidefAPIDefinition) GetIdpClientIdMappingDisabled() bool`

GetIdpClientIdMappingDisabled returns the IdpClientIdMappingDisabled field if non-nil, zero value otherwise.

### GetIdpClientIdMappingDisabledOk

`func (o *ApidefAPIDefinition) GetIdpClientIdMappingDisabledOk() (*bool, bool)`

GetIdpClientIdMappingDisabledOk returns a tuple with the IdpClientIdMappingDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpClientIdMappingDisabled

`func (o *ApidefAPIDefinition) SetIdpClientIdMappingDisabled(v bool)`

SetIdpClientIdMappingDisabled sets IdpClientIdMappingDisabled field to given value.

### HasIdpClientIdMappingDisabled

`func (o *ApidefAPIDefinition) HasIdpClientIdMappingDisabled() bool`

HasIdpClientIdMappingDisabled returns a boolean if a field has been set.

### GetInternal

`func (o *ApidefAPIDefinition) GetInternal() bool`

GetInternal returns the Internal field if non-nil, zero value otherwise.

### GetInternalOk

`func (o *ApidefAPIDefinition) GetInternalOk() (*bool, bool)`

GetInternalOk returns a tuple with the Internal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternal

`func (o *ApidefAPIDefinition) SetInternal(v bool)`

SetInternal sets Internal field to given value.

### HasInternal

`func (o *ApidefAPIDefinition) HasInternal() bool`

HasInternal returns a boolean if a field has been set.

### GetIsOas

`func (o *ApidefAPIDefinition) GetIsOas() bool`

GetIsOas returns the IsOas field if non-nil, zero value otherwise.

### GetIsOasOk

`func (o *ApidefAPIDefinition) GetIsOasOk() (*bool, bool)`

GetIsOasOk returns a tuple with the IsOas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOas

`func (o *ApidefAPIDefinition) SetIsOas(v bool)`

SetIsOas sets IsOas field to given value.

### HasIsOas

`func (o *ApidefAPIDefinition) HasIsOas() bool`

HasIsOas returns a boolean if a field has been set.

### GetJwtClientBaseField

`func (o *ApidefAPIDefinition) GetJwtClientBaseField() string`

GetJwtClientBaseField returns the JwtClientBaseField field if non-nil, zero value otherwise.

### GetJwtClientBaseFieldOk

`func (o *ApidefAPIDefinition) GetJwtClientBaseFieldOk() (*string, bool)`

GetJwtClientBaseFieldOk returns a tuple with the JwtClientBaseField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwtClientBaseField

`func (o *ApidefAPIDefinition) SetJwtClientBaseField(v string)`

SetJwtClientBaseField sets JwtClientBaseField field to given value.

### HasJwtClientBaseField

`func (o *ApidefAPIDefinition) HasJwtClientBaseField() bool`

HasJwtClientBaseField returns a boolean if a field has been set.

### GetJwtDefaultPolicies

`func (o *ApidefAPIDefinition) GetJwtDefaultPolicies() []string`

GetJwtDefaultPolicies returns the JwtDefaultPolicies field if non-nil, zero value otherwise.

### GetJwtDefaultPoliciesOk

`func (o *ApidefAPIDefinition) GetJwtDefaultPoliciesOk() (*[]string, bool)`

GetJwtDefaultPoliciesOk returns a tuple with the JwtDefaultPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwtDefaultPolicies

`func (o *ApidefAPIDefinition) SetJwtDefaultPolicies(v []string)`

SetJwtDefaultPolicies sets JwtDefaultPolicies field to given value.

### HasJwtDefaultPolicies

`func (o *ApidefAPIDefinition) HasJwtDefaultPolicies() bool`

HasJwtDefaultPolicies returns a boolean if a field has been set.

### SetJwtDefaultPoliciesNil

`func (o *ApidefAPIDefinition) SetJwtDefaultPoliciesNil(b bool)`

 SetJwtDefaultPoliciesNil sets the value for JwtDefaultPolicies to be an explicit nil

### UnsetJwtDefaultPolicies
`func (o *ApidefAPIDefinition) UnsetJwtDefaultPolicies()`

UnsetJwtDefaultPolicies ensures that no value is present for JwtDefaultPolicies, not even an explicit nil
### GetJwtExpiresAtValidationSkew

`func (o *ApidefAPIDefinition) GetJwtExpiresAtValidationSkew() int32`

GetJwtExpiresAtValidationSkew returns the JwtExpiresAtValidationSkew field if non-nil, zero value otherwise.

### GetJwtExpiresAtValidationSkewOk

`func (o *ApidefAPIDefinition) GetJwtExpiresAtValidationSkewOk() (*int32, bool)`

GetJwtExpiresAtValidationSkewOk returns a tuple with the JwtExpiresAtValidationSkew field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwtExpiresAtValidationSkew

`func (o *ApidefAPIDefinition) SetJwtExpiresAtValidationSkew(v int32)`

SetJwtExpiresAtValidationSkew sets JwtExpiresAtValidationSkew field to given value.

### HasJwtExpiresAtValidationSkew

`func (o *ApidefAPIDefinition) HasJwtExpiresAtValidationSkew() bool`

HasJwtExpiresAtValidationSkew returns a boolean if a field has been set.

### GetJwtIdentityBaseField

`func (o *ApidefAPIDefinition) GetJwtIdentityBaseField() string`

GetJwtIdentityBaseField returns the JwtIdentityBaseField field if non-nil, zero value otherwise.

### GetJwtIdentityBaseFieldOk

`func (o *ApidefAPIDefinition) GetJwtIdentityBaseFieldOk() (*string, bool)`

GetJwtIdentityBaseFieldOk returns a tuple with the JwtIdentityBaseField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwtIdentityBaseField

`func (o *ApidefAPIDefinition) SetJwtIdentityBaseField(v string)`

SetJwtIdentityBaseField sets JwtIdentityBaseField field to given value.

### HasJwtIdentityBaseField

`func (o *ApidefAPIDefinition) HasJwtIdentityBaseField() bool`

HasJwtIdentityBaseField returns a boolean if a field has been set.

### GetJwtIssuedAtValidationSkew

`func (o *ApidefAPIDefinition) GetJwtIssuedAtValidationSkew() int32`

GetJwtIssuedAtValidationSkew returns the JwtIssuedAtValidationSkew field if non-nil, zero value otherwise.

### GetJwtIssuedAtValidationSkewOk

`func (o *ApidefAPIDefinition) GetJwtIssuedAtValidationSkewOk() (*int32, bool)`

GetJwtIssuedAtValidationSkewOk returns a tuple with the JwtIssuedAtValidationSkew field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwtIssuedAtValidationSkew

`func (o *ApidefAPIDefinition) SetJwtIssuedAtValidationSkew(v int32)`

SetJwtIssuedAtValidationSkew sets JwtIssuedAtValidationSkew field to given value.

### HasJwtIssuedAtValidationSkew

`func (o *ApidefAPIDefinition) HasJwtIssuedAtValidationSkew() bool`

HasJwtIssuedAtValidationSkew returns a boolean if a field has been set.

### GetJwtNotBeforeValidationSkew

`func (o *ApidefAPIDefinition) GetJwtNotBeforeValidationSkew() int32`

GetJwtNotBeforeValidationSkew returns the JwtNotBeforeValidationSkew field if non-nil, zero value otherwise.

### GetJwtNotBeforeValidationSkewOk

`func (o *ApidefAPIDefinition) GetJwtNotBeforeValidationSkewOk() (*int32, bool)`

GetJwtNotBeforeValidationSkewOk returns a tuple with the JwtNotBeforeValidationSkew field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwtNotBeforeValidationSkew

`func (o *ApidefAPIDefinition) SetJwtNotBeforeValidationSkew(v int32)`

SetJwtNotBeforeValidationSkew sets JwtNotBeforeValidationSkew field to given value.

### HasJwtNotBeforeValidationSkew

`func (o *ApidefAPIDefinition) HasJwtNotBeforeValidationSkew() bool`

HasJwtNotBeforeValidationSkew returns a boolean if a field has been set.

### GetJwtPolicyFieldName

`func (o *ApidefAPIDefinition) GetJwtPolicyFieldName() string`

GetJwtPolicyFieldName returns the JwtPolicyFieldName field if non-nil, zero value otherwise.

### GetJwtPolicyFieldNameOk

`func (o *ApidefAPIDefinition) GetJwtPolicyFieldNameOk() (*string, bool)`

GetJwtPolicyFieldNameOk returns a tuple with the JwtPolicyFieldName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwtPolicyFieldName

`func (o *ApidefAPIDefinition) SetJwtPolicyFieldName(v string)`

SetJwtPolicyFieldName sets JwtPolicyFieldName field to given value.

### HasJwtPolicyFieldName

`func (o *ApidefAPIDefinition) HasJwtPolicyFieldName() bool`

HasJwtPolicyFieldName returns a boolean if a field has been set.

### GetJwtScopeClaimName

`func (o *ApidefAPIDefinition) GetJwtScopeClaimName() string`

GetJwtScopeClaimName returns the JwtScopeClaimName field if non-nil, zero value otherwise.

### GetJwtScopeClaimNameOk

`func (o *ApidefAPIDefinition) GetJwtScopeClaimNameOk() (*string, bool)`

GetJwtScopeClaimNameOk returns a tuple with the JwtScopeClaimName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwtScopeClaimName

`func (o *ApidefAPIDefinition) SetJwtScopeClaimName(v string)`

SetJwtScopeClaimName sets JwtScopeClaimName field to given value.

### HasJwtScopeClaimName

`func (o *ApidefAPIDefinition) HasJwtScopeClaimName() bool`

HasJwtScopeClaimName returns a boolean if a field has been set.

### GetJwtScopeToPolicyMapping

`func (o *ApidefAPIDefinition) GetJwtScopeToPolicyMapping() map[string]string`

GetJwtScopeToPolicyMapping returns the JwtScopeToPolicyMapping field if non-nil, zero value otherwise.

### GetJwtScopeToPolicyMappingOk

`func (o *ApidefAPIDefinition) GetJwtScopeToPolicyMappingOk() (*map[string]string, bool)`

GetJwtScopeToPolicyMappingOk returns a tuple with the JwtScopeToPolicyMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwtScopeToPolicyMapping

`func (o *ApidefAPIDefinition) SetJwtScopeToPolicyMapping(v map[string]string)`

SetJwtScopeToPolicyMapping sets JwtScopeToPolicyMapping field to given value.

### HasJwtScopeToPolicyMapping

`func (o *ApidefAPIDefinition) HasJwtScopeToPolicyMapping() bool`

HasJwtScopeToPolicyMapping returns a boolean if a field has been set.

### SetJwtScopeToPolicyMappingNil

`func (o *ApidefAPIDefinition) SetJwtScopeToPolicyMappingNil(b bool)`

 SetJwtScopeToPolicyMappingNil sets the value for JwtScopeToPolicyMapping to be an explicit nil

### UnsetJwtScopeToPolicyMapping
`func (o *ApidefAPIDefinition) UnsetJwtScopeToPolicyMapping()`

UnsetJwtScopeToPolicyMapping ensures that no value is present for JwtScopeToPolicyMapping, not even an explicit nil
### GetJwtSigningMethod

`func (o *ApidefAPIDefinition) GetJwtSigningMethod() string`

GetJwtSigningMethod returns the JwtSigningMethod field if non-nil, zero value otherwise.

### GetJwtSigningMethodOk

`func (o *ApidefAPIDefinition) GetJwtSigningMethodOk() (*string, bool)`

GetJwtSigningMethodOk returns a tuple with the JwtSigningMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwtSigningMethod

`func (o *ApidefAPIDefinition) SetJwtSigningMethod(v string)`

SetJwtSigningMethod sets JwtSigningMethod field to given value.

### HasJwtSigningMethod

`func (o *ApidefAPIDefinition) HasJwtSigningMethod() bool`

HasJwtSigningMethod returns a boolean if a field has been set.

### GetJwtSkipKid

`func (o *ApidefAPIDefinition) GetJwtSkipKid() bool`

GetJwtSkipKid returns the JwtSkipKid field if non-nil, zero value otherwise.

### GetJwtSkipKidOk

`func (o *ApidefAPIDefinition) GetJwtSkipKidOk() (*bool, bool)`

GetJwtSkipKidOk returns a tuple with the JwtSkipKid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwtSkipKid

`func (o *ApidefAPIDefinition) SetJwtSkipKid(v bool)`

SetJwtSkipKid sets JwtSkipKid field to given value.

### HasJwtSkipKid

`func (o *ApidefAPIDefinition) HasJwtSkipKid() bool`

HasJwtSkipKid returns a boolean if a field has been set.

### GetJwtSource

`func (o *ApidefAPIDefinition) GetJwtSource() string`

GetJwtSource returns the JwtSource field if non-nil, zero value otherwise.

### GetJwtSourceOk

`func (o *ApidefAPIDefinition) GetJwtSourceOk() (*string, bool)`

GetJwtSourceOk returns a tuple with the JwtSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwtSource

`func (o *ApidefAPIDefinition) SetJwtSource(v string)`

SetJwtSource sets JwtSource field to given value.

### HasJwtSource

`func (o *ApidefAPIDefinition) HasJwtSource() bool`

HasJwtSource returns a boolean if a field has been set.

### GetListenPort

`func (o *ApidefAPIDefinition) GetListenPort() int32`

GetListenPort returns the ListenPort field if non-nil, zero value otherwise.

### GetListenPortOk

`func (o *ApidefAPIDefinition) GetListenPortOk() (*int32, bool)`

GetListenPortOk returns a tuple with the ListenPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListenPort

`func (o *ApidefAPIDefinition) SetListenPort(v int32)`

SetListenPort sets ListenPort field to given value.

### HasListenPort

`func (o *ApidefAPIDefinition) HasListenPort() bool`

HasListenPort returns a boolean if a field has been set.

### GetName

`func (o *ApidefAPIDefinition) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApidefAPIDefinition) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApidefAPIDefinition) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ApidefAPIDefinition) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNotifications

`func (o *ApidefAPIDefinition) GetNotifications() ApidefNotificationsManager`

GetNotifications returns the Notifications field if non-nil, zero value otherwise.

### GetNotificationsOk

`func (o *ApidefAPIDefinition) GetNotificationsOk() (*ApidefNotificationsManager, bool)`

GetNotificationsOk returns a tuple with the Notifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifications

`func (o *ApidefAPIDefinition) SetNotifications(v ApidefNotificationsManager)`

SetNotifications sets Notifications field to given value.

### HasNotifications

`func (o *ApidefAPIDefinition) HasNotifications() bool`

HasNotifications returns a boolean if a field has been set.

### GetOauthMeta

`func (o *ApidefAPIDefinition) GetOauthMeta() ApidefAPIDefinitionOauthMeta`

GetOauthMeta returns the OauthMeta field if non-nil, zero value otherwise.

### GetOauthMetaOk

`func (o *ApidefAPIDefinition) GetOauthMetaOk() (*ApidefAPIDefinitionOauthMeta, bool)`

GetOauthMetaOk returns a tuple with the OauthMeta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauthMeta

`func (o *ApidefAPIDefinition) SetOauthMeta(v ApidefAPIDefinitionOauthMeta)`

SetOauthMeta sets OauthMeta field to given value.

### HasOauthMeta

`func (o *ApidefAPIDefinition) HasOauthMeta() bool`

HasOauthMeta returns a boolean if a field has been set.

### GetOpenidOptions

`func (o *ApidefAPIDefinition) GetOpenidOptions() ApidefOpenIDOptions`

GetOpenidOptions returns the OpenidOptions field if non-nil, zero value otherwise.

### GetOpenidOptionsOk

`func (o *ApidefAPIDefinition) GetOpenidOptionsOk() (*ApidefOpenIDOptions, bool)`

GetOpenidOptionsOk returns a tuple with the OpenidOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenidOptions

`func (o *ApidefAPIDefinition) SetOpenidOptions(v ApidefOpenIDOptions)`

SetOpenidOptions sets OpenidOptions field to given value.

### HasOpenidOptions

`func (o *ApidefAPIDefinition) HasOpenidOptions() bool`

HasOpenidOptions returns a boolean if a field has been set.

### GetOrgId

`func (o *ApidefAPIDefinition) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *ApidefAPIDefinition) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *ApidefAPIDefinition) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *ApidefAPIDefinition) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetPinnedPublicKeys

`func (o *ApidefAPIDefinition) GetPinnedPublicKeys() map[string]string`

GetPinnedPublicKeys returns the PinnedPublicKeys field if non-nil, zero value otherwise.

### GetPinnedPublicKeysOk

`func (o *ApidefAPIDefinition) GetPinnedPublicKeysOk() (*map[string]string, bool)`

GetPinnedPublicKeysOk returns a tuple with the PinnedPublicKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPinnedPublicKeys

`func (o *ApidefAPIDefinition) SetPinnedPublicKeys(v map[string]string)`

SetPinnedPublicKeys sets PinnedPublicKeys field to given value.

### HasPinnedPublicKeys

`func (o *ApidefAPIDefinition) HasPinnedPublicKeys() bool`

HasPinnedPublicKeys returns a boolean if a field has been set.

### SetPinnedPublicKeysNil

`func (o *ApidefAPIDefinition) SetPinnedPublicKeysNil(b bool)`

 SetPinnedPublicKeysNil sets the value for PinnedPublicKeys to be an explicit nil

### UnsetPinnedPublicKeys
`func (o *ApidefAPIDefinition) UnsetPinnedPublicKeys()`

UnsetPinnedPublicKeys ensures that no value is present for PinnedPublicKeys, not even an explicit nil
### GetProtocol

`func (o *ApidefAPIDefinition) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *ApidefAPIDefinition) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *ApidefAPIDefinition) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *ApidefAPIDefinition) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetProxy

`func (o *ApidefAPIDefinition) GetProxy() ApidefProxyConfig`

GetProxy returns the Proxy field if non-nil, zero value otherwise.

### GetProxyOk

`func (o *ApidefAPIDefinition) GetProxyOk() (*ApidefProxyConfig, bool)`

GetProxyOk returns a tuple with the Proxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxy

`func (o *ApidefAPIDefinition) SetProxy(v ApidefProxyConfig)`

SetProxy sets Proxy field to given value.

### HasProxy

`func (o *ApidefAPIDefinition) HasProxy() bool`

HasProxy returns a boolean if a field has been set.

### GetRequestSigning

`func (o *ApidefAPIDefinition) GetRequestSigning() ApidefRequestSigningMeta`

GetRequestSigning returns the RequestSigning field if non-nil, zero value otherwise.

### GetRequestSigningOk

`func (o *ApidefAPIDefinition) GetRequestSigningOk() (*ApidefRequestSigningMeta, bool)`

GetRequestSigningOk returns a tuple with the RequestSigning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestSigning

`func (o *ApidefAPIDefinition) SetRequestSigning(v ApidefRequestSigningMeta)`

SetRequestSigning sets RequestSigning field to given value.

### HasRequestSigning

`func (o *ApidefAPIDefinition) HasRequestSigning() bool`

HasRequestSigning returns a boolean if a field has been set.

### GetResponseProcessors

`func (o *ApidefAPIDefinition) GetResponseProcessors() []ApidefResponseProcessor`

GetResponseProcessors returns the ResponseProcessors field if non-nil, zero value otherwise.

### GetResponseProcessorsOk

`func (o *ApidefAPIDefinition) GetResponseProcessorsOk() (*[]ApidefResponseProcessor, bool)`

GetResponseProcessorsOk returns a tuple with the ResponseProcessors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseProcessors

`func (o *ApidefAPIDefinition) SetResponseProcessors(v []ApidefResponseProcessor)`

SetResponseProcessors sets ResponseProcessors field to given value.

### HasResponseProcessors

`func (o *ApidefAPIDefinition) HasResponseProcessors() bool`

HasResponseProcessors returns a boolean if a field has been set.

### SetResponseProcessorsNil

`func (o *ApidefAPIDefinition) SetResponseProcessorsNil(b bool)`

 SetResponseProcessorsNil sets the value for ResponseProcessors to be an explicit nil

### UnsetResponseProcessors
`func (o *ApidefAPIDefinition) UnsetResponseProcessors()`

UnsetResponseProcessors ensures that no value is present for ResponseProcessors, not even an explicit nil
### GetScopes

`func (o *ApidefAPIDefinition) GetScopes() ApidefScopes`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *ApidefAPIDefinition) GetScopesOk() (*ApidefScopes, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *ApidefAPIDefinition) SetScopes(v ApidefScopes)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *ApidefAPIDefinition) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### GetSessionLifetime

`func (o *ApidefAPIDefinition) GetSessionLifetime() int32`

GetSessionLifetime returns the SessionLifetime field if non-nil, zero value otherwise.

### GetSessionLifetimeOk

`func (o *ApidefAPIDefinition) GetSessionLifetimeOk() (*int32, bool)`

GetSessionLifetimeOk returns a tuple with the SessionLifetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionLifetime

`func (o *ApidefAPIDefinition) SetSessionLifetime(v int32)`

SetSessionLifetime sets SessionLifetime field to given value.

### HasSessionLifetime

`func (o *ApidefAPIDefinition) HasSessionLifetime() bool`

HasSessionLifetime returns a boolean if a field has been set.

### GetSessionLifetimeRespectsKeyExpiration

`func (o *ApidefAPIDefinition) GetSessionLifetimeRespectsKeyExpiration() bool`

GetSessionLifetimeRespectsKeyExpiration returns the SessionLifetimeRespectsKeyExpiration field if non-nil, zero value otherwise.

### GetSessionLifetimeRespectsKeyExpirationOk

`func (o *ApidefAPIDefinition) GetSessionLifetimeRespectsKeyExpirationOk() (*bool, bool)`

GetSessionLifetimeRespectsKeyExpirationOk returns a tuple with the SessionLifetimeRespectsKeyExpiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionLifetimeRespectsKeyExpiration

`func (o *ApidefAPIDefinition) SetSessionLifetimeRespectsKeyExpiration(v bool)`

SetSessionLifetimeRespectsKeyExpiration sets SessionLifetimeRespectsKeyExpiration field to given value.

### HasSessionLifetimeRespectsKeyExpiration

`func (o *ApidefAPIDefinition) HasSessionLifetimeRespectsKeyExpiration() bool`

HasSessionLifetimeRespectsKeyExpiration returns a boolean if a field has been set.

### GetSessionProvider

`func (o *ApidefAPIDefinition) GetSessionProvider() ApidefSessionProviderMeta`

GetSessionProvider returns the SessionProvider field if non-nil, zero value otherwise.

### GetSessionProviderOk

`func (o *ApidefAPIDefinition) GetSessionProviderOk() (*ApidefSessionProviderMeta, bool)`

GetSessionProviderOk returns a tuple with the SessionProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionProvider

`func (o *ApidefAPIDefinition) SetSessionProvider(v ApidefSessionProviderMeta)`

SetSessionProvider sets SessionProvider field to given value.

### HasSessionProvider

`func (o *ApidefAPIDefinition) HasSessionProvider() bool`

HasSessionProvider returns a boolean if a field has been set.

### GetSlug

`func (o *ApidefAPIDefinition) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *ApidefAPIDefinition) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *ApidefAPIDefinition) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *ApidefAPIDefinition) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetStripAuthData

`func (o *ApidefAPIDefinition) GetStripAuthData() bool`

GetStripAuthData returns the StripAuthData field if non-nil, zero value otherwise.

### GetStripAuthDataOk

`func (o *ApidefAPIDefinition) GetStripAuthDataOk() (*bool, bool)`

GetStripAuthDataOk returns a tuple with the StripAuthData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStripAuthData

`func (o *ApidefAPIDefinition) SetStripAuthData(v bool)`

SetStripAuthData sets StripAuthData field to given value.

### HasStripAuthData

`func (o *ApidefAPIDefinition) HasStripAuthData() bool`

HasStripAuthData returns a boolean if a field has been set.

### GetTagHeaders

`func (o *ApidefAPIDefinition) GetTagHeaders() []string`

GetTagHeaders returns the TagHeaders field if non-nil, zero value otherwise.

### GetTagHeadersOk

`func (o *ApidefAPIDefinition) GetTagHeadersOk() (*[]string, bool)`

GetTagHeadersOk returns a tuple with the TagHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagHeaders

`func (o *ApidefAPIDefinition) SetTagHeaders(v []string)`

SetTagHeaders sets TagHeaders field to given value.

### HasTagHeaders

`func (o *ApidefAPIDefinition) HasTagHeaders() bool`

HasTagHeaders returns a boolean if a field has been set.

### SetTagHeadersNil

`func (o *ApidefAPIDefinition) SetTagHeadersNil(b bool)`

 SetTagHeadersNil sets the value for TagHeaders to be an explicit nil

### UnsetTagHeaders
`func (o *ApidefAPIDefinition) UnsetTagHeaders()`

UnsetTagHeaders ensures that no value is present for TagHeaders, not even an explicit nil
### GetTags

`func (o *ApidefAPIDefinition) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ApidefAPIDefinition) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ApidefAPIDefinition) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ApidefAPIDefinition) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *ApidefAPIDefinition) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *ApidefAPIDefinition) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetTagsDisabled

`func (o *ApidefAPIDefinition) GetTagsDisabled() bool`

GetTagsDisabled returns the TagsDisabled field if non-nil, zero value otherwise.

### GetTagsDisabledOk

`func (o *ApidefAPIDefinition) GetTagsDisabledOk() (*bool, bool)`

GetTagsDisabledOk returns a tuple with the TagsDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagsDisabled

`func (o *ApidefAPIDefinition) SetTagsDisabled(v bool)`

SetTagsDisabled sets TagsDisabled field to given value.

### HasTagsDisabled

`func (o *ApidefAPIDefinition) HasTagsDisabled() bool`

HasTagsDisabled returns a boolean if a field has been set.

### GetUpstreamCertificates

`func (o *ApidefAPIDefinition) GetUpstreamCertificates() map[string]string`

GetUpstreamCertificates returns the UpstreamCertificates field if non-nil, zero value otherwise.

### GetUpstreamCertificatesOk

`func (o *ApidefAPIDefinition) GetUpstreamCertificatesOk() (*map[string]string, bool)`

GetUpstreamCertificatesOk returns a tuple with the UpstreamCertificates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpstreamCertificates

`func (o *ApidefAPIDefinition) SetUpstreamCertificates(v map[string]string)`

SetUpstreamCertificates sets UpstreamCertificates field to given value.

### HasUpstreamCertificates

`func (o *ApidefAPIDefinition) HasUpstreamCertificates() bool`

HasUpstreamCertificates returns a boolean if a field has been set.

### SetUpstreamCertificatesNil

`func (o *ApidefAPIDefinition) SetUpstreamCertificatesNil(b bool)`

 SetUpstreamCertificatesNil sets the value for UpstreamCertificates to be an explicit nil

### UnsetUpstreamCertificates
`func (o *ApidefAPIDefinition) UnsetUpstreamCertificates()`

UnsetUpstreamCertificates ensures that no value is present for UpstreamCertificates, not even an explicit nil
### GetUpstreamCertificatesDisabled

`func (o *ApidefAPIDefinition) GetUpstreamCertificatesDisabled() bool`

GetUpstreamCertificatesDisabled returns the UpstreamCertificatesDisabled field if non-nil, zero value otherwise.

### GetUpstreamCertificatesDisabledOk

`func (o *ApidefAPIDefinition) GetUpstreamCertificatesDisabledOk() (*bool, bool)`

GetUpstreamCertificatesDisabledOk returns a tuple with the UpstreamCertificatesDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpstreamCertificatesDisabled

`func (o *ApidefAPIDefinition) SetUpstreamCertificatesDisabled(v bool)`

SetUpstreamCertificatesDisabled sets UpstreamCertificatesDisabled field to given value.

### HasUpstreamCertificatesDisabled

`func (o *ApidefAPIDefinition) HasUpstreamCertificatesDisabled() bool`

HasUpstreamCertificatesDisabled returns a boolean if a field has been set.

### GetUptimeTests

`func (o *ApidefAPIDefinition) GetUptimeTests() ApidefUptimeTests`

GetUptimeTests returns the UptimeTests field if non-nil, zero value otherwise.

### GetUptimeTestsOk

`func (o *ApidefAPIDefinition) GetUptimeTestsOk() (*ApidefUptimeTests, bool)`

GetUptimeTestsOk returns a tuple with the UptimeTests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUptimeTests

`func (o *ApidefAPIDefinition) SetUptimeTests(v ApidefUptimeTests)`

SetUptimeTests sets UptimeTests field to given value.

### HasUptimeTests

`func (o *ApidefAPIDefinition) HasUptimeTests() bool`

HasUptimeTests returns a boolean if a field has been set.

### GetUseBasicAuth

`func (o *ApidefAPIDefinition) GetUseBasicAuth() bool`

GetUseBasicAuth returns the UseBasicAuth field if non-nil, zero value otherwise.

### GetUseBasicAuthOk

`func (o *ApidefAPIDefinition) GetUseBasicAuthOk() (*bool, bool)`

GetUseBasicAuthOk returns a tuple with the UseBasicAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseBasicAuth

`func (o *ApidefAPIDefinition) SetUseBasicAuth(v bool)`

SetUseBasicAuth sets UseBasicAuth field to given value.

### HasUseBasicAuth

`func (o *ApidefAPIDefinition) HasUseBasicAuth() bool`

HasUseBasicAuth returns a boolean if a field has been set.

### GetUseGoPluginAuth

`func (o *ApidefAPIDefinition) GetUseGoPluginAuth() bool`

GetUseGoPluginAuth returns the UseGoPluginAuth field if non-nil, zero value otherwise.

### GetUseGoPluginAuthOk

`func (o *ApidefAPIDefinition) GetUseGoPluginAuthOk() (*bool, bool)`

GetUseGoPluginAuthOk returns a tuple with the UseGoPluginAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseGoPluginAuth

`func (o *ApidefAPIDefinition) SetUseGoPluginAuth(v bool)`

SetUseGoPluginAuth sets UseGoPluginAuth field to given value.

### HasUseGoPluginAuth

`func (o *ApidefAPIDefinition) HasUseGoPluginAuth() bool`

HasUseGoPluginAuth returns a boolean if a field has been set.

### GetUseKeyless

`func (o *ApidefAPIDefinition) GetUseKeyless() bool`

GetUseKeyless returns the UseKeyless field if non-nil, zero value otherwise.

### GetUseKeylessOk

`func (o *ApidefAPIDefinition) GetUseKeylessOk() (*bool, bool)`

GetUseKeylessOk returns a tuple with the UseKeyless field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseKeyless

`func (o *ApidefAPIDefinition) SetUseKeyless(v bool)`

SetUseKeyless sets UseKeyless field to given value.

### HasUseKeyless

`func (o *ApidefAPIDefinition) HasUseKeyless() bool`

HasUseKeyless returns a boolean if a field has been set.

### GetUseMutualTlsAuth

`func (o *ApidefAPIDefinition) GetUseMutualTlsAuth() bool`

GetUseMutualTlsAuth returns the UseMutualTlsAuth field if non-nil, zero value otherwise.

### GetUseMutualTlsAuthOk

`func (o *ApidefAPIDefinition) GetUseMutualTlsAuthOk() (*bool, bool)`

GetUseMutualTlsAuthOk returns a tuple with the UseMutualTlsAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseMutualTlsAuth

`func (o *ApidefAPIDefinition) SetUseMutualTlsAuth(v bool)`

SetUseMutualTlsAuth sets UseMutualTlsAuth field to given value.

### HasUseMutualTlsAuth

`func (o *ApidefAPIDefinition) HasUseMutualTlsAuth() bool`

HasUseMutualTlsAuth returns a boolean if a field has been set.

### GetUseOauth2

`func (o *ApidefAPIDefinition) GetUseOauth2() bool`

GetUseOauth2 returns the UseOauth2 field if non-nil, zero value otherwise.

### GetUseOauth2Ok

`func (o *ApidefAPIDefinition) GetUseOauth2Ok() (*bool, bool)`

GetUseOauth2Ok returns a tuple with the UseOauth2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseOauth2

`func (o *ApidefAPIDefinition) SetUseOauth2(v bool)`

SetUseOauth2 sets UseOauth2 field to given value.

### HasUseOauth2

`func (o *ApidefAPIDefinition) HasUseOauth2() bool`

HasUseOauth2 returns a boolean if a field has been set.

### GetUseOpenid

`func (o *ApidefAPIDefinition) GetUseOpenid() bool`

GetUseOpenid returns the UseOpenid field if non-nil, zero value otherwise.

### GetUseOpenidOk

`func (o *ApidefAPIDefinition) GetUseOpenidOk() (*bool, bool)`

GetUseOpenidOk returns a tuple with the UseOpenid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseOpenid

`func (o *ApidefAPIDefinition) SetUseOpenid(v bool)`

SetUseOpenid sets UseOpenid field to given value.

### HasUseOpenid

`func (o *ApidefAPIDefinition) HasUseOpenid() bool`

HasUseOpenid returns a boolean if a field has been set.

### GetUseStandardAuth

`func (o *ApidefAPIDefinition) GetUseStandardAuth() bool`

GetUseStandardAuth returns the UseStandardAuth field if non-nil, zero value otherwise.

### GetUseStandardAuthOk

`func (o *ApidefAPIDefinition) GetUseStandardAuthOk() (*bool, bool)`

GetUseStandardAuthOk returns a tuple with the UseStandardAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseStandardAuth

`func (o *ApidefAPIDefinition) SetUseStandardAuth(v bool)`

SetUseStandardAuth sets UseStandardAuth field to given value.

### HasUseStandardAuth

`func (o *ApidefAPIDefinition) HasUseStandardAuth() bool`

HasUseStandardAuth returns a boolean if a field has been set.

### GetVersionData

`func (o *ApidefAPIDefinition) GetVersionData() ApidefVersionData`

GetVersionData returns the VersionData field if non-nil, zero value otherwise.

### GetVersionDataOk

`func (o *ApidefAPIDefinition) GetVersionDataOk() (*ApidefVersionData, bool)`

GetVersionDataOk returns a tuple with the VersionData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionData

`func (o *ApidefAPIDefinition) SetVersionData(v ApidefVersionData)`

SetVersionData sets VersionData field to given value.

### HasVersionData

`func (o *ApidefAPIDefinition) HasVersionData() bool`

HasVersionData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


