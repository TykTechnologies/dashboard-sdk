/*
 * Tyk Dashboard API
 *
 * ## <a name=\"introduction\"></a> Introduction  The Tyk Dashboard API offers granular, programmatic access to a centralised database of resources that your Tyk nodes can pull from. This API has a dynamic user administrative structure which means the secret key that is used to communicate with your Tyk nodes can be kept secret and access to the wider management functions can be handled on a user-by-user and organisation-by-organisation basis.  A common question around using a database-backed configuration is how to programatically add API definitions to your Tyk nodes, the Dashboard API allows much more fine-grained, secure and multi-user access to your Tyk cluster, and should be used to manage a database-backed Tyk node.  The Tyk Dashboard API works seamlessly with the Tyk Dashboard (and the two come bundled together).  ## <a name=\"security-hierarchy\"></a> Security Hierarchy  The Dashboard API provides a more structured security layer to managing Tyk nodes.  ### Organisations, APIs and Users  With the Dashboard API and a database-backed Tyk setup, (and to an extent with file-based API setups - if diligence is used in naming an creating definitions), the following security model is applied to the management of Upstream APIs:  * **Organisations**: All APIs are *owned* by an organisation, this is designated by the `OrgID` parameter in the API Definition. * **Users**: All users created in the Dashboard belong to an organisation (unless an exception is made for super-administrative access). * **APIs**: All APIs belong to an Organisation and only Users that belong to that organisation can see the analytics for those APIs and manage their configurations. * **API Keys**: API Keys are designated by organisation, this means an API key that has full access rights will not be allowed to access the APIs of another organisation on the same system, but can have full access to all APIs within the organisation. * **Access Rights**: Access rights are stored with the key, this enables a key to give access to multiple APIs, this is defined by the session object in the core Tyk API.  In order to use the Dashboard API, you'll need to get the `Tyk Dashboard API Access Credentials` secret from your user profile on the Dashboard UI.  The secret you set should then be sent along as a header with each Dashboard API Request in order for it to be successful:  ``` authorization: <your-secret> ```
 *
 * API version: 3.2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package dashboard

// There's a data structure that's based on this and it's used for Protocol Buffer support, make sure to update \"coprocess/proto/coprocess_session_state.proto\" and generate the bindings using: cd coprocess/proto && ./update_bindings.sh
type SessionState struct {
	Tags []string `json:"tags,omitempty"`
	AccessRights map[string]AccessDefinition `json:"access_rights,omitempty"`
	Alias string `json:"alias,omitempty"`
	Allowance float64 `json:"allowance,omitempty"`
	ApplyPolicies []string `json:"apply_policies,omitempty"`
	ApplyPolicyId string `json:"apply_policy_id,omitempty"`
	BasicAuthData *SessionStateBasicAuthData `json:"basic_auth_data,omitempty"`
	Certificate string `json:"certificate,omitempty"`
	DataExpires int64 `json:"data_expires,omitempty"`
	EnableDetailRecording bool `json:"enable_detail_recording,omitempty"`
	Expires int64 `json:"expires,omitempty"`
	HmacEnabled bool `json:"hmac_enabled,omitempty"`
	HmacString string `json:"hmac_string,omitempty"`
	IdExtractorDeadline int64 `json:"id_extractor_deadline,omitempty"`
	IsInactive bool `json:"is_inactive,omitempty"`
	JwtData *SessionStateJwtData `json:"jwt_data,omitempty"`
	LastCheck int64 `json:"last_check,omitempty"`
	LastUpdated string `json:"last_updated,omitempty"`
	MetaData map[string]interface{} `json:"meta_data,omitempty"`
	Monitor *SessionStateMonitor `json:"monitor,omitempty"`
	OauthClientId string `json:"oauth_client_id,omitempty"`
	OauthKeys map[string]string `json:"oauth_keys,omitempty"`
	OrgId string `json:"org_id,omitempty"`
	Per float64 `json:"per,omitempty"`
	QuotaMax int64 `json:"quota_max,omitempty"`
	QuotaRemaining int64 `json:"quota_remaining,omitempty"`
	QuotaRenewalRate int64 `json:"quota_renewal_rate,omitempty"`
	QuotaRenews int64 `json:"quota_renews,omitempty"`
	Rate float64 `json:"rate,omitempty"`
	SessionLifetime int64 `json:"session_lifetime,omitempty"`
	ThrottleInterval float64 `json:"throttle_interval,omitempty"`
	ThrottleRetryLimit int64 `json:"throttle_retry_limit,omitempty"`
}
