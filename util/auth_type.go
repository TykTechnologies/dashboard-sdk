package util

import "github.com/TykTechnologies/dashboard-sdk/pkg/dashboard"

type AuthType string

const (
	AuthTypeMultiAuth   AuthType = "multiAuth"
	AuthTypeKeyless     AuthType = "keyless"
	AuthTypeBasic       AuthType = "basic"
	AuthTypeBasicLegacy AuthType = "ba"
	AuthTypeHmac        AuthType = "hmac"
	AuthTypeJWT         AuthType = "jwt"
	AuthTypeOAuth       AuthType = "oauth"
	AuthTypeOpenID      AuthType = "openid"
	AuthTypeMutualTLS   AuthType = "mutualTLS"
	AuthTypeAuthToken   AuthType = "authToken"
	AuthTypeCustom      AuthType = "custom"
	AuthTypeOther       AuthType = "other"
)

func GetAuthType(apiDefinition *dashboard.APIDefinition) string {
	if apiDefinition.GetUseKeyless() {
		return string(AuthTypeKeyless)
	}

	count := 0
	authType := AuthTypeOther

	if apiDefinition.GetUseBasicAuth() {
		authType = AuthTypeBasic
		count++
	}

	if apiDefinition.GetEnableSignatureChecking() {
		authType = AuthTypeHmac
		count++
	}

	if apiDefinition.GetEnableJwt() {
		authType = AuthTypeJWT
		count++
	}

	if apiDefinition.GetUseOauth2() {
		authType = AuthTypeOAuth
		count++
	}

	if apiDefinition.GetUseOpenid() {
		authType = AuthTypeOpenID
		count++
	}

	if apiDefinition.GetCustomPluginAuthEnabled() || apiDefinition.GetUseGoPluginAuth() || apiDefinition.GetEnableCoprocessAuth() {
		authType = AuthTypeCustom
		count++
	}

	if apiDefinition.GetUseMutualTlsAuth() {
		authType = AuthTypeMutualTLS
		count++
	}

	if apiDefinition.GetUseStandardAuth() {
		authType = AuthTypeAuthToken
		count++
	}

	if count > 1 {
		return string(AuthTypeMultiAuth)
	}

	return string(authType)
}

func AvailableAuthTypes(apiDefinition dashboard.APIDefinition) (out []AuthType) {
	if apiDefinition.GetUseKeyless() {
		out = append(out, AuthTypeKeyless)
	}
	if apiDefinition.GetUseBasicAuth() {
		out = append(out, AuthTypeBasic)
	}
	if apiDefinition.GetEnableSignatureChecking() {
		out = append(out, AuthTypeHmac)
	}
	if apiDefinition.GetEnableJwt() {
		out = append(out, AuthTypeJWT)
	}
	if apiDefinition.GetUseOauth2() {
		out = append(out, AuthTypeOAuth)
	}
	if apiDefinition.GetUseOpenid() {
		out = append(out, AuthTypeOpenID)
	}
	if apiDefinition.GetUseGoPluginAuth() {
		out = append(out, AuthTypeCustom)
	}
	if apiDefinition.GetUseMutualTlsAuth() {
		out = append(out, AuthTypeMutualTLS)
	}
	if apiDefinition.GetUseStandardAuth() {
		out = append(out, AuthTypeAuthToken)
	}
	if apiDefinition.GetEnableCoprocessAuth() {
		out = append(out, AuthTypeCustom)
	}
	return
}
