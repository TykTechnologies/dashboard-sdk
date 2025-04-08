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


func  GetAuthType(a *dashboard.APIDefinition) string {
	if a.GetUseKeyless()  {
		return string(AuthTypeKeyless)
	}

	count := 0
	authType := AuthTypeOther

	if a.GetUseBasicAuth() {
		authType = AuthTypeBasic
		count++
	}

	if a.GetEnableSignatureChecking() {
		authType = AuthTypeHmac
		count++
	}

	if a.GetEnableJwt() {
		authType = AuthTypeJWT
		count++
	}

	if a.GetUseOauth2(){
		authType = AuthTypeOAuth
		count++
	}

	if a.GetUseOpenid() {
		authType = AuthTypeOpenID
		count++
	}

	if a.GetCustomPluginAuthEnabled() || a.GetUseGoPluginAuth()  || a.GetEnableCoprocessAuth()   {
		authType = AuthTypeCustom
		count++
	}

	if a.GetUseMutualTlsAuth()  {
		authType = AuthTypeMutualTLS
		count++
	}

	if a.GetUseStandardAuth() {
		authType = AuthTypeAuthToken
		count++
	}

	if count > 1 {
		return string(AuthTypeMultiAuth)
	}

	return string(authType)
}


func AvailableAuthTypes(a dashboard.APIDefinition) (out []AuthType) {
	if a.GetUseKeyless() {
		out = append(out, AuthTypeKeyless)
	}
	if a.GetUseBasicAuth() {
		out = append(out, AuthTypeBasic)
	}
	if a.GetEnableSignatureChecking() {
		out = append(out, AuthTypeHmac)
	}
	if a.GetEnableJwt() {
		out = append(out, AuthTypeJWT)
	}
	if a.GetUseOauth2()  {
		out = append(out, AuthTypeOAuth)
	}
	if  a.GetUseOpenid(){
		out = append(out, AuthTypeOpenID)
	}
	if a.GetUseGoPluginAuth() {
		out = append(out, AuthTypeCustom)
	}
	if a.GetUseMutualTlsAuth() {
		out = append(out, AuthTypeMutualTLS)
	}
	if a.GetUseStandardAuth() {
		out = append(out, AuthTypeAuthToken)
	}
	if a.GetEnableCoprocessAuth(){
		out = append(out, AuthTypeCustom)
	}
	return
}
